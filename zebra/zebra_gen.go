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
	const maxFields0ztpc = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0ztpc uint32
	totalEncodedFields0ztpc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0ztpc := totalEncodedFields0ztpc
	missingFieldsLeft0ztpc := maxFields0ztpc - totalEncodedFields0ztpc

	var nextMiss0ztpc int32 = -1
	var found0ztpc [maxFields0ztpc]bool
	var curField0ztpc string

doneWithStruct0ztpc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ztpc > 0 || missingFieldsLeft0ztpc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ztpc, missingFieldsLeft0ztpc, msgp.ShowFound(found0ztpc[:]), decodeMsgFieldOrder0ztpc)
		if encodedFieldsLeft0ztpc > 0 {
			encodedFieldsLeft0ztpc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0ztpc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ztpc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0ztpc = 0
			}
			for nextMiss0ztpc < maxFields0ztpc && (found0ztpc[nextMiss0ztpc] || decodeMsgFieldSkip0ztpc[nextMiss0ztpc]) {
				nextMiss0ztpc++
			}
			if nextMiss0ztpc == maxFields0ztpc {
				// filled all the empty fields!
				break doneWithStruct0ztpc
			}
			missingFieldsLeft0ztpc--
			curField0ztpc = decodeMsgFieldOrder0ztpc[nextMiss0ztpc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ztpc)
		switch curField0ztpc {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0ztpc[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0ztpc[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0ztpc[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0ztpc[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0ztpc[4] = true
			{
				var zfhz uint64
				zfhz, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zfhz)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0ztpc[5] = true
			{
				var ztsy uint64
				ztsy, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(ztsy)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0ztpc[6] = true
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
			found0ztpc[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0ztpc[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0ztpc[9] = true
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
	if nextMiss0ztpc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0ztpc = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0ztpc = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zrlo [10]bool
	fieldsInUse_zgez := z.fieldsNotEmpty(empty_zrlo[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgez)
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
	if !empty_zrlo[3] {
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

	if !empty_zrlo[4] {
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

	if !empty_zrlo[5] {
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

	if !empty_zrlo[6] {
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

	if !empty_zrlo[7] {
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

	if !empty_zrlo[8] {
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

	if !empty_zrlo[9] {
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
	const maxFields1zdkf = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zdkf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zdkf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zdkf := totalEncodedFields1zdkf
	missingFieldsLeft1zdkf := maxFields1zdkf - totalEncodedFields1zdkf

	var nextMiss1zdkf int32 = -1
	var found1zdkf [maxFields1zdkf]bool
	var curField1zdkf string

doneWithStruct1zdkf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zdkf > 0 || missingFieldsLeft1zdkf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zdkf, missingFieldsLeft1zdkf, msgp.ShowFound(found1zdkf[:]), unmarshalMsgFieldOrder1zdkf)
		if encodedFieldsLeft1zdkf > 0 {
			encodedFieldsLeft1zdkf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zdkf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zdkf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zdkf = 0
			}
			for nextMiss1zdkf < maxFields1zdkf && (found1zdkf[nextMiss1zdkf] || unmarshalMsgFieldSkip1zdkf[nextMiss1zdkf]) {
				nextMiss1zdkf++
			}
			if nextMiss1zdkf == maxFields1zdkf {
				// filled all the empty fields!
				break doneWithStruct1zdkf
			}
			missingFieldsLeft1zdkf--
			curField1zdkf = unmarshalMsgFieldOrder1zdkf[nextMiss1zdkf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zdkf)
		switch curField1zdkf {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zdkf[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zdkf[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zdkf[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zdkf[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zdkf[4] = true
			{
				var zxup uint64
				zxup, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zxup)
			}
		case "FieldPrimitive":
			found1zdkf[5] = true
			{
				var zdfr uint64
				zdfr, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zdfr)
			}
		case "FieldFullType":
			found1zdkf[6] = true
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
			found1zdkf[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zdkf[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zdkf[9] = true
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
	if nextMiss1zdkf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zdkf = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zdkf = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zkkp = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zkkp uint32
	totalEncodedFields2zkkp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zkkp := totalEncodedFields2zkkp
	missingFieldsLeft2zkkp := maxFields2zkkp - totalEncodedFields2zkkp

	var nextMiss2zkkp int32 = -1
	var found2zkkp [maxFields2zkkp]bool
	var curField2zkkp string

doneWithStruct2zkkp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zkkp > 0 || missingFieldsLeft2zkkp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zkkp, missingFieldsLeft2zkkp, msgp.ShowFound(found2zkkp[:]), decodeMsgFieldOrder2zkkp)
		if encodedFieldsLeft2zkkp > 0 {
			encodedFieldsLeft2zkkp--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zkkp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zkkp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zkkp = 0
			}
			for nextMiss2zkkp < maxFields2zkkp && (found2zkkp[nextMiss2zkkp] || decodeMsgFieldSkip2zkkp[nextMiss2zkkp]) {
				nextMiss2zkkp++
			}
			if nextMiss2zkkp == maxFields2zkkp {
				// filled all the empty fields!
				break doneWithStruct2zkkp
			}
			missingFieldsLeft2zkkp--
			curField2zkkp = decodeMsgFieldOrder2zkkp[nextMiss2zkkp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zkkp)
		switch curField2zkkp {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zkkp[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zkkp[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zkkp[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zkkp[3] = true
			var zupb uint32
			zupb, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zupb) {
				z.Structs = (z.Structs)[:zupb]
			} else {
				z.Structs = make([]Struct, zupb)
			}
			for zckc := range z.Structs {
				const maxFields3zcbl = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zcbl uint32
				totalEncodedFields3zcbl, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zcbl := totalEncodedFields3zcbl
				missingFieldsLeft3zcbl := maxFields3zcbl - totalEncodedFields3zcbl

				var nextMiss3zcbl int32 = -1
				var found3zcbl [maxFields3zcbl]bool
				var curField3zcbl string

			doneWithStruct3zcbl:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zcbl > 0 || missingFieldsLeft3zcbl > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zcbl, missingFieldsLeft3zcbl, msgp.ShowFound(found3zcbl[:]), decodeMsgFieldOrder3zcbl)
					if encodedFieldsLeft3zcbl > 0 {
						encodedFieldsLeft3zcbl--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zcbl = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zcbl < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zcbl = 0
						}
						for nextMiss3zcbl < maxFields3zcbl && (found3zcbl[nextMiss3zcbl] || decodeMsgFieldSkip3zcbl[nextMiss3zcbl]) {
							nextMiss3zcbl++
						}
						if nextMiss3zcbl == maxFields3zcbl {
							// filled all the empty fields!
							break doneWithStruct3zcbl
						}
						missingFieldsLeft3zcbl--
						curField3zcbl = decodeMsgFieldOrder3zcbl[nextMiss3zcbl]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zcbl)
					switch curField3zcbl {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zcbl[0] = true
						z.Structs[zckc].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zcbl[1] = true
						var zhvr uint32
						zhvr, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zckc].Fields) >= int(zhvr) {
							z.Structs[zckc].Fields = (z.Structs[zckc].Fields)[:zhvr]
						} else {
							z.Structs[zckc].Fields = make([]Field, zhvr)
						}
						for zfox := range z.Structs[zckc].Fields {
							err = z.Structs[zckc].Fields[zfox].DecodeMsg(dc)
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
				if nextMiss3zcbl != -1 {
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
	if nextMiss2zkkp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zkkp = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip2zkkp = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zcbl = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zcbl = []bool{false, false}

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
	for zckc := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zckc].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zckc].Fields)))
		if err != nil {
			panic(err)
		}
		for zfox := range z.Structs[zckc].Fields {
			err = z.Structs[zckc].Fields[zfox].EncodeMsg(en)
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
	for zckc := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zckc].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zckc].Fields)))
		for zfox := range z.Structs[zckc].Fields {
			o, err = z.Structs[zckc].Fields[zfox].MarshalMsg(o)
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
	const maxFields4zrbk = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zrbk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zrbk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zrbk := totalEncodedFields4zrbk
	missingFieldsLeft4zrbk := maxFields4zrbk - totalEncodedFields4zrbk

	var nextMiss4zrbk int32 = -1
	var found4zrbk [maxFields4zrbk]bool
	var curField4zrbk string

doneWithStruct4zrbk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zrbk > 0 || missingFieldsLeft4zrbk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zrbk, missingFieldsLeft4zrbk, msgp.ShowFound(found4zrbk[:]), unmarshalMsgFieldOrder4zrbk)
		if encodedFieldsLeft4zrbk > 0 {
			encodedFieldsLeft4zrbk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zrbk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zrbk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zrbk = 0
			}
			for nextMiss4zrbk < maxFields4zrbk && (found4zrbk[nextMiss4zrbk] || unmarshalMsgFieldSkip4zrbk[nextMiss4zrbk]) {
				nextMiss4zrbk++
			}
			if nextMiss4zrbk == maxFields4zrbk {
				// filled all the empty fields!
				break doneWithStruct4zrbk
			}
			missingFieldsLeft4zrbk--
			curField4zrbk = unmarshalMsgFieldOrder4zrbk[nextMiss4zrbk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zrbk)
		switch curField4zrbk {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zrbk[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zrbk[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zrbk[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zrbk[3] = true
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
				for zckc := range z.Structs {
					const maxFields5zdhr = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zdhr uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zdhr, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zdhr := totalEncodedFields5zdhr
					missingFieldsLeft5zdhr := maxFields5zdhr - totalEncodedFields5zdhr

					var nextMiss5zdhr int32 = -1
					var found5zdhr [maxFields5zdhr]bool
					var curField5zdhr string

				doneWithStruct5zdhr:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zdhr > 0 || missingFieldsLeft5zdhr > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zdhr, missingFieldsLeft5zdhr, msgp.ShowFound(found5zdhr[:]), unmarshalMsgFieldOrder5zdhr)
						if encodedFieldsLeft5zdhr > 0 {
							encodedFieldsLeft5zdhr--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zdhr = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zdhr < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zdhr = 0
							}
							for nextMiss5zdhr < maxFields5zdhr && (found5zdhr[nextMiss5zdhr] || unmarshalMsgFieldSkip5zdhr[nextMiss5zdhr]) {
								nextMiss5zdhr++
							}
							if nextMiss5zdhr == maxFields5zdhr {
								// filled all the empty fields!
								break doneWithStruct5zdhr
							}
							missingFieldsLeft5zdhr--
							curField5zdhr = unmarshalMsgFieldOrder5zdhr[nextMiss5zdhr]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zdhr)
						switch curField5zdhr {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zdhr[0] = true
							z.Structs[zckc].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zdhr[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zckc].Fields) = (z.Structs[zckc].Fields)[:0]
							} else {

								var zgwn uint32
								zgwn, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zckc].Fields) >= int(zgwn) {
									z.Structs[zckc].Fields = (z.Structs[zckc].Fields)[:zgwn]
								} else {
									z.Structs[zckc].Fields = make([]Field, zgwn)
								}
								for zfox := range z.Structs[zckc].Fields {
									bts, err = z.Structs[zckc].Fields[zfox].UnmarshalMsg(bts)
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
					if nextMiss5zdhr != -1 {
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
	if nextMiss4zrbk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zrbk = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip4zrbk = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zdhr = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zdhr = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zckc := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zckc].StructName) + 7 + msgp.ArrayHeaderSize
		for zfox := range z.Structs[zckc].Fields {
			s += z.Structs[zckc].Fields[zfox].Msgsize()
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
	const maxFields6zxps = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zxps uint32
	totalEncodedFields6zxps, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zxps := totalEncodedFields6zxps
	missingFieldsLeft6zxps := maxFields6zxps - totalEncodedFields6zxps

	var nextMiss6zxps int32 = -1
	var found6zxps [maxFields6zxps]bool
	var curField6zxps string

doneWithStruct6zxps:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zxps > 0 || missingFieldsLeft6zxps > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zxps, missingFieldsLeft6zxps, msgp.ShowFound(found6zxps[:]), decodeMsgFieldOrder6zxps)
		if encodedFieldsLeft6zxps > 0 {
			encodedFieldsLeft6zxps--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zxps = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zxps < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zxps = 0
			}
			for nextMiss6zxps < maxFields6zxps && (found6zxps[nextMiss6zxps] || decodeMsgFieldSkip6zxps[nextMiss6zxps]) {
				nextMiss6zxps++
			}
			if nextMiss6zxps == maxFields6zxps {
				// filled all the empty fields!
				break doneWithStruct6zxps
			}
			missingFieldsLeft6zxps--
			curField6zxps = decodeMsgFieldOrder6zxps[nextMiss6zxps]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zxps)
		switch curField6zxps {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zxps[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zxps[1] = true
			var zibo uint32
			zibo, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zibo) {
				z.Fields = (z.Fields)[:zibo]
			} else {
				z.Fields = make([]Field, zibo)
			}
			for ztkw := range z.Fields {
				err = z.Fields[ztkw].DecodeMsg(dc)
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
	if nextMiss6zxps != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zxps = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zxps = []bool{false, false}

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
	for ztkw := range z.Fields {
		err = z.Fields[ztkw].EncodeMsg(en)
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
	for ztkw := range z.Fields {
		o, err = z.Fields[ztkw].MarshalMsg(o)
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
	const maxFields7zdly = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zdly uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zdly, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zdly := totalEncodedFields7zdly
	missingFieldsLeft7zdly := maxFields7zdly - totalEncodedFields7zdly

	var nextMiss7zdly int32 = -1
	var found7zdly [maxFields7zdly]bool
	var curField7zdly string

doneWithStruct7zdly:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zdly > 0 || missingFieldsLeft7zdly > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zdly, missingFieldsLeft7zdly, msgp.ShowFound(found7zdly[:]), unmarshalMsgFieldOrder7zdly)
		if encodedFieldsLeft7zdly > 0 {
			encodedFieldsLeft7zdly--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zdly = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zdly < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zdly = 0
			}
			for nextMiss7zdly < maxFields7zdly && (found7zdly[nextMiss7zdly] || unmarshalMsgFieldSkip7zdly[nextMiss7zdly]) {
				nextMiss7zdly++
			}
			if nextMiss7zdly == maxFields7zdly {
				// filled all the empty fields!
				break doneWithStruct7zdly
			}
			missingFieldsLeft7zdly--
			curField7zdly = unmarshalMsgFieldOrder7zdly[nextMiss7zdly]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zdly)
		switch curField7zdly {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zdly[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zdly[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zqao uint32
				zqao, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zqao) {
					z.Fields = (z.Fields)[:zqao]
				} else {
					z.Fields = make([]Field, zqao)
				}
				for ztkw := range z.Fields {
					bts, err = z.Fields[ztkw].UnmarshalMsg(bts)
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
	if nextMiss7zdly != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zdly = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zdly = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for ztkw := range z.Fields {
		s += z.Fields[ztkw].Msgsize()
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
		var zryh uint64
		zryh, err = dc.ReadUint64()
		(*z) = Zkind(zryh)
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
		var zsjt uint64
		zsjt, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zsjt)
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
	const maxFields8zdqt = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zdqt uint32
	totalEncodedFields8zdqt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zdqt := totalEncodedFields8zdqt
	missingFieldsLeft8zdqt := maxFields8zdqt - totalEncodedFields8zdqt

	var nextMiss8zdqt int32 = -1
	var found8zdqt [maxFields8zdqt]bool
	var curField8zdqt string

doneWithStruct8zdqt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zdqt > 0 || missingFieldsLeft8zdqt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zdqt, missingFieldsLeft8zdqt, msgp.ShowFound(found8zdqt[:]), decodeMsgFieldOrder8zdqt)
		if encodedFieldsLeft8zdqt > 0 {
			encodedFieldsLeft8zdqt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zdqt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zdqt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zdqt = 0
			}
			for nextMiss8zdqt < maxFields8zdqt && (found8zdqt[nextMiss8zdqt] || decodeMsgFieldSkip8zdqt[nextMiss8zdqt]) {
				nextMiss8zdqt++
			}
			if nextMiss8zdqt == maxFields8zdqt {
				// filled all the empty fields!
				break doneWithStruct8zdqt
			}
			missingFieldsLeft8zdqt--
			curField8zdqt = decodeMsgFieldOrder8zdqt[nextMiss8zdqt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zdqt)
		switch curField8zdqt {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zdqt[0] = true
			{
				var zncu uint64
				zncu, err = dc.ReadUint64()
				z.Kind = Zkind(zncu)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zdqt[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zdqt[2] = true
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
			found8zdqt[3] = true
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
	if nextMiss8zdqt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zdqt = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zdqt = []bool{false, false, false, false}

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
	var empty_zphw [4]bool
	fieldsInUse_zdve := z.fieldsNotEmpty(empty_zphw[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zdve)
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
	if !empty_zphw[1] {
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

	if !empty_zphw[2] {
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

	if !empty_zphw[3] {
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
	const maxFields9zoiz = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zoiz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zoiz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zoiz := totalEncodedFields9zoiz
	missingFieldsLeft9zoiz := maxFields9zoiz - totalEncodedFields9zoiz

	var nextMiss9zoiz int32 = -1
	var found9zoiz [maxFields9zoiz]bool
	var curField9zoiz string

doneWithStruct9zoiz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zoiz > 0 || missingFieldsLeft9zoiz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zoiz, missingFieldsLeft9zoiz, msgp.ShowFound(found9zoiz[:]), unmarshalMsgFieldOrder9zoiz)
		if encodedFieldsLeft9zoiz > 0 {
			encodedFieldsLeft9zoiz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zoiz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zoiz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zoiz = 0
			}
			for nextMiss9zoiz < maxFields9zoiz && (found9zoiz[nextMiss9zoiz] || unmarshalMsgFieldSkip9zoiz[nextMiss9zoiz]) {
				nextMiss9zoiz++
			}
			if nextMiss9zoiz == maxFields9zoiz {
				// filled all the empty fields!
				break doneWithStruct9zoiz
			}
			missingFieldsLeft9zoiz--
			curField9zoiz = unmarshalMsgFieldOrder9zoiz[nextMiss9zoiz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zoiz)
		switch curField9zoiz {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zoiz[0] = true
			{
				var zeyr uint64
				zeyr, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zeyr)
			}
		case "Str":
			found9zoiz[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zoiz[2] = true
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
			found9zoiz[3] = true
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
	if nextMiss9zoiz != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zoiz = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zoiz = []bool{false, false, false, false}

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
