package zebra

//go:generate msgp

// Zcat describes the basic type category of the field.
// Implementation note: it should correspond to gen/Elem.
type Zcat uint8

const (
	InvalidCat Zcat = iota
	BaseElemCat
	MapCat
	StructCat
	SliceCat
	ArrayCat
	PtrCat
)

// Zprimitive describes the detailed type of the field
// Implementation note: this must correspond to
// gen/Primitive when Zcat is BaseElemCat.
type Zprimitive uint8

const (
	Invalid Zprimitive = iota
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

// ZebraSchema is the top level container
type Schema struct {
	SourcePath    string // reflect.TypeOf().PkgPath()
	ZebraSchemaId int64
	Structs       []Struct
}

// Struct represents a single message or struct.
type Struct struct {
	StructName string  // name of struct
	Fields     []Field // fields
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
	FieldCategory  Zcat       // will be InvalidCat if Skip is true
	FieldPrimitive Zprimitive // avail if FieldCategory == BaseElemCat

	// if OmitEmpty then we don't serialize
	// the field if it has its zero-value.
	OmitEmpty bool

	// Skip means either deprecated or type struct{}
	// or marked  as `msg:"-"`. In any case,
	// if Skip is true: do not serialize
	// this field when writing, and ignore it
	// when reading.
	Skip bool
}
