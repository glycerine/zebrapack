package testdata

import (
	//"fmt"
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
	"github.com/glycerine/zebrapack/msgp"
)

func Test010StructReuseDoesNotAllocat(t *testing.T) {

	cv.Convey("re-using a struct and unmarshalling a nil entry should not re-allocate the member fields", t, func() {

		v := S2{}
		bts, err := v.MarshalMsg(nil)
		if err != nil {
			t.Fatal(err)
		}
		//fmt.Printf("bts = '%x'\n", bts)
		v.P = 2
		v.Q = 7
		v.R = make(map[string]uint8)
		v.R["yo"] = 8
		before := &v.R
		//fmt.Printf("v = '%#v', yo = %v, v.R=%p\n", v, v.R["yo"], &v.R)
		left, err := v.UnmarshalMsg(bts)
		//fmt.Printf("after v = '%#v', yo = %v, v.R=%p\n", v, v.R["yo"], &v.R)
		after := &v.R
		cv.So(after, cv.ShouldEqual, before)
		cv.So(v.P, cv.ShouldEqual, 0)
		cv.So(v.Q, cv.ShouldEqual, 0)
		cv.So(v.R["yo"], cv.ShouldEqual, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(left) > 0 {
			t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
		}

		left, err = msgp.Skip(bts)
		if err != nil {
			t.Fatal(err)
		}
		if len(left) > 0 {
			t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
		}
	})
}
