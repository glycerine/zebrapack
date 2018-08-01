package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestZidSkip(t *testing.T) {

	cv.Convey("Given go: type Hi struct { A int `zid:\"skip\"`; B int  } ", t, func() {
		cv.Convey("then addzid should skip the A field and any negative numbered zid tags", func() {

			in1 := " type Hi struct { A int `zid:\"skip\"`; B int; C int `zid:\"-1\"`; D int `zid:\"-2\"`  } "

			expect1 := `
struct HiCapn { b  @0:   Int64; } 
`

			cv.So(ExtractString2String(in1), ShouldStartWithModuloWhiteSpace, expect1)
		})
	})
}
