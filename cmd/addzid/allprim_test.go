package main

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func Test012AllPrimFields(t *testing.T) {

	cv.Convey("Given a struct that contains all primitive fields", t, func() {
		cv.Convey("then the capnp schema should contain the fields and we should generate translation code for the all field types", func() {

			ex0 := `
type s1 struct {
     S string
     I int
     B bool
     I8 int8
     I16 int16
     I32 int32
     I64 int64

     Ui8 uint8
     Ui16 uint16
     Ui32 uint32
     Ui64 uint64

     F32 float32
     F64 float64
     By  byte
}
`

			expected0 := `
struct S1Capn {
 s     @0:   Text;
 i     @1:   Int64;
 b     @2:   Bool;
 i8    @3:   Int8;
 i16   @4:   Int16;
 i32   @5:   Int32;
 i64   @6:   Int64;
 ui8   @7:   UInt8;
 ui16  @8:   UInt16;
 ui32  @9:   UInt32;
 ui64  @10:  UInt64;
 f32   @11:  Float32;
 f64   @12:  Float64;
 by    @13:  UInt8;
 } 
    
  func (s *s1) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	s1GoToCapn(seg, s)
      _, err := seg.WriteTo(w)
      return err
  }
   
  func (s *s1) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
          return err
    	}
    	z := ReadRootS1Capn(capMsg)
        S1CapnToGo(z, s)
     return nil
  }
  
  func S1CapnToGo(src S1Capn, dest *s1) *s1 {
    if dest == nil {
      dest = &s1{}
    }
    dest.S = src.S()
    dest.I = int(src.I())
    dest.B = src.B()
    dest.I8 = src.I8()
    dest.I16 = src.I16()
    dest.I32 = src.I32()
    dest.I64 = src.I64()
    dest.Ui8 = src.Ui8()
    dest.Ui16 = src.Ui16()
    dest.Ui32 = src.Ui32()
    dest.Ui64 = src.Ui64()
    dest.F32 = src.F32()
    dest.F64 = src.F64()
    dest.By = src.By()
  
    return dest
  }
    
  func s1GoToCapn(seg *capn.Segment, src *s1) S1Capn {
    dest := AutoNewS1Capn(seg)
    dest.SetS(src.S)
    dest.SetI(int64(src.I))
    dest.SetB(src.B)
    dest.SetI8(src.I8)
    dest.SetI16(src.I16)
    dest.SetI32(src.I32)
    dest.SetI64(src.I64)
    dest.SetUi8(src.Ui8)
    dest.SetUi16(src.Ui16)
    dest.SetUi32(src.Ui32)
    dest.SetUi64(src.Ui64)
    dest.SetF32(src.F32)
    dest.SetF64(src.F64)
    dest.SetBy(src.By)
  
    return dest
  }
`

			cv.So(ExtractString2String(ex0), ShouldMatchModuloWhiteSpace, expected0)
			//cv.So(expected0, ShouldStartWithModuloWhiteSpace, ExtractString2String(ex0))

		})
	})
}
