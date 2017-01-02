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
	const maxFields0zdqq = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zdqq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zdqq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zdqq := totalEncodedFields0zdqq
	missingFieldsLeft0zdqq := maxFields0zdqq - totalEncodedFields0zdqq

	var nextMiss0zdqq int32 = -1
	var found0zdqq [maxFields0zdqq]bool
	var curField0zdqq string

doneWithStruct0zdqq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zdqq > 0 || missingFieldsLeft0zdqq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zdqq, missingFieldsLeft0zdqq, msgp.ShowFound(found0zdqq[:]), unmarshalMsgFieldOrder0zdqq)
		if encodedFieldsLeft0zdqq > 0 {
			encodedFieldsLeft0zdqq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zdqq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zdqq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zdqq = 0
			}
			for nextMiss0zdqq < maxFields0zdqq && (found0zdqq[nextMiss0zdqq] || unmarshalMsgFieldSkip0zdqq[nextMiss0zdqq]) {
				nextMiss0zdqq++
			}
			if nextMiss0zdqq == maxFields0zdqq {
				// filled all the empty fields!
				break doneWithStruct0zdqq
			}
			missingFieldsLeft0zdqq--
			curField0zdqq = unmarshalMsgFieldOrder0zdqq[nextMiss0zdqq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zdqq)
		switch curField0zdqq {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zdqq[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0zdqq[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0zdqq[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0zdqq[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0zdqq[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0zdqq[5] = true
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
	if nextMiss0zdqq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zdqq = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zdqq = []bool{false, false, false, false, false, false}

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
	for zvjy := range z.Slice {
		o, err = z.Slice[zvjy].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zhrp, zchz := range z.Transform {
		o = msgp.AppendInt(o, zhrp)
		if zchz == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zchz.MSGPMarshalMsg(o)
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
	for zpkw := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zpkw])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zgta := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zgta])
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
	const maxFields1zlzv = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zlzv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zlzv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zlzv := totalEncodedFields1zlzv
	missingFieldsLeft1zlzv := maxFields1zlzv - totalEncodedFields1zlzv

	var nextMiss1zlzv int32 = -1
	var found1zlzv [maxFields1zlzv]bool
	var curField1zlzv string

doneWithStruct1zlzv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zlzv > 0 || missingFieldsLeft1zlzv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zlzv, missingFieldsLeft1zlzv, msgp.ShowFound(found1zlzv[:]), unmarshalMsgFieldOrder1zlzv)
		if encodedFieldsLeft1zlzv > 0 {
			encodedFieldsLeft1zlzv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zlzv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zlzv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zlzv = 0
			}
			for nextMiss1zlzv < maxFields1zlzv && (found1zlzv[nextMiss1zlzv] || unmarshalMsgFieldSkip1zlzv[nextMiss1zlzv]) {
				nextMiss1zlzv++
			}
			if nextMiss1zlzv == maxFields1zlzv {
				// filled all the empty fields!
				break doneWithStruct1zlzv
			}
			missingFieldsLeft1zlzv--
			curField1zlzv = unmarshalMsgFieldOrder1zlzv[nextMiss1zlzv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zlzv)
		switch curField1zlzv {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zlzv[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zvmo uint32
				zvmo, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(zvmo) {
					z.Slice = (z.Slice)[:zvmo]
				} else {
					z.Slice = make([]S2, zvmo)
				}
				for zvjy := range z.Slice {
					bts, err = z.Slice[zvjy].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1zlzv[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zodt uint32
				zodt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zodt > 0 {
					z.Transform = make(map[int]*S2, zodt)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zodt > 0 {
					var zhrp int
					var zchz *S2
					zodt--
					zhrp, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zchz != nil {
							zchz.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zchz {
								zchz.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zchz == nil {
								zchz = new(S2)
							}
							bts, err = zchz.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zhrp] = zchz
				}
			}
		case "Myptr":
			found1zlzv[2] = true
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
			found1zlzv[3] = true
			var zhkk uint32
			zhkk, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zhkk != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zhkk}
				return
			}
			for zpkw := range z.Myarray {
				z.Myarray[zpkw], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1zlzv[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zvop uint32
				zvop, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zvop) {
					z.MySlice = (z.MySlice)[:zvop]
				} else {
					z.MySlice = make([]string, zvop)
				}
				for zgta := range z.MySlice {
					z.MySlice[zgta], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zlzv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder1zlzv = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zlzv = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zvjy := range z.Slice {
		s += z.Slice[zvjy].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zhrp, zchz := range z.Transform {
			_ = zchz
			_ = zhrp
			s += msgp.IntSize
			if zchz == nil {
				s += msgp.NilSize
			} else {
				s += zchz.MSGPMsgsize()
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
	for zpkw := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zpkw])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zgta := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zgta])
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
	for zuye, zgva := range z.R {
		o = msgp.AppendString(o, zuye)
		o = msgp.AppendUint8(o, zgva)
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
	for zsew := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zsew])
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
	const maxFields2zqex = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zqex uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zqex, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2zqex := totalEncodedFields2zqex
	missingFieldsLeft2zqex := maxFields2zqex - totalEncodedFields2zqex

	var nextMiss2zqex int32 = -1
	var found2zqex [maxFields2zqex]bool
	var curField2zqex string

doneWithStruct2zqex:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zqex > 0 || missingFieldsLeft2zqex > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zqex, missingFieldsLeft2zqex, msgp.ShowFound(found2zqex[:]), unmarshalMsgFieldOrder2zqex)
		if encodedFieldsLeft2zqex > 0 {
			encodedFieldsLeft2zqex--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2zqex = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zqex < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zqex = 0
			}
			for nextMiss2zqex < maxFields2zqex && (found2zqex[nextMiss2zqex] || unmarshalMsgFieldSkip2zqex[nextMiss2zqex]) {
				nextMiss2zqex++
			}
			if nextMiss2zqex == maxFields2zqex {
				// filled all the empty fields!
				break doneWithStruct2zqex
			}
			missingFieldsLeft2zqex--
			curField2zqex = unmarshalMsgFieldOrder2zqex[nextMiss2zqex]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zqex)
		switch curField2zqex {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zqex[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2zqex[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var znmv uint32
				znmv, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && znmv > 0 {
					z.R = make(map[string]uint8, znmv)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for znmv > 0 {
					var zuye string
					var zgva uint8
					znmv--
					zuye, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zgva, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zuye] = zgva
				}
			}
		case "P":
			found2zqex[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2zqex[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2zqex[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2zqex[6] = true
			var zwni uint32
			zwni, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zwni != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zwni}
				return
			}
			for zsew := range z.Arr {
				z.Arr[zsew], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2zqex[7] = true
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
	if nextMiss2zqex != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder2zqex = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zqex = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zuye, zgva := range z.R {
			_ = zgva
			_ = zuye
			s += msgp.StringPrefixSize + len(zuye) + msgp.Uint8Size
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
	const maxFields3zeux = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zeux uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zeux, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zeux := totalEncodedFields3zeux
	missingFieldsLeft3zeux := maxFields3zeux - totalEncodedFields3zeux

	var nextMiss3zeux int32 = -1
	var found3zeux [maxFields3zeux]bool
	var curField3zeux string

doneWithStruct3zeux:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zeux > 0 || missingFieldsLeft3zeux > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zeux, missingFieldsLeft3zeux, msgp.ShowFound(found3zeux[:]), unmarshalMsgFieldOrder3zeux)
		if encodedFieldsLeft3zeux > 0 {
			encodedFieldsLeft3zeux--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zeux = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zeux < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zeux = 0
			}
			for nextMiss3zeux < maxFields3zeux && (found3zeux[nextMiss3zeux] || unmarshalMsgFieldSkip3zeux[nextMiss3zeux]) {
				nextMiss3zeux++
			}
			if nextMiss3zeux == maxFields3zeux {
				// filled all the empty fields!
				break doneWithStruct3zeux
			}
			missingFieldsLeft3zeux--
			curField3zeux = unmarshalMsgFieldOrder3zeux[nextMiss3zeux]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zeux)
		switch curField3zeux {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zeux[0] = true
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
	if nextMiss3zeux != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder3zeux = []string{"F"}

var unmarshalMsgFieldSkip3zeux = []bool{false}

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
	for zmzd := range z.Chld {
		o, err = z.Chld[zmzd].MSGPMarshalMsg(o)
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
	const maxFields4zihr = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zihr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zihr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zihr := totalEncodedFields4zihr
	missingFieldsLeft4zihr := maxFields4zihr - totalEncodedFields4zihr

	var nextMiss4zihr int32 = -1
	var found4zihr [maxFields4zihr]bool
	var curField4zihr string

doneWithStruct4zihr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zihr > 0 || missingFieldsLeft4zihr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zihr, missingFieldsLeft4zihr, msgp.ShowFound(found4zihr[:]), unmarshalMsgFieldOrder4zihr)
		if encodedFieldsLeft4zihr > 0 {
			encodedFieldsLeft4zihr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zihr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zihr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zihr = 0
			}
			for nextMiss4zihr < maxFields4zihr && (found4zihr[nextMiss4zihr] || unmarshalMsgFieldSkip4zihr[nextMiss4zihr]) {
				nextMiss4zihr++
			}
			if nextMiss4zihr == maxFields4zihr {
				// filled all the empty fields!
				break doneWithStruct4zihr
			}
			missingFieldsLeft4zihr--
			curField4zihr = unmarshalMsgFieldOrder4zihr[nextMiss4zihr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zihr)
		switch curField4zihr {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zihr[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zwzt uint32
				zwzt, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zwzt) {
					z.Chld = (z.Chld)[:zwzt]
				} else {
					z.Chld = make([]Tree, zwzt)
				}
				for zmzd := range z.Chld {
					bts, err = z.Chld[zmzd].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zihr[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zihr[2] = true
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
	if nextMiss4zihr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zihr = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zihr = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zmzd := range z.Chld {
		s += z.Chld[zmzd].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
