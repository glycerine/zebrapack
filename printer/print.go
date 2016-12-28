package printer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/glycerine/zebrapack/cfg"
	"github.com/glycerine/zebrapack/gen"
	"github.com/glycerine/zebrapack/msgp"
	"github.com/glycerine/zebrapack/parse"
	"golang.org/x/tools/imports"
)

func infof(s string, v ...interface{}) {
	fmt.Printf(s, v...)
}

// PrintFile prints the methods for the provided list
// of elements to the given file name and canonical
// package path.
func PrintFile(
	file string,
	f *parse.FileSet,
	mode gen.Method,
	cfg *cfg.ZebraConfig,
	pathToGoSource string) error {

	out, tests, err := generate(f, mode, cfg)
	if err != nil {
		return err
	}

	if !cfg.NoEmbeddedSchema {
		// add the serialized msgpack2 zebra schema
		tr, err := parse.TranslateToZebraSchema(pathToGoSource, f)
		if err != nil {
			return err
		}
		//fmt.Printf("tr = %#v\n", tr)
		sby, err := tr.MarshalMsg(nil)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(out, "\n// ZebraSchemaInMsgpack2Format "+
			"provides the ZebraPack Schema in msgpack2 format, length "+
			"%v bytes\nvar ZebraSchemaInMsgpack2Format = []byte{",
			len(sby))
		if err != nil {
			return err
		}
		for i := range sby {
			if i%16 == 0 {
				fmt.Fprintf(out, "\n")
			}
			_, err = fmt.Fprintf(out, "0x%x,", sby[i])
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(out, "\n}\n\n")
		if err != nil {
			return err
		}

		// store as json and pretty-printed too
		buf := bytes.NewBuffer(sby)
		var compactJson, pretty bytes.Buffer
		_, err = msgp.CopyToJSON(&compactJson, buf)
		if err != nil {
			return err
		}

		jby := compactJson.Bytes()
		_, err = fmt.Fprintf(out, "\n// ZebraSchemaInJsonCompact provides "+
			"the ZebraPack schema in compact JSON format, length %v bytes\n"+
			"var ZebraSchemaInJsonCompact = `%s`\n",
			len(jby), string(jby))
		if err != nil {
			return err
		}

		err = json.Indent(&pretty, compactJson.Bytes(), "", "    ")
		if err != nil {
			return err
		}

		pby := pretty.Bytes()
		_, err = fmt.Fprintf(out, "\n// ZebraSchemaInJsonPretty provides"+
			" the ZebraPack schema in pretty JSON, length %v bytes\nvar "+
			"ZebraSchemaInJsonPretty = `%s`\n",
			len(pby), string(pby))
		if err != nil {
			return err
		}
	}

	// we'll run goimports on the main file
	// in another goroutine, and run it here
	// for the test file. empirically, this
	// takes about the same amount of time as
	// doing them in serial when GOMAXPROCS=1,
	// and faster otherwise.
	res := goformat(file, out.Bytes())
	if tests != nil {
		testfile := strings.TrimSuffix(file, ".go") + "_test.go"
		err = format(testfile, tests.Bytes())
		if err != nil {
			return err
		}
		infof(">>> Wrote and formatted \"%s\"\n", testfile)
	}
	err = <-res
	if err != nil {
		return err
	}
	return nil
}

func format(file string, data []byte) error {
	out, err := imports.Process(file, data, nil)
	if err != nil {
		fmt.Printf("\n\n debug: problem file:\n%s\n", file)
		ioutil.WriteFile(file, data, 0600)
		return err
	}
	return ioutil.WriteFile(file, out, 0600)
}

func goformat(file string, data []byte) <-chan error {
	out := make(chan error, 1)
	go func(file string, data []byte, end chan error) {
		end <- format(file, data)
		infof(">>> Wrote and formatted \"%s\"\n", file)
	}(file, data, out)
	return out
}

func dedupImports(imp []string) []string {
	m := make(map[string]struct{})
	for i := range imp {
		m[imp[i]] = struct{}{}
	}
	r := []string{}
	for k := range m {
		r = append(r, k)
	}
	return r
}

func generate(f *parse.FileSet, mode gen.Method, cfg *cfg.ZebraConfig) (*bytes.Buffer, *bytes.Buffer, error) {
	outbuf := bytes.NewBuffer(make([]byte, 0, 4096))
	writePkgHeader(outbuf, f.Package)

	myImports := []string{"fmt"}
	myImports = append(myImports, "github.com/glycerine/zebrapack/msgp")
	for _, imp := range f.Imports {
		if imp.Name != nil {
			// have an alias, include it.
			myImports = append(myImports, imp.Name.Name+` `+imp.Path.Value)
		} else {
			myImports = append(myImports, imp.Path.Value)
		}
	}
	dedup := dedupImports(myImports)
	writeImportHeader(outbuf, dedup...)

	var testbuf *bytes.Buffer
	var testwr io.Writer
	if mode&gen.Test == gen.Test {
		testbuf = bytes.NewBuffer(make([]byte, 0, 4096))
		writePkgHeader(testbuf, f.Package)
		if mode&(gen.Encode|gen.Decode) != 0 {
			writeImportHeader(testbuf, "bytes", "github.com/glycerine/zebrapack/msgp", "testing")
		} else {
			writeImportHeader(testbuf, "github.com/glycerine/zebrapack/msgp", "testing")
		}
		testwr = testbuf
	}
	return outbuf, testbuf, f.PrintTo(gen.NewPrinter(mode, outbuf, testwr, f.Cfg))
}

func writePkgHeader(b *bytes.Buffer, name string) {
	b.WriteString("package ")
	b.WriteString(name)
	b.WriteByte('\n')
	b.WriteString("// NOTE: THIS FILE WAS PRODUCED BY THE\n// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)\n// DO NOT EDIT\n\n")
}

func writeImportHeader(b *bytes.Buffer, imports ...string) {
	b.WriteString("import (\n")
	for _, im := range imports {
		if im[len(im)-1] == '"' {
			// support aliased imports
			fmt.Fprintf(b, "\t%s\n", im)
		} else {
			fmt.Fprintf(b, "\t%q\n", im)
		}
	}
	b.WriteString(")\n\n")
}
