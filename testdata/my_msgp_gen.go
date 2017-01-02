package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
)

// MSGPfieldsNotEmpty supports omitempty tags
func (z *A) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Sibs == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *A) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Bday"
	o = append(o, 0xa4, 0x42, 0x64, 0x61, 0x79)
	o = msgp.AppendTime(o, z.Bday)
	if !empty[2] {
		// string "phone"
		o = append(o, 0xa5, 0x70, 0x68, 0x6f, 0x6e, 0x65)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// string "Sibs"
		o = append(o, 0xa4, 0x53, 0x69, 0x62, 0x73)
		o = msgp.AppendInt(o, z.Sibs)
	}

	// string "GPA"
	o = append(o, 0xa3, 0x47, 0x50, 0x41)
	o = msgp.AppendFloat64(o, z.GPA)
	// string "Friend"
	o = append(o, 0xa6, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64)
	o = msgp.AppendBool(o, z.Friend)
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *A) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *A) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0zdnv = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zdnv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zdnv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zdnv := totalEncodedFields0zdnv
	missingFieldsLeft0zdnv := maxFields0zdnv - totalEncodedFields0zdnv

	var nextMiss0zdnv int32 = -1
	var found0zdnv [maxFields0zdnv]bool
	var curField0zdnv string

doneWithStruct0zdnv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zdnv > 0 || missingFieldsLeft0zdnv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zdnv, missingFieldsLeft0zdnv, msgp.ShowFound(found0zdnv[:]), unmarshalMsgFieldOrder0zdnv)
		if encodedFieldsLeft0zdnv > 0 {
			encodedFieldsLeft0zdnv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zdnv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zdnv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zdnv = 0
			}
			for nextMiss0zdnv < maxFields0zdnv && (found0zdnv[nextMiss0zdnv] || unmarshalMsgFieldSkip0zdnv[nextMiss0zdnv]) {
				nextMiss0zdnv++
			}
			if nextMiss0zdnv == maxFields0zdnv {
				// filled all the empty fields!
				break doneWithStruct0zdnv
			}
			missingFieldsLeft0zdnv--
			curField0zdnv = unmarshalMsgFieldOrder0zdnv[nextMiss0zdnv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zdnv)
		switch curField0zdnv {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zdnv[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0zdnv[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0zdnv[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0zdnv[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0zdnv[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0zdnv[5] = true
			z.Friend, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zdnv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zdnv = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zdnv = []bool{false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.TimeSize + 6 + msgp.StringPrefixSize + len(z.Phone) + 5 + msgp.IntSize + 4 + msgp.Float64Size + 7 + msgp.BoolSize
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Big) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Big) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 5
	// string "Slice"
	o = append(o, 0x85, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
	for zkqy := range z.Slice {
		o, err = z.Slice[zkqy].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zjwh, zwdg := range z.Transform {
		o = msgp.AppendInt(o, zjwh)
		if zwdg == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zwdg.MSGPMarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Myptr"
	o = append(o, 0xa5, 0x4d, 0x79, 0x70, 0x74, 0x72)
	if z.Myptr == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Myptr.MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Myarray"
	o = append(o, 0xa7, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, 3)
	for zvva := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zvva])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zdya := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zdya])
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Big) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Big) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1ztmi = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1ztmi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1ztmi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1ztmi := totalEncodedFields1ztmi
	missingFieldsLeft1ztmi := maxFields1ztmi - totalEncodedFields1ztmi

	var nextMiss1ztmi int32 = -1
	var found1ztmi [maxFields1ztmi]bool
	var curField1ztmi string

doneWithStruct1ztmi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1ztmi > 0 || missingFieldsLeft1ztmi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1ztmi, missingFieldsLeft1ztmi, msgp.ShowFound(found1ztmi[:]), unmarshalMsgFieldOrder1ztmi)
		if encodedFieldsLeft1ztmi > 0 {
			encodedFieldsLeft1ztmi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1ztmi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1ztmi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1ztmi = 0
			}
			for nextMiss1ztmi < maxFields1ztmi && (found1ztmi[nextMiss1ztmi] || unmarshalMsgFieldSkip1ztmi[nextMiss1ztmi]) {
				nextMiss1ztmi++
			}
			if nextMiss1ztmi == maxFields1ztmi {
				// filled all the empty fields!
				break doneWithStruct1ztmi
			}
			missingFieldsLeft1ztmi--
			curField1ztmi = unmarshalMsgFieldOrder1ztmi[nextMiss1ztmi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1ztmi)
		switch curField1ztmi {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1ztmi[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zhda uint32
				zhda, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(zhda) {
					z.Slice = (z.Slice)[:zhda]
				} else {
					z.Slice = make([]S2, zhda)
				}
				for zkqy := range z.Slice {
					bts, err = z.Slice[zkqy].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1ztmi[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zghg uint32
				zghg, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zghg > 0 {
					z.Transform = make(map[int]*S2, zghg)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zghg > 0 {
					var zjwh int
					var zwdg *S2
					zghg--
					zjwh, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zwdg != nil {
							zwdg.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zwdg {
								zwdg.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zwdg == nil {
								zwdg = new(S2)
							}
							bts, err = zwdg.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zjwh] = zwdg
				}
			}
		case "Myptr":
			found1ztmi[2] = true
			if nbs.AlwaysNil {
				if z.Myptr != nil {
					z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Myptr {
						z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Myptr == nil {
						z.Myptr = new(S2)
					}
					bts, err = z.Myptr.MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Myarray":
			found1ztmi[3] = true
			var zthj uint32
			zthj, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zthj != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zthj}
				return
			}
			for zvva := range z.Myarray {
				z.Myarray[zvva], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1ztmi[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zwfu uint32
				zwfu, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zwfu) {
					z.MySlice = (z.MySlice)[:zwfu]
				} else {
					z.MySlice = make([]string, zwfu)
				}
				for zdya := range z.MySlice {
					z.MySlice[zdya], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1ztmi != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder1ztmi = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1ztmi = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zkqy := range z.Slice {
		s += z.Slice[zkqy].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zjwh, zwdg := range z.Transform {
			_ = zwdg
			_ = zjwh
			s += msgp.IntSize
			if zwdg == nil {
				s += msgp.NilSize
			} else {
				s += zwdg.MSGPMsgsize()
			}
		}
	}
	s += 6
	if z.Myptr == nil {
		s += msgp.NilSize
	} else {
		s += z.Myptr.MSGPMsgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for zvva := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zvva])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zdya := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zdya])
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *S2) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 7
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *S2) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 7
	// string "beta"
	o = append(o, 0x87, 0xa4, 0x62, 0x65, 0x74, 0x61)
	o = msgp.AppendString(o, z.B)
	// string "ralph"
	o = append(o, 0xa5, 0x72, 0x61, 0x6c, 0x70, 0x68)
	o = msgp.AppendMapHeader(o, uint32(len(z.R)))
	for zams, zahj := range z.R {
		o = msgp.AppendString(o, zams)
		o = msgp.AppendUint8(o, zahj)
	}
	// string "P"
	o = append(o, 0xa1, 0x50)
	o = msgp.AppendUint16(o, z.P)
	// string "Q"
	o = append(o, 0xa1, 0x51)
	o = msgp.AppendUint32(o, z.Q)
	// string "T"
	o = append(o, 0xa1, 0x54)
	o = msgp.AppendInt64(o, z.T)
	// string "arr"
	o = append(o, 0xa3, 0x61, 0x72, 0x72)
	o = msgp.AppendArrayHeader(o, 6)
	for zeki := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zeki])
	}
	// string "MyTree"
	o = append(o, 0xa6, 0x4d, 0x79, 0x54, 0x72, 0x65, 0x65)
	if z.MyTree == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.MyTree.MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *S2) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *S2) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields2zumg = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zumg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zumg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2zumg := totalEncodedFields2zumg
	missingFieldsLeft2zumg := maxFields2zumg - totalEncodedFields2zumg

	var nextMiss2zumg int32 = -1
	var found2zumg [maxFields2zumg]bool
	var curField2zumg string

doneWithStruct2zumg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zumg > 0 || missingFieldsLeft2zumg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zumg, missingFieldsLeft2zumg, msgp.ShowFound(found2zumg[:]), unmarshalMsgFieldOrder2zumg)
		if encodedFieldsLeft2zumg > 0 {
			encodedFieldsLeft2zumg--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2zumg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zumg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zumg = 0
			}
			for nextMiss2zumg < maxFields2zumg && (found2zumg[nextMiss2zumg] || unmarshalMsgFieldSkip2zumg[nextMiss2zumg]) {
				nextMiss2zumg++
			}
			if nextMiss2zumg == maxFields2zumg {
				// filled all the empty fields!
				break doneWithStruct2zumg
			}
			missingFieldsLeft2zumg--
			curField2zumg = unmarshalMsgFieldOrder2zumg[nextMiss2zumg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zumg)
		switch curField2zumg {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zumg[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2zumg[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zhgt uint32
				zhgt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && zhgt > 0 {
					z.R = make(map[string]uint8, zhgt)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zhgt > 0 {
					var zams string
					var zahj uint8
					zhgt--
					zams, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zahj, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zams] = zahj
				}
			}
		case "P":
			found2zumg[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2zumg[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2zumg[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2zumg[6] = true
			var zzio uint32
			zzio, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zzio != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zzio}
				return
			}
			for zeki := range z.Arr {
				z.Arr[zeki], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2zumg[7] = true
			if nbs.AlwaysNil {
				if z.MyTree != nil {
					z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.MyTree {
						z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.MyTree == nil {
						z.MyTree = new(Tree)
					}
					bts, err = z.MyTree.MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss2zumg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder2zumg = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zumg = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zams, zahj := range z.R {
			_ = zahj
			_ = zams
			s += msgp.StringPrefixSize + len(zams) + msgp.Uint8Size
		}
	}
	s += 2 + msgp.Uint16Size + 2 + msgp.Uint32Size + 2 + msgp.Int64Size + 4 + msgp.ArrayHeaderSize + (6 * (msgp.Float64Size)) + 7
	if z.MyTree == nil {
		s += msgp.NilSize
	} else {
		s += z.MyTree.MSGPMsgsize()
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z Sys) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 1
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z Sys) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 1
	// string "F"
	o = append(o, 0x81, 0xa1, 0x46)
	o, err = msgp.AppendIntf(o, z.F)
	if err != nil {
		panic(err)
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Sys) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Sys) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zwwe = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zwwe uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zwwe, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zwwe := totalEncodedFields3zwwe
	missingFieldsLeft3zwwe := maxFields3zwwe - totalEncodedFields3zwwe

	var nextMiss3zwwe int32 = -1
	var found3zwwe [maxFields3zwwe]bool
	var curField3zwwe string

doneWithStruct3zwwe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zwwe > 0 || missingFieldsLeft3zwwe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zwwe, missingFieldsLeft3zwwe, msgp.ShowFound(found3zwwe[:]), unmarshalMsgFieldOrder3zwwe)
		if encodedFieldsLeft3zwwe > 0 {
			encodedFieldsLeft3zwwe--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zwwe = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zwwe < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zwwe = 0
			}
			for nextMiss3zwwe < maxFields3zwwe && (found3zwwe[nextMiss3zwwe] || unmarshalMsgFieldSkip3zwwe[nextMiss3zwwe]) {
				nextMiss3zwwe++
			}
			if nextMiss3zwwe == maxFields3zwwe {
				// filled all the empty fields!
				break doneWithStruct3zwwe
			}
			missingFieldsLeft3zwwe--
			curField3zwwe = unmarshalMsgFieldOrder3zwwe[nextMiss3zwwe]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zwwe)
		switch curField3zwwe {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zwwe[0] = true
			z.F, bts, err = nbs.ReadIntfBytes(bts)

			if err != nil {
				panic(err)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss3zwwe != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder3zwwe = []string{"F"}

var unmarshalMsgFieldSkip3zwwe = []bool{false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) MSGPMsgsize() (s int) {
	s = 1 + 2 + msgp.GuessSize(z.F)
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Tree) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Tree) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 3
	// string "Chld"
	o = append(o, 0x83, 0xa4, 0x43, 0x68, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
	for zuvu := range z.Chld {
		o, err = z.Chld[zuvu].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	// string "Par"
	o = append(o, 0xa3, 0x50, 0x61, 0x72)
	if z.Par == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Par.MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Tree) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Tree) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zwtj = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zwtj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zwtj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zwtj := totalEncodedFields4zwtj
	missingFieldsLeft4zwtj := maxFields4zwtj - totalEncodedFields4zwtj

	var nextMiss4zwtj int32 = -1
	var found4zwtj [maxFields4zwtj]bool
	var curField4zwtj string

doneWithStruct4zwtj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zwtj > 0 || missingFieldsLeft4zwtj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zwtj, missingFieldsLeft4zwtj, msgp.ShowFound(found4zwtj[:]), unmarshalMsgFieldOrder4zwtj)
		if encodedFieldsLeft4zwtj > 0 {
			encodedFieldsLeft4zwtj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zwtj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zwtj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zwtj = 0
			}
			for nextMiss4zwtj < maxFields4zwtj && (found4zwtj[nextMiss4zwtj] || unmarshalMsgFieldSkip4zwtj[nextMiss4zwtj]) {
				nextMiss4zwtj++
			}
			if nextMiss4zwtj == maxFields4zwtj {
				// filled all the empty fields!
				break doneWithStruct4zwtj
			}
			missingFieldsLeft4zwtj--
			curField4zwtj = unmarshalMsgFieldOrder4zwtj[nextMiss4zwtj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zwtj)
		switch curField4zwtj {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zwtj[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zkgl uint32
				zkgl, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zkgl) {
					z.Chld = (z.Chld)[:zkgl]
				} else {
					z.Chld = make([]Tree, zkgl)
				}
				for zuvu := range z.Chld {
					bts, err = z.Chld[zuvu].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zwtj[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zwtj[2] = true
			if nbs.AlwaysNil {
				if z.Par != nil {
					z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Par {
						z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Par == nil {
						z.Par = new(S2)
					}
					bts, err = z.Par.MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss4zwtj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zwtj = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zwtj = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zuvu := range z.Chld {
		s += z.Chld[zuvu].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
