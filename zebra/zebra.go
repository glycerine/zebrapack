package zebra

import (
	"github.com/glycerine/zebrapack/msgp"
)

type Raw msgp.Raw

//go:generate zebrapack

// Zprimitive describes the basic type category of the field
// It should match gen/Primitive
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
type ZebraSchema struct {
	PkgPath string // reflect.TypeOf().PkgPath()

	Structs []StructT
}

// Struct represents a single message or struct.
type StructT struct {
	StructName string   // name of struct
	Fld        []FieldT // fields
}

// Field represents fields within a struct.
type FieldT struct {

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
	Zid    int64
	Nam    string
	Typ    Zprimitive
	TypStr string            // for debugging
	Tag    map[string]string `msg:",omitempty"`
}
