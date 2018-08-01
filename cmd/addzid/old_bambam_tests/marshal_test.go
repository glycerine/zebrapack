package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestCapnpStructNaming(t *testing.T) {

	cv.Convey("Given go structs that we will marshal into capnp structs", t, func() {
		cv.Convey("then the names of the two types in go code should be distinct: Cpn suffix attached to the capnp structs", func() {

			ex0 := `
type Extra struct {
  A int
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct ExtraCapn { a @0: Int64; } `)
		})
	})
}

func TestMarshal(t *testing.T) {

	cv.Convey("Given go struct Extra", t, func() {
		cv.Convey("then the generated ExtraCapntoGo() code should copy content from an ExtraCapn to an Extra struct.", func() {
			cv.Convey("and should handle int fields", func() {
				ex0 := `
type Extra struct {
  A int
  B int
}`
				toGoCode := ExtractCapnToGoCode(ex0, "Extra")
				cv.So(toGoCode, ShouldMatchModuloWhiteSpace, `
func ExtraCapnToGo(src ExtraCapn, dest *Extra) *Extra { 
  if dest == nil { 
     dest = &Extra{} 
  }
  dest.A = int(src.A())
  dest.B = int(src.B())
  return dest } 
`)

				toCapnCode := ExtractGoToCapnCode(ex0, "Extra")
				cv.So(toCapnCode, ShouldMatchModuloWhiteSpace, `
func ExtraGoToCapn(seg *capn.Segment, src *Extra) ExtraCapn { 
  dest := AutoNewExtraCapn(seg)
  dest.SetA(int64(src.A))
  dest.SetB(int64(src.B))
  return dest } 
`)

			})
			cv.Convey("and should handle uint fields", func() {
				ex0 := `
type Extra struct {
  A uint64
}`
				toGoCode := ExtractCapnToGoCode(ex0, "Extra")
				cv.So(toGoCode, ShouldMatchModuloWhiteSpace, `
func ExtraCapnToGo(src ExtraCapn, dest *Extra) *Extra {
  if dest == nil {
     dest = &Extra{}
  }
  dest.A = src.A()
  return dest }
`)

				toCapnCode := ExtractGoToCapnCode(ex0, "Extra")
				cv.So(toCapnCode, ShouldMatchModuloWhiteSpace, `
func ExtraGoToCapn(seg *capn.Segment, src *Extra) ExtraCapn {
  dest := AutoNewExtraCapn(seg)
  dest.SetA(src.A)
  return dest }
`)
			})
		})
	})
}

func TestUnMarshal(t *testing.T) {

	cv.Convey("Given go structs", t, func() {
		cv.Convey("then the generated Unmarshal() code should copy from the capnp into the go structs", func() {

		})
	})
}
