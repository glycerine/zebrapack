package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

// PrettyPrint out the go source file we read in.
func (x *Extractor) PrettyPrint(fileSet *token.FileSet, astfile *ast.File, fn string) error {
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}

	printConfig := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 4}

	err = printConfig.Fprint(f, fileSet, astfile)
	if err != nil {
		panic(err)
	}
	return nil
}
