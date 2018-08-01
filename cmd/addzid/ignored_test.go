package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test911AddZidIgnoresUnserializedFields(t *testing.T) {

	cv.Convey("Given a struct with a field that has a msg:\"-\" tag, ", t, func() {
		cv.Convey("then addzid to should ignore that field, and not add a zid to it..", func() {
			in1 := "type HasUnserz struct {Jobid int ` msg: \"-\" `}"
			ex1 := "package main type HasUnserz struct {Jobid int `msg:\"-\"`}"
			added := ExtractStringAddZid(in1)
			//p("added='%v'", added)
			cv.So(added, ShouldStartWithModuloWhiteSpace, ex1)
		})
	})
}
