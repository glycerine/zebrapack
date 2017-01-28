// zebrapack is a code generation tool for
// creating methods to serialize and de-serialize
// Go data structures to and from ZebraPack (a
// schema-based serialization format that is derived
// from MessagePack2).
//
// This package is targeted at the `go generate` tool.
// To use it, include the following directive in a
// go source file with types requiring source generation:
//
//     //go:generate zebrapack
//
// The go generate tool should set the proper environment variables for
// the generator to execute without any command-line flags. However, the
// following options are supported, if you need them (See zebrapack -h):
//
//   $ zebrapack -h
//
//   Usage of zebrapack:
//
//  -msgp
//    	generate msgpack2 serializers instead of ZebraPack;
//      for backward compatiblity or serializing the zebra
//      schema itself.
//
//   -fast-strings
//     	for speed when reading a string in a message that won't be
//      reused, this flag means we'll use unsafe to cast the string
//      header and avoid allocation.
//
//   -file go generate
//     	input file (or directory); default is $GOFILE, which
//      is set by the go generate command.
//
//   -genid
//     	generate a fresh random zebraSchemaId64 value to
//      include in your Go source schema
//
//   -io
//     	create Encode and Decode methods (default true)
//
//   -marshal
//     	create Marshal and Unmarshal methods (default true)
//
//  -method-prefix string
//    	(optional) prefix that will be pre-prended to
//      the front of generated method names; useful when
//      you need to avoid namespace collisions, but the
//      generated tests will break/the msgp package
//      interfaces won't be satisfied.
//
//  -no-embedded-schema
//      don't embed the schema in the generated files
//
//  -no-structnames-onwire
//    	don't embed the name of the struct in the
//      serialized zebrapack. Skipping the embedded
//      struct names saves time and space and matches
//      what protocol buffers/thrift/capnproto/msgpack do.
//      You must know the type on the wire you expect;
//      or embed a type tag in one universal wrapper
//      struct. Embedded struct names are a feature
//      of ZebraPack to help with dynamic language
//      bindings.
//
//   -o string
//     	output file (default is {input_file}_gen.go
//
//   -schema-to-go string
//     	(standalone functionality) path to schema in msgpack2
//      format; we will convert it to Go, write the Go on stdout,
//      and exit immediately
//
//   -tests
//     	create tests and benchmarks (default true)
//
//   -unexported
//     	also process unexported types
//
//   -write-schema string
// 		write schema header to this file; - for stdout
//
//
// For more information, please read README.md, and the wiki at github.com/glycerine/zebrapack
//
package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/glycerine/zebrapack/cfg"
	"github.com/glycerine/zebrapack/gen"
	"github.com/glycerine/zebrapack/parse"
	"github.com/glycerine/zebrapack/printer"
	"github.com/glycerine/zebrapack/zebra"
)

func main() {
	myflags := flag.NewFlagSet("zebrapack", flag.ExitOnError)
	c := &cfg.ZebraConfig{}
	c.DefineFlags(myflags)

	err := myflags.Parse(os.Args[1:])
	err = c.ValidateConfig()
	if err != nil {
		fmt.Printf("zebrapack command line flag error: '%s'\n", err)
		os.Exit(1)
	}

	if c.GenSchemaId {
		var by [8]byte
		cryptorand.Read(by[:])
		n := binary.LittleEndian.Uint64(by[:])
		n &= 0x0001ffffffffffff // restrict to 53 bits so R and js work
		fmt.Printf("\n// This crypto-randomly generated zebraSchemaId64 is a 53-bit\n"+
			"// integer that identifies your namespace.\n"+
			"// Paste it into your Go source.\n"+
			"  const zebraSchemaId64 int64 = 0x%x // %v\n\n", n, n)
		os.Exit(0)
	}

	if c.SchemaToGo != "" {
		handleSchemaToGo(c)
		os.Exit(0)
	}

	// GOFILE is set by go generate
	if c.GoFile == "" {
		c.GoFile = os.Getenv("GOFILE")
		if c.GoFile == "" {
			fmt.Println("No file to parse.")
			os.Exit(1)
		}
	}

	var mode gen.Method
	if c.Encode {
		mode |= (gen.Encode | gen.Decode | gen.Size | gen.FieldsEmpty)
	}
	if c.Marshal {
		mode |= (gen.Marshal | gen.Unmarshal | gen.Size | gen.FieldsEmpty)
	}
	if c.Tests {
		mode |= gen.Test
	}

	if mode&^gen.Test == 0 {
		fmt.Println("No methods to generate; -io=false && -marshal=false")
		os.Exit(1)
	}

	if err := Run(mode, c); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Run writes all methods using the associated file or path, e.g.
//
//	err := msgp.Run("path/to/myfile.go", gen.Size|gen.Marshal|gen.Unmarshal|gen.Test, false)
//
func Run(mode gen.Method, c *cfg.ZebraConfig) error {
	if mode&^gen.Test == 0 {
		return nil
	}
	fmt.Println("======== ZebraPack Code Generator  =======")
	fmt.Printf(">>> Input: \"%s\"\n", c.GoFile)
	var fs *parse.FileSet
	var err error
	if c.NoLoad {
		fs, err = parse.FileNoLoad(c)
	} else {
		fs, err = parse.File(c)
	}
	if err != nil {
		return err
	}

	if len(fs.Identities) == 0 {
		fmt.Println("No types requiring code generation were found!")
		return nil
	}

	if c.WriteSchema != "" { // saveSchemaAsMsgpackToFile
		err := fs.SaveMsgpackFile(c.GoFile, c.WriteSchema)
		if err != nil {
			panic(err)
		}
	}

	return printer.PrintFile(newFilename(c.Out, c.GoFile, fs.Package), fs, mode, c, c.GoFile)
}

// picks a new file name based on input flags and input filename(s).
func newFilename(out, old, pkg string) string {
	if out != "" {
		if pre := strings.TrimPrefix(out, old); len(pre) > 0 &&
			!strings.HasSuffix(out, ".go") {
			return filepath.Join(old, out)
		}
		return out
	}

	if fi, err := os.Stat(old); err == nil && fi.IsDir() {
		old = filepath.Join(old, pkg)
	}
	// new file name is old file name + _gen.go
	return strings.TrimSuffix(old, ".go") + "_gen.go"
}

func fileExists(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}

func handleSchemaToGo(c *cfg.ZebraConfig) {
	if !fileExists(c.SchemaToGo) {
		fmt.Fprintf(os.Stderr, "error: -schema-to-go '%s' path not found\n", c.SchemaToGo)
		os.Exit(1)
	}
	by, err := ioutil.ReadFile(c.SchemaToGo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: -schema-to-go '%s' produced error on reading file: %v\n",
			c.SchemaToGo, err)
		os.Exit(1)
	}
	var sch zebra.Schema
	_, err = sch.UnmarshalMsg(by)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: -schema-to-go '%s' produced error on UnmarshalMsg: %v\n",
			c.SchemaToGo, err)
		os.Exit(1)
	}
	err = sch.WriteToGo(os.Stdout, c.SchemaToGo, "main")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: -schema-to-go '%s' produced error on UnmarshalMsg: %v\n",
			c.SchemaToGo, err)
		os.Exit(1)
	}
}
