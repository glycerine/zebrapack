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
	const maxFields0zyvy = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zyvy uint32
	totalEncodedFields0zyvy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zyvy := totalEncodedFields0zyvy
	missingFieldsLeft0zyvy := maxFields0zyvy - totalEncodedFields0zyvy

	var nextMiss0zyvy int32 = -1
	var found0zyvy [maxFields0zyvy]bool
	var curField0zyvy string

doneWithStruct0zyvy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zyvy > 0 || missingFieldsLeft0zyvy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zyvy, missingFieldsLeft0zyvy, msgp.ShowFound(found0zyvy[:]), decodeMsgFieldOrder0zyvy)
		if encodedFieldsLeft0zyvy > 0 {
			encodedFieldsLeft0zyvy--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zyvy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zyvy < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zyvy = 0
			}
			for nextMiss0zyvy < maxFields0zyvy && (found0zyvy[nextMiss0zyvy] || decodeMsgFieldSkip0zyvy[nextMiss0zyvy]) {
				nextMiss0zyvy++
			}
			if nextMiss0zyvy == maxFields0zyvy {
				// filled all the empty fields!
				break doneWithStruct0zyvy
			}
			missingFieldsLeft0zyvy--
			curField0zyvy = decodeMsgFieldOrder0zyvy[nextMiss0zyvy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zyvy)
		switch curField0zyvy {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zyvy[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zyvy[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zyvy[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zyvy[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zyvy[4] = true
			{
				var zbqy uint64
				zbqy, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zbqy)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zyvy[5] = true
			{
				var zgco uint64
				zgco, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zgco)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zyvy[6] = true
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
			found0zyvy[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zyvy[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zyvy[9] = true
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
	if nextMiss0zyvy != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zyvy = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zyvy = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zwsq [10]bool
	fieldsInUse_zycy := z.fieldsNotEmpty(empty_zwsq[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zycy)
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
	if !empty_zwsq[3] {
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

	if !empty_zwsq[4] {
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

	if !empty_zwsq[5] {
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

	if !empty_zwsq[6] {
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

	if !empty_zwsq[7] {
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

	if !empty_zwsq[8] {
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

	if !empty_zwsq[9] {
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
	const maxFields1zhjd = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zhjd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhjd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zhjd := totalEncodedFields1zhjd
	missingFieldsLeft1zhjd := maxFields1zhjd - totalEncodedFields1zhjd

	var nextMiss1zhjd int32 = -1
	var found1zhjd [maxFields1zhjd]bool
	var curField1zhjd string

doneWithStruct1zhjd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhjd > 0 || missingFieldsLeft1zhjd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhjd, missingFieldsLeft1zhjd, msgp.ShowFound(found1zhjd[:]), unmarshalMsgFieldOrder1zhjd)
		if encodedFieldsLeft1zhjd > 0 {
			encodedFieldsLeft1zhjd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zhjd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zhjd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhjd = 0
			}
			for nextMiss1zhjd < maxFields1zhjd && (found1zhjd[nextMiss1zhjd] || unmarshalMsgFieldSkip1zhjd[nextMiss1zhjd]) {
				nextMiss1zhjd++
			}
			if nextMiss1zhjd == maxFields1zhjd {
				// filled all the empty fields!
				break doneWithStruct1zhjd
			}
			missingFieldsLeft1zhjd--
			curField1zhjd = unmarshalMsgFieldOrder1zhjd[nextMiss1zhjd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhjd)
		switch curField1zhjd {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zhjd[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zhjd[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zhjd[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zhjd[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zhjd[4] = true
			{
				var zlvl uint64
				zlvl, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zlvl)
			}
		case "FieldPrimitive":
			found1zhjd[5] = true
			{
				var zbtd uint64
				zbtd, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zbtd)
			}
		case "FieldFullType":
			found1zhjd[6] = true
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
			found1zhjd[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zhjd[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zhjd[9] = true
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
	if nextMiss1zhjd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zhjd = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zhjd = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zdrw = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zdrw uint32
	totalEncodedFields2zdrw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zdrw := totalEncodedFields2zdrw
	missingFieldsLeft2zdrw := maxFields2zdrw - totalEncodedFields2zdrw

	var nextMiss2zdrw int32 = -1
	var found2zdrw [maxFields2zdrw]bool
	var curField2zdrw string

doneWithStruct2zdrw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zdrw > 0 || missingFieldsLeft2zdrw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zdrw, missingFieldsLeft2zdrw, msgp.ShowFound(found2zdrw[:]), decodeMsgFieldOrder2zdrw)
		if encodedFieldsLeft2zdrw > 0 {
			encodedFieldsLeft2zdrw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zdrw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zdrw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zdrw = 0
			}
			for nextMiss2zdrw < maxFields2zdrw && (found2zdrw[nextMiss2zdrw] || decodeMsgFieldSkip2zdrw[nextMiss2zdrw]) {
				nextMiss2zdrw++
			}
			if nextMiss2zdrw == maxFields2zdrw {
				// filled all the empty fields!
				break doneWithStruct2zdrw
			}
			missingFieldsLeft2zdrw--
			curField2zdrw = decodeMsgFieldOrder2zdrw[nextMiss2zdrw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zdrw)
		switch curField2zdrw {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zdrw[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zdrw[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zdrw[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zdrw[3] = true
			var zftk uint32
			zftk, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zftk) {
				z.Structs = (z.Structs)[:zftk]
			} else {
				z.Structs = make([]Struct, zftk)
			}
			for ztkx := range z.Structs {
				const maxFields3zeho = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zeho uint32
				totalEncodedFields3zeho, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zeho := totalEncodedFields3zeho
				missingFieldsLeft3zeho := maxFields3zeho - totalEncodedFields3zeho

				var nextMiss3zeho int32 = -1
				var found3zeho [maxFields3zeho]bool
				var curField3zeho string

			doneWithStruct3zeho:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zeho > 0 || missingFieldsLeft3zeho > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zeho, missingFieldsLeft3zeho, msgp.ShowFound(found3zeho[:]), decodeMsgFieldOrder3zeho)
					if encodedFieldsLeft3zeho > 0 {
						encodedFieldsLeft3zeho--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zeho = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zeho < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zeho = 0
						}
						for nextMiss3zeho < maxFields3zeho && (found3zeho[nextMiss3zeho] || decodeMsgFieldSkip3zeho[nextMiss3zeho]) {
							nextMiss3zeho++
						}
						if nextMiss3zeho == maxFields3zeho {
							// filled all the empty fields!
							break doneWithStruct3zeho
						}
						missingFieldsLeft3zeho--
						curField3zeho = decodeMsgFieldOrder3zeho[nextMiss3zeho]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zeho)
					switch curField3zeho {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zeho[0] = true
						z.Structs[ztkx].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zeho[1] = true
						var zvhb uint32
						zvhb, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[ztkx].Fields) >= int(zvhb) {
							z.Structs[ztkx].Fields = (z.Structs[ztkx].Fields)[:zvhb]
						} else {
							z.Structs[ztkx].Fields = make([]Field, zvhb)
						}
						for zkzw := range z.Structs[ztkx].Fields {
							err = z.Structs[ztkx].Fields[zkzw].DecodeMsg(dc)
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
				if nextMiss3zeho != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zdrw[4] = true
			var zdlc uint32
			zdlc, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zdlc) {
				z.Imports = (z.Imports)[:zdlc]
			} else {
				z.Imports = make([]string, zdlc)
			}
			for zkwv := range z.Imports {
				z.Imports[zkwv], err = dc.ReadString()
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
	if nextMiss2zdrw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zdrw = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zdrw = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zeho = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zeho = []bool{false, false}

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
	for ztkx := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[ztkx].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[ztkx].Fields)))
		if err != nil {
			panic(err)
		}
		for zkzw := range z.Structs[ztkx].Fields {
			err = z.Structs[ztkx].Fields[zkzw].EncodeMsg(en)
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
	for zkwv := range z.Imports {
		err = en.WriteString(z.Imports[zkwv])
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
	for ztkx := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[ztkx].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[ztkx].Fields)))
		for zkzw := range z.Structs[ztkx].Fields {
			o, err = z.Structs[ztkx].Fields[zkzw].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zkwv := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zkwv])
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
	const maxFields4zllc = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zllc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zllc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zllc := totalEncodedFields4zllc
	missingFieldsLeft4zllc := maxFields4zllc - totalEncodedFields4zllc

	var nextMiss4zllc int32 = -1
	var found4zllc [maxFields4zllc]bool
	var curField4zllc string

doneWithStruct4zllc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zllc > 0 || missingFieldsLeft4zllc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zllc, missingFieldsLeft4zllc, msgp.ShowFound(found4zllc[:]), unmarshalMsgFieldOrder4zllc)
		if encodedFieldsLeft4zllc > 0 {
			encodedFieldsLeft4zllc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zllc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zllc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zllc = 0
			}
			for nextMiss4zllc < maxFields4zllc && (found4zllc[nextMiss4zllc] || unmarshalMsgFieldSkip4zllc[nextMiss4zllc]) {
				nextMiss4zllc++
			}
			if nextMiss4zllc == maxFields4zllc {
				// filled all the empty fields!
				break doneWithStruct4zllc
			}
			missingFieldsLeft4zllc--
			curField4zllc = unmarshalMsgFieldOrder4zllc[nextMiss4zllc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zllc)
		switch curField4zllc {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zllc[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zllc[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zllc[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zllc[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zmrl uint32
				zmrl, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zmrl) {
					z.Structs = (z.Structs)[:zmrl]
				} else {
					z.Structs = make([]Struct, zmrl)
				}
				for ztkx := range z.Structs {
					const maxFields5zfyk = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zfyk uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zfyk, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zfyk := totalEncodedFields5zfyk
					missingFieldsLeft5zfyk := maxFields5zfyk - totalEncodedFields5zfyk

					var nextMiss5zfyk int32 = -1
					var found5zfyk [maxFields5zfyk]bool
					var curField5zfyk string

				doneWithStruct5zfyk:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zfyk > 0 || missingFieldsLeft5zfyk > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zfyk, missingFieldsLeft5zfyk, msgp.ShowFound(found5zfyk[:]), unmarshalMsgFieldOrder5zfyk)
						if encodedFieldsLeft5zfyk > 0 {
							encodedFieldsLeft5zfyk--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zfyk = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zfyk < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zfyk = 0
							}
							for nextMiss5zfyk < maxFields5zfyk && (found5zfyk[nextMiss5zfyk] || unmarshalMsgFieldSkip5zfyk[nextMiss5zfyk]) {
								nextMiss5zfyk++
							}
							if nextMiss5zfyk == maxFields5zfyk {
								// filled all the empty fields!
								break doneWithStruct5zfyk
							}
							missingFieldsLeft5zfyk--
							curField5zfyk = unmarshalMsgFieldOrder5zfyk[nextMiss5zfyk]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zfyk)
						switch curField5zfyk {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zfyk[0] = true
							z.Structs[ztkx].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zfyk[1] = true
							if nbs.AlwaysNil {
								(z.Structs[ztkx].Fields) = (z.Structs[ztkx].Fields)[:0]
							} else {

								var zehk uint32
								zehk, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[ztkx].Fields) >= int(zehk) {
									z.Structs[ztkx].Fields = (z.Structs[ztkx].Fields)[:zehk]
								} else {
									z.Structs[ztkx].Fields = make([]Field, zehk)
								}
								for zkzw := range z.Structs[ztkx].Fields {
									bts, err = z.Structs[ztkx].Fields[zkzw].UnmarshalMsg(bts)
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
					if nextMiss5zfyk != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4zllc[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zgzd uint32
				zgzd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zgzd) {
					z.Imports = (z.Imports)[:zgzd]
				} else {
					z.Imports = make([]string, zgzd)
				}
				for zkwv := range z.Imports {
					z.Imports[zkwv], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zllc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zllc = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zllc = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zfyk = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zfyk = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for ztkx := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[ztkx].StructName) + 7 + msgp.ArrayHeaderSize
		for zkzw := range z.Structs[ztkx].Fields {
			s += z.Structs[ztkx].Fields[zkzw].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zkwv := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zkwv])
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
	const maxFields6zgew = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zgew uint32
	totalEncodedFields6zgew, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zgew := totalEncodedFields6zgew
	missingFieldsLeft6zgew := maxFields6zgew - totalEncodedFields6zgew

	var nextMiss6zgew int32 = -1
	var found6zgew [maxFields6zgew]bool
	var curField6zgew string

doneWithStruct6zgew:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zgew > 0 || missingFieldsLeft6zgew > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zgew, missingFieldsLeft6zgew, msgp.ShowFound(found6zgew[:]), decodeMsgFieldOrder6zgew)
		if encodedFieldsLeft6zgew > 0 {
			encodedFieldsLeft6zgew--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zgew = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zgew < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zgew = 0
			}
			for nextMiss6zgew < maxFields6zgew && (found6zgew[nextMiss6zgew] || decodeMsgFieldSkip6zgew[nextMiss6zgew]) {
				nextMiss6zgew++
			}
			if nextMiss6zgew == maxFields6zgew {
				// filled all the empty fields!
				break doneWithStruct6zgew
			}
			missingFieldsLeft6zgew--
			curField6zgew = decodeMsgFieldOrder6zgew[nextMiss6zgew]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zgew)
		switch curField6zgew {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zgew[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zgew[1] = true
			var zgum uint32
			zgum, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zgum) {
				z.Fields = (z.Fields)[:zgum]
			} else {
				z.Fields = make([]Field, zgum)
			}
			for zwrg := range z.Fields {
				err = z.Fields[zwrg].DecodeMsg(dc)
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
	if nextMiss6zgew != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zgew = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zgew = []bool{false, false}

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
	for zwrg := range z.Fields {
		err = z.Fields[zwrg].EncodeMsg(en)
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
	for zwrg := range z.Fields {
		o, err = z.Fields[zwrg].MarshalMsg(o)
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
	const maxFields7zdls = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zdls uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zdls, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zdls := totalEncodedFields7zdls
	missingFieldsLeft7zdls := maxFields7zdls - totalEncodedFields7zdls

	var nextMiss7zdls int32 = -1
	var found7zdls [maxFields7zdls]bool
	var curField7zdls string

doneWithStruct7zdls:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zdls > 0 || missingFieldsLeft7zdls > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zdls, missingFieldsLeft7zdls, msgp.ShowFound(found7zdls[:]), unmarshalMsgFieldOrder7zdls)
		if encodedFieldsLeft7zdls > 0 {
			encodedFieldsLeft7zdls--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zdls = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zdls < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zdls = 0
			}
			for nextMiss7zdls < maxFields7zdls && (found7zdls[nextMiss7zdls] || unmarshalMsgFieldSkip7zdls[nextMiss7zdls]) {
				nextMiss7zdls++
			}
			if nextMiss7zdls == maxFields7zdls {
				// filled all the empty fields!
				break doneWithStruct7zdls
			}
			missingFieldsLeft7zdls--
			curField7zdls = unmarshalMsgFieldOrder7zdls[nextMiss7zdls]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zdls)
		switch curField7zdls {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zdls[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zdls[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zkua uint32
				zkua, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zkua) {
					z.Fields = (z.Fields)[:zkua]
				} else {
					z.Fields = make([]Field, zkua)
				}
				for zwrg := range z.Fields {
					bts, err = z.Fields[zwrg].UnmarshalMsg(bts)
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
	if nextMiss7zdls != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zdls = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zdls = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zwrg := range z.Fields {
		s += z.Fields[zwrg].Msgsize()
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
		var zzmc uint64
		zzmc, err = dc.ReadUint64()
		(*z) = Zkind(zzmc)
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
		var zonz uint64
		zonz, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zonz)
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
	const maxFields8zgnf = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zgnf uint32
	totalEncodedFields8zgnf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zgnf := totalEncodedFields8zgnf
	missingFieldsLeft8zgnf := maxFields8zgnf - totalEncodedFields8zgnf

	var nextMiss8zgnf int32 = -1
	var found8zgnf [maxFields8zgnf]bool
	var curField8zgnf string

doneWithStruct8zgnf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zgnf > 0 || missingFieldsLeft8zgnf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zgnf, missingFieldsLeft8zgnf, msgp.ShowFound(found8zgnf[:]), decodeMsgFieldOrder8zgnf)
		if encodedFieldsLeft8zgnf > 0 {
			encodedFieldsLeft8zgnf--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zgnf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zgnf < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zgnf = 0
			}
			for nextMiss8zgnf < maxFields8zgnf && (found8zgnf[nextMiss8zgnf] || decodeMsgFieldSkip8zgnf[nextMiss8zgnf]) {
				nextMiss8zgnf++
			}
			if nextMiss8zgnf == maxFields8zgnf {
				// filled all the empty fields!
				break doneWithStruct8zgnf
			}
			missingFieldsLeft8zgnf--
			curField8zgnf = decodeMsgFieldOrder8zgnf[nextMiss8zgnf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zgnf)
		switch curField8zgnf {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zgnf[0] = true
			{
				var zlav uint64
				zlav, err = dc.ReadUint64()
				z.Kind = Zkind(zlav)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zgnf[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zgnf[2] = true
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
			found8zgnf[3] = true
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
	if nextMiss8zgnf != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zgnf = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zgnf = []bool{false, false, false, false}

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
	var empty_zxce [4]bool
	fieldsInUse_zpce := z.fieldsNotEmpty(empty_zxce[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpce)
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
	if !empty_zxce[1] {
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

	if !empty_zxce[2] {
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

	if !empty_zxce[3] {
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
	const maxFields9zwzn = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zwzn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zwzn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zwzn := totalEncodedFields9zwzn
	missingFieldsLeft9zwzn := maxFields9zwzn - totalEncodedFields9zwzn

	var nextMiss9zwzn int32 = -1
	var found9zwzn [maxFields9zwzn]bool
	var curField9zwzn string

doneWithStruct9zwzn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zwzn > 0 || missingFieldsLeft9zwzn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zwzn, missingFieldsLeft9zwzn, msgp.ShowFound(found9zwzn[:]), unmarshalMsgFieldOrder9zwzn)
		if encodedFieldsLeft9zwzn > 0 {
			encodedFieldsLeft9zwzn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zwzn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zwzn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zwzn = 0
			}
			for nextMiss9zwzn < maxFields9zwzn && (found9zwzn[nextMiss9zwzn] || unmarshalMsgFieldSkip9zwzn[nextMiss9zwzn]) {
				nextMiss9zwzn++
			}
			if nextMiss9zwzn == maxFields9zwzn {
				// filled all the empty fields!
				break doneWithStruct9zwzn
			}
			missingFieldsLeft9zwzn--
			curField9zwzn = unmarshalMsgFieldOrder9zwzn[nextMiss9zwzn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zwzn)
		switch curField9zwzn {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zwzn[0] = true
			{
				var zwos uint64
				zwos, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zwos)
			}
		case "Str":
			found9zwzn[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zwzn[2] = true
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
			found9zwzn[3] = true
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
	if nextMiss9zwzn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zwzn = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zwzn = []bool{false, false, false, false}

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
