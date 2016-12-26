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

	WriteSchema     string
	GenSchemaId     bool
	UseZid          bool
	ReadStringsFast bool
	SchemaToGo      string
}

// call DefineFlags before myflags.Parse()
func (c *ZebraConfig) DefineFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Out, "o", "", "output file")
	fs.StringVar(&c.GoFile, "file", "", "input file")
	fs.BoolVar(&c.Encode, "io", true, "create Encode and Decode methods")
	fs.BoolVar(&c.Marshal, "marshal", true, "create Marshal and Unmarshal methods")
	fs.BoolVar(&c.Tests, "tests", true, "create tests and benchmarks")
	fs.BoolVar(&c.Unexported, "unexported", false, "also process unexported types")

	fs.StringVar(&c.WriteSchema, "write-schema", "", "write schema header to this file; - for stdout")
	fs.BoolVar(&c.GenSchemaId, "genid", false, "generate a fresh random zebraSchemaId64 value for your schema")
	fs.BoolVar(&c.UseZid, "fast", false, "for speed and type safety, instead of writing field names in structs, write the numeric zid into the msgpack. See also -write-schema to generate an external schema description to read/write in other languages.")
	fs.BoolVar(&c.ReadStringsFast, "fast-strings", false, "for speed when reading a string in a message that won't be reused, this flag means we'll use unsafe to cast the string header and avoid allocation.")
	fs.StringVar(&c.SchemaToGo, "schema-to-go", "", "path to schema that will be converted to Go and written on stdout")
}

// call c.ValidateConfig() after myflags.Parse()
func (c *ZebraConfig) ValidateConfig() error {
	return nil
}
