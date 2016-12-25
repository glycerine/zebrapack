package zebra

import (
	"fmt"
	"strings"
)

//go:generate msgp

// Zkind describes the detailed type of the field
// implentation note: must correspond to gen/Primitive
type Zkind uint64

const (
	// primitives
	// implementation note: must correspond to gen/Primitive.
	Invalid    Zkind = 0
	Bytes      Zkind = 1 // []byte
	String     Zkind = 2
	Float32    Zkind = 3
	Float64    Zkind = 4
	Complex64  Zkind = 5
	Complex128 Zkind = 6
	Uint       Zkind = 7 // 32 or 64 bit; as in Go, matches native word
	Uint8      Zkind = 8
	Uint16     Zkind = 9
	Uint32     Zkind = 10
	Uint64     Zkind = 11
	Byte       Zkind = 12
	Int        Zkind = 13 // as in Go, matches native word size.
	Int8       Zkind = 14
	Int16      Zkind = 15
	Int32      Zkind = 16
	Int64      Zkind = 17
	Bool       Zkind = 18
	Intf       Zkind = 19 // interface{}
	Time       Zkind = 20 // time.Time
	Ext        Zkind = 21 // extension
	IDENT      Zkind = 22 // IDENT means an unrecognized identifier

	// compound types
	// implementation note: should correspond to gen/Elem.
	InvalidCat  Zkind = 23
	BaseElemCat Zkind = 24
	MapCat      Zkind = 25
	StructCat   Zkind = 26
	SliceCat    Zkind = 27
	ArrayCat    Zkind = 28
	PtrCat      Zkind = 29
)

// KS is a building block for Ztype.
type KS struct {
	Kind Zkind
	Str  string
}

// Ztype describes any type, be it a BaseElem,
// Map, Struct, Slice, Array, or Ptr.
type Ztype struct {
	Name KS

	// key for maps. elem for ptr, slice, array.
	// Invalid when Name.Kind is < 23 (for a primitive).
	Domain KS

	// value for maps. otherwise typically zero (Invalid).
	// For an array, holds the fixed size.
	Range KS
}

// ZebraSchema is the top level container
type Schema struct {
	SourcePath    string
	SourcePackage string
	ZebraSchemaId int64
	Structs       []Struct
}

// Struct represents a single message or struct.
type Struct struct {
	StructName string
	Fields     []Field
}

// Field represents fields within a struct.
type Field struct {

	// Zid locates update collisions and ease resolution.
	//
	// Both follow Cap'nProto semantics: start at numbering at 0,
	// and strictly/always increase numbers monotically.
	//
	// No gaps and no duplicate Zid are allowed.
	//
	// Duplicate numbers are how collisions (between two
	// developers adding two distinct new fields independently
	// but at the same time) are detected.
	//
	// Therefore this ironclad rule: never delete a field or Zid number,
	// just mark it as deprecated with the `deprecated:"true"`
	// tag, and change its Go type to struct{}.
	//
	Zid          int64
	FieldGoName  string
	FieldTagName string

	// type info
	FieldTypeStr   string
	FieldCategory  Zkind // will be InvalidCat if Skip is true
	FieldPrimitive Zkind // avail if FieldCategory == BaseElemCat
	FieldFullType  Ztype

	// if OmitEmpty then we don't serialize
	// the field if it has its zero-value.
	OmitEmpty bool

	// Skip means either type struct{} or
	// other unserializable type;
	// or marked  as `msg:"-"`. In any case,
	// if Skip is true: do not serialize
	// this field when writing, and ignore it
	// when reading.
	Skip bool

	// Deprecated means tagged with `deprecated:"true"`
	// Compilers/libraries should discourage and warn
	// users away from writing to such fields, while
	// not making it impossible to either read or write
	// the field.
	Deprecated bool
}

func ZkindFromString(s string) Zkind {
	s = strings.ToLower(s)
	switch s {
	case "":
		return Invalid
	case "invalid":
		return Invalid
	case "bytes":
		return Bytes
	case "string":
		return String
	case "float32":
		return Float32
	case "float64":
		return Float64
	case "complex64":
		return Complex64
	case "complex128":
		return Complex128
	case "uint":
		return Uint
	case "uint8":
		return Uint8
	case "uint16":
		return Uint16
	case "uint32":
		return Uint32
	case "uint64":
		return Uint64
	case "byte":
		return Byte
	case "int":
		return Int
	case "int8":
		return Int8
	case "int16":
		return Int16
	case "int32":
		return Int32
	case "int64":
		return Int64
	case "bool":
		return Bool
	case "intf":
		return Intf
	case "time":
		return Time
	case "ext":
		return Ext
	case "ident":
		return IDENT
	case "invalidcat":
		return InvalidCat
	case "baseelemcat":
		return BaseElemCat
	case "mapcat":
		return MapCat
	case "structcat":
		return StructCat
	case "slicecat":
		return SliceCat
	case "arraycat":
		return ArrayCat
	case "ptrcat":
		return PtrCat
	}
	panic(fmt.Errorf("unrecognized arg '%s' to ZkindFromString()", s))
}

func (i Zkind) String() string {
	switch i {
	case Invalid:
		return ""
	case Bytes:
		return "bytes"
	case String:
		return "string"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Complex64:
		return "complex64"
	case Complex128:
		return "complex128"
	case Uint:
		return "uint"
	case Uint8:
		return "uint8"
	case Uint16:
		return "uint16"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"
	case Byte:
		return "byte"
	case Int:
		return "int"
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Bool:
		return "bool"

		// compound/non-primitives are uppercased
		// for readability
	case Intf:
		return "Intf"
	case Time:
		return "Time"
	case Ext:
		return "Ext"
	case IDENT:
		return "IDENT"
	case InvalidCat:
		return "InvalidCat"
	case BaseElemCat:
		return "BaseElemCat"
	case MapCat:
		return "MapCat"
	case StructCat:
		return "StructCat"
	case SliceCat:
		return "SliceCat"
	case ArrayCat:
		return "ArrayCat"
	case PtrCat:
		return "PtrCat"
	default:
		panic(fmt.Errorf("unrecognized Zkind value %#v", i))
	}
}
