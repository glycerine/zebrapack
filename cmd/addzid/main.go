/*
addzid automatically adds `zid:"0"`, `zid:"1"`, ... tags to your Go structs.

Given a set of golang (Go) source files, addzid will tag the public
struct fields with sequential zid tags. This prepares your source
so that it can be fed to the `zebrapack` codegen tool.

`addzid` was dervied from the author's `bambam` tool to support ZebraPack.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func use() {
	fmt.Fprintf(os.Stderr, "\nuse: addzid {-o outdir} myGoSourceFile.go myGoSourceFile2.go ...\n")
	fmt.Fprintf(os.Stderr, "     # addzid makes it easy to add ZebraPack serialization[1] to Go source files.\n")
	fmt.Fprintf(os.Stderr, "     # addzid reads .go files and adds `zid` tags to struct fields.\n")
	fmt.Fprintf(os.Stderr, "     #\n     # options:\n")
	fmt.Fprintf(os.Stderr, "     #   -o=\"odir\" specifies the directory to write to (created if need be).\n")
	fmt.Fprintf(os.Stderr, "     #   -unexported add tag to private fields of Go structs as well as public.\n")
	fmt.Fprintf(os.Stderr, "     #   -version   shows build version with git commit hash.\n")
	fmt.Fprintf(os.Stderr, "     #   -debug     print lots of debug info as we process.\n")
	fmt.Fprintf(os.Stderr, "     #   -OVERWRITE modify .go files in-place, adding zid tags (by default\n     #       we write to the to -o dir).\n")
	fmt.Fprintf(os.Stderr, "     #\n     # required: at least one .go source file for struct definitions.\n     #  note: the .go source file to process must be listed last, after any options.\n")
	fmt.Fprintf(os.Stderr, "     #\n")
	fmt.Fprintf(os.Stderr, "     # [1] https://github.com/glycerine/zebrapack \n")
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	MainArgs(os.Args)
}

// allow invocation from test
func MainArgs(args []string) {
	//fmt.Println(os.Args)
	os.Args = args

	flag.Usage = use
	if len(os.Args) < 2 {
		use()
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	debug := flag.Bool("debug", false, "print lots of debug info as we process.")
	verrequest := flag.Bool("version", false, "request git commit hash used to build this addzid")
	outdir := flag.String("o", "odir", "specify output directory")
	pkg := flag.String("p", "main", "specify package for generated code")
	privs := flag.Bool("unexported", false, "tag private as well as public struct fields")
	overwrite := flag.Bool("OVERWRITE", false, "replace named .go files with zid tagged versions.")
	flag.Parse()

	if debug != nil {
		Verbose = *debug
	}

	if verrequest != nil && *verrequest {
		fmt.Printf("%s\n", LASTGITCOMMITHASH)
		os.Exit(0)
	}

	if outdir == nil || *outdir == "" {
		fmt.Fprintf(os.Stderr, "required -o option missing. Use addzid -o <dirname> myfile.go # to specify the output directory.\n")
		use()
	}

	if !DirExists(*outdir) {
		err := os.MkdirAll(*outdir, 0755)
		if err != nil {
			panic(err)
		}
	}

	if pkg == nil || *pkg == "" {
		fmt.Fprintf(os.Stderr, "required -p option missing. Specify a package name for the generated go code with -p <pkgname>\n")
		use()
	}

	// all the rest are input .go files
	inputFiles := flag.Args()

	if len(inputFiles) == 0 {
		fmt.Fprintf(os.Stderr, "addzid needs at least one .go golang source file to process specified on the command line.\n")
		os.Exit(1)
	}

	for _, fn := range inputFiles {
		if !strings.HasSuffix(fn, ".go") && !strings.HasSuffix(fn, ".go.txt") {
			fmt.Fprintf(os.Stderr, "error: addzid input file '%s' did not end in '.go' or '.go.txt'.\n", fn)
			os.Exit(1)
		}
	}

	x := NewExtractor()
	x.fieldPrefix = "   "
	x.fieldSuffix = "\n"
	x.outDir = *outdir
	if privs != nil {
		x.extractPrivate = *privs
	}
	if overwrite != nil {
		x.overwrite = *overwrite
	}

	for _, inFile := range inputFiles {
		_, err := x.ExtractStructsFromOneFile(nil, inFile)
		if err != nil {
			panic(err)
		}
	}
	// get rid of default tmp dir
	x.compileDir.Cleanup()

	x.compileDir.DirPath = *outdir

	x.SetFinalFieldOrder()

	err := x.CopySourceFilesAddZidTag(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("generated files in '%s'\n", x.compileDir.DirPath)
}
