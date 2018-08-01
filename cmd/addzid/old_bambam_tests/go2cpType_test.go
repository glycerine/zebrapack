package main

import (
	"fmt"
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test001GoToCapTypeConversion(t *testing.T) {

	cv.Convey("Given a go structs:  type s1 struct { Ptrs []*int }", t, func() {
		cv.Convey(`then the goTypeSeq should be {"[]", "*", "int"} and the capTypeSeq should be {"List", "*", "Int64"} `, func() {

			src := `
type s1 struct {
	Ptrs []*int
}
`

			x := NewExtractor()
			defer x.Cleanup()
			_, err := ExtractStructs("", "package main; "+src, x)
			if err != nil {
				panic(err)
			}

			x.GenerateTranslators()

			capTypeSeq := []string{}
			goTypeSeq := []string{"[]", "*", "int"}
			var capTypeDisplay string
			capTypeSeq, capTypeDisplay = x.GoTypeToCapnpType(nil, goTypeSeq)

			fmt.Printf("capTypeSeq = '%v', goTypeSeq = '%v', capTypeDispaly = '%v'\n", capTypeSeq, goTypeSeq, capTypeDisplay)

			cv.So(len(capTypeSeq), cv.ShouldEqual, 3)
			cv.So(capTypeSeq[0], cv.ShouldEqual, "List")
			cv.So(capTypeSeq[1], cv.ShouldEqual, "*")
			cv.So(capTypeSeq[2], cv.ShouldEqual, "Int64")

			cv.So(capTypeDisplay, cv.ShouldEqual, "List(Int64)")
		})
	})
}
