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
	const maxFields0zbte = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zbte uint32
	totalEncodedFields0zbte, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zbte := totalEncodedFields0zbte
	missingFieldsLeft0zbte := maxFields0zbte - totalEncodedFields0zbte

	var nextMiss0zbte int32 = -1
	var found0zbte [maxFields0zbte]bool
	var curField0zbte string

doneWithStruct0zbte:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zbte > 0 || missingFieldsLeft0zbte > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zbte, missingFieldsLeft0zbte, msgp.ShowFound(found0zbte[:]), decodeMsgFieldOrder0zbte)
		if encodedFieldsLeft0zbte > 0 {
			encodedFieldsLeft0zbte--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zbte = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zbte < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zbte = 0
			}
			for nextMiss0zbte < maxFields0zbte && (found0zbte[nextMiss0zbte] || decodeMsgFieldSkip0zbte[nextMiss0zbte]) {
				nextMiss0zbte++
			}
			if nextMiss0zbte == maxFields0zbte {
				// filled all the empty fields!
				break doneWithStruct0zbte
			}
			missingFieldsLeft0zbte--
			curField0zbte = decodeMsgFieldOrder0zbte[nextMiss0zbte]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zbte)
		switch curField0zbte {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zbte[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zbte[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zbte[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zbte[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zbte[4] = true
			{
				var zljm uint64
				zljm, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zljm)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zbte[5] = true
			{
				var zhct uint64
				zhct, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zhct)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zbte[6] = true
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
			found0zbte[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zbte[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zbte[9] = true
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
	if nextMiss0zbte != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zbte = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zbte = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zsmt [10]bool
	fieldsInUse_zzdi := z.fieldsNotEmpty(empty_zsmt[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zzdi)
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
	if !empty_zsmt[3] {
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

	if !empty_zsmt[4] {
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

	if !empty_zsmt[5] {
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

	if !empty_zsmt[6] {
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

	if !empty_zsmt[7] {
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

	if !empty_zsmt[8] {
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

	if !empty_zsmt[9] {
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
	const maxFields1zcsp = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zcsp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zcsp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zcsp := totalEncodedFields1zcsp
	missingFieldsLeft1zcsp := maxFields1zcsp - totalEncodedFields1zcsp

	var nextMiss1zcsp int32 = -1
	var found1zcsp [maxFields1zcsp]bool
	var curField1zcsp string

doneWithStruct1zcsp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zcsp > 0 || missingFieldsLeft1zcsp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zcsp, missingFieldsLeft1zcsp, msgp.ShowFound(found1zcsp[:]), unmarshalMsgFieldOrder1zcsp)
		if encodedFieldsLeft1zcsp > 0 {
			encodedFieldsLeft1zcsp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zcsp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zcsp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zcsp = 0
			}
			for nextMiss1zcsp < maxFields1zcsp && (found1zcsp[nextMiss1zcsp] || unmarshalMsgFieldSkip1zcsp[nextMiss1zcsp]) {
				nextMiss1zcsp++
			}
			if nextMiss1zcsp == maxFields1zcsp {
				// filled all the empty fields!
				break doneWithStruct1zcsp
			}
			missingFieldsLeft1zcsp--
			curField1zcsp = unmarshalMsgFieldOrder1zcsp[nextMiss1zcsp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zcsp)
		switch curField1zcsp {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zcsp[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zcsp[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zcsp[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zcsp[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zcsp[4] = true
			{
				var zppe uint64
				zppe, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zppe)
			}
		case "FieldPrimitive":
			found1zcsp[5] = true
			{
				var zowt uint64
				zowt, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zowt)
			}
		case "FieldFullType":
			found1zcsp[6] = true
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
			found1zcsp[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zcsp[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zcsp[9] = true
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
	if nextMiss1zcsp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zcsp = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zcsp = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zqdy = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zqdy uint32
	totalEncodedFields2zqdy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zqdy := totalEncodedFields2zqdy
	missingFieldsLeft2zqdy := maxFields2zqdy - totalEncodedFields2zqdy

	var nextMiss2zqdy int32 = -1
	var found2zqdy [maxFields2zqdy]bool
	var curField2zqdy string

doneWithStruct2zqdy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zqdy > 0 || missingFieldsLeft2zqdy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zqdy, missingFieldsLeft2zqdy, msgp.ShowFound(found2zqdy[:]), decodeMsgFieldOrder2zqdy)
		if encodedFieldsLeft2zqdy > 0 {
			encodedFieldsLeft2zqdy--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zqdy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zqdy < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zqdy = 0
			}
			for nextMiss2zqdy < maxFields2zqdy && (found2zqdy[nextMiss2zqdy] || decodeMsgFieldSkip2zqdy[nextMiss2zqdy]) {
				nextMiss2zqdy++
			}
			if nextMiss2zqdy == maxFields2zqdy {
				// filled all the empty fields!
				break doneWithStruct2zqdy
			}
			missingFieldsLeft2zqdy--
			curField2zqdy = decodeMsgFieldOrder2zqdy[nextMiss2zqdy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zqdy)
		switch curField2zqdy {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found2zqdy[0] = true
			{
				var zvqk uint64
				zvqk, err = dc.ReadUint64()
				z.Kind = Zkind(zvqk)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found2zqdy[1] = true
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
	if nextMiss2zqdy != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of KS
var decodeMsgFieldOrder2zqdy = []string{"Kind", "Str"}

var decodeMsgFieldSkip2zqdy = []bool{false, false}

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
	const maxFields3zgog = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zgog uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zgog, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zgog := totalEncodedFields3zgog
	missingFieldsLeft3zgog := maxFields3zgog - totalEncodedFields3zgog

	var nextMiss3zgog int32 = -1
	var found3zgog [maxFields3zgog]bool
	var curField3zgog string

doneWithStruct3zgog:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zgog > 0 || missingFieldsLeft3zgog > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zgog, missingFieldsLeft3zgog, msgp.ShowFound(found3zgog[:]), unmarshalMsgFieldOrder3zgog)
		if encodedFieldsLeft3zgog > 0 {
			encodedFieldsLeft3zgog--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zgog = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zgog < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zgog = 0
			}
			for nextMiss3zgog < maxFields3zgog && (found3zgog[nextMiss3zgog] || unmarshalMsgFieldSkip3zgog[nextMiss3zgog]) {
				nextMiss3zgog++
			}
			if nextMiss3zgog == maxFields3zgog {
				// filled all the empty fields!
				break doneWithStruct3zgog
			}
			missingFieldsLeft3zgog--
			curField3zgog = unmarshalMsgFieldOrder3zgog[nextMiss3zgog]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zgog)
		switch curField3zgog {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found3zgog[0] = true
			{
				var zicz uint64
				zicz, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zicz)
			}
		case "Str":
			found3zgog[1] = true
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
	if nextMiss3zgog != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of KS
var unmarshalMsgFieldOrder3zgog = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip3zgog = []bool{false, false}

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
	const maxFields4zshc = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zshc uint32
	totalEncodedFields4zshc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zshc := totalEncodedFields4zshc
	missingFieldsLeft4zshc := maxFields4zshc - totalEncodedFields4zshc

	var nextMiss4zshc int32 = -1
	var found4zshc [maxFields4zshc]bool
	var curField4zshc string

doneWithStruct4zshc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zshc > 0 || missingFieldsLeft4zshc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zshc, missingFieldsLeft4zshc, msgp.ShowFound(found4zshc[:]), decodeMsgFieldOrder4zshc)
		if encodedFieldsLeft4zshc > 0 {
			encodedFieldsLeft4zshc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zshc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zshc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zshc = 0
			}
			for nextMiss4zshc < maxFields4zshc && (found4zshc[nextMiss4zshc] || decodeMsgFieldSkip4zshc[nextMiss4zshc]) {
				nextMiss4zshc++
			}
			if nextMiss4zshc == maxFields4zshc {
				// filled all the empty fields!
				break doneWithStruct4zshc
			}
			missingFieldsLeft4zshc--
			curField4zshc = decodeMsgFieldOrder4zshc[nextMiss4zshc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zshc)
		switch curField4zshc {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found4zshc[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zshc[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zshc[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zshc[3] = true
			var zdhc uint32
			zdhc, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zdhc) {
				z.Structs = (z.Structs)[:zdhc]
			} else {
				z.Structs = make([]Struct, zdhc)
			}
			for zoro := range z.Structs {
				const maxFields5zdkq = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zdkq uint32
				totalEncodedFields5zdkq, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zdkq := totalEncodedFields5zdkq
				missingFieldsLeft5zdkq := maxFields5zdkq - totalEncodedFields5zdkq

				var nextMiss5zdkq int32 = -1
				var found5zdkq [maxFields5zdkq]bool
				var curField5zdkq string

			doneWithStruct5zdkq:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zdkq > 0 || missingFieldsLeft5zdkq > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zdkq, missingFieldsLeft5zdkq, msgp.ShowFound(found5zdkq[:]), decodeMsgFieldOrder5zdkq)
					if encodedFieldsLeft5zdkq > 0 {
						encodedFieldsLeft5zdkq--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zdkq = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zdkq < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zdkq = 0
						}
						for nextMiss5zdkq < maxFields5zdkq && (found5zdkq[nextMiss5zdkq] || decodeMsgFieldSkip5zdkq[nextMiss5zdkq]) {
							nextMiss5zdkq++
						}
						if nextMiss5zdkq == maxFields5zdkq {
							// filled all the empty fields!
							break doneWithStruct5zdkq
						}
						missingFieldsLeft5zdkq--
						curField5zdkq = decodeMsgFieldOrder5zdkq[nextMiss5zdkq]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zdkq)
					switch curField5zdkq {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zdkq[0] = true
						z.Structs[zoro].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found5zdkq[1] = true
						var zckv uint32
						zckv, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zoro].Fields) >= int(zckv) {
							z.Structs[zoro].Fields = (z.Structs[zoro].Fields)[:zckv]
						} else {
							z.Structs[zoro].Fields = make([]Field, zckv)
						}
						for zbvt := range z.Structs[zoro].Fields {
							err = z.Structs[zoro].Fields[zbvt].DecodeMsg(dc)
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
				if nextMiss5zdkq != -1 {
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
	if nextMiss4zshc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder4zshc = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip4zshc = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder5zdkq = []string{"StructName", "Fields"}

var decodeMsgFieldSkip5zdkq = []bool{false, false}

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
	for zoro := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zoro].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zoro].Fields)))
		if err != nil {
			panic(err)
		}
		for zbvt := range z.Structs[zoro].Fields {
			err = z.Structs[zoro].Fields[zbvt].EncodeMsg(en)
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
	for zoro := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zoro].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zoro].Fields)))
		for zbvt := range z.Structs[zoro].Fields {
			o, err = z.Structs[zoro].Fields[zbvt].MarshalMsg(o)
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
	const maxFields6zdyr = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zdyr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zdyr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zdyr := totalEncodedFields6zdyr
	missingFieldsLeft6zdyr := maxFields6zdyr - totalEncodedFields6zdyr

	var nextMiss6zdyr int32 = -1
	var found6zdyr [maxFields6zdyr]bool
	var curField6zdyr string

doneWithStruct6zdyr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zdyr > 0 || missingFieldsLeft6zdyr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zdyr, missingFieldsLeft6zdyr, msgp.ShowFound(found6zdyr[:]), unmarshalMsgFieldOrder6zdyr)
		if encodedFieldsLeft6zdyr > 0 {
			encodedFieldsLeft6zdyr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zdyr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zdyr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zdyr = 0
			}
			for nextMiss6zdyr < maxFields6zdyr && (found6zdyr[nextMiss6zdyr] || unmarshalMsgFieldSkip6zdyr[nextMiss6zdyr]) {
				nextMiss6zdyr++
			}
			if nextMiss6zdyr == maxFields6zdyr {
				// filled all the empty fields!
				break doneWithStruct6zdyr
			}
			missingFieldsLeft6zdyr--
			curField6zdyr = unmarshalMsgFieldOrder6zdyr[nextMiss6zdyr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zdyr)
		switch curField6zdyr {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found6zdyr[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found6zdyr[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found6zdyr[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zdyr[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zsol uint32
				zsol, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zsol) {
					z.Structs = (z.Structs)[:zsol]
				} else {
					z.Structs = make([]Struct, zsol)
				}
				for zoro := range z.Structs {
					const maxFields7zcva = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zcva uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zcva, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zcva := totalEncodedFields7zcva
					missingFieldsLeft7zcva := maxFields7zcva - totalEncodedFields7zcva

					var nextMiss7zcva int32 = -1
					var found7zcva [maxFields7zcva]bool
					var curField7zcva string

				doneWithStruct7zcva:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zcva > 0 || missingFieldsLeft7zcva > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zcva, missingFieldsLeft7zcva, msgp.ShowFound(found7zcva[:]), unmarshalMsgFieldOrder7zcva)
						if encodedFieldsLeft7zcva > 0 {
							encodedFieldsLeft7zcva--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zcva = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zcva < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zcva = 0
							}
							for nextMiss7zcva < maxFields7zcva && (found7zcva[nextMiss7zcva] || unmarshalMsgFieldSkip7zcva[nextMiss7zcva]) {
								nextMiss7zcva++
							}
							if nextMiss7zcva == maxFields7zcva {
								// filled all the empty fields!
								break doneWithStruct7zcva
							}
							missingFieldsLeft7zcva--
							curField7zcva = unmarshalMsgFieldOrder7zcva[nextMiss7zcva]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zcva)
						switch curField7zcva {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zcva[0] = true
							z.Structs[zoro].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found7zcva[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zoro].Fields) = (z.Structs[zoro].Fields)[:0]
							} else {

								var zibw uint32
								zibw, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zoro].Fields) >= int(zibw) {
									z.Structs[zoro].Fields = (z.Structs[zoro].Fields)[:zibw]
								} else {
									z.Structs[zoro].Fields = make([]Field, zibw)
								}
								for zbvt := range z.Structs[zoro].Fields {
									bts, err = z.Structs[zoro].Fields[zbvt].UnmarshalMsg(bts)
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
					if nextMiss7zcva != -1 {
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
	if nextMiss6zdyr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder6zdyr = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip6zdyr = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder7zcva = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zcva = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zoro := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zoro].StructName) + 7 + msgp.ArrayHeaderSize
		for zbvt := range z.Structs[zoro].Fields {
			s += z.Structs[zoro].Fields[zbvt].Msgsize()
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
	const maxFields8zztd = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zztd uint32
	totalEncodedFields8zztd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zztd := totalEncodedFields8zztd
	missingFieldsLeft8zztd := maxFields8zztd - totalEncodedFields8zztd

	var nextMiss8zztd int32 = -1
	var found8zztd [maxFields8zztd]bool
	var curField8zztd string

doneWithStruct8zztd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zztd > 0 || missingFieldsLeft8zztd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zztd, missingFieldsLeft8zztd, msgp.ShowFound(found8zztd[:]), decodeMsgFieldOrder8zztd)
		if encodedFieldsLeft8zztd > 0 {
			encodedFieldsLeft8zztd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zztd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zztd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zztd = 0
			}
			for nextMiss8zztd < maxFields8zztd && (found8zztd[nextMiss8zztd] || decodeMsgFieldSkip8zztd[nextMiss8zztd]) {
				nextMiss8zztd++
			}
			if nextMiss8zztd == maxFields8zztd {
				// filled all the empty fields!
				break doneWithStruct8zztd
			}
			missingFieldsLeft8zztd--
			curField8zztd = decodeMsgFieldOrder8zztd[nextMiss8zztd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zztd)
		switch curField8zztd {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found8zztd[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found8zztd[1] = true
			var zlrx uint32
			zlrx, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zlrx) {
				z.Fields = (z.Fields)[:zlrx]
			} else {
				z.Fields = make([]Field, zlrx)
			}
			for zpaz := range z.Fields {
				err = z.Fields[zpaz].DecodeMsg(dc)
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
	if nextMiss8zztd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder8zztd = []string{"StructName", "Fields"}

var decodeMsgFieldSkip8zztd = []bool{false, false}

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
	for zpaz := range z.Fields {
		err = z.Fields[zpaz].EncodeMsg(en)
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
	for zpaz := range z.Fields {
		o, err = z.Fields[zpaz].MarshalMsg(o)
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
	const maxFields9zlpe = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zlpe uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zlpe, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zlpe := totalEncodedFields9zlpe
	missingFieldsLeft9zlpe := maxFields9zlpe - totalEncodedFields9zlpe

	var nextMiss9zlpe int32 = -1
	var found9zlpe [maxFields9zlpe]bool
	var curField9zlpe string

doneWithStruct9zlpe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zlpe > 0 || missingFieldsLeft9zlpe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zlpe, missingFieldsLeft9zlpe, msgp.ShowFound(found9zlpe[:]), unmarshalMsgFieldOrder9zlpe)
		if encodedFieldsLeft9zlpe > 0 {
			encodedFieldsLeft9zlpe--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zlpe = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zlpe < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zlpe = 0
			}
			for nextMiss9zlpe < maxFields9zlpe && (found9zlpe[nextMiss9zlpe] || unmarshalMsgFieldSkip9zlpe[nextMiss9zlpe]) {
				nextMiss9zlpe++
			}
			if nextMiss9zlpe == maxFields9zlpe {
				// filled all the empty fields!
				break doneWithStruct9zlpe
			}
			missingFieldsLeft9zlpe--
			curField9zlpe = unmarshalMsgFieldOrder9zlpe[nextMiss9zlpe]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zlpe)
		switch curField9zlpe {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found9zlpe[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found9zlpe[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zies uint32
				zies, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zies) {
					z.Fields = (z.Fields)[:zies]
				} else {
					z.Fields = make([]Field, zies)
				}
				for zpaz := range z.Fields {
					bts, err = z.Fields[zpaz].UnmarshalMsg(bts)
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
	if nextMiss9zlpe != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder9zlpe = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip9zlpe = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zpaz := range z.Fields {
		s += z.Fields[zpaz].Msgsize()
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
		var zrci uint64
		zrci, err = dc.ReadUint64()
		(*z) = Zkind(zrci)
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
		var zzvj uint64
		zzvj, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zzvj)
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
	const maxFields10zhyo = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zhyo uint32
	totalEncodedFields10zhyo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zhyo := totalEncodedFields10zhyo
	missingFieldsLeft10zhyo := maxFields10zhyo - totalEncodedFields10zhyo

	var nextMiss10zhyo int32 = -1
	var found10zhyo [maxFields10zhyo]bool
	var curField10zhyo string

doneWithStruct10zhyo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zhyo > 0 || missingFieldsLeft10zhyo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zhyo, missingFieldsLeft10zhyo, msgp.ShowFound(found10zhyo[:]), decodeMsgFieldOrder10zhyo)
		if encodedFieldsLeft10zhyo > 0 {
			encodedFieldsLeft10zhyo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zhyo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zhyo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zhyo = 0
			}
			for nextMiss10zhyo < maxFields10zhyo && (found10zhyo[nextMiss10zhyo] || decodeMsgFieldSkip10zhyo[nextMiss10zhyo]) {
				nextMiss10zhyo++
			}
			if nextMiss10zhyo == maxFields10zhyo {
				// filled all the empty fields!
				break doneWithStruct10zhyo
			}
			missingFieldsLeft10zhyo--
			curField10zhyo = decodeMsgFieldOrder10zhyo[nextMiss10zhyo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zhyo)
		switch curField10zhyo {
		// -- templateDecodeMsg ends here --

		case "Name":
			found10zhyo[0] = true
			const maxFields11zkzl = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields11zkzl uint32
			totalEncodedFields11zkzl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft11zkzl := totalEncodedFields11zkzl
			missingFieldsLeft11zkzl := maxFields11zkzl - totalEncodedFields11zkzl

			var nextMiss11zkzl int32 = -1
			var found11zkzl [maxFields11zkzl]bool
			var curField11zkzl string

		doneWithStruct11zkzl:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zkzl > 0 || missingFieldsLeft11zkzl > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zkzl, missingFieldsLeft11zkzl, msgp.ShowFound(found11zkzl[:]), decodeMsgFieldOrder11zkzl)
				if encodedFieldsLeft11zkzl > 0 {
					encodedFieldsLeft11zkzl--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField11zkzl = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zkzl < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss11zkzl = 0
					}
					for nextMiss11zkzl < maxFields11zkzl && (found11zkzl[nextMiss11zkzl] || decodeMsgFieldSkip11zkzl[nextMiss11zkzl]) {
						nextMiss11zkzl++
					}
					if nextMiss11zkzl == maxFields11zkzl {
						// filled all the empty fields!
						break doneWithStruct11zkzl
					}
					missingFieldsLeft11zkzl--
					curField11zkzl = decodeMsgFieldOrder11zkzl[nextMiss11zkzl]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zkzl)
				switch curField11zkzl {
				// -- templateDecodeMsg ends here --

				case "Kind":
					found11zkzl[0] = true
					{
						var zhkn uint64
						zhkn, err = dc.ReadUint64()
						z.Name.Kind = Zkind(zhkn)
					}
					if err != nil {
						panic(err)
					}
				case "Str":
					found11zkzl[1] = true
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
			if nextMiss11zkzl != -1 {
				dc.PopAlwaysNil()
			}

		case "Domain":
			found10zhyo[1] = true
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
			found10zhyo[2] = true
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
	if nextMiss10zhyo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder10zhyo = []string{"Name", "Domain", "Range"}

var decodeMsgFieldSkip10zhyo = []bool{false, false, false}

// fields of KS
var decodeMsgFieldOrder11zkzl = []string{"Kind", "Str"}

var decodeMsgFieldSkip11zkzl = []bool{false, false}

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
	var empty_zbyp [3]bool
	fieldsInUse_zibj := z.fieldsNotEmpty(empty_zbyp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zibj)
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
	if !empty_zbyp[1] {
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

	if !empty_zbyp[2] {
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
	const maxFields12zwih = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields12zwih uint32
	if !nbs.AlwaysNil {
		totalEncodedFields12zwih, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft12zwih := totalEncodedFields12zwih
	missingFieldsLeft12zwih := maxFields12zwih - totalEncodedFields12zwih

	var nextMiss12zwih int32 = -1
	var found12zwih [maxFields12zwih]bool
	var curField12zwih string

doneWithStruct12zwih:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zwih > 0 || missingFieldsLeft12zwih > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zwih, missingFieldsLeft12zwih, msgp.ShowFound(found12zwih[:]), unmarshalMsgFieldOrder12zwih)
		if encodedFieldsLeft12zwih > 0 {
			encodedFieldsLeft12zwih--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField12zwih = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zwih < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss12zwih = 0
			}
			for nextMiss12zwih < maxFields12zwih && (found12zwih[nextMiss12zwih] || unmarshalMsgFieldSkip12zwih[nextMiss12zwih]) {
				nextMiss12zwih++
			}
			if nextMiss12zwih == maxFields12zwih {
				// filled all the empty fields!
				break doneWithStruct12zwih
			}
			missingFieldsLeft12zwih--
			curField12zwih = unmarshalMsgFieldOrder12zwih[nextMiss12zwih]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zwih)
		switch curField12zwih {
		// -- templateUnmarshalMsg ends here --

		case "Name":
			found12zwih[0] = true
			const maxFields13zgmd = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields13zgmd uint32
			if !nbs.AlwaysNil {
				totalEncodedFields13zgmd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft13zgmd := totalEncodedFields13zgmd
			missingFieldsLeft13zgmd := maxFields13zgmd - totalEncodedFields13zgmd

			var nextMiss13zgmd int32 = -1
			var found13zgmd [maxFields13zgmd]bool
			var curField13zgmd string

		doneWithStruct13zgmd:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft13zgmd > 0 || missingFieldsLeft13zgmd > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zgmd, missingFieldsLeft13zgmd, msgp.ShowFound(found13zgmd[:]), unmarshalMsgFieldOrder13zgmd)
				if encodedFieldsLeft13zgmd > 0 {
					encodedFieldsLeft13zgmd--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField13zgmd = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss13zgmd < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss13zgmd = 0
					}
					for nextMiss13zgmd < maxFields13zgmd && (found13zgmd[nextMiss13zgmd] || unmarshalMsgFieldSkip13zgmd[nextMiss13zgmd]) {
						nextMiss13zgmd++
					}
					if nextMiss13zgmd == maxFields13zgmd {
						// filled all the empty fields!
						break doneWithStruct13zgmd
					}
					missingFieldsLeft13zgmd--
					curField13zgmd = unmarshalMsgFieldOrder13zgmd[nextMiss13zgmd]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField13zgmd)
				switch curField13zgmd {
				// -- templateUnmarshalMsg ends here --

				case "Kind":
					found13zgmd[0] = true
					{
						var zedx uint64
						zedx, bts, err = nbs.ReadUint64Bytes(bts)

						if err != nil {
							panic(err)
						}
						z.Name.Kind = Zkind(zedx)
					}
				case "Str":
					found13zgmd[1] = true
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
			if nextMiss13zgmd != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Domain":
			found12zwih[1] = true
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
			found12zwih[2] = true
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
	if nextMiss12zwih != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder12zwih = []string{"Name", "Domain", "Range"}

var unmarshalMsgFieldSkip12zwih = []bool{false, false, false}

// fields of KS
var unmarshalMsgFieldOrder13zgmd = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip13zgmd = []bool{false, false}

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
