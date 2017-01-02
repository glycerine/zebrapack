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
	const maxFields0ziwc = 11

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0ziwc uint32
	totalEncodedFields0ziwc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0ziwc := totalEncodedFields0ziwc
	missingFieldsLeft0ziwc := maxFields0ziwc - totalEncodedFields0ziwc

	var nextMiss0ziwc int32 = -1
	var found0ziwc [maxFields0ziwc]bool
	var curField0ziwc string

doneWithStruct0ziwc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ziwc > 0 || missingFieldsLeft0ziwc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ziwc, missingFieldsLeft0ziwc, msgp.ShowFound(found0ziwc[:]), decodeMsgFieldOrder0ziwc)
		if encodedFieldsLeft0ziwc > 0 {
			encodedFieldsLeft0ziwc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0ziwc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ziwc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0ziwc = 0
			}
			for nextMiss0ziwc < maxFields0ziwc && (found0ziwc[nextMiss0ziwc] || decodeMsgFieldSkip0ziwc[nextMiss0ziwc]) {
				nextMiss0ziwc++
			}
			if nextMiss0ziwc == maxFields0ziwc {
				// filled all the empty fields!
				break doneWithStruct0ziwc
			}
			missingFieldsLeft0ziwc--
			curField0ziwc = decodeMsgFieldOrder0ziwc[nextMiss0ziwc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ziwc)
		switch curField0ziwc {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0ziwc[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0ziwc[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0ziwc[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0ziwc[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0ziwc[4] = true
			{
				var zjzc uint64
				zjzc, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zjzc)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0ziwc[5] = true
			{
				var zwgl uint64
				zwgl, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zwgl)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0ziwc[6] = true
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
			found0ziwc[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0ziwc[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0ziwc[9] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "ShowZero":
			found0ziwc[10] = true
			z.ShowZero, err = dc.ReadBool()
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
	if nextMiss0ziwc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0ziwc = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated", "ShowZero"}

var decodeMsgFieldSkip0ziwc = []bool{false, false, false, false, false, false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Field) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 11
	}
	var fieldsInUse uint32 = 11
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
	isempty[10] = (!z.ShowZero) // bool, omitempty
	if isempty[10] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Field) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zutk [11]bool
	fieldsInUse_zmml := z.fieldsNotEmpty(empty_zutk[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zmml)
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
	if !empty_zutk[3] {
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

	if !empty_zutk[4] {
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

	if !empty_zutk[5] {
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

	if !empty_zutk[6] {
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

	if !empty_zutk[7] {
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

	if !empty_zutk[8] {
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

	if !empty_zutk[9] {
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

	if !empty_zutk[10] {
		// write "ShowZero"
		err = en.Append(0xa8, 0x53, 0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.ShowZero)
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
	var empty [11]bool
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

	if !empty[10] {
		// string "ShowZero"
		o = append(o, 0xa8, 0x53, 0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f)
		o = msgp.AppendBool(o, z.ShowZero)
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
	const maxFields1ztez = 11

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1ztez uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1ztez, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1ztez := totalEncodedFields1ztez
	missingFieldsLeft1ztez := maxFields1ztez - totalEncodedFields1ztez

	var nextMiss1ztez int32 = -1
	var found1ztez [maxFields1ztez]bool
	var curField1ztez string

doneWithStruct1ztez:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1ztez > 0 || missingFieldsLeft1ztez > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1ztez, missingFieldsLeft1ztez, msgp.ShowFound(found1ztez[:]), unmarshalMsgFieldOrder1ztez)
		if encodedFieldsLeft1ztez > 0 {
			encodedFieldsLeft1ztez--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1ztez = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1ztez < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1ztez = 0
			}
			for nextMiss1ztez < maxFields1ztez && (found1ztez[nextMiss1ztez] || unmarshalMsgFieldSkip1ztez[nextMiss1ztez]) {
				nextMiss1ztez++
			}
			if nextMiss1ztez == maxFields1ztez {
				// filled all the empty fields!
				break doneWithStruct1ztez
			}
			missingFieldsLeft1ztez--
			curField1ztez = unmarshalMsgFieldOrder1ztez[nextMiss1ztez]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1ztez)
		switch curField1ztez {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1ztez[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1ztez[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1ztez[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1ztez[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1ztez[4] = true
			{
				var zvhx uint64
				zvhx, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zvhx)
			}
		case "FieldPrimitive":
			found1ztez[5] = true
			{
				var zcuj uint64
				zcuj, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zcuj)
			}
		case "FieldFullType":
			found1ztez[6] = true
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
			found1ztez[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1ztez[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1ztez[9] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ShowZero":
			found1ztez[10] = true
			z.ShowZero, bts, err = nbs.ReadBoolBytes(bts)

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
	if nextMiss1ztez != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1ztez = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated", "ShowZero"}

var unmarshalMsgFieldSkip1ztez = []bool{false, false, false, false, false, false, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Field) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.FieldGoName) + 13 + msgp.StringPrefixSize + len(z.FieldTagName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 14 + msgp.Uint64Size + 15 + msgp.Uint64Size + 14
	if z.FieldFullType == nil {
		s += msgp.NilSize
	} else {
		s += z.FieldFullType.Msgsize()
	}
	s += 10 + msgp.BoolSize + 5 + msgp.BoolSize + 11 + msgp.BoolSize + 9 + msgp.BoolSize
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
	const maxFields2zzwv = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zzwv uint32
	totalEncodedFields2zzwv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zzwv := totalEncodedFields2zzwv
	missingFieldsLeft2zzwv := maxFields2zzwv - totalEncodedFields2zzwv

	var nextMiss2zzwv int32 = -1
	var found2zzwv [maxFields2zzwv]bool
	var curField2zzwv string

doneWithStruct2zzwv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zzwv > 0 || missingFieldsLeft2zzwv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zzwv, missingFieldsLeft2zzwv, msgp.ShowFound(found2zzwv[:]), decodeMsgFieldOrder2zzwv)
		if encodedFieldsLeft2zzwv > 0 {
			encodedFieldsLeft2zzwv--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zzwv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zzwv < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zzwv = 0
			}
			for nextMiss2zzwv < maxFields2zzwv && (found2zzwv[nextMiss2zzwv] || decodeMsgFieldSkip2zzwv[nextMiss2zzwv]) {
				nextMiss2zzwv++
			}
			if nextMiss2zzwv == maxFields2zzwv {
				// filled all the empty fields!
				break doneWithStruct2zzwv
			}
			missingFieldsLeft2zzwv--
			curField2zzwv = decodeMsgFieldOrder2zzwv[nextMiss2zzwv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zzwv)
		switch curField2zzwv {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zzwv[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zzwv[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zzwv[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zzwv[3] = true
			var zvpp uint32
			zvpp, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zvpp > 0 {
				z.Structs = make(map[string]*Struct, zvpp)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zvpp > 0 {
				zvpp--
				var zuhj string
				var zdju *Struct
				zuhj, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					zdju = nil
				} else {
					if zdju == nil {
						zdju = new(Struct)
					}
					const maxFields3zxlx = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zxlx uint32
					totalEncodedFields3zxlx, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zxlx := totalEncodedFields3zxlx
					missingFieldsLeft3zxlx := maxFields3zxlx - totalEncodedFields3zxlx

					var nextMiss3zxlx int32 = -1
					var found3zxlx [maxFields3zxlx]bool
					var curField3zxlx string

				doneWithStruct3zxlx:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zxlx > 0 || missingFieldsLeft3zxlx > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zxlx, missingFieldsLeft3zxlx, msgp.ShowFound(found3zxlx[:]), decodeMsgFieldOrder3zxlx)
						if encodedFieldsLeft3zxlx > 0 {
							encodedFieldsLeft3zxlx--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zxlx = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zxlx < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zxlx = 0
							}
							for nextMiss3zxlx < maxFields3zxlx && (found3zxlx[nextMiss3zxlx] || decodeMsgFieldSkip3zxlx[nextMiss3zxlx]) {
								nextMiss3zxlx++
							}
							if nextMiss3zxlx == maxFields3zxlx {
								// filled all the empty fields!
								break doneWithStruct3zxlx
							}
							missingFieldsLeft3zxlx--
							curField3zxlx = decodeMsgFieldOrder3zxlx[nextMiss3zxlx]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zxlx)
						switch curField3zxlx {
						// -- templateDecodeMsg ends here --

						case "StructName":
							found3zxlx[0] = true
							zdju.StructName, err = dc.ReadString()
							if err != nil {
								panic(err)
							}
						case "Fields":
							found3zxlx[1] = true
							var zeto uint32
							zeto, err = dc.ReadArrayHeader()
							if err != nil {
								panic(err)
							}
							if cap(zdju.Fields) >= int(zeto) {
								zdju.Fields = (zdju.Fields)[:zeto]
							} else {
								zdju.Fields = make([]Field, zeto)
							}
							for zxmj := range zdju.Fields {
								err = zdju.Fields[zxmj].DecodeMsg(dc)
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
					if nextMiss3zxlx != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[zuhj] = zdju
			}
		case "Imports":
			found2zzwv[4] = true
			var zuxm uint32
			zuxm, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zuxm) {
				z.Imports = (z.Imports)[:zuxm]
			} else {
				z.Imports = make([]string, zuxm)
			}
			for zgxg := range z.Imports {
				z.Imports[zgxg], err = dc.ReadString()
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
	if nextMiss2zzwv != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zzwv = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zzwv = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zxlx = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zxlx = []bool{false, false}

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
	for zuhj, zdju := range z.Structs {
		err = en.WriteString(zuhj)
		if err != nil {
			panic(err)
		}
		if zdju == nil {
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
			err = en.WriteString(zdju.StructName)
			if err != nil {
				panic(err)
			}
			// write "Fields"
			err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(zdju.Fields)))
			if err != nil {
				panic(err)
			}
			for zxmj := range zdju.Fields {
				err = zdju.Fields[zxmj].EncodeMsg(en)
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
	for zgxg := range z.Imports {
		err = en.WriteString(z.Imports[zgxg])
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
	for zuhj, zdju := range z.Structs {
		o = msgp.AppendString(o, zuhj)
		if zdju == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName"
			o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, zdju.StructName)
			// string "Fields"
			o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
			o = msgp.AppendArrayHeader(o, uint32(len(zdju.Fields)))
			for zxmj := range zdju.Fields {
				o, err = zdju.Fields[zxmj].MarshalMsg(o)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zgxg := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zgxg])
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
	const maxFields4zwdd = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zwdd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zwdd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zwdd := totalEncodedFields4zwdd
	missingFieldsLeft4zwdd := maxFields4zwdd - totalEncodedFields4zwdd

	var nextMiss4zwdd int32 = -1
	var found4zwdd [maxFields4zwdd]bool
	var curField4zwdd string

doneWithStruct4zwdd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zwdd > 0 || missingFieldsLeft4zwdd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zwdd, missingFieldsLeft4zwdd, msgp.ShowFound(found4zwdd[:]), unmarshalMsgFieldOrder4zwdd)
		if encodedFieldsLeft4zwdd > 0 {
			encodedFieldsLeft4zwdd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zwdd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zwdd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zwdd = 0
			}
			for nextMiss4zwdd < maxFields4zwdd && (found4zwdd[nextMiss4zwdd] || unmarshalMsgFieldSkip4zwdd[nextMiss4zwdd]) {
				nextMiss4zwdd++
			}
			if nextMiss4zwdd == maxFields4zwdd {
				// filled all the empty fields!
				break doneWithStruct4zwdd
			}
			missingFieldsLeft4zwdd--
			curField4zwdd = unmarshalMsgFieldOrder4zwdd[nextMiss4zwdd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zwdd)
		switch curField4zwdd {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zwdd[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zwdd[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zwdd[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zwdd[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zugk uint32
				zugk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zugk > 0 {
					z.Structs = make(map[string]*Struct, zugk)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zugk > 0 {
					var zuhj string
					var zdju *Struct
					zugk--
					zuhj, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && zdju == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if zdju == nil {
							zdju = new(Struct)
						}
						const maxFields5zxrc = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zxrc uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zxrc, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								panic(err)
								return
							}
						}
						encodedFieldsLeft5zxrc := totalEncodedFields5zxrc
						missingFieldsLeft5zxrc := maxFields5zxrc - totalEncodedFields5zxrc

						var nextMiss5zxrc int32 = -1
						var found5zxrc [maxFields5zxrc]bool
						var curField5zxrc string

					doneWithStruct5zxrc:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zxrc > 0 || missingFieldsLeft5zxrc > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxrc, missingFieldsLeft5zxrc, msgp.ShowFound(found5zxrc[:]), unmarshalMsgFieldOrder5zxrc)
							if encodedFieldsLeft5zxrc > 0 {
								encodedFieldsLeft5zxrc--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									panic(err)
									return
								}
								curField5zxrc = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zxrc < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zxrc = 0
								}
								for nextMiss5zxrc < maxFields5zxrc && (found5zxrc[nextMiss5zxrc] || unmarshalMsgFieldSkip5zxrc[nextMiss5zxrc]) {
									nextMiss5zxrc++
								}
								if nextMiss5zxrc == maxFields5zxrc {
									// filled all the empty fields!
									break doneWithStruct5zxrc
								}
								missingFieldsLeft5zxrc--
								curField5zxrc = unmarshalMsgFieldOrder5zxrc[nextMiss5zxrc]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zxrc)
							switch curField5zxrc {
							// -- templateUnmarshalMsg ends here --

							case "StructName":
								found5zxrc[0] = true
								zdju.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									panic(err)
								}
							case "Fields":
								found5zxrc[1] = true
								if nbs.AlwaysNil {
									(zdju.Fields) = (zdju.Fields)[:0]
								} else {

									var zpof uint32
									zpof, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										panic(err)
									}
									if cap(zdju.Fields) >= int(zpof) {
										zdju.Fields = (zdju.Fields)[:zpof]
									} else {
										zdju.Fields = make([]Field, zpof)
									}
									for zxmj := range zdju.Fields {
										bts, err = zdju.Fields[zxmj].UnmarshalMsg(bts)
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
						if nextMiss5zxrc != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[zuhj] = zdju
				}
			}
		case "Imports":
			found4zwdd[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zihd uint32
				zihd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zihd) {
					z.Imports = (z.Imports)[:zihd]
				} else {
					z.Imports = make([]string, zihd)
				}
				for zgxg := range z.Imports {
					z.Imports[zgxg], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zwdd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zwdd = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zwdd = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zxrc = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zxrc = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zuhj, zdju := range z.Structs {
			_ = zdju
			_ = zuhj
			s += msgp.StringPrefixSize + len(zuhj)
			if zdju == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 11 + msgp.StringPrefixSize + len(zdju.StructName) + 7 + msgp.ArrayHeaderSize
				for zxmj := range zdju.Fields {
					s += zdju.Fields[zxmj].Msgsize()
				}
			}
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zgxg := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zgxg])
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
	const maxFields6zepu = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zepu uint32
	totalEncodedFields6zepu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zepu := totalEncodedFields6zepu
	missingFieldsLeft6zepu := maxFields6zepu - totalEncodedFields6zepu

	var nextMiss6zepu int32 = -1
	var found6zepu [maxFields6zepu]bool
	var curField6zepu string

doneWithStruct6zepu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zepu > 0 || missingFieldsLeft6zepu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zepu, missingFieldsLeft6zepu, msgp.ShowFound(found6zepu[:]), decodeMsgFieldOrder6zepu)
		if encodedFieldsLeft6zepu > 0 {
			encodedFieldsLeft6zepu--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zepu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zepu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zepu = 0
			}
			for nextMiss6zepu < maxFields6zepu && (found6zepu[nextMiss6zepu] || decodeMsgFieldSkip6zepu[nextMiss6zepu]) {
				nextMiss6zepu++
			}
			if nextMiss6zepu == maxFields6zepu {
				// filled all the empty fields!
				break doneWithStruct6zepu
			}
			missingFieldsLeft6zepu--
			curField6zepu = decodeMsgFieldOrder6zepu[nextMiss6zepu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zepu)
		switch curField6zepu {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zepu[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zepu[1] = true
			var zzuv uint32
			zzuv, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zzuv) {
				z.Fields = (z.Fields)[:zzuv]
			} else {
				z.Fields = make([]Field, zzuv)
			}
			for zavw := range z.Fields {
				err = z.Fields[zavw].DecodeMsg(dc)
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
	if nextMiss6zepu != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zepu = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zepu = []bool{false, false}

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
	for zavw := range z.Fields {
		err = z.Fields[zavw].EncodeMsg(en)
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
	for zavw := range z.Fields {
		o, err = z.Fields[zavw].MarshalMsg(o)
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
	const maxFields7zldq = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zldq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zldq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zldq := totalEncodedFields7zldq
	missingFieldsLeft7zldq := maxFields7zldq - totalEncodedFields7zldq

	var nextMiss7zldq int32 = -1
	var found7zldq [maxFields7zldq]bool
	var curField7zldq string

doneWithStruct7zldq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zldq > 0 || missingFieldsLeft7zldq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zldq, missingFieldsLeft7zldq, msgp.ShowFound(found7zldq[:]), unmarshalMsgFieldOrder7zldq)
		if encodedFieldsLeft7zldq > 0 {
			encodedFieldsLeft7zldq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zldq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zldq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zldq = 0
			}
			for nextMiss7zldq < maxFields7zldq && (found7zldq[nextMiss7zldq] || unmarshalMsgFieldSkip7zldq[nextMiss7zldq]) {
				nextMiss7zldq++
			}
			if nextMiss7zldq == maxFields7zldq {
				// filled all the empty fields!
				break doneWithStruct7zldq
			}
			missingFieldsLeft7zldq--
			curField7zldq = unmarshalMsgFieldOrder7zldq[nextMiss7zldq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zldq)
		switch curField7zldq {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zldq[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zldq[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zhix uint32
				zhix, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zhix) {
					z.Fields = (z.Fields)[:zhix]
				} else {
					z.Fields = make([]Field, zhix)
				}
				for zavw := range z.Fields {
					bts, err = z.Fields[zavw].UnmarshalMsg(bts)
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
	if nextMiss7zldq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zldq = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zldq = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zavw := range z.Fields {
		s += z.Fields[zavw].Msgsize()
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
		var zlce uint64
		zlce, err = dc.ReadUint64()
		(*z) = Zkind(zlce)
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
		var zrcy uint64
		zrcy, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zrcy)
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
	const maxFields8zgkq = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zgkq uint32
	totalEncodedFields8zgkq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zgkq := totalEncodedFields8zgkq
	missingFieldsLeft8zgkq := maxFields8zgkq - totalEncodedFields8zgkq

	var nextMiss8zgkq int32 = -1
	var found8zgkq [maxFields8zgkq]bool
	var curField8zgkq string

doneWithStruct8zgkq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zgkq > 0 || missingFieldsLeft8zgkq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zgkq, missingFieldsLeft8zgkq, msgp.ShowFound(found8zgkq[:]), decodeMsgFieldOrder8zgkq)
		if encodedFieldsLeft8zgkq > 0 {
			encodedFieldsLeft8zgkq--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zgkq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zgkq < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zgkq = 0
			}
			for nextMiss8zgkq < maxFields8zgkq && (found8zgkq[nextMiss8zgkq] || decodeMsgFieldSkip8zgkq[nextMiss8zgkq]) {
				nextMiss8zgkq++
			}
			if nextMiss8zgkq == maxFields8zgkq {
				// filled all the empty fields!
				break doneWithStruct8zgkq
			}
			missingFieldsLeft8zgkq--
			curField8zgkq = decodeMsgFieldOrder8zgkq[nextMiss8zgkq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zgkq)
		switch curField8zgkq {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zgkq[0] = true
			{
				var zsmi uint64
				zsmi, err = dc.ReadUint64()
				z.Kind = Zkind(zsmi)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zgkq[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zgkq[2] = true
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
			found8zgkq[3] = true
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
	if nextMiss8zgkq != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zgkq = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zgkq = []bool{false, false, false, false}

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
	var empty_zqlg [4]bool
	fieldsInUse_zudz := z.fieldsNotEmpty(empty_zqlg[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zudz)
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
	if !empty_zqlg[1] {
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

	if !empty_zqlg[2] {
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

	if !empty_zqlg[3] {
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
	const maxFields9ztkc = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9ztkc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9ztkc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9ztkc := totalEncodedFields9ztkc
	missingFieldsLeft9ztkc := maxFields9ztkc - totalEncodedFields9ztkc

	var nextMiss9ztkc int32 = -1
	var found9ztkc [maxFields9ztkc]bool
	var curField9ztkc string

doneWithStruct9ztkc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9ztkc > 0 || missingFieldsLeft9ztkc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9ztkc, missingFieldsLeft9ztkc, msgp.ShowFound(found9ztkc[:]), unmarshalMsgFieldOrder9ztkc)
		if encodedFieldsLeft9ztkc > 0 {
			encodedFieldsLeft9ztkc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9ztkc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9ztkc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9ztkc = 0
			}
			for nextMiss9ztkc < maxFields9ztkc && (found9ztkc[nextMiss9ztkc] || unmarshalMsgFieldSkip9ztkc[nextMiss9ztkc]) {
				nextMiss9ztkc++
			}
			if nextMiss9ztkc == maxFields9ztkc {
				// filled all the empty fields!
				break doneWithStruct9ztkc
			}
			missingFieldsLeft9ztkc--
			curField9ztkc = unmarshalMsgFieldOrder9ztkc[nextMiss9ztkc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9ztkc)
		switch curField9ztkc {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9ztkc[0] = true
			{
				var zutu uint64
				zutu, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zutu)
			}
		case "Str":
			found9ztkc[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9ztkc[2] = true
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
			found9ztkc[3] = true
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
	if nextMiss9ztkc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9ztkc = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9ztkc = []bool{false, false, false, false}

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
