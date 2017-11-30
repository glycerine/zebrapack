package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/glycerine/zebrapack/zebra"
)

// This code defines the type declaration tree.
//
// Consider the following:
//
// type Marshaler struct {
// 	  Thing1 *float64 `msg:"thing1"`
// 	  Body   []byte   `msg:"body"`
// }
//
// A parser using this generator as a backend
// should parse the above into:
//
// var val Elem = &Ptr{
// 	name: "z",
// 	Value: &Struct{
// 		Name: "Marshaler",
// 		Fields: []StructField{
// 			{
// 				FieldTag: "thing1",
// 				FieldElem: &Ptr{
// 					name: "z.Thing1",
// 					Value: &BaseElem{
// 						name:    "*z.Thing1",
// 						Value:   Float64,
//						Convert: false,
// 					},
// 				},
// 			},
// 			{
// 				FieldTag: "body",
// 				FieldElem: &BaseElem{
// 					name:    "z.Body",
// 					Value:   Bytes,
// 					Convert: false,
// 				},
// 			},
// 		},
// 	},
// }

// Base is one of the
// base types
type Primitive uint8

// this is effectively the
// list of currently available
// ReadXxxx / WriteXxxx methods.
const (
	Invalid Primitive = iota
	Bytes
	String
	Float32
	Float64
	Complex64
	Complex128
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Byte
	Int
	Int8
	Int16
	Int32
	Int64
	Bool
	Intf // interface{}
	Time // time.Time
	Ext  // extension

	IDENT // IDENT means an unrecognized identifier
)

// all of the recognized identities
// that map to primitive types
var primitives = map[string]Primitive{
	"[]byte":         Bytes,
	"string":         String,
	"float32":        Float32,
	"float64":        Float64,
	"complex64":      Complex64,
	"complex128":     Complex128,
	"uint":           Uint,
	"uint8":          Uint8,
	"uint16":         Uint16,
	"uint32":         Uint32,
	"uint64":         Uint64,
	"byte":           Byte,
	"int":            Int,
	"int8":           Int8,
	"int16":          Int16,
	"int32":          Int32,
	"int64":          Int64,
	"bool":           Bool,
	"interface{}":    Intf,
	"time.Time":      Time,
	"msgp.Extension": Ext,
}

// types built into the library
// that satisfy all of the
// interfaces.
var builtins = map[string]struct{}{
	"msgp.Raw":    struct{}{},
	"msgp.Number": struct{}{},
}

// Common data/methods for every Elem
type Common struct {
	vname string
	alias string

	hmp  HasMethodPrefix
	Skip bool
}

type HasMethodPrefix interface {
	MethodPrefix() string
}

func (c *Common) SkipMe() bool        { return c.Skip }
func (c *Common) SetVarname(s string) { c.vname = s }
func (c *Common) Varname() string     { return c.vname }
func (c *Common) Alias(typ string)    { c.alias = typ }
func (c *Common) hidden()             {}
func (c *Common) MethodPrefix() string {
	if c.hmp == nil {
		return ""
	}
	return c.hmp.MethodPrefix()
}
func (c *Common) SetHasMethodPrefix(hmp HasMethodPrefix) {
	c.hmp = hmp
}

func IsPrintable(e Elem) bool {
	if be, ok := e.(*BaseElem); ok && !be.Printable() {
		return false
	}
	return true
}

// Elem is a go type capable of being
// serialized into MessagePack. It is
// implemented by *Ptr, *Struct, *Array,
// *Slice, *Map, and *BaseElem.
type Elem interface {
	// SetVarname sets this nodes
	// variable name and recursively
	// sets the names of all its children.
	// In general, this should only be
	// called on the parent of the tree.
	SetVarname(s string)

	// Varname returns the variable
	// name of the element.
	Varname() string

	// TypeName is the canonical
	// go type name of the node
	// e.g. "string", "int", "map[string]float64"
	// OR the alias name, if it has been set.
	TypeName() string

	// Alias sets a type (alias) name
	Alias(typ string)

	// Copy should perform a deep copy of the object
	Copy() Elem

	// Complexity returns a measure of the
	// complexity of element (greater than
	// or equal to 1.)
	Complexity() int

	// ZeroLiteral returns the literal expression
	// needed to initialize an element named v to its
	// zero value. e.g. 0 for numbers, "" for strings,
	// or Truck{} for a struct Truck.
	ZeroLiteral(v string) string

	// GetZtype provides type info in a uniform way.
	GetZtype() zebra.Ztype

	// for template instantiation with custom method prefix
	MethodPrefix() string
	SetHasMethodPrefix(hmp HasMethodPrefix)
	IsInterface() bool
	IsInInterfaceSlice() bool
	SetIsInInterfaceSlice()

	SkipMe() bool

	hidden()
}

// Ident returns the *BaseElem that corresponds
// to the provided identity.
func Ident(id string, isIface bool) *BaseElem {
	p, ok := primitives[id]
	if ok {
		return &BaseElem{Value: p, isIface: isIface}
	}
	be := &BaseElem{Value: IDENT, isIface: isIface}
	be.Alias(id)
	return be
}

type Array struct {
	Common
	Index        string // index variable name
	SizeNamed    string // array size
	SizeResolved string // array size
	Els          Elem   // child
}

func (s *Array) IsInterface() bool {
	return false
}

func (s *Array) IsInInterfaceSlice() bool {
	return false
}

func (s *Array) SetIsInInterfaceSlice() {}

func (a *Array) ZeroLiteral(v string) string {
	b := a.Els.ZeroLiteral(fmt.Sprintf("%s[i]", v))
	return fmt.Sprintf(`for i := range %s { %s }`, v, b)
}

func (a *Array) SetVarname(s string) {
	a.Common.SetVarname(s)
ridx:
	a.Index = gensym()

	// try to avoid using the same
	// index as a parent slice
	if strings.Contains(a.Varname(), a.Index) {
		goto ridx
	}

	a.Els.SetVarname(fmt.Sprintf("%s[%s]", a.Varname(), a.Index))
}

func (a *Array) TypeName() string {
	if a.Common.alias != "" {
		return a.Common.alias
	}
	a.Common.Alias(fmt.Sprintf("[%s]%s", a.SizeNamed, a.Els.TypeName()))
	return a.Common.alias
}

func (a *Array) GetZtype() (r zebra.Ztype) {

	r.Kind = zebra.ArrayCat
	r.Str = r.Kind.String()

	zt := a.Els.GetZtype()
	r.Range = &zt

	// set Domain to be the size of the array
	r.Domain = &zebra.Ztype{}
	r.Domain.Str = a.SizeNamed
	//fmt.Printf("a is '%#v'\n", a)
	n, err := strconv.Atoi(a.SizeResolved)
	if err != nil {
		panic(err)
	}
	r.Domain.Kind = zebra.Zkind(n)
	return
}

func (a *Array) Copy() Elem {
	b := *a
	b.Els = a.Els.Copy()
	return &b
}

func (a *Array) Complexity() int { return 1 + a.Els.Complexity() }

// Map is a map[string]Elem
type Map struct {
	Common
	Keyidx string // key variable name
	Validx string // value variable name
	Value  Elem   // value element

	KeyTyp     string
	KeyDeclTyp string
}

func (s *Map) IsInterface() bool {
	return false
}

func (s *Map) IsInInterfaceSlice() bool {
	return false
}
func (m *Map) SetIsInInterfaceSlice() {}

func (m *Map) GetZtype() (r zebra.Ztype) {

	r.Kind = zebra.MapCat
	r.Str = r.Kind.String()

	r.Domain = &zebra.Ztype{}
	r.Domain.Kind = zebra.ZkindFromString(m.KeyTyp)
	r.Domain.Str = r.Domain.Kind.String()

	rng := m.Value.GetZtype()
	r.Range = &rng
	return
}

func (m *Map) ZeroLiteral(v string) string {
	template := `
	            if len(%s) > 0 {
				for key, _ := range %s {
					delete(%s, key)
				}}
`
	return fmt.Sprintf(template, v, v, v)
}

func (m *Map) SetVarname(s string) {
	m.Common.SetVarname(s)
ridx:
	m.Keyidx = gensym()
	m.Validx = gensym()

	// just in case
	if m.Keyidx == m.Validx {
		goto ridx
	}

	m.Value.SetVarname(m.Validx)
}

func (m *Map) TypeName() string {
	if m.Common.alias != "" {
		return m.Common.alias
	}
	m.Common.Alias("map[" + m.KeyDeclTyp + "]" + m.Value.TypeName())
	return m.Common.alias
}

func (m *Map) Copy() Elem {
	g := *m
	g.Value = m.Value.Copy()
	return &g
}

func (m *Map) Complexity() int { return 2 + m.Value.Complexity() }

type Slice struct {
	Common
	Index string
	Els   Elem // The type of each element
}

func (s *Slice) IsInterface() bool {
	return s.Els.IsInterface()
}

func (s *Slice) IsInInterfaceSlice() bool {
	return false
}

func (s *Slice) SetIsInInterfaceSlice() {
	s.Els.SetIsInInterfaceSlice()
}

func (a *Slice) ZeroLiteral(v string) string {
	return fmt.Sprintf(`%s = %s[:0]`, v, v)
}

func (s *Slice) SetVarname(a string) {
	s.Common.SetVarname(a)
	s.Index = gensym()
	varName := s.Varname()
	if varName[0] == '*' {
		// Pointer-to-slice requires parenthesis for slicing.
		varName = "(" + varName + ")"
	}
	s.Els.SetVarname(fmt.Sprintf("%s[%s]", varName, s.Index))
}

func (s *Slice) GetZtype() (r zebra.Ztype) {
	r.Kind = zebra.SliceCat
	r.Str = r.Kind.String()

	dom := s.Els.GetZtype()
	r.Domain = &dom
	return
}

func (s *Slice) TypeName() string {
	if s.Common.alias != "" {
		return s.Common.alias
	}
	s.Common.Alias("[]" + s.Els.TypeName())
	return s.Common.alias
}

func (s *Slice) Copy() Elem {
	z := *s
	z.Els = s.Els.Copy()
	return &z
}

func (s *Slice) Complexity() int {
	return 1 + s.Els.Complexity()
}

type Ptr struct {
	Common
	Value Elem
}

func (s *Ptr) IsInterface() bool {
	return false
}

func (s *Ptr) IsInInterfaceSlice() bool {
	return false
}

func (s *Ptr) SetIsInInterfaceSlice() {}

func (s *Ptr) GetZtype() (r zebra.Ztype) {
	r.Kind = zebra.PointerCat
	r.Str = r.Kind.String()

	dom := s.Value.GetZtype()
	r.Domain = &dom
	return
}

func (s *Ptr) ZeroLiteral(v string) string {
	switch x := s.Value.(type) {
	case *Struct:
		if x.Common.alias != "" {
			return fmt.Sprintf(" _, err = %v.UnmarshalMsg(msgp.OnlyNilSlice); if err != nil { return };\n", v) // , s.TypeName())
		} else {
			return fmt.Sprintf("\n// what Ptr.ZeroLiteral(v='%v') goes here?????\n", v)
		}
	}
	return fmt.Sprintf("\n// Ptr.ZeroLiteral() was not a pointer to Struct, was %T/val=%#v \n %s = nil", s, s, v)
}

func (s *Ptr) SetVarname(a string) {
	s.Common.SetVarname(a)

	// struct fields are dereferenced
	// automatically...
	switch x := s.Value.(type) {
	case *Struct:
		// struct fields are automatically dereferenced
		x.SetVarname(a)
		return

	case *BaseElem:
		// identities have pointer receivers
		if x.Value == IDENT {
			x.SetVarname(a)
		} else {
			x.SetVarname("*" + a)
		}
		return

	default:
		s.Value.SetVarname("*" + a)
		return
	}
}

func (s *Ptr) TypeName() string {
	if s.Common.alias != "" {
		return s.Common.alias
	}
	s.Common.Alias("*" + s.Value.TypeName())
	return s.Common.alias
}

func (s *Ptr) Copy() Elem {
	v := *s
	v.Value = s.Value.Copy()
	return &v
}

func (s *Ptr) Complexity() int { return 1 + s.Value.Complexity() }

func (s *Ptr) Needsinit() bool {
	if be, ok := s.Value.(*BaseElem); ok && be.needsref {
		return false
	}
	return true
}

type Struct struct {
	Common
	Fields           []StructField // field list
	AsTuple          bool          // write as an array instead of a map
	hasOmitEmptyTags bool
	KeyTyp           string

	SkipCount int
}

func (s *Struct) IsInterface() bool {
	return false
}

func (s *Struct) IsInInterfaceSlice() bool {
	return false
}

func (s *Struct) SetIsInInterfaceSlice() {}

func (s *Struct) GetZtype() (r zebra.Ztype) {
	r.Kind = zebra.StructCat
	r.Str = r.Kind.String() // s.TypeName()
	return
}

func (s *Struct) ZeroLiteral(v string) string {
	return fmt.Sprintf(`%s = %s{}`, v, s.TypeName())
}

func (s *Struct) TypeName() string {
	if s.Common.alias != "" {
		return s.Common.alias
	}
	str := "struct{\n"
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		str += s.Fields[i].FieldName + " " + s.Fields[i].FieldElem.TypeName() + ";\n"
	}
	str += "}"
	s.Common.Alias(str)
	return s.Common.alias
}

func (s *Struct) SetVarname(a string) {
	s.Common.SetVarname(a)
	writeStructFields(s.Fields, a)
}

func (s *Struct) Copy() Elem {
	g := *s
	g.Fields = make([]StructField, len(s.Fields))
	copy(g.Fields, s.Fields)
	for i := range s.Fields {
		if s.Fields[i].Skip {
			// FieldElem will be nil for skipped fields.
			continue
		}
		g.Fields[i].FieldElem = s.Fields[i].FieldElem.Copy()
	}
	return &g
}

func (s *Struct) Complexity() int {
	c := 1
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		c += s.Fields[i].FieldElem.Complexity()
	}
	return c
}

type StructField struct {
	FieldTag   string // the string inside the `msg:""` tag
	FieldName  string // the name of the struct field
	FieldElem  Elem   // the field type
	OmitEmpty  bool   // if the tag `msg:",omitempty"` was found
	Deprecated bool   // if the tag `deprecated:"true"` was found
	Skip       bool   // if msg:"-" or field is type struct{}
	ShowZero   bool   // if msg:",showzero" tag was found.
	IsIface    bool   // the field type is an interface?

	// ZebraId defaults to -1, meaning not-tagged with a zebra id.
	// if ZebraId >= 0, then the tag `zebra:"N"` was found, with ZebraId == N.
	ZebraId int64
}

func (s *StructField) SkipMe() bool {
	return s.Skip
}

func (s *StructField) IsInterface() bool {
	return s.IsIface
}

// BaseElem is an element that
// can be represented by a primitive
// MessagePack type.
type BaseElem struct {
	Common
	ShimToBase     string    // shim to base type, or empty
	ShimFromBase   string    // shim from base type, or empty
	Value          Primitive // Type of element
	Convert        bool      // should we do an explicit conversion?
	mustinline     bool      // must inline; not printable
	needsref       bool      // needs reference for shim
	isIface        bool
	isInIfaceSlice bool
}

func (s *BaseElem) IsInterface() bool {
	return s.isIface
}

func (s *BaseElem) IsInInterfaceSlice() bool {
	return s.isInIfaceSlice
}

func (s *BaseElem) SetIsInInterfaceSlice() {
	s.isInIfaceSlice = true
}

func (s *BaseElem) GetZtype() (r zebra.Ztype) {
	r.Kind = zebra.Zkind(s.Value)
	if r.Kind != 22 {
		r.Str = r.Kind.String()
		return
	}
	r.Str = s.TypeName()
	return
}

func (s *BaseElem) Printable() bool { return !s.mustinline }

func (s *BaseElem) Alias(typ string) {
	s.Common.Alias(typ)
	if s.Value != IDENT {
		s.Convert = true
	}
	if strings.Contains(typ, ".") {
		s.mustinline = true
	}
}

func (s *BaseElem) SetVarname(a string) {
	// extensions whose parents
	// are not pointers need to
	// be explicitly referenced
	if s.Value == Ext || s.needsref {
		if strings.HasPrefix(a, "*") {
			s.Common.SetVarname(a[1:])
			return
		}
		s.Common.SetVarname("&" + a)
		return
	}

	s.Common.SetVarname(a)
}

// TypeName returns the syntactically correct Go
// type name for the base element.
func (s *BaseElem) TypeName() string {
	if s.Common.alias != "" {
		return s.Common.alias
	}
	s.Common.Alias(s.BaseType())
	return s.Common.alias
}

// ToBase, used if Convert==true, is used as tmp = {{ToBase}}({{Varname}})
func (s *BaseElem) ToBase() string {
	if s.ShimToBase != "" {
		return s.ShimToBase
	}
	return s.BaseType()
}

// FromBase, used if Convert==true, is used as {{Varname}} = {{FromBase}}(tmp)
func (s *BaseElem) FromBase() string {
	if s.ShimFromBase != "" {
		return s.ShimFromBase
	}
	return s.TypeName()
}

// BaseName returns the string form of the
// base type (e.g. Float64, Ident, etc)
func (s *BaseElem) BaseName() string {
	// time is a special case;
	// we strip the package prefix
	if s.Value == Time {
		return "Time"
	}
	return s.Value.String()
}

func (s *BaseElem) BaseType() string {
	switch s.Value {
	case IDENT:
		return s.TypeName()

	// exceptions to the naming/capitalization
	// rule:
	case Intf:
		return "interface{}"
	case Bytes:
		return "[]byte"
	case Time:
		return "time.Time"
	case Ext:
		return "msgp.Extension"

	// everything else is base.String() with
	// the first letter as lowercase
	default:
		return strings.ToLower(s.BaseName())
	}
}

func (s *BaseElem) Needsref(b bool) {
	s.needsref = b
}

func (s *BaseElem) Copy() Elem {
	g := *s
	return &g
}

func (s *BaseElem) Complexity() int {
	if s.Convert && !s.mustinline {
		return 2
	}
	// we need to return 1 if !printable(),
	// in order to make sure that stuff gets
	// inlined appropriately
	return 1
}

// Resolved returns whether or not
// the type of the element is
// a primitive or a builtin provided
// by the package.
func (s *BaseElem) Resolved() bool {
	if s.Value == IDENT {
		_, ok := builtins[s.TypeName()]
		return ok
	}
	return true
}

func (s *BaseElem) ZeroLiteral(v string) string {
	switch s.Value {
	case String:
		return fmt.Sprintf(`%s = ""`, v)
	case Float32, Float64, Complex64,
		Complex128, Uint, Uint8, Uint16,
		Uint32, Uint64, Byte, Int, Int8,
		Int16, Int32, Int64:
		return fmt.Sprintf(`%s = 0`, v)
	case Bool:
		return fmt.Sprintf(`%s = false`, v)
	case Intf:
		return fmt.Sprintf(`%s = nil`, v)
	case Time:
		return fmt.Sprintf(`%s = time.Time{}`, v)
	case IDENT:
		return fmt.Sprintf(`%s = %s{}`, v, s.TypeName())
	case Bytes:
		return fmt.Sprintf(`%s = %s[:0]`, v, v)
	default:
		return fmt.Sprintf("/*don't know how to zero value '%s' from BaseElem:'%#v'.*/", s.Value, s)
	}
}

func (k Primitive) String() string {
	switch k {
	case String:
		return "String"
	case Bytes:
		return "Bytes"
	case Float32:
		return "Float32"
	case Float64:
		return "Float64"
	case Complex64:
		return "Complex64"
	case Complex128:
		return "Complex128"
	case Uint:
		return "Uint"
	case Uint8:
		return "Uint8"
	case Uint16:
		return "Uint16"
	case Uint32:
		return "Uint32"
	case Uint64:
		return "Uint64"
	case Byte:
		return "Byte"
	case Int:
		return "Int"
	case Int8:
		return "Int8"
	case Int16:
		return "Int16"
	case Int32:
		return "Int32"
	case Int64:
		return "Int64"
	case Bool:
		return "Bool"
	case Intf:
		return "Intf"
	case Time:
		return "time.Time"
	case Ext:
		return "Extension"
	case IDENT:
		return "Ident"
	default:
		return "INVALID"
	}
}

// writeStructFields is a trampoline for writeBase for
// all of the fields in a struct
func writeStructFields(s []StructField, name string) {
	for i := range s {
		if !s[i].Skip {
			s[i].FieldElem.SetVarname(fmt.Sprintf("%s.%s", name, s[i].FieldName))
		}
	}
}
