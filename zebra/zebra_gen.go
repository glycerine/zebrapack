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
	const maxFields0zkcx = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zkcx uint32
	totalEncodedFields0zkcx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zkcx := totalEncodedFields0zkcx
	missingFieldsLeft0zkcx := maxFields0zkcx - totalEncodedFields0zkcx

	var nextMiss0zkcx int32 = -1
	var found0zkcx [maxFields0zkcx]bool
	var curField0zkcx string

doneWithStruct0zkcx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zkcx > 0 || missingFieldsLeft0zkcx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zkcx, missingFieldsLeft0zkcx, msgp.ShowFound(found0zkcx[:]), decodeMsgFieldOrder0zkcx)
		if encodedFieldsLeft0zkcx > 0 {
			encodedFieldsLeft0zkcx--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zkcx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zkcx < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zkcx = 0
			}
			for nextMiss0zkcx < maxFields0zkcx && (found0zkcx[nextMiss0zkcx] || decodeMsgFieldSkip0zkcx[nextMiss0zkcx]) {
				nextMiss0zkcx++
			}
			if nextMiss0zkcx == maxFields0zkcx {
				// filled all the empty fields!
				break doneWithStruct0zkcx
			}
			missingFieldsLeft0zkcx--
			curField0zkcx = decodeMsgFieldOrder0zkcx[nextMiss0zkcx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zkcx)
		switch curField0zkcx {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zkcx[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zkcx[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zkcx[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zkcx[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zkcx[4] = true
			{
				var ztct uint64
				ztct, err = dc.ReadUint64()
				z.FieldCategory = Zkind(ztct)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zkcx[5] = true
			{
				var zlcd uint64
				zlcd, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zlcd)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zkcx[6] = true
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
			found0zkcx[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zkcx[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zkcx[9] = true
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
	if nextMiss0zkcx != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zkcx = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zkcx = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zvcr [10]bool
	fieldsInUse_zhjx := z.fieldsNotEmpty(empty_zvcr[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zhjx)
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
	if !empty_zvcr[3] {
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

	if !empty_zvcr[4] {
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

	if !empty_zvcr[5] {
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

	if !empty_zvcr[6] {
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

	if !empty_zvcr[7] {
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

	if !empty_zvcr[8] {
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

	if !empty_zvcr[9] {
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
	const maxFields1zhnz = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zhnz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhnz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zhnz := totalEncodedFields1zhnz
	missingFieldsLeft1zhnz := maxFields1zhnz - totalEncodedFields1zhnz

	var nextMiss1zhnz int32 = -1
	var found1zhnz [maxFields1zhnz]bool
	var curField1zhnz string

doneWithStruct1zhnz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhnz > 0 || missingFieldsLeft1zhnz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhnz, missingFieldsLeft1zhnz, msgp.ShowFound(found1zhnz[:]), unmarshalMsgFieldOrder1zhnz)
		if encodedFieldsLeft1zhnz > 0 {
			encodedFieldsLeft1zhnz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zhnz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zhnz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhnz = 0
			}
			for nextMiss1zhnz < maxFields1zhnz && (found1zhnz[nextMiss1zhnz] || unmarshalMsgFieldSkip1zhnz[nextMiss1zhnz]) {
				nextMiss1zhnz++
			}
			if nextMiss1zhnz == maxFields1zhnz {
				// filled all the empty fields!
				break doneWithStruct1zhnz
			}
			missingFieldsLeft1zhnz--
			curField1zhnz = unmarshalMsgFieldOrder1zhnz[nextMiss1zhnz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhnz)
		switch curField1zhnz {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zhnz[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zhnz[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zhnz[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zhnz[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zhnz[4] = true
			{
				var zldz uint64
				zldz, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zldz)
			}
		case "FieldPrimitive":
			found1zhnz[5] = true
			{
				var zwsi uint64
				zwsi, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zwsi)
			}
		case "FieldFullType":
			found1zhnz[6] = true
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
			found1zhnz[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zhnz[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zhnz[9] = true
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
	if nextMiss1zhnz != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zhnz = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zhnz = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zmlf = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zmlf uint32
	totalEncodedFields2zmlf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zmlf := totalEncodedFields2zmlf
	missingFieldsLeft2zmlf := maxFields2zmlf - totalEncodedFields2zmlf

	var nextMiss2zmlf int32 = -1
	var found2zmlf [maxFields2zmlf]bool
	var curField2zmlf string

doneWithStruct2zmlf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zmlf > 0 || missingFieldsLeft2zmlf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zmlf, missingFieldsLeft2zmlf, msgp.ShowFound(found2zmlf[:]), decodeMsgFieldOrder2zmlf)
		if encodedFieldsLeft2zmlf > 0 {
			encodedFieldsLeft2zmlf--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zmlf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zmlf < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zmlf = 0
			}
			for nextMiss2zmlf < maxFields2zmlf && (found2zmlf[nextMiss2zmlf] || decodeMsgFieldSkip2zmlf[nextMiss2zmlf]) {
				nextMiss2zmlf++
			}
			if nextMiss2zmlf == maxFields2zmlf {
				// filled all the empty fields!
				break doneWithStruct2zmlf
			}
			missingFieldsLeft2zmlf--
			curField2zmlf = decodeMsgFieldOrder2zmlf[nextMiss2zmlf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zmlf)
		switch curField2zmlf {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zmlf[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zmlf[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zmlf[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zmlf[3] = true
			var zdin uint32
			zdin, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zdin) {
				z.Structs = (z.Structs)[:zdin]
			} else {
				z.Structs = make([]Struct, zdin)
			}
			for zkzo := range z.Structs {
				const maxFields3zhfw = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zhfw uint32
				totalEncodedFields3zhfw, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zhfw := totalEncodedFields3zhfw
				missingFieldsLeft3zhfw := maxFields3zhfw - totalEncodedFields3zhfw

				var nextMiss3zhfw int32 = -1
				var found3zhfw [maxFields3zhfw]bool
				var curField3zhfw string

			doneWithStruct3zhfw:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zhfw > 0 || missingFieldsLeft3zhfw > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zhfw, missingFieldsLeft3zhfw, msgp.ShowFound(found3zhfw[:]), decodeMsgFieldOrder3zhfw)
					if encodedFieldsLeft3zhfw > 0 {
						encodedFieldsLeft3zhfw--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zhfw = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zhfw < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zhfw = 0
						}
						for nextMiss3zhfw < maxFields3zhfw && (found3zhfw[nextMiss3zhfw] || decodeMsgFieldSkip3zhfw[nextMiss3zhfw]) {
							nextMiss3zhfw++
						}
						if nextMiss3zhfw == maxFields3zhfw {
							// filled all the empty fields!
							break doneWithStruct3zhfw
						}
						missingFieldsLeft3zhfw--
						curField3zhfw = decodeMsgFieldOrder3zhfw[nextMiss3zhfw]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zhfw)
					switch curField3zhfw {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zhfw[0] = true
						z.Structs[zkzo].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zhfw[1] = true
						var zvga uint32
						zvga, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zkzo].Fields) >= int(zvga) {
							z.Structs[zkzo].Fields = (z.Structs[zkzo].Fields)[:zvga]
						} else {
							z.Structs[zkzo].Fields = make([]Field, zvga)
						}
						for zzws := range z.Structs[zkzo].Fields {
							err = z.Structs[zkzo].Fields[zzws].DecodeMsg(dc)
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
				if nextMiss3zhfw != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zmlf[4] = true
			var zfcu uint32
			zfcu, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zfcu) {
				z.Imports = (z.Imports)[:zfcu]
			} else {
				z.Imports = make([]string, zfcu)
			}
			for zzgy := range z.Imports {
				z.Imports[zzgy], err = dc.ReadString()
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
	if nextMiss2zmlf != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zmlf = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zmlf = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zhfw = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zhfw = []bool{false, false}

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
	for zkzo := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zkzo].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zkzo].Fields)))
		if err != nil {
			panic(err)
		}
		for zzws := range z.Structs[zkzo].Fields {
			err = z.Structs[zkzo].Fields[zzws].EncodeMsg(en)
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
	for zzgy := range z.Imports {
		err = en.WriteString(z.Imports[zzgy])
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
	for zkzo := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zkzo].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zkzo].Fields)))
		for zzws := range z.Structs[zkzo].Fields {
			o, err = z.Structs[zkzo].Fields[zzws].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zzgy := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zzgy])
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
	const maxFields4zkus = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zkus uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zkus, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zkus := totalEncodedFields4zkus
	missingFieldsLeft4zkus := maxFields4zkus - totalEncodedFields4zkus

	var nextMiss4zkus int32 = -1
	var found4zkus [maxFields4zkus]bool
	var curField4zkus string

doneWithStruct4zkus:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zkus > 0 || missingFieldsLeft4zkus > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zkus, missingFieldsLeft4zkus, msgp.ShowFound(found4zkus[:]), unmarshalMsgFieldOrder4zkus)
		if encodedFieldsLeft4zkus > 0 {
			encodedFieldsLeft4zkus--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zkus = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zkus < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zkus = 0
			}
			for nextMiss4zkus < maxFields4zkus && (found4zkus[nextMiss4zkus] || unmarshalMsgFieldSkip4zkus[nextMiss4zkus]) {
				nextMiss4zkus++
			}
			if nextMiss4zkus == maxFields4zkus {
				// filled all the empty fields!
				break doneWithStruct4zkus
			}
			missingFieldsLeft4zkus--
			curField4zkus = unmarshalMsgFieldOrder4zkus[nextMiss4zkus]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zkus)
		switch curField4zkus {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zkus[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zkus[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zkus[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zkus[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zldt uint32
				zldt, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zldt) {
					z.Structs = (z.Structs)[:zldt]
				} else {
					z.Structs = make([]Struct, zldt)
				}
				for zkzo := range z.Structs {
					const maxFields5zcfy = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zcfy uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zcfy, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zcfy := totalEncodedFields5zcfy
					missingFieldsLeft5zcfy := maxFields5zcfy - totalEncodedFields5zcfy

					var nextMiss5zcfy int32 = -1
					var found5zcfy [maxFields5zcfy]bool
					var curField5zcfy string

				doneWithStruct5zcfy:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zcfy > 0 || missingFieldsLeft5zcfy > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zcfy, missingFieldsLeft5zcfy, msgp.ShowFound(found5zcfy[:]), unmarshalMsgFieldOrder5zcfy)
						if encodedFieldsLeft5zcfy > 0 {
							encodedFieldsLeft5zcfy--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zcfy = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zcfy < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zcfy = 0
							}
							for nextMiss5zcfy < maxFields5zcfy && (found5zcfy[nextMiss5zcfy] || unmarshalMsgFieldSkip5zcfy[nextMiss5zcfy]) {
								nextMiss5zcfy++
							}
							if nextMiss5zcfy == maxFields5zcfy {
								// filled all the empty fields!
								break doneWithStruct5zcfy
							}
							missingFieldsLeft5zcfy--
							curField5zcfy = unmarshalMsgFieldOrder5zcfy[nextMiss5zcfy]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zcfy)
						switch curField5zcfy {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zcfy[0] = true
							z.Structs[zkzo].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zcfy[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zkzo].Fields) = (z.Structs[zkzo].Fields)[:0]
							} else {

								var zhsg uint32
								zhsg, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zkzo].Fields) >= int(zhsg) {
									z.Structs[zkzo].Fields = (z.Structs[zkzo].Fields)[:zhsg]
								} else {
									z.Structs[zkzo].Fields = make([]Field, zhsg)
								}
								for zzws := range z.Structs[zkzo].Fields {
									bts, err = z.Structs[zkzo].Fields[zzws].UnmarshalMsg(bts)
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
					if nextMiss5zcfy != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4zkus[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zcty uint32
				zcty, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zcty) {
					z.Imports = (z.Imports)[:zcty]
				} else {
					z.Imports = make([]string, zcty)
				}
				for zzgy := range z.Imports {
					z.Imports[zzgy], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zkus != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zkus = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zkus = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zcfy = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zcfy = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zkzo := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zkzo].StructName) + 7 + msgp.ArrayHeaderSize
		for zzws := range z.Structs[zkzo].Fields {
			s += z.Structs[zkzo].Fields[zzws].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zzgy := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zzgy])
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
	const maxFields6zjig = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zjig uint32
	totalEncodedFields6zjig, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zjig := totalEncodedFields6zjig
	missingFieldsLeft6zjig := maxFields6zjig - totalEncodedFields6zjig

	var nextMiss6zjig int32 = -1
	var found6zjig [maxFields6zjig]bool
	var curField6zjig string

doneWithStruct6zjig:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zjig > 0 || missingFieldsLeft6zjig > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zjig, missingFieldsLeft6zjig, msgp.ShowFound(found6zjig[:]), decodeMsgFieldOrder6zjig)
		if encodedFieldsLeft6zjig > 0 {
			encodedFieldsLeft6zjig--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zjig = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zjig < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zjig = 0
			}
			for nextMiss6zjig < maxFields6zjig && (found6zjig[nextMiss6zjig] || decodeMsgFieldSkip6zjig[nextMiss6zjig]) {
				nextMiss6zjig++
			}
			if nextMiss6zjig == maxFields6zjig {
				// filled all the empty fields!
				break doneWithStruct6zjig
			}
			missingFieldsLeft6zjig--
			curField6zjig = decodeMsgFieldOrder6zjig[nextMiss6zjig]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zjig)
		switch curField6zjig {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zjig[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zjig[1] = true
			var zxoa uint32
			zxoa, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zxoa) {
				z.Fields = (z.Fields)[:zxoa]
			} else {
				z.Fields = make([]Field, zxoa)
			}
			for zafv := range z.Fields {
				err = z.Fields[zafv].DecodeMsg(dc)
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
	if nextMiss6zjig != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zjig = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zjig = []bool{false, false}

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
	for zafv := range z.Fields {
		err = z.Fields[zafv].EncodeMsg(en)
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
	for zafv := range z.Fields {
		o, err = z.Fields[zafv].MarshalMsg(o)
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
	const maxFields7zgct = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zgct uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zgct, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zgct := totalEncodedFields7zgct
	missingFieldsLeft7zgct := maxFields7zgct - totalEncodedFields7zgct

	var nextMiss7zgct int32 = -1
	var found7zgct [maxFields7zgct]bool
	var curField7zgct string

doneWithStruct7zgct:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zgct > 0 || missingFieldsLeft7zgct > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgct, missingFieldsLeft7zgct, msgp.ShowFound(found7zgct[:]), unmarshalMsgFieldOrder7zgct)
		if encodedFieldsLeft7zgct > 0 {
			encodedFieldsLeft7zgct--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zgct = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zgct < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zgct = 0
			}
			for nextMiss7zgct < maxFields7zgct && (found7zgct[nextMiss7zgct] || unmarshalMsgFieldSkip7zgct[nextMiss7zgct]) {
				nextMiss7zgct++
			}
			if nextMiss7zgct == maxFields7zgct {
				// filled all the empty fields!
				break doneWithStruct7zgct
			}
			missingFieldsLeft7zgct--
			curField7zgct = unmarshalMsgFieldOrder7zgct[nextMiss7zgct]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zgct)
		switch curField7zgct {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zgct[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zgct[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zpai uint32
				zpai, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zpai) {
					z.Fields = (z.Fields)[:zpai]
				} else {
					z.Fields = make([]Field, zpai)
				}
				for zafv := range z.Fields {
					bts, err = z.Fields[zafv].UnmarshalMsg(bts)
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
	if nextMiss7zgct != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zgct = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zgct = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zafv := range z.Fields {
		s += z.Fields[zafv].Msgsize()
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
		var zajw uint64
		zajw, err = dc.ReadUint64()
		(*z) = Zkind(zajw)
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
		var zhvy uint64
		zhvy, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zhvy)
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
	const maxFields8zaps = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zaps uint32
	totalEncodedFields8zaps, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zaps := totalEncodedFields8zaps
	missingFieldsLeft8zaps := maxFields8zaps - totalEncodedFields8zaps

	var nextMiss8zaps int32 = -1
	var found8zaps [maxFields8zaps]bool
	var curField8zaps string

doneWithStruct8zaps:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zaps > 0 || missingFieldsLeft8zaps > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zaps, missingFieldsLeft8zaps, msgp.ShowFound(found8zaps[:]), decodeMsgFieldOrder8zaps)
		if encodedFieldsLeft8zaps > 0 {
			encodedFieldsLeft8zaps--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zaps = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zaps < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zaps = 0
			}
			for nextMiss8zaps < maxFields8zaps && (found8zaps[nextMiss8zaps] || decodeMsgFieldSkip8zaps[nextMiss8zaps]) {
				nextMiss8zaps++
			}
			if nextMiss8zaps == maxFields8zaps {
				// filled all the empty fields!
				break doneWithStruct8zaps
			}
			missingFieldsLeft8zaps--
			curField8zaps = decodeMsgFieldOrder8zaps[nextMiss8zaps]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zaps)
		switch curField8zaps {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zaps[0] = true
			{
				var zhrs uint64
				zhrs, err = dc.ReadUint64()
				z.Kind = Zkind(zhrs)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zaps[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zaps[2] = true
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
			found8zaps[3] = true
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
	if nextMiss8zaps != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zaps = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zaps = []bool{false, false, false, false}

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
	var empty_zcfa [4]bool
	fieldsInUse_zwey := z.fieldsNotEmpty(empty_zcfa[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zwey)
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
	if !empty_zcfa[1] {
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

	if !empty_zcfa[2] {
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

	if !empty_zcfa[3] {
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
	const maxFields9zsjt = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zsjt uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zsjt, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zsjt := totalEncodedFields9zsjt
	missingFieldsLeft9zsjt := maxFields9zsjt - totalEncodedFields9zsjt

	var nextMiss9zsjt int32 = -1
	var found9zsjt [maxFields9zsjt]bool
	var curField9zsjt string

doneWithStruct9zsjt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zsjt > 0 || missingFieldsLeft9zsjt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zsjt, missingFieldsLeft9zsjt, msgp.ShowFound(found9zsjt[:]), unmarshalMsgFieldOrder9zsjt)
		if encodedFieldsLeft9zsjt > 0 {
			encodedFieldsLeft9zsjt--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zsjt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zsjt < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zsjt = 0
			}
			for nextMiss9zsjt < maxFields9zsjt && (found9zsjt[nextMiss9zsjt] || unmarshalMsgFieldSkip9zsjt[nextMiss9zsjt]) {
				nextMiss9zsjt++
			}
			if nextMiss9zsjt == maxFields9zsjt {
				// filled all the empty fields!
				break doneWithStruct9zsjt
			}
			missingFieldsLeft9zsjt--
			curField9zsjt = unmarshalMsgFieldOrder9zsjt[nextMiss9zsjt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zsjt)
		switch curField9zsjt {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zsjt[0] = true
			{
				var zwqb uint64
				zwqb, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zwqb)
			}
		case "Str":
			found9zsjt[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zsjt[2] = true
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
			found9zsjt[3] = true
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
	if nextMiss9zsjt != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zsjt = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zsjt = []bool{false, false, false, false}

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
