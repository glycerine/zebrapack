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
	const maxFields0znrt = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0znrt uint32
	totalEncodedFields0znrt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0znrt := totalEncodedFields0znrt
	missingFieldsLeft0znrt := maxFields0znrt - totalEncodedFields0znrt

	var nextMiss0znrt int32 = -1
	var found0znrt [maxFields0znrt]bool
	var curField0znrt string

doneWithStruct0znrt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0znrt > 0 || missingFieldsLeft0znrt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0znrt, missingFieldsLeft0znrt, msgp.ShowFound(found0znrt[:]), decodeMsgFieldOrder0znrt)
		if encodedFieldsLeft0znrt > 0 {
			encodedFieldsLeft0znrt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0znrt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0znrt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0znrt = 0
			}
			for nextMiss0znrt < maxFields0znrt && (found0znrt[nextMiss0znrt] || decodeMsgFieldSkip0znrt[nextMiss0znrt]) {
				nextMiss0znrt++
			}
			if nextMiss0znrt == maxFields0znrt {
				// filled all the empty fields!
				break doneWithStruct0znrt
			}
			missingFieldsLeft0znrt--
			curField0znrt = decodeMsgFieldOrder0znrt[nextMiss0znrt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0znrt)
		switch curField0znrt {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0znrt[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0znrt[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0znrt[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0znrt[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0znrt[4] = true
			{
				var zrjl uint64
				zrjl, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zrjl)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0znrt[5] = true
			{
				var zymg uint64
				zymg, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zymg)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0znrt[6] = true
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
			found0znrt[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0znrt[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0znrt[9] = true
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
	if nextMiss0znrt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0znrt = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0znrt = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zors [10]bool
	fieldsInUse_zunn := z.fieldsNotEmpty(empty_zors[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zunn)
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
	if !empty_zors[3] {
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

	if !empty_zors[4] {
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

	if !empty_zors[5] {
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

	if !empty_zors[6] {
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

	if !empty_zors[7] {
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

	if !empty_zors[8] {
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

	if !empty_zors[9] {
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
	const maxFields1zoka = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zoka uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zoka, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zoka := totalEncodedFields1zoka
	missingFieldsLeft1zoka := maxFields1zoka - totalEncodedFields1zoka

	var nextMiss1zoka int32 = -1
	var found1zoka [maxFields1zoka]bool
	var curField1zoka string

doneWithStruct1zoka:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zoka > 0 || missingFieldsLeft1zoka > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zoka, missingFieldsLeft1zoka, msgp.ShowFound(found1zoka[:]), unmarshalMsgFieldOrder1zoka)
		if encodedFieldsLeft1zoka > 0 {
			encodedFieldsLeft1zoka--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zoka = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zoka < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zoka = 0
			}
			for nextMiss1zoka < maxFields1zoka && (found1zoka[nextMiss1zoka] || unmarshalMsgFieldSkip1zoka[nextMiss1zoka]) {
				nextMiss1zoka++
			}
			if nextMiss1zoka == maxFields1zoka {
				// filled all the empty fields!
				break doneWithStruct1zoka
			}
			missingFieldsLeft1zoka--
			curField1zoka = unmarshalMsgFieldOrder1zoka[nextMiss1zoka]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zoka)
		switch curField1zoka {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zoka[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zoka[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zoka[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zoka[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zoka[4] = true
			{
				var znpi uint64
				znpi, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(znpi)
			}
		case "FieldPrimitive":
			found1zoka[5] = true
			{
				var zoez uint64
				zoez, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zoez)
			}
		case "FieldFullType":
			found1zoka[6] = true
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
			found1zoka[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zoka[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zoka[9] = true
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
	if nextMiss1zoka != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zoka = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zoka = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zgjl = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zgjl uint32
	totalEncodedFields2zgjl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zgjl := totalEncodedFields2zgjl
	missingFieldsLeft2zgjl := maxFields2zgjl - totalEncodedFields2zgjl

	var nextMiss2zgjl int32 = -1
	var found2zgjl [maxFields2zgjl]bool
	var curField2zgjl string

doneWithStruct2zgjl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zgjl > 0 || missingFieldsLeft2zgjl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zgjl, missingFieldsLeft2zgjl, msgp.ShowFound(found2zgjl[:]), decodeMsgFieldOrder2zgjl)
		if encodedFieldsLeft2zgjl > 0 {
			encodedFieldsLeft2zgjl--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zgjl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zgjl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zgjl = 0
			}
			for nextMiss2zgjl < maxFields2zgjl && (found2zgjl[nextMiss2zgjl] || decodeMsgFieldSkip2zgjl[nextMiss2zgjl]) {
				nextMiss2zgjl++
			}
			if nextMiss2zgjl == maxFields2zgjl {
				// filled all the empty fields!
				break doneWithStruct2zgjl
			}
			missingFieldsLeft2zgjl--
			curField2zgjl = decodeMsgFieldOrder2zgjl[nextMiss2zgjl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zgjl)
		switch curField2zgjl {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zgjl[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zgjl[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zgjl[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zgjl[3] = true
			var zjln uint32
			zjln, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zjln > 0 {
				z.Structs = make(map[string]*Struct, zjln)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zjln > 0 {
				zjln--
				var zumw string
				var zvqu *Struct
				zumw, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					zvqu = nil
				} else {
					if zvqu == nil {
						zvqu = new(Struct)
					}
					const maxFields3zwiu = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zwiu uint32
					totalEncodedFields3zwiu, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zwiu := totalEncodedFields3zwiu
					missingFieldsLeft3zwiu := maxFields3zwiu - totalEncodedFields3zwiu

					var nextMiss3zwiu int32 = -1
					var found3zwiu [maxFields3zwiu]bool
					var curField3zwiu string

				doneWithStruct3zwiu:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zwiu > 0 || missingFieldsLeft3zwiu > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zwiu, missingFieldsLeft3zwiu, msgp.ShowFound(found3zwiu[:]), decodeMsgFieldOrder3zwiu)
						if encodedFieldsLeft3zwiu > 0 {
							encodedFieldsLeft3zwiu--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zwiu = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zwiu < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zwiu = 0
							}
							for nextMiss3zwiu < maxFields3zwiu && (found3zwiu[nextMiss3zwiu] || decodeMsgFieldSkip3zwiu[nextMiss3zwiu]) {
								nextMiss3zwiu++
							}
							if nextMiss3zwiu == maxFields3zwiu {
								// filled all the empty fields!
								break doneWithStruct3zwiu
							}
							missingFieldsLeft3zwiu--
							curField3zwiu = decodeMsgFieldOrder3zwiu[nextMiss3zwiu]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zwiu)
						switch curField3zwiu {
						// -- templateDecodeMsg ends here --

						case "StructName":
							found3zwiu[0] = true
							zvqu.StructName, err = dc.ReadString()
							if err != nil {
								panic(err)
							}
						case "Fields":
							found3zwiu[1] = true
							var zodk uint32
							zodk, err = dc.ReadArrayHeader()
							if err != nil {
								panic(err)
							}
							if cap(zvqu.Fields) >= int(zodk) {
								zvqu.Fields = (zvqu.Fields)[:zodk]
							} else {
								zvqu.Fields = make([]Field, zodk)
							}
							for zcxm := range zvqu.Fields {
								err = zvqu.Fields[zcxm].DecodeMsg(dc)
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
					if nextMiss3zwiu != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[zumw] = zvqu
			}
		case "Imports":
			found2zgjl[4] = true
			var zdbn uint32
			zdbn, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zdbn) {
				z.Imports = (z.Imports)[:zdbn]
			} else {
				z.Imports = make([]string, zdbn)
			}
			for zuuf := range z.Imports {
				z.Imports[zuuf], err = dc.ReadString()
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
	if nextMiss2zgjl != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zgjl = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zgjl = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zwiu = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zwiu = []bool{false, false}

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
	for zumw, zvqu := range z.Structs {
		err = en.WriteString(zumw)
		if err != nil {
			panic(err)
		}
		if zvqu == nil {
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
			err = en.WriteString(zvqu.StructName)
			if err != nil {
				panic(err)
			}
			// write "Fields"
			err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(zvqu.Fields)))
			if err != nil {
				panic(err)
			}
			for zcxm := range zvqu.Fields {
				err = zvqu.Fields[zcxm].EncodeMsg(en)
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
	for zuuf := range z.Imports {
		err = en.WriteString(z.Imports[zuuf])
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
	for zumw, zvqu := range z.Structs {
		o = msgp.AppendString(o, zumw)
		if zvqu == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName"
			o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, zvqu.StructName)
			// string "Fields"
			o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			o = msgp.AppendArrayHeader(o, uint32(len(zvqu.Fields)))
			for zcxm := range zvqu.Fields {
				o, err = zvqu.Fields[zcxm].MarshalMsg(o)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zuuf := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zuuf])
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
	const maxFields4zutc = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zutc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zutc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zutc := totalEncodedFields4zutc
	missingFieldsLeft4zutc := maxFields4zutc - totalEncodedFields4zutc

	var nextMiss4zutc int32 = -1
	var found4zutc [maxFields4zutc]bool
	var curField4zutc string

doneWithStruct4zutc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zutc > 0 || missingFieldsLeft4zutc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zutc, missingFieldsLeft4zutc, msgp.ShowFound(found4zutc[:]), unmarshalMsgFieldOrder4zutc)
		if encodedFieldsLeft4zutc > 0 {
			encodedFieldsLeft4zutc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zutc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zutc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zutc = 0
			}
			for nextMiss4zutc < maxFields4zutc && (found4zutc[nextMiss4zutc] || unmarshalMsgFieldSkip4zutc[nextMiss4zutc]) {
				nextMiss4zutc++
			}
			if nextMiss4zutc == maxFields4zutc {
				// filled all the empty fields!
				break doneWithStruct4zutc
			}
			missingFieldsLeft4zutc--
			curField4zutc = unmarshalMsgFieldOrder4zutc[nextMiss4zutc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zutc)
		switch curField4zutc {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zutc[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zutc[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zutc[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zutc[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zvdo uint32
				zvdo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zvdo > 0 {
					z.Structs = make(map[string]*Struct, zvdo)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zvdo > 0 {
					var zumw string
					var zvqu *Struct
					zvdo--
					zumw, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && zvqu == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if zvqu == nil {
							zvqu = new(Struct)
						}
						const maxFields5zwaq = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zwaq uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zwaq, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								panic(err)
								return
							}
						}
						encodedFieldsLeft5zwaq := totalEncodedFields5zwaq
						missingFieldsLeft5zwaq := maxFields5zwaq - totalEncodedFields5zwaq

						var nextMiss5zwaq int32 = -1
						var found5zwaq [maxFields5zwaq]bool
						var curField5zwaq string

					doneWithStruct5zwaq:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zwaq > 0 || missingFieldsLeft5zwaq > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zwaq, missingFieldsLeft5zwaq, msgp.ShowFound(found5zwaq[:]), unmarshalMsgFieldOrder5zwaq)
							if encodedFieldsLeft5zwaq > 0 {
								encodedFieldsLeft5zwaq--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									panic(err)
									return
								}
								curField5zwaq = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zwaq < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zwaq = 0
								}
								for nextMiss5zwaq < maxFields5zwaq && (found5zwaq[nextMiss5zwaq] || unmarshalMsgFieldSkip5zwaq[nextMiss5zwaq]) {
									nextMiss5zwaq++
								}
								if nextMiss5zwaq == maxFields5zwaq {
									// filled all the empty fields!
									break doneWithStruct5zwaq
								}
								missingFieldsLeft5zwaq--
								curField5zwaq = unmarshalMsgFieldOrder5zwaq[nextMiss5zwaq]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zwaq)
							switch curField5zwaq {
							// -- templateUnmarshalMsg ends here --

							case "StructName":
								found5zwaq[0] = true
								zvqu.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									panic(err)
								}
							case "Fields":
								found5zwaq[1] = true
								if nbs.AlwaysNil {
									(zvqu.Fields) = (zvqu.Fields)[:0]
								} else {

									var zkpd uint32
									zkpd, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										panic(err)
									}
									if cap(zvqu.Fields) >= int(zkpd) {
										zvqu.Fields = (zvqu.Fields)[:zkpd]
									} else {
										zvqu.Fields = make([]Field, zkpd)
									}
									for zcxm := range zvqu.Fields {
										bts, err = zvqu.Fields[zcxm].UnmarshalMsg(bts)
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
						if nextMiss5zwaq != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[zumw] = zvqu
				}
			}
		case "Imports":
			found4zutc[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zucp uint32
				zucp, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zucp) {
					z.Imports = (z.Imports)[:zucp]
				} else {
					z.Imports = make([]string, zucp)
				}
				for zuuf := range z.Imports {
					z.Imports[zuuf], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zutc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zutc = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zutc = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zwaq = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zwaq = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zumw, zvqu := range z.Structs {
			_ = zvqu
			_ = zumw
			s += msgp.StringPrefixSize + len(zumw)
			if zvqu == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 11 + msgp.StringPrefixSize + len(zvqu.StructName) + 7 + msgp.ArrayHeaderSize
				for zcxm := range zvqu.Fields {
					s += zvqu.Fields[zcxm].Msgsize()
				}
			}
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zuuf := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zuuf])
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
	const maxFields6znee = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6znee uint32
	totalEncodedFields6znee, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6znee := totalEncodedFields6znee
	missingFieldsLeft6znee := maxFields6znee - totalEncodedFields6znee

	var nextMiss6znee int32 = -1
	var found6znee [maxFields6znee]bool
	var curField6znee string

doneWithStruct6znee:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6znee > 0 || missingFieldsLeft6znee > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6znee, missingFieldsLeft6znee, msgp.ShowFound(found6znee[:]), decodeMsgFieldOrder6znee)
		if encodedFieldsLeft6znee > 0 {
			encodedFieldsLeft6znee--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6znee = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6znee < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6znee = 0
			}
			for nextMiss6znee < maxFields6znee && (found6znee[nextMiss6znee] || decodeMsgFieldSkip6znee[nextMiss6znee]) {
				nextMiss6znee++
			}
			if nextMiss6znee == maxFields6znee {
				// filled all the empty fields!
				break doneWithStruct6znee
			}
			missingFieldsLeft6znee--
			curField6znee = decodeMsgFieldOrder6znee[nextMiss6znee]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6znee)
		switch curField6znee {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6znee[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6znee[1] = true
			var ztws uint32
			ztws, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(ztws) {
				z.Fields = (z.Fields)[:ztws]
			} else {
				z.Fields = make([]Field, ztws)
			}
			for zsph := range z.Fields {
				err = z.Fields[zsph].DecodeMsg(dc)
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
	if nextMiss6znee != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6znee = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6znee = []bool{false, false}

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
	for zsph := range z.Fields {
		err = z.Fields[zsph].EncodeMsg(en)
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
	for zsph := range z.Fields {
		o, err = z.Fields[zsph].MarshalMsg(o)
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
	const maxFields7zgsa = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zgsa uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zgsa, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zgsa := totalEncodedFields7zgsa
	missingFieldsLeft7zgsa := maxFields7zgsa - totalEncodedFields7zgsa

	var nextMiss7zgsa int32 = -1
	var found7zgsa [maxFields7zgsa]bool
	var curField7zgsa string

doneWithStruct7zgsa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zgsa > 0 || missingFieldsLeft7zgsa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgsa, missingFieldsLeft7zgsa, msgp.ShowFound(found7zgsa[:]), unmarshalMsgFieldOrder7zgsa)
		if encodedFieldsLeft7zgsa > 0 {
			encodedFieldsLeft7zgsa--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zgsa = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zgsa < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zgsa = 0
			}
			for nextMiss7zgsa < maxFields7zgsa && (found7zgsa[nextMiss7zgsa] || unmarshalMsgFieldSkip7zgsa[nextMiss7zgsa]) {
				nextMiss7zgsa++
			}
			if nextMiss7zgsa == maxFields7zgsa {
				// filled all the empty fields!
				break doneWithStruct7zgsa
			}
			missingFieldsLeft7zgsa--
			curField7zgsa = unmarshalMsgFieldOrder7zgsa[nextMiss7zgsa]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zgsa)
		switch curField7zgsa {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zgsa[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zgsa[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zqsd uint32
				zqsd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zqsd) {
					z.Fields = (z.Fields)[:zqsd]
				} else {
					z.Fields = make([]Field, zqsd)
				}
				for zsph := range z.Fields {
					bts, err = z.Fields[zsph].UnmarshalMsg(bts)
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
	if nextMiss7zgsa != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zgsa = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zgsa = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zsph := range z.Fields {
		s += z.Fields[zsph].Msgsize()
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
		var zodv uint64
		zodv, err = dc.ReadUint64()
		(*z) = Zkind(zodv)
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
		var zkca uint64
		zkca, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zkca)
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
	const maxFields8zjuc = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zjuc uint32
	totalEncodedFields8zjuc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zjuc := totalEncodedFields8zjuc
	missingFieldsLeft8zjuc := maxFields8zjuc - totalEncodedFields8zjuc

	var nextMiss8zjuc int32 = -1
	var found8zjuc [maxFields8zjuc]bool
	var curField8zjuc string

doneWithStruct8zjuc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zjuc > 0 || missingFieldsLeft8zjuc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zjuc, missingFieldsLeft8zjuc, msgp.ShowFound(found8zjuc[:]), decodeMsgFieldOrder8zjuc)
		if encodedFieldsLeft8zjuc > 0 {
			encodedFieldsLeft8zjuc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zjuc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zjuc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zjuc = 0
			}
			for nextMiss8zjuc < maxFields8zjuc && (found8zjuc[nextMiss8zjuc] || decodeMsgFieldSkip8zjuc[nextMiss8zjuc]) {
				nextMiss8zjuc++
			}
			if nextMiss8zjuc == maxFields8zjuc {
				// filled all the empty fields!
				break doneWithStruct8zjuc
			}
			missingFieldsLeft8zjuc--
			curField8zjuc = decodeMsgFieldOrder8zjuc[nextMiss8zjuc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zjuc)
		switch curField8zjuc {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zjuc[0] = true
			{
				var zxej uint64
				zxej, err = dc.ReadUint64()
				z.Kind = Zkind(zxej)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zjuc[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zjuc[2] = true
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
			found8zjuc[3] = true
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
	if nextMiss8zjuc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zjuc = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zjuc = []bool{false, false, false, false}

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
	var empty_zsuj [4]bool
	fieldsInUse_zeul := z.fieldsNotEmpty(empty_zsuj[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zeul)
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
	if !empty_zsuj[1] {
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

	if !empty_zsuj[2] {
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

	if !empty_zsuj[3] {
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
	const maxFields9zbbj = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zbbj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zbbj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zbbj := totalEncodedFields9zbbj
	missingFieldsLeft9zbbj := maxFields9zbbj - totalEncodedFields9zbbj

	var nextMiss9zbbj int32 = -1
	var found9zbbj [maxFields9zbbj]bool
	var curField9zbbj string

doneWithStruct9zbbj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zbbj > 0 || missingFieldsLeft9zbbj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zbbj, missingFieldsLeft9zbbj, msgp.ShowFound(found9zbbj[:]), unmarshalMsgFieldOrder9zbbj)
		if encodedFieldsLeft9zbbj > 0 {
			encodedFieldsLeft9zbbj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zbbj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zbbj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zbbj = 0
			}
			for nextMiss9zbbj < maxFields9zbbj && (found9zbbj[nextMiss9zbbj] || unmarshalMsgFieldSkip9zbbj[nextMiss9zbbj]) {
				nextMiss9zbbj++
			}
			if nextMiss9zbbj == maxFields9zbbj {
				// filled all the empty fields!
				break doneWithStruct9zbbj
			}
			missingFieldsLeft9zbbj--
			curField9zbbj = unmarshalMsgFieldOrder9zbbj[nextMiss9zbbj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zbbj)
		switch curField9zbbj {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zbbj[0] = true
			{
				var zvag uint64
				zvag, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zvag)
			}
		case "Str":
			found9zbbj[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zbbj[2] = true
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
			found9zbbj[3] = true
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
	if nextMiss9zbbj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zbbj = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zbbj = []bool{false, false, false, false}

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
