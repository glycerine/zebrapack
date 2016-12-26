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
	const maxFields0zvsb = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zvsb uint32
	totalEncodedFields0zvsb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zvsb := totalEncodedFields0zvsb
	missingFieldsLeft0zvsb := maxFields0zvsb - totalEncodedFields0zvsb

	var nextMiss0zvsb int32 = -1
	var found0zvsb [maxFields0zvsb]bool
	var curField0zvsb string

doneWithStruct0zvsb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zvsb > 0 || missingFieldsLeft0zvsb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zvsb, missingFieldsLeft0zvsb, msgp.ShowFound(found0zvsb[:]), decodeMsgFieldOrder0zvsb)
		if encodedFieldsLeft0zvsb > 0 {
			encodedFieldsLeft0zvsb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zvsb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zvsb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zvsb = 0
			}
			for nextMiss0zvsb < maxFields0zvsb && (found0zvsb[nextMiss0zvsb] || decodeMsgFieldSkip0zvsb[nextMiss0zvsb]) {
				nextMiss0zvsb++
			}
			if nextMiss0zvsb == maxFields0zvsb {
				// filled all the empty fields!
				break doneWithStruct0zvsb
			}
			missingFieldsLeft0zvsb--
			curField0zvsb = decodeMsgFieldOrder0zvsb[nextMiss0zvsb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zvsb)
		switch curField0zvsb {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zvsb[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zvsb[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zvsb[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zvsb[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zvsb[4] = true
			{
				var zvsh uint64
				zvsh, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zvsh)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zvsb[5] = true
			{
				var zztr uint64
				zztr, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zztr)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zvsb[6] = true
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
			found0zvsb[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zvsb[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zvsb[9] = true
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
	if nextMiss0zvsb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zvsb = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zvsb = []bool{false, false, false, false, false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Field) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 10
	}
	var fieldsInUse uint32 = 10
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

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Field) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zyhm [10]bool
	fieldsInUse_zlki := z.fieldsNotEmpty(empty_zyhm[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlki)
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
	if !empty_zyhm[3] {
		// write "FieldTypeStr"
		err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.FieldTypeStr)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zyhm[4] {
		// write "FieldCategory"
		err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.FieldCategory))
		if err != nil {
			panic(err)
		}
	}

	if !empty_zyhm[5] {
		// write "FieldPrimitive"
		err = en.Append(0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.FieldPrimitive))
		if err != nil {
			panic(err)
		}
	}

	if !empty_zyhm[6] {
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
	}

	if !empty_zyhm[7] {
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

	if !empty_zyhm[8] {
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

	if !empty_zyhm[9] {
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
				panic(err)
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
	const maxFields1zbjk = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zbjk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zbjk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zbjk := totalEncodedFields1zbjk
	missingFieldsLeft1zbjk := maxFields1zbjk - totalEncodedFields1zbjk

	var nextMiss1zbjk int32 = -1
	var found1zbjk [maxFields1zbjk]bool
	var curField1zbjk string

doneWithStruct1zbjk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zbjk > 0 || missingFieldsLeft1zbjk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zbjk, missingFieldsLeft1zbjk, msgp.ShowFound(found1zbjk[:]), unmarshalMsgFieldOrder1zbjk)
		if encodedFieldsLeft1zbjk > 0 {
			encodedFieldsLeft1zbjk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zbjk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zbjk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zbjk = 0
			}
			for nextMiss1zbjk < maxFields1zbjk && (found1zbjk[nextMiss1zbjk] || unmarshalMsgFieldSkip1zbjk[nextMiss1zbjk]) {
				nextMiss1zbjk++
			}
			if nextMiss1zbjk == maxFields1zbjk {
				// filled all the empty fields!
				break doneWithStruct1zbjk
			}
			missingFieldsLeft1zbjk--
			curField1zbjk = unmarshalMsgFieldOrder1zbjk[nextMiss1zbjk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zbjk)
		switch curField1zbjk {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zbjk[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zbjk[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zbjk[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zbjk[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zbjk[4] = true
			{
				var ziuf uint64
				ziuf, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(ziuf)
			}
		case "FieldPrimitive":
			found1zbjk[5] = true
			{
				var zaod uint64
				zaod, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zaod)
			}
		case "FieldFullType":
			found1zbjk[6] = true
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
			found1zbjk[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zbjk[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zbjk[9] = true
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
	if nextMiss1zbjk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zbjk = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zbjk = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zxao = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zxao uint32
	totalEncodedFields2zxao, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zxao := totalEncodedFields2zxao
	missingFieldsLeft2zxao := maxFields2zxao - totalEncodedFields2zxao

	var nextMiss2zxao int32 = -1
	var found2zxao [maxFields2zxao]bool
	var curField2zxao string

doneWithStruct2zxao:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zxao > 0 || missingFieldsLeft2zxao > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zxao, missingFieldsLeft2zxao, msgp.ShowFound(found2zxao[:]), decodeMsgFieldOrder2zxao)
		if encodedFieldsLeft2zxao > 0 {
			encodedFieldsLeft2zxao--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zxao = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zxao < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zxao = 0
			}
			for nextMiss2zxao < maxFields2zxao && (found2zxao[nextMiss2zxao] || decodeMsgFieldSkip2zxao[nextMiss2zxao]) {
				nextMiss2zxao++
			}
			if nextMiss2zxao == maxFields2zxao {
				// filled all the empty fields!
				break doneWithStruct2zxao
			}
			missingFieldsLeft2zxao--
			curField2zxao = decodeMsgFieldOrder2zxao[nextMiss2zxao]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zxao)
		switch curField2zxao {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zxao[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zxao[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zxao[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zxao[3] = true
			var zgsn uint32
			zgsn, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zgsn) {
				z.Structs = (z.Structs)[:zgsn]
			} else {
				z.Structs = make([]Struct, zgsn)
			}
			for zbic := range z.Structs {
				const maxFields3zglu = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zglu uint32
				totalEncodedFields3zglu, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zglu := totalEncodedFields3zglu
				missingFieldsLeft3zglu := maxFields3zglu - totalEncodedFields3zglu

				var nextMiss3zglu int32 = -1
				var found3zglu [maxFields3zglu]bool
				var curField3zglu string

			doneWithStruct3zglu:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zglu > 0 || missingFieldsLeft3zglu > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zglu, missingFieldsLeft3zglu, msgp.ShowFound(found3zglu[:]), decodeMsgFieldOrder3zglu)
					if encodedFieldsLeft3zglu > 0 {
						encodedFieldsLeft3zglu--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zglu = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zglu < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zglu = 0
						}
						for nextMiss3zglu < maxFields3zglu && (found3zglu[nextMiss3zglu] || decodeMsgFieldSkip3zglu[nextMiss3zglu]) {
							nextMiss3zglu++
						}
						if nextMiss3zglu == maxFields3zglu {
							// filled all the empty fields!
							break doneWithStruct3zglu
						}
						missingFieldsLeft3zglu--
						curField3zglu = decodeMsgFieldOrder3zglu[nextMiss3zglu]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zglu)
					switch curField3zglu {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zglu[0] = true
						z.Structs[zbic].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zglu[1] = true
						var zewj uint32
						zewj, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zbic].Fields) >= int(zewj) {
							z.Structs[zbic].Fields = (z.Structs[zbic].Fields)[:zewj]
						} else {
							z.Structs[zbic].Fields = make([]Field, zewj)
						}
						for zlsq := range z.Structs[zbic].Fields {
							err = z.Structs[zbic].Fields[zlsq].DecodeMsg(dc)
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
				if nextMiss3zglu != -1 {
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
	if nextMiss2zxao != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zxao = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip2zxao = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zglu = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zglu = []bool{false, false}

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
	for zbic := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zbic].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zbic].Fields)))
		if err != nil {
			panic(err)
		}
		for zlsq := range z.Structs[zbic].Fields {
			err = z.Structs[zbic].Fields[zlsq].EncodeMsg(en)
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
	for zbic := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zbic].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zbic].Fields)))
		for zlsq := range z.Structs[zbic].Fields {
			o, err = z.Structs[zbic].Fields[zlsq].MarshalMsg(o)
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
	const maxFields4zudb = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zudb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zudb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zudb := totalEncodedFields4zudb
	missingFieldsLeft4zudb := maxFields4zudb - totalEncodedFields4zudb

	var nextMiss4zudb int32 = -1
	var found4zudb [maxFields4zudb]bool
	var curField4zudb string

doneWithStruct4zudb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zudb > 0 || missingFieldsLeft4zudb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zudb, missingFieldsLeft4zudb, msgp.ShowFound(found4zudb[:]), unmarshalMsgFieldOrder4zudb)
		if encodedFieldsLeft4zudb > 0 {
			encodedFieldsLeft4zudb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zudb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zudb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zudb = 0
			}
			for nextMiss4zudb < maxFields4zudb && (found4zudb[nextMiss4zudb] || unmarshalMsgFieldSkip4zudb[nextMiss4zudb]) {
				nextMiss4zudb++
			}
			if nextMiss4zudb == maxFields4zudb {
				// filled all the empty fields!
				break doneWithStruct4zudb
			}
			missingFieldsLeft4zudb--
			curField4zudb = unmarshalMsgFieldOrder4zudb[nextMiss4zudb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zudb)
		switch curField4zudb {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zudb[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zudb[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zudb[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zudb[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var znon uint32
				znon, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(znon) {
					z.Structs = (z.Structs)[:znon]
				} else {
					z.Structs = make([]Struct, znon)
				}
				for zbic := range z.Structs {
					const maxFields5zwcy = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zwcy uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zwcy, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zwcy := totalEncodedFields5zwcy
					missingFieldsLeft5zwcy := maxFields5zwcy - totalEncodedFields5zwcy

					var nextMiss5zwcy int32 = -1
					var found5zwcy [maxFields5zwcy]bool
					var curField5zwcy string

				doneWithStruct5zwcy:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zwcy > 0 || missingFieldsLeft5zwcy > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zwcy, missingFieldsLeft5zwcy, msgp.ShowFound(found5zwcy[:]), unmarshalMsgFieldOrder5zwcy)
						if encodedFieldsLeft5zwcy > 0 {
							encodedFieldsLeft5zwcy--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zwcy = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zwcy < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zwcy = 0
							}
							for nextMiss5zwcy < maxFields5zwcy && (found5zwcy[nextMiss5zwcy] || unmarshalMsgFieldSkip5zwcy[nextMiss5zwcy]) {
								nextMiss5zwcy++
							}
							if nextMiss5zwcy == maxFields5zwcy {
								// filled all the empty fields!
								break doneWithStruct5zwcy
							}
							missingFieldsLeft5zwcy--
							curField5zwcy = unmarshalMsgFieldOrder5zwcy[nextMiss5zwcy]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zwcy)
						switch curField5zwcy {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zwcy[0] = true
							z.Structs[zbic].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zwcy[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zbic].Fields) = (z.Structs[zbic].Fields)[:0]
							} else {

								var zrcp uint32
								zrcp, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zbic].Fields) >= int(zrcp) {
									z.Structs[zbic].Fields = (z.Structs[zbic].Fields)[:zrcp]
								} else {
									z.Structs[zbic].Fields = make([]Field, zrcp)
								}
								for zlsq := range z.Structs[zbic].Fields {
									bts, err = z.Structs[zbic].Fields[zlsq].UnmarshalMsg(bts)
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
					if nextMiss5zwcy != -1 {
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
	if nextMiss4zudb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zudb = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip4zudb = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zwcy = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zwcy = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zbic := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zbic].StructName) + 7 + msgp.ArrayHeaderSize
		for zlsq := range z.Structs[zbic].Fields {
			s += z.Structs[zbic].Fields[zlsq].Msgsize()
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
	const maxFields6zndy = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zndy uint32
	totalEncodedFields6zndy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zndy := totalEncodedFields6zndy
	missingFieldsLeft6zndy := maxFields6zndy - totalEncodedFields6zndy

	var nextMiss6zndy int32 = -1
	var found6zndy [maxFields6zndy]bool
	var curField6zndy string

doneWithStruct6zndy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zndy > 0 || missingFieldsLeft6zndy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zndy, missingFieldsLeft6zndy, msgp.ShowFound(found6zndy[:]), decodeMsgFieldOrder6zndy)
		if encodedFieldsLeft6zndy > 0 {
			encodedFieldsLeft6zndy--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zndy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zndy < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zndy = 0
			}
			for nextMiss6zndy < maxFields6zndy && (found6zndy[nextMiss6zndy] || decodeMsgFieldSkip6zndy[nextMiss6zndy]) {
				nextMiss6zndy++
			}
			if nextMiss6zndy == maxFields6zndy {
				// filled all the empty fields!
				break doneWithStruct6zndy
			}
			missingFieldsLeft6zndy--
			curField6zndy = decodeMsgFieldOrder6zndy[nextMiss6zndy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zndy)
		switch curField6zndy {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zndy[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zndy[1] = true
			var zwbg uint32
			zwbg, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zwbg) {
				z.Fields = (z.Fields)[:zwbg]
			} else {
				z.Fields = make([]Field, zwbg)
			}
			for zvkj := range z.Fields {
				err = z.Fields[zvkj].DecodeMsg(dc)
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
	if nextMiss6zndy != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zndy = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zndy = []bool{false, false}

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
	for zvkj := range z.Fields {
		err = z.Fields[zvkj].EncodeMsg(en)
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
	for zvkj := range z.Fields {
		o, err = z.Fields[zvkj].MarshalMsg(o)
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
	const maxFields7zmdi = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zmdi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zmdi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zmdi := totalEncodedFields7zmdi
	missingFieldsLeft7zmdi := maxFields7zmdi - totalEncodedFields7zmdi

	var nextMiss7zmdi int32 = -1
	var found7zmdi [maxFields7zmdi]bool
	var curField7zmdi string

doneWithStruct7zmdi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zmdi > 0 || missingFieldsLeft7zmdi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zmdi, missingFieldsLeft7zmdi, msgp.ShowFound(found7zmdi[:]), unmarshalMsgFieldOrder7zmdi)
		if encodedFieldsLeft7zmdi > 0 {
			encodedFieldsLeft7zmdi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zmdi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zmdi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zmdi = 0
			}
			for nextMiss7zmdi < maxFields7zmdi && (found7zmdi[nextMiss7zmdi] || unmarshalMsgFieldSkip7zmdi[nextMiss7zmdi]) {
				nextMiss7zmdi++
			}
			if nextMiss7zmdi == maxFields7zmdi {
				// filled all the empty fields!
				break doneWithStruct7zmdi
			}
			missingFieldsLeft7zmdi--
			curField7zmdi = unmarshalMsgFieldOrder7zmdi[nextMiss7zmdi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zmdi)
		switch curField7zmdi {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zmdi[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zmdi[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zecr uint32
				zecr, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zecr) {
					z.Fields = (z.Fields)[:zecr]
				} else {
					z.Fields = make([]Field, zecr)
				}
				for zvkj := range z.Fields {
					bts, err = z.Fields[zvkj].UnmarshalMsg(bts)
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
	if nextMiss7zmdi != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zmdi = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zmdi = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zvkj := range z.Fields {
		s += z.Fields[zvkj].Msgsize()
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
		var zrcb uint64
		zrcb, err = dc.ReadUint64()
		(*z) = Zkind(zrcb)
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
		var zchb uint64
		zchb, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zchb)
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
	const maxFields8znqs = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8znqs uint32
	totalEncodedFields8znqs, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8znqs := totalEncodedFields8znqs
	missingFieldsLeft8znqs := maxFields8znqs - totalEncodedFields8znqs

	var nextMiss8znqs int32 = -1
	var found8znqs [maxFields8znqs]bool
	var curField8znqs string

doneWithStruct8znqs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8znqs > 0 || missingFieldsLeft8znqs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8znqs, missingFieldsLeft8znqs, msgp.ShowFound(found8znqs[:]), decodeMsgFieldOrder8znqs)
		if encodedFieldsLeft8znqs > 0 {
			encodedFieldsLeft8znqs--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8znqs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8znqs < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8znqs = 0
			}
			for nextMiss8znqs < maxFields8znqs && (found8znqs[nextMiss8znqs] || decodeMsgFieldSkip8znqs[nextMiss8znqs]) {
				nextMiss8znqs++
			}
			if nextMiss8znqs == maxFields8znqs {
				// filled all the empty fields!
				break doneWithStruct8znqs
			}
			missingFieldsLeft8znqs--
			curField8znqs = decodeMsgFieldOrder8znqs[nextMiss8znqs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8znqs)
		switch curField8znqs {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8znqs[0] = true
			{
				var zegx uint64
				zegx, err = dc.ReadUint64()
				z.Kind = Zkind(zegx)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8znqs[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8znqs[2] = true
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
			found8znqs[3] = true
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
	if nextMiss8znqs != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8znqs = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8znqs = []bool{false, false, false, false}

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

	// honor the omitempty tags
	var empty_zwlf [4]bool
	fieldsInUse_zbgf := z.fieldsNotEmpty(empty_zwlf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zbgf)
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
		panic(err)
	}
	if !empty_zwlf[1] {
		// write "Str"
		err = en.Append(0xa3, 0x53, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Str)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zwlf[2] {
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

	if !empty_zwlf[3] {
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
				panic(err)
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
	const maxFields9zmpo = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zmpo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zmpo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zmpo := totalEncodedFields9zmpo
	missingFieldsLeft9zmpo := maxFields9zmpo - totalEncodedFields9zmpo

	var nextMiss9zmpo int32 = -1
	var found9zmpo [maxFields9zmpo]bool
	var curField9zmpo string

doneWithStruct9zmpo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zmpo > 0 || missingFieldsLeft9zmpo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zmpo, missingFieldsLeft9zmpo, msgp.ShowFound(found9zmpo[:]), unmarshalMsgFieldOrder9zmpo)
		if encodedFieldsLeft9zmpo > 0 {
			encodedFieldsLeft9zmpo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zmpo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zmpo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zmpo = 0
			}
			for nextMiss9zmpo < maxFields9zmpo && (found9zmpo[nextMiss9zmpo] || unmarshalMsgFieldSkip9zmpo[nextMiss9zmpo]) {
				nextMiss9zmpo++
			}
			if nextMiss9zmpo == maxFields9zmpo {
				// filled all the empty fields!
				break doneWithStruct9zmpo
			}
			missingFieldsLeft9zmpo--
			curField9zmpo = unmarshalMsgFieldOrder9zmpo[nextMiss9zmpo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zmpo)
		switch curField9zmpo {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zmpo[0] = true
			{
				var zccn uint64
				zccn, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zccn)
			}
		case "Str":
			found9zmpo[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zmpo[2] = true
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
			found9zmpo[3] = true
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
	if nextMiss9zmpo != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zmpo = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zmpo = []bool{false, false, false, false}

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
