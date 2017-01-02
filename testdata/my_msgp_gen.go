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
	const maxFields0zhli = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zhli uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zhli, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zhli := totalEncodedFields0zhli
	missingFieldsLeft0zhli := maxFields0zhli - totalEncodedFields0zhli

	var nextMiss0zhli int32 = -1
	var found0zhli [maxFields0zhli]bool
	var curField0zhli string

doneWithStruct0zhli:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zhli > 0 || missingFieldsLeft0zhli > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zhli, missingFieldsLeft0zhli, msgp.ShowFound(found0zhli[:]), unmarshalMsgFieldOrder0zhli)
		if encodedFieldsLeft0zhli > 0 {
			encodedFieldsLeft0zhli--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zhli = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zhli < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zhli = 0
			}
			for nextMiss0zhli < maxFields0zhli && (found0zhli[nextMiss0zhli] || unmarshalMsgFieldSkip0zhli[nextMiss0zhli]) {
				nextMiss0zhli++
			}
			if nextMiss0zhli == maxFields0zhli {
				// filled all the empty fields!
				break doneWithStruct0zhli
			}
			missingFieldsLeft0zhli--
			curField0zhli = unmarshalMsgFieldOrder0zhli[nextMiss0zhli]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zhli)
		switch curField0zhli {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zhli[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0zhli[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0zhli[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0zhli[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0zhli[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0zhli[5] = true
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
	if nextMiss0zhli != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zhli = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zhli = []bool{false, false, false, false, false, false}

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
	for zwls := range z.Slice {
		o, err = z.Slice[zwls].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zvjb, zcvr := range z.Transform {
		o = msgp.AppendInt(o, zvjb)
		if zcvr == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zcvr.MSGPMarshalMsg(o)
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
	for zfvs := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zfvs])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zghi := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zghi])
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
	const maxFields1zqwv = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zqwv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zqwv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zqwv := totalEncodedFields1zqwv
	missingFieldsLeft1zqwv := maxFields1zqwv - totalEncodedFields1zqwv

	var nextMiss1zqwv int32 = -1
	var found1zqwv [maxFields1zqwv]bool
	var curField1zqwv string

doneWithStruct1zqwv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zqwv > 0 || missingFieldsLeft1zqwv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zqwv, missingFieldsLeft1zqwv, msgp.ShowFound(found1zqwv[:]), unmarshalMsgFieldOrder1zqwv)
		if encodedFieldsLeft1zqwv > 0 {
			encodedFieldsLeft1zqwv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zqwv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zqwv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zqwv = 0
			}
			for nextMiss1zqwv < maxFields1zqwv && (found1zqwv[nextMiss1zqwv] || unmarshalMsgFieldSkip1zqwv[nextMiss1zqwv]) {
				nextMiss1zqwv++
			}
			if nextMiss1zqwv == maxFields1zqwv {
				// filled all the empty fields!
				break doneWithStruct1zqwv
			}
			missingFieldsLeft1zqwv--
			curField1zqwv = unmarshalMsgFieldOrder1zqwv[nextMiss1zqwv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zqwv)
		switch curField1zqwv {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zqwv[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zrtk uint32
				zrtk, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(zrtk) {
					z.Slice = (z.Slice)[:zrtk]
				} else {
					z.Slice = make([]S2, zrtk)
				}
				for zwls := range z.Slice {
					bts, err = z.Slice[zwls].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1zqwv[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zueq uint32
				zueq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zueq > 0 {
					z.Transform = make(map[int]*S2, zueq)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zueq > 0 {
					var zvjb int
					var zcvr *S2
					zueq--
					zvjb, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zcvr != nil {
							zcvr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zcvr {
								zcvr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zcvr == nil {
								zcvr = new(S2)
							}
							bts, err = zcvr.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zvjb] = zcvr
				}
			}
		case "Myptr":
			found1zqwv[2] = true
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
			found1zqwv[3] = true
			var zuxf uint32
			zuxf, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zuxf != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zuxf}
				return
			}
			for zfvs := range z.Myarray {
				z.Myarray[zfvs], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1zqwv[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zndo uint32
				zndo, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zndo) {
					z.MySlice = (z.MySlice)[:zndo]
				} else {
					z.MySlice = make([]string, zndo)
				}
				for zghi := range z.MySlice {
					z.MySlice[zghi], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zqwv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder1zqwv = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zqwv = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zwls := range z.Slice {
		s += z.Slice[zwls].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zvjb, zcvr := range z.Transform {
			_ = zcvr
			_ = zvjb
			s += msgp.IntSize
			if zcvr == nil {
				s += msgp.NilSize
			} else {
				s += zcvr.MSGPMsgsize()
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
	for zfvs := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zfvs])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zghi := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zghi])
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
	for zxvc, ztvp := range z.R {
		o = msgp.AppendString(o, zxvc)
		o = msgp.AppendUint8(o, ztvp)
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
	for zymm := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zymm])
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
	const maxFields2zqws = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zqws uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zqws, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2zqws := totalEncodedFields2zqws
	missingFieldsLeft2zqws := maxFields2zqws - totalEncodedFields2zqws

	var nextMiss2zqws int32 = -1
	var found2zqws [maxFields2zqws]bool
	var curField2zqws string

doneWithStruct2zqws:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zqws > 0 || missingFieldsLeft2zqws > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zqws, missingFieldsLeft2zqws, msgp.ShowFound(found2zqws[:]), unmarshalMsgFieldOrder2zqws)
		if encodedFieldsLeft2zqws > 0 {
			encodedFieldsLeft2zqws--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2zqws = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zqws < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zqws = 0
			}
			for nextMiss2zqws < maxFields2zqws && (found2zqws[nextMiss2zqws] || unmarshalMsgFieldSkip2zqws[nextMiss2zqws]) {
				nextMiss2zqws++
			}
			if nextMiss2zqws == maxFields2zqws {
				// filled all the empty fields!
				break doneWithStruct2zqws
			}
			missingFieldsLeft2zqws--
			curField2zqws = unmarshalMsgFieldOrder2zqws[nextMiss2zqws]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zqws)
		switch curField2zqws {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zqws[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2zqws[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zsqa uint32
				zsqa, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && zsqa > 0 {
					z.R = make(map[string]uint8, zsqa)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zsqa > 0 {
					var zxvc string
					var ztvp uint8
					zsqa--
					zxvc, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					ztvp, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zxvc] = ztvp
				}
			}
		case "P":
			found2zqws[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2zqws[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2zqws[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2zqws[6] = true
			var zejj uint32
			zejj, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zejj != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zejj}
				return
			}
			for zymm := range z.Arr {
				z.Arr[zymm], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2zqws[7] = true
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
	if nextMiss2zqws != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder2zqws = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zqws = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zxvc, ztvp := range z.R {
			_ = ztvp
			_ = zxvc
			s += msgp.StringPrefixSize + len(zxvc) + msgp.Uint8Size
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
	const maxFields3zdbn = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zdbn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zdbn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zdbn := totalEncodedFields3zdbn
	missingFieldsLeft3zdbn := maxFields3zdbn - totalEncodedFields3zdbn

	var nextMiss3zdbn int32 = -1
	var found3zdbn [maxFields3zdbn]bool
	var curField3zdbn string

doneWithStruct3zdbn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zdbn > 0 || missingFieldsLeft3zdbn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zdbn, missingFieldsLeft3zdbn, msgp.ShowFound(found3zdbn[:]), unmarshalMsgFieldOrder3zdbn)
		if encodedFieldsLeft3zdbn > 0 {
			encodedFieldsLeft3zdbn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zdbn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zdbn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zdbn = 0
			}
			for nextMiss3zdbn < maxFields3zdbn && (found3zdbn[nextMiss3zdbn] || unmarshalMsgFieldSkip3zdbn[nextMiss3zdbn]) {
				nextMiss3zdbn++
			}
			if nextMiss3zdbn == maxFields3zdbn {
				// filled all the empty fields!
				break doneWithStruct3zdbn
			}
			missingFieldsLeft3zdbn--
			curField3zdbn = unmarshalMsgFieldOrder3zdbn[nextMiss3zdbn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zdbn)
		switch curField3zdbn {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zdbn[0] = true
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
	if nextMiss3zdbn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder3zdbn = []string{"F"}

var unmarshalMsgFieldSkip3zdbn = []bool{false}

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
	for zebq := range z.Chld {
		o, err = z.Chld[zebq].MSGPMarshalMsg(o)
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
	const maxFields4zluj = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zluj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zluj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zluj := totalEncodedFields4zluj
	missingFieldsLeft4zluj := maxFields4zluj - totalEncodedFields4zluj

	var nextMiss4zluj int32 = -1
	var found4zluj [maxFields4zluj]bool
	var curField4zluj string

doneWithStruct4zluj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zluj > 0 || missingFieldsLeft4zluj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zluj, missingFieldsLeft4zluj, msgp.ShowFound(found4zluj[:]), unmarshalMsgFieldOrder4zluj)
		if encodedFieldsLeft4zluj > 0 {
			encodedFieldsLeft4zluj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zluj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zluj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zluj = 0
			}
			for nextMiss4zluj < maxFields4zluj && (found4zluj[nextMiss4zluj] || unmarshalMsgFieldSkip4zluj[nextMiss4zluj]) {
				nextMiss4zluj++
			}
			if nextMiss4zluj == maxFields4zluj {
				// filled all the empty fields!
				break doneWithStruct4zluj
			}
			missingFieldsLeft4zluj--
			curField4zluj = unmarshalMsgFieldOrder4zluj[nextMiss4zluj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zluj)
		switch curField4zluj {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zluj[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zkuj uint32
				zkuj, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zkuj) {
					z.Chld = (z.Chld)[:zkuj]
				} else {
					z.Chld = make([]Tree, zkuj)
				}
				for zebq := range z.Chld {
					bts, err = z.Chld[zebq].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zluj[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zluj[2] = true
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
	if nextMiss4zluj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zluj = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zluj = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zebq := range z.Chld {
		s += z.Chld[zebq].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
