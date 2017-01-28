package main

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func Test909TimeFieldsAddZid(t *testing.T) {

	cv.Convey("Given go: type HasTime struct { A time.Time; B int  } ", t, func() {
		cv.Convey("then addzid should add a zid for both the time.Time and the int field.", func() {
			in1 := "type HasTime struct { A time.Time; B int } "
			ex1 := "type HasTime struct { A time.Time `zid:\"0\"`; B int `zid:\"1\"`}"
			cv.So(ExtractString2String(in1), ShouldStartWithModuloWhiteSpace, ex1)
		})
	})
}
