package testdata

import (
	"bytes"
	"testing"

	"github.com/glycerine/zebrapack/msgp"
	"github.com/glycerine/zebrapack/zebra"

	cv "github.com/glycerine/goconvey/convey"
)

func Test060ConvertZebraPackToMsgpack2(t *testing.T) {

	cv.Convey("ZebraPack + schema can be converted to msgpack2 via zSchema.ZebraToMsgp2()\n", t, func() {
		// read the schema
		msgp2schema := ZebraSchemaInMsgpack2Format()
		var zSchema zebra.Schema
		_, err := zSchema.UnmarshalMsg(msgp2schema)
		panicOn(err)

		// generate data
		var e S2
		e.B = "hello"
		e.P = 43
		data, err := e.MarshalMsg(nil)
		panicOn(err)

		// turn it into msgpack2 and display as json

		m2, _, err := zSchema.ZebraToMsgp2(data)
		panicOn(err)
		var json bytes.Buffer
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		cv.So(string(json.Bytes()), cv.ShouldEqual, `{"beta":"hello","P":43,"arr":[0,0,0,0,0,0]}`)
	})
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
