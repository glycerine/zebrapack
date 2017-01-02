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
	const maxFields0zxrf = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zxrf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zxrf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zxrf := totalEncodedFields0zxrf
	missingFieldsLeft0zxrf := maxFields0zxrf - totalEncodedFields0zxrf

	var nextMiss0zxrf int32 = -1
	var found0zxrf [maxFields0zxrf]bool
	var curField0zxrf string

doneWithStruct0zxrf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zxrf > 0 || missingFieldsLeft0zxrf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zxrf, missingFieldsLeft0zxrf, msgp.ShowFound(found0zxrf[:]), unmarshalMsgFieldOrder0zxrf)
		if encodedFieldsLeft0zxrf > 0 {
			encodedFieldsLeft0zxrf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zxrf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zxrf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zxrf = 0
			}
			for nextMiss0zxrf < maxFields0zxrf && (found0zxrf[nextMiss0zxrf] || unmarshalMsgFieldSkip0zxrf[nextMiss0zxrf]) {
				nextMiss0zxrf++
			}
			if nextMiss0zxrf == maxFields0zxrf {
				// filled all the empty fields!
				break doneWithStruct0zxrf
			}
			missingFieldsLeft0zxrf--
			curField0zxrf = unmarshalMsgFieldOrder0zxrf[nextMiss0zxrf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zxrf)
		switch curField0zxrf {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zxrf[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0zxrf[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0zxrf[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0zxrf[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0zxrf[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0zxrf[5] = true
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
	if nextMiss0zxrf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of A
var unmarshalMsgFieldOrder0zxrf = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zxrf = []bool{false, false, false, false, false, false}

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
	for zqhs := range z.Slice {
		o, err = z.Slice[zqhs].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zlzv, zsia := range z.Transform {
		o = msgp.AppendInt(o, zlzv)
		if zsia == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zsia.MSGPMarshalMsg(o)
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
	for zhza := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zhza])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zdtv := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zdtv])
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
	const maxFields1znuo = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1znuo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1znuo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1znuo := totalEncodedFields1znuo
	missingFieldsLeft1znuo := maxFields1znuo - totalEncodedFields1znuo

	var nextMiss1znuo int32 = -1
	var found1znuo [maxFields1znuo]bool
	var curField1znuo string

doneWithStruct1znuo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1znuo > 0 || missingFieldsLeft1znuo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1znuo, missingFieldsLeft1znuo, msgp.ShowFound(found1znuo[:]), unmarshalMsgFieldOrder1znuo)
		if encodedFieldsLeft1znuo > 0 {
			encodedFieldsLeft1znuo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1znuo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1znuo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1znuo = 0
			}
			for nextMiss1znuo < maxFields1znuo && (found1znuo[nextMiss1znuo] || unmarshalMsgFieldSkip1znuo[nextMiss1znuo]) {
				nextMiss1znuo++
			}
			if nextMiss1znuo == maxFields1znuo {
				// filled all the empty fields!
				break doneWithStruct1znuo
			}
			missingFieldsLeft1znuo--
			curField1znuo = unmarshalMsgFieldOrder1znuo[nextMiss1znuo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1znuo)
		switch curField1znuo {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1znuo[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var znta uint32
				znta, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(znta) {
					z.Slice = (z.Slice)[:znta]
				} else {
					z.Slice = make([]S2, znta)
				}
				for zqhs := range z.Slice {
					bts, err = z.Slice[zqhs].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1znuo[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zioi uint32
				zioi, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zioi > 0 {
					z.Transform = make(map[int]*S2, zioi)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zioi > 0 {
					var zlzv int
					var zsia *S2
					zioi--
					zlzv, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zsia != nil {
							zsia.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zsia {
								zsia.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zsia == nil {
								zsia = new(S2)
							}
							bts, err = zsia.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zlzv] = zsia
				}
			}
		case "Myptr":
			found1znuo[2] = true
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
			found1znuo[3] = true
			var zpti uint32
			zpti, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zpti != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zpti}
				return
			}
			for zhza := range z.Myarray {
				z.Myarray[zhza], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1znuo[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zcwc uint32
				zcwc, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zcwc) {
					z.MySlice = (z.MySlice)[:zcwc]
				} else {
					z.MySlice = make([]string, zcwc)
				}
				for zdtv := range z.MySlice {
					z.MySlice[zdtv], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1znuo != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Big
var unmarshalMsgFieldOrder1znuo = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1znuo = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zqhs := range z.Slice {
		s += z.Slice[zqhs].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zlzv, zsia := range z.Transform {
			_ = zsia
			_ = zlzv
			s += msgp.IntSize
			if zsia == nil {
				s += msgp.NilSize
			} else {
				s += zsia.MSGPMsgsize()
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
	for zhza := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zhza])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zdtv := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zdtv])
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
	for zlbd, zyro := range z.R {
		o = msgp.AppendString(o, zlbd)
		o = msgp.AppendUint8(o, zyro)
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
	for zmsm := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zmsm])
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
	const maxFields2zymp = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zymp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zymp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2zymp := totalEncodedFields2zymp
	missingFieldsLeft2zymp := maxFields2zymp - totalEncodedFields2zymp

	var nextMiss2zymp int32 = -1
	var found2zymp [maxFields2zymp]bool
	var curField2zymp string

doneWithStruct2zymp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zymp > 0 || missingFieldsLeft2zymp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zymp, missingFieldsLeft2zymp, msgp.ShowFound(found2zymp[:]), unmarshalMsgFieldOrder2zymp)
		if encodedFieldsLeft2zymp > 0 {
			encodedFieldsLeft2zymp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2zymp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zymp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zymp = 0
			}
			for nextMiss2zymp < maxFields2zymp && (found2zymp[nextMiss2zymp] || unmarshalMsgFieldSkip2zymp[nextMiss2zymp]) {
				nextMiss2zymp++
			}
			if nextMiss2zymp == maxFields2zymp {
				// filled all the empty fields!
				break doneWithStruct2zymp
			}
			missingFieldsLeft2zymp--
			curField2zymp = unmarshalMsgFieldOrder2zymp[nextMiss2zymp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zymp)
		switch curField2zymp {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zymp[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2zymp[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zopz uint32
				zopz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && zopz > 0 {
					z.R = make(map[string]uint8, zopz)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zopz > 0 {
					var zlbd string
					var zyro uint8
					zopz--
					zlbd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zyro, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zlbd] = zyro
				}
			}
		case "P":
			found2zymp[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2zymp[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2zymp[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2zymp[6] = true
			var zooo uint32
			zooo, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zooo != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zooo}
				return
			}
			for zmsm := range z.Arr {
				z.Arr[zmsm], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2zymp[7] = true
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
	if nextMiss2zymp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of S2
var unmarshalMsgFieldOrder2zymp = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zymp = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zlbd, zyro := range z.R {
			_ = zyro
			_ = zlbd
			s += msgp.StringPrefixSize + len(zlbd) + msgp.Uint8Size
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
	const maxFields3ziqn = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3ziqn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3ziqn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3ziqn := totalEncodedFields3ziqn
	missingFieldsLeft3ziqn := maxFields3ziqn - totalEncodedFields3ziqn

	var nextMiss3ziqn int32 = -1
	var found3ziqn [maxFields3ziqn]bool
	var curField3ziqn string

doneWithStruct3ziqn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3ziqn > 0 || missingFieldsLeft3ziqn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3ziqn, missingFieldsLeft3ziqn, msgp.ShowFound(found3ziqn[:]), unmarshalMsgFieldOrder3ziqn)
		if encodedFieldsLeft3ziqn > 0 {
			encodedFieldsLeft3ziqn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3ziqn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3ziqn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3ziqn = 0
			}
			for nextMiss3ziqn < maxFields3ziqn && (found3ziqn[nextMiss3ziqn] || unmarshalMsgFieldSkip3ziqn[nextMiss3ziqn]) {
				nextMiss3ziqn++
			}
			if nextMiss3ziqn == maxFields3ziqn {
				// filled all the empty fields!
				break doneWithStruct3ziqn
			}
			missingFieldsLeft3ziqn--
			curField3ziqn = unmarshalMsgFieldOrder3ziqn[nextMiss3ziqn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3ziqn)
		switch curField3ziqn {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3ziqn[0] = true
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
	if nextMiss3ziqn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Sys
var unmarshalMsgFieldOrder3ziqn = []string{"F"}

var unmarshalMsgFieldSkip3ziqn = []bool{false}

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
	for zgiq := range z.Chld {
		o, err = z.Chld[zgiq].MSGPMarshalMsg(o)
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
	const maxFields4zbkv = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zbkv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zbkv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zbkv := totalEncodedFields4zbkv
	missingFieldsLeft4zbkv := maxFields4zbkv - totalEncodedFields4zbkv

	var nextMiss4zbkv int32 = -1
	var found4zbkv [maxFields4zbkv]bool
	var curField4zbkv string

doneWithStruct4zbkv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zbkv > 0 || missingFieldsLeft4zbkv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zbkv, missingFieldsLeft4zbkv, msgp.ShowFound(found4zbkv[:]), unmarshalMsgFieldOrder4zbkv)
		if encodedFieldsLeft4zbkv > 0 {
			encodedFieldsLeft4zbkv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zbkv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zbkv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zbkv = 0
			}
			for nextMiss4zbkv < maxFields4zbkv && (found4zbkv[nextMiss4zbkv] || unmarshalMsgFieldSkip4zbkv[nextMiss4zbkv]) {
				nextMiss4zbkv++
			}
			if nextMiss4zbkv == maxFields4zbkv {
				// filled all the empty fields!
				break doneWithStruct4zbkv
			}
			missingFieldsLeft4zbkv--
			curField4zbkv = unmarshalMsgFieldOrder4zbkv[nextMiss4zbkv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zbkv)
		switch curField4zbkv {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zbkv[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zzla uint32
				zzla, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zzla) {
					z.Chld = (z.Chld)[:zzla]
				} else {
					z.Chld = make([]Tree, zzla)
				}
				for zgiq := range z.Chld {
					bts, err = z.Chld[zgiq].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zbkv[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zbkv[2] = true
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
	if nextMiss4zbkv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zbkv = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zbkv = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zgiq := range z.Chld {
		s += z.Chld[zgiq].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
