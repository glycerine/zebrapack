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
	const maxFields0zqoi = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zqoi uint32
	totalEncodedFields0zqoi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zqoi := totalEncodedFields0zqoi
	missingFieldsLeft0zqoi := maxFields0zqoi - totalEncodedFields0zqoi

	var nextMiss0zqoi int32 = -1
	var found0zqoi [maxFields0zqoi]bool
	var curField0zqoi string

doneWithStruct0zqoi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zqoi > 0 || missingFieldsLeft0zqoi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zqoi, missingFieldsLeft0zqoi, msgp.ShowFound(found0zqoi[:]), decodeMsgFieldOrder0zqoi)
		if encodedFieldsLeft0zqoi > 0 {
			encodedFieldsLeft0zqoi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zqoi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zqoi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zqoi = 0
			}
			for nextMiss0zqoi < maxFields0zqoi && (found0zqoi[nextMiss0zqoi] || decodeMsgFieldSkip0zqoi[nextMiss0zqoi]) {
				nextMiss0zqoi++
			}
			if nextMiss0zqoi == maxFields0zqoi {
				// filled all the empty fields!
				break doneWithStruct0zqoi
			}
			missingFieldsLeft0zqoi--
			curField0zqoi = decodeMsgFieldOrder0zqoi[nextMiss0zqoi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zqoi)
		switch curField0zqoi {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zqoi[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zqoi[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zqoi[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zqoi[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zqoi[4] = true
			{
				var zbtm uint64
				zbtm, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zbtm)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zqoi[5] = true
			{
				var znfu uint64
				znfu, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(znfu)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zqoi[6] = true
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
			found0zqoi[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zqoi[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zqoi[9] = true
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
	if nextMiss0zqoi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zqoi = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zqoi = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zvkb [10]bool
	fieldsInUse_zzkw := z.fieldsNotEmpty(empty_zvkb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zzkw)
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
	if !empty_zvkb[3] {
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

	if !empty_zvkb[4] {
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

	if !empty_zvkb[5] {
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

	if !empty_zvkb[6] {
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

	if !empty_zvkb[7] {
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

	if !empty_zvkb[8] {
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

	if !empty_zvkb[9] {
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
	const maxFields1zuye = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zuye uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zuye, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zuye := totalEncodedFields1zuye
	missingFieldsLeft1zuye := maxFields1zuye - totalEncodedFields1zuye

	var nextMiss1zuye int32 = -1
	var found1zuye [maxFields1zuye]bool
	var curField1zuye string

doneWithStruct1zuye:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zuye > 0 || missingFieldsLeft1zuye > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zuye, missingFieldsLeft1zuye, msgp.ShowFound(found1zuye[:]), unmarshalMsgFieldOrder1zuye)
		if encodedFieldsLeft1zuye > 0 {
			encodedFieldsLeft1zuye--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zuye = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zuye < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zuye = 0
			}
			for nextMiss1zuye < maxFields1zuye && (found1zuye[nextMiss1zuye] || unmarshalMsgFieldSkip1zuye[nextMiss1zuye]) {
				nextMiss1zuye++
			}
			if nextMiss1zuye == maxFields1zuye {
				// filled all the empty fields!
				break doneWithStruct1zuye
			}
			missingFieldsLeft1zuye--
			curField1zuye = unmarshalMsgFieldOrder1zuye[nextMiss1zuye]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zuye)
		switch curField1zuye {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zuye[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zuye[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zuye[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zuye[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zuye[4] = true
			{
				var zywp uint64
				zywp, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zywp)
			}
		case "FieldPrimitive":
			found1zuye[5] = true
			{
				var zdpx uint64
				zdpx, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zdpx)
			}
		case "FieldFullType":
			found1zuye[6] = true
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
			found1zuye[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zuye[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zuye[9] = true
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
	if nextMiss1zuye != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zuye = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zuye = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zlpd = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zlpd uint32
	totalEncodedFields2zlpd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zlpd := totalEncodedFields2zlpd
	missingFieldsLeft2zlpd := maxFields2zlpd - totalEncodedFields2zlpd

	var nextMiss2zlpd int32 = -1
	var found2zlpd [maxFields2zlpd]bool
	var curField2zlpd string

doneWithStruct2zlpd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zlpd > 0 || missingFieldsLeft2zlpd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zlpd, missingFieldsLeft2zlpd, msgp.ShowFound(found2zlpd[:]), decodeMsgFieldOrder2zlpd)
		if encodedFieldsLeft2zlpd > 0 {
			encodedFieldsLeft2zlpd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zlpd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zlpd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zlpd = 0
			}
			for nextMiss2zlpd < maxFields2zlpd && (found2zlpd[nextMiss2zlpd] || decodeMsgFieldSkip2zlpd[nextMiss2zlpd]) {
				nextMiss2zlpd++
			}
			if nextMiss2zlpd == maxFields2zlpd {
				// filled all the empty fields!
				break doneWithStruct2zlpd
			}
			missingFieldsLeft2zlpd--
			curField2zlpd = decodeMsgFieldOrder2zlpd[nextMiss2zlpd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zlpd)
		switch curField2zlpd {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zlpd[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zlpd[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zlpd[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zlpd[3] = true
			var zesr uint32
			zesr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zesr) {
				z.Structs = (z.Structs)[:zesr]
			} else {
				z.Structs = make([]Struct, zesr)
			}
			for zhdk := range z.Structs {
				const maxFields3zspl = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zspl uint32
				totalEncodedFields3zspl, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zspl := totalEncodedFields3zspl
				missingFieldsLeft3zspl := maxFields3zspl - totalEncodedFields3zspl

				var nextMiss3zspl int32 = -1
				var found3zspl [maxFields3zspl]bool
				var curField3zspl string

			doneWithStruct3zspl:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zspl > 0 || missingFieldsLeft3zspl > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zspl, missingFieldsLeft3zspl, msgp.ShowFound(found3zspl[:]), decodeMsgFieldOrder3zspl)
					if encodedFieldsLeft3zspl > 0 {
						encodedFieldsLeft3zspl--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zspl = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zspl < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zspl = 0
						}
						for nextMiss3zspl < maxFields3zspl && (found3zspl[nextMiss3zspl] || decodeMsgFieldSkip3zspl[nextMiss3zspl]) {
							nextMiss3zspl++
						}
						if nextMiss3zspl == maxFields3zspl {
							// filled all the empty fields!
							break doneWithStruct3zspl
						}
						missingFieldsLeft3zspl--
						curField3zspl = decodeMsgFieldOrder3zspl[nextMiss3zspl]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zspl)
					switch curField3zspl {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zspl[0] = true
						z.Structs[zhdk].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zspl[1] = true
						var zvom uint32
						zvom, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zhdk].Fields) >= int(zvom) {
							z.Structs[zhdk].Fields = (z.Structs[zhdk].Fields)[:zvom]
						} else {
							z.Structs[zhdk].Fields = make([]Field, zvom)
						}
						for zmpj := range z.Structs[zhdk].Fields {
							err = z.Structs[zhdk].Fields[zmpj].DecodeMsg(dc)
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
				if nextMiss3zspl != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zlpd[4] = true
			var zlqd uint32
			zlqd, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zlqd) {
				z.Imports = (z.Imports)[:zlqd]
			} else {
				z.Imports = make([]string, zlqd)
			}
			for zzku := range z.Imports {
				z.Imports[zzku], err = dc.ReadString()
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
	if nextMiss2zlpd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zlpd = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zlpd = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zspl = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zspl = []bool{false, false}

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
	for zhdk := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zhdk].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zhdk].Fields)))
		if err != nil {
			panic(err)
		}
		for zmpj := range z.Structs[zhdk].Fields {
			err = z.Structs[zhdk].Fields[zmpj].EncodeMsg(en)
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
	for zzku := range z.Imports {
		err = en.WriteString(z.Imports[zzku])
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
	for zhdk := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zhdk].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zhdk].Fields)))
		for zmpj := range z.Structs[zhdk].Fields {
			o, err = z.Structs[zhdk].Fields[zmpj].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zzku := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zzku])
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
	const maxFields4zgda = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zgda uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zgda, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zgda := totalEncodedFields4zgda
	missingFieldsLeft4zgda := maxFields4zgda - totalEncodedFields4zgda

	var nextMiss4zgda int32 = -1
	var found4zgda [maxFields4zgda]bool
	var curField4zgda string

doneWithStruct4zgda:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zgda > 0 || missingFieldsLeft4zgda > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zgda, missingFieldsLeft4zgda, msgp.ShowFound(found4zgda[:]), unmarshalMsgFieldOrder4zgda)
		if encodedFieldsLeft4zgda > 0 {
			encodedFieldsLeft4zgda--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zgda = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zgda < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zgda = 0
			}
			for nextMiss4zgda < maxFields4zgda && (found4zgda[nextMiss4zgda] || unmarshalMsgFieldSkip4zgda[nextMiss4zgda]) {
				nextMiss4zgda++
			}
			if nextMiss4zgda == maxFields4zgda {
				// filled all the empty fields!
				break doneWithStruct4zgda
			}
			missingFieldsLeft4zgda--
			curField4zgda = unmarshalMsgFieldOrder4zgda[nextMiss4zgda]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zgda)
		switch curField4zgda {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zgda[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zgda[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zgda[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zgda[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zlww uint32
				zlww, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zlww) {
					z.Structs = (z.Structs)[:zlww]
				} else {
					z.Structs = make([]Struct, zlww)
				}
				for zhdk := range z.Structs {
					const maxFields5zgwl = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zgwl uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zgwl, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zgwl := totalEncodedFields5zgwl
					missingFieldsLeft5zgwl := maxFields5zgwl - totalEncodedFields5zgwl

					var nextMiss5zgwl int32 = -1
					var found5zgwl [maxFields5zgwl]bool
					var curField5zgwl string

				doneWithStruct5zgwl:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zgwl > 0 || missingFieldsLeft5zgwl > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zgwl, missingFieldsLeft5zgwl, msgp.ShowFound(found5zgwl[:]), unmarshalMsgFieldOrder5zgwl)
						if encodedFieldsLeft5zgwl > 0 {
							encodedFieldsLeft5zgwl--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zgwl = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zgwl < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zgwl = 0
							}
							for nextMiss5zgwl < maxFields5zgwl && (found5zgwl[nextMiss5zgwl] || unmarshalMsgFieldSkip5zgwl[nextMiss5zgwl]) {
								nextMiss5zgwl++
							}
							if nextMiss5zgwl == maxFields5zgwl {
								// filled all the empty fields!
								break doneWithStruct5zgwl
							}
							missingFieldsLeft5zgwl--
							curField5zgwl = unmarshalMsgFieldOrder5zgwl[nextMiss5zgwl]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zgwl)
						switch curField5zgwl {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zgwl[0] = true
							z.Structs[zhdk].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zgwl[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zhdk].Fields) = (z.Structs[zhdk].Fields)[:0]
							} else {

								var zzio uint32
								zzio, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zhdk].Fields) >= int(zzio) {
									z.Structs[zhdk].Fields = (z.Structs[zhdk].Fields)[:zzio]
								} else {
									z.Structs[zhdk].Fields = make([]Field, zzio)
								}
								for zmpj := range z.Structs[zhdk].Fields {
									bts, err = z.Structs[zhdk].Fields[zmpj].UnmarshalMsg(bts)
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
					if nextMiss5zgwl != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4zgda[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zkrp uint32
				zkrp, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zkrp) {
					z.Imports = (z.Imports)[:zkrp]
				} else {
					z.Imports = make([]string, zkrp)
				}
				for zzku := range z.Imports {
					z.Imports[zzku], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zgda != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zgda = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zgda = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zgwl = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zgwl = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zhdk := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zhdk].StructName) + 7 + msgp.ArrayHeaderSize
		for zmpj := range z.Structs[zhdk].Fields {
			s += z.Structs[zhdk].Fields[zmpj].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zzku := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zzku])
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
	const maxFields6zrai = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zrai uint32
	totalEncodedFields6zrai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zrai := totalEncodedFields6zrai
	missingFieldsLeft6zrai := maxFields6zrai - totalEncodedFields6zrai

	var nextMiss6zrai int32 = -1
	var found6zrai [maxFields6zrai]bool
	var curField6zrai string

doneWithStruct6zrai:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zrai > 0 || missingFieldsLeft6zrai > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zrai, missingFieldsLeft6zrai, msgp.ShowFound(found6zrai[:]), decodeMsgFieldOrder6zrai)
		if encodedFieldsLeft6zrai > 0 {
			encodedFieldsLeft6zrai--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zrai = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zrai < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zrai = 0
			}
			for nextMiss6zrai < maxFields6zrai && (found6zrai[nextMiss6zrai] || decodeMsgFieldSkip6zrai[nextMiss6zrai]) {
				nextMiss6zrai++
			}
			if nextMiss6zrai == maxFields6zrai {
				// filled all the empty fields!
				break doneWithStruct6zrai
			}
			missingFieldsLeft6zrai--
			curField6zrai = decodeMsgFieldOrder6zrai[nextMiss6zrai]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zrai)
		switch curField6zrai {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zrai[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zrai[1] = true
			var zrkn uint32
			zrkn, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zrkn) {
				z.Fields = (z.Fields)[:zrkn]
			} else {
				z.Fields = make([]Field, zrkn)
			}
			for zrsr := range z.Fields {
				err = z.Fields[zrsr].DecodeMsg(dc)
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
	if nextMiss6zrai != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zrai = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zrai = []bool{false, false}

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
	for zrsr := range z.Fields {
		err = z.Fields[zrsr].EncodeMsg(en)
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
	for zrsr := range z.Fields {
		o, err = z.Fields[zrsr].MarshalMsg(o)
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
	const maxFields7zhon = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zhon uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zhon, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zhon := totalEncodedFields7zhon
	missingFieldsLeft7zhon := maxFields7zhon - totalEncodedFields7zhon

	var nextMiss7zhon int32 = -1
	var found7zhon [maxFields7zhon]bool
	var curField7zhon string

doneWithStruct7zhon:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zhon > 0 || missingFieldsLeft7zhon > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zhon, missingFieldsLeft7zhon, msgp.ShowFound(found7zhon[:]), unmarshalMsgFieldOrder7zhon)
		if encodedFieldsLeft7zhon > 0 {
			encodedFieldsLeft7zhon--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zhon = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zhon < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zhon = 0
			}
			for nextMiss7zhon < maxFields7zhon && (found7zhon[nextMiss7zhon] || unmarshalMsgFieldSkip7zhon[nextMiss7zhon]) {
				nextMiss7zhon++
			}
			if nextMiss7zhon == maxFields7zhon {
				// filled all the empty fields!
				break doneWithStruct7zhon
			}
			missingFieldsLeft7zhon--
			curField7zhon = unmarshalMsgFieldOrder7zhon[nextMiss7zhon]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zhon)
		switch curField7zhon {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zhon[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zhon[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zewy uint32
				zewy, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zewy) {
					z.Fields = (z.Fields)[:zewy]
				} else {
					z.Fields = make([]Field, zewy)
				}
				for zrsr := range z.Fields {
					bts, err = z.Fields[zrsr].UnmarshalMsg(bts)
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
	if nextMiss7zhon != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zhon = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zhon = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zrsr := range z.Fields {
		s += z.Fields[zrsr].Msgsize()
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
		var zluw uint64
		zluw, err = dc.ReadUint64()
		(*z) = Zkind(zluw)
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
		var zfhn uint64
		zfhn, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zfhn)
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
	const maxFields8zlau = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zlau uint32
	totalEncodedFields8zlau, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zlau := totalEncodedFields8zlau
	missingFieldsLeft8zlau := maxFields8zlau - totalEncodedFields8zlau

	var nextMiss8zlau int32 = -1
	var found8zlau [maxFields8zlau]bool
	var curField8zlau string

doneWithStruct8zlau:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zlau > 0 || missingFieldsLeft8zlau > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zlau, missingFieldsLeft8zlau, msgp.ShowFound(found8zlau[:]), decodeMsgFieldOrder8zlau)
		if encodedFieldsLeft8zlau > 0 {
			encodedFieldsLeft8zlau--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zlau = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zlau < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zlau = 0
			}
			for nextMiss8zlau < maxFields8zlau && (found8zlau[nextMiss8zlau] || decodeMsgFieldSkip8zlau[nextMiss8zlau]) {
				nextMiss8zlau++
			}
			if nextMiss8zlau == maxFields8zlau {
				// filled all the empty fields!
				break doneWithStruct8zlau
			}
			missingFieldsLeft8zlau--
			curField8zlau = decodeMsgFieldOrder8zlau[nextMiss8zlau]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zlau)
		switch curField8zlau {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zlau[0] = true
			{
				var zoyd uint64
				zoyd, err = dc.ReadUint64()
				z.Kind = Zkind(zoyd)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zlau[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zlau[2] = true
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
			found8zlau[3] = true
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
	if nextMiss8zlau != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zlau = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zlau = []bool{false, false, false, false}

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
	var empty_zzih [4]bool
	fieldsInUse_zszj := z.fieldsNotEmpty(empty_zzih[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zszj)
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
	if !empty_zzih[1] {
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

	if !empty_zzih[2] {
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

	if !empty_zzih[3] {
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
	const maxFields9zqge = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zqge uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zqge, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zqge := totalEncodedFields9zqge
	missingFieldsLeft9zqge := maxFields9zqge - totalEncodedFields9zqge

	var nextMiss9zqge int32 = -1
	var found9zqge [maxFields9zqge]bool
	var curField9zqge string

doneWithStruct9zqge:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zqge > 0 || missingFieldsLeft9zqge > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zqge, missingFieldsLeft9zqge, msgp.ShowFound(found9zqge[:]), unmarshalMsgFieldOrder9zqge)
		if encodedFieldsLeft9zqge > 0 {
			encodedFieldsLeft9zqge--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zqge = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zqge < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zqge = 0
			}
			for nextMiss9zqge < maxFields9zqge && (found9zqge[nextMiss9zqge] || unmarshalMsgFieldSkip9zqge[nextMiss9zqge]) {
				nextMiss9zqge++
			}
			if nextMiss9zqge == maxFields9zqge {
				// filled all the empty fields!
				break doneWithStruct9zqge
			}
			missingFieldsLeft9zqge--
			curField9zqge = unmarshalMsgFieldOrder9zqge[nextMiss9zqge]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zqge)
		switch curField9zqge {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zqge[0] = true
			{
				var zkfv uint64
				zkfv, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zkfv)
			}
		case "Str":
			found9zqge[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zqge[2] = true
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
			found9zqge[3] = true
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
	if nextMiss9zqge != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zqge = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zqge = []bool{false, false, false, false}

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
