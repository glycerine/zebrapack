package main

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func TestCapidSkip(t *testing.T) {

	cv.Convey("Given go: type Hi struct { A int `capid:\"skip\"`; B int  } ", t, func() {
		cv.Convey("then bambam should skip the A field and any negative numbered capid tags", func() {

			in1 := " type Hi struct { A int `capid:\"skip\"`; B int; C int `capid:\"-1\"`; D int `capid:\"-2\"`  } "

			expect1 := `
struct HiCapn { b  @0:   Int64; } 
`

			cv.So(ExtractString2String(in1), ShouldStartWithModuloWhiteSpace, expect1)
		})
	})
}
