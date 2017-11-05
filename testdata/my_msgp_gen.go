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
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

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
	const maxFields0zlbu = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zlbu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zlbu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zlbu := totalEncodedFields0zlbu
	missingFieldsLeft0zlbu := maxFields0zlbu - totalEncodedFields0zlbu

	var nextMiss0zlbu int32 = -1
	var found0zlbu [maxFields0zlbu]bool
	var curField0zlbu string

doneWithStruct0zlbu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zlbu > 0 || missingFieldsLeft0zlbu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zlbu, missingFieldsLeft0zlbu, msgp.ShowFound(found0zlbu[:]), unmarshalMsgFieldOrder0zlbu)
		if encodedFieldsLeft0zlbu > 0 {
			encodedFieldsLeft0zlbu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zlbu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zlbu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zlbu = 0
			}
			for nextMiss0zlbu < maxFields0zlbu && (found0zlbu[nextMiss0zlbu] || unmarshalMsgFieldSkip0zlbu[nextMiss0zlbu]) {
				nextMiss0zlbu++
			}
			if nextMiss0zlbu == maxFields0zlbu {
				// filled all the empty fields!
				break doneWithStruct0zlbu
			}
			missingFieldsLeft0zlbu--
			curField0zlbu = unmarshalMsgFieldOrder0zlbu[nextMiss0zlbu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zlbu)
		switch curField0zlbu {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zlbu[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday":
			found0zlbu[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone":
			found0zlbu[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs":
			found0zlbu[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA":
			found0zlbu[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend":
			found0zlbu[5] = true
			z.Friend, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zlbu != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of A
var unmarshalMsgFieldOrder0zlbu = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zlbu = []bool{false, false, false, false, false, false}

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
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 5
	// string "Slice"
	o = append(o, 0x85, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
	for zkkg := range z.Slice {
		o, err = z.Slice[zkkg].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zedk, zprr := range z.Transform {
		o = msgp.AppendInt(o, zedk)
		if zprr == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zprr.MSGPMarshalMsg(o)
			if err != nil {
				return
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
			return
		}
	}
	// string "Myarray"
	o = append(o, 0xa7, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, 3)
	for zoeb := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zoeb])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zshf := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zshf])
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
	const maxFields1zkum = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zkum uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zkum, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zkum := totalEncodedFields1zkum
	missingFieldsLeft1zkum := maxFields1zkum - totalEncodedFields1zkum

	var nextMiss1zkum int32 = -1
	var found1zkum [maxFields1zkum]bool
	var curField1zkum string

doneWithStruct1zkum:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zkum > 0 || missingFieldsLeft1zkum > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zkum, missingFieldsLeft1zkum, msgp.ShowFound(found1zkum[:]), unmarshalMsgFieldOrder1zkum)
		if encodedFieldsLeft1zkum > 0 {
			encodedFieldsLeft1zkum--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zkum = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zkum < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zkum = 0
			}
			for nextMiss1zkum < maxFields1zkum && (found1zkum[nextMiss1zkum] || unmarshalMsgFieldSkip1zkum[nextMiss1zkum]) {
				nextMiss1zkum++
			}
			if nextMiss1zkum == maxFields1zkum {
				// filled all the empty fields!
				break doneWithStruct1zkum
			}
			missingFieldsLeft1zkum--
			curField1zkum = unmarshalMsgFieldOrder1zkum[nextMiss1zkum]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zkum)
		switch curField1zkum {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zkum[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zujt uint32
				zujt, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zujt) {
					z.Slice = (z.Slice)[:zujt]
				} else {
					z.Slice = make([]S2, zujt)
				}
				for zkkg := range z.Slice {
					bts, err = z.Slice[zkkg].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform":
			found1zkum[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zdss uint32
				zdss, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zdss > 0 {
					z.Transform = make(map[int]*S2, zdss)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zdss > 0 {
					var zedk int
					var zprr *S2
					zdss--
					zedk, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zprr != nil {
							zprr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zprr {
								zprr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zprr == nil {
								zprr = new(S2)
							}
							bts, err = zprr.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zedk] = zprr
				}
			}
		case "Myptr":
			found1zkum[2] = true
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
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Myarray":
			found1zkum[3] = true
			var zrgr uint32
			zrgr, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zrgr != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zrgr}
				return
			}
			for zoeb := range z.Myarray {
				z.Myarray[zoeb], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice":
			found1zkum[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zhuo uint32
				zhuo, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zhuo) {
					z.MySlice = (z.MySlice)[:zhuo]
				} else {
					z.MySlice = make([]string, zhuo)
				}
				for zshf := range z.MySlice {
					z.MySlice[zshf], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss1zkum != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Big
var unmarshalMsgFieldOrder1zkum = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zkum = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zkkg := range z.Slice {
		s += z.Slice[zkkg].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zedk, zprr := range z.Transform {
			_ = zprr
			_ = zedk
			s += msgp.IntSize
			if zprr == nil {
				s += msgp.NilSize
			} else {
				s += zprr.MSGPMsgsize()
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
	for zoeb := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zoeb])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zshf := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zshf])
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *S2) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 7
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *S2) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 7
	// string "beta"
	o = append(o, 0x87, 0xa4, 0x62, 0x65, 0x74, 0x61)
	o = msgp.AppendString(o, z.B)
	// string "ralph"
	o = append(o, 0xa5, 0x72, 0x61, 0x6c, 0x70, 0x68)
	o = msgp.AppendMapHeader(o, uint32(len(z.R)))
	for zaug, zevt := range z.R {
		o = msgp.AppendString(o, zaug)
		o = msgp.AppendUint8(o, zevt)
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
	for zbsl := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zbsl])
	}
	// string "MyTree"
	o = append(o, 0xa6, 0x4d, 0x79, 0x54, 0x72, 0x65, 0x65)
	if z.MyTree == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.MyTree.MSGPMarshalMsg(o)
		if err != nil {
			return
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
	const maxFields2zrbi = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zrbi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zrbi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zrbi := totalEncodedFields2zrbi
	missingFieldsLeft2zrbi := maxFields2zrbi - totalEncodedFields2zrbi

	var nextMiss2zrbi int32 = -1
	var found2zrbi [maxFields2zrbi]bool
	var curField2zrbi string

doneWithStruct2zrbi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zrbi > 0 || missingFieldsLeft2zrbi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zrbi, missingFieldsLeft2zrbi, msgp.ShowFound(found2zrbi[:]), unmarshalMsgFieldOrder2zrbi)
		if encodedFieldsLeft2zrbi > 0 {
			encodedFieldsLeft2zrbi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zrbi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zrbi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zrbi = 0
			}
			for nextMiss2zrbi < maxFields2zrbi && (found2zrbi[nextMiss2zrbi] || unmarshalMsgFieldSkip2zrbi[nextMiss2zrbi]) {
				nextMiss2zrbi++
			}
			if nextMiss2zrbi == maxFields2zrbi {
				// filled all the empty fields!
				break doneWithStruct2zrbi
			}
			missingFieldsLeft2zrbi--
			curField2zrbi = unmarshalMsgFieldOrder2zrbi[nextMiss2zrbi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zrbi)
		switch curField2zrbi {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zrbi[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph":
			found2zrbi[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zkcr uint32
				zkcr, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zkcr > 0 {
					z.R = make(map[string]uint8, zkcr)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zkcr > 0 {
					var zaug string
					var zevt uint8
					zkcr--
					zaug, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zevt, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zaug] = zevt
				}
			}
		case "P":
			found2zrbi[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q":
			found2zrbi[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T":
			found2zrbi[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr":
			found2zrbi[6] = true
			var zlur uint32
			zlur, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zlur != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zlur}
				return
			}
			for zbsl := range z.Arr {
				z.Arr[zbsl], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree":
			found2zrbi[7] = true
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
						return
					}
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss2zrbi != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of S2
var unmarshalMsgFieldOrder2zrbi = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zrbi = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zaug, zevt := range z.R {
			_ = zevt
			_ = zaug
			s += msgp.StringPrefixSize + len(zaug) + msgp.Uint8Size
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
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 1
	// string "F"
	o = append(o, 0x81, 0xa1, 0x46)
	o, err = msgp.AppendIntf(o, z.F)
	if err != nil {
		return
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
	const maxFields3zyxs = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zyxs uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zyxs, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zyxs := totalEncodedFields3zyxs
	missingFieldsLeft3zyxs := maxFields3zyxs - totalEncodedFields3zyxs

	var nextMiss3zyxs int32 = -1
	var found3zyxs [maxFields3zyxs]bool
	var curField3zyxs string

doneWithStruct3zyxs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zyxs > 0 || missingFieldsLeft3zyxs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zyxs, missingFieldsLeft3zyxs, msgp.ShowFound(found3zyxs[:]), unmarshalMsgFieldOrder3zyxs)
		if encodedFieldsLeft3zyxs > 0 {
			encodedFieldsLeft3zyxs--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zyxs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zyxs < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zyxs = 0
			}
			for nextMiss3zyxs < maxFields3zyxs && (found3zyxs[nextMiss3zyxs] || unmarshalMsgFieldSkip3zyxs[nextMiss3zyxs]) {
				nextMiss3zyxs++
			}
			if nextMiss3zyxs == maxFields3zyxs {
				// filled all the empty fields!
				break doneWithStruct3zyxs
			}
			missingFieldsLeft3zyxs--
			curField3zyxs = unmarshalMsgFieldOrder3zyxs[nextMiss3zyxs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zyxs)
		switch curField3zyxs {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zyxs[0] = true
			z.F, bts, err = nbs.ReadIntfBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss3zyxs != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Sys
var unmarshalMsgFieldOrder3zyxs = []string{"F"}

var unmarshalMsgFieldSkip3zyxs = []bool{false}

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
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 3
	// string "Chld"
	o = append(o, 0x83, 0xa4, 0x43, 0x68, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
	for zusd := range z.Chld {
		o, err = z.Chld[zusd].MSGPMarshalMsg(o)
		if err != nil {
			return
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
			return
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
	const maxFields4zids = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zids uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zids, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zids := totalEncodedFields4zids
	missingFieldsLeft4zids := maxFields4zids - totalEncodedFields4zids

	var nextMiss4zids int32 = -1
	var found4zids [maxFields4zids]bool
	var curField4zids string

doneWithStruct4zids:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zids > 0 || missingFieldsLeft4zids > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zids, missingFieldsLeft4zids, msgp.ShowFound(found4zids[:]), unmarshalMsgFieldOrder4zids)
		if encodedFieldsLeft4zids > 0 {
			encodedFieldsLeft4zids--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zids = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zids < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zids = 0
			}
			for nextMiss4zids < maxFields4zids && (found4zids[nextMiss4zids] || unmarshalMsgFieldSkip4zids[nextMiss4zids]) {
				nextMiss4zids++
			}
			if nextMiss4zids == maxFields4zids {
				// filled all the empty fields!
				break doneWithStruct4zids
			}
			missingFieldsLeft4zids--
			curField4zids = unmarshalMsgFieldOrder4zids[nextMiss4zids]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zids)
		switch curField4zids {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zids[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zmrv uint32
				zmrv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zmrv) {
					z.Chld = (z.Chld)[:zmrv]
				} else {
					z.Chld = make([]Tree, zmrv)
				}
				for zusd := range z.Chld {
					bts, err = z.Chld[zusd].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str":
			found4zids[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par":
			found4zids[2] = true
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
						return
					}
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss4zids != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Tree
var unmarshalMsgFieldOrder4zids = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zids = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zusd := range z.Chld {
		s += z.Chld[zusd].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
