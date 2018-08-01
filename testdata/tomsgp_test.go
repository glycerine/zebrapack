package testdata

import (
	"bytes"
	"testing"

	"github.com/glycerine/zebrapack/msgp"
	"github.com/glycerine/zebrapack/zebra"

	cv "github.com/smartystreets/goconvey/convey"
)

func Test060ConvertZebraPackToMsgpack2(t *testing.T) {

	cv.Convey("ZebraPack + schema can be converted to msgpack2 via zSchema.ZebraToMsgp2()\n", t, func() {
		// read the schema
		msgp2schema := FileMy{}.ZebraSchemaInMsgpack2Format()
		var zSchema zebra.Schema
		_, err := zSchema.UnmarshalMsg(msgp2schema)
		panicOn(err)

		// generate data
		var e S2
		e.B = "hello"
		e.P = 43
		data, err := e.MarshalMsg(nil)
		panicOn(err)

		// confirm that encode works the same
		var encbuf bytes.Buffer
		err = msgp.Encode(&encbuf, &e)
		panicOn(err)
		cv.So(string(encbuf.Bytes()), cv.ShouldEqual, string(data))

		// turn it into msgpack2 and display as json

		m2, _, err := zSchema.ZebraToMsgp2(data, false)
		panicOn(err)
		var json bytes.Buffer
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		cv.So(string(json.Bytes()), cv.ShouldEqual, `{"beta":"hello","P":43,"T":0,"arr":[0,0,0,0,0,0]}`)

		// and test the recursive conversion with MyTree
		e.MyTree = &Tree{
			Str:  "root",
			Chld: []Tree{Tree{Str: "mid", Chld: []Tree{{Str: "leaf"}}}},
		}

		data, err = e.MarshalMsg(nil)
		panicOn(err)

		m2, _, err = zSchema.ZebraToMsgp2(data, false)
		panicOn(err)
		json.Reset()
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		cv.So(string(json.Bytes()), cv.ShouldEqual, `{"beta":"hello","P":43,"T":0,"arr":[0,0,0,0,0,0],"MyTree":{"Chld":[{"Chld":[{"Str":"leaf"}],"Str":"mid"}],"Str":"root"}}`)

		// confirm that Encode works the same as MarshalMsg
		var encbuf2 bytes.Buffer
		err = msgp.Encode(&encbuf2, &e)
		panicOn(err)
		cv.So(string(encbuf2.Bytes()), cv.ShouldEqual, string(data))

		// handle missing schema gracefully, test by
		// zero-ing out the Structs map.
		zSchema.Structs = make(map[string]*zebra.Struct)
		m2, _, err = zSchema.ZebraToMsgp2(data, true)
		panicOn(err)
		json.Reset()
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		cv.So(string(json.Bytes()), cv.ShouldEqual, `{"1":"hello","3":43,"6":[0,0,0,0,0,0],"7":{"0":[{"0":[{"1":"leaf"}],"1":"mid"}],"1":"root"}}`)

	})
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
