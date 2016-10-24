/*
structure of a ZebraPack file:

+---------------------------------------+
| SchemaT, msgpack2 serialized (a map). |
|  variable number of bytes             |
|                                       |
+---------------------------------------+
                 +
+---------------------------------------+}
|      Ztype (= StructInst)             | }
+---------------------------------------+  }
|      StructTypeId                     |   }
+---------------------------------------+    }- one of these for each struct on the wire.
|                                       |   }    These are ZPacket.
| Msgpack2 map, with integer keys       |  }
|                                       | }
+---------------------------------------+}


*/
package zebra

import (
	"github.com/glycerine/zebrapack/msgp"
)

type Raw msgp.Raw

//go:generate zebrapack

// ZPacket is the on-the-wire format for dat
type ZPacket struct {
	Ty Zkind              // type. e.g. StructInst, a
	Id StructTypeId       // TypeId, the specific type of struct.
	Da map[int64]msgp.Raw // data, keys are FieldId
}

// Zkind describes the basic type category of the field
type Zkind int32

const (
	Nil Zkind = iota // nil is a msgpack type, encoded as 0xc0
	Bool
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String // must be utf8
	Bytes  // []byte
	Array
	Map

	// metadata, but possibly not extensions.
	StructInst
	ZebraSchema

	// extensions
	Complex64  // per tinylib/msgp, extension 3
	Complex128 // per tinylib/msgp, extension 4
	Timestamp  // per tinylib/msgp, extension 5, time.Time as RFC3339 string

	// possible extensions
	InterfaceType // extension 8
	ZError        // extension 9
)

// Schema is the top level container for an ordered
// set of Structs and Interfaces.
type SchemaT struct {
	PkgPath string // reflect.TypeOf().PkgPath()

	// we use maps here so that we can publish
	// partial SchemaT and have them be cumulative
	// and without conflict. Just use different
	// integer keys to differentiate. keys for
	// the Structs should match the StructInstance.StructId,
	// likewise the integer key for Interfaces should
	// be referenced by the InterfaceInstance.IfaceId.
	Structs    map[int64]StructT
	Interfaces map[int64]InterfaceT
}

// Struct represents a single message or struct.
type StructT struct {
	Nam string   // name of struct
	Fld []FieldT // fields
}

// StructTypeId  precedes, on the wire, each
// occurance of a struct instance, identifying which previously
// published-to-the-wire record type is to follow.
//
// The StructId here is the key into the SchemaT.Structs map.
type StructTypeId int64

type InterfaceInstance struct {
	IfaceId int64
}

// Field represents fields within a struct.
type FieldT struct {

	// Field.Id locates update
	// collisions and eases resolution.
	// Both follow Cap'nProto semantics: start at numbering at 0,
	// increase numbers monotically, and no gaps or
	// duplicates are allowed.
	// Duplicate numbers are how collisions are detected.
	// Never delete a number, just mark it as deprecated and
	// change its type in the Go code to struct{}.
	FieldId int64

	Name string
	Ztyp Zkind
	Varg bool              `msg:",omitempty"` // is a var-arg ... input? only used in Method.Inputs
	Tag  map[string]string `msg:",omitempty"`
}

type MethodT struct {
	// NB no receiver element.
	// As in Go, these are methods on interfaces.

	MethodId int64
	Name     string
	Inputs   StructT
	Returns  StructT

	Deprecated bool
	Comment    string
}

type InterfaceT struct {
	Methods []MethodT

	Deprecated bool
	Comment    string
}
