/*
package zebra specifies the ZebraPack serialization format.
To bootstrap, the ZebraPack schema is itself stored
in msgpack2 (with optional JSON) format.
*/
package zebra

// in the generate command, we do not want -fast: we
// actually want to serialize our ZebraPack schema itself
// as simple msgpack2, rather than in ZebraPack format.

//go:generate zebrapack

// Zkind describes the detailed type of the field
// implentation note: must correspond to gen/Primitive
// Since it also stores the fixed size of a array type,
// it needs to be large. When serialized as msgpack2,
// it will be compressed.
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

	// IDENT means an unrecognized identifier;
	// it typically means a named struct type.
	IDENT Zkind = 22

	// compound types
	// implementation note: should correspond to gen/Elem.
	BaseElemCat Zkind = 23
	MapCat      Zkind = 24
	StructCat   Zkind = 25
	SliceCat    Zkind = 26
	ArrayCat    Zkind = 27
	PointerCat  Zkind = 28
)

// Ztype describes any type, be it a BaseElem,
// Map, Struct, Slice, Array, or Pointer.
type Ztype struct {

	// Kind gives the exact type for primitives,
	// and the category for compound types.
	Kind Zkind

	// Str holds the struct name when Kind == 22 (IDENT).
	// Otherwise it typically reflects Kind directly
	// which is useful for human readability.
	Str string `msg:",omitempty"`

	// Domain holds the key type for maps. For
	// pointers and slices it holds the element type.
	// For arrays, it holds the fixed size.
	// Domain is null when Kind is a primitive
	Domain *Ztype `msg:",omitempty"`

	// Range holds the value type for maps.
	// For arrays (always a fixed size), Range holds
	// the element type.  Otherwise Range is
	// typically null.
	Range *Ztype `msg:",omitempty"`
}

// Schema is the top level container
type Schema struct {
	SourcePath    string
	SourcePackage string
	ZebraSchemaId int64
	Structs       []Struct
	Imports       []string
}

// Struct represents a single message or struct.
type Struct struct {
	StructName string
	Fields     []Field
}

// Field represents fields within a struct.
type Field struct {

	// Zid numbering detects update collisions
	// when two developers simultaneously add two
	// new fields. Zid numbering allows sane
	// forward/backward data evolution, like protobufs
	// and Cap'nProto.
	//
	// Zid follows Cap'nProto numbering semantics:
	// start at numbering at 0, and strictly/always
	// increase numbers monotically.
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

	// Zid is the zebrapack id
	Zid int64

	// the name of the field in the Go schema/source file.
	FieldGoName string

	// optional wire-name designated by a
	// `msg:"tagname"` tag on the struct field.
	FieldTagName string

	// =======================
	// type info
	// =======================

	// human readable/Go type description
	FieldTypeStr string `msg:",omitempty"`

	// the broad category of this type. empty if Skip is true
	FieldCategory Zkind `msg:",omitempty"`

	// avail if FieldCategory == BaseElemCat
	FieldPrimitive Zkind `msg:",omitempty"`

	// the machine-parse-able type tree
	FieldFullType *Ztype `msg:",omitempty"`

	// =======================
	// field tag details:
	// =======================

	// if OmitEmpty then we don't serialize
	// the field if it has its zero-value.
	OmitEmpty bool `msg:",omitempty"`

	// Skip means either type struct{} or
	// other unserializable type;
	// or marked  as `msg:"-"`. In any case,
	// if Skip is true: do not serialize
	// this field when writing, and ignore it
	// when reading.
	Skip bool `msg:",omitempty"`

	// Deprecated means tagged with `deprecated:"true"`.
	// Compilers/libraries should discourage and warn
	// users away from writing to such fields, while
	// not making it impossible to either read or write
	// the field.
	Deprecated bool `msg:",omitempty"`
}
