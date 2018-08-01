package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestPointerInStruct(t *testing.T) {

	cv.Convey("Given a struct that contains a pointer to a struct", t, func() {
		cv.Convey("then the pointer element should be de-pointerized", func() {

			ex0 := `
type big struct {}
type s1 struct {
  MyBig *big
}`
			cv.So(ExtractString2String(ex0), ShouldStartWithModuloWhiteSpace, `struct BigCapn { } struct S1Capn { myBig @0: BigCapn; } `)

		})
	})
}

func TestPointerInSliceInStruct(t *testing.T) {

	cv.Convey("Given a struct that contains a slice of pointers", t, func() {
		cv.Convey("then the translation routines should still work", func() {

			in0 := `
type Big struct {
	A int
}
type s1 struct {
	Ptrs []*Big
}
`

			expect0 := `
struct BigCapn { a @0: Int64; }
struct S1Capn { ptrs @0: List(BigCapn); } 

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
	dest.A = int(src.A())

	return dest
}

func BigGoToCapn(seg *capn.Segment, src *Big) BigCapn {
	dest := AutoNewBigCapn(seg)
	dest.SetA(int64(src.A))

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

	// Ptrs
	n = src.Ptrs().Len()
	dest.Ptrs = make([]*Big, n)
	for i := 0; i < n; i++ {
		dest.Ptrs[i] = BigCapnToGo(src.Ptrs().At(i), nil)
	}

	return dest
}

func s1GoToCapn(seg *capn.Segment, src *s1) S1Capn {
	dest := AutoNewS1Capn(seg)

	// Ptrs -> BigCapn (go slice to capn list)
	if len(src.Ptrs) > 0 {
		typedList := NewBigCapnList(seg, len(src.Ptrs))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.Ptrs {
			plist.Set(i, capn.Object(BigGoToCapn(seg, ele)))
			i++
		}
		dest.SetPtrs(typedList)
	}

	return dest
}
`
			cv.So(ExtractString2String(in0), ShouldStartWithModuloWhiteSpace, expect0)
		})
	})
}

func TestPointerAndStraightInSliceInStruct(t *testing.T) {

	cv.Convey("Given a struct that contains a slice of pointers and a slice of struct", t, func() {
		cv.Convey("then the translation routines should still work", func() {

			in0 := `
type Big struct {
	A int
	B string
	C []string
}
type s1 struct {
	Ptrs     []*Big
	Straight []Big
}
`
			expect0 := `
struct BigCapn { a  @0:   Int64; b  @1:   Text; c  @2:   List(Text); } 
struct S1Capn { ptrs      @0:   List(BigCapn); straight  @1:   List(BigCapn); } 

  
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
	dest.A = int(src.A())
	dest.B = src.B()
	dest.C = src.C().ToArray()

	return dest
}

func BigGoToCapn(seg *capn.Segment, src *Big) BigCapn {
	dest := AutoNewBigCapn(seg)
	dest.SetA(int64(src.A))
	dest.SetB(src.B)

	mylist1 := seg.NewTextList(len(src.C))
	for i := range src.C {
		mylist1.Set(i, string(src.C[i]))
	}
	dest.SetC(mylist1)

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

	// Ptrs
	n = src.Ptrs().Len()
	dest.Ptrs = make([]*Big, n)
	for i := 0; i < n; i++ {
		dest.Ptrs[i] = BigCapnToGo(src.Ptrs().At(i), nil)
	}

	// Straight
	n = src.Straight().Len()
	dest.Straight = make([]Big, n)
	for i := 0; i < n; i++ {
		dest.Straight[i] = *BigCapnToGo(src.Straight().At(i), nil)
	}

	return dest
}

func s1GoToCapn(seg *capn.Segment, src *s1) S1Capn {
	dest := AutoNewS1Capn(seg)

	// Ptrs -> BigCapn (go slice to capn list)
	if len(src.Ptrs) > 0 {
		typedList := NewBigCapnList(seg, len(src.Ptrs))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.Ptrs {
			plist.Set(i, capn.Object(BigGoToCapn(seg, ele)))
			i++
		}
		dest.SetPtrs(typedList)
	}

	// Straight -> BigCapn (go slice to capn list)
	if len(src.Straight) > 0 {
		typedList := NewBigCapnList(seg, len(src.Straight))
		plist := capn.PointerList(typedList)
		i := 0
		for _, ele := range src.Straight {
			plist.Set(i, capn.Object(BigGoToCapn(seg, &ele)))
			i++
		}
		dest.SetStraight(typedList)
	}

	return dest
}
`
			cv.So(ExtractString2String(in0), ShouldStartWithModuloWhiteSpace, expect0)
		})
	})
}
