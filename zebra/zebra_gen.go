package zebra

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Field) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zoux = 11

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zoux uint32
	totalEncodedFields0zoux, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zoux := totalEncodedFields0zoux
	missingFieldsLeft0zoux := maxFields0zoux - totalEncodedFields0zoux

	var nextMiss0zoux int32 = -1
	var found0zoux [maxFields0zoux]bool
	var curField0zoux string

doneWithStruct0zoux:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zoux > 0 || missingFieldsLeft0zoux > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zoux, missingFieldsLeft0zoux, msgp.ShowFound(found0zoux[:]), decodeMsgFieldOrder0zoux)
		if encodedFieldsLeft0zoux > 0 {
			encodedFieldsLeft0zoux--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zoux = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zoux < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zoux = 0
			}
			for nextMiss0zoux < maxFields0zoux && (found0zoux[nextMiss0zoux] || decodeMsgFieldSkip0zoux[nextMiss0zoux]) {
				nextMiss0zoux++
			}
			if nextMiss0zoux == maxFields0zoux {
				// filled all the empty fields!
				break doneWithStruct0zoux
			}
			missingFieldsLeft0zoux--
			curField0zoux = decodeMsgFieldOrder0zoux[nextMiss0zoux]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zoux)
		switch curField0zoux {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zoux[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "FieldGoName":
			found0zoux[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTagName":
			found0zoux[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTypeStr":
			found0zoux[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldCategory":
			found0zoux[4] = true
			{
				var zlmq uint64
				zlmq, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zlmq)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive":
			found0zoux[5] = true
			{
				var zsyu uint64
				zsyu, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zsyu)
			}
			if err != nil {
				return
			}
		case "FieldFullType":
			found0zoux[6] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.FieldFullType != nil {
					dc.PushAlwaysNil()
					err = z.FieldFullType.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.FieldFullType == nil {
					z.FieldFullType = new(Ztype)
				}
				err = z.FieldFullType.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "OmitEmpty":
			found0zoux[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Skip":
			found0zoux[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Deprecated":
			found0zoux[9] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "ShowZero":
			found0zoux[10] = true
			z.ShowZero, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zoux != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zoux = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated", "ShowZero"}

var decodeMsgFieldSkip0zoux = []bool{false, false, false, false, false, false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Field) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 11
	}
	var fieldsInUse uint32 = 11
	isempty[2] = (len(z.FieldTagName) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.FieldTypeStr) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.FieldCategory == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.FieldPrimitive == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}
	isempty[6] = (z.FieldFullType == nil) // pointer, omitempty
	if isempty[6] {
		fieldsInUse--
	}
	isempty[7] = (!z.OmitEmpty) // bool, omitempty
	if isempty[7] {
		fieldsInUse--
	}
	isempty[8] = (!z.Skip) // bool, omitempty
	if isempty[8] {
		fieldsInUse--
	}
	isempty[9] = (!z.Deprecated) // bool, omitempty
	if isempty[9] {
		fieldsInUse--
	}
	isempty[10] = (!z.ShowZero) // bool, omitempty
	if isempty[10] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Field) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zusa [11]bool
	fieldsInUse_zykz := z.fieldsNotEmpty(empty_zusa[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zykz)
	if err != nil {
		return err
	}

	// write "Zid"
	err = en.Append(0xa3, 0x5a, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Zid)
	if err != nil {
		return
	}
	// write "FieldGoName"
	err = en.Append(0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldGoName)
	if err != nil {
		return
	}
	if !empty_zusa[2] {
		// write "FieldTagName"
		err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.FieldTagName)
		if err != nil {
			return
		}
	}

	if !empty_zusa[3] {
		// write "FieldTypeStr"
		err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.FieldTypeStr)
		if err != nil {
			return
		}
	}

	if !empty_zusa[4] {
		// write "FieldCategory"
		err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.FieldCategory))
		if err != nil {
			return
		}
	}

	if !empty_zusa[5] {
		// write "FieldPrimitive"
		err = en.Append(0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.FieldPrimitive))
		if err != nil {
			return
		}
	}

	if !empty_zusa[6] {
		// write "FieldFullType"
		err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65)
		if err != nil {
			return err
		}
		if z.FieldFullType == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.FieldFullType.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zusa[7] {
		// write "OmitEmpty"
		err = en.Append(0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.OmitEmpty)
		if err != nil {
			return
		}
	}

	if !empty_zusa[8] {
		// write "Skip"
		err = en.Append(0xa4, 0x53, 0x6b, 0x69, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Skip)
		if err != nil {
			return
		}
	}

	if !empty_zusa[9] {
		// write "Deprecated"
		err = en.Append(0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Deprecated)
		if err != nil {
			return
		}
	}

	if !empty_zusa[10] {
		// write "ShowZero"
		err = en.Append(0xa8, 0x53, 0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.ShowZero)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Field) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [11]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Zid"
	o = append(o, 0xa3, 0x5a, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Zid)
	// string "FieldGoName"
	o = append(o, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FieldGoName)
	if !empty[2] {
		// string "FieldTagName"
		o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.FieldTagName)
	}

	if !empty[3] {
		// string "FieldTypeStr"
		o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
		o = msgp.AppendString(o, z.FieldTypeStr)
	}

	if !empty[4] {
		// string "FieldCategory"
		o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
		o = msgp.AppendUint64(o, uint64(z.FieldCategory))
	}

	if !empty[5] {
		// string "FieldPrimitive"
		o = append(o, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
		o = msgp.AppendUint64(o, uint64(z.FieldPrimitive))
	}

	if !empty[6] {
		// string "FieldFullType"
		o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65)
		if z.FieldFullType == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.FieldFullType.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[7] {
		// string "OmitEmpty"
		o = append(o, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
		o = msgp.AppendBool(o, z.OmitEmpty)
	}

	if !empty[8] {
		// string "Skip"
		o = append(o, 0xa4, 0x53, 0x6b, 0x69, 0x70)
		o = msgp.AppendBool(o, z.Skip)
	}

	if !empty[9] {
		// string "Deprecated"
		o = append(o, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
		o = msgp.AppendBool(o, z.Deprecated)
	}

	if !empty[10] {
		// string "ShowZero"
		o = append(o, 0xa8, 0x53, 0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f)
		o = msgp.AppendBool(o, z.ShowZero)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Field) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Field) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zour = 11

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zour uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zour, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zour := totalEncodedFields1zour
	missingFieldsLeft1zour := maxFields1zour - totalEncodedFields1zour

	var nextMiss1zour int32 = -1
	var found1zour [maxFields1zour]bool
	var curField1zour string

doneWithStruct1zour:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zour > 0 || missingFieldsLeft1zour > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zour, missingFieldsLeft1zour, msgp.ShowFound(found1zour[:]), unmarshalMsgFieldOrder1zour)
		if encodedFieldsLeft1zour > 0 {
			encodedFieldsLeft1zour--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zour = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zour < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zour = 0
			}
			for nextMiss1zour < maxFields1zour && (found1zour[nextMiss1zour] || unmarshalMsgFieldSkip1zour[nextMiss1zour]) {
				nextMiss1zour++
			}
			if nextMiss1zour == maxFields1zour {
				// filled all the empty fields!
				break doneWithStruct1zour
			}
			missingFieldsLeft1zour--
			curField1zour = unmarshalMsgFieldOrder1zour[nextMiss1zour]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zour)
		switch curField1zour {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zour[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "FieldGoName":
			found1zour[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "FieldTagName":
			found1zour[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "FieldTypeStr":
			found1zour[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "FieldCategory":
			found1zour[4] = true
			{
				var zaay uint64
				zaay, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					return
				}
				z.FieldCategory = Zkind(zaay)
			}
		case "FieldPrimitive":
			found1zour[5] = true
			{
				var zppr uint64
				zppr, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					return
				}
				z.FieldPrimitive = Zkind(zppr)
			}
		case "FieldFullType":
			found1zour[6] = true
			if nbs.AlwaysNil {
				if z.FieldFullType != nil {
					z.FieldFullType.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.FieldFullType {
						z.FieldFullType.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.FieldFullType == nil {
						z.FieldFullType = new(Ztype)
					}
					bts, err = z.FieldFullType.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "OmitEmpty":
			found1zour[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "Skip":
			found1zour[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "Deprecated":
			found1zour[9] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "ShowZero":
			found1zour[10] = true
			z.ShowZero, bts, err = nbs.ReadBoolBytes(bts)

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
	if nextMiss1zour != -1 {
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

// fields of Field
var unmarshalMsgFieldOrder1zour = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated", "ShowZero"}

var unmarshalMsgFieldSkip1zour = []bool{false, false, false, false, false, false, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Field) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.FieldGoName) + 13 + msgp.StringPrefixSize + len(z.FieldTagName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 14 + msgp.Uint64Size + 15 + msgp.Uint64Size + 14
	if z.FieldFullType == nil {
		s += msgp.NilSize
	} else {
		s += z.FieldFullType.Msgsize()
	}
	s += 10 + msgp.BoolSize + 5 + msgp.BoolSize + 11 + msgp.BoolSize + 9 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Schema) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields2zzmc = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zzmc uint32
	totalEncodedFields2zzmc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zzmc := totalEncodedFields2zzmc
	missingFieldsLeft2zzmc := maxFields2zzmc - totalEncodedFields2zzmc

	var nextMiss2zzmc int32 = -1
	var found2zzmc [maxFields2zzmc]bool
	var curField2zzmc string

doneWithStruct2zzmc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zzmc > 0 || missingFieldsLeft2zzmc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zzmc, missingFieldsLeft2zzmc, msgp.ShowFound(found2zzmc[:]), decodeMsgFieldOrder2zzmc)
		if encodedFieldsLeft2zzmc > 0 {
			encodedFieldsLeft2zzmc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zzmc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zzmc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zzmc = 0
			}
			for nextMiss2zzmc < maxFields2zzmc && (found2zzmc[nextMiss2zzmc] || decodeMsgFieldSkip2zzmc[nextMiss2zzmc]) {
				nextMiss2zzmc++
			}
			if nextMiss2zzmc == maxFields2zzmc {
				// filled all the empty fields!
				break doneWithStruct2zzmc
			}
			missingFieldsLeft2zzmc--
			curField2zzmc = decodeMsgFieldOrder2zzmc[nextMiss2zzmc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zzmc)
		switch curField2zzmc {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zzmc[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SourcePackage":
			found2zzmc[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ZebraSchemaId":
			found2zzmc[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Structs":
			found2zzmc[3] = true
			var zgog uint32
			zgog, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Structs == nil && zgog > 0 {
				z.Structs = make(map[string]*Struct, zgog)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zgog > 0 {
				zgog--
				var zxxy string
				var zuiq *Struct
				zxxy, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					zuiq = nil
				} else {
					if zuiq == nil {
						zuiq = new(Struct)
					}
					const maxFields3zvst = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zvst uint32
					totalEncodedFields3zvst, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zvst := totalEncodedFields3zvst
					missingFieldsLeft3zvst := maxFields3zvst - totalEncodedFields3zvst

					var nextMiss3zvst int32 = -1
					var found3zvst [maxFields3zvst]bool
					var curField3zvst string

				doneWithStruct3zvst:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zvst > 0 || missingFieldsLeft3zvst > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvst, missingFieldsLeft3zvst, msgp.ShowFound(found3zvst[:]), decodeMsgFieldOrder3zvst)
						if encodedFieldsLeft3zvst > 0 {
							encodedFieldsLeft3zvst--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zvst = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zvst < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zvst = 0
							}
							for nextMiss3zvst < maxFields3zvst && (found3zvst[nextMiss3zvst] || decodeMsgFieldSkip3zvst[nextMiss3zvst]) {
								nextMiss3zvst++
							}
							if nextMiss3zvst == maxFields3zvst {
								// filled all the empty fields!
								break doneWithStruct3zvst
							}
							missingFieldsLeft3zvst--
							curField3zvst = decodeMsgFieldOrder3zvst[nextMiss3zvst]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zvst)
						switch curField3zvst {
						// -- templateDecodeMsg ends here --

						case "StructName":
							found3zvst[0] = true
							zuiq.StructName, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Fields":
							found3zvst[1] = true
							var zlot uint32
							zlot, err = dc.ReadArrayHeader()
							if err != nil {
								return
							}
							if cap(zuiq.Fields) >= int(zlot) {
								zuiq.Fields = (zuiq.Fields)[:zlot]
							} else {
								zuiq.Fields = make([]Field, zlot)
							}
							for zrvo := range zuiq.Fields {
								err = zuiq.Fields[zrvo].DecodeMsg(dc)
								if err != nil {
									return
								}
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
					if nextMiss3zvst != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[zxxy] = zuiq
			}
		case "Imports":
			found2zzmc[4] = true
			var zxba uint32
			zxba, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Imports) >= int(zxba) {
				z.Imports = (z.Imports)[:zxba]
			} else {
				z.Imports = make([]string, zxba)
			}
			for zoeo := range z.Imports {
				z.Imports[zoeo], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss2zzmc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zzmc = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zzmc = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zvst = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zvst = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Schema) fieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// EncodeMsg implements msgp.Encodable
func (z *Schema) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 5
	// write "SourcePath"
	err = en.Append(0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePath)
	if err != nil {
		return
	}
	// write "SourcePackage"
	err = en.Append(0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePackage)
	if err != nil {
		return
	}
	// write "ZebraSchemaId"
	err = en.Append(0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.ZebraSchemaId)
	if err != nil {
		return
	}
	// write "Structs"
	err = en.Append(0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Structs)))
	if err != nil {
		return
	}
	for zxxy, zuiq := range z.Structs {
		err = en.WriteString(zxxy)
		if err != nil {
			return
		}
		if zuiq == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "StructName"
			err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(zuiq.StructName)
			if err != nil {
				return
			}
			// write "Fields"
			err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(zuiq.Fields)))
			if err != nil {
				return
			}
			for zrvo := range zuiq.Fields {
				err = zuiq.Fields[zrvo].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}
	// write "Imports"
	err = en.Append(0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Imports)))
	if err != nil {
		return
	}
	for zoeo := range z.Imports {
		err = en.WriteString(z.Imports[zoeo])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Schema) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "SourcePath"
	o = append(o, 0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.SourcePath)
	// string "SourcePackage"
	o = append(o, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.SourcePackage)
	// string "ZebraSchemaId"
	o = append(o, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.ZebraSchemaId)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Structs)))
	for zxxy, zuiq := range z.Structs {
		o = msgp.AppendString(o, zxxy)
		if zuiq == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName"
			o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, zuiq.StructName)
			// string "Fields"
			o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			o = msgp.AppendArrayHeader(o, uint32(len(zuiq.Fields)))
			for zrvo := range zuiq.Fields {
				o, err = zuiq.Fields[zrvo].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zoeo := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zoeo])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Schema) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Schema) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zmxv = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zmxv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zmxv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zmxv := totalEncodedFields4zmxv
	missingFieldsLeft4zmxv := maxFields4zmxv - totalEncodedFields4zmxv

	var nextMiss4zmxv int32 = -1
	var found4zmxv [maxFields4zmxv]bool
	var curField4zmxv string

doneWithStruct4zmxv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zmxv > 0 || missingFieldsLeft4zmxv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zmxv, missingFieldsLeft4zmxv, msgp.ShowFound(found4zmxv[:]), unmarshalMsgFieldOrder4zmxv)
		if encodedFieldsLeft4zmxv > 0 {
			encodedFieldsLeft4zmxv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zmxv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zmxv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zmxv = 0
			}
			for nextMiss4zmxv < maxFields4zmxv && (found4zmxv[nextMiss4zmxv] || unmarshalMsgFieldSkip4zmxv[nextMiss4zmxv]) {
				nextMiss4zmxv++
			}
			if nextMiss4zmxv == maxFields4zmxv {
				// filled all the empty fields!
				break doneWithStruct4zmxv
			}
			missingFieldsLeft4zmxv--
			curField4zmxv = unmarshalMsgFieldOrder4zmxv[nextMiss4zmxv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zmxv)
		switch curField4zmxv {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zmxv[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "SourcePackage":
			found4zmxv[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ZebraSchemaId":
			found4zmxv[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "Structs":
			found4zmxv[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zvlz uint32
				zvlz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Structs == nil && zvlz > 0 {
					z.Structs = make(map[string]*Struct, zvlz)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zvlz > 0 {
					var zxxy string
					var zuiq *Struct
					zvlz--
					zxxy, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && zuiq == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if zuiq == nil {
							zuiq = new(Struct)
						}
						const maxFields5zqxx = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zqxx uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zqxx, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								return
							}
						}
						encodedFieldsLeft5zqxx := totalEncodedFields5zqxx
						missingFieldsLeft5zqxx := maxFields5zqxx - totalEncodedFields5zqxx

						var nextMiss5zqxx int32 = -1
						var found5zqxx [maxFields5zqxx]bool
						var curField5zqxx string

					doneWithStruct5zqxx:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zqxx > 0 || missingFieldsLeft5zqxx > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zqxx, missingFieldsLeft5zqxx, msgp.ShowFound(found5zqxx[:]), unmarshalMsgFieldOrder5zqxx)
							if encodedFieldsLeft5zqxx > 0 {
								encodedFieldsLeft5zqxx--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									return
								}
								curField5zqxx = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zqxx < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zqxx = 0
								}
								for nextMiss5zqxx < maxFields5zqxx && (found5zqxx[nextMiss5zqxx] || unmarshalMsgFieldSkip5zqxx[nextMiss5zqxx]) {
									nextMiss5zqxx++
								}
								if nextMiss5zqxx == maxFields5zqxx {
									// filled all the empty fields!
									break doneWithStruct5zqxx
								}
								missingFieldsLeft5zqxx--
								curField5zqxx = unmarshalMsgFieldOrder5zqxx[nextMiss5zqxx]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zqxx)
							switch curField5zqxx {
							// -- templateUnmarshalMsg ends here --

							case "StructName":
								found5zqxx[0] = true
								zuiq.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									return
								}
							case "Fields":
								found5zqxx[1] = true
								if nbs.AlwaysNil {
									(zuiq.Fields) = (zuiq.Fields)[:0]
								} else {

									var zjsd uint32
									zjsd, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										return
									}
									if cap(zuiq.Fields) >= int(zjsd) {
										zuiq.Fields = (zuiq.Fields)[:zjsd]
									} else {
										zuiq.Fields = make([]Field, zjsd)
									}
									for zrvo := range zuiq.Fields {
										bts, err = zuiq.Fields[zrvo].UnmarshalMsg(bts)
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
						if nextMiss5zqxx != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[zxxy] = zuiq
				}
			}
		case "Imports":
			found4zmxv[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zdoh uint32
				zdoh, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Imports) >= int(zdoh) {
					z.Imports = (z.Imports)[:zdoh]
				} else {
					z.Imports = make([]string, zdoh)
				}
				for zoeo := range z.Imports {
					z.Imports[zoeo], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zmxv != -1 {
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

// fields of Schema
var unmarshalMsgFieldOrder4zmxv = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zmxv = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zqxx = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zqxx = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zxxy, zuiq := range z.Structs {
			_ = zuiq
			_ = zxxy
			s += msgp.StringPrefixSize + len(zxxy)
			if zuiq == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 11 + msgp.StringPrefixSize + len(zuiq.StructName) + 7 + msgp.ArrayHeaderSize
				for zrvo := range zuiq.Fields {
					s += zuiq.Fields[zrvo].Msgsize()
				}
			}
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zoeo := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zoeo])
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields6zwfb = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zwfb uint32
	totalEncodedFields6zwfb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zwfb := totalEncodedFields6zwfb
	missingFieldsLeft6zwfb := maxFields6zwfb - totalEncodedFields6zwfb

	var nextMiss6zwfb int32 = -1
	var found6zwfb [maxFields6zwfb]bool
	var curField6zwfb string

doneWithStruct6zwfb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zwfb > 0 || missingFieldsLeft6zwfb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zwfb, missingFieldsLeft6zwfb, msgp.ShowFound(found6zwfb[:]), decodeMsgFieldOrder6zwfb)
		if encodedFieldsLeft6zwfb > 0 {
			encodedFieldsLeft6zwfb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zwfb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zwfb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zwfb = 0
			}
			for nextMiss6zwfb < maxFields6zwfb && (found6zwfb[nextMiss6zwfb] || decodeMsgFieldSkip6zwfb[nextMiss6zwfb]) {
				nextMiss6zwfb++
			}
			if nextMiss6zwfb == maxFields6zwfb {
				// filled all the empty fields!
				break doneWithStruct6zwfb
			}
			missingFieldsLeft6zwfb--
			curField6zwfb = decodeMsgFieldOrder6zwfb[nextMiss6zwfb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zwfb)
		switch curField6zwfb {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zwfb[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Fields":
			found6zwfb[1] = true
			var zvgk uint32
			zvgk, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zvgk) {
				z.Fields = (z.Fields)[:zvgk]
			} else {
				z.Fields = make([]Field, zvgk)
			}
			for zjfz := range z.Fields {
				err = z.Fields[zjfz].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zwfb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zwfb = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zwfb = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Struct) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *Struct) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 2
	// write "StructName"
	err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StructName)
	if err != nil {
		return
	}
	// write "Fields"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fields)))
	if err != nil {
		return
	}
	for zjfz := range z.Fields {
		err = z.Fields[zjfz].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Struct) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "StructName"
	o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.StructName)
	// string "Fields"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fields)))
	for zjfz := range z.Fields {
		o, err = z.Fields[zjfz].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Struct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Struct) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields7ziex = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7ziex uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7ziex, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft7ziex := totalEncodedFields7ziex
	missingFieldsLeft7ziex := maxFields7ziex - totalEncodedFields7ziex

	var nextMiss7ziex int32 = -1
	var found7ziex [maxFields7ziex]bool
	var curField7ziex string

doneWithStruct7ziex:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7ziex > 0 || missingFieldsLeft7ziex > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7ziex, missingFieldsLeft7ziex, msgp.ShowFound(found7ziex[:]), unmarshalMsgFieldOrder7ziex)
		if encodedFieldsLeft7ziex > 0 {
			encodedFieldsLeft7ziex--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField7ziex = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7ziex < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7ziex = 0
			}
			for nextMiss7ziex < maxFields7ziex && (found7ziex[nextMiss7ziex] || unmarshalMsgFieldSkip7ziex[nextMiss7ziex]) {
				nextMiss7ziex++
			}
			if nextMiss7ziex == maxFields7ziex {
				// filled all the empty fields!
				break doneWithStruct7ziex
			}
			missingFieldsLeft7ziex--
			curField7ziex = unmarshalMsgFieldOrder7ziex[nextMiss7ziex]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7ziex)
		switch curField7ziex {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7ziex[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Fields":
			found7ziex[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zqrv uint32
				zqrv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Fields) >= int(zqrv) {
					z.Fields = (z.Fields)[:zqrv]
				} else {
					z.Fields = make([]Field, zqrv)
				}
				for zjfz := range z.Fields {
					bts, err = z.Fields[zjfz].UnmarshalMsg(bts)
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
	if nextMiss7ziex != -1 {
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

// fields of Struct
var unmarshalMsgFieldOrder7ziex = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7ziex = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zjfz := range z.Fields {
		s += z.Fields[zjfz].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Zkind) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var zufi uint64
		zufi, err = dc.ReadUint64()
		(*z) = Zkind(zufi)
	}
	if err != nil {
		return
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z Zkind) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	err = en.WriteUint64(uint64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zkind) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zkind) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Zkind) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zsqw uint64
		zsqw, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			return
		}
		(*z) = Zkind(zsqw)
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

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zkind) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Ztype) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields8ztck = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8ztck uint32
	totalEncodedFields8ztck, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8ztck := totalEncodedFields8ztck
	missingFieldsLeft8ztck := maxFields8ztck - totalEncodedFields8ztck

	var nextMiss8ztck int32 = -1
	var found8ztck [maxFields8ztck]bool
	var curField8ztck string

doneWithStruct8ztck:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8ztck > 0 || missingFieldsLeft8ztck > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8ztck, missingFieldsLeft8ztck, msgp.ShowFound(found8ztck[:]), decodeMsgFieldOrder8ztck)
		if encodedFieldsLeft8ztck > 0 {
			encodedFieldsLeft8ztck--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8ztck = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8ztck < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8ztck = 0
			}
			for nextMiss8ztck < maxFields8ztck && (found8ztck[nextMiss8ztck] || decodeMsgFieldSkip8ztck[nextMiss8ztck]) {
				nextMiss8ztck++
			}
			if nextMiss8ztck == maxFields8ztck {
				// filled all the empty fields!
				break doneWithStruct8ztck
			}
			missingFieldsLeft8ztck--
			curField8ztck = decodeMsgFieldOrder8ztck[nextMiss8ztck]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8ztck)
		switch curField8ztck {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8ztck[0] = true
			{
				var zcgj uint64
				zcgj, err = dc.ReadUint64()
				z.Kind = Zkind(zcgj)
			}
			if err != nil {
				return
			}
		case "Str":
			found8ztck[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Domain":
			found8ztck[2] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.Domain != nil {
					dc.PushAlwaysNil()
					err = z.Domain.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.Domain == nil {
					z.Domain = new(Ztype)
				}
				err = z.Domain.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Range":
			found8ztck[3] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.Range != nil {
					dc.PushAlwaysNil()
					err = z.Range.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.Range == nil {
					z.Range = new(Ztype)
				}
				err = z.Range.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss8ztck != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8ztck = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8ztck = []bool{false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Ztype) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 4
	}
	var fieldsInUse uint32 = 4
	isempty[1] = (len(z.Str) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Domain == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Range == nil) // pointer, omitempty
	if isempty[3] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Ztype) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zicp [4]bool
	fieldsInUse_zbvn := z.fieldsNotEmpty(empty_zicp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zbvn)
	if err != nil {
		return err
	}

	// write "Kind"
	err = en.Append(0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Kind))
	if err != nil {
		return
	}
	if !empty_zicp[1] {
		// write "Str"
		err = en.Append(0xa3, 0x53, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Str)
		if err != nil {
			return
		}
	}

	if !empty_zicp[2] {
		// write "Domain"
		err = en.Append(0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
		if err != nil {
			return err
		}
		if z.Domain == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Domain.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zicp[3] {
		// write "Range"
		err = en.Append(0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65)
		if err != nil {
			return err
		}
		if z.Range == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Range.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Ztype) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [4]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Kind"
	o = append(o, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Kind))
	if !empty[1] {
		// string "Str"
		o = append(o, 0xa3, 0x53, 0x74, 0x72)
		o = msgp.AppendString(o, z.Str)
	}

	if !empty[2] {
		// string "Domain"
		o = append(o, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
		if z.Domain == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Domain.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[3] {
		// string "Range"
		o = append(o, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65)
		if z.Range == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Range.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Ztype) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Ztype) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields9zohm = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zohm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zohm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft9zohm := totalEncodedFields9zohm
	missingFieldsLeft9zohm := maxFields9zohm - totalEncodedFields9zohm

	var nextMiss9zohm int32 = -1
	var found9zohm [maxFields9zohm]bool
	var curField9zohm string

doneWithStruct9zohm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zohm > 0 || missingFieldsLeft9zohm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zohm, missingFieldsLeft9zohm, msgp.ShowFound(found9zohm[:]), unmarshalMsgFieldOrder9zohm)
		if encodedFieldsLeft9zohm > 0 {
			encodedFieldsLeft9zohm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField9zohm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zohm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zohm = 0
			}
			for nextMiss9zohm < maxFields9zohm && (found9zohm[nextMiss9zohm] || unmarshalMsgFieldSkip9zohm[nextMiss9zohm]) {
				nextMiss9zohm++
			}
			if nextMiss9zohm == maxFields9zohm {
				// filled all the empty fields!
				break doneWithStruct9zohm
			}
			missingFieldsLeft9zohm--
			curField9zohm = unmarshalMsgFieldOrder9zohm[nextMiss9zohm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zohm)
		switch curField9zohm {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zohm[0] = true
			{
				var zufv uint64
				zufv, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					return
				}
				z.Kind = Zkind(zufv)
			}
		case "Str":
			found9zohm[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Domain":
			found9zohm[2] = true
			if nbs.AlwaysNil {
				if z.Domain != nil {
					z.Domain.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Domain {
						z.Domain.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Domain == nil {
						z.Domain = new(Ztype)
					}
					bts, err = z.Domain.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Range":
			found9zohm[3] = true
			if nbs.AlwaysNil {
				if z.Range != nil {
					z.Range.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Range {
						z.Range.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Range == nil {
						z.Range = new(Ztype)
					}
					bts, err = z.Range.UnmarshalMsg(bts)
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
	if nextMiss9zohm != -1 {
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

// fields of Ztype
var unmarshalMsgFieldOrder9zohm = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zohm = []bool{false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Ztype) Msgsize() (s int) {
	s = 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Str) + 7
	if z.Domain == nil {
		s += msgp.NilSize
	} else {
		s += z.Domain.Msgsize()
	}
	s += 6
	if z.Range == nil {
		s += msgp.NilSize
	} else {
		s += z.Range.Msgsize()
	}
	return
}
