package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Extractor struct {
	fieldCount  int
	out         bytes.Buffer
	pkgName     string
	importDecl  string
	fieldPrefix string
	fieldSuffix string

	curStruct      *Struct
	heldComment    string
	extractPrivate bool

	// map structs' goName <-> capName
	goType2capTypeCache map[string]string
	capType2goType      map[string]string

	// key is goName
	srs        map[string]*Struct
	ToGoCode   map[string][]byte
	ToCapnCode map[string][]byte
	SaveCode   map[string][]byte
	LoadCode   map[string][]byte

	// key is CanonGoType(goTypeSeq)
	SliceToListCode map[string][]byte
	ListToSliceCode map[string][]byte

	compileDir *TempDir
	outDir     string
	srcFiles   []*SrcFile
	overwrite  bool

	// fields for testing zid tagging
	PubABC int `zid:"1"`
	PubYXZ int `zid:"0"`
	PubDEF int
}

func NewExtractor() *Extractor {
	return &Extractor{
		pkgName:             "testpkg",
		importDecl:          "testpkg",
		goType2capTypeCache: make(map[string]string),
		capType2goType:      make(map[string]string),

		// key is goTypeName
		ToGoCode:        make(map[string][]byte),
		ToCapnCode:      make(map[string][]byte),
		SaveCode:        make(map[string][]byte),
		LoadCode:        make(map[string][]byte),
		srs:             make(map[string]*Struct),
		compileDir:      NewTempDir(),
		srcFiles:        make([]*SrcFile, 0),
		SliceToListCode: make(map[string][]byte),
		ListToSliceCode: make(map[string][]byte),
	}
}

func (x *Extractor) Cleanup() {
	if x.compileDir != nil {
		x.compileDir.Cleanup()
	}
}

type Field struct {
	capname      string
	goCapGoName  string // Uppercased-first-letter of capname, as generated in go bindings.
	goCapGoType  string // int64 when goType is int, because capType is Int64.
	capType      string
	goName       string
	goType       string
	goTypePrefix string
	goToCapFunc  string

	goTypeSeq      []string
	capTypeSeq     []string
	goCapGoTypeSeq []string

	tagValue                   string
	isList                     bool
	zidFromTag                 int
	orderOfAppearance          int
	finalOrder                 int
	embedded                   bool
	astField                   *ast.Field
	canonGoType                string // key into SliceToListCode and ListToSliceCode
	canonGoTypeListToSliceFunc string
	canonGoTypeSliceToListFunc string
	singleCapListType          string
	baseIsIntrinsic            bool
	newListExpression          string
}

type Struct struct {
	capName              string
	goName               string
	fld                  []*Field
	longestField         int
	comment              string
	zidMap               map[int]*Field
	firstNonTextListSeen bool
	listNum              int
}

type SrcFile struct {
	filename string
	fset     *token.FileSet
	astFile  *ast.File
}

func (s *Struct) computeFinalOrder() {

	// assign Field.finalOrder to all values in s.fld, from 0..(len(s.fld)-1)
	max := len(s.fld) - 1
	// check for out of bounds requests
	for _, f := range s.fld {
		if f.zidFromTag > max {
			err := fmt.Errorf(`problem in zid tag '%s' on field '%s' in struct '%s': number '%d' is beyond the count of fields we have, largest available is %d`, f.tagValue, f.goName, s.goName, f.zidFromTag, max)
			panic(err)
		}
	}

	// only one field? already done
	if len(s.fld) == 1 {
		s.fld[0].finalOrder = 0
		return
	}
	// INVAR: no duplicate requests, and all are in bounds.

	// wipe slate clean
	for _, f := range s.fld {
		f.finalOrder = -1
	}

	final := make([]*Field, len(s.fld))

	// assign from map
	for _, v := range s.zidMap {
		v.finalOrder = v.zidFromTag
		final[v.zidFromTag] = v
	}

	appear := make([]*Field, len(s.fld))
	copy(appear, s.fld)

	sort.Sort(ByOrderOfAppearance(appear))

	//find next available slot, and fill in, in order of appearance
	write := 0

	for read := 0; read <= max; read++ {
		if appear[read].finalOrder != -1 {
			continue
		}
		// appear[read] needs an assignment

		done := advanceWrite(&write, final)
		if !done {
			// final[write] has a slot for an assignment
			final[write] = appear[read]
			final[write].finalOrder = write
		}

	}
}

// returns true if done. if false, then final[*pw] is available for writing.
func advanceWrite(pw *int, final []*Field) bool {
	for n := len(final); *pw < n; (*pw)++ {
		if final[*pw] == nil { // .finalOrder == -1 {
			return false
		}
	}
	return true
}

func NewStruct(capName, goName string) *Struct {
	return &Struct{
		capName: capName,
		goName:  goName,
		fld:     []*Field{},
		zidMap:  map[int]*Field{},
	}
}

func (x *Extractor) GenerateTranslators() {

	for _, s := range x.srs {

		x.SaveCode[s.goName] = []byte(fmt.Sprintf(`
func (s *%s) Save(w io.Writer) error {
  	seg := capn.NewBuffer(nil)
  	%sGoToCapn(seg, s)
    _, err := seg.WriteTo(w)
    return err
}
 `, s.goName, s.goName))

		x.LoadCode[s.goName] = []byte(fmt.Sprintf(`
func (s *%s) Load(r io.Reader) error {
  	capMsg, err := capn.ReadFromStream(r, nil)
  	if err != nil {
  		//panic(fmt.Errorf("capn.ReadFromStream error: %%s", err))
        return err
  	}
  	z := ReadRoot%s(capMsg)
      %sToGo(z, s)
   return nil
}
`, s.goName, s.capName, s.capName))

		x.ToGoCode[s.goName] = []byte(fmt.Sprintf(`
func %sToGo(src %s, dest *%s) *%s {
  if dest == nil {
    dest = &%s{}
  }
%s
  return dest
}
`, s.capName, s.capName, s.goName, s.goName, s.goName, x.SettersToGo(s.goName)))

		x.ToCapnCode[s.goName] = []byte(fmt.Sprintf(`
func %sGoToCapn(seg *capn.Segment, src *%s) %s {
  dest := AutoNew%s(seg)
%s
  return dest
}
`, s.goName, s.goName, s.capName, s.capName, x.SettersToCapn(s.goName)))

	}
}

func (x *Extractor) packageDot() string {
	if x.pkgName == "" || x.pkgName == "main" {
		return ""
	}
	return x.pkgName + "."
}

func (x *Extractor) SettersToGo(goName string) string {
	var buf bytes.Buffer
	myStruct := x.srs[goName]
	if myStruct == nil {
		panic(fmt.Sprintf("bad goName '%s'", goName))
	}
	VPrintf("\n\n SettersToGo running on myStruct = %#v\n", myStruct)
	i := 0
	for _, f := range myStruct.fld {
		VPrintf("\n\n SettersToGo running on myStruct.fld[%d] = %#v\n", i, f)

		n := len(f.goTypeSeq)

		if n >= 2 && f.goTypeSeq[0] == "[]" {
			x.SettersToGoListHelper(&buf, myStruct, f)
		} else {

			var isCapType bool = false
			if _, ok := x.goType2capTypeCache[f.goTypeSeq[0]]; !ok {
				if len(f.goTypeSeq) > 1 {
					if _, ok := x.goType2capTypeCache[f.goTypeSeq[1]]; ok {
						isCapType = true
					}
				}
			} else {
				isCapType = true
			}

			if isCapType {
				if f.goTypeSeq[0] == "*" {
					fmt.Fprintf(&buf, "  dest.%s = %sCapnToGo(src.%s(), nil)\n",
						f.goName, f.goType, f.goCapGoName)
				} else {
					fmt.Fprintf(&buf, "  dest.%s = *%sCapnToGo(src.%s(), nil)\n",
						f.goName, f.goType, f.goCapGoName)
				}

				continue
			}

			switch f.goType {
			case "int":
				fmt.Fprintf(&buf, "  dest.%s = int(src.%s())\n", f.goName, f.goCapGoName)

			default:
				fmt.Fprintf(&buf, "  dest.%s = src.%s()\n", f.goName, f.goCapGoName)
			}
		}
		i++
	}
	return string(buf.Bytes())
}

func (x *Extractor) SettersToGoListHelper(buf io.Writer, myStruct *Struct, f *Field) {

	VPrintf("\n in SettersToGoListHelper(): debug: field f = %#v\n", f)
	VPrintf("\n in SettersToGoListHelper(): debug: myStruct = %#v\n", myStruct)

	// special case Text / string slices
	if f.capType == "List(Text)" {
		fmt.Fprintf(buf, "  dest.%s = src.%s().ToArray()\n", f.goName, f.goCapGoName)
		return
	}
	if !myStruct.firstNonTextListSeen {
		fmt.Fprintf(buf, "\n  var n int\n")
		myStruct.firstNonTextListSeen = true
	}
	// add a dereference (*) in from of the ToGo() invocation for go types that aren't pointers.
	addStar := "*"
	if isPointerType(f.goTypePrefix) {
		addStar = ""
	}

	fmt.Fprintf(buf, `
    // %s
	n = src.%s().Len()
	dest.%s = make(%s%s, n)
	for i := 0; i < n; i++ {
        dest.%s[i] = %s
    }

`, f.goName, f.goCapGoName, f.goName, f.goTypePrefix, f.goType, f.goName, x.ElemStarCapToGo(addStar, f))

}

func (x *Extractor) ElemStarCapToGo(addStar string, f *Field) string {
	f.goToCapFunc = x.goToCapTypeFunction(f.capTypeSeq)

	VPrintf("\n\n f = %#v   addStar = '%v'    f.goToCapFunc = '%s'\n", f, addStar, f.goToCapFunc)

	// list of list special handling, try to generalize it, as it needs
	// to work for intrinsics and structs
	if f.goTypePrefix == "[][]" {
		VPrintf("\n slice of slice / ListList detected.\n")
		return fmt.Sprintf("%s(%s(src.%s().At(i)))", f.canonGoTypeListToSliceFunc, f.singleCapListType, f.goName)
	}

	if IsIntrinsicGoType(f.goType) {
		VPrintf("\n intrinsic detected.\n")
		return fmt.Sprintf("%s(src.%s().At(i))", f.goType, f.goName)
	} else {
		VPrintf("\n non-intrinsic detected. f.goType = '%v'\n", f.goType)
		return fmt.Sprintf("%s%sToGo(src.%s().At(i), nil)", addStar, f.goToCapFunc, f.goName)
	}
}

func (x *Extractor) goToCapTypeFunction(capTypeSeq []string) string {
	var r string
	for _, s := range capTypeSeq {
		if s != "*" && s != "List" {
			r = r + s
		}
	}
	return r
}

func isPointerType(goTypePrefix string) bool {
	if len(goTypePrefix) == 0 {
		return false
	}
	prefix := []rune(goTypePrefix)
	if prefix[len(prefix)-1] == '*' {
		return true
	}
	return false
}

func last(slc []string) string {
	n := len(slc)
	if n == 0 {
		panic("last of empty slice undefined")
	}
	return slc[n-1]
}

func IsDoubleList(f *Field) bool {
	if len(f.capTypeSeq) > 2 {
		if f.capTypeSeq[0] == f.capTypeSeq[1] && f.capTypeSeq[1] == "List" {
			return true
		}
	}
	return false
}

func (x *Extractor) SettersToCapn(goName string) string {
	var buf bytes.Buffer
	t := x.srs[goName]
	if t == nil {
		panic(fmt.Sprintf("bad goName '%s'", goName))
	}
	VPrintf("\n\n SettersToCapn running on myStruct = %#v\n", t)
	for i, f := range t.fld {
		VPrintf("\n\n SettersToCapn running on t.fld[%d] = %#v\n", i, f)

		if f.isList {
			t.listNum++
			if IsIntrinsicGoType(f.goType) {
				VPrintf("\n intrinsic detected in SettersToCapn.\n")

				if IsDoubleList(f) {
					VPrintf("\n yes IsDoubleList(f: '%s') in SettersToCapn.\n", f.goName)

					fmt.Fprintf(&buf, `

  mylist%d := seg.NewPointerList(len(src.%s))
  for i := range src.%s {
     mylist%d.Set(i, capn.Object(%s(seg, src.%s[i])))
  }
  dest.Set%s(mylist%d)
`, t.listNum, f.goName, f.goName, t.listNum, f.canonGoTypeSliceToListFunc, f.goName, f.goCapGoName, t.listNum)

				} else {
					VPrintf("\n\n  at non-List(List()), yes intrinsic list in SettersToCap(): f = %#v\n", f)
					fmt.Fprintf(&buf, `

  mylist%d := seg.New%sList(len(src.%s))
  for i := range src.%s {
     mylist%d.Set(i, %s(src.%s[i]))
  }
  dest.Set%s(mylist%d)
`, t.listNum, last(f.capTypeSeq), f.goName, f.goName, t.listNum, last(f.goCapGoTypeSeq), f.goName, f.goCapGoName, t.listNum)
				}
			} else {
				// handle list of struct
				VPrintf("\n\n  at struct list in SettersToCap(): f = %#v\n", f)
				addAmpersand := "&"
				if isPointerType(f.goTypePrefix) {
					addAmpersand = ""
				}

				// list of list needs special casing
				if isListList(f.goTypePrefix) {
					fmt.Fprintf(&buf, `
  // %s -> %s (go slice to capn list)
  if len(src.%s) > 0 {
		plist := seg.NewPointerList(len(src.%s))
		i := 0
		for _, ele := range src.%s {
			plist.Set(i, capn.Object(%s(seg, ele)))
			i++
		}
		dest.Set%s(plist)
	}
`, f.goName, f.goToCapFunc, f.goName, f.goName, f.goName, f.canonGoTypeSliceToListFunc, f.goCapGoName)
				} else {
					fmt.Fprintf(&buf, `
  // %s -> %s (go slice to capn list)
  if len(src.%s) > 0 {
		typedList := New%sList(seg, len(src.%s))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.%s {
			plist.Set(i, capn.Object(%sGoToCapn(seg, %sele)))
			i++
		}
		dest.Set%s(typedList)
	}
`, f.goName, f.goToCapFunc, f.goName, f.goToCapFunc, f.goName, f.goName, f.goType, addAmpersand, f.goCapGoName)
				}
			} // end switch f.goType

		} else {

			var isCapType bool = false
			if _, ok := x.goType2capTypeCache[f.goTypeSeq[0]]; !ok {
				if len(f.goTypeSeq) > 1 {
					if _, ok := x.goType2capTypeCache[f.goTypeSeq[1]]; ok {
						isCapType = true
					}
				}
			} else {
				isCapType = true
			}

			if isCapType {
				if f.goTypeSeq[0] == "*" {
					fmt.Fprintf(&buf, "  dest.Set%s(%sGoToCapn(seg, src.%s))\n",
						f.goName, f.goType, f.goName)
				} else {
					fmt.Fprintf(&buf, "  dest.Set%s(%sGoToCapn(seg, &src.%s))\n",
						f.goName, f.goType, f.goName)
				}

				continue
			}

			switch f.goType {
			case "int":
				fmt.Fprintf(&buf, "  dest.Set%s(int64(src.%s))\n", f.goCapGoName, f.goName)
			default:
				fmt.Fprintf(&buf, "  dest.Set%s(src.%s)\n", f.goCapGoName, f.goName)
			}
		}
	}
	return string(buf.Bytes())
}

func isListList(goTypePrefix string) bool {
	return strings.HasPrefix(goTypePrefix, "[][]")
}

func (x *Extractor) ToGoCodeFor(goName string) []byte {
	return x.ToGoCode[goName]
}

func (x *Extractor) ToCapnCodeFor(goStructName string) []byte {
	return x.ToCapnCode[goStructName]
}

func (x *Extractor) WriteToSchema(w io.Writer) (n int64, err error) {

	var m int
	var spaces string

	// sort structs alphabetically to get a stable (testable) ordering.
	sortedStructs := ByGoName(make([]*Struct, 0, len(x.srs)))
	for _, strct := range x.srs {
		sortedStructs = append(sortedStructs, strct)
	}
	sort.Sort(ByGoName(sortedStructs))

	for _, s := range sortedStructs {

		m, err = fmt.Fprintf(w, "%sstruct %s { %s", x.fieldSuffix, s.capName, x.fieldSuffix)
		n += int64(m)
		if err != nil {
			return
		}

		s.computeFinalOrder()

		sort.Sort(ByFinalOrder(s.fld))

		for i, fld := range s.fld {

			VPrintf("\n\n debug in WriteToSchema(), fld = %#v\n", fld)

			SetSpaces(&spaces, s.longestField, len(fld.capname))

			capType, already := x.goType2capTypeCache[fld.goType]
			if !already {
				VPrintf("\n\n debug: setting capType = '%s', instead of '%s', already = false\n", fld.capType, capType)
				capType = fld.capType
			} else {
				VPrintf("\n\n debug: already = true, capType = '%s'   fld.capType = %v\n", capType, fld.capType)
			}

			m, err = fmt.Fprintf(w, "%s%s  %s@%d: %s%s; %s", x.fieldPrefix, fld.capname, spaces, fld.finalOrder, ExtraSpaces(i), fld.capType, x.fieldSuffix)
			//m, err = fmt.Fprintf(w, "%s%s  %s@%d: %s%s; %s", x.fieldPrefix, fld.capname, spaces, fld.finalOrder, ExtraSpaces(i), capType, x.fieldSuffix)
			n += int64(m)
			if err != nil {
				return
			}

		} // end field loop

		m, err = fmt.Fprintf(w, "} %s", x.fieldSuffix)
		n += int64(m)
		if err != nil {
			return
		}

	} // end loop over structs

	return
}

func (x *Extractor) GenZidTag(f *Field) string {
	if f.astField == nil {
		f.astField = &ast.Field{}
	}
	if f.astField.Tag == nil {
		f.astField.Tag = &ast.BasicLit{}
	}

	curTag := f.astField.Tag.Value

	if hasZidTag(curTag) {
		return curTag
	}
	// else add one
	addme := fmt.Sprintf(`zid:"%d"`, f.finalOrder)
	if curTag == "" || curTag == "``" {
		return fmt.Sprintf("`%s`", addme)
	}
	return fmt.Sprintf("`%s %s`", stripBackticks(curTag), addme)
}

func hasZidTag(s string) bool {
	return strings.Contains(s, "zid")
}

func stripBackticks(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	if r[0] == '`' {
		r = r[1:]
	}
	if len(r) > 0 && r[len(r)-1] == '`' {
		r = r[:len(r)-1]
	}
	return string(r)
}

func (x *Extractor) CopySourceFilesAddZidTag() error {

	// run through struct fields, adding tags
	for _, s := range x.srs {
		for _, f := range s.fld {

			VPrintf("\n\n\n ********** before  f.astField.Tag = %#v\n", f.astField.Tag)
			f.astField.Tag.Value = x.GenZidTag(f)
			VPrintf("\n\n\n ********** AFTER:  f.astField.Tag = %#v\n", f.astField.Tag)
		}
	}

	// run through files, printing
	for _, s := range x.srcFiles {
		if s.filename != "" {
			err := x.PrettyPrint(s.fset, s.astFile, x.compileDir.DirPath+string(os.PathSeparator)+s.filename)
			if err != nil {
				return err
			}
		}
	}

	if x.overwrite {
		bk := fmt.Sprintf("%s%cbk%c", x.compileDir.DirPath, os.PathSeparator, os.PathSeparator)
		err := os.MkdirAll(bk, 0755)
		if err != nil {
			panic(err)
		}
		for _, s := range x.srcFiles {
			if s.filename != "" {
				// make a backup
				err := Cp(s.filename, bk+s.filename)
				if err != nil {
					panic(err)
				}
				// overwrite
				err = Cp(x.compileDir.DirPath+string(os.PathSeparator)+s.filename, s.filename)
				if err != nil {
					panic(err)
				}
			}
		}

	}

	return nil
}

func (x *Extractor) WriteToTranslators(w io.Writer) (n int64, err error) {

	var m int

	x.GenerateTranslators()

	// sort structs alphabetically to get a stable (testable) ordering.
	sortedStructs := ByGoName(make([]*Struct, 0, len(x.srs)))
	for _, strct := range x.srs {
		sortedStructs = append(sortedStructs, strct)
	}
	sort.Sort(ByGoName(sortedStructs))

	// now print the translating methods, in a second pass over structures, to accomodate
	// our test structure
	for _, s := range sortedStructs {

		m, err = fmt.Fprintf(w, "\n\n")
		n += int64(m)
		if err != nil {
			return
		}

		m, err = w.Write(x.SaveCode[s.goName])
		n += int64(m)
		if err != nil {
			return
		}

		m, err = fmt.Fprintf(w, "\n\n")
		n += int64(m)
		if err != nil {
			return
		}

		m, err = w.Write(x.LoadCode[s.goName])
		n += int64(m)
		if err != nil {
			return
		}

		m, err = fmt.Fprintf(w, "\n\n")
		n += int64(m)
		if err != nil {
			return
		}

		m, err = w.Write(x.ToGoCodeFor(s.goName))
		n += int64(m)
		if err != nil {
			return
		}

		m, err = fmt.Fprintf(w, "\n\n")
		n += int64(m)
		if err != nil {
			return
		}

		m, err = w.Write(x.ToCapnCodeFor(s.goName))
		n += int64(m)
		if err != nil {
			return
		}

	} // end second loop over structs for translating methods.

	// print the helpers made from x.GenerateListHelpers(capListTypeSeq, goTypeSeq)
	// sort helper functions to get consistent (testable) order.
	a := make([]AlphaHelper, len(x.SliceToListCode)+len(x.ListToSliceCode))
	i := 0
	for k, v := range x.SliceToListCode {
		a[i].Name = k
		a[i].Code = v
		i++
	}
	for k, v := range x.ListToSliceCode {
		a[i].Name = k
		a[i].Code = v
		i++
	}

	sort.Sort(AlphaHelperSlice(a))

	for _, help := range a {

		m, err = fmt.Fprintf(w, "\n\n")
		n += int64(m)
		if err != nil {
			return
		}

		m, err = w.Write(help.Code)
		n += int64(m)
		if err != nil {
			return
		}

	}

	return
}

func ExtractFromString(src string) ([]byte, error) {
	return ExtractStructs("", "package main; "+src, nil)
}

func ExtractString2String(src string) string {

	x := NewExtractor()
	defer x.Cleanup()
	_, err := ExtractStructs("", "package main; "+src, x)
	if err != nil {
		panic(err)
	}

	//goon.Dump(x.srs)

	// final write, this time accounting for zid tag ordering
	var buf bytes.Buffer
	_, err = x.WriteToSchema(&buf)
	if err != nil {
		panic(err)
	}
	_, err = x.WriteToTranslators(&buf)
	if err != nil {
		panic(err)
	}

	return string(buf.Bytes())
}

func ExtractCapnToGoCode(src string, goName string) string {

	x := NewExtractor()
	defer x.Cleanup()
	_, err := ExtractStructs("", "package main; "+src, x)
	if err != nil {
		panic(err)
	}
	x.GenerateTranslators()
	return string(x.ToGoCodeFor(goName))
}

func ExtractGoToCapnCode(src string, goName string) string {

	x := NewExtractor()
	defer x.Cleanup()
	_, err := ExtractStructs("", "package main; "+src, x)
	if err != nil {
		panic(err)
	}
	x.GenerateTranslators()
	return string(x.ToCapnCodeFor(goName))
}

// ExtractStructs pulls out the struct definitions from a golang source file.
//
// src has to be string, []byte, or io.Reader, as in parser.ParseFile(). src
// can be nil if fname is provided. See http://golang.org/pkg/go/parser/#ParseFile
//
func ExtractStructs(fname string, src interface{}, x *Extractor) ([]byte, error) {
	if x == nil {
		x = NewExtractor()
		defer x.Cleanup()
	}

	return x.ExtractStructsFromOneFile(src, fname)
}

func (x *Extractor) ExtractStructsFromOneFile(src interface{}, fname string) ([]byte, error) {

	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, fname, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	if fname != "" {
		x.srcFiles = append(x.srcFiles, &SrcFile{filename: fname, fset: fset, astFile: f})
	}

	//	VPrintf("parsed output f.Decls is:\n")
	//VPrintf("len(f.Decls) = %d\n", len(f.Decls))

	for _, v := range f.Decls {
		switch v.(type) {
		case *ast.GenDecl:
			d := v.(*ast.GenDecl)
			//VPrintf("dump of d, an %#v = \n", ty)
			//goon.Dump(d)

			//VPrintf("\n\n\n  detail dump of d.Specs\n")
			for _, spe := range d.Specs {
				switch spe.(type) {
				case (*ast.TypeSpec):

					// go back and print the comments
					if d.Doc != nil && d.Doc.List != nil && len(d.Doc.List) > 0 {
						for _, com := range d.Doc.List {
							x.GenerateComment(com.Text)
						}
					}

					typeSpec := spe.(*ast.TypeSpec)
					//VPrintf("\n\n *ast.TypeSpec spe = \n")

					if typeSpec.Name.Obj.Kind == ast.Typ {

						switch typeSpec.Name.Obj.Decl.(type) {
						case (*ast.TypeSpec):
							//curStructName := typeSpec.Name.String()
							curStructName := typeSpec.Name.Name
							ts2 := typeSpec.Name.Obj.Decl.(*ast.TypeSpec)

							//VPrintf("\n\n  in ts2 = %#v\n", ts2)
							//goon.Dump(ts2)

							switch ty := ts2.Type.(type) {
							default:
								// *ast.InterfaceType and *ast.Ident end up here.
								//VPrintf("\n\n unrecog type ty = %#v\n", ty)
							case (*ast.Ident):
								goNewTypeName := ts2.Name.Obj.Name
								goTargetTypeName := ty.Name
								x.NoteTypedef(goNewTypeName, goTargetTypeName)

							case (*ast.StructType):
								stru := ts2.Type.(*ast.StructType)

								err = x.StartStruct(curStructName)
								if err != nil {
									return []byte{}, err
								}
								//VPrintf("\n\n stru = %#v\n", stru)
								//goon.Dump(stru)

								if stru.Fields != nil {
									for _, fld := range stru.Fields.List {
										if fld != nil {
											//VPrintf("\n\n    fld.Names = %#v\n", fld.Names) // looking for
											//goon.Dump(fld.Names)

											if len(fld.Names) == 0 {
												// field without name: embedded/anonymous struct
												var typeName string
												switch nmmm := fld.Type.(type) {
												case *ast.StarExpr:
													// ???
												case *ast.Ident:
													typeName = nmmm.Name
													err = x.GenerateStructField(typeName, "", typeName, fld, NotList, fld.Tag, YesEmbedded, []string{typeName})
													if err != nil {
														return []byte{}, err
													}
												}

											} else {
												// field with name
												for _, ident := range fld.Names {

													switch ident.Obj.Decl.(type) {
													case (*ast.Field):
														// named field
														fld2 := ident.Obj.Decl.(*ast.Field)

														//VPrintf("\n\n    fld2 = %#v\n", fld2)
														//goon.Dump(fld2)

														typeNamePrefix, ident4, gotypeseq := GetTypeAsString(fld2.Type, "", []string{})
														//VPrintf("\n\n tnas = %#v, ident4 = %s\n", typeNamePrefix, ident4)

														err = x.GenerateStructField(ident.Name, typeNamePrefix, ident4, fld2, IsSlice(typeNamePrefix), fld2.Tag, NotEmbedded, gotypeseq)
														if err != nil {
															return []byte{}, err
														}
													}
												}
											}

										}
									}
								}

								//VPrintf("} // end of %s \n\n", typeSpec.Name) // prod
								x.EndStruct()

								//goon.Dump(stru)
								//VPrintf("\n =========== end stru =======\n\n\n")
							}
						}
					}

				}
			}
		}
	}

	return x.out.Bytes(), err
}

func IsSlice(tnas string) bool {
	return strings.HasPrefix(tnas, "[]")
}

func (x *Extractor) NoteTypedef(goNewTypeName string, goTargetTypeName string) {
	// we just want to preserve the mapping, without adding Capn suffix
	//VPrintf("\n\n noting typedef: goNewTypeName: '%s', goTargetTypeName: '%s'\n", goNewTypeName, goTargetTypeName)
	var capTypeSeq []string
	capTypeSeq, x.goType2capTypeCache[goNewTypeName] = x.GoTypeToCapnpType(nil, []string{goTargetTypeName})
	VPrintf("\n\n 888 noting typedef: goNewTypeName: '%s', goTargetTypeName: '%s'   x.goType2capTypeCache[goNewTypeName]: '%v'  capTypeSeq: '%v'\n", goNewTypeName, goTargetTypeName, x.goType2capTypeCache[goNewTypeName], capTypeSeq)
}

var regexCapname = regexp.MustCompile(`capname:[ \t]*\"([^\"]+)\"`)

var regexZid = regexp.MustCompile(`zid:[ \t]*\"([^\"]+)\"`)

func GoType2CapnType(gotypeName string) string {
	return UppercaseFirstLetter(gotypeName) + "Capn"
}

func (x *Extractor) StartStruct(goName string) error {
	x.fieldCount = 0

	capname := GoType2CapnType(goName)
	x.goType2capTypeCache[goName] = capname

	VPrintf("\n\n debug 777 setting x.goType2capTypeCache['%s'] = '%s'\n", goName, capname)

	x.capType2goType[capname] = goName

	// check for rename comment, capname:"newCapName"
	if x.heldComment != "" {

		match := regexCapname.FindStringSubmatch(x.heldComment)
		if match != nil {
			if len(match) == 2 {
				capname = match[1]
			}
		}
	}

	if isCapnpKeyword(capname) {
		err := fmt.Errorf(`after uppercasing the first letter, struct '%s' becomes '%s' but this is a reserved capnp word, so please write a comment annotation just before the struct definition in go (e.g. // capname:"capName") to rename it.`, goName, capname)
		panic(err)
		return err
	}

	fmt.Fprintf(&x.out, "struct %s { %s", capname, x.fieldSuffix)

	x.curStruct = NewStruct(capname, goName)
	x.curStruct.comment = x.heldComment
	x.heldComment = ""
	x.srs[goName] = x.curStruct

	return nil
}
func (x *Extractor) EndStruct() {
	fmt.Fprintf(&x.out, "} %s", x.fieldSuffix)
}

func (x *Extractor) GenerateComment(c string) {
	x.heldComment = x.heldComment + c + "\n"
}

func UppercaseFirstLetter(name string) string {
	if len(name) == 0 {
		return name
	}

	// gotta upercase the first letter of type (struct) names
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)

}

func LowercaseCapnpFieldName(name string) string {
	if len(name) == 0 {
		return name
	}

	// gotta lowercase the first letter of field names
	runes := []rune(name)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

const YesIsList = true
const NotList = false

const NotEmbedded = false
const YesEmbedded = true

func (x *Extractor) GenerateStructField(goFieldName string, goFieldTypePrefix string, goFieldTypeName string, astfld *ast.Field, isList bool, tag *ast.BasicLit, IsEmbedded bool, goTypeSeq []string) error {

	if goFieldTypeName == "" {
		return nil
	}

	// skip protobuf side effects
	if goFieldTypeName == "XXX_unrecognized" {
		fmt.Printf("detected XXX_unrecognized\n")
		return nil
	}

	VPrintf("\n\n\n GenerateStructField called with goFieldName = '%s', goFieldTypeName = '%s', astfld = %#v, tag = %#v\n\n", goFieldName, goFieldTypeName, astfld, tag)

	// if we are ignoring private (lowercase first letter) fields, then stop here.
	if !IsEmbedded {
		if len(goFieldName) > 0 && unicode.IsLower([]rune(goFieldName)[0]) && !x.extractPrivate {
			return nil
		}
	}

	curField := &Field{orderOfAppearance: x.fieldCount, embedded: IsEmbedded, astField: astfld, goTypeSeq: goTypeSeq, capTypeSeq: []string{}}

	var tagValue string
	loweredName := underToCamelCase(LowercaseCapnpFieldName(goFieldName))

	if tag != nil {
		VPrintf("tag = %#v\n", tag)

		if tag.Value != "" {

			// capname tag
			match := regexCapname.FindStringSubmatch(tag.Value)
			if match != nil {
				if len(match) == 2 {
					VPrintf("matched, using '%s' instead of '%s'\n", match[1], goFieldName)
					loweredName = match[1]
					tagValue = tag.Value

					if isCapnpKeyword(loweredName) {
						err := fmt.Errorf(`problem detected after applying the capname tag '%s' found on field '%s': '%s' is a reserved capnp word, so please use a *different* struct field tag (e.g. capname:"capnpName") to rename it`, tag.Value, goFieldName, loweredName)
						return err
					}

				}
			}

			// zid tag
			match2 := regexZid.FindStringSubmatch(tag.Value)
			if match2 != nil {
				if len(match2) == 2 {
					if match2[1] == "skip" {
						VPrintf("skipping field '%s' marked with zid:\"skip\"", loweredName)
						return nil
					}
					VPrintf("matched, applying zid tag '%s' for field '%s'\n", match2[1], loweredName)
					n, err := strconv.Atoi(match2[1])
					if err != nil {
						err := fmt.Errorf(`problem in zid tag '%s' on field '%s' in struct '%s': could not convert to number, error: '%s'`, match2[1], goFieldName, x.curStruct.goName, err)
						panic(err)
						return err
					}
					if n < 0 {
						VPrintf("skipping field '%s' marked with negative zid:\"%d\"", loweredName, n)
						return nil
					}
					fld, already := x.curStruct.zidMap[n]
					if already {
						err := fmt.Errorf(`problem in zid tag '%s' on field '%s' in struct '%s': number '%d' is already taken by field '%s'`, match2[1], goFieldName, x.curStruct.goName, n, fld.goName)
						panic(err)
						return err

					} else {
						x.curStruct.zidMap[n] = curField
						curField.zidFromTag = n
					}
				}
			}
		}

	}

	VPrintf("\n\n\n GenerateStructField: goFieldName:'%s' -> loweredName:'%s'\n\n", goFieldName, loweredName)

	if isCapnpKeyword(loweredName) {
		err := fmt.Errorf(`after lowercasing the first letter, field '%s' becomes '%s' but this is a reserved capnp word, so please use a struct field tag (e.g. capname:"capnpName") to rename it`, goFieldName, loweredName)
		return err
	}

	var capnTypeDisplayed string
	curField.capTypeSeq, capnTypeDisplayed = x.GoTypeToCapnpType(curField, goTypeSeq)

	VPrintf("\n\n\n DEBUG:  '%s' '%s' @%d: %s; %s\n\n", x.fieldPrefix, loweredName, x.fieldCount, capnTypeDisplayed, x.fieldSuffix)

	sz := len(loweredName)
	if sz > x.curStruct.longestField {
		x.curStruct.longestField = sz
	}

	curField.capname = loweredName

	curField.goCapGoName = UppercaseFirstLetter(loweredName)

	curField.goCapGoTypeSeq, curField.goCapGoType = x.CapnTypeToGoType(curField.capTypeSeq)

	curField.capType = capnTypeDisplayed
	curField.goName = goFieldName
	curField.goType = goFieldTypeName
	if len(curField.capTypeSeq) > 0 && curField.capTypeSeq[0] == "List" {
		curField.isList = true
	}
	curField.tagValue = tagValue
	curField.goTypePrefix = goFieldTypePrefix

	x.curStruct.fld = append(x.curStruct.fld, curField)
	x.fieldCount++

	VPrintf("\n\n curField = %#v\n", curField)

	return nil
}

func (x *Extractor) CapnTypeToGoType(capTypeSeq []string) (goTypeSeq []string, displayGoCapGoType string) {

	for _, c := range capTypeSeq {
		goType, special := x.c2g(c)

		if special {
			if goType == "Data" {
				goTypeSeq = append(goTypeSeq, "[]", "byte")
				continue
			}

			if x.capType2goType[c] != "" {
				goType = x.capType2goType[c]
			}
		}
		goTypeSeq = append(goTypeSeq, goType)
	}
	return goTypeSeq, x.assembleGoType(goTypeSeq)
}

func (x *Extractor) assembleGoType(goTypeSeq []string) string {
	// make a legitimate go type
	return strings.Join(goTypeSeq, "")
}

// special flags Data <-> []byte, and types we couldn't convert
func (x *Extractor) c2g(capType string) (goType string, special bool) {

	switch capType {
	default:
		return capType, true
	case "Data":
		return "Data", true
	case "List":
		return "[]", false
	case "Text":
		return "string", false
	case "Bool":
		return "bool", false
	case "Int8":
		return "int8", false
	case "Int16":
		return "int16", false
	case "Int32":
		return "int32", false
	case "Int64":
		return "int64", false
	case "UInt8":
		return "uint8", false
	case "UInt16":
		return "uint16", false
	case "UInt32":
		return "uint32", false
	case "UInt64":
		return "uint64", false
	case "Float32":
		return "float32", false
	case "Float64":
		return "float64", false
	}
}

func (x *Extractor) GoTypeToCapnpType(curField *Field, goTypeSeq []string) (capTypeSeq []string, capnTypeDisplayed string) {

	VPrintf("\n\n In GoTypeToCapnpType() : goTypeSeq=%#v)\n", goTypeSeq)

	capTypeSeq = make([]string, len(goTypeSeq))
	for i, t := range goTypeSeq {
		capTypeSeq[i] = x.g2c(t)
	}

	// now that the capTypeSeq is completely generated, check for lists
	// currently only do List(primitive or struct type); no List(List(prim)) or List(List(struct))
	n := len(capTypeSeq)
	for i, ty := range capTypeSeq {
		if ty == "List" && i == n-2 {
			VPrintf("\n\n generating List helpers at i=%d, capTypeSeq = '%#v\n", i, capTypeSeq)
			x.GenerateListHelpers(curField, capTypeSeq[i:], goTypeSeq[i:])
		}
	}

	return capTypeSeq, x.assembleCapType(capTypeSeq)
}

func (x *Extractor) assembleCapType(capTypeSeq []string) string {
	// make a legitimate capnp type
	switch capTypeSeq[0] {
	case "List":
		return "List(" + x.assembleCapType(capTypeSeq[1:]) + ")"
	case "*":
		return x.assembleCapType(capTypeSeq[1:])
	default:
		return capTypeSeq[0]
	}
}

func (x *Extractor) g2c(goFieldTypeName string) string {

	switch goFieldTypeName {
	case "[]":
		return "List"
	case "*":
		return "*"
	case "string":
		return "Text"
	case "int":
		return "Int64"
	case "bool":
		return "Bool"
	case "int8":
		return "Int8"
	case "int16":
		return "Int16"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint8":
		return "UInt8"
	case "uint16":
		return "UInt16"
	case "uint32":
		return "UInt32"
	case "uint64":
		return "UInt64"
	case "float32":
		return "Float32"
	case "float64":
		return "Float64"
	case "byte":
		return "UInt8"
	}

	var capnTypeDisplayed string
	alreadyKnownCapnType := x.goType2capTypeCache[goFieldTypeName]
	if alreadyKnownCapnType != "" {
		VPrintf("\n\n debug: x.goType2capTypeCache[goFieldTypeName='%s'] -> '%s'\n", goFieldTypeName, alreadyKnownCapnType)
		capnTypeDisplayed = alreadyKnownCapnType
	} else {
		capnTypeDisplayed = GoType2CapnType(goFieldTypeName)
		VPrintf("\n\n 999 debug: adding to  x.goType2capTypeCache[goFieldTypeName='%s'] = '%s'\n", goFieldTypeName, capnTypeDisplayed)
		x.goType2capTypeCache[goFieldTypeName] = capnTypeDisplayed
	}

	return capnTypeDisplayed
}

func (x *Extractor) GenerateEmbedded(typeName string) {
	fmt.Fprintf(&x.out, "%s; ", typeName) // prod
}

func getNewCapnpId() string {
	id, err := exec.Command("capnp", "id").CombinedOutput()
	if err != nil {
		panic(err)
	}
	n := len(id)
	if n > 0 {
		id = id[:n-1]
	}

	return string(id)
}

func (x *Extractor) GenCapnpHeader() *bytes.Buffer {
	var by bytes.Buffer

	id := getNewCapnpId()

	fmt.Fprintf(&by, `%s;
using Go = import "go.capnp";
$Go.package("%s");
$Go.import("%s");
%s`, id, x.pkgName, x.importDecl, x.fieldSuffix)

	return &by
}

func (x *Extractor) AssembleCapnpFile(in []byte) *bytes.Buffer {
	by := x.GenCapnpHeader()

	by.Write(in)
	fmt.Fprintf(by, "\n")

	return by
}

func CapnpCompileFragment(in []byte) ([]byte, error, *Extractor) {
	x := NewExtractor()
	out, err := x.CapnpCompileFragment(in)
	return out, err, x
}

func (x *Extractor) CapnpCompileFragment(in []byte) ([]byte, error) {

	if x.compileDir != nil {
		x.compileDir.Cleanup()
	}
	x.compileDir = NewTempDir()

	f := x.compileDir.TempFile()

	by := x.AssembleCapnpFile(in)
	debug := string(by.Bytes())

	f.Write(by.Bytes())
	f.Close()

	compiled, combinedOut, err := CapnpCompilePath(f.Name())
	if err != nil {
		errmsg := fmt.Sprintf("error compiling the generated capnp code: '%s'; error: '%s'\n", debug, err) + string(combinedOut)
		return []byte(errmsg), fmt.Errorf(errmsg)
	}

	return compiled, nil
}

func CapnpCompilePath(fname string) (generatedGoFile []byte, comboOut []byte, err error) {
	goOutFn := fname + ".go"

	by, err := exec.Command("capnp", "compile", "-ogo", fname).CombinedOutput()
	if err != nil {
		return []byte{}, by, err
	}

	generatedGoFile, err = ioutil.ReadFile(goOutFn)

	return generatedGoFile, by, err
}

func SetSpaces(spaces *string, Max int, Len int) {
	if Len >= Max {
		*spaces = ""
		return
	}
	*spaces = strings.Repeat(" ", Max-Len)
}

func ExtraSpaces(fieldNum int) string {
	if fieldNum < 10 {
		return "  "
	}
	if fieldNum < 100 {
		return " "
	}
	return ""
}

func IsIntrinsicGoType(goFieldTypeName string) bool {
	VPrintf("\n IsIntrinsic called with '%s'\n", goFieldTypeName)

	switch goFieldTypeName {
	case "string":
		return true
	case "int":
		return true
	case "bool":
		return true
	case "int8":
		return true
	case "int16":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "uint8":
		return true
	case "uint16":
		return true
	case "uint32":
		return true
	case "uint64":
		return true
	case "float32":
		return true
	case "float64":
		return true
	case "byte":
		return true
	default:
		return false
	}
	return false
}

func CanonGoType(goTypeSeq []string) string {
	var r string
	for _, s := range goTypeSeq {
		if s == "[]" {
			r += "Slice"
		} else {
			r += UppercaseFirstLetter(s)
		}
	}
	return r
}

func CanonCapType(capTypeSeq []string) string {
	var r string
	for _, s := range capTypeSeq {
		r += s
	}
	return r
}

func (x *Extractor) GenerateListHelpers(f *Field, capListTypeSeq []string, goTypeSeq []string) {

	canonGoType := CanonGoType(goTypeSeq)
	//canonCapType := CanonCapType(capListTypeSeq)

	// already done before? but maybe f was nil and so needs re-doing!
	// so don't return early even if partially set before!

	VPrintf("\n\n debug GenerateListHelper: called with capListTypeSeq = '%#v'\n", capListTypeSeq)

	n := len(capListTypeSeq)
	capBaseType := capListTypeSeq[n-1]
	capTypeThenList := strings.Join(capListTypeSeq, "")
	if n > 1 {
		capTypeThenList = capBaseType + strings.Join(capListTypeSeq[:n-1], "")
	}

	if IsIntrinsicGoType(last(goTypeSeq)) {
		f.baseIsIntrinsic = true
		capTypeThenList = capBaseType + "List"
		f.singleCapListType = fmt.Sprintf("capn.%sList", capBaseType)
		f.newListExpression = fmt.Sprintf("seg.New%sList(len(m))", capBaseType)
	} else {
		capTypeThenList = capBaseType + strings.Join(capListTypeSeq[:n-1], "")
		f.singleCapListType = fmt.Sprintf("%s_List", capBaseType)
		f.newListExpression = fmt.Sprintf("New%s(seg, len(m))", capTypeThenList)
	}
	VPrintf("\n capTypeThenList is set to : '%s'\n\n", capTypeThenList)

	collapGoType := strings.Join(goTypeSeq, "")
	m := len(goTypeSeq)
	goBaseType := goTypeSeq[m-1]

	c2g, _ := x.c2g(capBaseType)

	f.canonGoType = canonGoType
	f.canonGoTypeListToSliceFunc = fmt.Sprintf("%sTo%s", capTypeThenList, canonGoType)
	f.canonGoTypeSliceToListFunc = fmt.Sprintf("%sTo%s", canonGoType, capTypeThenList)

	x.SliceToListCode[canonGoType] = []byte(fmt.Sprintf(`
func %sTo%s(seg *capn.Segment, m %s) %s {
	lst := %s
	for i := range m {
		lst.Set(i, %s)
	}
	return lst
}
`, canonGoType, capTypeThenList, collapGoType, f.singleCapListType, f.newListExpression, x.SliceToListSetRHS(f.baseIsIntrinsic, goBaseType, c2g)))

	x.ListToSliceCode[canonGoType] = []byte(fmt.Sprintf(`
func %sTo%s(p %s) %s {
	v := make(%s, p.Len())
	for i := range v {
        %s
	}
	return v
}
`, capTypeThenList, canonGoType, f.singleCapListType, collapGoType, collapGoType, x.ListToSliceSetLHS_RHS(f.baseIsIntrinsic, capBaseType, goBaseType)))

	VPrintf("\n\n GenerateListHelpers done for field '%#v'\n\n", f)
}

func (x *Extractor) SliceToListSetRHS(baseIsIntrinsic bool, goName string, c2g string) string {
	if baseIsIntrinsic {
		return fmt.Sprintf("%s(m[i])", c2g)
	} else {
		return fmt.Sprintf("%sGoToCapn(seg, &m[i])", goName)
	}
}

func (x *Extractor) ListToSliceSetLHS_RHS(baseIsIntrinsic bool, capName string, goBaseType string) string {
	if baseIsIntrinsic {
		return fmt.Sprintf("v[i] = %s(p.At(i))", goBaseType) // e.g. int64
	} else {
		return fmt.Sprintf("%sToGo(p.At(i), &v[i])", capName)
	}
}
