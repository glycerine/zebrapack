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
	const maxFields0zmlq = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zmlq uint32
	totalEncodedFields0zmlq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zmlq := totalEncodedFields0zmlq
	missingFieldsLeft0zmlq := maxFields0zmlq - totalEncodedFields0zmlq

	var nextMiss0zmlq int32 = -1
	var found0zmlq [maxFields0zmlq]bool
	var curField0zmlq string

doneWithStruct0zmlq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zmlq > 0 || missingFieldsLeft0zmlq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zmlq, missingFieldsLeft0zmlq, msgp.ShowFound(found0zmlq[:]), decodeMsgFieldOrder0zmlq)
		if encodedFieldsLeft0zmlq > 0 {
			encodedFieldsLeft0zmlq--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zmlq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zmlq < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zmlq = 0
			}
			for nextMiss0zmlq < maxFields0zmlq && (found0zmlq[nextMiss0zmlq] || decodeMsgFieldSkip0zmlq[nextMiss0zmlq]) {
				nextMiss0zmlq++
			}
			if nextMiss0zmlq == maxFields0zmlq {
				// filled all the empty fields!
				break doneWithStruct0zmlq
			}
			missingFieldsLeft0zmlq--
			curField0zmlq = decodeMsgFieldOrder0zmlq[nextMiss0zmlq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zmlq)
		switch curField0zmlq {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zmlq[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zmlq[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zmlq[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zmlq[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zmlq[4] = true
			{
				var zeiz uint64
				zeiz, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zeiz)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zmlq[5] = true
			{
				var zvbn uint64
				zvbn, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zvbn)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zmlq[6] = true
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
			found0zmlq[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zmlq[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zmlq[9] = true
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
	if nextMiss0zmlq != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zmlq = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zmlq = []bool{false, false, false, false, false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Field) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 10
	}
	var fieldsInUse uint32 = 10
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
	var empty_zgcv [10]bool
	fieldsInUse_zzai := z.fieldsNotEmpty(empty_zgcv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zzai)
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
	// write "FieldTypeStr"
	err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldTypeStr)
	if err != nil {
		panic(err)
	}
	// write "FieldCategory"
	err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldCategory))
	if err != nil {
		panic(err)
	}
	// write "FieldPrimitive"
	err = en.Append(0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldPrimitive))
	if err != nil {
		panic(err)
	}
	if !empty_zgcv[6] {
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

	if !empty_zgcv[7] {
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

	if !empty_zgcv[8] {
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

	if !empty_zgcv[9] {
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
	// string "FieldTypeStr"
	o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldTypeStr)
	// string "FieldCategory"
	o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	o = msgp.AppendUint64(o, uint64(z.FieldCategory))
	// string "FieldPrimitive"
	o = append(o, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendUint64(o, uint64(z.FieldPrimitive))
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
	const maxFields1zvja = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zvja uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zvja, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zvja := totalEncodedFields1zvja
	missingFieldsLeft1zvja := maxFields1zvja - totalEncodedFields1zvja

	var nextMiss1zvja int32 = -1
	var found1zvja [maxFields1zvja]bool
	var curField1zvja string

doneWithStruct1zvja:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zvja > 0 || missingFieldsLeft1zvja > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zvja, missingFieldsLeft1zvja, msgp.ShowFound(found1zvja[:]), unmarshalMsgFieldOrder1zvja)
		if encodedFieldsLeft1zvja > 0 {
			encodedFieldsLeft1zvja--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zvja = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zvja < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zvja = 0
			}
			for nextMiss1zvja < maxFields1zvja && (found1zvja[nextMiss1zvja] || unmarshalMsgFieldSkip1zvja[nextMiss1zvja]) {
				nextMiss1zvja++
			}
			if nextMiss1zvja == maxFields1zvja {
				// filled all the empty fields!
				break doneWithStruct1zvja
			}
			missingFieldsLeft1zvja--
			curField1zvja = unmarshalMsgFieldOrder1zvja[nextMiss1zvja]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zvja)
		switch curField1zvja {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zvja[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zvja[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zvja[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zvja[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zvja[4] = true
			{
				var zezq uint64
				zezq, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zezq)
			}
		case "FieldPrimitive":
			found1zvja[5] = true
			{
				var zpwa uint64
				zpwa, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zpwa)
			}
		case "FieldFullType":
			found1zvja[6] = true
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
			found1zvja[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zvja[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zvja[9] = true
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
	if nextMiss1zvja != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zvja = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zvja = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2znmq = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2znmq uint32
	totalEncodedFields2znmq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2znmq := totalEncodedFields2znmq
	missingFieldsLeft2znmq := maxFields2znmq - totalEncodedFields2znmq

	var nextMiss2znmq int32 = -1
	var found2znmq [maxFields2znmq]bool
	var curField2znmq string

doneWithStruct2znmq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2znmq > 0 || missingFieldsLeft2znmq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2znmq, missingFieldsLeft2znmq, msgp.ShowFound(found2znmq[:]), decodeMsgFieldOrder2znmq)
		if encodedFieldsLeft2znmq > 0 {
			encodedFieldsLeft2znmq--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2znmq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2znmq < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2znmq = 0
			}
			for nextMiss2znmq < maxFields2znmq && (found2znmq[nextMiss2znmq] || decodeMsgFieldSkip2znmq[nextMiss2znmq]) {
				nextMiss2znmq++
			}
			if nextMiss2znmq == maxFields2znmq {
				// filled all the empty fields!
				break doneWithStruct2znmq
			}
			missingFieldsLeft2znmq--
			curField2znmq = decodeMsgFieldOrder2znmq[nextMiss2znmq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2znmq)
		switch curField2znmq {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found2znmq[0] = true
			{
				var znyu uint64
				znyu, err = dc.ReadUint64()
				z.Kind = Zkind(znyu)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found2znmq[1] = true
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
	if nextMiss2znmq != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of KS
var decodeMsgFieldOrder2znmq = []string{"Kind", "Str"}

var decodeMsgFieldSkip2znmq = []bool{false, false}

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
	const maxFields3zflm = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zflm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zflm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zflm := totalEncodedFields3zflm
	missingFieldsLeft3zflm := maxFields3zflm - totalEncodedFields3zflm

	var nextMiss3zflm int32 = -1
	var found3zflm [maxFields3zflm]bool
	var curField3zflm string

doneWithStruct3zflm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zflm > 0 || missingFieldsLeft3zflm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zflm, missingFieldsLeft3zflm, msgp.ShowFound(found3zflm[:]), unmarshalMsgFieldOrder3zflm)
		if encodedFieldsLeft3zflm > 0 {
			encodedFieldsLeft3zflm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zflm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zflm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zflm = 0
			}
			for nextMiss3zflm < maxFields3zflm && (found3zflm[nextMiss3zflm] || unmarshalMsgFieldSkip3zflm[nextMiss3zflm]) {
				nextMiss3zflm++
			}
			if nextMiss3zflm == maxFields3zflm {
				// filled all the empty fields!
				break doneWithStruct3zflm
			}
			missingFieldsLeft3zflm--
			curField3zflm = unmarshalMsgFieldOrder3zflm[nextMiss3zflm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zflm)
		switch curField3zflm {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found3zflm[0] = true
			{
				var zqzu uint64
				zqzu, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zqzu)
			}
		case "Str":
			found3zflm[1] = true
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
	if nextMiss3zflm != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of KS
var unmarshalMsgFieldOrder3zflm = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip3zflm = []bool{false, false}

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
	const maxFields4zrbx = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zrbx uint32
	totalEncodedFields4zrbx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zrbx := totalEncodedFields4zrbx
	missingFieldsLeft4zrbx := maxFields4zrbx - totalEncodedFields4zrbx

	var nextMiss4zrbx int32 = -1
	var found4zrbx [maxFields4zrbx]bool
	var curField4zrbx string

doneWithStruct4zrbx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zrbx > 0 || missingFieldsLeft4zrbx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zrbx, missingFieldsLeft4zrbx, msgp.ShowFound(found4zrbx[:]), decodeMsgFieldOrder4zrbx)
		if encodedFieldsLeft4zrbx > 0 {
			encodedFieldsLeft4zrbx--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zrbx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zrbx < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zrbx = 0
			}
			for nextMiss4zrbx < maxFields4zrbx && (found4zrbx[nextMiss4zrbx] || decodeMsgFieldSkip4zrbx[nextMiss4zrbx]) {
				nextMiss4zrbx++
			}
			if nextMiss4zrbx == maxFields4zrbx {
				// filled all the empty fields!
				break doneWithStruct4zrbx
			}
			missingFieldsLeft4zrbx--
			curField4zrbx = decodeMsgFieldOrder4zrbx[nextMiss4zrbx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zrbx)
		switch curField4zrbx {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found4zrbx[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zrbx[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zrbx[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zrbx[3] = true
			var zkmv uint32
			zkmv, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zkmv) {
				z.Structs = (z.Structs)[:zkmv]
			} else {
				z.Structs = make([]Struct, zkmv)
			}
			for zuoh := range z.Structs {
				const maxFields5zuxt = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zuxt uint32
				totalEncodedFields5zuxt, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zuxt := totalEncodedFields5zuxt
				missingFieldsLeft5zuxt := maxFields5zuxt - totalEncodedFields5zuxt

				var nextMiss5zuxt int32 = -1
				var found5zuxt [maxFields5zuxt]bool
				var curField5zuxt string

			doneWithStruct5zuxt:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zuxt > 0 || missingFieldsLeft5zuxt > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zuxt, missingFieldsLeft5zuxt, msgp.ShowFound(found5zuxt[:]), decodeMsgFieldOrder5zuxt)
					if encodedFieldsLeft5zuxt > 0 {
						encodedFieldsLeft5zuxt--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zuxt = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zuxt < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zuxt = 0
						}
						for nextMiss5zuxt < maxFields5zuxt && (found5zuxt[nextMiss5zuxt] || decodeMsgFieldSkip5zuxt[nextMiss5zuxt]) {
							nextMiss5zuxt++
						}
						if nextMiss5zuxt == maxFields5zuxt {
							// filled all the empty fields!
							break doneWithStruct5zuxt
						}
						missingFieldsLeft5zuxt--
						curField5zuxt = decodeMsgFieldOrder5zuxt[nextMiss5zuxt]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zuxt)
					switch curField5zuxt {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zuxt[0] = true
						z.Structs[zuoh].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found5zuxt[1] = true
						var zzrs uint32
						zzrs, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zuoh].Fields) >= int(zzrs) {
							z.Structs[zuoh].Fields = (z.Structs[zuoh].Fields)[:zzrs]
						} else {
							z.Structs[zuoh].Fields = make([]Field, zzrs)
						}
						for zyff := range z.Structs[zuoh].Fields {
							err = z.Structs[zuoh].Fields[zyff].DecodeMsg(dc)
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
				if nextMiss5zuxt != -1 {
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
	if nextMiss4zrbx != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder4zrbx = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var decodeMsgFieldSkip4zrbx = []bool{false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder5zuxt = []string{"StructName", "Fields"}

var decodeMsgFieldSkip5zuxt = []bool{false, false}

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
	for zuoh := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zuoh].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zuoh].Fields)))
		if err != nil {
			panic(err)
		}
		for zyff := range z.Structs[zuoh].Fields {
			err = z.Structs[zuoh].Fields[zyff].EncodeMsg(en)
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
	for zuoh := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zuoh].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zuoh].Fields)))
		for zyff := range z.Structs[zuoh].Fields {
			o, err = z.Structs[zuoh].Fields[zyff].MarshalMsg(o)
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
	const maxFields6zevt = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zevt uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zevt, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zevt := totalEncodedFields6zevt
	missingFieldsLeft6zevt := maxFields6zevt - totalEncodedFields6zevt

	var nextMiss6zevt int32 = -1
	var found6zevt [maxFields6zevt]bool
	var curField6zevt string

doneWithStruct6zevt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zevt > 0 || missingFieldsLeft6zevt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zevt, missingFieldsLeft6zevt, msgp.ShowFound(found6zevt[:]), unmarshalMsgFieldOrder6zevt)
		if encodedFieldsLeft6zevt > 0 {
			encodedFieldsLeft6zevt--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zevt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zevt < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zevt = 0
			}
			for nextMiss6zevt < maxFields6zevt && (found6zevt[nextMiss6zevt] || unmarshalMsgFieldSkip6zevt[nextMiss6zevt]) {
				nextMiss6zevt++
			}
			if nextMiss6zevt == maxFields6zevt {
				// filled all the empty fields!
				break doneWithStruct6zevt
			}
			missingFieldsLeft6zevt--
			curField6zevt = unmarshalMsgFieldOrder6zevt[nextMiss6zevt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zevt)
		switch curField6zevt {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found6zevt[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found6zevt[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found6zevt[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zevt[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zuhi uint32
				zuhi, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zuhi) {
					z.Structs = (z.Structs)[:zuhi]
				} else {
					z.Structs = make([]Struct, zuhi)
				}
				for zuoh := range z.Structs {
					const maxFields7zndl = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zndl uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zndl, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zndl := totalEncodedFields7zndl
					missingFieldsLeft7zndl := maxFields7zndl - totalEncodedFields7zndl

					var nextMiss7zndl int32 = -1
					var found7zndl [maxFields7zndl]bool
					var curField7zndl string

				doneWithStruct7zndl:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zndl > 0 || missingFieldsLeft7zndl > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zndl, missingFieldsLeft7zndl, msgp.ShowFound(found7zndl[:]), unmarshalMsgFieldOrder7zndl)
						if encodedFieldsLeft7zndl > 0 {
							encodedFieldsLeft7zndl--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zndl = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zndl < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zndl = 0
							}
							for nextMiss7zndl < maxFields7zndl && (found7zndl[nextMiss7zndl] || unmarshalMsgFieldSkip7zndl[nextMiss7zndl]) {
								nextMiss7zndl++
							}
							if nextMiss7zndl == maxFields7zndl {
								// filled all the empty fields!
								break doneWithStruct7zndl
							}
							missingFieldsLeft7zndl--
							curField7zndl = unmarshalMsgFieldOrder7zndl[nextMiss7zndl]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zndl)
						switch curField7zndl {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zndl[0] = true
							z.Structs[zuoh].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found7zndl[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zuoh].Fields) = (z.Structs[zuoh].Fields)[:0]
							} else {

								var zoog uint32
								zoog, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zuoh].Fields) >= int(zoog) {
									z.Structs[zuoh].Fields = (z.Structs[zuoh].Fields)[:zoog]
								} else {
									z.Structs[zuoh].Fields = make([]Field, zoog)
								}
								for zyff := range z.Structs[zuoh].Fields {
									bts, err = z.Structs[zuoh].Fields[zyff].UnmarshalMsg(bts)
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
					if nextMiss7zndl != -1 {
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
	if nextMiss6zevt != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder6zevt = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs"}

var unmarshalMsgFieldSkip6zevt = []bool{false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder7zndl = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zndl = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zuoh := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zuoh].StructName) + 7 + msgp.ArrayHeaderSize
		for zyff := range z.Structs[zuoh].Fields {
			s += z.Structs[zuoh].Fields[zyff].Msgsize()
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
	const maxFields8zbji = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zbji uint32
	totalEncodedFields8zbji, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zbji := totalEncodedFields8zbji
	missingFieldsLeft8zbji := maxFields8zbji - totalEncodedFields8zbji

	var nextMiss8zbji int32 = -1
	var found8zbji [maxFields8zbji]bool
	var curField8zbji string

doneWithStruct8zbji:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zbji > 0 || missingFieldsLeft8zbji > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zbji, missingFieldsLeft8zbji, msgp.ShowFound(found8zbji[:]), decodeMsgFieldOrder8zbji)
		if encodedFieldsLeft8zbji > 0 {
			encodedFieldsLeft8zbji--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zbji = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zbji < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zbji = 0
			}
			for nextMiss8zbji < maxFields8zbji && (found8zbji[nextMiss8zbji] || decodeMsgFieldSkip8zbji[nextMiss8zbji]) {
				nextMiss8zbji++
			}
			if nextMiss8zbji == maxFields8zbji {
				// filled all the empty fields!
				break doneWithStruct8zbji
			}
			missingFieldsLeft8zbji--
			curField8zbji = decodeMsgFieldOrder8zbji[nextMiss8zbji]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zbji)
		switch curField8zbji {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found8zbji[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found8zbji[1] = true
			var zsim uint32
			zsim, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zsim) {
				z.Fields = (z.Fields)[:zsim]
			} else {
				z.Fields = make([]Field, zsim)
			}
			for zqkl := range z.Fields {
				err = z.Fields[zqkl].DecodeMsg(dc)
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
	if nextMiss8zbji != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder8zbji = []string{"StructName", "Fields"}

var decodeMsgFieldSkip8zbji = []bool{false, false}

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
	for zqkl := range z.Fields {
		err = z.Fields[zqkl].EncodeMsg(en)
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
	for zqkl := range z.Fields {
		o, err = z.Fields[zqkl].MarshalMsg(o)
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
	const maxFields9zrsa = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zrsa uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zrsa, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zrsa := totalEncodedFields9zrsa
	missingFieldsLeft9zrsa := maxFields9zrsa - totalEncodedFields9zrsa

	var nextMiss9zrsa int32 = -1
	var found9zrsa [maxFields9zrsa]bool
	var curField9zrsa string

doneWithStruct9zrsa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zrsa > 0 || missingFieldsLeft9zrsa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zrsa, missingFieldsLeft9zrsa, msgp.ShowFound(found9zrsa[:]), unmarshalMsgFieldOrder9zrsa)
		if encodedFieldsLeft9zrsa > 0 {
			encodedFieldsLeft9zrsa--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zrsa = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zrsa < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zrsa = 0
			}
			for nextMiss9zrsa < maxFields9zrsa && (found9zrsa[nextMiss9zrsa] || unmarshalMsgFieldSkip9zrsa[nextMiss9zrsa]) {
				nextMiss9zrsa++
			}
			if nextMiss9zrsa == maxFields9zrsa {
				// filled all the empty fields!
				break doneWithStruct9zrsa
			}
			missingFieldsLeft9zrsa--
			curField9zrsa = unmarshalMsgFieldOrder9zrsa[nextMiss9zrsa]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zrsa)
		switch curField9zrsa {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found9zrsa[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found9zrsa[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zcvp uint32
				zcvp, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zcvp) {
					z.Fields = (z.Fields)[:zcvp]
				} else {
					z.Fields = make([]Field, zcvp)
				}
				for zqkl := range z.Fields {
					bts, err = z.Fields[zqkl].UnmarshalMsg(bts)
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
	if nextMiss9zrsa != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder9zrsa = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip9zrsa = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zqkl := range z.Fields {
		s += z.Fields[zqkl].Msgsize()
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
		var zazf uint64
		zazf, err = dc.ReadUint64()
		(*z) = Zkind(zazf)
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
		var zgtw uint64
		zgtw, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zgtw)
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
	const maxFields10zgre = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zgre uint32
	totalEncodedFields10zgre, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zgre := totalEncodedFields10zgre
	missingFieldsLeft10zgre := maxFields10zgre - totalEncodedFields10zgre

	var nextMiss10zgre int32 = -1
	var found10zgre [maxFields10zgre]bool
	var curField10zgre string

doneWithStruct10zgre:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zgre > 0 || missingFieldsLeft10zgre > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zgre, missingFieldsLeft10zgre, msgp.ShowFound(found10zgre[:]), decodeMsgFieldOrder10zgre)
		if encodedFieldsLeft10zgre > 0 {
			encodedFieldsLeft10zgre--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zgre = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zgre < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zgre = 0
			}
			for nextMiss10zgre < maxFields10zgre && (found10zgre[nextMiss10zgre] || decodeMsgFieldSkip10zgre[nextMiss10zgre]) {
				nextMiss10zgre++
			}
			if nextMiss10zgre == maxFields10zgre {
				// filled all the empty fields!
				break doneWithStruct10zgre
			}
			missingFieldsLeft10zgre--
			curField10zgre = decodeMsgFieldOrder10zgre[nextMiss10zgre]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zgre)
		switch curField10zgre {
		// -- templateDecodeMsg ends here --

		case "Name":
			found10zgre[0] = true
			const maxFields11zlos = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields11zlos uint32
			totalEncodedFields11zlos, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft11zlos := totalEncodedFields11zlos
			missingFieldsLeft11zlos := maxFields11zlos - totalEncodedFields11zlos

			var nextMiss11zlos int32 = -1
			var found11zlos [maxFields11zlos]bool
			var curField11zlos string

		doneWithStruct11zlos:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zlos > 0 || missingFieldsLeft11zlos > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zlos, missingFieldsLeft11zlos, msgp.ShowFound(found11zlos[:]), decodeMsgFieldOrder11zlos)
				if encodedFieldsLeft11zlos > 0 {
					encodedFieldsLeft11zlos--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField11zlos = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zlos < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss11zlos = 0
					}
					for nextMiss11zlos < maxFields11zlos && (found11zlos[nextMiss11zlos] || decodeMsgFieldSkip11zlos[nextMiss11zlos]) {
						nextMiss11zlos++
					}
					if nextMiss11zlos == maxFields11zlos {
						// filled all the empty fields!
						break doneWithStruct11zlos
					}
					missingFieldsLeft11zlos--
					curField11zlos = decodeMsgFieldOrder11zlos[nextMiss11zlos]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zlos)
				switch curField11zlos {
				// -- templateDecodeMsg ends here --

				case "Kind":
					found11zlos[0] = true
					{
						var zjjj uint64
						zjjj, err = dc.ReadUint64()
						z.Name.Kind = Zkind(zjjj)
					}
					if err != nil {
						panic(err)
					}
				case "Str":
					found11zlos[1] = true
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
			if nextMiss11zlos != -1 {
				dc.PopAlwaysNil()
			}

		case "Domain":
			found10zgre[1] = true
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
			found10zgre[2] = true
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
	if nextMiss10zgre != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder10zgre = []string{"Name", "Domain", "Range"}

var decodeMsgFieldSkip10zgre = []bool{false, false, false}

// fields of KS
var decodeMsgFieldOrder11zlos = []string{"Kind", "Str"}

var decodeMsgFieldSkip11zlos = []bool{false, false}

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
	var empty_zjkx [3]bool
	fieldsInUse_zqhi := z.fieldsNotEmpty(empty_zjkx[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqhi)
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
	if !empty_zjkx[1] {
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

	if !empty_zjkx[2] {
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
	const maxFields12zrwo = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields12zrwo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields12zrwo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft12zrwo := totalEncodedFields12zrwo
	missingFieldsLeft12zrwo := maxFields12zrwo - totalEncodedFields12zrwo

	var nextMiss12zrwo int32 = -1
	var found12zrwo [maxFields12zrwo]bool
	var curField12zrwo string

doneWithStruct12zrwo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zrwo > 0 || missingFieldsLeft12zrwo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zrwo, missingFieldsLeft12zrwo, msgp.ShowFound(found12zrwo[:]), unmarshalMsgFieldOrder12zrwo)
		if encodedFieldsLeft12zrwo > 0 {
			encodedFieldsLeft12zrwo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField12zrwo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zrwo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss12zrwo = 0
			}
			for nextMiss12zrwo < maxFields12zrwo && (found12zrwo[nextMiss12zrwo] || unmarshalMsgFieldSkip12zrwo[nextMiss12zrwo]) {
				nextMiss12zrwo++
			}
			if nextMiss12zrwo == maxFields12zrwo {
				// filled all the empty fields!
				break doneWithStruct12zrwo
			}
			missingFieldsLeft12zrwo--
			curField12zrwo = unmarshalMsgFieldOrder12zrwo[nextMiss12zrwo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zrwo)
		switch curField12zrwo {
		// -- templateUnmarshalMsg ends here --

		case "Name":
			found12zrwo[0] = true
			const maxFields13znut = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields13znut uint32
			if !nbs.AlwaysNil {
				totalEncodedFields13znut, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft13znut := totalEncodedFields13znut
			missingFieldsLeft13znut := maxFields13znut - totalEncodedFields13znut

			var nextMiss13znut int32 = -1
			var found13znut [maxFields13znut]bool
			var curField13znut string

		doneWithStruct13znut:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft13znut > 0 || missingFieldsLeft13znut > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13znut, missingFieldsLeft13znut, msgp.ShowFound(found13znut[:]), unmarshalMsgFieldOrder13znut)
				if encodedFieldsLeft13znut > 0 {
					encodedFieldsLeft13znut--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField13znut = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss13znut < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss13znut = 0
					}
					for nextMiss13znut < maxFields13znut && (found13znut[nextMiss13znut] || unmarshalMsgFieldSkip13znut[nextMiss13znut]) {
						nextMiss13znut++
					}
					if nextMiss13znut == maxFields13znut {
						// filled all the empty fields!
						break doneWithStruct13znut
					}
					missingFieldsLeft13znut--
					curField13znut = unmarshalMsgFieldOrder13znut[nextMiss13znut]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField13znut)
				switch curField13znut {
				// -- templateUnmarshalMsg ends here --

				case "Kind":
					found13znut[0] = true
					{
						var znww uint64
						znww, bts, err = nbs.ReadUint64Bytes(bts)

						if err != nil {
							panic(err)
						}
						z.Name.Kind = Zkind(znww)
					}
				case "Str":
					found13znut[1] = true
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
			if nextMiss13znut != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Domain":
			found12zrwo[1] = true
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
			found12zrwo[2] = true
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
	if nextMiss12zrwo != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder12zrwo = []string{"Name", "Domain", "Range"}

var unmarshalMsgFieldSkip12zrwo = []bool{false, false, false}

// fields of KS
var unmarshalMsgFieldOrder13znut = []string{"Kind", "Str"}

var unmarshalMsgFieldSkip13znut = []bool{false, false}

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
