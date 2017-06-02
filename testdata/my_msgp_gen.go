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
	const maxFields0znpa = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0znpa uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0znpa, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0znpa := totalEncodedFields0znpa
	missingFieldsLeft0znpa := maxFields0znpa - totalEncodedFields0znpa

	var nextMiss0znpa int32 = -1
	var found0znpa [maxFields0znpa]bool
	var curField0znpa string

doneWithStruct0znpa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0znpa > 0 || missingFieldsLeft0znpa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0znpa, missingFieldsLeft0znpa, msgp.ShowFound(found0znpa[:]), unmarshalMsgFieldOrder0znpa)
		if encodedFieldsLeft0znpa > 0 {
			encodedFieldsLeft0znpa--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0znpa = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0znpa < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0znpa = 0
			}
			for nextMiss0znpa < maxFields0znpa && (found0znpa[nextMiss0znpa] || unmarshalMsgFieldSkip0znpa[nextMiss0znpa]) {
				nextMiss0znpa++
			}
			if nextMiss0znpa == maxFields0znpa {
				// filled all the empty fields!
				break doneWithStruct0znpa
			}
			missingFieldsLeft0znpa--
			curField0znpa = unmarshalMsgFieldOrder0znpa[nextMiss0znpa]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0znpa)
		switch curField0znpa {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0znpa[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday":
			found0znpa[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone":
			found0znpa[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs":
			found0znpa[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA":
			found0znpa[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend":
			found0znpa[5] = true
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
	if nextMiss0znpa != -1 {
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
var unmarshalMsgFieldOrder0znpa = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0znpa = []bool{false, false, false, false, false, false}

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
	for zenh := range z.Slice {
		o, err = z.Slice[zenh].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zbec, zhxf := range z.Transform {
		o = msgp.AppendInt(o, zbec)
		if zhxf == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zhxf.MSGPMarshalMsg(o)
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
	for zrkq := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zrkq])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zcui := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zcui])
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
	const maxFields1zgme = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zgme uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zgme, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zgme := totalEncodedFields1zgme
	missingFieldsLeft1zgme := maxFields1zgme - totalEncodedFields1zgme

	var nextMiss1zgme int32 = -1
	var found1zgme [maxFields1zgme]bool
	var curField1zgme string

doneWithStruct1zgme:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgme > 0 || missingFieldsLeft1zgme > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgme, missingFieldsLeft1zgme, msgp.ShowFound(found1zgme[:]), unmarshalMsgFieldOrder1zgme)
		if encodedFieldsLeft1zgme > 0 {
			encodedFieldsLeft1zgme--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zgme = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zgme < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zgme = 0
			}
			for nextMiss1zgme < maxFields1zgme && (found1zgme[nextMiss1zgme] || unmarshalMsgFieldSkip1zgme[nextMiss1zgme]) {
				nextMiss1zgme++
			}
			if nextMiss1zgme == maxFields1zgme {
				// filled all the empty fields!
				break doneWithStruct1zgme
			}
			missingFieldsLeft1zgme--
			curField1zgme = unmarshalMsgFieldOrder1zgme[nextMiss1zgme]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgme)
		switch curField1zgme {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zgme[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zygg uint32
				zygg, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zygg) {
					z.Slice = (z.Slice)[:zygg]
				} else {
					z.Slice = make([]S2, zygg)
				}
				for zenh := range z.Slice {
					bts, err = z.Slice[zenh].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform":
			found1zgme[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zbko uint32
				zbko, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zbko > 0 {
					z.Transform = make(map[int]*S2, zbko)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zbko > 0 {
					var zbec int
					var zhxf *S2
					zbko--
					zbec, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zhxf != nil {
							zhxf.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zhxf {
								zhxf.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zhxf == nil {
								zhxf = new(S2)
							}
							bts, err = zhxf.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zbec] = zhxf
				}
			}
		case "Myptr":
			found1zgme[2] = true
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
			found1zgme[3] = true
			var zwcn uint32
			zwcn, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zwcn != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zwcn}
				return
			}
			for zrkq := range z.Myarray {
				z.Myarray[zrkq], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice":
			found1zgme[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zfns uint32
				zfns, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zfns) {
					z.MySlice = (z.MySlice)[:zfns]
				} else {
					z.MySlice = make([]string, zfns)
				}
				for zcui := range z.MySlice {
					z.MySlice[zcui], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zgme != -1 {
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
var unmarshalMsgFieldOrder1zgme = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zgme = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zenh := range z.Slice {
		s += z.Slice[zenh].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zbec, zhxf := range z.Transform {
			_ = zhxf
			_ = zbec
			s += msgp.IntSize
			if zhxf == nil {
				s += msgp.NilSize
			} else {
				s += zhxf.MSGPMsgsize()
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
	for zrkq := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zrkq])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zcui := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zcui])
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
	for zhsq, zota := range z.R {
		o = msgp.AppendString(o, zhsq)
		o = msgp.AppendUint8(o, zota)
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
	for zcpf := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zcpf])
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
	const maxFields2zzaf = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zzaf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zzaf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zzaf := totalEncodedFields2zzaf
	missingFieldsLeft2zzaf := maxFields2zzaf - totalEncodedFields2zzaf

	var nextMiss2zzaf int32 = -1
	var found2zzaf [maxFields2zzaf]bool
	var curField2zzaf string

doneWithStruct2zzaf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zzaf > 0 || missingFieldsLeft2zzaf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zzaf, missingFieldsLeft2zzaf, msgp.ShowFound(found2zzaf[:]), unmarshalMsgFieldOrder2zzaf)
		if encodedFieldsLeft2zzaf > 0 {
			encodedFieldsLeft2zzaf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zzaf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zzaf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zzaf = 0
			}
			for nextMiss2zzaf < maxFields2zzaf && (found2zzaf[nextMiss2zzaf] || unmarshalMsgFieldSkip2zzaf[nextMiss2zzaf]) {
				nextMiss2zzaf++
			}
			if nextMiss2zzaf == maxFields2zzaf {
				// filled all the empty fields!
				break doneWithStruct2zzaf
			}
			missingFieldsLeft2zzaf--
			curField2zzaf = unmarshalMsgFieldOrder2zzaf[nextMiss2zzaf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zzaf)
		switch curField2zzaf {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zzaf[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph":
			found2zzaf[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var ztua uint32
				ztua, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && ztua > 0 {
					z.R = make(map[string]uint8, ztua)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for ztua > 0 {
					var zhsq string
					var zota uint8
					ztua--
					zhsq, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zota, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zhsq] = zota
				}
			}
		case "P":
			found2zzaf[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q":
			found2zzaf[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T":
			found2zzaf[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr":
			found2zzaf[6] = true
			var zitt uint32
			zitt, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zitt != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zitt}
				return
			}
			for zcpf := range z.Arr {
				z.Arr[zcpf], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree":
			found2zzaf[7] = true
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
	if nextMiss2zzaf != -1 {
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
var unmarshalMsgFieldOrder2zzaf = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zzaf = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zhsq, zota := range z.R {
			_ = zota
			_ = zhsq
			s += msgp.StringPrefixSize + len(zhsq) + msgp.Uint8Size
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
	const maxFields3zaqt = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zaqt uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zaqt, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zaqt := totalEncodedFields3zaqt
	missingFieldsLeft3zaqt := maxFields3zaqt - totalEncodedFields3zaqt

	var nextMiss3zaqt int32 = -1
	var found3zaqt [maxFields3zaqt]bool
	var curField3zaqt string

doneWithStruct3zaqt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zaqt > 0 || missingFieldsLeft3zaqt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zaqt, missingFieldsLeft3zaqt, msgp.ShowFound(found3zaqt[:]), unmarshalMsgFieldOrder3zaqt)
		if encodedFieldsLeft3zaqt > 0 {
			encodedFieldsLeft3zaqt--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zaqt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zaqt < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zaqt = 0
			}
			for nextMiss3zaqt < maxFields3zaqt && (found3zaqt[nextMiss3zaqt] || unmarshalMsgFieldSkip3zaqt[nextMiss3zaqt]) {
				nextMiss3zaqt++
			}
			if nextMiss3zaqt == maxFields3zaqt {
				// filled all the empty fields!
				break doneWithStruct3zaqt
			}
			missingFieldsLeft3zaqt--
			curField3zaqt = unmarshalMsgFieldOrder3zaqt[nextMiss3zaqt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zaqt)
		switch curField3zaqt {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zaqt[0] = true
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
	if nextMiss3zaqt != -1 {
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
var unmarshalMsgFieldOrder3zaqt = []string{"F"}

var unmarshalMsgFieldSkip3zaqt = []bool{false}

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
	for zrfx := range z.Chld {
		o, err = z.Chld[zrfx].MSGPMarshalMsg(o)
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
	const maxFields4zmlp = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zmlp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zmlp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zmlp := totalEncodedFields4zmlp
	missingFieldsLeft4zmlp := maxFields4zmlp - totalEncodedFields4zmlp

	var nextMiss4zmlp int32 = -1
	var found4zmlp [maxFields4zmlp]bool
	var curField4zmlp string

doneWithStruct4zmlp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zmlp > 0 || missingFieldsLeft4zmlp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zmlp, missingFieldsLeft4zmlp, msgp.ShowFound(found4zmlp[:]), unmarshalMsgFieldOrder4zmlp)
		if encodedFieldsLeft4zmlp > 0 {
			encodedFieldsLeft4zmlp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zmlp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zmlp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zmlp = 0
			}
			for nextMiss4zmlp < maxFields4zmlp && (found4zmlp[nextMiss4zmlp] || unmarshalMsgFieldSkip4zmlp[nextMiss4zmlp]) {
				nextMiss4zmlp++
			}
			if nextMiss4zmlp == maxFields4zmlp {
				// filled all the empty fields!
				break doneWithStruct4zmlp
			}
			missingFieldsLeft4zmlp--
			curField4zmlp = unmarshalMsgFieldOrder4zmlp[nextMiss4zmlp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zmlp)
		switch curField4zmlp {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zmlp[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zojv uint32
				zojv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zojv) {
					z.Chld = (z.Chld)[:zojv]
				} else {
					z.Chld = make([]Tree, zojv)
				}
				for zrfx := range z.Chld {
					bts, err = z.Chld[zrfx].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str":
			found4zmlp[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par":
			found4zmlp[2] = true
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
	if nextMiss4zmlp != -1 {
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
var unmarshalMsgFieldOrder4zmlp = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zmlp = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zrfx := range z.Chld {
		s += z.Chld[zrfx].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
