package main

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func TestCapidTagOrderAssignment(t *testing.T) {

	cv.Convey("Given capid tag requests for field numbering", t, func() {
		cv.Convey("then when assigning field numbers in the capnproto struct schema, assign capids first, then proceed in order of field appearance", func() {

			ex0 := "type Z struct { A int `capid:\"2\"`; B int `capid:\"0\"`; C int `capid:\"1\"` }"
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct ZCapn { b  @0:   Int64; c  @1:   Int64; a  @2:   Int64; } `)

			ex1 := "type Z struct { A int `capid:\"1\"`; B int; C int `capid:\"0\"` }"
			cv.So(ExtractString2String(ex1), ShouldStartWithModuloWhiteSpace, `struct ZCapn { c  @0:   Int64; a  @1:   Int64; b  @2:   Int64; } `)

		})

	})
}
