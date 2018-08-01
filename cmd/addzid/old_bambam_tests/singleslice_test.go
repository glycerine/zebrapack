package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test004SingleSlice(t *testing.T) {

	cv.Convey("Given a parsable golang source file with type Vector struct { M []int }", t, func() {
		cv.Convey("then we should generate translation utility methods for List(Int64) <-> []int in the capnp output", func() {

			ex0 := `
type Vector struct {
  M []int
}`
			cv.So(ExtractString2String(ex0), ShouldContainModuloWhiteSpace, `

func Int64ListToSliceInt(p capn.Int64List) []int {
	v := make([]int, p.Len())
	for i := range v {
		v[i] = int(p.At(i))
	}
	return v
}
`)

			cv.So(ExtractString2String(ex0), ShouldContainModuloWhiteSpace, `
func SliceIntToInt64List(seg *capn.Segment, m []int) capn.Int64List {
	lst := seg.NewInt64List(len(m))
	for i := range m {
		lst.Set(i, int64(m[i]))
	}
	return lst
}

`)
		})
	})
}
