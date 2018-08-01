package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test010ListListStruct(t *testing.T) {

	cv.Convey("Given a go [][]struct that maps to List(List(Struct)) in capnproto", t, func() {
		cv.Convey("then out list of list of struct serialization code should work", func() {

			ex0 := `
type Nester1 struct {
  Strs []string
}
type RWTest struct {
    NestMatrix [][]Nester1
}
`
			cv.So(ExtractString2String(ex0), ShouldContainModuloWhiteSpace, `

func RWTestGoToCapn(seg *capn.Segment, src *RWTest) RWTestCapn {
	dest := AutoNewRWTestCapn(seg)

	// NestMatrix -> Nester1Capn (go slice to capn list)
	if len(src.NestMatrix) > 0 {
		plist := seg.NewPointerList(len(src.NestMatrix))
		i := 0
		for _, ele := range src.NestMatrix {
			plist.Set(i, capn.Object(SliceNester1ToNester1CapnList(seg, ele)))
			i++
		}
		dest.SetNestMatrix(plist)
	}

	return dest
}
`)

		})
	})
}
