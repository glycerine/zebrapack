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
	const maxFields0zewi = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zewi uint32
	totalEncodedFields0zewi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zewi := totalEncodedFields0zewi
	missingFieldsLeft0zewi := maxFields0zewi - totalEncodedFields0zewi

	var nextMiss0zewi int32 = -1
	var found0zewi [maxFields0zewi]bool
	var curField0zewi string

doneWithStruct0zewi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zewi > 0 || missingFieldsLeft0zewi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zewi, missingFieldsLeft0zewi, msgp.ShowFound(found0zewi[:]), decodeMsgFieldOrder0zewi)
		if encodedFieldsLeft0zewi > 0 {
			encodedFieldsLeft0zewi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zewi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zewi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zewi = 0
			}
			for nextMiss0zewi < maxFields0zewi && (found0zewi[nextMiss0zewi] || decodeMsgFieldSkip0zewi[nextMiss0zewi]) {
				nextMiss0zewi++
			}
			if nextMiss0zewi == maxFields0zewi {
				// filled all the empty fields!
				break doneWithStruct0zewi
			}
			missingFieldsLeft0zewi--
			curField0zewi = decodeMsgFieldOrder0zewi[nextMiss0zewi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zewi)
		switch curField0zewi {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zewi[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zewi[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zewi[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zewi[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zewi[4] = true
			{
				var zjdh uint64
				zjdh, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zjdh)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zewi[5] = true
			{
				var ztzu uint64
				ztzu, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(ztzu)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zewi[6] = true
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
			found0zewi[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zewi[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zewi[9] = true
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
	if nextMiss0zewi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zewi = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zewi = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zwrr [10]bool
	fieldsInUse_zmux := z.fieldsNotEmpty(empty_zwrr[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zmux)
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
	if !empty_zwrr[3] {
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

	if !empty_zwrr[4] {
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

	if !empty_zwrr[5] {
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

	if !empty_zwrr[6] {
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

	if !empty_zwrr[7] {
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

	if !empty_zwrr[8] {
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

	if !empty_zwrr[9] {
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
	const maxFields1zclj = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zclj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zclj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zclj := totalEncodedFields1zclj
	missingFieldsLeft1zclj := maxFields1zclj - totalEncodedFields1zclj

	var nextMiss1zclj int32 = -1
	var found1zclj [maxFields1zclj]bool
	var curField1zclj string

doneWithStruct1zclj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zclj > 0 || missingFieldsLeft1zclj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zclj, missingFieldsLeft1zclj, msgp.ShowFound(found1zclj[:]), unmarshalMsgFieldOrder1zclj)
		if encodedFieldsLeft1zclj > 0 {
			encodedFieldsLeft1zclj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zclj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zclj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zclj = 0
			}
			for nextMiss1zclj < maxFields1zclj && (found1zclj[nextMiss1zclj] || unmarshalMsgFieldSkip1zclj[nextMiss1zclj]) {
				nextMiss1zclj++
			}
			if nextMiss1zclj == maxFields1zclj {
				// filled all the empty fields!
				break doneWithStruct1zclj
			}
			missingFieldsLeft1zclj--
			curField1zclj = unmarshalMsgFieldOrder1zclj[nextMiss1zclj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zclj)
		switch curField1zclj {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zclj[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zclj[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zclj[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zclj[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zclj[4] = true
			{
				var znrz uint64
				znrz, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(znrz)
			}
		case "FieldPrimitive":
			found1zclj[5] = true
			{
				var zczo uint64
				zczo, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zczo)
			}
		case "FieldFullType":
			found1zclj[6] = true
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
			found1zclj[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zclj[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zclj[9] = true
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
	if nextMiss1zclj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zclj = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zclj = []bool{false, false, false, false, false, false, false, false, false, false}

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
func (z *KS) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zvbl = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zvbl uint32
	totalEncodedFields2zvbl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zvbl := totalEncodedFields2zvbl
	missingFieldsLeft2zvbl := maxFields2zvbl - totalEncodedFields2zvbl

	var nextMiss2zvbl int32 = -1
	var found2zvbl [maxFields2zvbl]bool
	var curField2zvbl string

doneWithStruct2zvbl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zvbl > 0 || missingFieldsLeft2zvbl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zvbl, missingFieldsLeft2zvbl, msgp.ShowFound(found2zvbl[:]), decodeMsgFieldOrder2zvbl)
		if encodedFieldsLeft2zvbl > 0 {
			encodedFieldsLeft2zvbl--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zvbl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zvbl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zvbl = 0
			}
			for nextMiss2zvbl < maxFields2zvbl && (found2zvbl[nextMiss2zvbl] || decodeMsgFieldSkip2zvbl[nextMiss2zvbl]) {
				nextMiss2zvbl++
			}
			if nextMiss2zvbl == maxFields2zvbl {
				// filled all the empty fields!
				break doneWithStruct2zvbl
			}
			missingFieldsLeft2zvbl--
			curField2zvbl = decodeMsgFieldOrder2zvbl[nextMiss2zvbl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zvbl)
		switch curField2zvbl {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found2zvbl[0] = true
			{
				var zkyz uint64
				zkyz, err = dc.ReadUint64()
				z.Kind = Zkind(zkyz)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found2zvbl[1] = true
			z.Str, err = dc.ReadString()
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
	if nextMiss2zvbl != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of KS
var decodeMsgFieldOrder2zvbl = []string{"Kind", "Str"}

var decodeMsgFieldSkip2zvbl = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z KS) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z KS) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Kind"
	err = en.Append(0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Kind))
	if err != nil {
		panic(err)
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Str)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z KS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Kind"
	o = append(o, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KS) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *KS) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zlfm = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zlfm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zlfm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zlfm := totalEncodedFields3zlfm
	missingFieldsLeft3zlfm := maxFields3zlfm - totalEncodedFields3zlfm

	var nextMiss3zlfm int32 = -1
	var found3zlfm [maxFields3zlfm]bool
	var curField3zlfm string

doneWithStruct3zlfm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zlfm > 0 || missingFieldsLeft3zlfm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zlfm, missingFieldsLeft3zlfm, msgp.ShowFound(found3zlfm[:]), unmarshalMsgFieldOrder3zlfm)
		if encodedFieldsLeft3zlfm > 0 {
			encodedFieldsLeft3zlfm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zlfm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zlfm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zlfm = 0
			}
			for nextMiss3zlfm < maxFields3zlfm && (found3zlfm[nextMiss3zlfm] || unmarshalMsgFieldSkip3zlfm[nextMiss3zlfm]) {
				nextMiss3zlfm++
			}
			if nextMiss3zlfm == maxFields3zlfm {
				// filled all the empty fields!
				break doneWithStruct3zlfm
			}
			missingFieldsLeft3zlfm--
			curField3zlfm = unmarshalMsgFieldOrder3zlfm[nextMiss3zlfm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zlfm)
		switch curField3zlfm {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found3zlfm[0] = true
			{
				var zjuz uint64
				zjuz, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zjuz)
			}
		case "Str":
			found3zlfm[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss3zlfm != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of KS
var unmarshalMsgFieldOrder3zlfm = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip3zlfm = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z KS) Msgsize() (s int) {
	s = 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Str)
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
	const maxFields4zaeo = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zaeo uint32
	totalEncodedFields4zaeo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zaeo := totalEncodedFields4zaeo
	missingFieldsLeft4zaeo := maxFields4zaeo - totalEncodedFields4zaeo

	var nextMiss4zaeo int32 = -1
	var found4zaeo [maxFields4zaeo]bool
	var curField4zaeo string

doneWithStruct4zaeo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zaeo > 0 || missingFieldsLeft4zaeo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zaeo, missingFieldsLeft4zaeo, msgp.ShowFound(found4zaeo[:]), decodeMsgFieldOrder4zaeo)
		if encodedFieldsLeft4zaeo > 0 {
			encodedFieldsLeft4zaeo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zaeo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zaeo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zaeo = 0
			}
			for nextMiss4zaeo < maxFields4zaeo && (found4zaeo[nextMiss4zaeo] || decodeMsgFieldSkip4zaeo[nextMiss4zaeo]) {
				nextMiss4zaeo++
			}
			if nextMiss4zaeo == maxFields4zaeo {
				// filled all the empty fields!
				break doneWithStruct4zaeo
			}
			missingFieldsLeft4zaeo--
			curField4zaeo = decodeMsgFieldOrder4zaeo[nextMiss4zaeo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zaeo)
		switch curField4zaeo {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found4zaeo[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zaeo[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zaeo[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zaeo[3] = true
			var zpiy uint32
			zpiy, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zpiy) {
				z.Structs = (z.Structs)[:zpiy]
			} else {
				z.Structs = make([]Struct, zpiy)
			}
			for zchs := range z.Structs {
				const maxFields5ztpq = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5ztpq uint32
				totalEncodedFields5ztpq, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5ztpq := totalEncodedFields5ztpq
				missingFieldsLeft5ztpq := maxFields5ztpq - totalEncodedFields5ztpq

				var nextMiss5ztpq int32 = -1
				var found5ztpq [maxFields5ztpq]bool
				var curField5ztpq string

			doneWithStruct5ztpq:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5ztpq > 0 || missingFieldsLeft5ztpq > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5ztpq, missingFieldsLeft5ztpq, msgp.ShowFound(found5ztpq[:]), decodeMsgFieldOrder5ztpq)
					if encodedFieldsLeft5ztpq > 0 {
						encodedFieldsLeft5ztpq--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5ztpq = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5ztpq < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5ztpq = 0
						}
						for nextMiss5ztpq < maxFields5ztpq && (found5ztpq[nextMiss5ztpq] || decodeMsgFieldSkip5ztpq[nextMiss5ztpq]) {
							nextMiss5ztpq++
						}
						if nextMiss5ztpq == maxFields5ztpq {
							// filled all the empty fields!
							break doneWithStruct5ztpq
						}
						missingFieldsLeft5ztpq--
						curField5ztpq = decodeMsgFieldOrder5ztpq[nextMiss5ztpq]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5ztpq)
					switch curField5ztpq {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5ztpq[0] = true
						z.Structs[zchs].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found5ztpq[1] = true
						var zatd uint32
						zatd, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zchs].Fields) >= int(zatd) {
							z.Structs[zchs].Fields = (z.Structs[zchs].Fields)[:zatd]
						} else {
							z.Structs[zchs].Fields = make([]Field, zatd)
						}
						for zrzm := range z.Structs[zchs].Fields {
							err = z.Structs[zchs].Fields[zrzm].DecodeMsg(dc)
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
				if nextMiss5ztpq != -1 {
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
	if nextMiss4zaeo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder4zaeo = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip4zaeo = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder5ztpq = []string{"StructName", "Fields"}

var decodeMsgFieldSkip5ztpq = []bool{false, false}

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
	for zchs := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zchs].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zchs].Fields)))
		if err != nil {
			panic(err)
		}
		for zrzm := range z.Structs[zchs].Fields {
			err = z.Structs[zchs].Fields[zrzm].EncodeMsg(en)
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
	for zchs := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zchs].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zchs].Fields)))
		for zrzm := range z.Structs[zchs].Fields {
			o, err = z.Structs[zchs].Fields[zrzm].MarshalMsg(o)
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
	const maxFields6zrqh = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zrqh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zrqh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zrqh := totalEncodedFields6zrqh
	missingFieldsLeft6zrqh := maxFields6zrqh - totalEncodedFields6zrqh

	var nextMiss6zrqh int32 = -1
	var found6zrqh [maxFields6zrqh]bool
	var curField6zrqh string

doneWithStruct6zrqh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zrqh > 0 || missingFieldsLeft6zrqh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zrqh, missingFieldsLeft6zrqh, msgp.ShowFound(found6zrqh[:]), unmarshalMsgFieldOrder6zrqh)
		if encodedFieldsLeft6zrqh > 0 {
			encodedFieldsLeft6zrqh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zrqh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zrqh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zrqh = 0
			}
			for nextMiss6zrqh < maxFields6zrqh && (found6zrqh[nextMiss6zrqh] || unmarshalMsgFieldSkip6zrqh[nextMiss6zrqh]) {
				nextMiss6zrqh++
			}
			if nextMiss6zrqh == maxFields6zrqh {
				// filled all the empty fields!
				break doneWithStruct6zrqh
			}
			missingFieldsLeft6zrqh--
			curField6zrqh = unmarshalMsgFieldOrder6zrqh[nextMiss6zrqh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zrqh)
		switch curField6zrqh {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found6zrqh[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found6zrqh[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found6zrqh[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zrqh[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zcnn uint32
				zcnn, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zcnn) {
					z.Structs = (z.Structs)[:zcnn]
				} else {
					z.Structs = make([]Struct, zcnn)
				}
				for zchs := range z.Structs {
					const maxFields7zmvq = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zmvq uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zmvq, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zmvq := totalEncodedFields7zmvq
					missingFieldsLeft7zmvq := maxFields7zmvq - totalEncodedFields7zmvq

					var nextMiss7zmvq int32 = -1
					var found7zmvq [maxFields7zmvq]bool
					var curField7zmvq string

				doneWithStruct7zmvq:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zmvq > 0 || missingFieldsLeft7zmvq > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zmvq, missingFieldsLeft7zmvq, msgp.ShowFound(found7zmvq[:]), unmarshalMsgFieldOrder7zmvq)
						if encodedFieldsLeft7zmvq > 0 {
							encodedFieldsLeft7zmvq--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zmvq = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zmvq < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zmvq = 0
							}
							for nextMiss7zmvq < maxFields7zmvq && (found7zmvq[nextMiss7zmvq] || unmarshalMsgFieldSkip7zmvq[nextMiss7zmvq]) {
								nextMiss7zmvq++
							}
							if nextMiss7zmvq == maxFields7zmvq {
								// filled all the empty fields!
								break doneWithStruct7zmvq
							}
							missingFieldsLeft7zmvq--
							curField7zmvq = unmarshalMsgFieldOrder7zmvq[nextMiss7zmvq]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zmvq)
						switch curField7zmvq {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zmvq[0] = true
							z.Structs[zchs].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found7zmvq[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zchs].Fields) = (z.Structs[zchs].Fields)[:0]
							} else {

								var zosg uint32
								zosg, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zchs].Fields) >= int(zosg) {
									z.Structs[zchs].Fields = (z.Structs[zchs].Fields)[:zosg]
								} else {
									z.Structs[zchs].Fields = make([]Field, zosg)
								}
								for zrzm := range z.Structs[zchs].Fields {
									bts, err = z.Structs[zchs].Fields[zrzm].UnmarshalMsg(bts)
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
					if nextMiss7zmvq != -1 {
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
	if nextMiss6zrqh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder6zrqh = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip6zrqh = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder7zmvq = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zmvq = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zchs := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zchs].StructName) + 7 + msgp.ArrayHeaderSize
		for zrzm := range z.Structs[zchs].Fields {
			s += z.Structs[zchs].Fields[zrzm].Msgsize()
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
	const maxFields8zgmv = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zgmv uint32
	totalEncodedFields8zgmv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zgmv := totalEncodedFields8zgmv
	missingFieldsLeft8zgmv := maxFields8zgmv - totalEncodedFields8zgmv

	var nextMiss8zgmv int32 = -1
	var found8zgmv [maxFields8zgmv]bool
	var curField8zgmv string

doneWithStruct8zgmv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zgmv > 0 || missingFieldsLeft8zgmv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zgmv, missingFieldsLeft8zgmv, msgp.ShowFound(found8zgmv[:]), decodeMsgFieldOrder8zgmv)
		if encodedFieldsLeft8zgmv > 0 {
			encodedFieldsLeft8zgmv--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zgmv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zgmv < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zgmv = 0
			}
			for nextMiss8zgmv < maxFields8zgmv && (found8zgmv[nextMiss8zgmv] || decodeMsgFieldSkip8zgmv[nextMiss8zgmv]) {
				nextMiss8zgmv++
			}
			if nextMiss8zgmv == maxFields8zgmv {
				// filled all the empty fields!
				break doneWithStruct8zgmv
			}
			missingFieldsLeft8zgmv--
			curField8zgmv = decodeMsgFieldOrder8zgmv[nextMiss8zgmv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zgmv)
		switch curField8zgmv {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found8zgmv[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found8zgmv[1] = true
			var zmlz uint32
			zmlz, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zmlz) {
				z.Fields = (z.Fields)[:zmlz]
			} else {
				z.Fields = make([]Field, zmlz)
			}
			for zxhk := range z.Fields {
				err = z.Fields[zxhk].DecodeMsg(dc)
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
	if nextMiss8zgmv != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder8zgmv = []string{"StructName", "Fields"}

var decodeMsgFieldSkip8zgmv = []bool{false, false}

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
	for zxhk := range z.Fields {
		err = z.Fields[zxhk].EncodeMsg(en)
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
	for zxhk := range z.Fields {
		o, err = z.Fields[zxhk].MarshalMsg(o)
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
	const maxFields9zigr = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zigr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zigr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zigr := totalEncodedFields9zigr
	missingFieldsLeft9zigr := maxFields9zigr - totalEncodedFields9zigr

	var nextMiss9zigr int32 = -1
	var found9zigr [maxFields9zigr]bool
	var curField9zigr string

doneWithStruct9zigr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zigr > 0 || missingFieldsLeft9zigr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zigr, missingFieldsLeft9zigr, msgp.ShowFound(found9zigr[:]), unmarshalMsgFieldOrder9zigr)
		if encodedFieldsLeft9zigr > 0 {
			encodedFieldsLeft9zigr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zigr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zigr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zigr = 0
			}
			for nextMiss9zigr < maxFields9zigr && (found9zigr[nextMiss9zigr] || unmarshalMsgFieldSkip9zigr[nextMiss9zigr]) {
				nextMiss9zigr++
			}
			if nextMiss9zigr == maxFields9zigr {
				// filled all the empty fields!
				break doneWithStruct9zigr
			}
			missingFieldsLeft9zigr--
			curField9zigr = unmarshalMsgFieldOrder9zigr[nextMiss9zigr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zigr)
		switch curField9zigr {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found9zigr[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found9zigr[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zolj uint32
				zolj, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zolj) {
					z.Fields = (z.Fields)[:zolj]
				} else {
					z.Fields = make([]Field, zolj)
				}
				for zxhk := range z.Fields {
					bts, err = z.Fields[zxhk].UnmarshalMsg(bts)
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
	if nextMiss9zigr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder9zigr = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip9zigr = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zxhk := range z.Fields {
		s += z.Fields[zxhk].Msgsize()
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
		var zdvh uint64
		zdvh, err = dc.ReadUint64()
		(*z) = Zkind(zdvh)
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
		var zguu uint64
		zguu, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zguu)
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
	const maxFields10zzbs = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zzbs uint32
	totalEncodedFields10zzbs, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zzbs := totalEncodedFields10zzbs
	missingFieldsLeft10zzbs := maxFields10zzbs - totalEncodedFields10zzbs

	var nextMiss10zzbs int32 = -1
	var found10zzbs [maxFields10zzbs]bool
	var curField10zzbs string

doneWithStruct10zzbs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zzbs > 0 || missingFieldsLeft10zzbs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zzbs, missingFieldsLeft10zzbs, msgp.ShowFound(found10zzbs[:]), decodeMsgFieldOrder10zzbs)
		if encodedFieldsLeft10zzbs > 0 {
			encodedFieldsLeft10zzbs--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zzbs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zzbs < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zzbs = 0
			}
			for nextMiss10zzbs < maxFields10zzbs && (found10zzbs[nextMiss10zzbs] || decodeMsgFieldSkip10zzbs[nextMiss10zzbs]) {
				nextMiss10zzbs++
			}
			if nextMiss10zzbs == maxFields10zzbs {
				// filled all the empty fields!
				break doneWithStruct10zzbs
			}
			missingFieldsLeft10zzbs--
			curField10zzbs = decodeMsgFieldOrder10zzbs[nextMiss10zzbs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zzbs)
		switch curField10zzbs {
		// -- templateDecodeMsg ends here --

		case "Name":
			found10zzbs[0] = true
			const maxFields11zowb = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields11zowb uint32
			totalEncodedFields11zowb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft11zowb := totalEncodedFields11zowb
			missingFieldsLeft11zowb := maxFields11zowb - totalEncodedFields11zowb

			var nextMiss11zowb int32 = -1
			var found11zowb [maxFields11zowb]bool
			var curField11zowb string

		doneWithStruct11zowb:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zowb > 0 || missingFieldsLeft11zowb > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zowb, missingFieldsLeft11zowb, msgp.ShowFound(found11zowb[:]), decodeMsgFieldOrder11zowb)
				if encodedFieldsLeft11zowb > 0 {
					encodedFieldsLeft11zowb--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField11zowb = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zowb < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss11zowb = 0
					}
					for nextMiss11zowb < maxFields11zowb && (found11zowb[nextMiss11zowb] || decodeMsgFieldSkip11zowb[nextMiss11zowb]) {
						nextMiss11zowb++
					}
					if nextMiss11zowb == maxFields11zowb {
						// filled all the empty fields!
						break doneWithStruct11zowb
					}
					missingFieldsLeft11zowb--
					curField11zowb = decodeMsgFieldOrder11zowb[nextMiss11zowb]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zowb)
				switch curField11zowb {
				// -- templateDecodeMsg ends here --

				case "Kind":
					found11zowb[0] = true
					{
						var ztdd uint64
						ztdd, err = dc.ReadUint64()
						z.Name.Kind = Zkind(ztdd)
					}
					if err != nil {
						panic(err)
					}
				case "Str":
					found11zowb[1] = true
					z.Name.Str, err = dc.ReadString()
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
			if nextMiss11zowb != -1 {
				dc.PopAlwaysNil()
			}

		case "Domain":
			found10zzbs[1] = true
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
			found10zzbs[2] = true
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
	if nextMiss10zzbs != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder10zzbs = []string{"Name", "Domain", "Range"}

var decodeMsgFieldSkip10zzbs = []bool{false, false, false}

// fields of KS
var decodeMsgFieldOrder11zowb = []string{"Kind", "Str"}

var decodeMsgFieldSkip11zowb = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Ztype) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[1] = (z.Domain == nil) // pointer, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Range == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Ztype) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zdye [3]bool
	fieldsInUse_zjbs := z.fieldsNotEmpty(empty_zdye[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjbs)
	if err != nil {
		return err
	}

	// write "Name"
	// map header, size 2
	// write "Kind"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Name.Kind))
	if err != nil {
		panic(err)
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name.Str)
	if err != nil {
		panic(err)
	}
	if !empty_zdye[1] {
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

	if !empty_zdye[2] {
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
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Name"
	// map header, size 2
	// string "Kind"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Name.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Name.Str)
	if !empty[1] {
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

	if !empty[2] {
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
	const maxFields12znop = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields12znop uint32
	if !nbs.AlwaysNil {
		totalEncodedFields12znop, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft12znop := totalEncodedFields12znop
	missingFieldsLeft12znop := maxFields12znop - totalEncodedFields12znop

	var nextMiss12znop int32 = -1
	var found12znop [maxFields12znop]bool
	var curField12znop string

doneWithStruct12znop:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12znop > 0 || missingFieldsLeft12znop > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12znop, missingFieldsLeft12znop, msgp.ShowFound(found12znop[:]), unmarshalMsgFieldOrder12znop)
		if encodedFieldsLeft12znop > 0 {
			encodedFieldsLeft12znop--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField12znop = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12znop < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss12znop = 0
			}
			for nextMiss12znop < maxFields12znop && (found12znop[nextMiss12znop] || unmarshalMsgFieldSkip12znop[nextMiss12znop]) {
				nextMiss12znop++
			}
			if nextMiss12znop == maxFields12znop {
				// filled all the empty fields!
				break doneWithStruct12znop
			}
			missingFieldsLeft12znop--
			curField12znop = unmarshalMsgFieldOrder12znop[nextMiss12znop]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12znop)
		switch curField12znop {
		// -- templateUnmarshalMsg ends here --

		case "Name":
			found12znop[0] = true
			const maxFields13zoer = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields13zoer uint32
			if !nbs.AlwaysNil {
				totalEncodedFields13zoer, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft13zoer := totalEncodedFields13zoer
			missingFieldsLeft13zoer := maxFields13zoer - totalEncodedFields13zoer

			var nextMiss13zoer int32 = -1
			var found13zoer [maxFields13zoer]bool
			var curField13zoer string

		doneWithStruct13zoer:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft13zoer > 0 || missingFieldsLeft13zoer > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zoer, missingFieldsLeft13zoer, msgp.ShowFound(found13zoer[:]), unmarshalMsgFieldOrder13zoer)
				if encodedFieldsLeft13zoer > 0 {
					encodedFieldsLeft13zoer--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField13zoer = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss13zoer < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss13zoer = 0
					}
					for nextMiss13zoer < maxFields13zoer && (found13zoer[nextMiss13zoer] || unmarshalMsgFieldSkip13zoer[nextMiss13zoer]) {
						nextMiss13zoer++
					}
					if nextMiss13zoer == maxFields13zoer {
						// filled all the empty fields!
						break doneWithStruct13zoer
					}
					missingFieldsLeft13zoer--
					curField13zoer = unmarshalMsgFieldOrder13zoer[nextMiss13zoer]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField13zoer)
				switch curField13zoer {
				// -- templateUnmarshalMsg ends here --

				case "Kind":
					found13zoer[0] = true
					{
						var zcor uint64
						zcor, bts, err = nbs.ReadUint64Bytes(bts)

						if err != nil {
							panic(err)
						}
						z.Name.Kind = Zkind(zcor)
					}
				case "Str":
					found13zoer[1] = true
					z.Name.Str, bts, err = nbs.ReadStringBytes(bts)

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
			if nextMiss13zoer != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Domain":
			found12znop[1] = true
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
			found12znop[2] = true
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
	if nextMiss12znop != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder12znop = []string{"Name", "Domain", "Range"}

var unmarshalMsgFieldSkip12znop = []bool{false, false, false}

// fields of KS
var unmarshalMsgFieldOrder13zoer = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip13zoer = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Ztype) Msgsize() (s int) {
	s = 1 + 5 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Name.Str) + 7
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
