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
	const maxFields0zwxm = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zwxm uint32
	totalEncodedFields0zwxm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zwxm := totalEncodedFields0zwxm
	missingFieldsLeft0zwxm := maxFields0zwxm - totalEncodedFields0zwxm

	var nextMiss0zwxm int32 = -1
	var found0zwxm [maxFields0zwxm]bool
	var curField0zwxm string

doneWithStruct0zwxm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zwxm > 0 || missingFieldsLeft0zwxm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zwxm, missingFieldsLeft0zwxm, msgp.ShowFound(found0zwxm[:]), decodeMsgFieldOrder0zwxm)
		if encodedFieldsLeft0zwxm > 0 {
			encodedFieldsLeft0zwxm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zwxm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zwxm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zwxm = 0
			}
			for nextMiss0zwxm < maxFields0zwxm && (found0zwxm[nextMiss0zwxm] || decodeMsgFieldSkip0zwxm[nextMiss0zwxm]) {
				nextMiss0zwxm++
			}
			if nextMiss0zwxm == maxFields0zwxm {
				// filled all the empty fields!
				break doneWithStruct0zwxm
			}
			missingFieldsLeft0zwxm--
			curField0zwxm = decodeMsgFieldOrder0zwxm[nextMiss0zwxm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zwxm)
		switch curField0zwxm {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zwxm[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zwxm[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zwxm[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zwxm[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zwxm[4] = true
			{
				var ztaj uint64
				ztaj, err = dc.ReadUint64()
				z.FieldCategory = Zkind(ztaj)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zwxm[5] = true
			{
				var zsqs uint64
				zsqs, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zsqs)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zwxm[6] = true
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
			found0zwxm[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zwxm[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zwxm[9] = true
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
	if nextMiss0zwxm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zwxm = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zwxm = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zfeh [10]bool
	fieldsInUse_znme := z.fieldsNotEmpty(empty_zfeh[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_znme)
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
	if !empty_zfeh[3] {
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

	if !empty_zfeh[4] {
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

	if !empty_zfeh[5] {
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

	if !empty_zfeh[6] {
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

	if !empty_zfeh[7] {
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

	if !empty_zfeh[8] {
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

	if !empty_zfeh[9] {
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
	const maxFields1zypx = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zypx uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zypx, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zypx := totalEncodedFields1zypx
	missingFieldsLeft1zypx := maxFields1zypx - totalEncodedFields1zypx

	var nextMiss1zypx int32 = -1
	var found1zypx [maxFields1zypx]bool
	var curField1zypx string

doneWithStruct1zypx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zypx > 0 || missingFieldsLeft1zypx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zypx, missingFieldsLeft1zypx, msgp.ShowFound(found1zypx[:]), unmarshalMsgFieldOrder1zypx)
		if encodedFieldsLeft1zypx > 0 {
			encodedFieldsLeft1zypx--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zypx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zypx < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zypx = 0
			}
			for nextMiss1zypx < maxFields1zypx && (found1zypx[nextMiss1zypx] || unmarshalMsgFieldSkip1zypx[nextMiss1zypx]) {
				nextMiss1zypx++
			}
			if nextMiss1zypx == maxFields1zypx {
				// filled all the empty fields!
				break doneWithStruct1zypx
			}
			missingFieldsLeft1zypx--
			curField1zypx = unmarshalMsgFieldOrder1zypx[nextMiss1zypx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zypx)
		switch curField1zypx {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zypx[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zypx[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zypx[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zypx[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zypx[4] = true
			{
				var zmes uint64
				zmes, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zmes)
			}
		case "FieldPrimitive":
			found1zypx[5] = true
			{
				var zryg uint64
				zryg, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zryg)
			}
		case "FieldFullType":
			found1zypx[6] = true
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
			found1zypx[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zypx[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zypx[9] = true
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
	if nextMiss1zypx != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zypx = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zypx = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zjtw = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zjtw uint32
	totalEncodedFields2zjtw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zjtw := totalEncodedFields2zjtw
	missingFieldsLeft2zjtw := maxFields2zjtw - totalEncodedFields2zjtw

	var nextMiss2zjtw int32 = -1
	var found2zjtw [maxFields2zjtw]bool
	var curField2zjtw string

doneWithStruct2zjtw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zjtw > 0 || missingFieldsLeft2zjtw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zjtw, missingFieldsLeft2zjtw, msgp.ShowFound(found2zjtw[:]), decodeMsgFieldOrder2zjtw)
		if encodedFieldsLeft2zjtw > 0 {
			encodedFieldsLeft2zjtw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zjtw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zjtw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zjtw = 0
			}
			for nextMiss2zjtw < maxFields2zjtw && (found2zjtw[nextMiss2zjtw] || decodeMsgFieldSkip2zjtw[nextMiss2zjtw]) {
				nextMiss2zjtw++
			}
			if nextMiss2zjtw == maxFields2zjtw {
				// filled all the empty fields!
				break doneWithStruct2zjtw
			}
			missingFieldsLeft2zjtw--
			curField2zjtw = decodeMsgFieldOrder2zjtw[nextMiss2zjtw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zjtw)
		switch curField2zjtw {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zjtw[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zjtw[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zjtw[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zjtw[3] = true
			var zvqa uint32
			zvqa, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zvqa > 0 {
				z.Structs = make(map[string]*Struct, zvqa)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zvqa > 0 {
				zvqa--
				var zkdp string
				var zluc *Struct
				zkdp, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					zluc = nil
				} else {
					if zluc == nil {
						zluc = new(Struct)
					}
					const maxFields3zuis = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zuis uint32
					totalEncodedFields3zuis, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zuis := totalEncodedFields3zuis
					missingFieldsLeft3zuis := maxFields3zuis - totalEncodedFields3zuis

					var nextMiss3zuis int32 = -1
					var found3zuis [maxFields3zuis]bool
					var curField3zuis string

				doneWithStruct3zuis:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zuis > 0 || missingFieldsLeft3zuis > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zuis, missingFieldsLeft3zuis, msgp.ShowFound(found3zuis[:]), decodeMsgFieldOrder3zuis)
						if encodedFieldsLeft3zuis > 0 {
							encodedFieldsLeft3zuis--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zuis = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zuis < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zuis = 0
							}
							for nextMiss3zuis < maxFields3zuis && (found3zuis[nextMiss3zuis] || decodeMsgFieldSkip3zuis[nextMiss3zuis]) {
								nextMiss3zuis++
							}
							if nextMiss3zuis == maxFields3zuis {
								// filled all the empty fields!
								break doneWithStruct3zuis
							}
							missingFieldsLeft3zuis--
							curField3zuis = decodeMsgFieldOrder3zuis[nextMiss3zuis]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zuis)
						switch curField3zuis {
						// -- templateDecodeMsg ends here --

						case "StructName":
							found3zuis[0] = true
							zluc.StructName, err = dc.ReadString()
							if err != nil {
								panic(err)
							}
						case "Fields":
							found3zuis[1] = true
							var zkzk uint32
							zkzk, err = dc.ReadArrayHeader()
							if err != nil {
								panic(err)
							}
							if cap(zluc.Fields) >= int(zkzk) {
								zluc.Fields = (zluc.Fields)[:zkzk]
							} else {
								zluc.Fields = make([]Field, zkzk)
							}
							for ztfp := range zluc.Fields {
								err = zluc.Fields[ztfp].DecodeMsg(dc)
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
					if nextMiss3zuis != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[zkdp] = zluc
			}
		case "Imports":
			found2zjtw[4] = true
			var zrrv uint32
			zrrv, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zrrv) {
				z.Imports = (z.Imports)[:zrrv]
			} else {
				z.Imports = make([]string, zrrv)
			}
			for zunz := range z.Imports {
				z.Imports[zunz], err = dc.ReadString()
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
	if nextMiss2zjtw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zjtw = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zjtw = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zuis = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zuis = []bool{false, false}

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
	err = en.WriteMapHeader(uint32(len(z.Structs)))
	if err != nil {
		panic(err)
	}
	for zkdp, zluc := range z.Structs {
		err = en.WriteString(zkdp)
		if err != nil {
			panic(err)
		}
		if zluc == nil {
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
			err = en.WriteString(zluc.StructName)
			if err != nil {
				panic(err)
			}
			// write "Fields"
			err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(zluc.Fields)))
			if err != nil {
				panic(err)
			}
			for ztfp := range zluc.Fields {
				err = zluc.Fields[ztfp].EncodeMsg(en)
				if err != nil {
					panic(err)
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
		panic(err)
	}
	for zunz := range z.Imports {
		err = en.WriteString(z.Imports[zunz])
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
	o = msgp.AppendMapHeader(o, uint32(len(z.Structs)))
	for zkdp, zluc := range z.Structs {
		o = msgp.AppendString(o, zkdp)
		if zluc == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName"
			o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, zluc.StructName)
			// string "Fields"
			o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			o = msgp.AppendArrayHeader(o, uint32(len(zluc.Fields)))
			for ztfp := range zluc.Fields {
				o, err = zluc.Fields[ztfp].MarshalMsg(o)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zunz := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zunz])
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
	const maxFields4zkog = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zkog uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zkog, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zkog := totalEncodedFields4zkog
	missingFieldsLeft4zkog := maxFields4zkog - totalEncodedFields4zkog

	var nextMiss4zkog int32 = -1
	var found4zkog [maxFields4zkog]bool
	var curField4zkog string

doneWithStruct4zkog:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zkog > 0 || missingFieldsLeft4zkog > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zkog, missingFieldsLeft4zkog, msgp.ShowFound(found4zkog[:]), unmarshalMsgFieldOrder4zkog)
		if encodedFieldsLeft4zkog > 0 {
			encodedFieldsLeft4zkog--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zkog = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zkog < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zkog = 0
			}
			for nextMiss4zkog < maxFields4zkog && (found4zkog[nextMiss4zkog] || unmarshalMsgFieldSkip4zkog[nextMiss4zkog]) {
				nextMiss4zkog++
			}
			if nextMiss4zkog == maxFields4zkog {
				// filled all the empty fields!
				break doneWithStruct4zkog
			}
			missingFieldsLeft4zkog--
			curField4zkog = unmarshalMsgFieldOrder4zkog[nextMiss4zkog]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zkog)
		switch curField4zkog {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zkog[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zkog[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zkog[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zkog[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zpyb uint32
				zpyb, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zpyb > 0 {
					z.Structs = make(map[string]*Struct, zpyb)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zpyb > 0 {
					var zkdp string
					var zluc *Struct
					zpyb--
					zkdp, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && zluc == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if zluc == nil {
							zluc = new(Struct)
						}
						const maxFields5zrlo = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zrlo uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zrlo, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								panic(err)
								return
							}
						}
						encodedFieldsLeft5zrlo := totalEncodedFields5zrlo
						missingFieldsLeft5zrlo := maxFields5zrlo - totalEncodedFields5zrlo

						var nextMiss5zrlo int32 = -1
						var found5zrlo [maxFields5zrlo]bool
						var curField5zrlo string

					doneWithStruct5zrlo:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zrlo > 0 || missingFieldsLeft5zrlo > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zrlo, missingFieldsLeft5zrlo, msgp.ShowFound(found5zrlo[:]), unmarshalMsgFieldOrder5zrlo)
							if encodedFieldsLeft5zrlo > 0 {
								encodedFieldsLeft5zrlo--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									panic(err)
									return
								}
								curField5zrlo = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zrlo < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zrlo = 0
								}
								for nextMiss5zrlo < maxFields5zrlo && (found5zrlo[nextMiss5zrlo] || unmarshalMsgFieldSkip5zrlo[nextMiss5zrlo]) {
									nextMiss5zrlo++
								}
								if nextMiss5zrlo == maxFields5zrlo {
									// filled all the empty fields!
									break doneWithStruct5zrlo
								}
								missingFieldsLeft5zrlo--
								curField5zrlo = unmarshalMsgFieldOrder5zrlo[nextMiss5zrlo]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zrlo)
							switch curField5zrlo {
							// -- templateUnmarshalMsg ends here --

							case "StructName":
								found5zrlo[0] = true
								zluc.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									panic(err)
								}
							case "Fields":
								found5zrlo[1] = true
								if nbs.AlwaysNil {
									(zluc.Fields) = (zluc.Fields)[:0]
								} else {

									var zpia uint32
									zpia, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										panic(err)
									}
									if cap(zluc.Fields) >= int(zpia) {
										zluc.Fields = (zluc.Fields)[:zpia]
									} else {
										zluc.Fields = make([]Field, zpia)
									}
									for ztfp := range zluc.Fields {
										bts, err = zluc.Fields[ztfp].UnmarshalMsg(bts)
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
						if nextMiss5zrlo != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[zkdp] = zluc
				}
			}
		case "Imports":
			found4zkog[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var ztdd uint32
				ztdd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(ztdd) {
					z.Imports = (z.Imports)[:ztdd]
				} else {
					z.Imports = make([]string, ztdd)
				}
				for zunz := range z.Imports {
					z.Imports[zunz], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zkog != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zkog = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zkog = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zrlo = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zrlo = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zkdp, zluc := range z.Structs {
			_ = zluc
			_ = zkdp
			s += msgp.StringPrefixSize + len(zkdp)
			if zluc == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 11 + msgp.StringPrefixSize + len(zluc.StructName) + 7 + msgp.ArrayHeaderSize
				for ztfp := range zluc.Fields {
					s += zluc.Fields[ztfp].Msgsize()
				}
			}
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zunz := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zunz])
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
	const maxFields6znuj = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6znuj uint32
	totalEncodedFields6znuj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6znuj := totalEncodedFields6znuj
	missingFieldsLeft6znuj := maxFields6znuj - totalEncodedFields6znuj

	var nextMiss6znuj int32 = -1
	var found6znuj [maxFields6znuj]bool
	var curField6znuj string

doneWithStruct6znuj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6znuj > 0 || missingFieldsLeft6znuj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6znuj, missingFieldsLeft6znuj, msgp.ShowFound(found6znuj[:]), decodeMsgFieldOrder6znuj)
		if encodedFieldsLeft6znuj > 0 {
			encodedFieldsLeft6znuj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6znuj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6znuj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6znuj = 0
			}
			for nextMiss6znuj < maxFields6znuj && (found6znuj[nextMiss6znuj] || decodeMsgFieldSkip6znuj[nextMiss6znuj]) {
				nextMiss6znuj++
			}
			if nextMiss6znuj == maxFields6znuj {
				// filled all the empty fields!
				break doneWithStruct6znuj
			}
			missingFieldsLeft6znuj--
			curField6znuj = decodeMsgFieldOrder6znuj[nextMiss6znuj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6znuj)
		switch curField6znuj {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6znuj[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6znuj[1] = true
			var zmcr uint32
			zmcr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zmcr) {
				z.Fields = (z.Fields)[:zmcr]
			} else {
				z.Fields = make([]Field, zmcr)
			}
			for zlak := range z.Fields {
				err = z.Fields[zlak].DecodeMsg(dc)
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
	if nextMiss6znuj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6znuj = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6znuj = []bool{false, false}

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
	for zlak := range z.Fields {
		err = z.Fields[zlak].EncodeMsg(en)
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
	for zlak := range z.Fields {
		o, err = z.Fields[zlak].MarshalMsg(o)
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
	const maxFields7zqyf = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zqyf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zqyf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zqyf := totalEncodedFields7zqyf
	missingFieldsLeft7zqyf := maxFields7zqyf - totalEncodedFields7zqyf

	var nextMiss7zqyf int32 = -1
	var found7zqyf [maxFields7zqyf]bool
	var curField7zqyf string

doneWithStruct7zqyf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zqyf > 0 || missingFieldsLeft7zqyf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zqyf, missingFieldsLeft7zqyf, msgp.ShowFound(found7zqyf[:]), unmarshalMsgFieldOrder7zqyf)
		if encodedFieldsLeft7zqyf > 0 {
			encodedFieldsLeft7zqyf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zqyf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zqyf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zqyf = 0
			}
			for nextMiss7zqyf < maxFields7zqyf && (found7zqyf[nextMiss7zqyf] || unmarshalMsgFieldSkip7zqyf[nextMiss7zqyf]) {
				nextMiss7zqyf++
			}
			if nextMiss7zqyf == maxFields7zqyf {
				// filled all the empty fields!
				break doneWithStruct7zqyf
			}
			missingFieldsLeft7zqyf--
			curField7zqyf = unmarshalMsgFieldOrder7zqyf[nextMiss7zqyf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zqyf)
		switch curField7zqyf {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zqyf[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zqyf[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zlzv uint32
				zlzv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zlzv) {
					z.Fields = (z.Fields)[:zlzv]
				} else {
					z.Fields = make([]Field, zlzv)
				}
				for zlak := range z.Fields {
					bts, err = z.Fields[zlak].UnmarshalMsg(bts)
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
	if nextMiss7zqyf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zqyf = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zqyf = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zlak := range z.Fields {
		s += z.Fields[zlak].Msgsize()
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
		var zdvw uint64
		zdvw, err = dc.ReadUint64()
		(*z) = Zkind(zdvw)
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
		var zmbi uint64
		zmbi, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zmbi)
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
	const maxFields8zugo = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zugo uint32
	totalEncodedFields8zugo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zugo := totalEncodedFields8zugo
	missingFieldsLeft8zugo := maxFields8zugo - totalEncodedFields8zugo

	var nextMiss8zugo int32 = -1
	var found8zugo [maxFields8zugo]bool
	var curField8zugo string

doneWithStruct8zugo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zugo > 0 || missingFieldsLeft8zugo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zugo, missingFieldsLeft8zugo, msgp.ShowFound(found8zugo[:]), decodeMsgFieldOrder8zugo)
		if encodedFieldsLeft8zugo > 0 {
			encodedFieldsLeft8zugo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zugo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zugo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zugo = 0
			}
			for nextMiss8zugo < maxFields8zugo && (found8zugo[nextMiss8zugo] || decodeMsgFieldSkip8zugo[nextMiss8zugo]) {
				nextMiss8zugo++
			}
			if nextMiss8zugo == maxFields8zugo {
				// filled all the empty fields!
				break doneWithStruct8zugo
			}
			missingFieldsLeft8zugo--
			curField8zugo = decodeMsgFieldOrder8zugo[nextMiss8zugo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zugo)
		switch curField8zugo {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zugo[0] = true
			{
				var zraj uint64
				zraj, err = dc.ReadUint64()
				z.Kind = Zkind(zraj)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zugo[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zugo[2] = true
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
			found8zugo[3] = true
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
	if nextMiss8zugo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zugo = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zugo = []bool{false, false, false, false}

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
	var empty_zuuu [4]bool
	fieldsInUse_zmpw := z.fieldsNotEmpty(empty_zuuu[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zmpw)
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
	if !empty_zuuu[1] {
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

	if !empty_zuuu[2] {
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

	if !empty_zuuu[3] {
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
	const maxFields9zyip = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zyip uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zyip, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zyip := totalEncodedFields9zyip
	missingFieldsLeft9zyip := maxFields9zyip - totalEncodedFields9zyip

	var nextMiss9zyip int32 = -1
	var found9zyip [maxFields9zyip]bool
	var curField9zyip string

doneWithStruct9zyip:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zyip > 0 || missingFieldsLeft9zyip > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zyip, missingFieldsLeft9zyip, msgp.ShowFound(found9zyip[:]), unmarshalMsgFieldOrder9zyip)
		if encodedFieldsLeft9zyip > 0 {
			encodedFieldsLeft9zyip--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zyip = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zyip < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zyip = 0
			}
			for nextMiss9zyip < maxFields9zyip && (found9zyip[nextMiss9zyip] || unmarshalMsgFieldSkip9zyip[nextMiss9zyip]) {
				nextMiss9zyip++
			}
			if nextMiss9zyip == maxFields9zyip {
				// filled all the empty fields!
				break doneWithStruct9zyip
			}
			missingFieldsLeft9zyip--
			curField9zyip = unmarshalMsgFieldOrder9zyip[nextMiss9zyip]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zyip)
		switch curField9zyip {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zyip[0] = true
			{
				var zmst uint64
				zmst, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zmst)
			}
		case "Str":
			found9zyip[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zyip[2] = true
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
			found9zyip[3] = true
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
	if nextMiss9zyip != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zyip = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zyip = []bool{false, false, false, false}

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

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 3224 bytes
func ZebraSchemaInMsgpack2Format() []byte { return zebraSchemaInMsgpack2Format }

var zebraSchemaInMsgpack2Format = []byte{
	0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xa8, 0x7a, 0x65, 0x62,
	0x72, 0x61, 0x2e, 0x67, 0x6f, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0xa5, 0x7a, 0x65, 0x62, 0x72, 0x61, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0xd3, 0x0, 0x1, 0xa5, 0xa9, 0x4b, 0xd4, 0x96, 0x24,
	0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x84, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0x82,
	0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x5a, 0x74, 0x79, 0x70,
	0x65, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x94, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x4b, 0x69,
	0x6e, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5,
	0x5a, 0x6b, 0x69, 0x6e, 0x64, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0xb, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x53, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x53, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69,
	0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa9, 0x4f,
	0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53,
	0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x5a, 0x74, 0x79,
	0x70, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x87, 0xa3, 0x5a,
	0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xad, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xad, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c,
	0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x5a,
	0x74, 0x79, 0x70, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0xa6,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0xa6, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x95, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53,
	0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72,
	0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x49, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
	0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0xac, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb2,
	0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84,
	0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72,
	0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4, 0x4b,
	0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3, 0x53,
	0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x74, 0x72, 0xa8, 0x5b, 0x5d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e,
	0x64, 0x1a, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xa6, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x82, 0xaa, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0xa6,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x92, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
	0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x86,
	0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0xac, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa7, 0x5b, 0x5d, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4,
	0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xa6,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74,
	0x72, 0xa5, 0x46, 0x69, 0x65, 0x6c, 0x64, 0xa5, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x82, 0xaa, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x46, 0x69, 0x65, 0x6c, 0x64, 0xa6,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x9a, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x5a, 0x69, 0x64, 0xac, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x5a, 0x69, 0x64, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17,
	0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
	0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x87,
	0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53,
	0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82,
	0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x88, 0xa3, 0x5a, 0x69,
	0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0xac, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x74, 0x72, 0xa5, 0x5a, 0x6b, 0x69, 0x6e, 0x64, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0xb, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb, 0xa3, 0x53,
	0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0xc3, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x5a,
	0x6b, 0x69, 0x6e, 0x64, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0xb, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x87, 0xa3,
	0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xad, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xad, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
	0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa5,
	0x5a, 0x74, 0x79, 0x70, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3,
	0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xac, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f,
	0x6f, 0x6c, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x88, 0xa3, 0x5a,
	0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
	0x53, 0x6b, 0x69, 0x70, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0xa4, 0x53, 0x6b, 0x69, 0x70, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72,
	0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3,
	0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0xac, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4,
	0x62, 0x6f, 0x6f, 0x6c, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0xa7,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x90,
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 4152 bytes
func ZebraSchemaInJsonCompact() []byte { return zebraSchemaInJsonCompact }

var zebraSchemaInJsonCompact = []byte(`{"SourcePath":"zebra.go","SourcePackage":"zebra","ZebraSchemaId":463621516989988,"Structs":{"Ztype":{"StructName":"Ztype","Fields":[{"Zid":-1,"FieldGoName":"Kind","FieldTagName":"Kind","FieldTypeStr":"Zkind","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}},{"Zid":-1,"FieldGoName":"Str","FieldTagName":"Str","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Domain","FieldTagName":"Domain","FieldTypeStr":"*Ztype","FieldCategory":28,"FieldFullType":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Ztype"}},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Range","FieldTagName":"Range","FieldTypeStr":"*Ztype","FieldCategory":28,"FieldFullType":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Ztype"}},"OmitEmpty":true}]},"Schema":{"StructName":"Schema","Fields":[{"Zid":-1,"FieldGoName":"SourcePath","FieldTagName":"SourcePath","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"SourcePackage","FieldTagName":"SourcePackage","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"ZebraSchemaId","FieldTagName":"ZebraSchemaId","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":-1,"FieldGoName":"Structs","FieldTagName":"Structs","FieldTypeStr":"map[string]*Struct","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":25,"Str":"Struct"}}}},{"Zid":-1,"FieldGoName":"Imports","FieldTagName":"Imports","FieldTypeStr":"[]string","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":2,"Str":"string"}}}]},"Struct":{"StructName":"Struct","Fields":[{"Zid":-1,"FieldGoName":"StructName","FieldTagName":"StructName","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"Fields","FieldTagName":"Fields","FieldTypeStr":"[]Field","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":22,"Str":"Field"}}}]},"Field":{"StructName":"Field","Fields":[{"Zid":-1,"FieldGoName":"Zid","FieldTagName":"Zid","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":-1,"FieldGoName":"FieldGoName","FieldTagName":"FieldGoName","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"FieldTagName","FieldTagName":"FieldTagName","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"FieldTypeStr","FieldTagName":"FieldTypeStr","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"FieldCategory","FieldTagName":"FieldCategory","FieldTypeStr":"Zkind","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"FieldPrimitive","FieldTagName":"FieldPrimitive","FieldTypeStr":"Zkind","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"FieldFullType","FieldTagName":"FieldFullType","FieldTypeStr":"*Ztype","FieldCategory":28,"FieldFullType":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Ztype"}},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"OmitEmpty","FieldTagName":"OmitEmpty","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Skip","FieldTagName":"Skip","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Deprecated","FieldTagName":"Deprecated","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"},"OmitEmpty":true}]}},"Imports":[]}`)

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 10901 bytes
func ZebraSchemaInJsonPretty() []byte { return zebraSchemaInJsonPretty }

var zebraSchemaInJsonPretty = []byte(`{
    "SourcePath": "zebra.go",
    "SourcePackage": "zebra",
    "ZebraSchemaId": 463621516989988,
    "Structs": {
        "Ztype": {
            "StructName": "Ztype",
            "Fields": [
                {
                    "Zid": -1,
                    "FieldGoName": "Kind",
                    "FieldTagName": "Kind",
                    "FieldTypeStr": "Zkind",
                    "FieldCategory": 23,
                    "FieldPrimitive": 11,
                    "FieldFullType": {
                        "Kind": 11,
                        "Str": "uint64"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Str",
                    "FieldTagName": "Str",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Domain",
                    "FieldTagName": "Domain",
                    "FieldTypeStr": "*Ztype",
                    "FieldCategory": 28,
                    "FieldFullType": {
                        "Kind": 28,
                        "Str": "Pointer",
                        "Domain": {
                            "Kind": 22,
                            "Str": "Ztype"
                        }
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Range",
                    "FieldTagName": "Range",
                    "FieldTypeStr": "*Ztype",
                    "FieldCategory": 28,
                    "FieldFullType": {
                        "Kind": 28,
                        "Str": "Pointer",
                        "Domain": {
                            "Kind": 22,
                            "Str": "Ztype"
                        }
                    },
                    "OmitEmpty": true
                }
            ]
        },
        "Schema": {
            "StructName": "Schema",
            "Fields": [
                {
                    "Zid": -1,
                    "FieldGoName": "SourcePath",
                    "FieldTagName": "SourcePath",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "SourcePackage",
                    "FieldTagName": "SourcePackage",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "ZebraSchemaId",
                    "FieldTagName": "ZebraSchemaId",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Structs",
                    "FieldTagName": "Structs",
                    "FieldTypeStr": "map[string]*Struct",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 25,
                                "Str": "Struct"
                            }
                        }
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Imports",
                    "FieldTagName": "Imports",
                    "FieldTypeStr": "[]string",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        }
                    }
                }
            ]
        },
        "Struct": {
            "StructName": "Struct",
            "Fields": [
                {
                    "Zid": -1,
                    "FieldGoName": "StructName",
                    "FieldTagName": "StructName",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Fields",
                    "FieldTagName": "Fields",
                    "FieldTypeStr": "[]Field",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 22,
                            "Str": "Field"
                        }
                    }
                }
            ]
        },
        "Field": {
            "StructName": "Field",
            "Fields": [
                {
                    "Zid": -1,
                    "FieldGoName": "Zid",
                    "FieldTagName": "Zid",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "FieldGoName",
                    "FieldTagName": "FieldGoName",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "FieldTagName",
                    "FieldTagName": "FieldTagName",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "FieldTypeStr",
                    "FieldTagName": "FieldTypeStr",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "FieldCategory",
                    "FieldTagName": "FieldCategory",
                    "FieldTypeStr": "Zkind",
                    "FieldCategory": 23,
                    "FieldPrimitive": 11,
                    "FieldFullType": {
                        "Kind": 11,
                        "Str": "uint64"
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "FieldPrimitive",
                    "FieldTagName": "FieldPrimitive",
                    "FieldTypeStr": "Zkind",
                    "FieldCategory": 23,
                    "FieldPrimitive": 11,
                    "FieldFullType": {
                        "Kind": 11,
                        "Str": "uint64"
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "FieldFullType",
                    "FieldTagName": "FieldFullType",
                    "FieldTypeStr": "*Ztype",
                    "FieldCategory": 28,
                    "FieldFullType": {
                        "Kind": 28,
                        "Str": "Pointer",
                        "Domain": {
                            "Kind": 22,
                            "Str": "Ztype"
                        }
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "OmitEmpty",
                    "FieldTagName": "OmitEmpty",
                    "FieldTypeStr": "bool",
                    "FieldCategory": 23,
                    "FieldPrimitive": 18,
                    "FieldFullType": {
                        "Kind": 18,
                        "Str": "bool"
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Skip",
                    "FieldTagName": "Skip",
                    "FieldTypeStr": "bool",
                    "FieldCategory": 23,
                    "FieldPrimitive": 18,
                    "FieldFullType": {
                        "Kind": 18,
                        "Str": "bool"
                    },
                    "OmitEmpty": true
                },
                {
                    "Zid": -1,
                    "FieldGoName": "Deprecated",
                    "FieldTagName": "Deprecated",
                    "FieldTypeStr": "bool",
                    "FieldCategory": 23,
                    "FieldPrimitive": 18,
                    "FieldFullType": {
                        "Kind": 18,
                        "Str": "bool"
                    },
                    "OmitEmpty": true
                }
            ]
        }
    },
    "Imports": []
}`)
