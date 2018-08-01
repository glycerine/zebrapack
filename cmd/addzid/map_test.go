package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test910MapAddZid(t *testing.T) {

	cv.Convey("Given a struct with a map , ", t, func() {
		cv.Convey("with map[string]bool, then addzid to should add a zid for the map field.", func() {
			in1 := "type HasMap struct {Jobid map[string]bool }"
			ex1 := "package main type HasMap struct { Jobid map[string]bool `zid:\"0\"`}"
			added := ExtractStringAddZid(in1)
			//p("added='%v'", added)
			cv.So(added, ShouldStartWithModuloWhiteSpace, ex1)
		})

		cv.Convey("with a more complex map, then addzid to should add a zid for the map field.", func() {
			in1 := "type HasMap struct {S string  `zid:\"0\"`; Jobid map[int]map[string]Jobid }"
			ex1 := "package main type HasMap struct { S string  `zid:\"0\"`\nJobid map[int]map[string]Jobid  `zid:\"1\"`}"
			added := ExtractStringAddZid(in1)
			//p("added='%v'", added)
			cv.So(added, ShouldStartWithModuloWhiteSpace, ex1)
		})

	})
}
