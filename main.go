// msgp is a code generation tool for
// creating methods to serialize and de-serialize
// Go data structures to and from MessagePack.
//
// This package is targeted at the `go generate` tool.
// To use it, include the following directive in a
// go source file with types requiring source generation:
//
//     //go:generate msgp
//
// The go generate tool should set the proper environment variables for
// the generator to execute without any command-line flags. However, the
// following options are supported, if you need them:
//
//  -o = output file name (default is {input}_gen.go)
//  -file = input file name (or directory; default is $GOFILE, which is set by the `go generate` command)
//  -io = satisfy the `msgp.Decodable` and `msgp.Encodable` interfaces (default is true)
//  -marshal = satisfy the `msgp.Marshaler` and `msgp.Unmarshaler` interfaces (default is true)
//  -tests = generate tests and benchmarks (default is true)
//
// For more information, please read README.md, and the wiki at github.com/glycerine/zebrapack
//
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/glycerine/zebrapack/cfg"
	"github.com/glycerine/zebrapack/gen"
	"github.com/glycerine/zebrapack/parse"
	"github.com/glycerine/zebrapack/printer"

	"github.com/tinylib/msgp/msgp"
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
	fs, err := parse.File(c)
	if err != nil {
		return err
	}

	if len(fs.Identities) == 0 {
		fmt.Println("No types requiring code generation were found!")
		return nil
	} else {
		if c.WriteSchema != "" { // saveSchemaAsMsgpackToFile
			err := saveMsgpackFile(c.GoFile, c.WriteSchema, fs)
			if err != nil {
				panic(err)
			}
		}
	}

	return printer.PrintFile(newFilename(c.Out, c.GoFile, fs.Package), fs, mode, c)
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

func saveMsgpackFile(parsedPath, path string, fs *parse.FileSet) error {
	//fmt.Printf("\n saveMsgpackFile saving to: '%v'", path)

	var f *os.File
	var err error
	if path == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	w := msgp.NewWriter(f)
	defer w.Flush()

	tr, err := parse.TranslateToZebraSchema(parsedPath, fs)
	if err != nil {
		return err
	}
	err = tr.EncodeMsg(w)
	if err != nil {
		return err
	}
	if path != "-" {
		by, err := tr.MarshalMsg(nil)
		if err != nil {
			return err
		}
		fjson, err := os.Create(path + ".json")
		if err != nil {
			return err
		}
		defer fjson.Close()

		// and write out pretty json version too.
		buf := bytes.NewBuffer(by)
		var compactJson, pretty bytes.Buffer
		_, err = msgp.CopyToJSON(&compactJson, buf)
		if err != nil {
			return err
		}

		err = json.Indent(&pretty, compactJson.Bytes(), "", "    ")
		if err != nil {
			return err
		}
		_, err = io.Copy(fjson, &pretty)
		fmt.Fprintf(fjson, "\n")
		return err
	}
	return nil
}
