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
	const maxFields0zeub = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zeub uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zeub, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zeub := totalEncodedFields0zeub
	missingFieldsLeft0zeub := maxFields0zeub - totalEncodedFields0zeub

	var nextMiss0zeub int32 = -1
	var found0zeub [maxFields0zeub]bool
	var curField0zeub string

doneWithStruct0zeub:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zeub > 0 || missingFieldsLeft0zeub > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zeub, missingFieldsLeft0zeub, msgp.ShowFound(found0zeub[:]), unmarshalMsgFieldOrder0zeub)
		if encodedFieldsLeft0zeub > 0 {
			encodedFieldsLeft0zeub--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zeub = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zeub < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zeub = 0
			}
			for nextMiss0zeub < maxFields0zeub && (found0zeub[nextMiss0zeub] || unmarshalMsgFieldSkip0zeub[nextMiss0zeub]) {
				nextMiss0zeub++
			}
			if nextMiss0zeub == maxFields0zeub {
				// filled all the empty fields!
				break doneWithStruct0zeub
			}
			missingFieldsLeft0zeub--
			curField0zeub = unmarshalMsgFieldOrder0zeub[nextMiss0zeub]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zeub)
		switch curField0zeub {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zeub[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday":
			found0zeub[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone":
			found0zeub[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs":
			found0zeub[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA":
			found0zeub[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend":
			found0zeub[5] = true
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
	if nextMiss0zeub != -1 {
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
var unmarshalMsgFieldOrder0zeub = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zeub = []bool{false, false, false, false, false, false}

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
	for zglx := range z.Slice {
		o, err = z.Slice[zglx].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zgld, zoqb := range z.Transform {
		o = msgp.AppendInt(o, zgld)
		if zoqb == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zoqb.MSGPMarshalMsg(o)
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
	for zimc := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zimc])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zdvw := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zdvw])
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
	const maxFields1zedr = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zedr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zedr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zedr := totalEncodedFields1zedr
	missingFieldsLeft1zedr := maxFields1zedr - totalEncodedFields1zedr

	var nextMiss1zedr int32 = -1
	var found1zedr [maxFields1zedr]bool
	var curField1zedr string

doneWithStruct1zedr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zedr > 0 || missingFieldsLeft1zedr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zedr, missingFieldsLeft1zedr, msgp.ShowFound(found1zedr[:]), unmarshalMsgFieldOrder1zedr)
		if encodedFieldsLeft1zedr > 0 {
			encodedFieldsLeft1zedr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zedr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zedr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zedr = 0
			}
			for nextMiss1zedr < maxFields1zedr && (found1zedr[nextMiss1zedr] || unmarshalMsgFieldSkip1zedr[nextMiss1zedr]) {
				nextMiss1zedr++
			}
			if nextMiss1zedr == maxFields1zedr {
				// filled all the empty fields!
				break doneWithStruct1zedr
			}
			missingFieldsLeft1zedr--
			curField1zedr = unmarshalMsgFieldOrder1zedr[nextMiss1zedr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zedr)
		switch curField1zedr {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zedr[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zgto uint32
				zgto, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zgto) {
					z.Slice = (z.Slice)[:zgto]
				} else {
					z.Slice = make([]S2, zgto)
				}
				for zglx := range z.Slice {
					bts, err = z.Slice[zglx].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform":
			found1zedr[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zhkz uint32
				zhkz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zhkz > 0 {
					z.Transform = make(map[int]*S2, zhkz)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zhkz > 0 {
					var zgld int
					var zoqb *S2
					zhkz--
					zgld, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zoqb != nil {
							zoqb.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zoqb {
								zoqb.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zoqb == nil {
								zoqb = new(S2)
							}
							bts, err = zoqb.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zgld] = zoqb
				}
			}
		case "Myptr":
			found1zedr[2] = true
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
			found1zedr[3] = true
			var zump uint32
			zump, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zump != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zump}
				return
			}
			for zimc := range z.Myarray {
				z.Myarray[zimc], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice":
			found1zedr[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var znbh uint32
				znbh, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(znbh) {
					z.MySlice = (z.MySlice)[:znbh]
				} else {
					z.MySlice = make([]string, znbh)
				}
				for zdvw := range z.MySlice {
					z.MySlice[zdvw], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zedr != -1 {
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
var unmarshalMsgFieldOrder1zedr = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zedr = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zglx := range z.Slice {
		s += z.Slice[zglx].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zgld, zoqb := range z.Transform {
			_ = zoqb
			_ = zgld
			s += msgp.IntSize
			if zoqb == nil {
				s += msgp.NilSize
			} else {
				s += zoqb.MSGPMsgsize()
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
	for zimc := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zimc])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zdvw := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zdvw])
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
	for zusn, zmcq := range z.R {
		o = msgp.AppendString(o, zusn)
		o = msgp.AppendUint8(o, zmcq)
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
	for zmci := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zmci])
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
	const maxFields2zamu = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zamu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zamu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zamu := totalEncodedFields2zamu
	missingFieldsLeft2zamu := maxFields2zamu - totalEncodedFields2zamu

	var nextMiss2zamu int32 = -1
	var found2zamu [maxFields2zamu]bool
	var curField2zamu string

doneWithStruct2zamu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zamu > 0 || missingFieldsLeft2zamu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zamu, missingFieldsLeft2zamu, msgp.ShowFound(found2zamu[:]), unmarshalMsgFieldOrder2zamu)
		if encodedFieldsLeft2zamu > 0 {
			encodedFieldsLeft2zamu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zamu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zamu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zamu = 0
			}
			for nextMiss2zamu < maxFields2zamu && (found2zamu[nextMiss2zamu] || unmarshalMsgFieldSkip2zamu[nextMiss2zamu]) {
				nextMiss2zamu++
			}
			if nextMiss2zamu == maxFields2zamu {
				// filled all the empty fields!
				break doneWithStruct2zamu
			}
			missingFieldsLeft2zamu--
			curField2zamu = unmarshalMsgFieldOrder2zamu[nextMiss2zamu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zamu)
		switch curField2zamu {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zamu[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph":
			found2zamu[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zpad uint32
				zpad, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zpad > 0 {
					z.R = make(map[string]uint8, zpad)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zpad > 0 {
					var zusn string
					var zmcq uint8
					zpad--
					zusn, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zmcq, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zusn] = zmcq
				}
			}
		case "P":
			found2zamu[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q":
			found2zamu[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T":
			found2zamu[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr":
			found2zamu[6] = true
			var zjcx uint32
			zjcx, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zjcx != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zjcx}
				return
			}
			for zmci := range z.Arr {
				z.Arr[zmci], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree":
			found2zamu[7] = true
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
	if nextMiss2zamu != -1 {
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
var unmarshalMsgFieldOrder2zamu = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zamu = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zusn, zmcq := range z.R {
			_ = zmcq
			_ = zusn
			s += msgp.StringPrefixSize + len(zusn) + msgp.Uint8Size
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
	const maxFields3zrky = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zrky uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zrky, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zrky := totalEncodedFields3zrky
	missingFieldsLeft3zrky := maxFields3zrky - totalEncodedFields3zrky

	var nextMiss3zrky int32 = -1
	var found3zrky [maxFields3zrky]bool
	var curField3zrky string

doneWithStruct3zrky:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zrky > 0 || missingFieldsLeft3zrky > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zrky, missingFieldsLeft3zrky, msgp.ShowFound(found3zrky[:]), unmarshalMsgFieldOrder3zrky)
		if encodedFieldsLeft3zrky > 0 {
			encodedFieldsLeft3zrky--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zrky = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zrky < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zrky = 0
			}
			for nextMiss3zrky < maxFields3zrky && (found3zrky[nextMiss3zrky] || unmarshalMsgFieldSkip3zrky[nextMiss3zrky]) {
				nextMiss3zrky++
			}
			if nextMiss3zrky == maxFields3zrky {
				// filled all the empty fields!
				break doneWithStruct3zrky
			}
			missingFieldsLeft3zrky--
			curField3zrky = unmarshalMsgFieldOrder3zrky[nextMiss3zrky]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zrky)
		switch curField3zrky {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zrky[0] = true
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
	if nextMiss3zrky != -1 {
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
var unmarshalMsgFieldOrder3zrky = []string{"F"}

var unmarshalMsgFieldSkip3zrky = []bool{false}

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
	for zzpc := range z.Chld {
		o, err = z.Chld[zzpc].MSGPMarshalMsg(o)
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
	const maxFields4zrpc = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zrpc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zrpc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zrpc := totalEncodedFields4zrpc
	missingFieldsLeft4zrpc := maxFields4zrpc - totalEncodedFields4zrpc

	var nextMiss4zrpc int32 = -1
	var found4zrpc [maxFields4zrpc]bool
	var curField4zrpc string

doneWithStruct4zrpc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zrpc > 0 || missingFieldsLeft4zrpc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zrpc, missingFieldsLeft4zrpc, msgp.ShowFound(found4zrpc[:]), unmarshalMsgFieldOrder4zrpc)
		if encodedFieldsLeft4zrpc > 0 {
			encodedFieldsLeft4zrpc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zrpc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zrpc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zrpc = 0
			}
			for nextMiss4zrpc < maxFields4zrpc && (found4zrpc[nextMiss4zrpc] || unmarshalMsgFieldSkip4zrpc[nextMiss4zrpc]) {
				nextMiss4zrpc++
			}
			if nextMiss4zrpc == maxFields4zrpc {
				// filled all the empty fields!
				break doneWithStruct4zrpc
			}
			missingFieldsLeft4zrpc--
			curField4zrpc = unmarshalMsgFieldOrder4zrpc[nextMiss4zrpc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zrpc)
		switch curField4zrpc {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zrpc[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zrzj uint32
				zrzj, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zrzj) {
					z.Chld = (z.Chld)[:zrzj]
				} else {
					z.Chld = make([]Tree, zrzj)
				}
				for zzpc := range z.Chld {
					bts, err = z.Chld[zzpc].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str":
			found4zrpc[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par":
			found4zrpc[2] = true
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
	if nextMiss4zrpc != -1 {
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
var unmarshalMsgFieldOrder4zrpc = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zrpc = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zzpc := range z.Chld {
		s += z.Chld[zzpc].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
