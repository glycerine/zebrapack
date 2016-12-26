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
	const maxFields0zlmm = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zlmm uint32
	totalEncodedFields0zlmm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zlmm := totalEncodedFields0zlmm
	missingFieldsLeft0zlmm := maxFields0zlmm - totalEncodedFields0zlmm

	var nextMiss0zlmm int32 = -1
	var found0zlmm [maxFields0zlmm]bool
	var curField0zlmm string

doneWithStruct0zlmm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zlmm > 0 || missingFieldsLeft0zlmm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zlmm, missingFieldsLeft0zlmm, msgp.ShowFound(found0zlmm[:]), decodeMsgFieldOrder0zlmm)
		if encodedFieldsLeft0zlmm > 0 {
			encodedFieldsLeft0zlmm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zlmm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zlmm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zlmm = 0
			}
			for nextMiss0zlmm < maxFields0zlmm && (found0zlmm[nextMiss0zlmm] || decodeMsgFieldSkip0zlmm[nextMiss0zlmm]) {
				nextMiss0zlmm++
			}
			if nextMiss0zlmm == maxFields0zlmm {
				// filled all the empty fields!
				break doneWithStruct0zlmm
			}
			missingFieldsLeft0zlmm--
			curField0zlmm = decodeMsgFieldOrder0zlmm[nextMiss0zlmm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zlmm)
		switch curField0zlmm {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zlmm[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zlmm[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zlmm[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zlmm[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zlmm[4] = true
			{
				var zywg uint64
				zywg, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zywg)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zlmm[5] = true
			{
				var zhqy uint64
				zhqy, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zhqy)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zlmm[6] = true
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
			found0zlmm[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zlmm[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zlmm[9] = true
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
	if nextMiss0zlmm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zlmm = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zlmm = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zvot [10]bool
	fieldsInUse_zktj := z.fieldsNotEmpty(empty_zvot[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zktj)
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
	if !empty_zvot[3] {
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

	if !empty_zvot[4] {
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

	if !empty_zvot[5] {
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

	if !empty_zvot[6] {
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

	if !empty_zvot[7] {
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

	if !empty_zvot[8] {
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

	if !empty_zvot[9] {
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
	const maxFields1zfgo = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zfgo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zfgo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zfgo := totalEncodedFields1zfgo
	missingFieldsLeft1zfgo := maxFields1zfgo - totalEncodedFields1zfgo

	var nextMiss1zfgo int32 = -1
	var found1zfgo [maxFields1zfgo]bool
	var curField1zfgo string

doneWithStruct1zfgo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zfgo > 0 || missingFieldsLeft1zfgo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zfgo, missingFieldsLeft1zfgo, msgp.ShowFound(found1zfgo[:]), unmarshalMsgFieldOrder1zfgo)
		if encodedFieldsLeft1zfgo > 0 {
			encodedFieldsLeft1zfgo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zfgo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zfgo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zfgo = 0
			}
			for nextMiss1zfgo < maxFields1zfgo && (found1zfgo[nextMiss1zfgo] || unmarshalMsgFieldSkip1zfgo[nextMiss1zfgo]) {
				nextMiss1zfgo++
			}
			if nextMiss1zfgo == maxFields1zfgo {
				// filled all the empty fields!
				break doneWithStruct1zfgo
			}
			missingFieldsLeft1zfgo--
			curField1zfgo = unmarshalMsgFieldOrder1zfgo[nextMiss1zfgo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zfgo)
		switch curField1zfgo {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zfgo[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zfgo[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zfgo[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zfgo[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zfgo[4] = true
			{
				var zzft uint64
				zzft, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zzft)
			}
		case "FieldPrimitive":
			found1zfgo[5] = true
			{
				var zasv uint64
				zasv, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zasv)
			}
		case "FieldFullType":
			found1zfgo[6] = true
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
			found1zfgo[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zfgo[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zfgo[9] = true
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
	if nextMiss1zfgo != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zfgo = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zfgo = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zrif = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zrif uint32
	totalEncodedFields2zrif, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zrif := totalEncodedFields2zrif
	missingFieldsLeft2zrif := maxFields2zrif - totalEncodedFields2zrif

	var nextMiss2zrif int32 = -1
	var found2zrif [maxFields2zrif]bool
	var curField2zrif string

doneWithStruct2zrif:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zrif > 0 || missingFieldsLeft2zrif > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zrif, missingFieldsLeft2zrif, msgp.ShowFound(found2zrif[:]), decodeMsgFieldOrder2zrif)
		if encodedFieldsLeft2zrif > 0 {
			encodedFieldsLeft2zrif--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zrif = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zrif < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zrif = 0
			}
			for nextMiss2zrif < maxFields2zrif && (found2zrif[nextMiss2zrif] || decodeMsgFieldSkip2zrif[nextMiss2zrif]) {
				nextMiss2zrif++
			}
			if nextMiss2zrif == maxFields2zrif {
				// filled all the empty fields!
				break doneWithStruct2zrif
			}
			missingFieldsLeft2zrif--
			curField2zrif = decodeMsgFieldOrder2zrif[nextMiss2zrif]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zrif)
		switch curField2zrif {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zrif[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zrif[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zrif[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zrif[3] = true
			var zwdk uint32
			zwdk, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zwdk) {
				z.Structs = (z.Structs)[:zwdk]
			} else {
				z.Structs = make([]Struct, zwdk)
			}
			for zsqq := range z.Structs {
				const maxFields3zesm = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zesm uint32
				totalEncodedFields3zesm, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zesm := totalEncodedFields3zesm
				missingFieldsLeft3zesm := maxFields3zesm - totalEncodedFields3zesm

				var nextMiss3zesm int32 = -1
				var found3zesm [maxFields3zesm]bool
				var curField3zesm string

			doneWithStruct3zesm:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zesm > 0 || missingFieldsLeft3zesm > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zesm, missingFieldsLeft3zesm, msgp.ShowFound(found3zesm[:]), decodeMsgFieldOrder3zesm)
					if encodedFieldsLeft3zesm > 0 {
						encodedFieldsLeft3zesm--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zesm = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zesm < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zesm = 0
						}
						for nextMiss3zesm < maxFields3zesm && (found3zesm[nextMiss3zesm] || decodeMsgFieldSkip3zesm[nextMiss3zesm]) {
							nextMiss3zesm++
						}
						if nextMiss3zesm == maxFields3zesm {
							// filled all the empty fields!
							break doneWithStruct3zesm
						}
						missingFieldsLeft3zesm--
						curField3zesm = decodeMsgFieldOrder3zesm[nextMiss3zesm]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zesm)
					switch curField3zesm {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zesm[0] = true
						z.Structs[zsqq].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zesm[1] = true
						var zoiw uint32
						zoiw, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zsqq].Fields) >= int(zoiw) {
							z.Structs[zsqq].Fields = (z.Structs[zsqq].Fields)[:zoiw]
						} else {
							z.Structs[zsqq].Fields = make([]Field, zoiw)
						}
						for zuiz := range z.Structs[zsqq].Fields {
							err = z.Structs[zsqq].Fields[zuiz].DecodeMsg(dc)
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
				if nextMiss3zesm != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zrif[4] = true
			var zkhc uint32
			zkhc, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zkhc) {
				z.Imports = (z.Imports)[:zkhc]
			} else {
				z.Imports = make([]string, zkhc)
			}
			for ztru := range z.Imports {
				z.Imports[ztru], err = dc.ReadString()
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
	if nextMiss2zrif != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zrif = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zrif = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zesm = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zesm = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Schema) fieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// EncodeMsg implements msgp.Encodable
func (z *Schema) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "SourcePath"
	err = en.Append(0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
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
	for zsqq := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zsqq].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zsqq].Fields)))
		if err != nil {
			panic(err)
		}
		for zuiz := range z.Structs[zsqq].Fields {
			err = z.Structs[zsqq].Fields[zuiz].EncodeMsg(en)
			if err != nil {
				panic(err)
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
		panic(err)
	}
	for ztru := range z.Imports {
		err = en.WriteString(z.Imports[ztru])
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Schema) MarshalMsg(b []byte) (o []byte, err error) {
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
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zsqq := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zsqq].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zsqq].Fields)))
		for zuiz := range z.Structs[zsqq].Fields {
			o, err = z.Structs[zsqq].Fields[zuiz].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for ztru := range z.Imports {
		o = msgp.AppendString(o, z.Imports[ztru])
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
	const maxFields4ztni = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4ztni uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4ztni, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4ztni := totalEncodedFields4ztni
	missingFieldsLeft4ztni := maxFields4ztni - totalEncodedFields4ztni

	var nextMiss4ztni int32 = -1
	var found4ztni [maxFields4ztni]bool
	var curField4ztni string

doneWithStruct4ztni:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4ztni > 0 || missingFieldsLeft4ztni > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4ztni, missingFieldsLeft4ztni, msgp.ShowFound(found4ztni[:]), unmarshalMsgFieldOrder4ztni)
		if encodedFieldsLeft4ztni > 0 {
			encodedFieldsLeft4ztni--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4ztni = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4ztni < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4ztni = 0
			}
			for nextMiss4ztni < maxFields4ztni && (found4ztni[nextMiss4ztni] || unmarshalMsgFieldSkip4ztni[nextMiss4ztni]) {
				nextMiss4ztni++
			}
			if nextMiss4ztni == maxFields4ztni {
				// filled all the empty fields!
				break doneWithStruct4ztni
			}
			missingFieldsLeft4ztni--
			curField4ztni = unmarshalMsgFieldOrder4ztni[nextMiss4ztni]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4ztni)
		switch curField4ztni {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4ztni[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4ztni[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4ztni[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4ztni[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zjdw uint32
				zjdw, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zjdw) {
					z.Structs = (z.Structs)[:zjdw]
				} else {
					z.Structs = make([]Struct, zjdw)
				}
				for zsqq := range z.Structs {
					const maxFields5zwdx = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zwdx uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zwdx, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zwdx := totalEncodedFields5zwdx
					missingFieldsLeft5zwdx := maxFields5zwdx - totalEncodedFields5zwdx

					var nextMiss5zwdx int32 = -1
					var found5zwdx [maxFields5zwdx]bool
					var curField5zwdx string

				doneWithStruct5zwdx:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zwdx > 0 || missingFieldsLeft5zwdx > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zwdx, missingFieldsLeft5zwdx, msgp.ShowFound(found5zwdx[:]), unmarshalMsgFieldOrder5zwdx)
						if encodedFieldsLeft5zwdx > 0 {
							encodedFieldsLeft5zwdx--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zwdx = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zwdx < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zwdx = 0
							}
							for nextMiss5zwdx < maxFields5zwdx && (found5zwdx[nextMiss5zwdx] || unmarshalMsgFieldSkip5zwdx[nextMiss5zwdx]) {
								nextMiss5zwdx++
							}
							if nextMiss5zwdx == maxFields5zwdx {
								// filled all the empty fields!
								break doneWithStruct5zwdx
							}
							missingFieldsLeft5zwdx--
							curField5zwdx = unmarshalMsgFieldOrder5zwdx[nextMiss5zwdx]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zwdx)
						switch curField5zwdx {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zwdx[0] = true
							z.Structs[zsqq].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zwdx[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zsqq].Fields) = (z.Structs[zsqq].Fields)[:0]
							} else {

								var zedz uint32
								zedz, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zsqq].Fields) >= int(zedz) {
									z.Structs[zsqq].Fields = (z.Structs[zsqq].Fields)[:zedz]
								} else {
									z.Structs[zsqq].Fields = make([]Field, zedz)
								}
								for zuiz := range z.Structs[zsqq].Fields {
									bts, err = z.Structs[zsqq].Fields[zuiz].UnmarshalMsg(bts)
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
					if nextMiss5zwdx != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4ztni[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zeng uint32
				zeng, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zeng) {
					z.Imports = (z.Imports)[:zeng]
				} else {
					z.Imports = make([]string, zeng)
				}
				for ztru := range z.Imports {
					z.Imports[ztru], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4ztni != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4ztni = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4ztni = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zwdx = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zwdx = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zsqq := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zsqq].StructName) + 7 + msgp.ArrayHeaderSize
		for zuiz := range z.Structs[zsqq].Fields {
			s += z.Structs[zsqq].Fields[zuiz].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for ztru := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[ztru])
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
	const maxFields6zqiu = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zqiu uint32
	totalEncodedFields6zqiu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zqiu := totalEncodedFields6zqiu
	missingFieldsLeft6zqiu := maxFields6zqiu - totalEncodedFields6zqiu

	var nextMiss6zqiu int32 = -1
	var found6zqiu [maxFields6zqiu]bool
	var curField6zqiu string

doneWithStruct6zqiu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zqiu > 0 || missingFieldsLeft6zqiu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zqiu, missingFieldsLeft6zqiu, msgp.ShowFound(found6zqiu[:]), decodeMsgFieldOrder6zqiu)
		if encodedFieldsLeft6zqiu > 0 {
			encodedFieldsLeft6zqiu--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zqiu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zqiu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zqiu = 0
			}
			for nextMiss6zqiu < maxFields6zqiu && (found6zqiu[nextMiss6zqiu] || decodeMsgFieldSkip6zqiu[nextMiss6zqiu]) {
				nextMiss6zqiu++
			}
			if nextMiss6zqiu == maxFields6zqiu {
				// filled all the empty fields!
				break doneWithStruct6zqiu
			}
			missingFieldsLeft6zqiu--
			curField6zqiu = decodeMsgFieldOrder6zqiu[nextMiss6zqiu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zqiu)
		switch curField6zqiu {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zqiu[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zqiu[1] = true
			var zgeg uint32
			zgeg, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zgeg) {
				z.Fields = (z.Fields)[:zgeg]
			} else {
				z.Fields = make([]Field, zgeg)
			}
			for zbat := range z.Fields {
				err = z.Fields[zbat].DecodeMsg(dc)
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
	if nextMiss6zqiu != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zqiu = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zqiu = []bool{false, false}

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
	for zbat := range z.Fields {
		err = z.Fields[zbat].EncodeMsg(en)
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
	for zbat := range z.Fields {
		o, err = z.Fields[zbat].MarshalMsg(o)
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
	const maxFields7zdce = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zdce uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zdce, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zdce := totalEncodedFields7zdce
	missingFieldsLeft7zdce := maxFields7zdce - totalEncodedFields7zdce

	var nextMiss7zdce int32 = -1
	var found7zdce [maxFields7zdce]bool
	var curField7zdce string

doneWithStruct7zdce:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zdce > 0 || missingFieldsLeft7zdce > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zdce, missingFieldsLeft7zdce, msgp.ShowFound(found7zdce[:]), unmarshalMsgFieldOrder7zdce)
		if encodedFieldsLeft7zdce > 0 {
			encodedFieldsLeft7zdce--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zdce = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zdce < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zdce = 0
			}
			for nextMiss7zdce < maxFields7zdce && (found7zdce[nextMiss7zdce] || unmarshalMsgFieldSkip7zdce[nextMiss7zdce]) {
				nextMiss7zdce++
			}
			if nextMiss7zdce == maxFields7zdce {
				// filled all the empty fields!
				break doneWithStruct7zdce
			}
			missingFieldsLeft7zdce--
			curField7zdce = unmarshalMsgFieldOrder7zdce[nextMiss7zdce]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zdce)
		switch curField7zdce {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zdce[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zdce[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zbvv uint32
				zbvv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zbvv) {
					z.Fields = (z.Fields)[:zbvv]
				} else {
					z.Fields = make([]Field, zbvv)
				}
				for zbat := range z.Fields {
					bts, err = z.Fields[zbat].UnmarshalMsg(bts)
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
	if nextMiss7zdce != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zdce = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zdce = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zbat := range z.Fields {
		s += z.Fields[zbat].Msgsize()
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
		var zzdn uint64
		zzdn, err = dc.ReadUint64()
		(*z) = Zkind(zzdn)
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
		var zkdc uint64
		zkdc, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zkdc)
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
	const maxFields8zxka = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zxka uint32
	totalEncodedFields8zxka, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zxka := totalEncodedFields8zxka
	missingFieldsLeft8zxka := maxFields8zxka - totalEncodedFields8zxka

	var nextMiss8zxka int32 = -1
	var found8zxka [maxFields8zxka]bool
	var curField8zxka string

doneWithStruct8zxka:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zxka > 0 || missingFieldsLeft8zxka > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zxka, missingFieldsLeft8zxka, msgp.ShowFound(found8zxka[:]), decodeMsgFieldOrder8zxka)
		if encodedFieldsLeft8zxka > 0 {
			encodedFieldsLeft8zxka--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zxka = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zxka < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zxka = 0
			}
			for nextMiss8zxka < maxFields8zxka && (found8zxka[nextMiss8zxka] || decodeMsgFieldSkip8zxka[nextMiss8zxka]) {
				nextMiss8zxka++
			}
			if nextMiss8zxka == maxFields8zxka {
				// filled all the empty fields!
				break doneWithStruct8zxka
			}
			missingFieldsLeft8zxka--
			curField8zxka = decodeMsgFieldOrder8zxka[nextMiss8zxka]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zxka)
		switch curField8zxka {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zxka[0] = true
			{
				var zwdk uint64
				zwdk, err = dc.ReadUint64()
				z.Kind = Zkind(zwdk)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zxka[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zxka[2] = true
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
			found8zxka[3] = true
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
	if nextMiss8zxka != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zxka = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zxka = []bool{false, false, false, false}

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
	var empty_zvfk [4]bool
	fieldsInUse_zjql := z.fieldsNotEmpty(empty_zvfk[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjql)
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
	if !empty_zvfk[1] {
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

	if !empty_zvfk[2] {
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

	if !empty_zvfk[3] {
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
	const maxFields9zovy = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zovy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zovy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zovy := totalEncodedFields9zovy
	missingFieldsLeft9zovy := maxFields9zovy - totalEncodedFields9zovy

	var nextMiss9zovy int32 = -1
	var found9zovy [maxFields9zovy]bool
	var curField9zovy string

doneWithStruct9zovy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zovy > 0 || missingFieldsLeft9zovy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zovy, missingFieldsLeft9zovy, msgp.ShowFound(found9zovy[:]), unmarshalMsgFieldOrder9zovy)
		if encodedFieldsLeft9zovy > 0 {
			encodedFieldsLeft9zovy--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zovy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zovy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zovy = 0
			}
			for nextMiss9zovy < maxFields9zovy && (found9zovy[nextMiss9zovy] || unmarshalMsgFieldSkip9zovy[nextMiss9zovy]) {
				nextMiss9zovy++
			}
			if nextMiss9zovy == maxFields9zovy {
				// filled all the empty fields!
				break doneWithStruct9zovy
			}
			missingFieldsLeft9zovy--
			curField9zovy = unmarshalMsgFieldOrder9zovy[nextMiss9zovy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zovy)
		switch curField9zovy {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zovy[0] = true
			{
				var zixl uint64
				zixl, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zixl)
			}
		case "Str":
			found9zovy[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zovy[2] = true
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
			found9zovy[3] = true
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
	if nextMiss9zovy != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zovy = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zovy = []bool{false, false, false, false}

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
