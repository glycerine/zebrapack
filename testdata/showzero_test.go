package testdata

import (
	"bytes"
	//"fmt"
	"testing"

	"github.com/glycerine/zebrapack/msgp"
	"github.com/glycerine/zebrapack/zebra"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test080ShowZero(t *testing.T) {

	cv.Convey(`Given a field tagged with msg:",showzero", when the field has a zero value in it, it should still be printed`, t, func() {

		// read the schema
		msgp2schema := FileMy{}.ZebraSchemaInMsgpack2Format()
		var zSchema zebra.Schema
		_, err := zSchema.UnmarshalMsg(msgp2schema)
		panicOn(err)

		// generate data
		var e S2
		data, err := e.MarshalMsg(nil)
		panicOn(err)

		// turn it into msgpack2 and display as json
		m2, _, err := zSchema.ZebraToMsgp2(data, false)
		/*
					fmt.Printf("m2 = raw(")
			    	for i := range m2 {
						fmt.Printf("0x%x,", m2[i])
					}
		*/
		panicOn(err)
		var json bytes.Buffer
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		s := string(json.Bytes())
		//fmt.Printf("s = '%v'\n", s)
		cv.So(s, cv.ShouldEqual, `{"T":0,"arr":[0,0,0,0,0,0]}`)
	})
}

func Test081ShowZeroLastField(t *testing.T) {

	cv.Convey(`Given a field tagged with msg:",showzero" of type slice and pointer, json version should show/display the field name and null`, t, func() {

		// read the schema
		msgp2schema := FileMy{}.ZebraSchemaInMsgpack2Format()
		var zSchema zebra.Schema
		_, err := zSchema.UnmarshalMsg(msgp2schema)
		panicOn(err)

		// generate data
		var e Big
		data, err := e.MarshalMsg(nil)
		panicOn(err)

		// turn it into msgpack2 and display as json
		m2, _, err := zSchema.ZebraToMsgp2(data, false)
		//fmt.Printf("m2 = raw(")
		//for i := range m2 {
		//	fmt.Printf("0x%x,", m2[i])
		//}
		panicOn(err)
		var json bytes.Buffer
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		s := string(json.Bytes())
		//fmt.Printf("s = '%v'\n", s)
		cv.So(s, cv.ShouldEqual, `{"Slice":null,"Transform":null,"Myptr":null,"Myarray":["","",""],"MySlice":null}`)
	})
}
