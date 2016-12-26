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
	const maxFields0zase = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zase uint32
	totalEncodedFields0zase, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zase := totalEncodedFields0zase
	missingFieldsLeft0zase := maxFields0zase - totalEncodedFields0zase

	var nextMiss0zase int32 = -1
	var found0zase [maxFields0zase]bool
	var curField0zase string

doneWithStruct0zase:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zase > 0 || missingFieldsLeft0zase > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zase, missingFieldsLeft0zase, msgp.ShowFound(found0zase[:]), decodeMsgFieldOrder0zase)
		if encodedFieldsLeft0zase > 0 {
			encodedFieldsLeft0zase--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zase = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zase < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zase = 0
			}
			for nextMiss0zase < maxFields0zase && (found0zase[nextMiss0zase] || decodeMsgFieldSkip0zase[nextMiss0zase]) {
				nextMiss0zase++
			}
			if nextMiss0zase == maxFields0zase {
				// filled all the empty fields!
				break doneWithStruct0zase
			}
			missingFieldsLeft0zase--
			curField0zase = decodeMsgFieldOrder0zase[nextMiss0zase]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zase)
		switch curField0zase {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zase[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zase[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zase[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zase[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zase[4] = true
			{
				var zjqn uint64
				zjqn, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zjqn)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zase[5] = true
			{
				var zlmm uint64
				zlmm, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zlmm)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zase[6] = true
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
			found0zase[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zase[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zase[9] = true
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
	if nextMiss0zase != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zase = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zase = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zypg [10]bool
	fieldsInUse_zeoz := z.fieldsNotEmpty(empty_zypg[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zeoz)
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
	if !empty_zypg[3] {
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

	if !empty_zypg[4] {
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

	if !empty_zypg[5] {
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

	if !empty_zypg[6] {
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

	if !empty_zypg[7] {
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

	if !empty_zypg[8] {
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

	if !empty_zypg[9] {
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
	const maxFields1zaii = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zaii uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zaii, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zaii := totalEncodedFields1zaii
	missingFieldsLeft1zaii := maxFields1zaii - totalEncodedFields1zaii

	var nextMiss1zaii int32 = -1
	var found1zaii [maxFields1zaii]bool
	var curField1zaii string

doneWithStruct1zaii:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zaii > 0 || missingFieldsLeft1zaii > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zaii, missingFieldsLeft1zaii, msgp.ShowFound(found1zaii[:]), unmarshalMsgFieldOrder1zaii)
		if encodedFieldsLeft1zaii > 0 {
			encodedFieldsLeft1zaii--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zaii = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zaii < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zaii = 0
			}
			for nextMiss1zaii < maxFields1zaii && (found1zaii[nextMiss1zaii] || unmarshalMsgFieldSkip1zaii[nextMiss1zaii]) {
				nextMiss1zaii++
			}
			if nextMiss1zaii == maxFields1zaii {
				// filled all the empty fields!
				break doneWithStruct1zaii
			}
			missingFieldsLeft1zaii--
			curField1zaii = unmarshalMsgFieldOrder1zaii[nextMiss1zaii]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zaii)
		switch curField1zaii {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zaii[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zaii[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zaii[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zaii[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zaii[4] = true
			{
				var zsoa uint64
				zsoa, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zsoa)
			}
		case "FieldPrimitive":
			found1zaii[5] = true
			{
				var zhfe uint64
				zhfe, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zhfe)
			}
		case "FieldFullType":
			found1zaii[6] = true
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
			found1zaii[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zaii[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zaii[9] = true
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
	if nextMiss1zaii != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zaii = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zaii = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zdkm = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zdkm uint32
	totalEncodedFields2zdkm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zdkm := totalEncodedFields2zdkm
	missingFieldsLeft2zdkm := maxFields2zdkm - totalEncodedFields2zdkm

	var nextMiss2zdkm int32 = -1
	var found2zdkm [maxFields2zdkm]bool
	var curField2zdkm string

doneWithStruct2zdkm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zdkm > 0 || missingFieldsLeft2zdkm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zdkm, missingFieldsLeft2zdkm, msgp.ShowFound(found2zdkm[:]), decodeMsgFieldOrder2zdkm)
		if encodedFieldsLeft2zdkm > 0 {
			encodedFieldsLeft2zdkm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zdkm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zdkm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zdkm = 0
			}
			for nextMiss2zdkm < maxFields2zdkm && (found2zdkm[nextMiss2zdkm] || decodeMsgFieldSkip2zdkm[nextMiss2zdkm]) {
				nextMiss2zdkm++
			}
			if nextMiss2zdkm == maxFields2zdkm {
				// filled all the empty fields!
				break doneWithStruct2zdkm
			}
			missingFieldsLeft2zdkm--
			curField2zdkm = decodeMsgFieldOrder2zdkm[nextMiss2zdkm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zdkm)
		switch curField2zdkm {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zdkm[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zdkm[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zdkm[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zdkm[3] = true
			var ztsw uint32
			ztsw, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(ztsw) {
				z.Structs = (z.Structs)[:ztsw]
			} else {
				z.Structs = make([]Struct, ztsw)
			}
			for zlof := range z.Structs {
				const maxFields3zczb = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zczb uint32
				totalEncodedFields3zczb, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zczb := totalEncodedFields3zczb
				missingFieldsLeft3zczb := maxFields3zczb - totalEncodedFields3zczb

				var nextMiss3zczb int32 = -1
				var found3zczb [maxFields3zczb]bool
				var curField3zczb string

			doneWithStruct3zczb:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zczb > 0 || missingFieldsLeft3zczb > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zczb, missingFieldsLeft3zczb, msgp.ShowFound(found3zczb[:]), decodeMsgFieldOrder3zczb)
					if encodedFieldsLeft3zczb > 0 {
						encodedFieldsLeft3zczb--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zczb = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zczb < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zczb = 0
						}
						for nextMiss3zczb < maxFields3zczb && (found3zczb[nextMiss3zczb] || decodeMsgFieldSkip3zczb[nextMiss3zczb]) {
							nextMiss3zczb++
						}
						if nextMiss3zczb == maxFields3zczb {
							// filled all the empty fields!
							break doneWithStruct3zczb
						}
						missingFieldsLeft3zczb--
						curField3zczb = decodeMsgFieldOrder3zczb[nextMiss3zczb]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zczb)
					switch curField3zczb {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zczb[0] = true
						z.Structs[zlof].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zczb[1] = true
						var zsgb uint32
						zsgb, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zlof].Fields) >= int(zsgb) {
							z.Structs[zlof].Fields = (z.Structs[zlof].Fields)[:zsgb]
						} else {
							z.Structs[zlof].Fields = make([]Field, zsgb)
						}
						for zrao := range z.Structs[zlof].Fields {
							err = z.Structs[zlof].Fields[zrao].DecodeMsg(dc)
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
				if nextMiss3zczb != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zdkm[4] = true
			var zcxz uint32
			zcxz, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zcxz) {
				z.Imports = (z.Imports)[:zcxz]
			} else {
				z.Imports = make([]string, zcxz)
			}
			for zrdo := range z.Imports {
				z.Imports[zrdo], err = dc.ReadString()
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
	if nextMiss2zdkm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zdkm = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zdkm = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zczb = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zczb = []bool{false, false}

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
	for zlof := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zlof].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zlof].Fields)))
		if err != nil {
			panic(err)
		}
		for zrao := range z.Structs[zlof].Fields {
			err = z.Structs[zlof].Fields[zrao].EncodeMsg(en)
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
	for zrdo := range z.Imports {
		err = en.WriteString(z.Imports[zrdo])
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
	for zlof := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zlof].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zlof].Fields)))
		for zrao := range z.Structs[zlof].Fields {
			o, err = z.Structs[zlof].Fields[zrao].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zrdo := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zrdo])
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
	const maxFields4zlbk = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zlbk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zlbk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zlbk := totalEncodedFields4zlbk
	missingFieldsLeft4zlbk := maxFields4zlbk - totalEncodedFields4zlbk

	var nextMiss4zlbk int32 = -1
	var found4zlbk [maxFields4zlbk]bool
	var curField4zlbk string

doneWithStruct4zlbk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zlbk > 0 || missingFieldsLeft4zlbk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zlbk, missingFieldsLeft4zlbk, msgp.ShowFound(found4zlbk[:]), unmarshalMsgFieldOrder4zlbk)
		if encodedFieldsLeft4zlbk > 0 {
			encodedFieldsLeft4zlbk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zlbk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zlbk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zlbk = 0
			}
			for nextMiss4zlbk < maxFields4zlbk && (found4zlbk[nextMiss4zlbk] || unmarshalMsgFieldSkip4zlbk[nextMiss4zlbk]) {
				nextMiss4zlbk++
			}
			if nextMiss4zlbk == maxFields4zlbk {
				// filled all the empty fields!
				break doneWithStruct4zlbk
			}
			missingFieldsLeft4zlbk--
			curField4zlbk = unmarshalMsgFieldOrder4zlbk[nextMiss4zlbk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zlbk)
		switch curField4zlbk {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zlbk[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zlbk[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zlbk[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zlbk[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zajc uint32
				zajc, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zajc) {
					z.Structs = (z.Structs)[:zajc]
				} else {
					z.Structs = make([]Struct, zajc)
				}
				for zlof := range z.Structs {
					const maxFields5zosm = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zosm uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zosm, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zosm := totalEncodedFields5zosm
					missingFieldsLeft5zosm := maxFields5zosm - totalEncodedFields5zosm

					var nextMiss5zosm int32 = -1
					var found5zosm [maxFields5zosm]bool
					var curField5zosm string

				doneWithStruct5zosm:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zosm > 0 || missingFieldsLeft5zosm > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zosm, missingFieldsLeft5zosm, msgp.ShowFound(found5zosm[:]), unmarshalMsgFieldOrder5zosm)
						if encodedFieldsLeft5zosm > 0 {
							encodedFieldsLeft5zosm--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zosm = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zosm < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zosm = 0
							}
							for nextMiss5zosm < maxFields5zosm && (found5zosm[nextMiss5zosm] || unmarshalMsgFieldSkip5zosm[nextMiss5zosm]) {
								nextMiss5zosm++
							}
							if nextMiss5zosm == maxFields5zosm {
								// filled all the empty fields!
								break doneWithStruct5zosm
							}
							missingFieldsLeft5zosm--
							curField5zosm = unmarshalMsgFieldOrder5zosm[nextMiss5zosm]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zosm)
						switch curField5zosm {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zosm[0] = true
							z.Structs[zlof].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zosm[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zlof].Fields) = (z.Structs[zlof].Fields)[:0]
							} else {

								var ziaj uint32
								ziaj, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zlof].Fields) >= int(ziaj) {
									z.Structs[zlof].Fields = (z.Structs[zlof].Fields)[:ziaj]
								} else {
									z.Structs[zlof].Fields = make([]Field, ziaj)
								}
								for zrao := range z.Structs[zlof].Fields {
									bts, err = z.Structs[zlof].Fields[zrao].UnmarshalMsg(bts)
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
					if nextMiss5zosm != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4zlbk[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zjjh uint32
				zjjh, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zjjh) {
					z.Imports = (z.Imports)[:zjjh]
				} else {
					z.Imports = make([]string, zjjh)
				}
				for zrdo := range z.Imports {
					z.Imports[zrdo], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zlbk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zlbk = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zlbk = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zosm = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zosm = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zlof := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zlof].StructName) + 7 + msgp.ArrayHeaderSize
		for zrao := range z.Structs[zlof].Fields {
			s += z.Structs[zlof].Fields[zrao].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zrdo := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zrdo])
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
	const maxFields6zknm = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zknm uint32
	totalEncodedFields6zknm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zknm := totalEncodedFields6zknm
	missingFieldsLeft6zknm := maxFields6zknm - totalEncodedFields6zknm

	var nextMiss6zknm int32 = -1
	var found6zknm [maxFields6zknm]bool
	var curField6zknm string

doneWithStruct6zknm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zknm > 0 || missingFieldsLeft6zknm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zknm, missingFieldsLeft6zknm, msgp.ShowFound(found6zknm[:]), decodeMsgFieldOrder6zknm)
		if encodedFieldsLeft6zknm > 0 {
			encodedFieldsLeft6zknm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zknm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zknm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zknm = 0
			}
			for nextMiss6zknm < maxFields6zknm && (found6zknm[nextMiss6zknm] || decodeMsgFieldSkip6zknm[nextMiss6zknm]) {
				nextMiss6zknm++
			}
			if nextMiss6zknm == maxFields6zknm {
				// filled all the empty fields!
				break doneWithStruct6zknm
			}
			missingFieldsLeft6zknm--
			curField6zknm = decodeMsgFieldOrder6zknm[nextMiss6zknm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zknm)
		switch curField6zknm {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zknm[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zknm[1] = true
			var zkiz uint32
			zkiz, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zkiz) {
				z.Fields = (z.Fields)[:zkiz]
			} else {
				z.Fields = make([]Field, zkiz)
			}
			for zrxf := range z.Fields {
				err = z.Fields[zrxf].DecodeMsg(dc)
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
	if nextMiss6zknm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zknm = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zknm = []bool{false, false}

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
	for zrxf := range z.Fields {
		err = z.Fields[zrxf].EncodeMsg(en)
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
	for zrxf := range z.Fields {
		o, err = z.Fields[zrxf].MarshalMsg(o)
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
	const maxFields7zvbt = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zvbt uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zvbt, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zvbt := totalEncodedFields7zvbt
	missingFieldsLeft7zvbt := maxFields7zvbt - totalEncodedFields7zvbt

	var nextMiss7zvbt int32 = -1
	var found7zvbt [maxFields7zvbt]bool
	var curField7zvbt string

doneWithStruct7zvbt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zvbt > 0 || missingFieldsLeft7zvbt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zvbt, missingFieldsLeft7zvbt, msgp.ShowFound(found7zvbt[:]), unmarshalMsgFieldOrder7zvbt)
		if encodedFieldsLeft7zvbt > 0 {
			encodedFieldsLeft7zvbt--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zvbt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zvbt < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zvbt = 0
			}
			for nextMiss7zvbt < maxFields7zvbt && (found7zvbt[nextMiss7zvbt] || unmarshalMsgFieldSkip7zvbt[nextMiss7zvbt]) {
				nextMiss7zvbt++
			}
			if nextMiss7zvbt == maxFields7zvbt {
				// filled all the empty fields!
				break doneWithStruct7zvbt
			}
			missingFieldsLeft7zvbt--
			curField7zvbt = unmarshalMsgFieldOrder7zvbt[nextMiss7zvbt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zvbt)
		switch curField7zvbt {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zvbt[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zvbt[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zrxa uint32
				zrxa, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zrxa) {
					z.Fields = (z.Fields)[:zrxa]
				} else {
					z.Fields = make([]Field, zrxa)
				}
				for zrxf := range z.Fields {
					bts, err = z.Fields[zrxf].UnmarshalMsg(bts)
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
	if nextMiss7zvbt != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zvbt = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zvbt = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zrxf := range z.Fields {
		s += z.Fields[zrxf].Msgsize()
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
		var zudb uint64
		zudb, err = dc.ReadUint64()
		(*z) = Zkind(zudb)
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
		var zqpe uint64
		zqpe, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zqpe)
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
	const maxFields8zwyr = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zwyr uint32
	totalEncodedFields8zwyr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zwyr := totalEncodedFields8zwyr
	missingFieldsLeft8zwyr := maxFields8zwyr - totalEncodedFields8zwyr

	var nextMiss8zwyr int32 = -1
	var found8zwyr [maxFields8zwyr]bool
	var curField8zwyr string

doneWithStruct8zwyr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zwyr > 0 || missingFieldsLeft8zwyr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zwyr, missingFieldsLeft8zwyr, msgp.ShowFound(found8zwyr[:]), decodeMsgFieldOrder8zwyr)
		if encodedFieldsLeft8zwyr > 0 {
			encodedFieldsLeft8zwyr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zwyr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zwyr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zwyr = 0
			}
			for nextMiss8zwyr < maxFields8zwyr && (found8zwyr[nextMiss8zwyr] || decodeMsgFieldSkip8zwyr[nextMiss8zwyr]) {
				nextMiss8zwyr++
			}
			if nextMiss8zwyr == maxFields8zwyr {
				// filled all the empty fields!
				break doneWithStruct8zwyr
			}
			missingFieldsLeft8zwyr--
			curField8zwyr = decodeMsgFieldOrder8zwyr[nextMiss8zwyr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zwyr)
		switch curField8zwyr {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zwyr[0] = true
			{
				var zpjc uint64
				zpjc, err = dc.ReadUint64()
				z.Kind = Zkind(zpjc)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zwyr[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zwyr[2] = true
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
			found8zwyr[3] = true
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
	if nextMiss8zwyr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zwyr = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zwyr = []bool{false, false, false, false}

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
	var empty_zdxd [4]bool
	fieldsInUse_zyff := z.fieldsNotEmpty(empty_zdxd[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zyff)
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
	if !empty_zdxd[1] {
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

	if !empty_zdxd[2] {
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

	if !empty_zdxd[3] {
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
	const maxFields9zxji = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zxji uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zxji, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zxji := totalEncodedFields9zxji
	missingFieldsLeft9zxji := maxFields9zxji - totalEncodedFields9zxji

	var nextMiss9zxji int32 = -1
	var found9zxji [maxFields9zxji]bool
	var curField9zxji string

doneWithStruct9zxji:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zxji > 0 || missingFieldsLeft9zxji > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zxji, missingFieldsLeft9zxji, msgp.ShowFound(found9zxji[:]), unmarshalMsgFieldOrder9zxji)
		if encodedFieldsLeft9zxji > 0 {
			encodedFieldsLeft9zxji--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zxji = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zxji < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zxji = 0
			}
			for nextMiss9zxji < maxFields9zxji && (found9zxji[nextMiss9zxji] || unmarshalMsgFieldSkip9zxji[nextMiss9zxji]) {
				nextMiss9zxji++
			}
			if nextMiss9zxji == maxFields9zxji {
				// filled all the empty fields!
				break doneWithStruct9zxji
			}
			missingFieldsLeft9zxji--
			curField9zxji = unmarshalMsgFieldOrder9zxji[nextMiss9zxji]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zxji)
		switch curField9zxji {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zxji[0] = true
			{
				var zzlb uint64
				zzlb, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zzlb)
			}
		case "Str":
			found9zxji[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zxji[2] = true
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
			found9zxji[3] = true
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
	if nextMiss9zxji != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zxji = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zxji = []bool{false, false, false, false}

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
