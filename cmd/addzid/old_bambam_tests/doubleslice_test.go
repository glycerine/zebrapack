package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test002SlistSliceIntToListListInt64(t *testing.T) {

	cv.Convey("Given a parsable golang source file with type Matrix struct { M [][]int }", t, func() {
		cv.Convey("then the slice should be converted to a List(List(Int64)) in the capnp output", func() {

			ex0 := `
type Matrix struct {
  M [][]int
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `
struct MatrixCapn { 
  m  @0:   List(List(Int64)); 
} 

  func (s *Matrix) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	MatrixGoToCapn(seg, s)
    	_, err := seg.WriteTo(w)
        return err
  }
   
  func (s *Matrix) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
            return err
    	}
    	z := ReadRootMatrixCapn(capMsg)
        MatrixCapnToGo(z, s)
        return nil
  }
  
  func MatrixCapnToGo(src MatrixCapn, dest *Matrix) *Matrix { 
    if dest == nil { 
      dest = &Matrix{} 
    }
  
    var n int
  
      // M
  	n = src.M().Len()
  	dest.M = make([][]int, n)
  	for i := 0; i < n; i++ {
		dest.M[i] = Int64ListToSliceInt(capn.Int64List(src.M().At(i)))
    }
    
    return dest
  } 

  func MatrixGoToCapn(seg *capn.Segment, src *Matrix) MatrixCapn { 
    dest := AutoNewMatrixCapn(seg)
  
	mylist1 := seg.NewPointerList(len(src.M))
	for i := range src.M {
		mylist1.Set(i, capn.Object(SliceIntToInt64List(seg, src.M[i])))
	}
	dest.SetM(mylist1)

    return dest
  } 

func SliceIntToInt64List(seg *capn.Segment, m []int) capn.Int64List {
	lst := seg.NewInt64List(len(m))
	for i := range m {
		lst.Set(i, int64(m[i]))
	}
	return lst
}

func Int64ListToSliceInt(p capn.Int64List) []int {
	v := make([]int, p.Len())
	for i := range v {
		v[i] = int(p.At(i))
	}
	return v
}
`)
		})
	})
}
