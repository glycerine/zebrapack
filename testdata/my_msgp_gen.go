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
	const maxFields0zcwq = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zcwq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zcwq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zcwq := totalEncodedFields0zcwq
	missingFieldsLeft0zcwq := maxFields0zcwq - totalEncodedFields0zcwq

	var nextMiss0zcwq int32 = -1
	var found0zcwq [maxFields0zcwq]bool
	var curField0zcwq string

doneWithStruct0zcwq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zcwq > 0 || missingFieldsLeft0zcwq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zcwq, missingFieldsLeft0zcwq, msgp.ShowFound(found0zcwq[:]), unmarshalMsgFieldOrder0zcwq)
		if encodedFieldsLeft0zcwq > 0 {
			encodedFieldsLeft0zcwq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zcwq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zcwq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zcwq = 0
			}
			for nextMiss0zcwq < maxFields0zcwq && (found0zcwq[nextMiss0zcwq] || unmarshalMsgFieldSkip0zcwq[nextMiss0zcwq]) {
				nextMiss0zcwq++
			}
			if nextMiss0zcwq == maxFields0zcwq {
				// filled all the empty fields!
				break doneWithStruct0zcwq
			}
			missingFieldsLeft0zcwq--
			curField0zcwq = unmarshalMsgFieldOrder0zcwq[nextMiss0zcwq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zcwq)
		switch curField0zcwq {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zcwq[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0zcwq[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0zcwq[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0zcwq[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0zcwq[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0zcwq[5] = true
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
	if nextMiss0zcwq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zcwq = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zcwq = []bool{false, false, false, false, false, false}

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
	for zdiv := range z.Slice {
		o, err = z.Slice[zdiv].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zthr, zydc := range z.Transform {
		o = msgp.AppendInt(o, zthr)
		if zydc == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zydc.MSGPMarshalMsg(o)
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
	for zmnd := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zmnd])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zlmh := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zlmh])
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
	const maxFields1zlly = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zlly uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zlly, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zlly := totalEncodedFields1zlly
	missingFieldsLeft1zlly := maxFields1zlly - totalEncodedFields1zlly

	var nextMiss1zlly int32 = -1
	var found1zlly [maxFields1zlly]bool
	var curField1zlly string

doneWithStruct1zlly:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zlly > 0 || missingFieldsLeft1zlly > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zlly, missingFieldsLeft1zlly, msgp.ShowFound(found1zlly[:]), unmarshalMsgFieldOrder1zlly)
		if encodedFieldsLeft1zlly > 0 {
			encodedFieldsLeft1zlly--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zlly = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zlly < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zlly = 0
			}
			for nextMiss1zlly < maxFields1zlly && (found1zlly[nextMiss1zlly] || unmarshalMsgFieldSkip1zlly[nextMiss1zlly]) {
				nextMiss1zlly++
			}
			if nextMiss1zlly == maxFields1zlly {
				// filled all the empty fields!
				break doneWithStruct1zlly
			}
			missingFieldsLeft1zlly--
			curField1zlly = unmarshalMsgFieldOrder1zlly[nextMiss1zlly]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zlly)
		switch curField1zlly {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zlly[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zhem uint32
				zhem, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(zhem) {
					z.Slice = (z.Slice)[:zhem]
				} else {
					z.Slice = make([]S2, zhem)
				}
				for zdiv := range z.Slice {
					bts, err = z.Slice[zdiv].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1zlly[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zpbo uint32
				zpbo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zpbo > 0 {
					z.Transform = make(map[int]*S2, zpbo)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zpbo > 0 {
					var zthr int
					var zydc *S2
					zpbo--
					zthr, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zydc != nil {
							zydc.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zydc {
								zydc.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zydc == nil {
								zydc = new(S2)
							}
							bts, err = zydc.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zthr] = zydc
				}
			}
		case "Myptr":
			found1zlly[2] = true
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
			found1zlly[3] = true
			var zhou uint32
			zhou, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zhou != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zhou}
				return
			}
			for zmnd := range z.Myarray {
				z.Myarray[zmnd], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1zlly[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zmgr uint32
				zmgr, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zmgr) {
					z.MySlice = (z.MySlice)[:zmgr]
				} else {
					z.MySlice = make([]string, zmgr)
				}
				for zlmh := range z.MySlice {
					z.MySlice[zlmh], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zlly != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder1zlly = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zlly = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zdiv := range z.Slice {
		s += z.Slice[zdiv].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zthr, zydc := range z.Transform {
			_ = zydc
			_ = zthr
			s += msgp.IntSize
			if zydc == nil {
				s += msgp.NilSize
			} else {
				s += zydc.MSGPMsgsize()
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
	for zmnd := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zmnd])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zlmh := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zlmh])
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
	for zoge, zlno := range z.R {
		o = msgp.AppendString(o, zoge)
		o = msgp.AppendUint8(o, zlno)
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
	for zfve := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zfve])
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
	const maxFields2ztkt = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2ztkt uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2ztkt, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2ztkt := totalEncodedFields2ztkt
	missingFieldsLeft2ztkt := maxFields2ztkt - totalEncodedFields2ztkt

	var nextMiss2ztkt int32 = -1
	var found2ztkt [maxFields2ztkt]bool
	var curField2ztkt string

doneWithStruct2ztkt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2ztkt > 0 || missingFieldsLeft2ztkt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2ztkt, missingFieldsLeft2ztkt, msgp.ShowFound(found2ztkt[:]), unmarshalMsgFieldOrder2ztkt)
		if encodedFieldsLeft2ztkt > 0 {
			encodedFieldsLeft2ztkt--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2ztkt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2ztkt < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2ztkt = 0
			}
			for nextMiss2ztkt < maxFields2ztkt && (found2ztkt[nextMiss2ztkt] || unmarshalMsgFieldSkip2ztkt[nextMiss2ztkt]) {
				nextMiss2ztkt++
			}
			if nextMiss2ztkt == maxFields2ztkt {
				// filled all the empty fields!
				break doneWithStruct2ztkt
			}
			missingFieldsLeft2ztkt--
			curField2ztkt = unmarshalMsgFieldOrder2ztkt[nextMiss2ztkt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2ztkt)
		switch curField2ztkt {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2ztkt[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2ztkt[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zaae uint32
				zaae, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && zaae > 0 {
					z.R = make(map[string]uint8, zaae)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zaae > 0 {
					var zoge string
					var zlno uint8
					zaae--
					zoge, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zlno, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zoge] = zlno
				}
			}
		case "P":
			found2ztkt[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2ztkt[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2ztkt[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2ztkt[6] = true
			var zvsm uint32
			zvsm, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zvsm != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zvsm}
				return
			}
			for zfve := range z.Arr {
				z.Arr[zfve], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2ztkt[7] = true
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
	if nextMiss2ztkt != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder2ztkt = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2ztkt = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zoge, zlno := range z.R {
			_ = zlno
			_ = zoge
			s += msgp.StringPrefixSize + len(zoge) + msgp.Uint8Size
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
	const maxFields3zesb = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zesb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zesb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zesb := totalEncodedFields3zesb
	missingFieldsLeft3zesb := maxFields3zesb - totalEncodedFields3zesb

	var nextMiss3zesb int32 = -1
	var found3zesb [maxFields3zesb]bool
	var curField3zesb string

doneWithStruct3zesb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zesb > 0 || missingFieldsLeft3zesb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zesb, missingFieldsLeft3zesb, msgp.ShowFound(found3zesb[:]), unmarshalMsgFieldOrder3zesb)
		if encodedFieldsLeft3zesb > 0 {
			encodedFieldsLeft3zesb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zesb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zesb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zesb = 0
			}
			for nextMiss3zesb < maxFields3zesb && (found3zesb[nextMiss3zesb] || unmarshalMsgFieldSkip3zesb[nextMiss3zesb]) {
				nextMiss3zesb++
			}
			if nextMiss3zesb == maxFields3zesb {
				// filled all the empty fields!
				break doneWithStruct3zesb
			}
			missingFieldsLeft3zesb--
			curField3zesb = unmarshalMsgFieldOrder3zesb[nextMiss3zesb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zesb)
		switch curField3zesb {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zesb[0] = true
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
	if nextMiss3zesb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder3zesb = []string{"F"}

var unmarshalMsgFieldSkip3zesb = []bool{false}

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
	for zukc := range z.Chld {
		o, err = z.Chld[zukc].MSGPMarshalMsg(o)
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
	const maxFields4zexa = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zexa uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zexa, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zexa := totalEncodedFields4zexa
	missingFieldsLeft4zexa := maxFields4zexa - totalEncodedFields4zexa

	var nextMiss4zexa int32 = -1
	var found4zexa [maxFields4zexa]bool
	var curField4zexa string

doneWithStruct4zexa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zexa > 0 || missingFieldsLeft4zexa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zexa, missingFieldsLeft4zexa, msgp.ShowFound(found4zexa[:]), unmarshalMsgFieldOrder4zexa)
		if encodedFieldsLeft4zexa > 0 {
			encodedFieldsLeft4zexa--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zexa = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zexa < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zexa = 0
			}
			for nextMiss4zexa < maxFields4zexa && (found4zexa[nextMiss4zexa] || unmarshalMsgFieldSkip4zexa[nextMiss4zexa]) {
				nextMiss4zexa++
			}
			if nextMiss4zexa == maxFields4zexa {
				// filled all the empty fields!
				break doneWithStruct4zexa
			}
			missingFieldsLeft4zexa--
			curField4zexa = unmarshalMsgFieldOrder4zexa[nextMiss4zexa]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zexa)
		switch curField4zexa {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zexa[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zupq uint32
				zupq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zupq) {
					z.Chld = (z.Chld)[:zupq]
				} else {
					z.Chld = make([]Tree, zupq)
				}
				for zukc := range z.Chld {
					bts, err = z.Chld[zukc].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zexa[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zexa[2] = true
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
	if nextMiss4zexa != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zexa = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zexa = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zukc := range z.Chld {
		s += z.Chld[zukc].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
