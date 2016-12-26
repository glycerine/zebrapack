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
	const maxFields0ztsn = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0ztsn uint32
	totalEncodedFields0ztsn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0ztsn := totalEncodedFields0ztsn
	missingFieldsLeft0ztsn := maxFields0ztsn - totalEncodedFields0ztsn

	var nextMiss0ztsn int32 = -1
	var found0ztsn [maxFields0ztsn]bool
	var curField0ztsn string

doneWithStruct0ztsn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ztsn > 0 || missingFieldsLeft0ztsn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ztsn, missingFieldsLeft0ztsn, msgp.ShowFound(found0ztsn[:]), decodeMsgFieldOrder0ztsn)
		if encodedFieldsLeft0ztsn > 0 {
			encodedFieldsLeft0ztsn--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0ztsn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ztsn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0ztsn = 0
			}
			for nextMiss0ztsn < maxFields0ztsn && (found0ztsn[nextMiss0ztsn] || decodeMsgFieldSkip0ztsn[nextMiss0ztsn]) {
				nextMiss0ztsn++
			}
			if nextMiss0ztsn == maxFields0ztsn {
				// filled all the empty fields!
				break doneWithStruct0ztsn
			}
			missingFieldsLeft0ztsn--
			curField0ztsn = decodeMsgFieldOrder0ztsn[nextMiss0ztsn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ztsn)
		switch curField0ztsn {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0ztsn[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0ztsn[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0ztsn[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0ztsn[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0ztsn[4] = true
			{
				var zvvz uint64
				zvvz, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zvvz)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0ztsn[5] = true
			{
				var zitm uint64
				zitm, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zitm)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0ztsn[6] = true
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
			found0ztsn[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0ztsn[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0ztsn[9] = true
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
	if nextMiss0ztsn != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0ztsn = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0ztsn = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zwic [10]bool
	fieldsInUse_zeio := z.fieldsNotEmpty(empty_zwic[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zeio)
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
	if !empty_zwic[3] {
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

	if !empty_zwic[4] {
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

	if !empty_zwic[5] {
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

	if !empty_zwic[6] {
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

	if !empty_zwic[7] {
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

	if !empty_zwic[8] {
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

	if !empty_zwic[9] {
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
	const maxFields1zrfn = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zrfn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zrfn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zrfn := totalEncodedFields1zrfn
	missingFieldsLeft1zrfn := maxFields1zrfn - totalEncodedFields1zrfn

	var nextMiss1zrfn int32 = -1
	var found1zrfn [maxFields1zrfn]bool
	var curField1zrfn string

doneWithStruct1zrfn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zrfn > 0 || missingFieldsLeft1zrfn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zrfn, missingFieldsLeft1zrfn, msgp.ShowFound(found1zrfn[:]), unmarshalMsgFieldOrder1zrfn)
		if encodedFieldsLeft1zrfn > 0 {
			encodedFieldsLeft1zrfn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zrfn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zrfn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zrfn = 0
			}
			for nextMiss1zrfn < maxFields1zrfn && (found1zrfn[nextMiss1zrfn] || unmarshalMsgFieldSkip1zrfn[nextMiss1zrfn]) {
				nextMiss1zrfn++
			}
			if nextMiss1zrfn == maxFields1zrfn {
				// filled all the empty fields!
				break doneWithStruct1zrfn
			}
			missingFieldsLeft1zrfn--
			curField1zrfn = unmarshalMsgFieldOrder1zrfn[nextMiss1zrfn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zrfn)
		switch curField1zrfn {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zrfn[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zrfn[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zrfn[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zrfn[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zrfn[4] = true
			{
				var zffv uint64
				zffv, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zffv)
			}
		case "FieldPrimitive":
			found1zrfn[5] = true
			{
				var zzwr uint64
				zzwr, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zzwr)
			}
		case "FieldFullType":
			found1zrfn[6] = true
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
			found1zrfn[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zrfn[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zrfn[9] = true
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
	if nextMiss1zrfn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zrfn = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zrfn = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zapj = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zapj uint32
	totalEncodedFields2zapj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zapj := totalEncodedFields2zapj
	missingFieldsLeft2zapj := maxFields2zapj - totalEncodedFields2zapj

	var nextMiss2zapj int32 = -1
	var found2zapj [maxFields2zapj]bool
	var curField2zapj string

doneWithStruct2zapj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zapj > 0 || missingFieldsLeft2zapj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zapj, missingFieldsLeft2zapj, msgp.ShowFound(found2zapj[:]), decodeMsgFieldOrder2zapj)
		if encodedFieldsLeft2zapj > 0 {
			encodedFieldsLeft2zapj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zapj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zapj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zapj = 0
			}
			for nextMiss2zapj < maxFields2zapj && (found2zapj[nextMiss2zapj] || decodeMsgFieldSkip2zapj[nextMiss2zapj]) {
				nextMiss2zapj++
			}
			if nextMiss2zapj == maxFields2zapj {
				// filled all the empty fields!
				break doneWithStruct2zapj
			}
			missingFieldsLeft2zapj--
			curField2zapj = decodeMsgFieldOrder2zapj[nextMiss2zapj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zapj)
		switch curField2zapj {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zapj[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zapj[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zapj[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zapj[3] = true
			var zubr uint32
			zubr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zubr) {
				z.Structs = (z.Structs)[:zubr]
			} else {
				z.Structs = make([]Struct, zubr)
			}
			for zpug := range z.Structs {
				const maxFields3zvll = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zvll uint32
				totalEncodedFields3zvll, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zvll := totalEncodedFields3zvll
				missingFieldsLeft3zvll := maxFields3zvll - totalEncodedFields3zvll

				var nextMiss3zvll int32 = -1
				var found3zvll [maxFields3zvll]bool
				var curField3zvll string

			doneWithStruct3zvll:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zvll > 0 || missingFieldsLeft3zvll > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvll, missingFieldsLeft3zvll, msgp.ShowFound(found3zvll[:]), decodeMsgFieldOrder3zvll)
					if encodedFieldsLeft3zvll > 0 {
						encodedFieldsLeft3zvll--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zvll = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zvll < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zvll = 0
						}
						for nextMiss3zvll < maxFields3zvll && (found3zvll[nextMiss3zvll] || decodeMsgFieldSkip3zvll[nextMiss3zvll]) {
							nextMiss3zvll++
						}
						if nextMiss3zvll == maxFields3zvll {
							// filled all the empty fields!
							break doneWithStruct3zvll
						}
						missingFieldsLeft3zvll--
						curField3zvll = decodeMsgFieldOrder3zvll[nextMiss3zvll]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zvll)
					switch curField3zvll {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zvll[0] = true
						z.Structs[zpug].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zvll[1] = true
						var zcyb uint32
						zcyb, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zpug].Fields) >= int(zcyb) {
							z.Structs[zpug].Fields = (z.Structs[zpug].Fields)[:zcyb]
						} else {
							z.Structs[zpug].Fields = make([]Field, zcyb)
						}
						for zooz := range z.Structs[zpug].Fields {
							err = z.Structs[zpug].Fields[zooz].DecodeMsg(dc)
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
				if nextMiss3zvll != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zapj[4] = true
			var zqpe uint32
			zqpe, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zqpe) {
				z.Imports = (z.Imports)[:zqpe]
			} else {
				z.Imports = make([]string, zqpe)
			}
			for zvsj := range z.Imports {
				z.Imports[zvsj], err = dc.ReadString()
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
	if nextMiss2zapj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zapj = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zapj = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zvll = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zvll = []bool{false, false}

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
	for zpug := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zpug].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zpug].Fields)))
		if err != nil {
			panic(err)
		}
		for zooz := range z.Structs[zpug].Fields {
			err = z.Structs[zpug].Fields[zooz].EncodeMsg(en)
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
	for zvsj := range z.Imports {
		err = en.WriteString(z.Imports[zvsj])
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
	for zpug := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zpug].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zpug].Fields)))
		for zooz := range z.Structs[zpug].Fields {
			o, err = z.Structs[zpug].Fields[zooz].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zvsj := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zvsj])
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
	const maxFields4zuiq = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zuiq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zuiq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zuiq := totalEncodedFields4zuiq
	missingFieldsLeft4zuiq := maxFields4zuiq - totalEncodedFields4zuiq

	var nextMiss4zuiq int32 = -1
	var found4zuiq [maxFields4zuiq]bool
	var curField4zuiq string

doneWithStruct4zuiq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zuiq > 0 || missingFieldsLeft4zuiq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zuiq, missingFieldsLeft4zuiq, msgp.ShowFound(found4zuiq[:]), unmarshalMsgFieldOrder4zuiq)
		if encodedFieldsLeft4zuiq > 0 {
			encodedFieldsLeft4zuiq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zuiq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zuiq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zuiq = 0
			}
			for nextMiss4zuiq < maxFields4zuiq && (found4zuiq[nextMiss4zuiq] || unmarshalMsgFieldSkip4zuiq[nextMiss4zuiq]) {
				nextMiss4zuiq++
			}
			if nextMiss4zuiq == maxFields4zuiq {
				// filled all the empty fields!
				break doneWithStruct4zuiq
			}
			missingFieldsLeft4zuiq--
			curField4zuiq = unmarshalMsgFieldOrder4zuiq[nextMiss4zuiq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zuiq)
		switch curField4zuiq {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zuiq[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zuiq[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zuiq[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zuiq[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zgva uint32
				zgva, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zgva) {
					z.Structs = (z.Structs)[:zgva]
				} else {
					z.Structs = make([]Struct, zgva)
				}
				for zpug := range z.Structs {
					const maxFields5zgaz = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zgaz uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zgaz, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zgaz := totalEncodedFields5zgaz
					missingFieldsLeft5zgaz := maxFields5zgaz - totalEncodedFields5zgaz

					var nextMiss5zgaz int32 = -1
					var found5zgaz [maxFields5zgaz]bool
					var curField5zgaz string

				doneWithStruct5zgaz:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zgaz > 0 || missingFieldsLeft5zgaz > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zgaz, missingFieldsLeft5zgaz, msgp.ShowFound(found5zgaz[:]), unmarshalMsgFieldOrder5zgaz)
						if encodedFieldsLeft5zgaz > 0 {
							encodedFieldsLeft5zgaz--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zgaz = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zgaz < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zgaz = 0
							}
							for nextMiss5zgaz < maxFields5zgaz && (found5zgaz[nextMiss5zgaz] || unmarshalMsgFieldSkip5zgaz[nextMiss5zgaz]) {
								nextMiss5zgaz++
							}
							if nextMiss5zgaz == maxFields5zgaz {
								// filled all the empty fields!
								break doneWithStruct5zgaz
							}
							missingFieldsLeft5zgaz--
							curField5zgaz = unmarshalMsgFieldOrder5zgaz[nextMiss5zgaz]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zgaz)
						switch curField5zgaz {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zgaz[0] = true
							z.Structs[zpug].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zgaz[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zpug].Fields) = (z.Structs[zpug].Fields)[:0]
							} else {

								var zgkk uint32
								zgkk, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zpug].Fields) >= int(zgkk) {
									z.Structs[zpug].Fields = (z.Structs[zpug].Fields)[:zgkk]
								} else {
									z.Structs[zpug].Fields = make([]Field, zgkk)
								}
								for zooz := range z.Structs[zpug].Fields {
									bts, err = z.Structs[zpug].Fields[zooz].UnmarshalMsg(bts)
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
					if nextMiss5zgaz != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4zuiq[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zhwq uint32
				zhwq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zhwq) {
					z.Imports = (z.Imports)[:zhwq]
				} else {
					z.Imports = make([]string, zhwq)
				}
				for zvsj := range z.Imports {
					z.Imports[zvsj], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zuiq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zuiq = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zuiq = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zgaz = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zgaz = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zpug := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zpug].StructName) + 7 + msgp.ArrayHeaderSize
		for zooz := range z.Structs[zpug].Fields {
			s += z.Structs[zpug].Fields[zooz].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zvsj := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zvsj])
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
	const maxFields6zsef = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zsef uint32
	totalEncodedFields6zsef, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zsef := totalEncodedFields6zsef
	missingFieldsLeft6zsef := maxFields6zsef - totalEncodedFields6zsef

	var nextMiss6zsef int32 = -1
	var found6zsef [maxFields6zsef]bool
	var curField6zsef string

doneWithStruct6zsef:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zsef > 0 || missingFieldsLeft6zsef > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zsef, missingFieldsLeft6zsef, msgp.ShowFound(found6zsef[:]), decodeMsgFieldOrder6zsef)
		if encodedFieldsLeft6zsef > 0 {
			encodedFieldsLeft6zsef--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zsef = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zsef < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zsef = 0
			}
			for nextMiss6zsef < maxFields6zsef && (found6zsef[nextMiss6zsef] || decodeMsgFieldSkip6zsef[nextMiss6zsef]) {
				nextMiss6zsef++
			}
			if nextMiss6zsef == maxFields6zsef {
				// filled all the empty fields!
				break doneWithStruct6zsef
			}
			missingFieldsLeft6zsef--
			curField6zsef = decodeMsgFieldOrder6zsef[nextMiss6zsef]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zsef)
		switch curField6zsef {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zsef[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zsef[1] = true
			var zeib uint32
			zeib, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zeib) {
				z.Fields = (z.Fields)[:zeib]
			} else {
				z.Fields = make([]Field, zeib)
			}
			for zdem := range z.Fields {
				err = z.Fields[zdem].DecodeMsg(dc)
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
	if nextMiss6zsef != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zsef = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zsef = []bool{false, false}

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
	for zdem := range z.Fields {
		err = z.Fields[zdem].EncodeMsg(en)
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
	for zdem := range z.Fields {
		o, err = z.Fields[zdem].MarshalMsg(o)
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
	const maxFields7zhvg = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zhvg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zhvg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zhvg := totalEncodedFields7zhvg
	missingFieldsLeft7zhvg := maxFields7zhvg - totalEncodedFields7zhvg

	var nextMiss7zhvg int32 = -1
	var found7zhvg [maxFields7zhvg]bool
	var curField7zhvg string

doneWithStruct7zhvg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zhvg > 0 || missingFieldsLeft7zhvg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zhvg, missingFieldsLeft7zhvg, msgp.ShowFound(found7zhvg[:]), unmarshalMsgFieldOrder7zhvg)
		if encodedFieldsLeft7zhvg > 0 {
			encodedFieldsLeft7zhvg--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zhvg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zhvg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zhvg = 0
			}
			for nextMiss7zhvg < maxFields7zhvg && (found7zhvg[nextMiss7zhvg] || unmarshalMsgFieldSkip7zhvg[nextMiss7zhvg]) {
				nextMiss7zhvg++
			}
			if nextMiss7zhvg == maxFields7zhvg {
				// filled all the empty fields!
				break doneWithStruct7zhvg
			}
			missingFieldsLeft7zhvg--
			curField7zhvg = unmarshalMsgFieldOrder7zhvg[nextMiss7zhvg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zhvg)
		switch curField7zhvg {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zhvg[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zhvg[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zpjb uint32
				zpjb, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zpjb) {
					z.Fields = (z.Fields)[:zpjb]
				} else {
					z.Fields = make([]Field, zpjb)
				}
				for zdem := range z.Fields {
					bts, err = z.Fields[zdem].UnmarshalMsg(bts)
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
	if nextMiss7zhvg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zhvg = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zhvg = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zdem := range z.Fields {
		s += z.Fields[zdem].Msgsize()
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
		var zduh uint64
		zduh, err = dc.ReadUint64()
		(*z) = Zkind(zduh)
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
		var zyvq uint64
		zyvq, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zyvq)
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
	const maxFields8zjhi = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zjhi uint32
	totalEncodedFields8zjhi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zjhi := totalEncodedFields8zjhi
	missingFieldsLeft8zjhi := maxFields8zjhi - totalEncodedFields8zjhi

	var nextMiss8zjhi int32 = -1
	var found8zjhi [maxFields8zjhi]bool
	var curField8zjhi string

doneWithStruct8zjhi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zjhi > 0 || missingFieldsLeft8zjhi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zjhi, missingFieldsLeft8zjhi, msgp.ShowFound(found8zjhi[:]), decodeMsgFieldOrder8zjhi)
		if encodedFieldsLeft8zjhi > 0 {
			encodedFieldsLeft8zjhi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zjhi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zjhi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zjhi = 0
			}
			for nextMiss8zjhi < maxFields8zjhi && (found8zjhi[nextMiss8zjhi] || decodeMsgFieldSkip8zjhi[nextMiss8zjhi]) {
				nextMiss8zjhi++
			}
			if nextMiss8zjhi == maxFields8zjhi {
				// filled all the empty fields!
				break doneWithStruct8zjhi
			}
			missingFieldsLeft8zjhi--
			curField8zjhi = decodeMsgFieldOrder8zjhi[nextMiss8zjhi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zjhi)
		switch curField8zjhi {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zjhi[0] = true
			{
				var zhfx uint64
				zhfx, err = dc.ReadUint64()
				z.Kind = Zkind(zhfx)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zjhi[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zjhi[2] = true
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
			found8zjhi[3] = true
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
	if nextMiss8zjhi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zjhi = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zjhi = []bool{false, false, false, false}

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
	var empty_zxcb [4]bool
	fieldsInUse_zsih := z.fieldsNotEmpty(empty_zxcb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zsih)
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
	if !empty_zxcb[1] {
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

	if !empty_zxcb[2] {
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

	if !empty_zxcb[3] {
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
	const maxFields9zllc = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zllc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zllc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zllc := totalEncodedFields9zllc
	missingFieldsLeft9zllc := maxFields9zllc - totalEncodedFields9zllc

	var nextMiss9zllc int32 = -1
	var found9zllc [maxFields9zllc]bool
	var curField9zllc string

doneWithStruct9zllc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zllc > 0 || missingFieldsLeft9zllc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zllc, missingFieldsLeft9zllc, msgp.ShowFound(found9zllc[:]), unmarshalMsgFieldOrder9zllc)
		if encodedFieldsLeft9zllc > 0 {
			encodedFieldsLeft9zllc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zllc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zllc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zllc = 0
			}
			for nextMiss9zllc < maxFields9zllc && (found9zllc[nextMiss9zllc] || unmarshalMsgFieldSkip9zllc[nextMiss9zllc]) {
				nextMiss9zllc++
			}
			if nextMiss9zllc == maxFields9zllc {
				// filled all the empty fields!
				break doneWithStruct9zllc
			}
			missingFieldsLeft9zllc--
			curField9zllc = unmarshalMsgFieldOrder9zllc[nextMiss9zllc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zllc)
		switch curField9zllc {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zllc[0] = true
			{
				var zxqs uint64
				zxqs, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zxqs)
			}
		case "Str":
			found9zllc[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zllc[2] = true
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
			found9zllc[3] = true
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
	if nextMiss9zllc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zllc = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zllc = []bool{false, false, false, false}

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
