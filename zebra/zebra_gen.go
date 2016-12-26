package zebra

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

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
	const maxFields0zwtk = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zwtk uint32
	totalEncodedFields0zwtk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zwtk := totalEncodedFields0zwtk
	missingFieldsLeft0zwtk := maxFields0zwtk - totalEncodedFields0zwtk

	var nextMiss0zwtk int32 = -1
	var found0zwtk [maxFields0zwtk]bool
	var curField0zwtk string

doneWithStruct0zwtk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zwtk > 0 || missingFieldsLeft0zwtk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zwtk, missingFieldsLeft0zwtk, msgp.ShowFound(found0zwtk[:]), decodeMsgFieldOrder0zwtk)
		if encodedFieldsLeft0zwtk > 0 {
			encodedFieldsLeft0zwtk--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zwtk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zwtk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zwtk = 0
			}
			for nextMiss0zwtk < maxFields0zwtk && (found0zwtk[nextMiss0zwtk] || decodeMsgFieldSkip0zwtk[nextMiss0zwtk]) {
				nextMiss0zwtk++
			}
			if nextMiss0zwtk == maxFields0zwtk {
				// filled all the empty fields!
				break doneWithStruct0zwtk
			}
			missingFieldsLeft0zwtk--
			curField0zwtk = decodeMsgFieldOrder0zwtk[nextMiss0zwtk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zwtk)
		switch curField0zwtk {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zwtk[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zwtk[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zwtk[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zwtk[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zwtk[4] = true
			{
				var zaof uint64
				zaof, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zaof)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zwtk[5] = true
			{
				var zrlr uint64
				zrlr, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zrlr)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zwtk[6] = true
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
					panic(err)
				}
			}
		case "OmitEmpty":
			found0zwtk[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zwtk[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zwtk[9] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zwtk != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zwtk = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zwtk = []bool{false, false, false, false, false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Field) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 10
	}
	var fieldsInUse uint32 = 10
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

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Field) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zhjp [10]bool
	fieldsInUse_zlhr := z.fieldsNotEmpty(empty_zhjp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlhr)
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
		panic(err)
	}
	// write "FieldGoName"
	err = en.Append(0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldGoName)
	if err != nil {
		panic(err)
	}
	// write "FieldTagName"
	err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldTagName)
	if err != nil {
		panic(err)
	}
	// write "FieldTypeStr"
	err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldTypeStr)
	if err != nil {
		panic(err)
	}
	// write "FieldCategory"
	err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldCategory))
	if err != nil {
		panic(err)
	}
	// write "FieldPrimitive"
	err = en.Append(0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldPrimitive))
	if err != nil {
		panic(err)
	}
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
			panic(err)
		}
	}
	if !empty_zhjp[7] {
		// write "OmitEmpty"
		err = en.Append(0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.OmitEmpty)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zhjp[8] {
		// write "Skip"
		err = en.Append(0xa4, 0x53, 0x6b, 0x69, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Skip)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zhjp[9] {
		// write "Deprecated"
		err = en.Append(0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Deprecated)
		if err != nil {
			panic(err)
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Field) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [10]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Zid"
	o = append(o, 0xa3, 0x5a, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Zid)
	// string "FieldGoName"
	o = append(o, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FieldGoName)
	// string "FieldTagName"
	o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FieldTagName)
	// string "FieldTypeStr"
	o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldTypeStr)
	// string "FieldCategory"
	o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	o = msgp.AppendUint64(o, uint64(z.FieldCategory))
	// string "FieldPrimitive"
	o = append(o, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendUint64(o, uint64(z.FieldPrimitive))
	// string "FieldFullType"
	o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65)
	if z.FieldFullType == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.FieldFullType.MarshalMsg(o)
		if err != nil {
			panic(err)
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
	const maxFields1zkty = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zkty uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zkty, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zkty := totalEncodedFields1zkty
	missingFieldsLeft1zkty := maxFields1zkty - totalEncodedFields1zkty

	var nextMiss1zkty int32 = -1
	var found1zkty [maxFields1zkty]bool
	var curField1zkty string

doneWithStruct1zkty:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zkty > 0 || missingFieldsLeft1zkty > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zkty, missingFieldsLeft1zkty, msgp.ShowFound(found1zkty[:]), unmarshalMsgFieldOrder1zkty)
		if encodedFieldsLeft1zkty > 0 {
			encodedFieldsLeft1zkty--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zkty = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zkty < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zkty = 0
			}
			for nextMiss1zkty < maxFields1zkty && (found1zkty[nextMiss1zkty] || unmarshalMsgFieldSkip1zkty[nextMiss1zkty]) {
				nextMiss1zkty++
			}
			if nextMiss1zkty == maxFields1zkty {
				// filled all the empty fields!
				break doneWithStruct1zkty
			}
			missingFieldsLeft1zkty--
			curField1zkty = unmarshalMsgFieldOrder1zkty[nextMiss1zkty]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zkty)
		switch curField1zkty {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zkty[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zkty[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zkty[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zkty[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zkty[4] = true
			{
				var zjby uint64
				zjby, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zjby)
			}
		case "FieldPrimitive":
			found1zkty[5] = true
			{
				var zgsn uint64
				zgsn, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zgsn)
			}
		case "FieldFullType":
			found1zkty[6] = true
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
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "OmitEmpty":
			found1zkty[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zkty[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zkty[9] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

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
	if nextMiss1zkty != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zkty = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zkty = []bool{false, false, false, false, false, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Field) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.FieldGoName) + 13 + msgp.StringPrefixSize + len(z.FieldTagName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 14 + msgp.Uint64Size + 15 + msgp.Uint64Size + 14
	if z.FieldFullType == nil {
		s += msgp.NilSize
	} else {
		s += z.FieldFullType.Msgsize()
	}
	s += 10 + msgp.BoolSize + 5 + msgp.BoolSize + 11 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *KS) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zeib = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zeib uint32
	totalEncodedFields2zeib, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zeib := totalEncodedFields2zeib
	missingFieldsLeft2zeib := maxFields2zeib - totalEncodedFields2zeib

	var nextMiss2zeib int32 = -1
	var found2zeib [maxFields2zeib]bool
	var curField2zeib string

doneWithStruct2zeib:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zeib > 0 || missingFieldsLeft2zeib > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zeib, missingFieldsLeft2zeib, msgp.ShowFound(found2zeib[:]), decodeMsgFieldOrder2zeib)
		if encodedFieldsLeft2zeib > 0 {
			encodedFieldsLeft2zeib--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zeib = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zeib < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zeib = 0
			}
			for nextMiss2zeib < maxFields2zeib && (found2zeib[nextMiss2zeib] || decodeMsgFieldSkip2zeib[nextMiss2zeib]) {
				nextMiss2zeib++
			}
			if nextMiss2zeib == maxFields2zeib {
				// filled all the empty fields!
				break doneWithStruct2zeib
			}
			missingFieldsLeft2zeib--
			curField2zeib = decodeMsgFieldOrder2zeib[nextMiss2zeib]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zeib)
		switch curField2zeib {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found2zeib[0] = true
			{
				var zysw uint64
				zysw, err = dc.ReadUint64()
				z.Kind = Zkind(zysw)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found2zeib[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss2zeib != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of KS
var decodeMsgFieldOrder2zeib = []string{"Kind", "Str"}

var decodeMsgFieldSkip2zeib = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z KS) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z KS) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Kind"
	err = en.Append(0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Kind))
	if err != nil {
		panic(err)
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Str)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z KS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Kind"
	o = append(o, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *KS) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zxep = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zxep uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zxep, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zxep := totalEncodedFields3zxep
	missingFieldsLeft3zxep := maxFields3zxep - totalEncodedFields3zxep

	var nextMiss3zxep int32 = -1
	var found3zxep [maxFields3zxep]bool
	var curField3zxep string

doneWithStruct3zxep:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zxep > 0 || missingFieldsLeft3zxep > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zxep, missingFieldsLeft3zxep, msgp.ShowFound(found3zxep[:]), unmarshalMsgFieldOrder3zxep)
		if encodedFieldsLeft3zxep > 0 {
			encodedFieldsLeft3zxep--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zxep = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zxep < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zxep = 0
			}
			for nextMiss3zxep < maxFields3zxep && (found3zxep[nextMiss3zxep] || unmarshalMsgFieldSkip3zxep[nextMiss3zxep]) {
				nextMiss3zxep++
			}
			if nextMiss3zxep == maxFields3zxep {
				// filled all the empty fields!
				break doneWithStruct3zxep
			}
			missingFieldsLeft3zxep--
			curField3zxep = unmarshalMsgFieldOrder3zxep[nextMiss3zxep]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zxep)
		switch curField3zxep {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found3zxep[0] = true
			{
				var zpwm uint64
				zpwm, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zpwm)
			}
		case "Str":
			found3zxep[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss3zxep != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of KS
var unmarshalMsgFieldOrder3zxep = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip3zxep = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z KS) Msgsize() (s int) {
	s = 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Str)
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
	const maxFields4zdnb = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zdnb uint32
	totalEncodedFields4zdnb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zdnb := totalEncodedFields4zdnb
	missingFieldsLeft4zdnb := maxFields4zdnb - totalEncodedFields4zdnb

	var nextMiss4zdnb int32 = -1
	var found4zdnb [maxFields4zdnb]bool
	var curField4zdnb string

doneWithStruct4zdnb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zdnb > 0 || missingFieldsLeft4zdnb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zdnb, missingFieldsLeft4zdnb, msgp.ShowFound(found4zdnb[:]), decodeMsgFieldOrder4zdnb)
		if encodedFieldsLeft4zdnb > 0 {
			encodedFieldsLeft4zdnb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zdnb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zdnb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zdnb = 0
			}
			for nextMiss4zdnb < maxFields4zdnb && (found4zdnb[nextMiss4zdnb] || decodeMsgFieldSkip4zdnb[nextMiss4zdnb]) {
				nextMiss4zdnb++
			}
			if nextMiss4zdnb == maxFields4zdnb {
				// filled all the empty fields!
				break doneWithStruct4zdnb
			}
			missingFieldsLeft4zdnb--
			curField4zdnb = decodeMsgFieldOrder4zdnb[nextMiss4zdnb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zdnb)
		switch curField4zdnb {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found4zdnb[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zdnb[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zdnb[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zdnb[3] = true
			var zbqf uint32
			zbqf, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zbqf) {
				z.Structs = (z.Structs)[:zbqf]
			} else {
				z.Structs = make([]Struct, zbqf)
			}
			for zvoe := range z.Structs {
				const maxFields5zeul = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zeul uint32
				totalEncodedFields5zeul, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zeul := totalEncodedFields5zeul
				missingFieldsLeft5zeul := maxFields5zeul - totalEncodedFields5zeul

				var nextMiss5zeul int32 = -1
				var found5zeul [maxFields5zeul]bool
				var curField5zeul string

			doneWithStruct5zeul:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zeul > 0 || missingFieldsLeft5zeul > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zeul, missingFieldsLeft5zeul, msgp.ShowFound(found5zeul[:]), decodeMsgFieldOrder5zeul)
					if encodedFieldsLeft5zeul > 0 {
						encodedFieldsLeft5zeul--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zeul = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zeul < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zeul = 0
						}
						for nextMiss5zeul < maxFields5zeul && (found5zeul[nextMiss5zeul] || decodeMsgFieldSkip5zeul[nextMiss5zeul]) {
							nextMiss5zeul++
						}
						if nextMiss5zeul == maxFields5zeul {
							// filled all the empty fields!
							break doneWithStruct5zeul
						}
						missingFieldsLeft5zeul--
						curField5zeul = decodeMsgFieldOrder5zeul[nextMiss5zeul]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zeul)
					switch curField5zeul {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zeul[0] = true
						z.Structs[zvoe].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found5zeul[1] = true
						var zchd uint32
						zchd, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zvoe].Fields) >= int(zchd) {
							z.Structs[zvoe].Fields = (z.Structs[zvoe].Fields)[:zchd]
						} else {
							z.Structs[zvoe].Fields = make([]Field, zchd)
						}
						for zweb := range z.Structs[zvoe].Fields {
							err = z.Structs[zvoe].Fields[zweb].DecodeMsg(dc)
							if err != nil {
								panic(err)
							}
						}
					default:
						err = dc.Skip()
						if err != nil {
							panic(err)
						}
					}
				}
				if nextMiss5zeul != -1 {
					dc.PopAlwaysNil()
				}

			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss4zdnb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder4zdnb = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip4zdnb = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder5zeul = []string{"StructName", "Fields"}

var decodeMsgFieldSkip5zeul = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Schema) fieldsNotEmpty(isempty []bool) uint32 {
	return 4
}

// EncodeMsg implements msgp.Encodable
func (z *Schema) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "SourcePath"
	err = en.Append(0x84, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePath)
	if err != nil {
		panic(err)
	}
	// write "SourcePackage"
	err = en.Append(0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePackage)
	if err != nil {
		panic(err)
	}
	// write "ZebraSchemaId"
	err = en.Append(0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.ZebraSchemaId)
	if err != nil {
		panic(err)
	}
	// write "Structs"
	err = en.Append(0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Structs)))
	if err != nil {
		panic(err)
	}
	for zvoe := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zvoe].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zvoe].Fields)))
		if err != nil {
			panic(err)
		}
		for zweb := range z.Structs[zvoe].Fields {
			err = z.Structs[zvoe].Fields[zweb].EncodeMsg(en)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Schema) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "SourcePath"
	o = append(o, 0x84, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.SourcePath)
	// string "SourcePackage"
	o = append(o, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.SourcePackage)
	// string "ZebraSchemaId"
	o = append(o, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.ZebraSchemaId)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zvoe := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zvoe].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zvoe].Fields)))
		for zweb := range z.Structs[zvoe].Fields {
			o, err = z.Structs[zvoe].Fields[zweb].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
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
	const maxFields6zoji = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zoji uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zoji, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zoji := totalEncodedFields6zoji
	missingFieldsLeft6zoji := maxFields6zoji - totalEncodedFields6zoji

	var nextMiss6zoji int32 = -1
	var found6zoji [maxFields6zoji]bool
	var curField6zoji string

doneWithStruct6zoji:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zoji > 0 || missingFieldsLeft6zoji > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zoji, missingFieldsLeft6zoji, msgp.ShowFound(found6zoji[:]), unmarshalMsgFieldOrder6zoji)
		if encodedFieldsLeft6zoji > 0 {
			encodedFieldsLeft6zoji--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zoji = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zoji < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zoji = 0
			}
			for nextMiss6zoji < maxFields6zoji && (found6zoji[nextMiss6zoji] || unmarshalMsgFieldSkip6zoji[nextMiss6zoji]) {
				nextMiss6zoji++
			}
			if nextMiss6zoji == maxFields6zoji {
				// filled all the empty fields!
				break doneWithStruct6zoji
			}
			missingFieldsLeft6zoji--
			curField6zoji = unmarshalMsgFieldOrder6zoji[nextMiss6zoji]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zoji)
		switch curField6zoji {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found6zoji[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found6zoji[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found6zoji[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zoji[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zgpf uint32
				zgpf, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zgpf) {
					z.Structs = (z.Structs)[:zgpf]
				} else {
					z.Structs = make([]Struct, zgpf)
				}
				for zvoe := range z.Structs {
					const maxFields7zock = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zock uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zock, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zock := totalEncodedFields7zock
					missingFieldsLeft7zock := maxFields7zock - totalEncodedFields7zock

					var nextMiss7zock int32 = -1
					var found7zock [maxFields7zock]bool
					var curField7zock string

				doneWithStruct7zock:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zock > 0 || missingFieldsLeft7zock > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zock, missingFieldsLeft7zock, msgp.ShowFound(found7zock[:]), unmarshalMsgFieldOrder7zock)
						if encodedFieldsLeft7zock > 0 {
							encodedFieldsLeft7zock--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zock = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zock < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zock = 0
							}
							for nextMiss7zock < maxFields7zock && (found7zock[nextMiss7zock] || unmarshalMsgFieldSkip7zock[nextMiss7zock]) {
								nextMiss7zock++
							}
							if nextMiss7zock == maxFields7zock {
								// filled all the empty fields!
								break doneWithStruct7zock
							}
							missingFieldsLeft7zock--
							curField7zock = unmarshalMsgFieldOrder7zock[nextMiss7zock]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zock)
						switch curField7zock {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zock[0] = true
							z.Structs[zvoe].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found7zock[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zvoe].Fields) = (z.Structs[zvoe].Fields)[:0]
							} else {

								var ztpy uint32
								ztpy, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zvoe].Fields) >= int(ztpy) {
									z.Structs[zvoe].Fields = (z.Structs[zvoe].Fields)[:ztpy]
								} else {
									z.Structs[zvoe].Fields = make([]Field, ztpy)
								}
								for zweb := range z.Structs[zvoe].Fields {
									bts, err = z.Structs[zvoe].Fields[zweb].UnmarshalMsg(bts)
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
					if nextMiss7zock != -1 {
						bts = nbs.PopAlwaysNil()
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
	if nextMiss6zoji != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder6zoji = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip6zoji = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder7zock = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zock = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zvoe := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zvoe].StructName) + 7 + msgp.ArrayHeaderSize
		for zweb := range z.Structs[zvoe].Fields {
			s += z.Structs[zvoe].Fields[zweb].Msgsize()
		}
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
	const maxFields8zmyd = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zmyd uint32
	totalEncodedFields8zmyd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zmyd := totalEncodedFields8zmyd
	missingFieldsLeft8zmyd := maxFields8zmyd - totalEncodedFields8zmyd

	var nextMiss8zmyd int32 = -1
	var found8zmyd [maxFields8zmyd]bool
	var curField8zmyd string

doneWithStruct8zmyd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zmyd > 0 || missingFieldsLeft8zmyd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zmyd, missingFieldsLeft8zmyd, msgp.ShowFound(found8zmyd[:]), decodeMsgFieldOrder8zmyd)
		if encodedFieldsLeft8zmyd > 0 {
			encodedFieldsLeft8zmyd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zmyd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zmyd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zmyd = 0
			}
			for nextMiss8zmyd < maxFields8zmyd && (found8zmyd[nextMiss8zmyd] || decodeMsgFieldSkip8zmyd[nextMiss8zmyd]) {
				nextMiss8zmyd++
			}
			if nextMiss8zmyd == maxFields8zmyd {
				// filled all the empty fields!
				break doneWithStruct8zmyd
			}
			missingFieldsLeft8zmyd--
			curField8zmyd = decodeMsgFieldOrder8zmyd[nextMiss8zmyd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zmyd)
		switch curField8zmyd {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found8zmyd[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found8zmyd[1] = true
			var zylg uint32
			zylg, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zylg) {
				z.Fields = (z.Fields)[:zylg]
			} else {
				z.Fields = make([]Field, zylg)
			}
			for zamf := range z.Fields {
				err = z.Fields[zamf].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss8zmyd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder8zmyd = []string{"StructName", "Fields"}

var decodeMsgFieldSkip8zmyd = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Struct) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *Struct) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "StructName"
	err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StructName)
	if err != nil {
		panic(err)
	}
	// write "Fields"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fields)))
	if err != nil {
		panic(err)
	}
	for zamf := range z.Fields {
		err = z.Fields[zamf].EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Struct) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "StructName"
	o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.StructName)
	// string "Fields"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fields)))
	for zamf := range z.Fields {
		o, err = z.Fields[zamf].MarshalMsg(o)
		if err != nil {
			panic(err)
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
	const maxFields9zyro = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zyro uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zyro, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zyro := totalEncodedFields9zyro
	missingFieldsLeft9zyro := maxFields9zyro - totalEncodedFields9zyro

	var nextMiss9zyro int32 = -1
	var found9zyro [maxFields9zyro]bool
	var curField9zyro string

doneWithStruct9zyro:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zyro > 0 || missingFieldsLeft9zyro > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zyro, missingFieldsLeft9zyro, msgp.ShowFound(found9zyro[:]), unmarshalMsgFieldOrder9zyro)
		if encodedFieldsLeft9zyro > 0 {
			encodedFieldsLeft9zyro--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zyro = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zyro < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zyro = 0
			}
			for nextMiss9zyro < maxFields9zyro && (found9zyro[nextMiss9zyro] || unmarshalMsgFieldSkip9zyro[nextMiss9zyro]) {
				nextMiss9zyro++
			}
			if nextMiss9zyro == maxFields9zyro {
				// filled all the empty fields!
				break doneWithStruct9zyro
			}
			missingFieldsLeft9zyro--
			curField9zyro = unmarshalMsgFieldOrder9zyro[nextMiss9zyro]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zyro)
		switch curField9zyro {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found9zyro[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found9zyro[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zgwt uint32
				zgwt, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zgwt) {
					z.Fields = (z.Fields)[:zgwt]
				} else {
					z.Fields = make([]Field, zgwt)
				}
				for zamf := range z.Fields {
					bts, err = z.Fields[zamf].UnmarshalMsg(bts)
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
	if nextMiss9zyro != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder9zyro = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip9zyro = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zamf := range z.Fields {
		s += z.Fields[zamf].Msgsize()
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
		var zybg uint64
		zybg, err = dc.ReadUint64()
		(*z) = Zkind(zybg)
	}
	if err != nil {
		panic(err)
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z Zkind) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint64(uint64(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zkind) MarshalMsg(b []byte) (o []byte, err error) {
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
		var zhog uint64
		zhog, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zhog)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
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
	const maxFields10zshg = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zshg uint32
	totalEncodedFields10zshg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zshg := totalEncodedFields10zshg
	missingFieldsLeft10zshg := maxFields10zshg - totalEncodedFields10zshg

	var nextMiss10zshg int32 = -1
	var found10zshg [maxFields10zshg]bool
	var curField10zshg string

doneWithStruct10zshg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zshg > 0 || missingFieldsLeft10zshg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zshg, missingFieldsLeft10zshg, msgp.ShowFound(found10zshg[:]), decodeMsgFieldOrder10zshg)
		if encodedFieldsLeft10zshg > 0 {
			encodedFieldsLeft10zshg--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zshg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zshg < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zshg = 0
			}
			for nextMiss10zshg < maxFields10zshg && (found10zshg[nextMiss10zshg] || decodeMsgFieldSkip10zshg[nextMiss10zshg]) {
				nextMiss10zshg++
			}
			if nextMiss10zshg == maxFields10zshg {
				// filled all the empty fields!
				break doneWithStruct10zshg
			}
			missingFieldsLeft10zshg--
			curField10zshg = decodeMsgFieldOrder10zshg[nextMiss10zshg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zshg)
		switch curField10zshg {
		// -- templateDecodeMsg ends here --

		case "Name":
			found10zshg[0] = true
			const maxFields11zslz = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields11zslz uint32
			totalEncodedFields11zslz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft11zslz := totalEncodedFields11zslz
			missingFieldsLeft11zslz := maxFields11zslz - totalEncodedFields11zslz

			var nextMiss11zslz int32 = -1
			var found11zslz [maxFields11zslz]bool
			var curField11zslz string

		doneWithStruct11zslz:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zslz > 0 || missingFieldsLeft11zslz > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zslz, missingFieldsLeft11zslz, msgp.ShowFound(found11zslz[:]), decodeMsgFieldOrder11zslz)
				if encodedFieldsLeft11zslz > 0 {
					encodedFieldsLeft11zslz--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField11zslz = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zslz < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss11zslz = 0
					}
					for nextMiss11zslz < maxFields11zslz && (found11zslz[nextMiss11zslz] || decodeMsgFieldSkip11zslz[nextMiss11zslz]) {
						nextMiss11zslz++
					}
					if nextMiss11zslz == maxFields11zslz {
						// filled all the empty fields!
						break doneWithStruct11zslz
					}
					missingFieldsLeft11zslz--
					curField11zslz = decodeMsgFieldOrder11zslz[nextMiss11zslz]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zslz)
				switch curField11zslz {
				// -- templateDecodeMsg ends here --

				case "Kind":
					found11zslz[0] = true
					{
						var zefm uint64
						zefm, err = dc.ReadUint64()
						z.Name.Kind = Zkind(zefm)
					}
					if err != nil {
						panic(err)
					}
				case "Str":
					found11zslz[1] = true
					z.Name.Str, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				default:
					err = dc.Skip()
					if err != nil {
						panic(err)
					}
				}
			}
			if nextMiss11zslz != -1 {
				dc.PopAlwaysNil()
			}

		case "Domain":
			found10zshg[1] = true
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
					panic(err)
				}
			}
		case "Range":
			found10zshg[2] = true
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
					panic(err)
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss10zshg != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder10zshg = []string{"Name", "Domain", "Range"}

var decodeMsgFieldSkip10zshg = []bool{false, false, false}

// fields of KS
var decodeMsgFieldOrder11zslz = []string{"Kind", "Str"}

var decodeMsgFieldSkip11zslz = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Ztype) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[1] = (z.Domain == nil) // pointer, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Range == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Ztype) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zugl [3]bool
	fieldsInUse_zejz := z.fieldsNotEmpty(empty_zugl[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zejz)
	if err != nil {
		return err
	}

	// write "Name"
	// map header, size 2
	// write "Kind"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Name.Kind))
	if err != nil {
		panic(err)
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name.Str)
	if err != nil {
		panic(err)
	}
	if !empty_zugl[1] {
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
				panic(err)
			}
		}
	}

	if !empty_zugl[2] {
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
				panic(err)
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Ztype) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Name"
	// map header, size 2
	// string "Kind"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Name.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Name.Str)
	if !empty[1] {
		// string "Domain"
		o = append(o, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
		if z.Domain == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Domain.MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty[2] {
		// string "Range"
		o = append(o, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65)
		if z.Range == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Range.MarshalMsg(o)
			if err != nil {
				panic(err)
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
	const maxFields12znee = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields12znee uint32
	if !nbs.AlwaysNil {
		totalEncodedFields12znee, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft12znee := totalEncodedFields12znee
	missingFieldsLeft12znee := maxFields12znee - totalEncodedFields12znee

	var nextMiss12znee int32 = -1
	var found12znee [maxFields12znee]bool
	var curField12znee string

doneWithStruct12znee:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12znee > 0 || missingFieldsLeft12znee > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12znee, missingFieldsLeft12znee, msgp.ShowFound(found12znee[:]), unmarshalMsgFieldOrder12znee)
		if encodedFieldsLeft12znee > 0 {
			encodedFieldsLeft12znee--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField12znee = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12znee < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss12znee = 0
			}
			for nextMiss12znee < maxFields12znee && (found12znee[nextMiss12znee] || unmarshalMsgFieldSkip12znee[nextMiss12znee]) {
				nextMiss12znee++
			}
			if nextMiss12znee == maxFields12znee {
				// filled all the empty fields!
				break doneWithStruct12znee
			}
			missingFieldsLeft12znee--
			curField12znee = unmarshalMsgFieldOrder12znee[nextMiss12znee]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12znee)
		switch curField12znee {
		// -- templateUnmarshalMsg ends here --

		case "Name":
			found12znee[0] = true
			const maxFields13zxfi = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields13zxfi uint32
			if !nbs.AlwaysNil {
				totalEncodedFields13zxfi, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft13zxfi := totalEncodedFields13zxfi
			missingFieldsLeft13zxfi := maxFields13zxfi - totalEncodedFields13zxfi

			var nextMiss13zxfi int32 = -1
			var found13zxfi [maxFields13zxfi]bool
			var curField13zxfi string

		doneWithStruct13zxfi:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft13zxfi > 0 || missingFieldsLeft13zxfi > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zxfi, missingFieldsLeft13zxfi, msgp.ShowFound(found13zxfi[:]), unmarshalMsgFieldOrder13zxfi)
				if encodedFieldsLeft13zxfi > 0 {
					encodedFieldsLeft13zxfi--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField13zxfi = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss13zxfi < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss13zxfi = 0
					}
					for nextMiss13zxfi < maxFields13zxfi && (found13zxfi[nextMiss13zxfi] || unmarshalMsgFieldSkip13zxfi[nextMiss13zxfi]) {
						nextMiss13zxfi++
					}
					if nextMiss13zxfi == maxFields13zxfi {
						// filled all the empty fields!
						break doneWithStruct13zxfi
					}
					missingFieldsLeft13zxfi--
					curField13zxfi = unmarshalMsgFieldOrder13zxfi[nextMiss13zxfi]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField13zxfi)
				switch curField13zxfi {
				// -- templateUnmarshalMsg ends here --

				case "Kind":
					found13zxfi[0] = true
					{
						var ztak uint64
						ztak, bts, err = nbs.ReadUint64Bytes(bts)

						if err != nil {
							panic(err)
						}
						z.Name.Kind = Zkind(ztak)
					}
				case "Str":
					found13zxfi[1] = true
					z.Name.Str, bts, err = nbs.ReadStringBytes(bts)

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
			if nextMiss13zxfi != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Domain":
			found12znee[1] = true
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
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Range":
			found12znee[2] = true
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
	if nextMiss12znee != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder12znee = []string{"Name", "Domain", "Range"}

var unmarshalMsgFieldSkip12znee = []bool{false, false, false}

// fields of KS
var unmarshalMsgFieldOrder13zxfi = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip13zxfi = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Ztype) Msgsize() (s int) {
	s = 1 + 5 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Name.Str) + 7
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
