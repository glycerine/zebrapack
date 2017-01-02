package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

// PrettyPrint out the go source file we read in.
func (x *Extractor) PrettyPrint(output io.Writer, fileSet *token.FileSet, astfile *ast.File) error {

	printConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 4}

	err := printConfig.Fprint(output, fileSet, astfile)
	if err != nil {
		panic(err)
	}
	return nil
}
