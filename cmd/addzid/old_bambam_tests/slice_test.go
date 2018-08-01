package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestSliceToList(t *testing.T) {

	cv.Convey("Given a parsable golang source file with struct containing a slice", t, func() {
		cv.Convey("then the slice should be converted to a List() in the capnp output", func() {

			ex0 := `
type s1 struct {
  MyInts []int
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct S1Capn { myInts  @0:   List(Int64); } `)

		})
	})
}

func TestSliceOfStructToList(t *testing.T) {

	cv.Convey("Given a parsable golang source file with struct containing a slice of struct bbb", t, func() {
		cv.Convey("then the slice should be converted to a List(Bbb) in the capnp output", func() {

			ex0 := `
type bbb struct {}
type s1 struct {
  MyBees []bbb
}`
			out0 := ExtractString2String(ex0)

			VPrintf("out0: '%s'\n", out0)

			cv.So(out0, ShouldStartWithModuloWhiteSpace, `struct BbbCapn { } struct S1Capn { myBees  @0:   List(BbbCapn); } `)

		})
	})
}

func TestSliceOfPointerToList(t *testing.T) {

	cv.Convey("Given a parsable golang source file with struct containing a slice of pointers to struct big", t, func() {
		cv.Convey("then the slice should be converted to a List(Big) in the capnp output", func() {

			ex0 := `
type big struct {}
type s1 struct {
  MyBigs []*big
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct BigCapn { } struct S1Capn { myBigs  @0:   List(BigCapn); } `)

		})
	})
}

func TestSliceOfByteBecomesData(t *testing.T) {

	cv.Convey("Given golang src with []byte", t, func() {
		cv.Convey("then the slice should be converted to Data, not List(Byte), in the capnp output", func() {

			ex0 := `
type s1 struct {
  MyData []byte
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct S1Capn { myData  @0:   List(UInt8); } `)

		})
	})
}

func TestStructWithSliceOfOtherStructs(t *testing.T) {

	cv.Convey("Given a go struct containing MyBigs []Big, where Big is another struct", t, func() {
		cv.Convey("then then the CapnToGo() translation code should call the CapnToGo translation function over each slice member during translation", func() {

			in0 := `
type Big struct {}
type s1 struct {
  MyBigs []Big
}`

			expect0 := `
struct BigCapn { }
struct S1Capn { myBigs  @0:   List(BigCapn); } 


    func (s *Big) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	BigGoToCapn(seg, s)
    	_, err := seg.WriteTo(w)
        return err
    }
   
  
   
    func (s *Big) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
            return err
    	}
    	z := ReadRootBigCapn(capMsg)
        BigCapnToGo(z, s)
        return nil
    }

func BigCapnToGo(src BigCapn, dest *Big) *Big { 
    if dest == nil { 
      dest = &Big{} 
    }
    return dest
}
func BigGoToCapn(seg *capn.Segment, src *Big) BigCapn { 
    dest := AutoNewBigCapn(seg)
    return dest
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
    var n int

    // MyBigs
	n = src.MyBigs().Len()
	dest.MyBigs = make([]Big, n)
	for i := 0; i < n; i++ {
        dest.MyBigs[i] = *BigCapnToGo(src.MyBigs().At(i), nil)
    }

`
			cv.So(ExtractString2String(in0), ShouldStartWithModuloWhiteSpace, expect0)

		})
	})
}

// ==========================================
// ==========================================

func Test008SliceOfSliceOfStruct(t *testing.T) {

	cv.Convey("Given a go struct a slice of slice of int: type Cooper struct { Formation [][]Mini } ", t, func() {
		cv.Convey("then then List(List(Mini)) should be generated in the capnp schema", func() {

			in0 := `
type Mini struct {
  A int64
}
type Cooper struct {
  Downey      []Mini
  Formation [][]Mini
}`

			expect0 := `
struct CooperCapn { 
  downey     @0:   List(MiniCapn);
  formation  @1:   List(List(MiniCapn));
} 

struct MiniCapn { 
  a  @0:   Int64; 
} 
  
  
  func (s *Cooper) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	CooperGoToCapn(seg, s)
    	_, err := seg.WriteTo(w)
        return err
  }
   
  
   
  func (s *Cooper) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
            return err
    	}
    	z := ReadRootCooperCapn(capMsg)
        CooperCapnToGo(z, s)
        return nil
  }
  
  
  
  func CooperCapnToGo(src CooperCapn, dest *Cooper) *Cooper { 
    if dest == nil { 
      dest = &Cooper{} 
    }
  
    var n int
  
      // Downey
  	n = src.Downey().Len()
  	dest.Downey = make([]Mini, n)
  	for i := 0; i < n; i++ {
          dest.Downey[i] = *MiniCapnToGo(src.Downey().At(i), nil)
      }
  
  
      // Formation
  	n = src.Formation().Len()
  	dest.Formation = make([][]Mini, n)
  	for i := 0; i < n; i++ {
          dest.Formation[i] = MiniCapnListToSliceMini(MiniCapn_List(src.Formation().At(i)))
      }
  
    return dest
  } 
  
  
  
  func CooperGoToCapn(seg *capn.Segment, src *Cooper) CooperCapn { 
    dest := AutoNewCooperCapn(seg)
  
    // Downey -> MiniCapn (go slice to capn list)
    if len(src.Downey) > 0 {
  		typedList := NewMiniCapnList(seg, len(src.Downey))
  		plist := capn.PointerList(typedList)
  		i := 0
  		for _, ele := range src.Downey {
  			plist.Set(i, capn.Object(MiniGoToCapn(seg, &ele)))
  			i++
  		}
  		dest.SetDowney(typedList)
  	}
  
    // Formation -> MiniCapn (go slice to capn list)
    if len(src.Formation) > 0 {
  		plist := seg.NewPointerList(len(src.Formation))
  		i := 0
  		for _, ele := range src.Formation {
  			plist.Set(i, capn.Object(SliceMiniToMiniCapnList(seg, ele)))
  			i++
  		}
  		dest.SetFormation(plist)
  	}
  
    return dest
  } 
  
  
  
  func (s *Mini) Save(w io.Writer) error {
    	seg := capn.NewBuffer(nil)
    	MiniGoToCapn(seg, s)
    	_, err := seg.WriteTo(w)
        return err
  }
   
  
   
  func (s *Mini) Load(r io.Reader) error {
    	capMsg, err := capn.ReadFromStream(r, nil)
    	if err != nil {
    		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
            return err
    	}
    	z := ReadRootMiniCapn(capMsg)
        MiniCapnToGo(z, s)
        return nil
  }
  
  
  
  func MiniCapnToGo(src MiniCapn, dest *Mini) *Mini { 
    if dest == nil { 
      dest = &Mini{} 
    }
    dest.A = src.A()
  
    return dest
  } 
  
  
  
  func MiniGoToCapn(seg *capn.Segment, src *Mini) MiniCapn { 
    dest := AutoNewMiniCapn(seg)
    dest.SetA(src.A)
  
    return dest
  } 

  
  func SliceMiniToMiniCapnList(seg *capn.Segment, m []Mini) MiniCapn_List {
  	lst := NewMiniCapnList(seg, len(m))
  	for i := range m {
  		lst.Set(i, MiniGoToCapn(seg, &m[i]))
  	}
  	return lst
  }
  
  
  
  func MiniCapnListToSliceMini(p MiniCapn_List) []Mini {
  	v := make([]Mini, p.Len())
  	for i := range v {
       MiniCapnToGo(p.At(i), &v[i])
  	}
  	return v
  } 

`
			cv.So(ExtractString2String(in0), ShouldMatchModuloWhiteSpace, expect0)

		})
	})
}

// ==========================================
// ==========================================

func Test009SliceOfSliceOfInt(t *testing.T) {

	cv.Convey("Given a go struct a slice of slice: type Cooper struct { A [][]int } ", t, func() {
		cv.Convey("then then List(List(Int64)) should be generated in the capnp schema", func() {

			in0 := `
type Cooper struct {
  A [][]int
}`

			expect0 := `
struct CooperCapn { 
   a  @0:   List(List(Int64)); 
} 

func (s *Cooper) Save(w io.Writer) error {
	seg := capn.NewBuffer(nil)
	CooperGoToCapn(seg, s)
	_, err := seg.WriteTo(w)
    return err
}

func (s *Cooper) Load(r io.Reader) error {
	capMsg, err := capn.ReadFromStream(r, nil)
	if err != nil {
		//panic(fmt.Errorf("capn.ReadFromStream error: %s", err))
        return err
	}
	z := ReadRootCooperCapn(capMsg)
	CooperCapnToGo(z, s)
    return nil
}

func CooperCapnToGo(src CooperCapn, dest *Cooper) *Cooper {
	if dest == nil {
		dest = &Cooper{}
	}

	var n int

	// A
	n = src.A().Len()
	dest.A = make([][]int, n)
	for i := 0; i < n; i++ {
        dest.A[i] = Int64ListToSliceInt(capn.Int64List(src.A().At(i)))
	}

	return dest
}

func CooperGoToCapn(seg *capn.Segment, src *Cooper) CooperCapn {
	dest := AutoNewCooperCapn(seg)

	mylist1 := seg.NewPointerList(len(src.A))
	for i := range src.A {
		mylist1.Set(i, capn.Object(SliceIntToInt64List(seg, src.A[i])))
	}
	dest.SetA(mylist1)

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
`

			cv.So(ExtractString2String(in0), ShouldStartWithModuloWhiteSpace, expect0)

		})
	})
}
