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
	const maxFields0zycw = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zycw uint32
	totalEncodedFields0zycw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zycw := totalEncodedFields0zycw
	missingFieldsLeft0zycw := maxFields0zycw - totalEncodedFields0zycw

	var nextMiss0zycw int32 = -1
	var found0zycw [maxFields0zycw]bool
	var curField0zycw string

doneWithStruct0zycw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zycw > 0 || missingFieldsLeft0zycw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zycw, missingFieldsLeft0zycw, msgp.ShowFound(found0zycw[:]), decodeMsgFieldOrder0zycw)
		if encodedFieldsLeft0zycw > 0 {
			encodedFieldsLeft0zycw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zycw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zycw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zycw = 0
			}
			for nextMiss0zycw < maxFields0zycw && (found0zycw[nextMiss0zycw] || decodeMsgFieldSkip0zycw[nextMiss0zycw]) {
				nextMiss0zycw++
			}
			if nextMiss0zycw == maxFields0zycw {
				// filled all the empty fields!
				break doneWithStruct0zycw
			}
			missingFieldsLeft0zycw--
			curField0zycw = decodeMsgFieldOrder0zycw[nextMiss0zycw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zycw)
		switch curField0zycw {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zycw[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zycw[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zycw[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zycw[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zycw[4] = true
			{
				var zdbv uint64
				zdbv, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zdbv)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zycw[5] = true
			{
				var zmsj uint64
				zmsj, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zmsj)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zycw[6] = true
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
			found0zycw[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zycw[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zycw[9] = true
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
	if nextMiss0zycw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zycw = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zycw = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zwld [10]bool
	fieldsInUse_zcjb := z.fieldsNotEmpty(empty_zwld[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zcjb)
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
	if !empty_zwld[3] {
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

	if !empty_zwld[4] {
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

	if !empty_zwld[5] {
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

	if !empty_zwld[6] {
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

	if !empty_zwld[7] {
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

	if !empty_zwld[8] {
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

	if !empty_zwld[9] {
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
	const maxFields1zkum = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zkum uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zkum, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zkum := totalEncodedFields1zkum
	missingFieldsLeft1zkum := maxFields1zkum - totalEncodedFields1zkum

	var nextMiss1zkum int32 = -1
	var found1zkum [maxFields1zkum]bool
	var curField1zkum string

doneWithStruct1zkum:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zkum > 0 || missingFieldsLeft1zkum > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zkum, missingFieldsLeft1zkum, msgp.ShowFound(found1zkum[:]), unmarshalMsgFieldOrder1zkum)
		if encodedFieldsLeft1zkum > 0 {
			encodedFieldsLeft1zkum--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zkum = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zkum < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zkum = 0
			}
			for nextMiss1zkum < maxFields1zkum && (found1zkum[nextMiss1zkum] || unmarshalMsgFieldSkip1zkum[nextMiss1zkum]) {
				nextMiss1zkum++
			}
			if nextMiss1zkum == maxFields1zkum {
				// filled all the empty fields!
				break doneWithStruct1zkum
			}
			missingFieldsLeft1zkum--
			curField1zkum = unmarshalMsgFieldOrder1zkum[nextMiss1zkum]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zkum)
		switch curField1zkum {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zkum[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zkum[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zkum[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zkum[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zkum[4] = true
			{
				var zjop uint64
				zjop, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zjop)
			}
		case "FieldPrimitive":
			found1zkum[5] = true
			{
				var zbsl uint64
				zbsl, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zbsl)
			}
		case "FieldFullType":
			found1zkum[6] = true
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
			found1zkum[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zkum[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zkum[9] = true
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
	if nextMiss1zkum != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zkum = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zkum = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2znje = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2znje uint32
	totalEncodedFields2znje, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2znje := totalEncodedFields2znje
	missingFieldsLeft2znje := maxFields2znje - totalEncodedFields2znje

	var nextMiss2znje int32 = -1
	var found2znje [maxFields2znje]bool
	var curField2znje string

doneWithStruct2znje:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2znje > 0 || missingFieldsLeft2znje > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2znje, missingFieldsLeft2znje, msgp.ShowFound(found2znje[:]), decodeMsgFieldOrder2znje)
		if encodedFieldsLeft2znje > 0 {
			encodedFieldsLeft2znje--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2znje = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2znje < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2znje = 0
			}
			for nextMiss2znje < maxFields2znje && (found2znje[nextMiss2znje] || decodeMsgFieldSkip2znje[nextMiss2znje]) {
				nextMiss2znje++
			}
			if nextMiss2znje == maxFields2znje {
				// filled all the empty fields!
				break doneWithStruct2znje
			}
			missingFieldsLeft2znje--
			curField2znje = decodeMsgFieldOrder2znje[nextMiss2znje]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2znje)
		switch curField2znje {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2znje[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2znje[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2znje[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2znje[3] = true
			var zoxy uint32
			zoxy, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zoxy) {
				z.Structs = (z.Structs)[:zoxy]
			} else {
				z.Structs = make([]Struct, zoxy)
			}
			for zmkb := range z.Structs {
				const maxFields3zzyy = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zzyy uint32
				totalEncodedFields3zzyy, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zzyy := totalEncodedFields3zzyy
				missingFieldsLeft3zzyy := maxFields3zzyy - totalEncodedFields3zzyy

				var nextMiss3zzyy int32 = -1
				var found3zzyy [maxFields3zzyy]bool
				var curField3zzyy string

			doneWithStruct3zzyy:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zzyy > 0 || missingFieldsLeft3zzyy > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zzyy, missingFieldsLeft3zzyy, msgp.ShowFound(found3zzyy[:]), decodeMsgFieldOrder3zzyy)
					if encodedFieldsLeft3zzyy > 0 {
						encodedFieldsLeft3zzyy--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zzyy = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zzyy < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zzyy = 0
						}
						for nextMiss3zzyy < maxFields3zzyy && (found3zzyy[nextMiss3zzyy] || decodeMsgFieldSkip3zzyy[nextMiss3zzyy]) {
							nextMiss3zzyy++
						}
						if nextMiss3zzyy == maxFields3zzyy {
							// filled all the empty fields!
							break doneWithStruct3zzyy
						}
						missingFieldsLeft3zzyy--
						curField3zzyy = decodeMsgFieldOrder3zzyy[nextMiss3zzyy]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zzyy)
					switch curField3zzyy {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zzyy[0] = true
						z.Structs[zmkb].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zzyy[1] = true
						var zozj uint32
						zozj, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zmkb].Fields) >= int(zozj) {
							z.Structs[zmkb].Fields = (z.Structs[zmkb].Fields)[:zozj]
						} else {
							z.Structs[zmkb].Fields = make([]Field, zozj)
						}
						for zpno := range z.Structs[zmkb].Fields {
							err = z.Structs[zmkb].Fields[zpno].DecodeMsg(dc)
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
				if nextMiss3zzyy != -1 {
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
	if nextMiss2znje != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2znje = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip2znje = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zzyy = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zzyy = []bool{false, false}

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
	for zmkb := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zmkb].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zmkb].Fields)))
		if err != nil {
			panic(err)
		}
		for zpno := range z.Structs[zmkb].Fields {
			err = z.Structs[zmkb].Fields[zpno].EncodeMsg(en)
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
	for zmkb := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zmkb].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zmkb].Fields)))
		for zpno := range z.Structs[zmkb].Fields {
			o, err = z.Structs[zmkb].Fields[zpno].MarshalMsg(o)
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
	const maxFields4zrgq = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zrgq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zrgq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zrgq := totalEncodedFields4zrgq
	missingFieldsLeft4zrgq := maxFields4zrgq - totalEncodedFields4zrgq

	var nextMiss4zrgq int32 = -1
	var found4zrgq [maxFields4zrgq]bool
	var curField4zrgq string

doneWithStruct4zrgq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zrgq > 0 || missingFieldsLeft4zrgq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zrgq, missingFieldsLeft4zrgq, msgp.ShowFound(found4zrgq[:]), unmarshalMsgFieldOrder4zrgq)
		if encodedFieldsLeft4zrgq > 0 {
			encodedFieldsLeft4zrgq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zrgq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zrgq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zrgq = 0
			}
			for nextMiss4zrgq < maxFields4zrgq && (found4zrgq[nextMiss4zrgq] || unmarshalMsgFieldSkip4zrgq[nextMiss4zrgq]) {
				nextMiss4zrgq++
			}
			if nextMiss4zrgq == maxFields4zrgq {
				// filled all the empty fields!
				break doneWithStruct4zrgq
			}
			missingFieldsLeft4zrgq--
			curField4zrgq = unmarshalMsgFieldOrder4zrgq[nextMiss4zrgq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zrgq)
		switch curField4zrgq {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zrgq[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zrgq[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zrgq[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zrgq[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var ztcx uint32
				ztcx, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(ztcx) {
					z.Structs = (z.Structs)[:ztcx]
				} else {
					z.Structs = make([]Struct, ztcx)
				}
				for zmkb := range z.Structs {
					const maxFields5zojp = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zojp uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zojp, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zojp := totalEncodedFields5zojp
					missingFieldsLeft5zojp := maxFields5zojp - totalEncodedFields5zojp

					var nextMiss5zojp int32 = -1
					var found5zojp [maxFields5zojp]bool
					var curField5zojp string

				doneWithStruct5zojp:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zojp > 0 || missingFieldsLeft5zojp > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zojp, missingFieldsLeft5zojp, msgp.ShowFound(found5zojp[:]), unmarshalMsgFieldOrder5zojp)
						if encodedFieldsLeft5zojp > 0 {
							encodedFieldsLeft5zojp--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zojp = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zojp < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zojp = 0
							}
							for nextMiss5zojp < maxFields5zojp && (found5zojp[nextMiss5zojp] || unmarshalMsgFieldSkip5zojp[nextMiss5zojp]) {
								nextMiss5zojp++
							}
							if nextMiss5zojp == maxFields5zojp {
								// filled all the empty fields!
								break doneWithStruct5zojp
							}
							missingFieldsLeft5zojp--
							curField5zojp = unmarshalMsgFieldOrder5zojp[nextMiss5zojp]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zojp)
						switch curField5zojp {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zojp[0] = true
							z.Structs[zmkb].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zojp[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zmkb].Fields) = (z.Structs[zmkb].Fields)[:0]
							} else {

								var zziy uint32
								zziy, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zmkb].Fields) >= int(zziy) {
									z.Structs[zmkb].Fields = (z.Structs[zmkb].Fields)[:zziy]
								} else {
									z.Structs[zmkb].Fields = make([]Field, zziy)
								}
								for zpno := range z.Structs[zmkb].Fields {
									bts, err = z.Structs[zmkb].Fields[zpno].UnmarshalMsg(bts)
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
					if nextMiss5zojp != -1 {
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
	if nextMiss4zrgq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zrgq = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip4zrgq = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zojp = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zojp = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zmkb := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zmkb].StructName) + 7 + msgp.ArrayHeaderSize
		for zpno := range z.Structs[zmkb].Fields {
			s += z.Structs[zmkb].Fields[zpno].Msgsize()
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
	const maxFields6zxvw = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zxvw uint32
	totalEncodedFields6zxvw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zxvw := totalEncodedFields6zxvw
	missingFieldsLeft6zxvw := maxFields6zxvw - totalEncodedFields6zxvw

	var nextMiss6zxvw int32 = -1
	var found6zxvw [maxFields6zxvw]bool
	var curField6zxvw string

doneWithStruct6zxvw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zxvw > 0 || missingFieldsLeft6zxvw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zxvw, missingFieldsLeft6zxvw, msgp.ShowFound(found6zxvw[:]), decodeMsgFieldOrder6zxvw)
		if encodedFieldsLeft6zxvw > 0 {
			encodedFieldsLeft6zxvw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zxvw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zxvw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zxvw = 0
			}
			for nextMiss6zxvw < maxFields6zxvw && (found6zxvw[nextMiss6zxvw] || decodeMsgFieldSkip6zxvw[nextMiss6zxvw]) {
				nextMiss6zxvw++
			}
			if nextMiss6zxvw == maxFields6zxvw {
				// filled all the empty fields!
				break doneWithStruct6zxvw
			}
			missingFieldsLeft6zxvw--
			curField6zxvw = decodeMsgFieldOrder6zxvw[nextMiss6zxvw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zxvw)
		switch curField6zxvw {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zxvw[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zxvw[1] = true
			var znon uint32
			znon, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(znon) {
				z.Fields = (z.Fields)[:znon]
			} else {
				z.Fields = make([]Field, znon)
			}
			for zrdf := range z.Fields {
				err = z.Fields[zrdf].DecodeMsg(dc)
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
	if nextMiss6zxvw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zxvw = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zxvw = []bool{false, false}

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
	for zrdf := range z.Fields {
		err = z.Fields[zrdf].EncodeMsg(en)
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
	for zrdf := range z.Fields {
		o, err = z.Fields[zrdf].MarshalMsg(o)
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
	const maxFields7zstn = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zstn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zstn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zstn := totalEncodedFields7zstn
	missingFieldsLeft7zstn := maxFields7zstn - totalEncodedFields7zstn

	var nextMiss7zstn int32 = -1
	var found7zstn [maxFields7zstn]bool
	var curField7zstn string

doneWithStruct7zstn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zstn > 0 || missingFieldsLeft7zstn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zstn, missingFieldsLeft7zstn, msgp.ShowFound(found7zstn[:]), unmarshalMsgFieldOrder7zstn)
		if encodedFieldsLeft7zstn > 0 {
			encodedFieldsLeft7zstn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zstn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zstn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zstn = 0
			}
			for nextMiss7zstn < maxFields7zstn && (found7zstn[nextMiss7zstn] || unmarshalMsgFieldSkip7zstn[nextMiss7zstn]) {
				nextMiss7zstn++
			}
			if nextMiss7zstn == maxFields7zstn {
				// filled all the empty fields!
				break doneWithStruct7zstn
			}
			missingFieldsLeft7zstn--
			curField7zstn = unmarshalMsgFieldOrder7zstn[nextMiss7zstn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zstn)
		switch curField7zstn {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zstn[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zstn[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zzkh uint32
				zzkh, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zzkh) {
					z.Fields = (z.Fields)[:zzkh]
				} else {
					z.Fields = make([]Field, zzkh)
				}
				for zrdf := range z.Fields {
					bts, err = z.Fields[zrdf].UnmarshalMsg(bts)
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
	if nextMiss7zstn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zstn = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zstn = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zrdf := range z.Fields {
		s += z.Fields[zrdf].Msgsize()
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
		var zmns uint64
		zmns, err = dc.ReadUint64()
		(*z) = Zkind(zmns)
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
		var zzvk uint64
		zzvk, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zzvk)
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
	const maxFields8zmea = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zmea uint32
	totalEncodedFields8zmea, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zmea := totalEncodedFields8zmea
	missingFieldsLeft8zmea := maxFields8zmea - totalEncodedFields8zmea

	var nextMiss8zmea int32 = -1
	var found8zmea [maxFields8zmea]bool
	var curField8zmea string

doneWithStruct8zmea:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zmea > 0 || missingFieldsLeft8zmea > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zmea, missingFieldsLeft8zmea, msgp.ShowFound(found8zmea[:]), decodeMsgFieldOrder8zmea)
		if encodedFieldsLeft8zmea > 0 {
			encodedFieldsLeft8zmea--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zmea = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zmea < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zmea = 0
			}
			for nextMiss8zmea < maxFields8zmea && (found8zmea[nextMiss8zmea] || decodeMsgFieldSkip8zmea[nextMiss8zmea]) {
				nextMiss8zmea++
			}
			if nextMiss8zmea == maxFields8zmea {
				// filled all the empty fields!
				break doneWithStruct8zmea
			}
			missingFieldsLeft8zmea--
			curField8zmea = decodeMsgFieldOrder8zmea[nextMiss8zmea]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zmea)
		switch curField8zmea {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zmea[0] = true
			{
				var zhrz uint64
				zhrz, err = dc.ReadUint64()
				z.Kind = Zkind(zhrz)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zmea[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zmea[2] = true
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
			found8zmea[3] = true
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
	if nextMiss8zmea != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zmea = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zmea = []bool{false, false, false, false}

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
	var empty_zjbp [4]bool
	fieldsInUse_zlat := z.fieldsNotEmpty(empty_zjbp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlat)
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
	if !empty_zjbp[1] {
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

	if !empty_zjbp[2] {
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

	if !empty_zjbp[3] {
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
	const maxFields9zdvb = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zdvb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zdvb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zdvb := totalEncodedFields9zdvb
	missingFieldsLeft9zdvb := maxFields9zdvb - totalEncodedFields9zdvb

	var nextMiss9zdvb int32 = -1
	var found9zdvb [maxFields9zdvb]bool
	var curField9zdvb string

doneWithStruct9zdvb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zdvb > 0 || missingFieldsLeft9zdvb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zdvb, missingFieldsLeft9zdvb, msgp.ShowFound(found9zdvb[:]), unmarshalMsgFieldOrder9zdvb)
		if encodedFieldsLeft9zdvb > 0 {
			encodedFieldsLeft9zdvb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zdvb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zdvb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zdvb = 0
			}
			for nextMiss9zdvb < maxFields9zdvb && (found9zdvb[nextMiss9zdvb] || unmarshalMsgFieldSkip9zdvb[nextMiss9zdvb]) {
				nextMiss9zdvb++
			}
			if nextMiss9zdvb == maxFields9zdvb {
				// filled all the empty fields!
				break doneWithStruct9zdvb
			}
			missingFieldsLeft9zdvb--
			curField9zdvb = unmarshalMsgFieldOrder9zdvb[nextMiss9zdvb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zdvb)
		switch curField9zdvb {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zdvb[0] = true
			{
				var zgxk uint64
				zgxk, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zgxk)
			}
		case "Str":
			found9zdvb[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zdvb[2] = true
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
			found9zdvb[3] = true
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
	if nextMiss9zdvb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zdvb = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zdvb = []bool{false, false, false, false}

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
