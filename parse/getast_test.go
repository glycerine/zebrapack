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

		{
			gofile, err := ioutil.TempFile(".", "tmp-test-001")
			panicOn(err)
			ofile := gofile.Name() + ".out"

			s := "\npackage fred\n\n" +
				"type Flint struct {\n" +
				"   Barney string `zid:\"0\"`\n" +
				"   Wilma  string `zid:\"0\"`\n" +
				"}\n"
			fmt.Fprintf(gofile, s)
			gofile.Close()

			fmt.Printf("\n checking:\n%v\n", s)

			cfg := cfg.ZebraConfig{
				Out:        ofile,
				GoFile:     gofile.Name(),
				Encode:     true,
				Marshal:    true,
				Tests:      true,
				Unexported: false,
			}
			_, err = File(&cfg)
			cv.So(err, cv.ShouldNotBeNil)
			os.Remove(gofile.Name())
		}

		{
			gofile, err := ioutil.TempFile(".", "tmp-test-001")
			panicOn(err)
			ofile := gofile.Name() + ".out"

			s := "\npackage fred\n\n" +
				"type Flint struct {\n" +
				"   Barney string `zid:\"1\"`\n" +
				"   Wilma  string `zid:\"2\"`\n" +
				"}\n"
			fmt.Fprintf(gofile, s)
			gofile.Close()

			fmt.Printf("\n checking:\n%v\n", s)

			cfg := cfg.ZebraConfig{
				Out:        ofile,
				GoFile:     gofile.Name(),
				Encode:     true,
				Marshal:    true,
				Tests:      true,
				Unexported: false,
			}
			_, err = File(&cfg)
			cv.So(err, cv.ShouldNotBeNil)
			os.Remove(gofile.Name())
		}

		{
			gofile, err := ioutil.TempFile(".", "tmp-test-001")
			panicOn(err)
			ofile := gofile.Name() + ".out"

			s := "\npackage fred\n\n" +
				"type Flint struct {\n" +
				"   Barney string `zid:\"0\"`\n" +
				"   Wilma  string `zid:\"1\"`\n" +
				"}\n"
			fmt.Fprintf(gofile, s)
			gofile.Close()

			fmt.Printf("\n in file '%s', checking:\n%v\n", gofile.Name(), s)

			cfg := cfg.ZebraConfig{
				Out:        ofile,
				GoFile:     gofile.Name(),
				Encode:     true,
				Marshal:    true,
				Tests:      true,
				Unexported: false,
			}
			_, err = File(&cfg)
			cv.So(err, cv.ShouldBeNil)
			os.Remove(gofile.Name())
			os.Remove(ofile)
		}

		{
			gofile, err := ioutil.TempFile(".", "tmp-test-001")
			panicOn(err)
			ofile := gofile.Name() + ".out"

			s := "package p; type Flint struct {}"
			fmt.Fprintf(gofile, s)
			gofile.Close()

			fmt.Printf("\n in file '%s', checking:\n%v\n", gofile.Name(), s)

			cfg := cfg.ZebraConfig{
				Out:        ofile,
				GoFile:     gofile.Name(),
				Encode:     true,
				Marshal:    true,
				Tests:      true,
				Unexported: false,
			}
			_, err = File(&cfg)
			cv.So(err, cv.ShouldBeNil) // is only a warning.
			os.Remove(gofile.Name())
			os.Remove(ofile)
		}

	})
}
