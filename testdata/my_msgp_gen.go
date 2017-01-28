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
	const maxFields0ziaw = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0ziaw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0ziaw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0ziaw := totalEncodedFields0ziaw
	missingFieldsLeft0ziaw := maxFields0ziaw - totalEncodedFields0ziaw

	var nextMiss0ziaw int32 = -1
	var found0ziaw [maxFields0ziaw]bool
	var curField0ziaw string

doneWithStruct0ziaw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ziaw > 0 || missingFieldsLeft0ziaw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ziaw, missingFieldsLeft0ziaw, msgp.ShowFound(found0ziaw[:]), unmarshalMsgFieldOrder0ziaw)
		if encodedFieldsLeft0ziaw > 0 {
			encodedFieldsLeft0ziaw--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0ziaw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ziaw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0ziaw = 0
			}
			for nextMiss0ziaw < maxFields0ziaw && (found0ziaw[nextMiss0ziaw] || unmarshalMsgFieldSkip0ziaw[nextMiss0ziaw]) {
				nextMiss0ziaw++
			}
			if nextMiss0ziaw == maxFields0ziaw {
				// filled all the empty fields!
				break doneWithStruct0ziaw
			}
			missingFieldsLeft0ziaw--
			curField0ziaw = unmarshalMsgFieldOrder0ziaw[nextMiss0ziaw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ziaw)
		switch curField0ziaw {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0ziaw[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0ziaw[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0ziaw[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0ziaw[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0ziaw[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0ziaw[5] = true
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
	if nextMiss0ziaw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0ziaw = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0ziaw = []bool{false, false, false, false, false, false}

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
	for zkmp := range z.Slice {
		o, err = z.Slice[zkmp].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zzrj, znmd := range z.Transform {
		o = msgp.AppendInt(o, zzrj)
		if znmd == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = znmd.MSGPMarshalMsg(o)
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
	for zggp := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zggp])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zlfk := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zlfk])
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
	const maxFields1zkuw = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zkuw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zkuw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zkuw := totalEncodedFields1zkuw
	missingFieldsLeft1zkuw := maxFields1zkuw - totalEncodedFields1zkuw

	var nextMiss1zkuw int32 = -1
	var found1zkuw [maxFields1zkuw]bool
	var curField1zkuw string

doneWithStruct1zkuw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zkuw > 0 || missingFieldsLeft1zkuw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zkuw, missingFieldsLeft1zkuw, msgp.ShowFound(found1zkuw[:]), unmarshalMsgFieldOrder1zkuw)
		if encodedFieldsLeft1zkuw > 0 {
			encodedFieldsLeft1zkuw--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zkuw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zkuw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zkuw = 0
			}
			for nextMiss1zkuw < maxFields1zkuw && (found1zkuw[nextMiss1zkuw] || unmarshalMsgFieldSkip1zkuw[nextMiss1zkuw]) {
				nextMiss1zkuw++
			}
			if nextMiss1zkuw == maxFields1zkuw {
				// filled all the empty fields!
				break doneWithStruct1zkuw
			}
			missingFieldsLeft1zkuw--
			curField1zkuw = unmarshalMsgFieldOrder1zkuw[nextMiss1zkuw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zkuw)
		switch curField1zkuw {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zkuw[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zcwq uint32
				zcwq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(zcwq) {
					z.Slice = (z.Slice)[:zcwq]
				} else {
					z.Slice = make([]S2, zcwq)
				}
				for zkmp := range z.Slice {
					bts, err = z.Slice[zkmp].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1zkuw[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zfwz uint32
				zfwz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zfwz > 0 {
					z.Transform = make(map[int]*S2, zfwz)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zfwz > 0 {
					var zzrj int
					var znmd *S2
					zfwz--
					zzrj, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if znmd != nil {
							znmd.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != znmd {
								znmd.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if znmd == nil {
								znmd = new(S2)
							}
							bts, err = znmd.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zzrj] = znmd
				}
			}
		case "Myptr":
			found1zkuw[2] = true
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
			found1zkuw[3] = true
			var zgar uint32
			zgar, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zgar != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zgar}
				return
			}
			for zggp := range z.Myarray {
				z.Myarray[zggp], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1zkuw[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zudd uint32
				zudd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zudd) {
					z.MySlice = (z.MySlice)[:zudd]
				} else {
					z.MySlice = make([]string, zudd)
				}
				for zlfk := range z.MySlice {
					z.MySlice[zlfk], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zkuw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder1zkuw = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zkuw = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zkmp := range z.Slice {
		s += z.Slice[zkmp].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zzrj, znmd := range z.Transform {
			_ = znmd
			_ = zzrj
			s += msgp.IntSize
			if znmd == nil {
				s += msgp.NilSize
			} else {
				s += znmd.MSGPMsgsize()
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
	for zggp := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zggp])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zlfk := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zlfk])
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
	for zwcp, zxct := range z.R {
		o = msgp.AppendString(o, zwcp)
		o = msgp.AppendUint8(o, zxct)
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
	for zjuo := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zjuo])
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
	const maxFields2ztbp = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2ztbp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2ztbp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2ztbp := totalEncodedFields2ztbp
	missingFieldsLeft2ztbp := maxFields2ztbp - totalEncodedFields2ztbp

	var nextMiss2ztbp int32 = -1
	var found2ztbp [maxFields2ztbp]bool
	var curField2ztbp string

doneWithStruct2ztbp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2ztbp > 0 || missingFieldsLeft2ztbp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2ztbp, missingFieldsLeft2ztbp, msgp.ShowFound(found2ztbp[:]), unmarshalMsgFieldOrder2ztbp)
		if encodedFieldsLeft2ztbp > 0 {
			encodedFieldsLeft2ztbp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2ztbp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2ztbp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2ztbp = 0
			}
			for nextMiss2ztbp < maxFields2ztbp && (found2ztbp[nextMiss2ztbp] || unmarshalMsgFieldSkip2ztbp[nextMiss2ztbp]) {
				nextMiss2ztbp++
			}
			if nextMiss2ztbp == maxFields2ztbp {
				// filled all the empty fields!
				break doneWithStruct2ztbp
			}
			missingFieldsLeft2ztbp--
			curField2ztbp = unmarshalMsgFieldOrder2ztbp[nextMiss2ztbp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2ztbp)
		switch curField2ztbp {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2ztbp[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2ztbp[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var znqw uint32
				znqw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && znqw > 0 {
					z.R = make(map[string]uint8, znqw)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for znqw > 0 {
					var zwcp string
					var zxct uint8
					znqw--
					zwcp, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zxct, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zwcp] = zxct
				}
			}
		case "P":
			found2ztbp[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2ztbp[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2ztbp[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2ztbp[6] = true
			var zuup uint32
			zuup, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zuup != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zuup}
				return
			}
			for zjuo := range z.Arr {
				z.Arr[zjuo], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2ztbp[7] = true
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
	if nextMiss2ztbp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder2ztbp = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2ztbp = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zwcp, zxct := range z.R {
			_ = zxct
			_ = zwcp
			s += msgp.StringPrefixSize + len(zwcp) + msgp.Uint8Size
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
	const maxFields3zuvy = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zuvy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zuvy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zuvy := totalEncodedFields3zuvy
	missingFieldsLeft3zuvy := maxFields3zuvy - totalEncodedFields3zuvy

	var nextMiss3zuvy int32 = -1
	var found3zuvy [maxFields3zuvy]bool
	var curField3zuvy string

doneWithStruct3zuvy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zuvy > 0 || missingFieldsLeft3zuvy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zuvy, missingFieldsLeft3zuvy, msgp.ShowFound(found3zuvy[:]), unmarshalMsgFieldOrder3zuvy)
		if encodedFieldsLeft3zuvy > 0 {
			encodedFieldsLeft3zuvy--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zuvy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zuvy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zuvy = 0
			}
			for nextMiss3zuvy < maxFields3zuvy && (found3zuvy[nextMiss3zuvy] || unmarshalMsgFieldSkip3zuvy[nextMiss3zuvy]) {
				nextMiss3zuvy++
			}
			if nextMiss3zuvy == maxFields3zuvy {
				// filled all the empty fields!
				break doneWithStruct3zuvy
			}
			missingFieldsLeft3zuvy--
			curField3zuvy = unmarshalMsgFieldOrder3zuvy[nextMiss3zuvy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zuvy)
		switch curField3zuvy {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zuvy[0] = true
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
	if nextMiss3zuvy != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder3zuvy = []string{"F"}

var unmarshalMsgFieldSkip3zuvy = []bool{false}

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
	for zxnr := range z.Chld {
		o, err = z.Chld[zxnr].MSGPMarshalMsg(o)
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
	const maxFields4zifq = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zifq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zifq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zifq := totalEncodedFields4zifq
	missingFieldsLeft4zifq := maxFields4zifq - totalEncodedFields4zifq

	var nextMiss4zifq int32 = -1
	var found4zifq [maxFields4zifq]bool
	var curField4zifq string

doneWithStruct4zifq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zifq > 0 || missingFieldsLeft4zifq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zifq, missingFieldsLeft4zifq, msgp.ShowFound(found4zifq[:]), unmarshalMsgFieldOrder4zifq)
		if encodedFieldsLeft4zifq > 0 {
			encodedFieldsLeft4zifq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zifq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zifq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zifq = 0
			}
			for nextMiss4zifq < maxFields4zifq && (found4zifq[nextMiss4zifq] || unmarshalMsgFieldSkip4zifq[nextMiss4zifq]) {
				nextMiss4zifq++
			}
			if nextMiss4zifq == maxFields4zifq {
				// filled all the empty fields!
				break doneWithStruct4zifq
			}
			missingFieldsLeft4zifq--
			curField4zifq = unmarshalMsgFieldOrder4zifq[nextMiss4zifq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zifq)
		switch curField4zifq {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zifq[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zpwb uint32
				zpwb, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zpwb) {
					z.Chld = (z.Chld)[:zpwb]
				} else {
					z.Chld = make([]Tree, zpwb)
				}
				for zxnr := range z.Chld {
					bts, err = z.Chld[zxnr].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zifq[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zifq[2] = true
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
	if nextMiss4zifq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zifq = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zifq = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zxnr := range z.Chld {
		s += z.Chld[zxnr].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
