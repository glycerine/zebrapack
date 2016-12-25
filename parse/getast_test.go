package parse

import (
	"fmt"
	"os"

	cv "github.com/glycerine/goconvey/convey"
	"github.com/glycerine/zebrapack/cfg"
	"io/ioutil"
	"testing"
)

func Test001DuplicateOrMissingGapZidFieldsNotAllowed(t *testing.T) {

	cv.Convey("duplicate zid fields and sequences of zid fields with gaps are detected and forbidden", t, func() {

		fmt.Printf("\n  duplicate zid not allowed\n")
		s := "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"0\"`\n" +
			"}\n"
		cv.So(testCode(s), cv.ShouldNotBeNil)

		fmt.Printf("\n  zid must start at 0\n")
		s = "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"1\"`\n" +
			"   Wilma  string `zid:\"2\"`\n" +
			"}\n"
		cv.So(testCode(s), cv.ShouldNotBeNil)

		fmt.Printf("\n  negative zid not allowed\n")
		s = "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"-1\"`\n" +
			"}\n"
		cv.So(testCode(s), cv.ShouldNotBeNil)

		fmt.Printf("\n  0, 1 should be fine and not error\n")
		s = "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"1\"`\n" +
			"}\n"
		cv.So(testCode(s), cv.ShouldBeNil)

		fmt.Printf("\n  empty struct{} values should be allowed without error\n")
		s = "package p; type Flint struct {}"
		cv.So(testCode(s), cv.ShouldBeNil)

		fmt.Printf("\n  struct{} with zid should be allowed and count in the zid sequence\n")
		s = "package hoo; type S2 struct {A struct{} `zid:\"0\"`;B string   `zid:\"1\"`;}"
		cv.So(testCode(s), cv.ShouldBeNil)

		fmt.Printf("\n  can't have one zid on 2 fields\n")
		s = "package hoo; type S2 struct {A string `zid:\"0\"`; " +
			"B, C string   `zid:\"1\"`;}" // multiple fields, one zid.
		cv.So(testCode(s), cv.ShouldNotBeNil)

		fmt.Printf("\n  can't have one zid on 3 fields\n")
		s = "package hoo; type S2 struct {A string `zid:\"0\"`; " +
			"B, C, D string   `zid:\"1\"`;}" // multiple fields, one zid.
		cv.So(testCode(s), cv.ShouldNotBeNil)
	})
}

func Test002EmptyStructsNotMarshalled(t *testing.T) {

	cv.Convey("skipped fields (e.g. struct{} empty values) shouldn't be marshalled", t, func() {

		s := "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"1\"`\n" +
			"   Skiperoo  func()   \n" +
			"   Skiperoo2  struct{}   \n" +
			"}\n"
		cv.So(testCode(s), cv.ShouldBeNil)
	})
}

// testCode is a helper for checking for parsing errors
func testCode(code string) error {
	// struct{} fields should still count in their zid
	gofile, err := ioutil.TempFile(".", "tmp-test-001")
	panicOn(err)
	ofile := gofile.Name() + ".out"

	fmt.Fprintf(gofile, code)
	gofile.Close()

	fmt.Printf("\n in file '%s', checking:\n%v\n", gofile.Name(),
		code)

	cfg := cfg.ZebraConfig{
		Out:        ofile,
		GoFile:     gofile.Name(),
		Encode:     true,
		Marshal:    true,
		Tests:      true,
		Unexported: false,
	}
	_, err = File(&cfg)
	os.Remove(gofile.Name())
	os.Remove(ofile)
	return err
}
