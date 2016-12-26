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
	const maxFields0zums = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zums uint32
	totalEncodedFields0zums, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zums := totalEncodedFields0zums
	missingFieldsLeft0zums := maxFields0zums - totalEncodedFields0zums

	var nextMiss0zums int32 = -1
	var found0zums [maxFields0zums]bool
	var curField0zums string

doneWithStruct0zums:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zums > 0 || missingFieldsLeft0zums > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zums, missingFieldsLeft0zums, msgp.ShowFound(found0zums[:]), decodeMsgFieldOrder0zums)
		if encodedFieldsLeft0zums > 0 {
			encodedFieldsLeft0zums--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zums = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zums < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zums = 0
			}
			for nextMiss0zums < maxFields0zums && (found0zums[nextMiss0zums] || decodeMsgFieldSkip0zums[nextMiss0zums]) {
				nextMiss0zums++
			}
			if nextMiss0zums == maxFields0zums {
				// filled all the empty fields!
				break doneWithStruct0zums
			}
			missingFieldsLeft0zums--
			curField0zums = decodeMsgFieldOrder0zums[nextMiss0zums]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zums)
		switch curField0zums {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zums[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zums[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zums[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zums[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zums[4] = true
			{
				var zrfc uint64
				zrfc, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zrfc)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zums[5] = true
			{
				var zogo uint64
				zogo, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zogo)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zums[6] = true
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
			found0zums[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zums[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zums[9] = true
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
	if nextMiss0zums != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zums = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zums = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_ztfq [10]bool
	fieldsInUse_zybw := z.fieldsNotEmpty(empty_ztfq[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zybw)
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
	if !empty_ztfq[3] {
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

	if !empty_ztfq[4] {
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

	if !empty_ztfq[5] {
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

	if !empty_ztfq[6] {
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

	if !empty_ztfq[7] {
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

	if !empty_ztfq[8] {
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

	if !empty_ztfq[9] {
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
	const maxFields1zugj = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zugj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zugj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zugj := totalEncodedFields1zugj
	missingFieldsLeft1zugj := maxFields1zugj - totalEncodedFields1zugj

	var nextMiss1zugj int32 = -1
	var found1zugj [maxFields1zugj]bool
	var curField1zugj string

doneWithStruct1zugj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zugj > 0 || missingFieldsLeft1zugj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zugj, missingFieldsLeft1zugj, msgp.ShowFound(found1zugj[:]), unmarshalMsgFieldOrder1zugj)
		if encodedFieldsLeft1zugj > 0 {
			encodedFieldsLeft1zugj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zugj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zugj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zugj = 0
			}
			for nextMiss1zugj < maxFields1zugj && (found1zugj[nextMiss1zugj] || unmarshalMsgFieldSkip1zugj[nextMiss1zugj]) {
				nextMiss1zugj++
			}
			if nextMiss1zugj == maxFields1zugj {
				// filled all the empty fields!
				break doneWithStruct1zugj
			}
			missingFieldsLeft1zugj--
			curField1zugj = unmarshalMsgFieldOrder1zugj[nextMiss1zugj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zugj)
		switch curField1zugj {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zugj[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zugj[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zugj[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zugj[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zugj[4] = true
			{
				var zudf uint64
				zudf, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zudf)
			}
		case "FieldPrimitive":
			found1zugj[5] = true
			{
				var zqiv uint64
				zqiv, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zqiv)
			}
		case "FieldFullType":
			found1zugj[6] = true
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
			found1zugj[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zugj[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zugj[9] = true
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
	if nextMiss1zugj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zugj = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zugj = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zpcc = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zpcc uint32
	totalEncodedFields2zpcc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zpcc := totalEncodedFields2zpcc
	missingFieldsLeft2zpcc := maxFields2zpcc - totalEncodedFields2zpcc

	var nextMiss2zpcc int32 = -1
	var found2zpcc [maxFields2zpcc]bool
	var curField2zpcc string

doneWithStruct2zpcc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zpcc > 0 || missingFieldsLeft2zpcc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zpcc, missingFieldsLeft2zpcc, msgp.ShowFound(found2zpcc[:]), decodeMsgFieldOrder2zpcc)
		if encodedFieldsLeft2zpcc > 0 {
			encodedFieldsLeft2zpcc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zpcc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zpcc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zpcc = 0
			}
			for nextMiss2zpcc < maxFields2zpcc && (found2zpcc[nextMiss2zpcc] || decodeMsgFieldSkip2zpcc[nextMiss2zpcc]) {
				nextMiss2zpcc++
			}
			if nextMiss2zpcc == maxFields2zpcc {
				// filled all the empty fields!
				break doneWithStruct2zpcc
			}
			missingFieldsLeft2zpcc--
			curField2zpcc = decodeMsgFieldOrder2zpcc[nextMiss2zpcc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zpcc)
		switch curField2zpcc {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zpcc[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zpcc[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zpcc[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zpcc[3] = true
			var zljg uint32
			zljg, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zljg) {
				z.Structs = (z.Structs)[:zljg]
			} else {
				z.Structs = make([]Struct, zljg)
			}
			for zmry := range z.Structs {
				const maxFields3zjbh = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zjbh uint32
				totalEncodedFields3zjbh, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zjbh := totalEncodedFields3zjbh
				missingFieldsLeft3zjbh := maxFields3zjbh - totalEncodedFields3zjbh

				var nextMiss3zjbh int32 = -1
				var found3zjbh [maxFields3zjbh]bool
				var curField3zjbh string

			doneWithStruct3zjbh:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zjbh > 0 || missingFieldsLeft3zjbh > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zjbh, missingFieldsLeft3zjbh, msgp.ShowFound(found3zjbh[:]), decodeMsgFieldOrder3zjbh)
					if encodedFieldsLeft3zjbh > 0 {
						encodedFieldsLeft3zjbh--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zjbh = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zjbh < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zjbh = 0
						}
						for nextMiss3zjbh < maxFields3zjbh && (found3zjbh[nextMiss3zjbh] || decodeMsgFieldSkip3zjbh[nextMiss3zjbh]) {
							nextMiss3zjbh++
						}
						if nextMiss3zjbh == maxFields3zjbh {
							// filled all the empty fields!
							break doneWithStruct3zjbh
						}
						missingFieldsLeft3zjbh--
						curField3zjbh = decodeMsgFieldOrder3zjbh[nextMiss3zjbh]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zjbh)
					switch curField3zjbh {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zjbh[0] = true
						z.Structs[zmry].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zjbh[1] = true
						var zllf uint32
						zllf, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zmry].Fields) >= int(zllf) {
							z.Structs[zmry].Fields = (z.Structs[zmry].Fields)[:zllf]
						} else {
							z.Structs[zmry].Fields = make([]Field, zllf)
						}
						for zodo := range z.Structs[zmry].Fields {
							err = z.Structs[zmry].Fields[zodo].DecodeMsg(dc)
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
				if nextMiss3zjbh != -1 {
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
	if nextMiss2zpcc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zpcc = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip2zpcc = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zjbh = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zjbh = []bool{false, false}

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
	for zmry := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zmry].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zmry].Fields)))
		if err != nil {
			panic(err)
		}
		for zodo := range z.Structs[zmry].Fields {
			err = z.Structs[zmry].Fields[zodo].EncodeMsg(en)
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
	for zmry := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zmry].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zmry].Fields)))
		for zodo := range z.Structs[zmry].Fields {
			o, err = z.Structs[zmry].Fields[zodo].MarshalMsg(o)
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
	const maxFields4zeqs = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zeqs uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zeqs, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zeqs := totalEncodedFields4zeqs
	missingFieldsLeft4zeqs := maxFields4zeqs - totalEncodedFields4zeqs

	var nextMiss4zeqs int32 = -1
	var found4zeqs [maxFields4zeqs]bool
	var curField4zeqs string

doneWithStruct4zeqs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zeqs > 0 || missingFieldsLeft4zeqs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zeqs, missingFieldsLeft4zeqs, msgp.ShowFound(found4zeqs[:]), unmarshalMsgFieldOrder4zeqs)
		if encodedFieldsLeft4zeqs > 0 {
			encodedFieldsLeft4zeqs--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zeqs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zeqs < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zeqs = 0
			}
			for nextMiss4zeqs < maxFields4zeqs && (found4zeqs[nextMiss4zeqs] || unmarshalMsgFieldSkip4zeqs[nextMiss4zeqs]) {
				nextMiss4zeqs++
			}
			if nextMiss4zeqs == maxFields4zeqs {
				// filled all the empty fields!
				break doneWithStruct4zeqs
			}
			missingFieldsLeft4zeqs--
			curField4zeqs = unmarshalMsgFieldOrder4zeqs[nextMiss4zeqs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zeqs)
		switch curField4zeqs {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zeqs[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zeqs[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zeqs[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zeqs[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zpze uint32
				zpze, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zpze) {
					z.Structs = (z.Structs)[:zpze]
				} else {
					z.Structs = make([]Struct, zpze)
				}
				for zmry := range z.Structs {
					const maxFields5zqww = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zqww uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zqww, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zqww := totalEncodedFields5zqww
					missingFieldsLeft5zqww := maxFields5zqww - totalEncodedFields5zqww

					var nextMiss5zqww int32 = -1
					var found5zqww [maxFields5zqww]bool
					var curField5zqww string

				doneWithStruct5zqww:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zqww > 0 || missingFieldsLeft5zqww > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zqww, missingFieldsLeft5zqww, msgp.ShowFound(found5zqww[:]), unmarshalMsgFieldOrder5zqww)
						if encodedFieldsLeft5zqww > 0 {
							encodedFieldsLeft5zqww--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zqww = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zqww < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zqww = 0
							}
							for nextMiss5zqww < maxFields5zqww && (found5zqww[nextMiss5zqww] || unmarshalMsgFieldSkip5zqww[nextMiss5zqww]) {
								nextMiss5zqww++
							}
							if nextMiss5zqww == maxFields5zqww {
								// filled all the empty fields!
								break doneWithStruct5zqww
							}
							missingFieldsLeft5zqww--
							curField5zqww = unmarshalMsgFieldOrder5zqww[nextMiss5zqww]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zqww)
						switch curField5zqww {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zqww[0] = true
							z.Structs[zmry].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zqww[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zmry].Fields) = (z.Structs[zmry].Fields)[:0]
							} else {

								var zybi uint32
								zybi, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zmry].Fields) >= int(zybi) {
									z.Structs[zmry].Fields = (z.Structs[zmry].Fields)[:zybi]
								} else {
									z.Structs[zmry].Fields = make([]Field, zybi)
								}
								for zodo := range z.Structs[zmry].Fields {
									bts, err = z.Structs[zmry].Fields[zodo].UnmarshalMsg(bts)
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
					if nextMiss5zqww != -1 {
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
	if nextMiss4zeqs != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zeqs = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip4zeqs = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zqww = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zqww = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zmry := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zmry].StructName) + 7 + msgp.ArrayHeaderSize
		for zodo := range z.Structs[zmry].Fields {
			s += z.Structs[zmry].Fields[zodo].Msgsize()
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
	const maxFields6zbry = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zbry uint32
	totalEncodedFields6zbry, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zbry := totalEncodedFields6zbry
	missingFieldsLeft6zbry := maxFields6zbry - totalEncodedFields6zbry

	var nextMiss6zbry int32 = -1
	var found6zbry [maxFields6zbry]bool
	var curField6zbry string

doneWithStruct6zbry:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zbry > 0 || missingFieldsLeft6zbry > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zbry, missingFieldsLeft6zbry, msgp.ShowFound(found6zbry[:]), decodeMsgFieldOrder6zbry)
		if encodedFieldsLeft6zbry > 0 {
			encodedFieldsLeft6zbry--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zbry = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zbry < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zbry = 0
			}
			for nextMiss6zbry < maxFields6zbry && (found6zbry[nextMiss6zbry] || decodeMsgFieldSkip6zbry[nextMiss6zbry]) {
				nextMiss6zbry++
			}
			if nextMiss6zbry == maxFields6zbry {
				// filled all the empty fields!
				break doneWithStruct6zbry
			}
			missingFieldsLeft6zbry--
			curField6zbry = decodeMsgFieldOrder6zbry[nextMiss6zbry]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zbry)
		switch curField6zbry {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zbry[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zbry[1] = true
			var zcnk uint32
			zcnk, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zcnk) {
				z.Fields = (z.Fields)[:zcnk]
			} else {
				z.Fields = make([]Field, zcnk)
			}
			for zrhi := range z.Fields {
				err = z.Fields[zrhi].DecodeMsg(dc)
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
	if nextMiss6zbry != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zbry = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zbry = []bool{false, false}

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
	for zrhi := range z.Fields {
		err = z.Fields[zrhi].EncodeMsg(en)
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
	for zrhi := range z.Fields {
		o, err = z.Fields[zrhi].MarshalMsg(o)
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
	const maxFields7zokc = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zokc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zokc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zokc := totalEncodedFields7zokc
	missingFieldsLeft7zokc := maxFields7zokc - totalEncodedFields7zokc

	var nextMiss7zokc int32 = -1
	var found7zokc [maxFields7zokc]bool
	var curField7zokc string

doneWithStruct7zokc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zokc > 0 || missingFieldsLeft7zokc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zokc, missingFieldsLeft7zokc, msgp.ShowFound(found7zokc[:]), unmarshalMsgFieldOrder7zokc)
		if encodedFieldsLeft7zokc > 0 {
			encodedFieldsLeft7zokc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zokc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zokc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zokc = 0
			}
			for nextMiss7zokc < maxFields7zokc && (found7zokc[nextMiss7zokc] || unmarshalMsgFieldSkip7zokc[nextMiss7zokc]) {
				nextMiss7zokc++
			}
			if nextMiss7zokc == maxFields7zokc {
				// filled all the empty fields!
				break doneWithStruct7zokc
			}
			missingFieldsLeft7zokc--
			curField7zokc = unmarshalMsgFieldOrder7zokc[nextMiss7zokc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zokc)
		switch curField7zokc {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zokc[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zokc[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zclb uint32
				zclb, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zclb) {
					z.Fields = (z.Fields)[:zclb]
				} else {
					z.Fields = make([]Field, zclb)
				}
				for zrhi := range z.Fields {
					bts, err = z.Fields[zrhi].UnmarshalMsg(bts)
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
	if nextMiss7zokc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zokc = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zokc = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zrhi := range z.Fields {
		s += z.Fields[zrhi].Msgsize()
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
		var zgdr uint64
		zgdr, err = dc.ReadUint64()
		(*z) = Zkind(zgdr)
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
		var zxml uint64
		zxml, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zxml)
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
	const maxFields8zhrh = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zhrh uint32
	totalEncodedFields8zhrh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zhrh := totalEncodedFields8zhrh
	missingFieldsLeft8zhrh := maxFields8zhrh - totalEncodedFields8zhrh

	var nextMiss8zhrh int32 = -1
	var found8zhrh [maxFields8zhrh]bool
	var curField8zhrh string

doneWithStruct8zhrh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zhrh > 0 || missingFieldsLeft8zhrh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zhrh, missingFieldsLeft8zhrh, msgp.ShowFound(found8zhrh[:]), decodeMsgFieldOrder8zhrh)
		if encodedFieldsLeft8zhrh > 0 {
			encodedFieldsLeft8zhrh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zhrh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zhrh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zhrh = 0
			}
			for nextMiss8zhrh < maxFields8zhrh && (found8zhrh[nextMiss8zhrh] || decodeMsgFieldSkip8zhrh[nextMiss8zhrh]) {
				nextMiss8zhrh++
			}
			if nextMiss8zhrh == maxFields8zhrh {
				// filled all the empty fields!
				break doneWithStruct8zhrh
			}
			missingFieldsLeft8zhrh--
			curField8zhrh = decodeMsgFieldOrder8zhrh[nextMiss8zhrh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zhrh)
		switch curField8zhrh {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zhrh[0] = true
			{
				var zfiu uint64
				zfiu, err = dc.ReadUint64()
				z.Kind = Zkind(zfiu)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zhrh[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zhrh[2] = true
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
			found8zhrh[3] = true
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
	if nextMiss8zhrh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zhrh = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zhrh = []bool{false, false, false, false}

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
	var empty_zqrf [4]bool
	fieldsInUse_zlsh := z.fieldsNotEmpty(empty_zqrf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlsh)
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
	if !empty_zqrf[1] {
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

	if !empty_zqrf[2] {
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

	if !empty_zqrf[3] {
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
	const maxFields9zrmq = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zrmq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zrmq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zrmq := totalEncodedFields9zrmq
	missingFieldsLeft9zrmq := maxFields9zrmq - totalEncodedFields9zrmq

	var nextMiss9zrmq int32 = -1
	var found9zrmq [maxFields9zrmq]bool
	var curField9zrmq string

doneWithStruct9zrmq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zrmq > 0 || missingFieldsLeft9zrmq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zrmq, missingFieldsLeft9zrmq, msgp.ShowFound(found9zrmq[:]), unmarshalMsgFieldOrder9zrmq)
		if encodedFieldsLeft9zrmq > 0 {
			encodedFieldsLeft9zrmq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zrmq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zrmq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zrmq = 0
			}
			for nextMiss9zrmq < maxFields9zrmq && (found9zrmq[nextMiss9zrmq] || unmarshalMsgFieldSkip9zrmq[nextMiss9zrmq]) {
				nextMiss9zrmq++
			}
			if nextMiss9zrmq == maxFields9zrmq {
				// filled all the empty fields!
				break doneWithStruct9zrmq
			}
			missingFieldsLeft9zrmq--
			curField9zrmq = unmarshalMsgFieldOrder9zrmq[nextMiss9zrmq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zrmq)
		switch curField9zrmq {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zrmq[0] = true
			{
				var zkbq uint64
				zkbq, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zkbq)
			}
		case "Str":
			found9zrmq[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zrmq[2] = true
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
			found9zrmq[3] = true
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
	if nextMiss9zrmq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zrmq = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zrmq = []bool{false, false, false, false}

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
