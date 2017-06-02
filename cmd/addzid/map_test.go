package main

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func Test910MapAddZid(t *testing.T) {

	cv.Convey("Given a struct with a map[string]bool, ", t, func() {
		cv.Convey("then addzid should add a zid for the map field.", func() {
			in1 := "type HasMap struct {Jobid map[string]bool }"
			ex1 := "package main type HasMap struct { Jobid map[string]bool `zid:\"0\"`}"
			added := ExtractStringAddZid(in1)
			p("added='%v'", added)
			cv.So(added, ShouldStartWithModuloWhiteSpace, ex1)
		})
	})
}
