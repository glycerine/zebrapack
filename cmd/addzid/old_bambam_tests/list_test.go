package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test005SliceOfStringToListOfText(t *testing.T) {

	cv.Convey("Given a struct that contains a slice of string", t, func() {
		cv.Convey("then the capnp schema should contain a list of string and list translating code should be generated", func() {

			ex0 := `
type s1 struct {
   Names []string
}`

			expected0 := `
  struct S1Capn { names  @0:   List(Text); } 

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

    dest.Names = src.Names().ToArray()
  
    return dest
  } 
  
  func s1GoToCapn(seg *capn.Segment, src *s1) S1Capn { 
    dest := AutoNewS1Capn(seg)
  
    mylist1 := seg.NewTextList(len(src.Names))
    for i := range src.Names {
       mylist1.Set(i, string(src.Names[i]))
    }
    dest.SetNames(mylist1)
  
    return dest
  } 

  func SliceStringToTextList(seg *capn.Segment, m []string) capn.TextList {
  	lst := seg.NewTextList(len(m))
  	for i := range m {
  		lst.Set(i, string(m[i]))
  	}
  	return lst
  }
  
  
  
  func TextListToSliceString(p capn.TextList) []string {
  	v := make([]string, p.Len())
  	for i := range v {
  		v[i] = string(p.At(i))
  	}
  	return v
  } 

`

			cv.So(ExtractString2String(ex0), ShouldMatchModuloWhiteSpace, expected0)
			//cv.So(expected0, ShouldStartWithModuloWhiteSpace, ExtractString2String(ex0))

		})
	})
}
