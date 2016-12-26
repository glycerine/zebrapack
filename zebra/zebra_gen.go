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
	const maxFields0zzsz = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zzsz uint32
	totalEncodedFields0zzsz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zzsz := totalEncodedFields0zzsz
	missingFieldsLeft0zzsz := maxFields0zzsz - totalEncodedFields0zzsz

	var nextMiss0zzsz int32 = -1
	var found0zzsz [maxFields0zzsz]bool
	var curField0zzsz string

doneWithStruct0zzsz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zzsz > 0 || missingFieldsLeft0zzsz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zzsz, missingFieldsLeft0zzsz, msgp.ShowFound(found0zzsz[:]), decodeMsgFieldOrder0zzsz)
		if encodedFieldsLeft0zzsz > 0 {
			encodedFieldsLeft0zzsz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zzsz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zzsz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zzsz = 0
			}
			for nextMiss0zzsz < maxFields0zzsz && (found0zzsz[nextMiss0zzsz] || decodeMsgFieldSkip0zzsz[nextMiss0zzsz]) {
				nextMiss0zzsz++
			}
			if nextMiss0zzsz == maxFields0zzsz {
				// filled all the empty fields!
				break doneWithStruct0zzsz
			}
			missingFieldsLeft0zzsz--
			curField0zzsz = decodeMsgFieldOrder0zzsz[nextMiss0zzsz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zzsz)
		switch curField0zzsz {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zzsz[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zzsz[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zzsz[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zzsz[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zzsz[4] = true
			{
				var zsih uint64
				zsih, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zsih)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zzsz[5] = true
			{
				var zoig uint64
				zoig, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zoig)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zzsz[6] = true
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
			found0zzsz[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zzsz[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zzsz[9] = true
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
	if nextMiss0zzsz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zzsz = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zzsz = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zfgd [10]bool
	fieldsInUse_zbqw := z.fieldsNotEmpty(empty_zfgd[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zbqw)
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
	if !empty_zfgd[3] {
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

	if !empty_zfgd[4] {
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

	if !empty_zfgd[5] {
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

	if !empty_zfgd[6] {
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

	if !empty_zfgd[7] {
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

	if !empty_zfgd[8] {
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

	if !empty_zfgd[9] {
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
	const maxFields1zyux = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zyux uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zyux, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zyux := totalEncodedFields1zyux
	missingFieldsLeft1zyux := maxFields1zyux - totalEncodedFields1zyux

	var nextMiss1zyux int32 = -1
	var found1zyux [maxFields1zyux]bool
	var curField1zyux string

doneWithStruct1zyux:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zyux > 0 || missingFieldsLeft1zyux > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zyux, missingFieldsLeft1zyux, msgp.ShowFound(found1zyux[:]), unmarshalMsgFieldOrder1zyux)
		if encodedFieldsLeft1zyux > 0 {
			encodedFieldsLeft1zyux--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zyux = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zyux < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zyux = 0
			}
			for nextMiss1zyux < maxFields1zyux && (found1zyux[nextMiss1zyux] || unmarshalMsgFieldSkip1zyux[nextMiss1zyux]) {
				nextMiss1zyux++
			}
			if nextMiss1zyux == maxFields1zyux {
				// filled all the empty fields!
				break doneWithStruct1zyux
			}
			missingFieldsLeft1zyux--
			curField1zyux = unmarshalMsgFieldOrder1zyux[nextMiss1zyux]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zyux)
		switch curField1zyux {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zyux[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zyux[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zyux[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zyux[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zyux[4] = true
			{
				var zmyk uint64
				zmyk, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zmyk)
			}
		case "FieldPrimitive":
			found1zyux[5] = true
			{
				var zebr uint64
				zebr, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zebr)
			}
		case "FieldFullType":
			found1zyux[6] = true
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
			found1zyux[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zyux[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zyux[9] = true
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
	if nextMiss1zyux != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zyux = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zyux = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zecr = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zecr uint32
	totalEncodedFields2zecr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zecr := totalEncodedFields2zecr
	missingFieldsLeft2zecr := maxFields2zecr - totalEncodedFields2zecr

	var nextMiss2zecr int32 = -1
	var found2zecr [maxFields2zecr]bool
	var curField2zecr string

doneWithStruct2zecr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zecr > 0 || missingFieldsLeft2zecr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zecr, missingFieldsLeft2zecr, msgp.ShowFound(found2zecr[:]), decodeMsgFieldOrder2zecr)
		if encodedFieldsLeft2zecr > 0 {
			encodedFieldsLeft2zecr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zecr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zecr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zecr = 0
			}
			for nextMiss2zecr < maxFields2zecr && (found2zecr[nextMiss2zecr] || decodeMsgFieldSkip2zecr[nextMiss2zecr]) {
				nextMiss2zecr++
			}
			if nextMiss2zecr == maxFields2zecr {
				// filled all the empty fields!
				break doneWithStruct2zecr
			}
			missingFieldsLeft2zecr--
			curField2zecr = decodeMsgFieldOrder2zecr[nextMiss2zecr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zecr)
		switch curField2zecr {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zecr[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zecr[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zecr[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zecr[3] = true
			var zegl uint32
			zegl, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zegl) {
				z.Structs = (z.Structs)[:zegl]
			} else {
				z.Structs = make([]Struct, zegl)
			}
			for zfyi := range z.Structs {
				const maxFields3ztov = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3ztov uint32
				totalEncodedFields3ztov, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3ztov := totalEncodedFields3ztov
				missingFieldsLeft3ztov := maxFields3ztov - totalEncodedFields3ztov

				var nextMiss3ztov int32 = -1
				var found3ztov [maxFields3ztov]bool
				var curField3ztov string

			doneWithStruct3ztov:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3ztov > 0 || missingFieldsLeft3ztov > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3ztov, missingFieldsLeft3ztov, msgp.ShowFound(found3ztov[:]), decodeMsgFieldOrder3ztov)
					if encodedFieldsLeft3ztov > 0 {
						encodedFieldsLeft3ztov--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3ztov = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3ztov < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3ztov = 0
						}
						for nextMiss3ztov < maxFields3ztov && (found3ztov[nextMiss3ztov] || decodeMsgFieldSkip3ztov[nextMiss3ztov]) {
							nextMiss3ztov++
						}
						if nextMiss3ztov == maxFields3ztov {
							// filled all the empty fields!
							break doneWithStruct3ztov
						}
						missingFieldsLeft3ztov--
						curField3ztov = decodeMsgFieldOrder3ztov[nextMiss3ztov]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3ztov)
					switch curField3ztov {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3ztov[0] = true
						z.Structs[zfyi].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3ztov[1] = true
						var zjbo uint32
						zjbo, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zfyi].Fields) >= int(zjbo) {
							z.Structs[zfyi].Fields = (z.Structs[zfyi].Fields)[:zjbo]
						} else {
							z.Structs[zfyi].Fields = make([]Field, zjbo)
						}
						for zoox := range z.Structs[zfyi].Fields {
							err = z.Structs[zfyi].Fields[zoox].DecodeMsg(dc)
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
				if nextMiss3ztov != -1 {
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
	if nextMiss2zecr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zecr = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip2zecr = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3ztov = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3ztov = []bool{false, false}

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
	for zfyi := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zfyi].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zfyi].Fields)))
		if err != nil {
			panic(err)
		}
		for zoox := range z.Structs[zfyi].Fields {
			err = z.Structs[zfyi].Fields[zoox].EncodeMsg(en)
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
	for zfyi := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zfyi].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zfyi].Fields)))
		for zoox := range z.Structs[zfyi].Fields {
			o, err = z.Structs[zfyi].Fields[zoox].MarshalMsg(o)
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
	const maxFields4zzzi = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zzzi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zzzi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zzzi := totalEncodedFields4zzzi
	missingFieldsLeft4zzzi := maxFields4zzzi - totalEncodedFields4zzzi

	var nextMiss4zzzi int32 = -1
	var found4zzzi [maxFields4zzzi]bool
	var curField4zzzi string

doneWithStruct4zzzi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zzzi > 0 || missingFieldsLeft4zzzi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zzzi, missingFieldsLeft4zzzi, msgp.ShowFound(found4zzzi[:]), unmarshalMsgFieldOrder4zzzi)
		if encodedFieldsLeft4zzzi > 0 {
			encodedFieldsLeft4zzzi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zzzi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zzzi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zzzi = 0
			}
			for nextMiss4zzzi < maxFields4zzzi && (found4zzzi[nextMiss4zzzi] || unmarshalMsgFieldSkip4zzzi[nextMiss4zzzi]) {
				nextMiss4zzzi++
			}
			if nextMiss4zzzi == maxFields4zzzi {
				// filled all the empty fields!
				break doneWithStruct4zzzi
			}
			missingFieldsLeft4zzzi--
			curField4zzzi = unmarshalMsgFieldOrder4zzzi[nextMiss4zzzi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zzzi)
		switch curField4zzzi {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zzzi[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zzzi[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zzzi[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zzzi[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zlls uint32
				zlls, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zlls) {
					z.Structs = (z.Structs)[:zlls]
				} else {
					z.Structs = make([]Struct, zlls)
				}
				for zfyi := range z.Structs {
					const maxFields5zqxm = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zqxm uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zqxm, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zqxm := totalEncodedFields5zqxm
					missingFieldsLeft5zqxm := maxFields5zqxm - totalEncodedFields5zqxm

					var nextMiss5zqxm int32 = -1
					var found5zqxm [maxFields5zqxm]bool
					var curField5zqxm string

				doneWithStruct5zqxm:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zqxm > 0 || missingFieldsLeft5zqxm > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zqxm, missingFieldsLeft5zqxm, msgp.ShowFound(found5zqxm[:]), unmarshalMsgFieldOrder5zqxm)
						if encodedFieldsLeft5zqxm > 0 {
							encodedFieldsLeft5zqxm--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zqxm = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zqxm < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zqxm = 0
							}
							for nextMiss5zqxm < maxFields5zqxm && (found5zqxm[nextMiss5zqxm] || unmarshalMsgFieldSkip5zqxm[nextMiss5zqxm]) {
								nextMiss5zqxm++
							}
							if nextMiss5zqxm == maxFields5zqxm {
								// filled all the empty fields!
								break doneWithStruct5zqxm
							}
							missingFieldsLeft5zqxm--
							curField5zqxm = unmarshalMsgFieldOrder5zqxm[nextMiss5zqxm]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zqxm)
						switch curField5zqxm {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zqxm[0] = true
							z.Structs[zfyi].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zqxm[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zfyi].Fields) = (z.Structs[zfyi].Fields)[:0]
							} else {

								var zqzk uint32
								zqzk, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zfyi].Fields) >= int(zqzk) {
									z.Structs[zfyi].Fields = (z.Structs[zfyi].Fields)[:zqzk]
								} else {
									z.Structs[zfyi].Fields = make([]Field, zqzk)
								}
								for zoox := range z.Structs[zfyi].Fields {
									bts, err = z.Structs[zfyi].Fields[zoox].UnmarshalMsg(bts)
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
					if nextMiss5zqxm != -1 {
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
	if nextMiss4zzzi != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zzzi = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip4zzzi = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zqxm = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zqxm = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zfyi := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zfyi].StructName) + 7 + msgp.ArrayHeaderSize
		for zoox := range z.Structs[zfyi].Fields {
			s += z.Structs[zfyi].Fields[zoox].Msgsize()
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
	const maxFields6zwfo = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zwfo uint32
	totalEncodedFields6zwfo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zwfo := totalEncodedFields6zwfo
	missingFieldsLeft6zwfo := maxFields6zwfo - totalEncodedFields6zwfo

	var nextMiss6zwfo int32 = -1
	var found6zwfo [maxFields6zwfo]bool
	var curField6zwfo string

doneWithStruct6zwfo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zwfo > 0 || missingFieldsLeft6zwfo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zwfo, missingFieldsLeft6zwfo, msgp.ShowFound(found6zwfo[:]), decodeMsgFieldOrder6zwfo)
		if encodedFieldsLeft6zwfo > 0 {
			encodedFieldsLeft6zwfo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zwfo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zwfo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zwfo = 0
			}
			for nextMiss6zwfo < maxFields6zwfo && (found6zwfo[nextMiss6zwfo] || decodeMsgFieldSkip6zwfo[nextMiss6zwfo]) {
				nextMiss6zwfo++
			}
			if nextMiss6zwfo == maxFields6zwfo {
				// filled all the empty fields!
				break doneWithStruct6zwfo
			}
			missingFieldsLeft6zwfo--
			curField6zwfo = decodeMsgFieldOrder6zwfo[nextMiss6zwfo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zwfo)
		switch curField6zwfo {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zwfo[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zwfo[1] = true
			var zekl uint32
			zekl, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zekl) {
				z.Fields = (z.Fields)[:zekl]
			} else {
				z.Fields = make([]Field, zekl)
			}
			for zkdb := range z.Fields {
				err = z.Fields[zkdb].DecodeMsg(dc)
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
	if nextMiss6zwfo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zwfo = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zwfo = []bool{false, false}

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
	for zkdb := range z.Fields {
		err = z.Fields[zkdb].EncodeMsg(en)
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
	for zkdb := range z.Fields {
		o, err = z.Fields[zkdb].MarshalMsg(o)
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
	const maxFields7zifm = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zifm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zifm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zifm := totalEncodedFields7zifm
	missingFieldsLeft7zifm := maxFields7zifm - totalEncodedFields7zifm

	var nextMiss7zifm int32 = -1
	var found7zifm [maxFields7zifm]bool
	var curField7zifm string

doneWithStruct7zifm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zifm > 0 || missingFieldsLeft7zifm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zifm, missingFieldsLeft7zifm, msgp.ShowFound(found7zifm[:]), unmarshalMsgFieldOrder7zifm)
		if encodedFieldsLeft7zifm > 0 {
			encodedFieldsLeft7zifm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zifm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zifm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zifm = 0
			}
			for nextMiss7zifm < maxFields7zifm && (found7zifm[nextMiss7zifm] || unmarshalMsgFieldSkip7zifm[nextMiss7zifm]) {
				nextMiss7zifm++
			}
			if nextMiss7zifm == maxFields7zifm {
				// filled all the empty fields!
				break doneWithStruct7zifm
			}
			missingFieldsLeft7zifm--
			curField7zifm = unmarshalMsgFieldOrder7zifm[nextMiss7zifm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zifm)
		switch curField7zifm {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zifm[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zifm[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zgsg uint32
				zgsg, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zgsg) {
					z.Fields = (z.Fields)[:zgsg]
				} else {
					z.Fields = make([]Field, zgsg)
				}
				for zkdb := range z.Fields {
					bts, err = z.Fields[zkdb].UnmarshalMsg(bts)
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
	if nextMiss7zifm != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zifm = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zifm = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zkdb := range z.Fields {
		s += z.Fields[zkdb].Msgsize()
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
		var znlo uint64
		znlo, err = dc.ReadUint64()
		(*z) = Zkind(znlo)
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
		var zlhy uint64
		zlhy, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zlhy)
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
	const maxFields8zbdl = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zbdl uint32
	totalEncodedFields8zbdl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zbdl := totalEncodedFields8zbdl
	missingFieldsLeft8zbdl := maxFields8zbdl - totalEncodedFields8zbdl

	var nextMiss8zbdl int32 = -1
	var found8zbdl [maxFields8zbdl]bool
	var curField8zbdl string

doneWithStruct8zbdl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zbdl > 0 || missingFieldsLeft8zbdl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zbdl, missingFieldsLeft8zbdl, msgp.ShowFound(found8zbdl[:]), decodeMsgFieldOrder8zbdl)
		if encodedFieldsLeft8zbdl > 0 {
			encodedFieldsLeft8zbdl--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zbdl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zbdl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zbdl = 0
			}
			for nextMiss8zbdl < maxFields8zbdl && (found8zbdl[nextMiss8zbdl] || decodeMsgFieldSkip8zbdl[nextMiss8zbdl]) {
				nextMiss8zbdl++
			}
			if nextMiss8zbdl == maxFields8zbdl {
				// filled all the empty fields!
				break doneWithStruct8zbdl
			}
			missingFieldsLeft8zbdl--
			curField8zbdl = decodeMsgFieldOrder8zbdl[nextMiss8zbdl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zbdl)
		switch curField8zbdl {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zbdl[0] = true
			{
				var zfmt uint64
				zfmt, err = dc.ReadUint64()
				z.Kind = Zkind(zfmt)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zbdl[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zbdl[2] = true
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
			found8zbdl[3] = true
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
	if nextMiss8zbdl != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zbdl = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zbdl = []bool{false, false, false, false}

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
	var empty_zmhl [4]bool
	fieldsInUse_zhvt := z.fieldsNotEmpty(empty_zmhl[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zhvt)
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
	if !empty_zmhl[1] {
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

	if !empty_zmhl[2] {
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

	if !empty_zmhl[3] {
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
	const maxFields9ztis = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9ztis uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9ztis, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9ztis := totalEncodedFields9ztis
	missingFieldsLeft9ztis := maxFields9ztis - totalEncodedFields9ztis

	var nextMiss9ztis int32 = -1
	var found9ztis [maxFields9ztis]bool
	var curField9ztis string

doneWithStruct9ztis:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9ztis > 0 || missingFieldsLeft9ztis > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9ztis, missingFieldsLeft9ztis, msgp.ShowFound(found9ztis[:]), unmarshalMsgFieldOrder9ztis)
		if encodedFieldsLeft9ztis > 0 {
			encodedFieldsLeft9ztis--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9ztis = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9ztis < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9ztis = 0
			}
			for nextMiss9ztis < maxFields9ztis && (found9ztis[nextMiss9ztis] || unmarshalMsgFieldSkip9ztis[nextMiss9ztis]) {
				nextMiss9ztis++
			}
			if nextMiss9ztis == maxFields9ztis {
				// filled all the empty fields!
				break doneWithStruct9ztis
			}
			missingFieldsLeft9ztis--
			curField9ztis = unmarshalMsgFieldOrder9ztis[nextMiss9ztis]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9ztis)
		switch curField9ztis {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9ztis[0] = true
			{
				var zxiu uint64
				zxiu, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zxiu)
			}
		case "Str":
			found9ztis[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9ztis[2] = true
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
			found9ztis[3] = true
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
	if nextMiss9ztis != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9ztis = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9ztis = []bool{false, false, false, false}

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
