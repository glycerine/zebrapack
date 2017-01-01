package cfg

import (
	"flag"
)

type ZebraConfig struct {
	Out        string
	GoFile     string
	Encode     bool
	Marshal    bool
	Tests      bool
	Unexported bool

	WriteSchema string
	GenSchemaId bool

	// for back-compatibity, generate msgpack2
	// encoders instead of ZebraPack
	UseMsgp2 bool

	ReadStringsFast  bool
	SchemaToGo       string
	NoEmbeddedSchema bool
	NoRTTI           bool
}

// call DefineFlags before myflags.Parse()
func (c *ZebraConfig) DefineFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Out, "o", "", "output file (default is {input_file}_gen.go")
	fs.StringVar(&c.GoFile, "file", "", "input file (or directory); default is $GOFILE, which is set by the `go generate` command.")
	fs.BoolVar(&c.Encode, "io", true, "create Encode and Decode methods")
	fs.BoolVar(&c.Marshal, "marshal", true, "create Marshal and Unmarshal methods")
	fs.BoolVar(&c.Tests, "tests", true, "create tests and benchmarks")
	fs.BoolVar(&c.Unexported, "unexported", false, "also process unexported types")

	fs.StringVar(&c.WriteSchema, "write-schema", "", "write schema (in msgpack2 format) to this file; - for stdout")
	fs.BoolVar(&c.GenSchemaId, "genid", false, "generate a fresh random zebraSchemaId64 value to include in your Go source schema")
	fs.BoolVar(&c.UseMsgp2, "msgp", false, "generate msgpack2 serializers instead of ZebraPack; for backward compatiblity and serializing the zebra.go schema itself.")
	fs.BoolVar(&c.ReadStringsFast, "fast-strings", false, "for speed when reading a string in a message that won't be reused, this flag means we'll use unsafe to cast the string header and avoid allocation.")
	fs.StringVar(&c.SchemaToGo, "schema-to-go", "", "(standalone functionality) path to schema in msgpack2 format; we will convert it to Go, write the Go on stdout, and exit immediately")
	fs.BoolVar(&c.NoEmbeddedSchema, "no-embedded-schema", false, "don't embed the schema in the generated files")
	fs.BoolVar(&c.NoRTTI, "no-rtti", false, "don't embed the name of the struct in the serialized zebrapack. Saves space and time, but you must know the type on the wire in advance.")
}

// call c.ValidateConfig() after myflags.Parse()
func (c *ZebraConfig) ValidateConfig() error {
	return nil
}
