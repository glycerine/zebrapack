package testdata

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/glycerine/zebrapack/msgp"
	"github.com/glycerine/zebrapack/zebra"

	cv "github.com/glycerine/goconvey/convey"
)

func Test070PuttingZebrapackSchemaAtTopOfFileOnce(t *testing.T) {

	cv.Convey("In a typical use case for ZebraPack, the schema can be written once at the top of the file, and then afterwards used to interpret the rest of the file.\n", t, func() {

		msgp2schema := FileHeaderplus{}.ZebraSchemaInMsgpack2Format()
		var zSchema zebra.Schema
		_, err := zSchema.UnmarshalMsg(msgp2schema)
		panicOn(err)

		var pseudoFile bytes.Buffer
		w := msgp.NewWriter(&pseudoFile)

		// write header
		h := &Header{
			S1: "hello",
			S2: "universe",
			C:  43,
		}
		err = h.EncodeMsg(w)
		panicOn(err)
		hedwig := &Owl{Nm: "Hedwig"}

		// write many Pluses
		n := 0
		f := 0.0
		for i := 0; i < 3; i++ {
			plus := &Plus{
				Tm:  time.Now(),
				N:   int64(n + i),
				F:   f - float64(i),
				S:   "salutations",
				Pet: hedwig,
			}
			err = plus.EncodeMsg(w)
			panicOn(err)
		}
		w.Flush()
		by := pseudoFile.Bytes()
		fmt.Printf("by='%#v' / '%s'\n", by, string(by))

		// turn it into msgpack2 and display as json
		m2, _, err := zSchema.ZebraToMsgp2(by, false)
		panicOn(err)
		var json bytes.Buffer
		_, err = msgp.CopyToJSON(&json, bytes.NewBuffer(m2))
		panicOn(err)
		cv.So(string(json.Bytes()), cv.ShouldEqual, `{"S1":"hello","S2":"universe","C":43}`)

		var h2 Header
		var p2 Plus
		o, err := h2.UnmarshalMsg(by)
		panicOn(err)

		o, err = p2.UnmarshalMsg(o)
		panicOn(err)
		cv.So(p2.Pet.Name(), cv.ShouldResemble, "Hedwig")
	})
}
