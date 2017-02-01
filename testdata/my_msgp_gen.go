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
	const maxFields0zjzn = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zjzn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zjzn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft0zjzn := totalEncodedFields0zjzn
	missingFieldsLeft0zjzn := maxFields0zjzn - totalEncodedFields0zjzn

	var nextMiss0zjzn int32 = -1
	var found0zjzn [maxFields0zjzn]bool
	var curField0zjzn string

doneWithStruct0zjzn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zjzn > 0 || missingFieldsLeft0zjzn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zjzn, missingFieldsLeft0zjzn, msgp.ShowFound(found0zjzn[:]), unmarshalMsgFieldOrder0zjzn)
		if encodedFieldsLeft0zjzn > 0 {
			encodedFieldsLeft0zjzn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField0zjzn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zjzn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zjzn = 0
			}
			for nextMiss0zjzn < maxFields0zjzn && (found0zjzn[nextMiss0zjzn] || unmarshalMsgFieldSkip0zjzn[nextMiss0zjzn]) {
				nextMiss0zjzn++
			}
			if nextMiss0zjzn == maxFields0zjzn {
				// filled all the empty fields!
				break doneWithStruct0zjzn
			}
			missingFieldsLeft0zjzn--
			curField0zjzn = unmarshalMsgFieldOrder0zjzn[nextMiss0zjzn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zjzn)
		switch curField0zjzn {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zjzn[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Bday":
			found0zjzn[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				panic(err)
			}
		case "phone":
			found0zjzn[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Sibs":
			found0zjzn[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case "GPA":
			found0zjzn[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Friend":
			found0zjzn[5] = true
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
	if nextMiss0zjzn != -1 {
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
var unmarshalMsgFieldOrder0zjzn = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zjzn = []bool{false, false, false, false, false, false}

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
	for zezc := range z.Slice {
		o, err = z.Slice[zezc].MSGPMarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zzgo, zewj := range z.Transform {
		o = msgp.AppendInt(o, zzgo)
		if zewj == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zewj.MSGPMarshalMsg(o)
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
	for zkqh := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zkqh])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zjwv := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zjwv])
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
	const maxFields1zwji = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zwji uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zwji, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zwji := totalEncodedFields1zwji
	missingFieldsLeft1zwji := maxFields1zwji - totalEncodedFields1zwji

	var nextMiss1zwji int32 = -1
	var found1zwji [maxFields1zwji]bool
	var curField1zwji string

doneWithStruct1zwji:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zwji > 0 || missingFieldsLeft1zwji > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zwji, missingFieldsLeft1zwji, msgp.ShowFound(found1zwji[:]), unmarshalMsgFieldOrder1zwji)
		if encodedFieldsLeft1zwji > 0 {
			encodedFieldsLeft1zwji--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zwji = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zwji < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zwji = 0
			}
			for nextMiss1zwji < maxFields1zwji && (found1zwji[nextMiss1zwji] || unmarshalMsgFieldSkip1zwji[nextMiss1zwji]) {
				nextMiss1zwji++
			}
			if nextMiss1zwji == maxFields1zwji {
				// filled all the empty fields!
				break doneWithStruct1zwji
			}
			missingFieldsLeft1zwji--
			curField1zwji = unmarshalMsgFieldOrder1zwji[nextMiss1zwji]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zwji)
		switch curField1zwji {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zwji[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var znnn uint32
				znnn, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Slice) >= int(znnn) {
					z.Slice = (z.Slice)[:znnn]
				} else {
					z.Slice = make([]S2, znnn)
				}
				for zezc := range z.Slice {
					bts, err = z.Slice[zezc].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Transform":
			found1zwji[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zuie uint32
				zuie, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Transform == nil && zuie > 0 {
					z.Transform = make(map[int]*S2, zuie)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zuie > 0 {
					var zzgo int
					var zewj *S2
					zuie--
					zzgo, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zewj != nil {
							zewj.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zewj {
								zewj.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zewj == nil {
								zewj = new(S2)
							}
							bts, err = zewj.MSGPUnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.Transform[zzgo] = zewj
				}
			}
		case "Myptr":
			found1zwji[2] = true
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
			found1zwji[3] = true
			var zarg uint32
			zarg, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zarg != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zarg}
				return
			}
			for zkqh := range z.Myarray {
				z.Myarray[zkqh], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MySlice":
			found1zwji[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zgke uint32
				zgke, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.MySlice) >= int(zgke) {
					z.MySlice = (z.MySlice)[:zgke]
				} else {
					z.MySlice = make([]string, zgke)
				}
				for zjwv := range z.MySlice {
					z.MySlice[zjwv], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zwji != -1 {
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
var unmarshalMsgFieldOrder1zwji = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zwji = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zezc := range z.Slice {
		s += z.Slice[zezc].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zzgo, zewj := range z.Transform {
			_ = zewj
			_ = zzgo
			s += msgp.IntSize
			if zewj == nil {
				s += msgp.NilSize
			} else {
				s += zewj.MSGPMsgsize()
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
	for zkqh := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zkqh])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zjwv := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zjwv])
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
	for zoqn, zsab := range z.R {
		o = msgp.AppendString(o, zoqn)
		o = msgp.AppendUint8(o, zsab)
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
	for zurg := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zurg])
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
	const maxFields2zysb = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zysb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zysb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft2zysb := totalEncodedFields2zysb
	missingFieldsLeft2zysb := maxFields2zysb - totalEncodedFields2zysb

	var nextMiss2zysb int32 = -1
	var found2zysb [maxFields2zysb]bool
	var curField2zysb string

doneWithStruct2zysb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zysb > 0 || missingFieldsLeft2zysb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zysb, missingFieldsLeft2zysb, msgp.ShowFound(found2zysb[:]), unmarshalMsgFieldOrder2zysb)
		if encodedFieldsLeft2zysb > 0 {
			encodedFieldsLeft2zysb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField2zysb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zysb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zysb = 0
			}
			for nextMiss2zysb < maxFields2zysb && (found2zysb[nextMiss2zysb] || unmarshalMsgFieldSkip2zysb[nextMiss2zysb]) {
				nextMiss2zysb++
			}
			if nextMiss2zysb == maxFields2zysb {
				// filled all the empty fields!
				break doneWithStruct2zysb
			}
			missingFieldsLeft2zysb--
			curField2zysb = unmarshalMsgFieldOrder2zysb[nextMiss2zysb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zysb)
		switch curField2zysb {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zysb[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ralph":
			found2zysb[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zmxf uint32
				zmxf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.R == nil && zmxf > 0 {
					z.R = make(map[string]uint8, zmxf)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zmxf > 0 {
					var zoqn string
					var zsab uint8
					zmxf--
					zoqn, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zsab, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						panic(err)
					}
					z.R[zoqn] = zsab
				}
			}
		case "P":
			found2zysb[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Q":
			found2zysb[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "T":
			found2zysb[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "arr":
			found2zysb[6] = true
			var zlqy uint32
			zlqy, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				panic(err)
			}
			if !nbs.IsNil(bts) && zlqy != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zlqy}
				return
			}
			for zurg := range z.Arr {
				z.Arr[zurg], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					panic(err)
				}
			}
		case "MyTree":
			found2zysb[7] = true
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
	if nextMiss2zysb != -1 {
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
var unmarshalMsgFieldOrder2zysb = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zysb = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zoqn, zsab := range z.R {
			_ = zsab
			_ = zoqn
			s += msgp.StringPrefixSize + len(zoqn) + msgp.Uint8Size
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
	const maxFields3zzpi = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zzpi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zzpi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zzpi := totalEncodedFields3zzpi
	missingFieldsLeft3zzpi := maxFields3zzpi - totalEncodedFields3zzpi

	var nextMiss3zzpi int32 = -1
	var found3zzpi [maxFields3zzpi]bool
	var curField3zzpi string

doneWithStruct3zzpi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zzpi > 0 || missingFieldsLeft3zzpi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zzpi, missingFieldsLeft3zzpi, msgp.ShowFound(found3zzpi[:]), unmarshalMsgFieldOrder3zzpi)
		if encodedFieldsLeft3zzpi > 0 {
			encodedFieldsLeft3zzpi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zzpi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zzpi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zzpi = 0
			}
			for nextMiss3zzpi < maxFields3zzpi && (found3zzpi[nextMiss3zzpi] || unmarshalMsgFieldSkip3zzpi[nextMiss3zzpi]) {
				nextMiss3zzpi++
			}
			if nextMiss3zzpi == maxFields3zzpi {
				// filled all the empty fields!
				break doneWithStruct3zzpi
			}
			missingFieldsLeft3zzpi--
			curField3zzpi = unmarshalMsgFieldOrder3zzpi[nextMiss3zzpi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zzpi)
		switch curField3zzpi {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zzpi[0] = true
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
	if nextMiss3zzpi != -1 {
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
var unmarshalMsgFieldOrder3zzpi = []string{"F"}

var unmarshalMsgFieldSkip3zzpi = []bool{false}

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
	for zirp := range z.Chld {
		o, err = z.Chld[zirp].MSGPMarshalMsg(o)
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
	const maxFields4zvaj = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zvaj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zvaj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zvaj := totalEncodedFields4zvaj
	missingFieldsLeft4zvaj := maxFields4zvaj - totalEncodedFields4zvaj

	var nextMiss4zvaj int32 = -1
	var found4zvaj [maxFields4zvaj]bool
	var curField4zvaj string

doneWithStruct4zvaj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zvaj > 0 || missingFieldsLeft4zvaj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zvaj, missingFieldsLeft4zvaj, msgp.ShowFound(found4zvaj[:]), unmarshalMsgFieldOrder4zvaj)
		if encodedFieldsLeft4zvaj > 0 {
			encodedFieldsLeft4zvaj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zvaj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zvaj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zvaj = 0
			}
			for nextMiss4zvaj < maxFields4zvaj && (found4zvaj[nextMiss4zvaj] || unmarshalMsgFieldSkip4zvaj[nextMiss4zvaj]) {
				nextMiss4zvaj++
			}
			if nextMiss4zvaj == maxFields4zvaj {
				// filled all the empty fields!
				break doneWithStruct4zvaj
			}
			missingFieldsLeft4zvaj--
			curField4zvaj = unmarshalMsgFieldOrder4zvaj[nextMiss4zvaj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zvaj)
		switch curField4zvaj {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zvaj[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zkld uint32
				zkld, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Chld) >= int(zkld) {
					z.Chld = (z.Chld)[:zkld]
				} else {
					z.Chld = make([]Tree, zkld)
				}
				for zirp := range z.Chld {
					bts, err = z.Chld[zirp].MSGPUnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Str":
			found4zvaj[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Par":
			found4zvaj[2] = true
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
	if nextMiss4zvaj != -1 {
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
var unmarshalMsgFieldOrder4zvaj = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zvaj = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zirp := range z.Chld {
		s += z.Chld[zirp].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
