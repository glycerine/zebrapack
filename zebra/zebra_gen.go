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
	const maxFields0zaei = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zaei uint32
	totalEncodedFields0zaei, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zaei := totalEncodedFields0zaei
	missingFieldsLeft0zaei := maxFields0zaei - totalEncodedFields0zaei

	var nextMiss0zaei int32 = -1
	var found0zaei [maxFields0zaei]bool
	var curField0zaei string

doneWithStruct0zaei:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaei > 0 || missingFieldsLeft0zaei > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaei, missingFieldsLeft0zaei, msgp.ShowFound(found0zaei[:]), decodeMsgFieldOrder0zaei)
		if encodedFieldsLeft0zaei > 0 {
			encodedFieldsLeft0zaei--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zaei = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zaei < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zaei = 0
			}
			for nextMiss0zaei < maxFields0zaei && (found0zaei[nextMiss0zaei] || decodeMsgFieldSkip0zaei[nextMiss0zaei]) {
				nextMiss0zaei++
			}
			if nextMiss0zaei == maxFields0zaei {
				// filled all the empty fields!
				break doneWithStruct0zaei
			}
			missingFieldsLeft0zaei--
			curField0zaei = decodeMsgFieldOrder0zaei[nextMiss0zaei]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaei)
		switch curField0zaei {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zaei[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zaei[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zaei[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zaei[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zaei[4] = true
			{
				var zyrp uint64
				zyrp, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zyrp)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zaei[5] = true
			{
				var zsog uint64
				zsog, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zsog)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zaei[6] = true
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
			found0zaei[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zaei[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zaei[9] = true
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
	if nextMiss0zaei != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zaei = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zaei = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zwlj [10]bool
	fieldsInUse_zomz := z.fieldsNotEmpty(empty_zwlj[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zomz)
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
	if !empty_zwlj[7] {
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

	if !empty_zwlj[8] {
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

	if !empty_zwlj[9] {
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
	const maxFields1zrgz = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zrgz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zrgz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zrgz := totalEncodedFields1zrgz
	missingFieldsLeft1zrgz := maxFields1zrgz - totalEncodedFields1zrgz

	var nextMiss1zrgz int32 = -1
	var found1zrgz [maxFields1zrgz]bool
	var curField1zrgz string

doneWithStruct1zrgz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zrgz > 0 || missingFieldsLeft1zrgz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zrgz, missingFieldsLeft1zrgz, msgp.ShowFound(found1zrgz[:]), unmarshalMsgFieldOrder1zrgz)
		if encodedFieldsLeft1zrgz > 0 {
			encodedFieldsLeft1zrgz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zrgz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zrgz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zrgz = 0
			}
			for nextMiss1zrgz < maxFields1zrgz && (found1zrgz[nextMiss1zrgz] || unmarshalMsgFieldSkip1zrgz[nextMiss1zrgz]) {
				nextMiss1zrgz++
			}
			if nextMiss1zrgz == maxFields1zrgz {
				// filled all the empty fields!
				break doneWithStruct1zrgz
			}
			missingFieldsLeft1zrgz--
			curField1zrgz = unmarshalMsgFieldOrder1zrgz[nextMiss1zrgz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zrgz)
		switch curField1zrgz {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zrgz[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zrgz[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zrgz[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zrgz[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zrgz[4] = true
			{
				var znch uint64
				znch, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(znch)
			}
		case "FieldPrimitive":
			found1zrgz[5] = true
			{
				var zspy uint64
				zspy, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zspy)
			}
		case "FieldFullType":
			found1zrgz[6] = true
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
			found1zrgz[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zrgz[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zrgz[9] = true
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
	if nextMiss1zrgz != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zrgz = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zrgz = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2ztch = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2ztch uint32
	totalEncodedFields2ztch, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2ztch := totalEncodedFields2ztch
	missingFieldsLeft2ztch := maxFields2ztch - totalEncodedFields2ztch

	var nextMiss2ztch int32 = -1
	var found2ztch [maxFields2ztch]bool
	var curField2ztch string

doneWithStruct2ztch:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2ztch > 0 || missingFieldsLeft2ztch > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2ztch, missingFieldsLeft2ztch, msgp.ShowFound(found2ztch[:]), decodeMsgFieldOrder2ztch)
		if encodedFieldsLeft2ztch > 0 {
			encodedFieldsLeft2ztch--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2ztch = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2ztch < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2ztch = 0
			}
			for nextMiss2ztch < maxFields2ztch && (found2ztch[nextMiss2ztch] || decodeMsgFieldSkip2ztch[nextMiss2ztch]) {
				nextMiss2ztch++
			}
			if nextMiss2ztch == maxFields2ztch {
				// filled all the empty fields!
				break doneWithStruct2ztch
			}
			missingFieldsLeft2ztch--
			curField2ztch = decodeMsgFieldOrder2ztch[nextMiss2ztch]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2ztch)
		switch curField2ztch {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found2ztch[0] = true
			{
				var zwip uint64
				zwip, err = dc.ReadUint64()
				z.Kind = Zkind(zwip)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found2ztch[1] = true
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
	if nextMiss2ztch != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of KS
var decodeMsgFieldOrder2ztch = []string{"Kind", "Str"}

var decodeMsgFieldSkip2ztch = []bool{false, false}

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
	const maxFields3zzcv = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zzcv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zzcv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zzcv := totalEncodedFields3zzcv
	missingFieldsLeft3zzcv := maxFields3zzcv - totalEncodedFields3zzcv

	var nextMiss3zzcv int32 = -1
	var found3zzcv [maxFields3zzcv]bool
	var curField3zzcv string

doneWithStruct3zzcv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zzcv > 0 || missingFieldsLeft3zzcv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zzcv, missingFieldsLeft3zzcv, msgp.ShowFound(found3zzcv[:]), unmarshalMsgFieldOrder3zzcv)
		if encodedFieldsLeft3zzcv > 0 {
			encodedFieldsLeft3zzcv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zzcv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zzcv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zzcv = 0
			}
			for nextMiss3zzcv < maxFields3zzcv && (found3zzcv[nextMiss3zzcv] || unmarshalMsgFieldSkip3zzcv[nextMiss3zzcv]) {
				nextMiss3zzcv++
			}
			if nextMiss3zzcv == maxFields3zzcv {
				// filled all the empty fields!
				break doneWithStruct3zzcv
			}
			missingFieldsLeft3zzcv--
			curField3zzcv = unmarshalMsgFieldOrder3zzcv[nextMiss3zzcv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zzcv)
		switch curField3zzcv {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found3zzcv[0] = true
			{
				var zesu uint64
				zesu, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zesu)
			}
		case "Str":
			found3zzcv[1] = true
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
	if nextMiss3zzcv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of KS
var unmarshalMsgFieldOrder3zzcv = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip3zzcv = []bool{false, false}

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
	const maxFields4zhbr = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zhbr uint32
	totalEncodedFields4zhbr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zhbr := totalEncodedFields4zhbr
	missingFieldsLeft4zhbr := maxFields4zhbr - totalEncodedFields4zhbr

	var nextMiss4zhbr int32 = -1
	var found4zhbr [maxFields4zhbr]bool
	var curField4zhbr string

doneWithStruct4zhbr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zhbr > 0 || missingFieldsLeft4zhbr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zhbr, missingFieldsLeft4zhbr, msgp.ShowFound(found4zhbr[:]), decodeMsgFieldOrder4zhbr)
		if encodedFieldsLeft4zhbr > 0 {
			encodedFieldsLeft4zhbr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zhbr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zhbr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zhbr = 0
			}
			for nextMiss4zhbr < maxFields4zhbr && (found4zhbr[nextMiss4zhbr] || decodeMsgFieldSkip4zhbr[nextMiss4zhbr]) {
				nextMiss4zhbr++
			}
			if nextMiss4zhbr == maxFields4zhbr {
				// filled all the empty fields!
				break doneWithStruct4zhbr
			}
			missingFieldsLeft4zhbr--
			curField4zhbr = decodeMsgFieldOrder4zhbr[nextMiss4zhbr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zhbr)
		switch curField4zhbr {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found4zhbr[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zhbr[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zhbr[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zhbr[3] = true
			var znnw uint32
			znnw, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(znnw) {
				z.Structs = (z.Structs)[:znnw]
			} else {
				z.Structs = make([]Struct, znnw)
			}
			for zdnp := range z.Structs {
				const maxFields5zuqf = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zuqf uint32
				totalEncodedFields5zuqf, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zuqf := totalEncodedFields5zuqf
				missingFieldsLeft5zuqf := maxFields5zuqf - totalEncodedFields5zuqf

				var nextMiss5zuqf int32 = -1
				var found5zuqf [maxFields5zuqf]bool
				var curField5zuqf string

			doneWithStruct5zuqf:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zuqf > 0 || missingFieldsLeft5zuqf > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zuqf, missingFieldsLeft5zuqf, msgp.ShowFound(found5zuqf[:]), decodeMsgFieldOrder5zuqf)
					if encodedFieldsLeft5zuqf > 0 {
						encodedFieldsLeft5zuqf--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zuqf = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zuqf < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zuqf = 0
						}
						for nextMiss5zuqf < maxFields5zuqf && (found5zuqf[nextMiss5zuqf] || decodeMsgFieldSkip5zuqf[nextMiss5zuqf]) {
							nextMiss5zuqf++
						}
						if nextMiss5zuqf == maxFields5zuqf {
							// filled all the empty fields!
							break doneWithStruct5zuqf
						}
						missingFieldsLeft5zuqf--
						curField5zuqf = decodeMsgFieldOrder5zuqf[nextMiss5zuqf]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zuqf)
					switch curField5zuqf {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zuqf[0] = true
						z.Structs[zdnp].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found5zuqf[1] = true
						var zxxl uint32
						zxxl, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zdnp].Fields) >= int(zxxl) {
							z.Structs[zdnp].Fields = (z.Structs[zdnp].Fields)[:zxxl]
						} else {
							z.Structs[zdnp].Fields = make([]Field, zxxl)
						}
						for zefr := range z.Structs[zdnp].Fields {
							err = z.Structs[zdnp].Fields[zefr].DecodeMsg(dc)
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
				if nextMiss5zuqf != -1 {
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
	if nextMiss4zhbr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder4zhbr = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip4zhbr = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder5zuqf = []string{"StructName", "Fields"}

var decodeMsgFieldSkip5zuqf = []bool{false, false}

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
	for zdnp := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zdnp].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zdnp].Fields)))
		if err != nil {
			panic(err)
		}
		for zefr := range z.Structs[zdnp].Fields {
			err = z.Structs[zdnp].Fields[zefr].EncodeMsg(en)
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
	for zdnp := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zdnp].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zdnp].Fields)))
		for zefr := range z.Structs[zdnp].Fields {
			o, err = z.Structs[zdnp].Fields[zefr].MarshalMsg(o)
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
	const maxFields6zxmb = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zxmb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zxmb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zxmb := totalEncodedFields6zxmb
	missingFieldsLeft6zxmb := maxFields6zxmb - totalEncodedFields6zxmb

	var nextMiss6zxmb int32 = -1
	var found6zxmb [maxFields6zxmb]bool
	var curField6zxmb string

doneWithStruct6zxmb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zxmb > 0 || missingFieldsLeft6zxmb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zxmb, missingFieldsLeft6zxmb, msgp.ShowFound(found6zxmb[:]), unmarshalMsgFieldOrder6zxmb)
		if encodedFieldsLeft6zxmb > 0 {
			encodedFieldsLeft6zxmb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zxmb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zxmb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zxmb = 0
			}
			for nextMiss6zxmb < maxFields6zxmb && (found6zxmb[nextMiss6zxmb] || unmarshalMsgFieldSkip6zxmb[nextMiss6zxmb]) {
				nextMiss6zxmb++
			}
			if nextMiss6zxmb == maxFields6zxmb {
				// filled all the empty fields!
				break doneWithStruct6zxmb
			}
			missingFieldsLeft6zxmb--
			curField6zxmb = unmarshalMsgFieldOrder6zxmb[nextMiss6zxmb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zxmb)
		switch curField6zxmb {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found6zxmb[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found6zxmb[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found6zxmb[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zxmb[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zhwb uint32
				zhwb, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zhwb) {
					z.Structs = (z.Structs)[:zhwb]
				} else {
					z.Structs = make([]Struct, zhwb)
				}
				for zdnp := range z.Structs {
					const maxFields7zgpo = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zgpo uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zgpo, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zgpo := totalEncodedFields7zgpo
					missingFieldsLeft7zgpo := maxFields7zgpo - totalEncodedFields7zgpo

					var nextMiss7zgpo int32 = -1
					var found7zgpo [maxFields7zgpo]bool
					var curField7zgpo string

				doneWithStruct7zgpo:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zgpo > 0 || missingFieldsLeft7zgpo > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgpo, missingFieldsLeft7zgpo, msgp.ShowFound(found7zgpo[:]), unmarshalMsgFieldOrder7zgpo)
						if encodedFieldsLeft7zgpo > 0 {
							encodedFieldsLeft7zgpo--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zgpo = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zgpo < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zgpo = 0
							}
							for nextMiss7zgpo < maxFields7zgpo && (found7zgpo[nextMiss7zgpo] || unmarshalMsgFieldSkip7zgpo[nextMiss7zgpo]) {
								nextMiss7zgpo++
							}
							if nextMiss7zgpo == maxFields7zgpo {
								// filled all the empty fields!
								break doneWithStruct7zgpo
							}
							missingFieldsLeft7zgpo--
							curField7zgpo = unmarshalMsgFieldOrder7zgpo[nextMiss7zgpo]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zgpo)
						switch curField7zgpo {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zgpo[0] = true
							z.Structs[zdnp].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found7zgpo[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zdnp].Fields) = (z.Structs[zdnp].Fields)[:0]
							} else {

								var zpmd uint32
								zpmd, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zdnp].Fields) >= int(zpmd) {
									z.Structs[zdnp].Fields = (z.Structs[zdnp].Fields)[:zpmd]
								} else {
									z.Structs[zdnp].Fields = make([]Field, zpmd)
								}
								for zefr := range z.Structs[zdnp].Fields {
									bts, err = z.Structs[zdnp].Fields[zefr].UnmarshalMsg(bts)
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
					if nextMiss7zgpo != -1 {
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
	if nextMiss6zxmb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder6zxmb = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip6zxmb = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder7zgpo = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zgpo = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zdnp := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zdnp].StructName) + 7 + msgp.ArrayHeaderSize
		for zefr := range z.Structs[zdnp].Fields {
			s += z.Structs[zdnp].Fields[zefr].Msgsize()
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
	const maxFields8zfxk = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zfxk uint32
	totalEncodedFields8zfxk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zfxk := totalEncodedFields8zfxk
	missingFieldsLeft8zfxk := maxFields8zfxk - totalEncodedFields8zfxk

	var nextMiss8zfxk int32 = -1
	var found8zfxk [maxFields8zfxk]bool
	var curField8zfxk string

doneWithStruct8zfxk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zfxk > 0 || missingFieldsLeft8zfxk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zfxk, missingFieldsLeft8zfxk, msgp.ShowFound(found8zfxk[:]), decodeMsgFieldOrder8zfxk)
		if encodedFieldsLeft8zfxk > 0 {
			encodedFieldsLeft8zfxk--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zfxk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zfxk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zfxk = 0
			}
			for nextMiss8zfxk < maxFields8zfxk && (found8zfxk[nextMiss8zfxk] || decodeMsgFieldSkip8zfxk[nextMiss8zfxk]) {
				nextMiss8zfxk++
			}
			if nextMiss8zfxk == maxFields8zfxk {
				// filled all the empty fields!
				break doneWithStruct8zfxk
			}
			missingFieldsLeft8zfxk--
			curField8zfxk = decodeMsgFieldOrder8zfxk[nextMiss8zfxk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zfxk)
		switch curField8zfxk {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found8zfxk[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found8zfxk[1] = true
			var zckk uint32
			zckk, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zckk) {
				z.Fields = (z.Fields)[:zckk]
			} else {
				z.Fields = make([]Field, zckk)
			}
			for zdkm := range z.Fields {
				err = z.Fields[zdkm].DecodeMsg(dc)
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
	if nextMiss8zfxk != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder8zfxk = []string{"StructName", "Fields"}

var decodeMsgFieldSkip8zfxk = []bool{false, false}

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
	for zdkm := range z.Fields {
		err = z.Fields[zdkm].EncodeMsg(en)
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
	for zdkm := range z.Fields {
		o, err = z.Fields[zdkm].MarshalMsg(o)
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
	const maxFields9zaus = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zaus uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zaus, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zaus := totalEncodedFields9zaus
	missingFieldsLeft9zaus := maxFields9zaus - totalEncodedFields9zaus

	var nextMiss9zaus int32 = -1
	var found9zaus [maxFields9zaus]bool
	var curField9zaus string

doneWithStruct9zaus:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zaus > 0 || missingFieldsLeft9zaus > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zaus, missingFieldsLeft9zaus, msgp.ShowFound(found9zaus[:]), unmarshalMsgFieldOrder9zaus)
		if encodedFieldsLeft9zaus > 0 {
			encodedFieldsLeft9zaus--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zaus = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zaus < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zaus = 0
			}
			for nextMiss9zaus < maxFields9zaus && (found9zaus[nextMiss9zaus] || unmarshalMsgFieldSkip9zaus[nextMiss9zaus]) {
				nextMiss9zaus++
			}
			if nextMiss9zaus == maxFields9zaus {
				// filled all the empty fields!
				break doneWithStruct9zaus
			}
			missingFieldsLeft9zaus--
			curField9zaus = unmarshalMsgFieldOrder9zaus[nextMiss9zaus]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zaus)
		switch curField9zaus {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found9zaus[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found9zaus[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zdqa uint32
				zdqa, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zdqa) {
					z.Fields = (z.Fields)[:zdqa]
				} else {
					z.Fields = make([]Field, zdqa)
				}
				for zdkm := range z.Fields {
					bts, err = z.Fields[zdkm].UnmarshalMsg(bts)
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
	if nextMiss9zaus != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder9zaus = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip9zaus = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zdkm := range z.Fields {
		s += z.Fields[zdkm].Msgsize()
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
		var zqyt uint64
		zqyt, err = dc.ReadUint64()
		(*z) = Zkind(zqyt)
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
		var ziig uint64
		ziig, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(ziig)
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
	const maxFields10zvjc = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zvjc uint32
	totalEncodedFields10zvjc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zvjc := totalEncodedFields10zvjc
	missingFieldsLeft10zvjc := maxFields10zvjc - totalEncodedFields10zvjc

	var nextMiss10zvjc int32 = -1
	var found10zvjc [maxFields10zvjc]bool
	var curField10zvjc string

doneWithStruct10zvjc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zvjc > 0 || missingFieldsLeft10zvjc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zvjc, missingFieldsLeft10zvjc, msgp.ShowFound(found10zvjc[:]), decodeMsgFieldOrder10zvjc)
		if encodedFieldsLeft10zvjc > 0 {
			encodedFieldsLeft10zvjc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zvjc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zvjc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zvjc = 0
			}
			for nextMiss10zvjc < maxFields10zvjc && (found10zvjc[nextMiss10zvjc] || decodeMsgFieldSkip10zvjc[nextMiss10zvjc]) {
				nextMiss10zvjc++
			}
			if nextMiss10zvjc == maxFields10zvjc {
				// filled all the empty fields!
				break doneWithStruct10zvjc
			}
			missingFieldsLeft10zvjc--
			curField10zvjc = decodeMsgFieldOrder10zvjc[nextMiss10zvjc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zvjc)
		switch curField10zvjc {
		// -- templateDecodeMsg ends here --

		case "Name":
			found10zvjc[0] = true
			const maxFields11zvxd = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields11zvxd uint32
			totalEncodedFields11zvxd, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft11zvxd := totalEncodedFields11zvxd
			missingFieldsLeft11zvxd := maxFields11zvxd - totalEncodedFields11zvxd

			var nextMiss11zvxd int32 = -1
			var found11zvxd [maxFields11zvxd]bool
			var curField11zvxd string

		doneWithStruct11zvxd:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zvxd > 0 || missingFieldsLeft11zvxd > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zvxd, missingFieldsLeft11zvxd, msgp.ShowFound(found11zvxd[:]), decodeMsgFieldOrder11zvxd)
				if encodedFieldsLeft11zvxd > 0 {
					encodedFieldsLeft11zvxd--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField11zvxd = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zvxd < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss11zvxd = 0
					}
					for nextMiss11zvxd < maxFields11zvxd && (found11zvxd[nextMiss11zvxd] || decodeMsgFieldSkip11zvxd[nextMiss11zvxd]) {
						nextMiss11zvxd++
					}
					if nextMiss11zvxd == maxFields11zvxd {
						// filled all the empty fields!
						break doneWithStruct11zvxd
					}
					missingFieldsLeft11zvxd--
					curField11zvxd = decodeMsgFieldOrder11zvxd[nextMiss11zvxd]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zvxd)
				switch curField11zvxd {
				// -- templateDecodeMsg ends here --

				case "Kind":
					found11zvxd[0] = true
					{
						var ziga uint64
						ziga, err = dc.ReadUint64()
						z.Name.Kind = Zkind(ziga)
					}
					if err != nil {
						panic(err)
					}
				case "Str":
					found11zvxd[1] = true
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
			if nextMiss11zvxd != -1 {
				dc.PopAlwaysNil()
			}

		case "Domain":
			found10zvjc[1] = true
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
			found10zvjc[2] = true
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
	if nextMiss10zvjc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder10zvjc = []string{"Name", "Domain", "Range"}

var decodeMsgFieldSkip10zvjc = []bool{false, false, false}

// fields of KS
var decodeMsgFieldOrder11zvxd = []string{"Kind", "Str"}

var decodeMsgFieldSkip11zvxd = []bool{false, false}

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
	var empty_zadv [3]bool
	fieldsInUse_zxyk := z.fieldsNotEmpty(empty_zadv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxyk)
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
	if !empty_zadv[1] {
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

	if !empty_zadv[2] {
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
	const maxFields12zlnc = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields12zlnc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields12zlnc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft12zlnc := totalEncodedFields12zlnc
	missingFieldsLeft12zlnc := maxFields12zlnc - totalEncodedFields12zlnc

	var nextMiss12zlnc int32 = -1
	var found12zlnc [maxFields12zlnc]bool
	var curField12zlnc string

doneWithStruct12zlnc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zlnc > 0 || missingFieldsLeft12zlnc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zlnc, missingFieldsLeft12zlnc, msgp.ShowFound(found12zlnc[:]), unmarshalMsgFieldOrder12zlnc)
		if encodedFieldsLeft12zlnc > 0 {
			encodedFieldsLeft12zlnc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField12zlnc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zlnc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss12zlnc = 0
			}
			for nextMiss12zlnc < maxFields12zlnc && (found12zlnc[nextMiss12zlnc] || unmarshalMsgFieldSkip12zlnc[nextMiss12zlnc]) {
				nextMiss12zlnc++
			}
			if nextMiss12zlnc == maxFields12zlnc {
				// filled all the empty fields!
				break doneWithStruct12zlnc
			}
			missingFieldsLeft12zlnc--
			curField12zlnc = unmarshalMsgFieldOrder12zlnc[nextMiss12zlnc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zlnc)
		switch curField12zlnc {
		// -- templateUnmarshalMsg ends here --

		case "Name":
			found12zlnc[0] = true
			const maxFields13zlxb = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields13zlxb uint32
			if !nbs.AlwaysNil {
				totalEncodedFields13zlxb, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft13zlxb := totalEncodedFields13zlxb
			missingFieldsLeft13zlxb := maxFields13zlxb - totalEncodedFields13zlxb

			var nextMiss13zlxb int32 = -1
			var found13zlxb [maxFields13zlxb]bool
			var curField13zlxb string

		doneWithStruct13zlxb:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft13zlxb > 0 || missingFieldsLeft13zlxb > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zlxb, missingFieldsLeft13zlxb, msgp.ShowFound(found13zlxb[:]), unmarshalMsgFieldOrder13zlxb)
				if encodedFieldsLeft13zlxb > 0 {
					encodedFieldsLeft13zlxb--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField13zlxb = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss13zlxb < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss13zlxb = 0
					}
					for nextMiss13zlxb < maxFields13zlxb && (found13zlxb[nextMiss13zlxb] || unmarshalMsgFieldSkip13zlxb[nextMiss13zlxb]) {
						nextMiss13zlxb++
					}
					if nextMiss13zlxb == maxFields13zlxb {
						// filled all the empty fields!
						break doneWithStruct13zlxb
					}
					missingFieldsLeft13zlxb--
					curField13zlxb = unmarshalMsgFieldOrder13zlxb[nextMiss13zlxb]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField13zlxb)
				switch curField13zlxb {
				// -- templateUnmarshalMsg ends here --

				case "Kind":
					found13zlxb[0] = true
					{
						var zqcy uint64
						zqcy, bts, err = nbs.ReadUint64Bytes(bts)

						if err != nil {
							panic(err)
						}
						z.Name.Kind = Zkind(zqcy)
					}
				case "Str":
					found13zlxb[1] = true
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
			if nextMiss13zlxb != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Domain":
			found12zlnc[1] = true
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
			found12zlnc[2] = true
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
	if nextMiss12zlnc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder12zlnc = []string{"Name", "Domain", "Range"}

var unmarshalMsgFieldSkip12zlnc = []bool{false, false, false}

// fields of KS
var unmarshalMsgFieldOrder13zlxb = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip13zlxb = []bool{false, false}

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
