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
	const maxFields0zerc = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zerc uint32
	totalEncodedFields0zerc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zerc := totalEncodedFields0zerc
	missingFieldsLeft0zerc := maxFields0zerc - totalEncodedFields0zerc

	var nextMiss0zerc int32 = -1
	var found0zerc [maxFields0zerc]bool
	var curField0zerc string

doneWithStruct0zerc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zerc > 0 || missingFieldsLeft0zerc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zerc, missingFieldsLeft0zerc, msgp.ShowFound(found0zerc[:]), decodeMsgFieldOrder0zerc)
		if encodedFieldsLeft0zerc > 0 {
			encodedFieldsLeft0zerc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zerc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zerc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zerc = 0
			}
			for nextMiss0zerc < maxFields0zerc && (found0zerc[nextMiss0zerc] || decodeMsgFieldSkip0zerc[nextMiss0zerc]) {
				nextMiss0zerc++
			}
			if nextMiss0zerc == maxFields0zerc {
				// filled all the empty fields!
				break doneWithStruct0zerc
			}
			missingFieldsLeft0zerc--
			curField0zerc = decodeMsgFieldOrder0zerc[nextMiss0zerc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zerc)
		switch curField0zerc {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zerc[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zerc[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zerc[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zerc[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zerc[4] = true
			{
				var zywl uint64
				zywl, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zywl)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zerc[5] = true
			{
				var zanb uint64
				zanb, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zanb)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zerc[6] = true
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
			found0zerc[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zerc[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zerc[9] = true
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
	if nextMiss0zerc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zerc = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zerc = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zsbf [10]bool
	fieldsInUse_zxml := z.fieldsNotEmpty(empty_zsbf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxml)
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
	if !empty_zsbf[3] {
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

	if !empty_zsbf[4] {
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

	if !empty_zsbf[5] {
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

	if !empty_zsbf[6] {
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

	if !empty_zsbf[7] {
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

	if !empty_zsbf[8] {
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

	if !empty_zsbf[9] {
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
	const maxFields1zubv = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zubv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zubv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zubv := totalEncodedFields1zubv
	missingFieldsLeft1zubv := maxFields1zubv - totalEncodedFields1zubv

	var nextMiss1zubv int32 = -1
	var found1zubv [maxFields1zubv]bool
	var curField1zubv string

doneWithStruct1zubv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zubv > 0 || missingFieldsLeft1zubv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zubv, missingFieldsLeft1zubv, msgp.ShowFound(found1zubv[:]), unmarshalMsgFieldOrder1zubv)
		if encodedFieldsLeft1zubv > 0 {
			encodedFieldsLeft1zubv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zubv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zubv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zubv = 0
			}
			for nextMiss1zubv < maxFields1zubv && (found1zubv[nextMiss1zubv] || unmarshalMsgFieldSkip1zubv[nextMiss1zubv]) {
				nextMiss1zubv++
			}
			if nextMiss1zubv == maxFields1zubv {
				// filled all the empty fields!
				break doneWithStruct1zubv
			}
			missingFieldsLeft1zubv--
			curField1zubv = unmarshalMsgFieldOrder1zubv[nextMiss1zubv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zubv)
		switch curField1zubv {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zubv[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zubv[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zubv[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zubv[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zubv[4] = true
			{
				var zgyo uint64
				zgyo, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zgyo)
			}
		case "FieldPrimitive":
			found1zubv[5] = true
			{
				var zdrw uint64
				zdrw, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zdrw)
			}
		case "FieldFullType":
			found1zubv[6] = true
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
			found1zubv[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zubv[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zubv[9] = true
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
	if nextMiss1zubv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zubv = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zubv = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zccg = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zccg uint32
	totalEncodedFields2zccg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zccg := totalEncodedFields2zccg
	missingFieldsLeft2zccg := maxFields2zccg - totalEncodedFields2zccg

	var nextMiss2zccg int32 = -1
	var found2zccg [maxFields2zccg]bool
	var curField2zccg string

doneWithStruct2zccg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zccg > 0 || missingFieldsLeft2zccg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zccg, missingFieldsLeft2zccg, msgp.ShowFound(found2zccg[:]), decodeMsgFieldOrder2zccg)
		if encodedFieldsLeft2zccg > 0 {
			encodedFieldsLeft2zccg--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zccg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zccg < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zccg = 0
			}
			for nextMiss2zccg < maxFields2zccg && (found2zccg[nextMiss2zccg] || decodeMsgFieldSkip2zccg[nextMiss2zccg]) {
				nextMiss2zccg++
			}
			if nextMiss2zccg == maxFields2zccg {
				// filled all the empty fields!
				break doneWithStruct2zccg
			}
			missingFieldsLeft2zccg--
			curField2zccg = decodeMsgFieldOrder2zccg[nextMiss2zccg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zccg)
		switch curField2zccg {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zccg[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zccg[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zccg[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zccg[3] = true
			var znxc uint32
			znxc, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && znxc > 0 {
				z.Structs = make(map[string]*Struct, znxc)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for znxc > 0 {
				znxc--
				var zfun string
				var zvkz *Struct
				zfun, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					zvkz = nil
				} else {
					if zvkz == nil {
						zvkz = new(Struct)
					}
					const maxFields3zadu = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zadu uint32
					totalEncodedFields3zadu, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zadu := totalEncodedFields3zadu
					missingFieldsLeft3zadu := maxFields3zadu - totalEncodedFields3zadu

					var nextMiss3zadu int32 = -1
					var found3zadu [maxFields3zadu]bool
					var curField3zadu string

				doneWithStruct3zadu:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zadu > 0 || missingFieldsLeft3zadu > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zadu, missingFieldsLeft3zadu, msgp.ShowFound(found3zadu[:]), decodeMsgFieldOrder3zadu)
						if encodedFieldsLeft3zadu > 0 {
							encodedFieldsLeft3zadu--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zadu = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zadu < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zadu = 0
							}
							for nextMiss3zadu < maxFields3zadu && (found3zadu[nextMiss3zadu] || decodeMsgFieldSkip3zadu[nextMiss3zadu]) {
								nextMiss3zadu++
							}
							if nextMiss3zadu == maxFields3zadu {
								// filled all the empty fields!
								break doneWithStruct3zadu
							}
							missingFieldsLeft3zadu--
							curField3zadu = decodeMsgFieldOrder3zadu[nextMiss3zadu]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zadu)
						switch curField3zadu {
						// -- templateDecodeMsg ends here --

						case "StructName":
							found3zadu[0] = true
							zvkz.StructName, err = dc.ReadString()
							if err != nil {
								panic(err)
							}
						case "Fields":
							found3zadu[1] = true
							var zmrm uint32
							zmrm, err = dc.ReadArrayHeader()
							if err != nil {
								panic(err)
							}
							if cap(zvkz.Fields) >= int(zmrm) {
								zvkz.Fields = (zvkz.Fields)[:zmrm]
							} else {
								zvkz.Fields = make([]Field, zmrm)
							}
							for zsis := range zvkz.Fields {
								err = zvkz.Fields[zsis].DecodeMsg(dc)
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
					if nextMiss3zadu != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[zfun] = zvkz
			}
		case "Imports":
			found2zccg[4] = true
			var zcil uint32
			zcil, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zcil) {
				z.Imports = (z.Imports)[:zcil]
			} else {
				z.Imports = make([]string, zcil)
			}
			for zpnh := range z.Imports {
				z.Imports[zpnh], err = dc.ReadString()
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
	if nextMiss2zccg != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zccg = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zccg = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zadu = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zadu = []bool{false, false}

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
	for zfun, zvkz := range z.Structs {
		err = en.WriteString(zfun)
		if err != nil {
			panic(err)
		}
		if zvkz == nil {
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
			err = en.WriteString(zvkz.StructName)
			if err != nil {
				panic(err)
			}
			// write "Fields"
			err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(zvkz.Fields)))
			if err != nil {
				panic(err)
			}
			for zsis := range zvkz.Fields {
				err = zvkz.Fields[zsis].EncodeMsg(en)
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
	for zpnh := range z.Imports {
		err = en.WriteString(z.Imports[zpnh])
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
	for zfun, zvkz := range z.Structs {
		o = msgp.AppendString(o, zfun)
		if zvkz == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName"
			o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, zvkz.StructName)
			// string "Fields"
			o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			o = msgp.AppendArrayHeader(o, uint32(len(zvkz.Fields)))
			for zsis := range zvkz.Fields {
				o, err = zvkz.Fields[zsis].MarshalMsg(o)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zpnh := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zpnh])
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
	const maxFields4ztnk = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4ztnk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4ztnk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4ztnk := totalEncodedFields4ztnk
	missingFieldsLeft4ztnk := maxFields4ztnk - totalEncodedFields4ztnk

	var nextMiss4ztnk int32 = -1
	var found4ztnk [maxFields4ztnk]bool
	var curField4ztnk string

doneWithStruct4ztnk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4ztnk > 0 || missingFieldsLeft4ztnk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4ztnk, missingFieldsLeft4ztnk, msgp.ShowFound(found4ztnk[:]), unmarshalMsgFieldOrder4ztnk)
		if encodedFieldsLeft4ztnk > 0 {
			encodedFieldsLeft4ztnk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4ztnk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4ztnk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4ztnk = 0
			}
			for nextMiss4ztnk < maxFields4ztnk && (found4ztnk[nextMiss4ztnk] || unmarshalMsgFieldSkip4ztnk[nextMiss4ztnk]) {
				nextMiss4ztnk++
			}
			if nextMiss4ztnk == maxFields4ztnk {
				// filled all the empty fields!
				break doneWithStruct4ztnk
			}
			missingFieldsLeft4ztnk--
			curField4ztnk = unmarshalMsgFieldOrder4ztnk[nextMiss4ztnk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4ztnk)
		switch curField4ztnk {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4ztnk[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4ztnk[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4ztnk[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4ztnk[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zuig uint32
				zuig, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zuig > 0 {
					z.Structs = make(map[string]*Struct, zuig)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zuig > 0 {
					var zfun string
					var zvkz *Struct
					zuig--
					zfun, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && zvkz == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if zvkz == nil {
							zvkz = new(Struct)
						}
						const maxFields5zxcs = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zxcs uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zxcs, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								panic(err)
								return
							}
						}
						encodedFieldsLeft5zxcs := totalEncodedFields5zxcs
						missingFieldsLeft5zxcs := maxFields5zxcs - totalEncodedFields5zxcs

						var nextMiss5zxcs int32 = -1
						var found5zxcs [maxFields5zxcs]bool
						var curField5zxcs string

					doneWithStruct5zxcs:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zxcs > 0 || missingFieldsLeft5zxcs > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxcs, missingFieldsLeft5zxcs, msgp.ShowFound(found5zxcs[:]), unmarshalMsgFieldOrder5zxcs)
							if encodedFieldsLeft5zxcs > 0 {
								encodedFieldsLeft5zxcs--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									panic(err)
									return
								}
								curField5zxcs = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zxcs < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zxcs = 0
								}
								for nextMiss5zxcs < maxFields5zxcs && (found5zxcs[nextMiss5zxcs] || unmarshalMsgFieldSkip5zxcs[nextMiss5zxcs]) {
									nextMiss5zxcs++
								}
								if nextMiss5zxcs == maxFields5zxcs {
									// filled all the empty fields!
									break doneWithStruct5zxcs
								}
								missingFieldsLeft5zxcs--
								curField5zxcs = unmarshalMsgFieldOrder5zxcs[nextMiss5zxcs]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zxcs)
							switch curField5zxcs {
							// -- templateUnmarshalMsg ends here --

							case "StructName":
								found5zxcs[0] = true
								zvkz.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									panic(err)
								}
							case "Fields":
								found5zxcs[1] = true
								if nbs.AlwaysNil {
									(zvkz.Fields) = (zvkz.Fields)[:0]
								} else {

									var zwba uint32
									zwba, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										panic(err)
									}
									if cap(zvkz.Fields) >= int(zwba) {
										zvkz.Fields = (zvkz.Fields)[:zwba]
									} else {
										zvkz.Fields = make([]Field, zwba)
									}
									for zsis := range zvkz.Fields {
										bts, err = zvkz.Fields[zsis].UnmarshalMsg(bts)
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
						if nextMiss5zxcs != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[zfun] = zvkz
				}
			}
		case "Imports":
			found4ztnk[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zgvq uint32
				zgvq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zgvq) {
					z.Imports = (z.Imports)[:zgvq]
				} else {
					z.Imports = make([]string, zgvq)
				}
				for zpnh := range z.Imports {
					z.Imports[zpnh], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4ztnk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4ztnk = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4ztnk = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zxcs = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zxcs = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zfun, zvkz := range z.Structs {
			_ = zvkz
			_ = zfun
			s += msgp.StringPrefixSize + len(zfun)
			if zvkz == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 11 + msgp.StringPrefixSize + len(zvkz.StructName) + 7 + msgp.ArrayHeaderSize
				for zsis := range zvkz.Fields {
					s += zvkz.Fields[zsis].Msgsize()
				}
			}
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zpnh := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zpnh])
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
	const maxFields6zzbr = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zzbr uint32
	totalEncodedFields6zzbr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zzbr := totalEncodedFields6zzbr
	missingFieldsLeft6zzbr := maxFields6zzbr - totalEncodedFields6zzbr

	var nextMiss6zzbr int32 = -1
	var found6zzbr [maxFields6zzbr]bool
	var curField6zzbr string

doneWithStruct6zzbr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zzbr > 0 || missingFieldsLeft6zzbr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zzbr, missingFieldsLeft6zzbr, msgp.ShowFound(found6zzbr[:]), decodeMsgFieldOrder6zzbr)
		if encodedFieldsLeft6zzbr > 0 {
			encodedFieldsLeft6zzbr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zzbr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zzbr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zzbr = 0
			}
			for nextMiss6zzbr < maxFields6zzbr && (found6zzbr[nextMiss6zzbr] || decodeMsgFieldSkip6zzbr[nextMiss6zzbr]) {
				nextMiss6zzbr++
			}
			if nextMiss6zzbr == maxFields6zzbr {
				// filled all the empty fields!
				break doneWithStruct6zzbr
			}
			missingFieldsLeft6zzbr--
			curField6zzbr = decodeMsgFieldOrder6zzbr[nextMiss6zzbr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zzbr)
		switch curField6zzbr {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zzbr[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zzbr[1] = true
			var zgix uint32
			zgix, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zgix) {
				z.Fields = (z.Fields)[:zgix]
			} else {
				z.Fields = make([]Field, zgix)
			}
			for zkns := range z.Fields {
				err = z.Fields[zkns].DecodeMsg(dc)
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
	if nextMiss6zzbr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zzbr = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zzbr = []bool{false, false}

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
	for zkns := range z.Fields {
		err = z.Fields[zkns].EncodeMsg(en)
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
	for zkns := range z.Fields {
		o, err = z.Fields[zkns].MarshalMsg(o)
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
	const maxFields7zhge = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zhge uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zhge, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zhge := totalEncodedFields7zhge
	missingFieldsLeft7zhge := maxFields7zhge - totalEncodedFields7zhge

	var nextMiss7zhge int32 = -1
	var found7zhge [maxFields7zhge]bool
	var curField7zhge string

doneWithStruct7zhge:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zhge > 0 || missingFieldsLeft7zhge > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zhge, missingFieldsLeft7zhge, msgp.ShowFound(found7zhge[:]), unmarshalMsgFieldOrder7zhge)
		if encodedFieldsLeft7zhge > 0 {
			encodedFieldsLeft7zhge--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zhge = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zhge < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zhge = 0
			}
			for nextMiss7zhge < maxFields7zhge && (found7zhge[nextMiss7zhge] || unmarshalMsgFieldSkip7zhge[nextMiss7zhge]) {
				nextMiss7zhge++
			}
			if nextMiss7zhge == maxFields7zhge {
				// filled all the empty fields!
				break doneWithStruct7zhge
			}
			missingFieldsLeft7zhge--
			curField7zhge = unmarshalMsgFieldOrder7zhge[nextMiss7zhge]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zhge)
		switch curField7zhge {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zhge[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zhge[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zvto uint32
				zvto, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zvto) {
					z.Fields = (z.Fields)[:zvto]
				} else {
					z.Fields = make([]Field, zvto)
				}
				for zkns := range z.Fields {
					bts, err = z.Fields[zkns].UnmarshalMsg(bts)
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
	if nextMiss7zhge != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zhge = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zhge = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zkns := range z.Fields {
		s += z.Fields[zkns].Msgsize()
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
		var zgvz uint64
		zgvz, err = dc.ReadUint64()
		(*z) = Zkind(zgvz)
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
		var zens uint64
		zens, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zens)
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
	const maxFields8zajc = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zajc uint32
	totalEncodedFields8zajc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zajc := totalEncodedFields8zajc
	missingFieldsLeft8zajc := maxFields8zajc - totalEncodedFields8zajc

	var nextMiss8zajc int32 = -1
	var found8zajc [maxFields8zajc]bool
	var curField8zajc string

doneWithStruct8zajc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zajc > 0 || missingFieldsLeft8zajc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zajc, missingFieldsLeft8zajc, msgp.ShowFound(found8zajc[:]), decodeMsgFieldOrder8zajc)
		if encodedFieldsLeft8zajc > 0 {
			encodedFieldsLeft8zajc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zajc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zajc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zajc = 0
			}
			for nextMiss8zajc < maxFields8zajc && (found8zajc[nextMiss8zajc] || decodeMsgFieldSkip8zajc[nextMiss8zajc]) {
				nextMiss8zajc++
			}
			if nextMiss8zajc == maxFields8zajc {
				// filled all the empty fields!
				break doneWithStruct8zajc
			}
			missingFieldsLeft8zajc--
			curField8zajc = decodeMsgFieldOrder8zajc[nextMiss8zajc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zajc)
		switch curField8zajc {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zajc[0] = true
			{
				var zuym uint64
				zuym, err = dc.ReadUint64()
				z.Kind = Zkind(zuym)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zajc[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zajc[2] = true
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
			found8zajc[3] = true
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
	if nextMiss8zajc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zajc = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zajc = []bool{false, false, false, false}

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
	var empty_zihh [4]bool
	fieldsInUse_zpnx := z.fieldsNotEmpty(empty_zihh[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpnx)
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
	if !empty_zihh[1] {
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

	if !empty_zihh[2] {
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

	if !empty_zihh[3] {
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
	const maxFields9zfgf = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zfgf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zfgf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zfgf := totalEncodedFields9zfgf
	missingFieldsLeft9zfgf := maxFields9zfgf - totalEncodedFields9zfgf

	var nextMiss9zfgf int32 = -1
	var found9zfgf [maxFields9zfgf]bool
	var curField9zfgf string

doneWithStruct9zfgf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zfgf > 0 || missingFieldsLeft9zfgf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zfgf, missingFieldsLeft9zfgf, msgp.ShowFound(found9zfgf[:]), unmarshalMsgFieldOrder9zfgf)
		if encodedFieldsLeft9zfgf > 0 {
			encodedFieldsLeft9zfgf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zfgf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zfgf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zfgf = 0
			}
			for nextMiss9zfgf < maxFields9zfgf && (found9zfgf[nextMiss9zfgf] || unmarshalMsgFieldSkip9zfgf[nextMiss9zfgf]) {
				nextMiss9zfgf++
			}
			if nextMiss9zfgf == maxFields9zfgf {
				// filled all the empty fields!
				break doneWithStruct9zfgf
			}
			missingFieldsLeft9zfgf--
			curField9zfgf = unmarshalMsgFieldOrder9zfgf[nextMiss9zfgf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zfgf)
		switch curField9zfgf {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zfgf[0] = true
			{
				var zsgq uint64
				zsgq, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zsgq)
			}
		case "Str":
			found9zfgf[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zfgf[2] = true
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
			found9zfgf[3] = true
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
	if nextMiss9zfgf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zfgf = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zfgf = []bool{false, false, false, false}

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
