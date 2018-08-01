package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestUnderScoreFieldNamesRenamed(t *testing.T) {

	cv.Convey("underToCamelCase should remove underscores and leave camelCase", t, func() {
		cv.So(underToCamelCase(`one_two`), cv.ShouldEqual, `oneTwo`)
		cv.So(underToCamelCase(`one__two`), cv.ShouldEqual, `oneTwo`)
	})

	cv.Convey("Given a struct that contains fields with underscores", t, func() {
		cv.Convey("then these field names should be camel cased", func() {

			ex0 := `
type s1 struct {
  Hello_world int
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct S1Capn { helloWorld @0: Int64; } `)
		})
	})
}
