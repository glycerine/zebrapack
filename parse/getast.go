package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/glycerine/zebrapack/cfg"
	"github.com/glycerine/zebrapack/gen"
	"github.com/ttacon/chalk"
)

var StopOnError bool

// A FileSet is the in-memory representation of a
// parsed file.
type FileSet struct {
	Package    string              // package name
	Specs      map[string]ast.Expr // type specs in file
	Identities map[string]gen.Elem // processed from specs
	Directives []string            // raw preprocessor directives
	Imports    []*ast.ImportSpec   // imports
	Cfg        *cfg.ZebraConfig
}

// File parses a file at the relative path
// provided and produces a new *FileSet.
// If you pass in a path to a directory, the entire
// directory will be parsed.
// If unexport is false, only exported identifiers are included in the FileSet.
// If the resulting FileSet would be empty, an error is returned.
func File(c *cfg.ZebraConfig) (*FileSet, error) {
	name := c.GoFile
	pushstate(name)
	defer popstate()
	fs := &FileSet{
		Specs:      make(map[string]ast.Expr),
		Identities: make(map[string]gen.Elem),
		Cfg:        c,
	}

	fset := token.NewFileSet()
	finfo, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	if finfo.IsDir() {
		pkgs, err := parser.ParseDir(fset, name, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		if len(pkgs) != 1 {
			return nil, fmt.Errorf("multiple packages in directory: %s", name)
		}
		var one *ast.Package
		for _, nm := range pkgs {
			one = nm
			break
		}
		fs.Package = one.Name
		for _, fl := range one.Files {
			pushstate(fl.Name.Name)
			fs.Directives = append(fs.Directives, yieldComments(fl.Comments)...)
			if !c.Unexported {
				ast.FileExports(fl)
			}
			fs.getTypeSpecs(fl)
			popstate()
		}
	} else {
		f, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
		if err != nil {
			return nil, err
		}
		fs.Package = f.Name.Name
		fs.Directives = yieldComments(f.Comments)
		if !c.Unexported {
			ast.FileExports(f)
		}
		fs.getTypeSpecs(f)
	}

	if len(fs.Specs) == 0 {
		return nil, fmt.Errorf("no definitions in %s", name)
	}

	err = fs.process()
	if err != nil {
		return nil, err
	}
	fs.applyDirectives()
	fs.propInline()

	return fs, nil
}

// applyDirectives applies all of the directives that
// are known to the parser. additional method-specific
// directives remain in f.Directives
func (f *FileSet) applyDirectives() {
	newdirs := make([]string, 0, len(f.Directives))
	for _, d := range f.Directives {
		chunks := strings.Split(d, " ")
		if len(chunks) > 0 {
			if fn, ok := directives[chunks[0]]; ok {
				pushstate(chunks[0])
				err := fn(chunks, f)
				if err != nil {
					warnln(err.Error())
				}
				popstate()
			} else {
				newdirs = append(newdirs, d)
			}
		}
	}
	f.Directives = newdirs
}

// A linkset is a graph of unresolved
// identities.
//
// Since gen.Ident can only represent
// one level of type indirection (e.g. Foo -> uint8),
// type declarations like `type Foo Bar`
// aren't resolve-able until we've processed
// everything else.
//
// The goal of this dependency resolution
// is to distill the type declaration
// into just one level of indirection.
// In other words, if we have:
//
//  type A uint64
//  type B A
//  type C B
//  type D C
//
// ... then we want to end up
// figuring out that D is just a uint64.
type linkset map[string]*gen.BaseElem

func (f *FileSet) resolve(ls linkset) {
	progress := true
	for progress && len(ls) > 0 {
		progress = false
		for name, elem := range ls {
			real, ok := f.Identities[elem.TypeName()]
			if ok {
				// copy the old type descriptor,
				// alias it to the new value,
				// and insert it into the resolved
				// identities list
				progress = true
				nt := real.Copy()
				nt.Alias(name)
				f.Identities[name] = nt
				delete(ls, name)
			}
		}
	}

	// what's left can't be resolved
	for name, elem := range ls {
		warnf("couldn't resolve type %s (%s)\n", name, elem.TypeName())
	}
}

// process takes the contents of f.Specs and
// uses them to populate f.Identities
func (f *FileSet) process() error {

	deferred := make(linkset)
parse:
	for name, def := range f.Specs {
		pushstate(name)
		el, err := f.parseExpr(def)
		if err != nil {
			return err
		}
		if el == nil {
			warnln("failed to parse")
			popstate()
			continue parse
		}
		// push unresolved identities into
		// the graph of links and resolve after
		// we've handled every possible named type.
		if be, ok := el.(*gen.BaseElem); ok && be.Value == gen.IDENT {
			deferred[name] = be
			popstate()
			continue parse
		}
		el.Alias(name)
		f.Identities[name] = el
		popstate()
	}

	if len(deferred) > 0 {
		f.resolve(deferred)
	}
	return nil
}

func strToMethod(s string) gen.Method {
	switch s {
	case "encode":
		return gen.Encode
	case "decode":
		return gen.Decode
	case "test":
		return gen.Test
	case "size":
		return gen.Size
	case "marshal":
		return gen.Marshal
	case "unmarshal":
		return gen.Unmarshal
	default:
		return 0
	}
}

func (f *FileSet) applyDirs(p *gen.Printer) {
	// apply directives of the form
	//
	// 	//msgp:encode ignore {{TypeName}}
	//
loop:
	for _, d := range f.Directives {
		chunks := strings.Split(d, " ")
		if len(chunks) > 1 {
			for i := range chunks {
				chunks[i] = strings.TrimSpace(chunks[i])
			}
			m := strToMethod(chunks[0])
			if m == 0 {
				warnf("unknown pass name: %q\n", chunks[0])
				continue loop
			}
			if fn, ok := passDirectives[chunks[1]]; ok {
				pushstate(chunks[1])
				err := fn(m, chunks[2:], p)
				if err != nil {
					warnf("error applying directive: %s\n", err)
				}
				popstate()
			} else {
				warnf("unrecognized directive %q\n", chunks[1])
			}
		} else {
			warnf("empty directive: %q\n", d)
		}
	}
}

func (f *FileSet) PrintTo(p *gen.Printer) error {
	f.applyDirs(p)
	names := make([]string, 0, len(f.Identities))
	for name := range f.Identities {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		el := f.Identities[name]
		el.SetVarname("z")
		pushstate(el.TypeName())
		err := p.Print(el)
		popstate()
		if err != nil {
			return err
		}
	}
	return nil
}

// getTypeSpecs extracts all of the *ast.TypeSpecs in the file
// into fs.Identities, but does not set the actual element
func (fs *FileSet) getTypeSpecs(f *ast.File) {

	// collect all imports...
	fs.Imports = append(fs.Imports, f.Imports...)

	// check all declarations...
	for i := range f.Decls {

		// for GenDecls...
		if g, ok := f.Decls[i].(*ast.GenDecl); ok {

			// and check the specs...
			for _, s := range g.Specs {

				// for ast.TypeSpecs....
				if ts, ok := s.(*ast.TypeSpec); ok {
					switch ts.Type.(type) {

					// this is the list of parse-able
					// type specs
					case *ast.StructType,
						*ast.ArrayType,
						*ast.StarExpr,
						*ast.MapType,
						*ast.Ident:
						fs.Specs[ts.Name.Name] = ts.Type

					}
				}
			}
		}
	}
}

func fieldName(f *ast.Field) string {
	switch len(f.Names) {
	case 0:
		return stringify(f.Type)
	case 1:
		return f.Names[0].Name
	default:
		return f.Names[0].Name + " (and others)"
	}
}

type zid struct {
	zid       int64
	fieldName string
}

type zidSetSlice []zid

func (p zidSetSlice) Len() int           { return len(p) }
func (p zidSetSlice) Less(i, j int) bool { return p[i].zid < p[j].zid }
func (p zidSetSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (fs *FileSet) parseFieldList(fl *ast.FieldList) ([]gen.StructField, error) {
	if fl == nil || fl.NumFields() == 0 {
		return nil, nil
	}
	out := make([]gen.StructField, 0, fl.NumFields())
	var zidSet []zid
	for _, field := range fl.List {
		pushstate(fieldName(field))
		fds, err := fs.getField(field)
		if err != nil {
			fatalf(err.Error())
			return nil, err
		}
		for _, x := range fds {
			//fmt.Printf("\n on field '%#v'\n", x)
			if x.ZebraId >= 0 {
				zidSet = append(zidSet, zid{zid: x.ZebraId, fieldName: x.FieldName})
			}
		}
		if len(fds) > 0 {
			out = append(out, fds...)
		} else {
			warnln(fmt.Sprintf("ignored. heh, on '%#v'", fs))
		}
		popstate()
	}
	// check zidSet sequential from 0, no gaps, no duplicates
	if len(zidSet) > 0 {
		sort.Sort(zidSetSlice(zidSet))
		if zidSet[0].zid != 0 {
			return nil, fmt.Errorf("zid (zebra id tags on struct fields) must start at 0; lowest zid was '%v' at field '%v'", zidSet[0].zid, zidSet[0].fieldName)
		}
		for i := range zidSet {
			if zidSet[i].zid != int64(i) {
				return nil, fmt.Errorf("zid sequence interrupted - commit conflict possible! gap or duplicate in zid sequence (saw %v; expected %v), near field '%v'", zidSet[i].zid, i, zidSet[i].fieldName)
			}
		}
	}
	return out, nil
}

func anyMatches(haystack []string, needle string) bool {
	needle = strings.TrimSpace(needle)
	for _, v := range haystack {
		tr := strings.TrimSpace(v)
		if tr == needle {
			return true
		}
	}
	return false
}

// translate *ast.Field into []gen.StructField
func (fs *FileSet) getField(f *ast.Field) ([]gen.StructField, error) {
	sf := make([]gen.StructField, 1)
	var extension bool
	var omitempty bool

	var skip bool
	var deprecated bool
	var zebraId int64 = -1

	// parse tag; otherwise field name is field tag
	if f.Tag != nil {
		alltags := reflect.StructTag(strings.Trim(f.Tag.Value, "`"))
		body := alltags.Get("msg")
		tags := strings.Split(body, ",")
		if len(tags) == 2 && tags[1] == "extension" {
			extension = true
		}
		// must use msg:",omitempty" if no alt name, to
		// mark a field omitempty. this avoids confusion
		// with any alt name, which always comes first.
		if len(tags) > 1 && anyMatches(tags[1:], "omitempty") {
			omitempty = true
		}
		// ignore "-" fields
		if tags[0] == "-" {
			skip = true
			// can't return early, need to track deprecated zid.
			//return nil, nil
		}
		if len(tags[0]) > 0 {
			sf[0].FieldTag = tags[0]
		}

		// check deprecated
		dep := alltags.Get("deprecated")
		if dep == "true" {
			deprecated = true
			// ignore these too, but still need them to detect
			// gaps in the zebra:id fields
		}

		// check zebra
		zebra := alltags.Get("zid")
		if zebra != "" {
			// must be a non-negative number
			id, err := strconv.Atoi(zebra)
			if err != nil || id < 0 {
				where := ""
				if len(f.Names) > 0 {
					where = " on '" + f.Names[0].Name + "'"
				}
				err2 := fmt.Errorf("bad `zid` tag%s, could not convert"+
					" '%v' to non-zero integer: %v", where, zebra, err)
				fatalf(err2.Error())
				return nil, err2
			}
			zebraId = int64(id)
			//fmt.Printf("\n we see zebraId: %v\n", zebraId)
		}

	}

	ex, err := fs.parseExpr(f.Type)
	if err != nil {
		fatalf(err.Error())
		return nil, err
	}
	if ex == nil {
		skip = true
		// struct{} type fields, must track for zid checking.
		// so we can't return early here.
	}

	sf[0].Deprecated = deprecated
	sf[0].OmitEmpty = omitempty
	sf[0].ZebraId = zebraId
	sf[0].Skip = skip

	// parse field name
	switch len(f.Names) {
	case 0:
		sf[0].FieldName = embedded(f.Type)
	case 1:
		sf[0].FieldName = f.Names[0].Name
	default:
		// this is for a multiple in-line declaration,
		// e.g. type A struct { One, Two int }
		sf = sf[0:0]
		for _, nm := range f.Names {
			sf = append(sf, gen.StructField{
				FieldTag:   nm.Name,
				FieldName:  nm.Name,
				FieldElem:  ex.Copy(),
				OmitEmpty:  omitempty,
				Deprecated: deprecated,
				ZebraId:    zebraId,
				Skip:       skip,
			})
		}
		return sf, nil
	}
	sf[0].FieldElem = ex
	if sf[0].FieldTag == "" {
		sf[0].FieldTag = sf[0].FieldName
	}

	// validate extension
	if extension {
		switch ex := ex.(type) {
		case *gen.Ptr:
			if b, ok := ex.Value.(*gen.BaseElem); ok {
				b.Value = gen.Ext
			} else {
				warnln("couldn't cast to extension.")
				return nil, nil
			}
		case *gen.BaseElem:
			ex.Value = gen.Ext
		default:
			warnln("couldn't cast to extension.")
			return nil, nil
		}
	}
	return sf, nil
}

// extract embedded field name
//
// so, for a struct like
//
//	type A struct {
//		io.Writer
//  }
//
// we want "Writer"
func embedded(f ast.Expr) string {
	switch f := f.(type) {
	case *ast.Ident:
		return f.Name
	case *ast.StarExpr:
		return embedded(f.X)
	case *ast.SelectorExpr:
		return f.Sel.Name
	default:
		// other possibilities are disallowed
		return ""
	}
}

// stringify a field type name
func stringify(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return "*" + stringify(e.X)
	case *ast.SelectorExpr:
		return stringify(e.X) + "." + e.Sel.Name
	case *ast.ArrayType:
		if e.Len == nil {
			return "[]" + stringify(e.Elt)
		}
		return fmt.Sprintf("[%s]%s", stringify(e.Len), stringify(e.Elt))
	case *ast.InterfaceType:
		if e.Methods == nil || e.Methods.NumFields() == 0 {
			return "interface{}"
		}
	}
	return "<BAD>"
}

// recursively translate ast.Expr to gen.Elem; nil means type not supported
// expected input types:
// - *ast.MapType (map[T]J)
// - *ast.Ident (name)
// - *ast.ArrayType ([(sz)]T)
// - *ast.StarExpr (*T)
// - *ast.StructType (struct {})
// - *ast.SelectorExpr (a.B)
// - *ast.InterfaceType (interface {})
func (fs *FileSet) parseExpr(e ast.Expr) (gen.Elem, error) {
	switch e := e.(type) {

	case *ast.MapType:
		if k, ok := e.Key.(*ast.Ident); ok && k.Name == "string" {
			in, err := fs.parseExpr(e.Value)
			panicOn(err)
			if in != nil {
				return &gen.Map{Value: in, KeyTyp: "String", KeyDeclTyp: "string"}, nil
			}
		}

		// support int64/int32/int keys
		if k, ok := e.Key.(*ast.Ident); ok {
			in, err := fs.parseExpr(e.Value)
			if err != nil {
				fatalf(err.Error())
			}
			if in != nil {
				switch k.Name {
				case "int64":
					return &gen.Map{Value: in, KeyTyp: "Int64", KeyDeclTyp: "int64"}, nil
				case "int32":
					return &gen.Map{Value: in, KeyTyp: "Int32", KeyDeclTyp: "int32"}, nil
				case "int":
					return &gen.Map{Value: in, KeyTyp: "Int", KeyDeclTyp: "int"}, nil
				}
			}
		}

		return nil, nil

	case *ast.Ident:
		b := gen.Ident(e.Name)

		// work to resove this expression
		// can be done later, once we've resolved
		// everything else.
		if b.Value == gen.IDENT {
			if _, ok := fs.Specs[e.Name]; !ok {
				warnf("non-local identifier: %s\n", e.Name)
			}
		}
		return b, nil

	case *ast.ArrayType:

		// special case for []byte
		if e.Len == nil {
			if i, ok := e.Elt.(*ast.Ident); ok && i.Name == "byte" {
				return &gen.BaseElem{Value: gen.Bytes}, nil
			}
		}

		// return early if we don't know
		// what the slice element type is
		els, err := fs.parseExpr(e.Elt)
		if err != nil {
			return nil, err
		}
		if els == nil {
			return nil, nil
		}

		// array and not a slice
		if e.Len != nil {
			switch s := e.Len.(type) {
			case *ast.BasicLit:
				return &gen.Array{
					Size: s.Value,
					Els:  els,
				}, nil

			case *ast.Ident:
				return &gen.Array{
					Size: s.String(),
					Els:  els,
				}, nil

			case *ast.SelectorExpr:
				return &gen.Array{
					Size: stringify(s),
					Els:  els,
				}, nil

			default:
				return nil, nil
			}
		}
		return &gen.Slice{Els: els}, nil

	case *ast.StarExpr:
		v, err := fs.parseExpr(e.X)
		if err != nil {
			return nil, err
		}
		if v != nil {
			return &gen.Ptr{Value: v}, nil
		}
		return nil, nil

	case *ast.StructType:
		fields, err := fs.parseFieldList(e.Fields)
		if err != nil {
			return nil, err
		}
		if len(fields) > 0 {
			return &gen.Struct{Fields: fields}, nil
		}
		return nil, nil

	case *ast.SelectorExpr:
		return gen.Ident(stringify(e)), nil

	case *ast.InterfaceType:
		// support `interface{}`
		if len(e.Methods.List) == 0 {
			return &gen.BaseElem{Value: gen.Intf}, nil
		}
		return nil, nil

	default: // other types not supported
		return nil, nil
	}
}

func infof(s string, v ...interface{}) {
	pushstate(s)
	fmt.Printf(chalk.Green.Color(strings.Join(logctx, ": ")), v...)
	popstate()
}

func infoln(s string) {
	pushstate(s)
	fmt.Println(chalk.Green.Color(strings.Join(logctx, ": ")))
	popstate()
}

func warnf(s string, v ...interface{}) {
	pushstate(s)
	fmt.Printf(chalk.Yellow.Color(strings.Join(logctx, ": ")), v...)
	popstate()
}

func warnln(s string) {
	pushstate(s)
	fmt.Println(chalk.Yellow.Color(strings.Join(logctx, ": ")))
	popstate()
}

func fatalf(s string, v ...interface{}) {
	pushstate(s)
	fmt.Printf(chalk.Red.Color(strings.Join(logctx, ": ")), v...)
	popstate()
}

var logctx []string

// push logging state
func pushstate(s string) {
	logctx = append(logctx, s)
}

// pop logging state
func popstate() {
	logctx = logctx[:len(logctx)-1]
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
