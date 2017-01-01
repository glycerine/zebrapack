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
	const maxFields0zgum = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zgum uint32
	totalEncodedFields0zgum, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zgum := totalEncodedFields0zgum
	missingFieldsLeft0zgum := maxFields0zgum - totalEncodedFields0zgum

	var nextMiss0zgum int32 = -1
	var found0zgum [maxFields0zgum]bool
	var curField0zgum string

doneWithStruct0zgum:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgum > 0 || missingFieldsLeft0zgum > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgum, missingFieldsLeft0zgum, msgp.ShowFound(found0zgum[:]), decodeMsgFieldOrder0zgum)
		if encodedFieldsLeft0zgum > 0 {
			encodedFieldsLeft0zgum--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zgum = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgum < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zgum = 0
			}
			for nextMiss0zgum < maxFields0zgum && (found0zgum[nextMiss0zgum] || decodeMsgFieldSkip0zgum[nextMiss0zgum]) {
				nextMiss0zgum++
			}
			if nextMiss0zgum == maxFields0zgum {
				// filled all the empty fields!
				break doneWithStruct0zgum
			}
			missingFieldsLeft0zgum--
			curField0zgum = decodeMsgFieldOrder0zgum[nextMiss0zgum]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgum)
		switch curField0zgum {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zgum[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zgum[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zgum[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zgum[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zgum[4] = true
			{
				var zgcn uint64
				zgcn, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zgcn)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zgum[5] = true
			{
				var zewo uint64
				zewo, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zewo)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zgum[6] = true
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
			found0zgum[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zgum[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zgum[9] = true
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
	if nextMiss0zgum != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zgum = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zgum = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zvvr [10]bool
	fieldsInUse_zhbz := z.fieldsNotEmpty(empty_zvvr[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zhbz)
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
	if !empty_zvvr[3] {
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

	if !empty_zvvr[4] {
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

	if !empty_zvvr[5] {
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

	if !empty_zvvr[6] {
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

	if !empty_zvvr[7] {
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

	if !empty_zvvr[8] {
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

	if !empty_zvvr[9] {
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
	const maxFields1zdqs = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zdqs uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zdqs, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zdqs := totalEncodedFields1zdqs
	missingFieldsLeft1zdqs := maxFields1zdqs - totalEncodedFields1zdqs

	var nextMiss1zdqs int32 = -1
	var found1zdqs [maxFields1zdqs]bool
	var curField1zdqs string

doneWithStruct1zdqs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zdqs > 0 || missingFieldsLeft1zdqs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zdqs, missingFieldsLeft1zdqs, msgp.ShowFound(found1zdqs[:]), unmarshalMsgFieldOrder1zdqs)
		if encodedFieldsLeft1zdqs > 0 {
			encodedFieldsLeft1zdqs--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zdqs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zdqs < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zdqs = 0
			}
			for nextMiss1zdqs < maxFields1zdqs && (found1zdqs[nextMiss1zdqs] || unmarshalMsgFieldSkip1zdqs[nextMiss1zdqs]) {
				nextMiss1zdqs++
			}
			if nextMiss1zdqs == maxFields1zdqs {
				// filled all the empty fields!
				break doneWithStruct1zdqs
			}
			missingFieldsLeft1zdqs--
			curField1zdqs = unmarshalMsgFieldOrder1zdqs[nextMiss1zdqs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zdqs)
		switch curField1zdqs {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zdqs[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zdqs[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zdqs[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zdqs[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zdqs[4] = true
			{
				var zbro uint64
				zbro, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zbro)
			}
		case "FieldPrimitive":
			found1zdqs[5] = true
			{
				var zbmj uint64
				zbmj, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zbmj)
			}
		case "FieldFullType":
			found1zdqs[6] = true
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
			found1zdqs[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zdqs[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zdqs[9] = true
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
	if nextMiss1zdqs != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zdqs = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zdqs = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zrjt = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zrjt uint32
	totalEncodedFields2zrjt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zrjt := totalEncodedFields2zrjt
	missingFieldsLeft2zrjt := maxFields2zrjt - totalEncodedFields2zrjt

	var nextMiss2zrjt int32 = -1
	var found2zrjt [maxFields2zrjt]bool
	var curField2zrjt string

doneWithStruct2zrjt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zrjt > 0 || missingFieldsLeft2zrjt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zrjt, missingFieldsLeft2zrjt, msgp.ShowFound(found2zrjt[:]), decodeMsgFieldOrder2zrjt)
		if encodedFieldsLeft2zrjt > 0 {
			encodedFieldsLeft2zrjt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zrjt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zrjt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zrjt = 0
			}
			for nextMiss2zrjt < maxFields2zrjt && (found2zrjt[nextMiss2zrjt] || decodeMsgFieldSkip2zrjt[nextMiss2zrjt]) {
				nextMiss2zrjt++
			}
			if nextMiss2zrjt == maxFields2zrjt {
				// filled all the empty fields!
				break doneWithStruct2zrjt
			}
			missingFieldsLeft2zrjt--
			curField2zrjt = decodeMsgFieldOrder2zrjt[nextMiss2zrjt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zrjt)
		switch curField2zrjt {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zrjt[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zrjt[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zrjt[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zrjt[3] = true
			var zqze uint32
			zqze, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zqze > 0 {
				z.Structs = make(map[string]*Struct, zqze)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zqze > 0 {
				zqze--
				var zksi string
				var zfky *Struct
				zksi, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					zfky = nil
				} else {
					if zfky == nil {
						zfky = new(Struct)
					}
					const maxFields3zvns = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zvns uint32
					totalEncodedFields3zvns, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zvns := totalEncodedFields3zvns
					missingFieldsLeft3zvns := maxFields3zvns - totalEncodedFields3zvns

					var nextMiss3zvns int32 = -1
					var found3zvns [maxFields3zvns]bool
					var curField3zvns string

				doneWithStruct3zvns:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zvns > 0 || missingFieldsLeft3zvns > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvns, missingFieldsLeft3zvns, msgp.ShowFound(found3zvns[:]), decodeMsgFieldOrder3zvns)
						if encodedFieldsLeft3zvns > 0 {
							encodedFieldsLeft3zvns--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zvns = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zvns < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zvns = 0
							}
							for nextMiss3zvns < maxFields3zvns && (found3zvns[nextMiss3zvns] || decodeMsgFieldSkip3zvns[nextMiss3zvns]) {
								nextMiss3zvns++
							}
							if nextMiss3zvns == maxFields3zvns {
								// filled all the empty fields!
								break doneWithStruct3zvns
							}
							missingFieldsLeft3zvns--
							curField3zvns = decodeMsgFieldOrder3zvns[nextMiss3zvns]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zvns)
						switch curField3zvns {
						// -- templateDecodeMsg ends here --

						case "StructName":
							found3zvns[0] = true
							zfky.StructName, err = dc.ReadString()
							if err != nil {
								panic(err)
							}
						case "Fields":
							found3zvns[1] = true
							var zdls uint32
							zdls, err = dc.ReadArrayHeader()
							if err != nil {
								panic(err)
							}
							if cap(zfky.Fields) >= int(zdls) {
								zfky.Fields = (zfky.Fields)[:zdls]
							} else {
								zfky.Fields = make([]Field, zdls)
							}
							for zbzy := range zfky.Fields {
								err = zfky.Fields[zbzy].DecodeMsg(dc)
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
					if nextMiss3zvns != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[zksi] = zfky
			}
		case "Imports":
			found2zrjt[4] = true
			var zffe uint32
			zffe, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zffe) {
				z.Imports = (z.Imports)[:zffe]
			} else {
				z.Imports = make([]string, zffe)
			}
			for zkkm := range z.Imports {
				z.Imports[zkkm], err = dc.ReadString()
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
	if nextMiss2zrjt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zrjt = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zrjt = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zvns = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zvns = []bool{false, false}

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
	for zksi, zfky := range z.Structs {
		err = en.WriteString(zksi)
		if err != nil {
			panic(err)
		}
		if zfky == nil {
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
			err = en.WriteString(zfky.StructName)
			if err != nil {
				panic(err)
			}
			// write "Fields"
			err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(zfky.Fields)))
			if err != nil {
				panic(err)
			}
			for zbzy := range zfky.Fields {
				err = zfky.Fields[zbzy].EncodeMsg(en)
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
	for zkkm := range z.Imports {
		err = en.WriteString(z.Imports[zkkm])
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
	for zksi, zfky := range z.Structs {
		o = msgp.AppendString(o, zksi)
		if zfky == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName"
			o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, zfky.StructName)
			// string "Fields"
			o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			o = msgp.AppendArrayHeader(o, uint32(len(zfky.Fields)))
			for zbzy := range zfky.Fields {
				o, err = zfky.Fields[zbzy].MarshalMsg(o)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zkkm := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zkkm])
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
	const maxFields4zlza = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zlza uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zlza, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zlza := totalEncodedFields4zlza
	missingFieldsLeft4zlza := maxFields4zlza - totalEncodedFields4zlza

	var nextMiss4zlza int32 = -1
	var found4zlza [maxFields4zlza]bool
	var curField4zlza string

doneWithStruct4zlza:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zlza > 0 || missingFieldsLeft4zlza > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zlza, missingFieldsLeft4zlza, msgp.ShowFound(found4zlza[:]), unmarshalMsgFieldOrder4zlza)
		if encodedFieldsLeft4zlza > 0 {
			encodedFieldsLeft4zlza--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zlza = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zlza < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zlza = 0
			}
			for nextMiss4zlza < maxFields4zlza && (found4zlza[nextMiss4zlza] || unmarshalMsgFieldSkip4zlza[nextMiss4zlza]) {
				nextMiss4zlza++
			}
			if nextMiss4zlza == maxFields4zlza {
				// filled all the empty fields!
				break doneWithStruct4zlza
			}
			missingFieldsLeft4zlza--
			curField4zlza = unmarshalMsgFieldOrder4zlza[nextMiss4zlza]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zlza)
		switch curField4zlza {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zlza[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zlza[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zlza[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zlza[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zaiz uint32
				zaiz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zaiz > 0 {
					z.Structs = make(map[string]*Struct, zaiz)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zaiz > 0 {
					var zksi string
					var zfky *Struct
					zaiz--
					zksi, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && zfky == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if zfky == nil {
							zfky = new(Struct)
						}
						const maxFields5zdrt = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zdrt uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zdrt, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								panic(err)
								return
							}
						}
						encodedFieldsLeft5zdrt := totalEncodedFields5zdrt
						missingFieldsLeft5zdrt := maxFields5zdrt - totalEncodedFields5zdrt

						var nextMiss5zdrt int32 = -1
						var found5zdrt [maxFields5zdrt]bool
						var curField5zdrt string

					doneWithStruct5zdrt:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zdrt > 0 || missingFieldsLeft5zdrt > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zdrt, missingFieldsLeft5zdrt, msgp.ShowFound(found5zdrt[:]), unmarshalMsgFieldOrder5zdrt)
							if encodedFieldsLeft5zdrt > 0 {
								encodedFieldsLeft5zdrt--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									panic(err)
									return
								}
								curField5zdrt = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zdrt < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zdrt = 0
								}
								for nextMiss5zdrt < maxFields5zdrt && (found5zdrt[nextMiss5zdrt] || unmarshalMsgFieldSkip5zdrt[nextMiss5zdrt]) {
									nextMiss5zdrt++
								}
								if nextMiss5zdrt == maxFields5zdrt {
									// filled all the empty fields!
									break doneWithStruct5zdrt
								}
								missingFieldsLeft5zdrt--
								curField5zdrt = unmarshalMsgFieldOrder5zdrt[nextMiss5zdrt]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zdrt)
							switch curField5zdrt {
							// -- templateUnmarshalMsg ends here --

							case "StructName":
								found5zdrt[0] = true
								zfky.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									panic(err)
								}
							case "Fields":
								found5zdrt[1] = true
								if nbs.AlwaysNil {
									(zfky.Fields) = (zfky.Fields)[:0]
								} else {

									var zsau uint32
									zsau, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										panic(err)
									}
									if cap(zfky.Fields) >= int(zsau) {
										zfky.Fields = (zfky.Fields)[:zsau]
									} else {
										zfky.Fields = make([]Field, zsau)
									}
									for zbzy := range zfky.Fields {
										bts, err = zfky.Fields[zbzy].UnmarshalMsg(bts)
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
						if nextMiss5zdrt != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[zksi] = zfky
				}
			}
		case "Imports":
			found4zlza[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zkpy uint32
				zkpy, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zkpy) {
					z.Imports = (z.Imports)[:zkpy]
				} else {
					z.Imports = make([]string, zkpy)
				}
				for zkkm := range z.Imports {
					z.Imports[zkkm], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zlza != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zlza = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zlza = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zdrt = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zdrt = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zksi, zfky := range z.Structs {
			_ = zfky
			_ = zksi
			s += msgp.StringPrefixSize + len(zksi)
			if zfky == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 11 + msgp.StringPrefixSize + len(zfky.StructName) + 7 + msgp.ArrayHeaderSize
				for zbzy := range zfky.Fields {
					s += zfky.Fields[zbzy].Msgsize()
				}
			}
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zkkm := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zkkm])
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
	const maxFields6ztaz = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6ztaz uint32
	totalEncodedFields6ztaz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6ztaz := totalEncodedFields6ztaz
	missingFieldsLeft6ztaz := maxFields6ztaz - totalEncodedFields6ztaz

	var nextMiss6ztaz int32 = -1
	var found6ztaz [maxFields6ztaz]bool
	var curField6ztaz string

doneWithStruct6ztaz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6ztaz > 0 || missingFieldsLeft6ztaz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6ztaz, missingFieldsLeft6ztaz, msgp.ShowFound(found6ztaz[:]), decodeMsgFieldOrder6ztaz)
		if encodedFieldsLeft6ztaz > 0 {
			encodedFieldsLeft6ztaz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6ztaz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6ztaz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6ztaz = 0
			}
			for nextMiss6ztaz < maxFields6ztaz && (found6ztaz[nextMiss6ztaz] || decodeMsgFieldSkip6ztaz[nextMiss6ztaz]) {
				nextMiss6ztaz++
			}
			if nextMiss6ztaz == maxFields6ztaz {
				// filled all the empty fields!
				break doneWithStruct6ztaz
			}
			missingFieldsLeft6ztaz--
			curField6ztaz = decodeMsgFieldOrder6ztaz[nextMiss6ztaz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6ztaz)
		switch curField6ztaz {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6ztaz[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6ztaz[1] = true
			var zeat uint32
			zeat, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zeat) {
				z.Fields = (z.Fields)[:zeat]
			} else {
				z.Fields = make([]Field, zeat)
			}
			for zxln := range z.Fields {
				err = z.Fields[zxln].DecodeMsg(dc)
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
	if nextMiss6ztaz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6ztaz = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6ztaz = []bool{false, false}

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
	for zxln := range z.Fields {
		err = z.Fields[zxln].EncodeMsg(en)
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
	for zxln := range z.Fields {
		o, err = z.Fields[zxln].MarshalMsg(o)
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
	const maxFields7zwsw = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zwsw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zwsw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zwsw := totalEncodedFields7zwsw
	missingFieldsLeft7zwsw := maxFields7zwsw - totalEncodedFields7zwsw

	var nextMiss7zwsw int32 = -1
	var found7zwsw [maxFields7zwsw]bool
	var curField7zwsw string

doneWithStruct7zwsw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zwsw > 0 || missingFieldsLeft7zwsw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zwsw, missingFieldsLeft7zwsw, msgp.ShowFound(found7zwsw[:]), unmarshalMsgFieldOrder7zwsw)
		if encodedFieldsLeft7zwsw > 0 {
			encodedFieldsLeft7zwsw--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zwsw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zwsw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zwsw = 0
			}
			for nextMiss7zwsw < maxFields7zwsw && (found7zwsw[nextMiss7zwsw] || unmarshalMsgFieldSkip7zwsw[nextMiss7zwsw]) {
				nextMiss7zwsw++
			}
			if nextMiss7zwsw == maxFields7zwsw {
				// filled all the empty fields!
				break doneWithStruct7zwsw
			}
			missingFieldsLeft7zwsw--
			curField7zwsw = unmarshalMsgFieldOrder7zwsw[nextMiss7zwsw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zwsw)
		switch curField7zwsw {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zwsw[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zwsw[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zjyl uint32
				zjyl, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zjyl) {
					z.Fields = (z.Fields)[:zjyl]
				} else {
					z.Fields = make([]Field, zjyl)
				}
				for zxln := range z.Fields {
					bts, err = z.Fields[zxln].UnmarshalMsg(bts)
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
	if nextMiss7zwsw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zwsw = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zwsw = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zxln := range z.Fields {
		s += z.Fields[zxln].Msgsize()
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
		var zgyg uint64
		zgyg, err = dc.ReadUint64()
		(*z) = Zkind(zgyg)
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
		var zrke uint64
		zrke, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zrke)
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
	const maxFields8zihi = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zihi uint32
	totalEncodedFields8zihi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zihi := totalEncodedFields8zihi
	missingFieldsLeft8zihi := maxFields8zihi - totalEncodedFields8zihi

	var nextMiss8zihi int32 = -1
	var found8zihi [maxFields8zihi]bool
	var curField8zihi string

doneWithStruct8zihi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zihi > 0 || missingFieldsLeft8zihi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zihi, missingFieldsLeft8zihi, msgp.ShowFound(found8zihi[:]), decodeMsgFieldOrder8zihi)
		if encodedFieldsLeft8zihi > 0 {
			encodedFieldsLeft8zihi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zihi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zihi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zihi = 0
			}
			for nextMiss8zihi < maxFields8zihi && (found8zihi[nextMiss8zihi] || decodeMsgFieldSkip8zihi[nextMiss8zihi]) {
				nextMiss8zihi++
			}
			if nextMiss8zihi == maxFields8zihi {
				// filled all the empty fields!
				break doneWithStruct8zihi
			}
			missingFieldsLeft8zihi--
			curField8zihi = decodeMsgFieldOrder8zihi[nextMiss8zihi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zihi)
		switch curField8zihi {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zihi[0] = true
			{
				var zbgi uint64
				zbgi, err = dc.ReadUint64()
				z.Kind = Zkind(zbgi)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zihi[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zihi[2] = true
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
			found8zihi[3] = true
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
	if nextMiss8zihi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zihi = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zihi = []bool{false, false, false, false}

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
	var empty_zrmc [4]bool
	fieldsInUse_zaik := z.fieldsNotEmpty(empty_zrmc[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zaik)
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
	if !empty_zrmc[1] {
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

	if !empty_zrmc[2] {
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

	if !empty_zrmc[3] {
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
	const maxFields9zufj = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zufj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zufj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zufj := totalEncodedFields9zufj
	missingFieldsLeft9zufj := maxFields9zufj - totalEncodedFields9zufj

	var nextMiss9zufj int32 = -1
	var found9zufj [maxFields9zufj]bool
	var curField9zufj string

doneWithStruct9zufj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zufj > 0 || missingFieldsLeft9zufj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zufj, missingFieldsLeft9zufj, msgp.ShowFound(found9zufj[:]), unmarshalMsgFieldOrder9zufj)
		if encodedFieldsLeft9zufj > 0 {
			encodedFieldsLeft9zufj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zufj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zufj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zufj = 0
			}
			for nextMiss9zufj < maxFields9zufj && (found9zufj[nextMiss9zufj] || unmarshalMsgFieldSkip9zufj[nextMiss9zufj]) {
				nextMiss9zufj++
			}
			if nextMiss9zufj == maxFields9zufj {
				// filled all the empty fields!
				break doneWithStruct9zufj
			}
			missingFieldsLeft9zufj--
			curField9zufj = unmarshalMsgFieldOrder9zufj[nextMiss9zufj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zufj)
		switch curField9zufj {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zufj[0] = true
			{
				var zpez uint64
				zpez, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zpez)
			}
		case "Str":
			found9zufj[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zufj[2] = true
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
			found9zufj[3] = true
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
	if nextMiss9zufj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zufj = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zufj = []bool{false, false, false, false}

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
	0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x84, 0xa5, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x82,
	0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x9a, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x5a, 0x69, 0x64, 0xac,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x5a, 0x69, 0x64,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2,
	0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64,
	0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72,
	0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x88, 0xa3,
	0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x5a, 0x6b, 0x69, 0x6e, 0x64, 0xad, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0xb, 0xad, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb,
	0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa9, 0x4f, 0x6d, 0x69, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0xa5, 0x5a, 0x6b, 0x69, 0x6e, 0x64, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0xb, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3,
	0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xad,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xad, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69,
	0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74,
	0x72, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0xc3, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4,
	0x62, 0x6f, 0x6f, 0x6c, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0x88,
	0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0xa4, 0x53, 0x6b, 0x69, 0x70, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6b, 0x69, 0x70, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53,
	0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0xc3, 0x88, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0xac,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74,
	0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0xc3, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x94, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x5a, 0x6b, 0x69, 0x6e, 0x64, 0xad, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0xb, 0xad, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
	0xb, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x88, 0xa3, 0x5a, 0x69,
	0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x53,
	0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa3,
	0x53, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79,
	0x70, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x1c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83,
	0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16,
	0xa3, 0x53, 0x74, 0x72, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0xac, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x2a, 0x5a,
	0x74, 0x79, 0x70, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x1c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
	0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xa9, 0x4f, 0x6d, 0x69,
	0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xc3, 0xa6, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x82, 0xaa,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x95, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xac, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82,
	0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xac, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
	0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87,
	0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0xac,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xad, 0x5a, 0x65, 0x62,
	0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
	0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x86, 0xa3, 0x5a, 0x69, 0x64,
	0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb2, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5d, 0x2a, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53,
	0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
	0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72,
	0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
	0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa8, 0x5b, 0x5d,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
	0x64, 0x2, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa6, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x92, 0x87,
	0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x2, 0xa3, 0x53, 0x74, 0x72,
	0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x72, 0xa7, 0x5b, 0x5d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xa3, 0x53, 0x74,
	0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4,
	0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x46, 0x69, 0x65, 0x6c, 0x64, 0xa7,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x90,
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 4152 bytes
func ZebraSchemaInJsonCompact() []byte { return zebraSchemaInJsonCompact }

var zebraSchemaInJsonCompact = []byte(`{"SourcePath":"zebra.go","SourcePackage":"zebra","ZebraSchemaId":463621516989988,"Structs":{"Field":{"StructName":"Field","Fields":[{"Zid":-1,"FieldGoName":"Zid","FieldTagName":"Zid","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":-1,"FieldGoName":"FieldGoName","FieldTagName":"FieldGoName","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"FieldTagName","FieldTagName":"FieldTagName","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"FieldTypeStr","FieldTagName":"FieldTypeStr","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"FieldCategory","FieldTagName":"FieldCategory","FieldTypeStr":"Zkind","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"FieldPrimitive","FieldTagName":"FieldPrimitive","FieldTypeStr":"Zkind","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"FieldFullType","FieldTagName":"FieldFullType","FieldTypeStr":"*Ztype","FieldCategory":28,"FieldFullType":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Ztype"}},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"OmitEmpty","FieldTagName":"OmitEmpty","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Skip","FieldTagName":"Skip","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Deprecated","FieldTagName":"Deprecated","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"},"OmitEmpty":true}]},"Ztype":{"StructName":"Ztype","Fields":[{"Zid":-1,"FieldGoName":"Kind","FieldTagName":"Kind","FieldTypeStr":"Zkind","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}},{"Zid":-1,"FieldGoName":"Str","FieldTagName":"Str","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Domain","FieldTagName":"Domain","FieldTypeStr":"*Ztype","FieldCategory":28,"FieldFullType":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Ztype"}},"OmitEmpty":true},{"Zid":-1,"FieldGoName":"Range","FieldTagName":"Range","FieldTypeStr":"*Ztype","FieldCategory":28,"FieldFullType":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Ztype"}},"OmitEmpty":true}]},"Schema":{"StructName":"Schema","Fields":[{"Zid":-1,"FieldGoName":"SourcePath","FieldTagName":"SourcePath","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"SourcePackage","FieldTagName":"SourcePackage","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"ZebraSchemaId","FieldTagName":"ZebraSchemaId","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":-1,"FieldGoName":"Structs","FieldTagName":"Structs","FieldTypeStr":"map[string]*Struct","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":25,"Str":"Struct"}}}},{"Zid":-1,"FieldGoName":"Imports","FieldTagName":"Imports","FieldTypeStr":"[]string","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":2,"Str":"string"}}}]},"Struct":{"StructName":"Struct","Fields":[{"Zid":-1,"FieldGoName":"StructName","FieldTagName":"StructName","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"Fields","FieldTagName":"Fields","FieldTypeStr":"[]Field","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":22,"Str":"Field"}}}]}},"Imports":[]}`)

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 10901 bytes
func ZebraSchemaInJsonPretty() []byte { return zebraSchemaInJsonPretty }

var zebraSchemaInJsonPretty = []byte(`{
    "SourcePath": "zebra.go",
    "SourcePackage": "zebra",
    "ZebraSchemaId": 463621516989988,
    "Structs": {
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
        },
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
        }
    },
    "Imports": []
}`)
