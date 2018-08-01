package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestNestedStructs016(t *testing.T) {

	cv.Convey("Given go structs that contain structs \n"+
		"         then these should marshal and unmarshal their inner struct members", t, func() {

		ex0 := `
type Inner struct {
 C int
}
type Outer struct {
  A int
  B Inner
}
`
		cv.So(ExtractString2String(ex0), ShouldMatchModuloWhiteSpace, `
struct InnerCapn {
  c  @0:   Int64;
}

struct OuterCapn {
  a  @0:   Int64; b  @1:   InnerCapn;
}

  func (s *Inner) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	InnerGoToCapn(seg, s)
      _, err := seg.WriteTo(w)
      return err
  }



  func (s *Inner) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
          return err
    	}
    	z := ReadRootInnerCapn(capMsg)
        InnerCapnToGo(z, s)
     return nil
  }



  func InnerCapnToGo(src InnerCapn, dest *Inner) *Inner {
    if dest == nil {
      dest = &Inner{}
    }
    dest.C = int(src.C())

    return dest
  }



  func InnerGoToCapn(seg *capn.Segment, src *Inner) InnerCapn {
    dest := AutoNewInnerCapn(seg)
    dest.SetC(int64(src.C))

    return dest
  }



  func (s *Outer) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	OuterGoToCapn(seg, s)
      _, err := seg.WriteTo(w)
      return err
  }



  func (s *Outer) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
          return err
    	}
    	z := ReadRootOuterCapn(capMsg)
        OuterCapnToGo(z, s)
     return nil
  }



  func OuterCapnToGo(src OuterCapn, dest *Outer) *Outer {
    if dest == nil {
      dest = &Outer{}
    }
    dest.A = int(src.A())
    dest.B = *InnerCapnToGo(src.B(), nil)

    return dest
  }



  func OuterGoToCapn(seg *capn.Segment, src *Outer) OuterCapn {
    dest := AutoNewOuterCapn(seg)
    dest.SetA(int64(src.A))
    dest.SetB(InnerGoToCapn(seg, &src.B))

    return dest
  }
`)
	})
}
