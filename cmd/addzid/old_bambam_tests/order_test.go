package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestOrderOfDeclIrrel(t *testing.T) {

	cv.Convey("Given a golang file with a struct In1 that refers to struct In2 that is defined *later* in the file", t, func() {
		cv.Convey("then the first capnp definition for In1Capn consistently refer to the In2Capn defined later", func() {

			ex0 := `
type In1 struct {
   In2 In2
}
type In2 struct {}
`
			expect0 := `struct In1Capn { in2 @0: In2Capn; } struct In2Capn {}`
			act0 := ExtractString2String(ex0)
			//fmt.Printf("\n\n act0 = %#v\n expected = %#v\n", act0, expect0)

			cv.So(act0,
				ShouldStartWithModuloWhiteSpace, expect0)
		})
	})
}
