package main

import (
	"testing"
	"time"

	cv "github.com/smartystreets/goconvey/convey"
)

// run 'addzid -unexported' to see zid being added.
type unused struct {
	tm time.Time
}

func Test909TimeFieldsAddZid(t *testing.T) {

	cv.Convey("Given go: type HasTime struct { A time.Time; B int  } ", t, func() {
		cv.Convey("then addzid should add a zid for both the time.Time and the int field.", func() {
			in1 := "type HasTime struct { A time.Time; B int; } "
			ex1 := "package main type HasTime struct { A time.Time `zid:\"0\"` B int `zid:\"1\"`}"
			added := ExtractStringAddZid(in1)
			//p("added='%v'", added)
			cv.So(added, ShouldStartWithModuloWhiteSpace, ex1)
		})
	})
}
