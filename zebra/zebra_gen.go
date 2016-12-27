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
	const maxFields0zkyy = 10

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zkyy uint32
	totalEncodedFields0zkyy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zkyy := totalEncodedFields0zkyy
	missingFieldsLeft0zkyy := maxFields0zkyy - totalEncodedFields0zkyy

	var nextMiss0zkyy int32 = -1
	var found0zkyy [maxFields0zkyy]bool
	var curField0zkyy string

doneWithStruct0zkyy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zkyy > 0 || missingFieldsLeft0zkyy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zkyy, missingFieldsLeft0zkyy, msgp.ShowFound(found0zkyy[:]), decodeMsgFieldOrder0zkyy)
		if encodedFieldsLeft0zkyy > 0 {
			encodedFieldsLeft0zkyy--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zkyy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zkyy < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zkyy = 0
			}
			for nextMiss0zkyy < maxFields0zkyy && (found0zkyy[nextMiss0zkyy] || decodeMsgFieldSkip0zkyy[nextMiss0zkyy]) {
				nextMiss0zkyy++
			}
			if nextMiss0zkyy == maxFields0zkyy {
				// filled all the empty fields!
				break doneWithStruct0zkyy
			}
			missingFieldsLeft0zkyy--
			curField0zkyy = decodeMsgFieldOrder0zkyy[nextMiss0zkyy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zkyy)
		switch curField0zkyy {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zkyy[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found0zkyy[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found0zkyy[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found0zkyy[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found0zkyy[4] = true
			{
				var zpgf uint64
				zpgf, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zpgf)
			}
			if err != nil {
				panic(err)
			}
		case "FieldPrimitive":
			found0zkyy[5] = true
			{
				var zuxy uint64
				zuxy, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zuxy)
			}
			if err != nil {
				panic(err)
			}
		case "FieldFullType":
			found0zkyy[6] = true
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
			found0zkyy[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Skip":
			found0zkyy[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found0zkyy[9] = true
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
	if nextMiss0zkyy != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Field
var decodeMsgFieldOrder0zkyy = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var decodeMsgFieldSkip0zkyy = []bool{false, false, false, false, false, false, false, false, false, false}

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
	var empty_zadh [10]bool
	fieldsInUse_zcqw := z.fieldsNotEmpty(empty_zadh[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zcqw)
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
	if !empty_zadh[3] {
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

	if !empty_zadh[4] {
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

	if !empty_zadh[5] {
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

	if !empty_zadh[6] {
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

	if !empty_zadh[7] {
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

	if !empty_zadh[8] {
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

	if !empty_zadh[9] {
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
	const maxFields1zlck = 10

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zlck uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zlck, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zlck := totalEncodedFields1zlck
	missingFieldsLeft1zlck := maxFields1zlck - totalEncodedFields1zlck

	var nextMiss1zlck int32 = -1
	var found1zlck [maxFields1zlck]bool
	var curField1zlck string

doneWithStruct1zlck:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zlck > 0 || missingFieldsLeft1zlck > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zlck, missingFieldsLeft1zlck, msgp.ShowFound(found1zlck[:]), unmarshalMsgFieldOrder1zlck)
		if encodedFieldsLeft1zlck > 0 {
			encodedFieldsLeft1zlck--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zlck = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zlck < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zlck = 0
			}
			for nextMiss1zlck < maxFields1zlck && (found1zlck[nextMiss1zlck] || unmarshalMsgFieldSkip1zlck[nextMiss1zlck]) {
				nextMiss1zlck++
			}
			if nextMiss1zlck == maxFields1zlck {
				// filled all the empty fields!
				break doneWithStruct1zlck
			}
			missingFieldsLeft1zlck--
			curField1zlck = unmarshalMsgFieldOrder1zlck[nextMiss1zlck]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zlck)
		switch curField1zlck {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zlck[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldGoName":
			found1zlck[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTagName":
			found1zlck[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldTypeStr":
			found1zlck[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "FieldCategory":
			found1zlck[4] = true
			{
				var zkdo uint64
				zkdo, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldCategory = Zkind(zkdo)
			}
		case "FieldPrimitive":
			found1zlck[5] = true
			{
				var zeiu uint64
				zeiu, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.FieldPrimitive = Zkind(zeiu)
			}
		case "FieldFullType":
			found1zlck[6] = true
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
			found1zlck[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Skip":
			found1zlck[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Deprecated":
			found1zlck[9] = true
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
	if nextMiss1zlck != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Field
var unmarshalMsgFieldOrder1zlck = []string{"Zid", "FieldGoName", "FieldTagName", "FieldTypeStr", "FieldCategory", "FieldPrimitive", "FieldFullType", "OmitEmpty", "Skip", "Deprecated"}

var unmarshalMsgFieldSkip1zlck = []bool{false, false, false, false, false, false, false, false, false, false}

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
	const maxFields2zdrj = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zdrj uint32
	totalEncodedFields2zdrj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zdrj := totalEncodedFields2zdrj
	missingFieldsLeft2zdrj := maxFields2zdrj - totalEncodedFields2zdrj

	var nextMiss2zdrj int32 = -1
	var found2zdrj [maxFields2zdrj]bool
	var curField2zdrj string

doneWithStruct2zdrj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zdrj > 0 || missingFieldsLeft2zdrj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zdrj, missingFieldsLeft2zdrj, msgp.ShowFound(found2zdrj[:]), decodeMsgFieldOrder2zdrj)
		if encodedFieldsLeft2zdrj > 0 {
			encodedFieldsLeft2zdrj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zdrj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zdrj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zdrj = 0
			}
			for nextMiss2zdrj < maxFields2zdrj && (found2zdrj[nextMiss2zdrj] || decodeMsgFieldSkip2zdrj[nextMiss2zdrj]) {
				nextMiss2zdrj++
			}
			if nextMiss2zdrj == maxFields2zdrj {
				// filled all the empty fields!
				break doneWithStruct2zdrj
			}
			missingFieldsLeft2zdrj--
			curField2zdrj = decodeMsgFieldOrder2zdrj[nextMiss2zdrj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zdrj)
		switch curField2zdrj {
		// -- templateDecodeMsg ends here --

		case "SourcePath":
			found2zdrj[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found2zdrj[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found2zdrj[2] = true
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found2zdrj[3] = true
			var zhej uint32
			zhej, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zhej) {
				z.Structs = (z.Structs)[:zhej]
			} else {
				z.Structs = make([]Struct, zhej)
			}
			for zgwx := range z.Structs {
				const maxFields3zcdb = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields3zcdb uint32
				totalEncodedFields3zcdb, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft3zcdb := totalEncodedFields3zcdb
				missingFieldsLeft3zcdb := maxFields3zcdb - totalEncodedFields3zcdb

				var nextMiss3zcdb int32 = -1
				var found3zcdb [maxFields3zcdb]bool
				var curField3zcdb string

			doneWithStruct3zcdb:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft3zcdb > 0 || missingFieldsLeft3zcdb > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zcdb, missingFieldsLeft3zcdb, msgp.ShowFound(found3zcdb[:]), decodeMsgFieldOrder3zcdb)
					if encodedFieldsLeft3zcdb > 0 {
						encodedFieldsLeft3zcdb--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField3zcdb = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss3zcdb < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss3zcdb = 0
						}
						for nextMiss3zcdb < maxFields3zcdb && (found3zcdb[nextMiss3zcdb] || decodeMsgFieldSkip3zcdb[nextMiss3zcdb]) {
							nextMiss3zcdb++
						}
						if nextMiss3zcdb == maxFields3zcdb {
							// filled all the empty fields!
							break doneWithStruct3zcdb
						}
						missingFieldsLeft3zcdb--
						curField3zcdb = decodeMsgFieldOrder3zcdb[nextMiss3zcdb]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField3zcdb)
					switch curField3zcdb {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found3zcdb[0] = true
						z.Structs[zgwx].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fields":
						found3zcdb[1] = true
						var zrko uint32
						zrko, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zgwx].Fields) >= int(zrko) {
							z.Structs[zgwx].Fields = (z.Structs[zgwx].Fields)[:zrko]
						} else {
							z.Structs[zgwx].Fields = make([]Field, zrko)
						}
						for zbei := range z.Structs[zgwx].Fields {
							err = z.Structs[zgwx].Fields[zbei].DecodeMsg(dc)
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
				if nextMiss3zcdb != -1 {
					dc.PopAlwaysNil()
				}

			}
		case "Imports":
			found2zdrj[4] = true
			var zkak uint32
			zkak, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Imports) >= int(zkak) {
				z.Imports = (z.Imports)[:zkak]
			} else {
				z.Imports = make([]string, zkak)
			}
			for zzfc := range z.Imports {
				z.Imports[zzfc], err = dc.ReadString()
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
	if nextMiss2zdrj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Schema
var decodeMsgFieldOrder2zdrj = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var decodeMsgFieldSkip2zdrj = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zcdb = []string{"StructName", "Fields"}

var decodeMsgFieldSkip3zcdb = []bool{false, false}

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
	for zgwx := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zgwx].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zgwx].Fields)))
		if err != nil {
			panic(err)
		}
		for zbei := range z.Structs[zgwx].Fields {
			err = z.Structs[zgwx].Fields[zbei].EncodeMsg(en)
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
	for zzfc := range z.Imports {
		err = en.WriteString(z.Imports[zzfc])
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
	for zgwx := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zgwx].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zgwx].Fields)))
		for zbei := range z.Structs[zgwx].Fields {
			o, err = z.Structs[zgwx].Fields[zbei].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Imports"
	o = append(o, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zzfc := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zzfc])
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
	const maxFields4zulh = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zulh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zulh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft4zulh := totalEncodedFields4zulh
	missingFieldsLeft4zulh := maxFields4zulh - totalEncodedFields4zulh

	var nextMiss4zulh int32 = -1
	var found4zulh [maxFields4zulh]bool
	var curField4zulh string

doneWithStruct4zulh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zulh > 0 || missingFieldsLeft4zulh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zulh, missingFieldsLeft4zulh, msgp.ShowFound(found4zulh[:]), unmarshalMsgFieldOrder4zulh)
		if encodedFieldsLeft4zulh > 0 {
			encodedFieldsLeft4zulh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField4zulh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zulh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zulh = 0
			}
			for nextMiss4zulh < maxFields4zulh && (found4zulh[nextMiss4zulh] || unmarshalMsgFieldSkip4zulh[nextMiss4zulh]) {
				nextMiss4zulh++
			}
			if nextMiss4zulh == maxFields4zulh {
				// filled all the empty fields!
				break doneWithStruct4zulh
			}
			missingFieldsLeft4zulh--
			curField4zulh = unmarshalMsgFieldOrder4zulh[nextMiss4zulh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zulh)
		switch curField4zulh {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath":
			found4zulh[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "SourcePackage":
			found4zulh[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "ZebraSchemaId":
			found4zulh[2] = true
			z.ZebraSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zulh[3] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zlqz uint32
				zlqz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zlqz) {
					z.Structs = (z.Structs)[:zlqz]
				} else {
					z.Structs = make([]Struct, zlqz)
				}
				for zgwx := range z.Structs {
					const maxFields5zxzr = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields5zxzr uint32
					if !nbs.AlwaysNil {
						totalEncodedFields5zxzr, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft5zxzr := totalEncodedFields5zxzr
					missingFieldsLeft5zxzr := maxFields5zxzr - totalEncodedFields5zxzr

					var nextMiss5zxzr int32 = -1
					var found5zxzr [maxFields5zxzr]bool
					var curField5zxzr string

				doneWithStruct5zxzr:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft5zxzr > 0 || missingFieldsLeft5zxzr > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxzr, missingFieldsLeft5zxzr, msgp.ShowFound(found5zxzr[:]), unmarshalMsgFieldOrder5zxzr)
						if encodedFieldsLeft5zxzr > 0 {
							encodedFieldsLeft5zxzr--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField5zxzr = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss5zxzr < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss5zxzr = 0
							}
							for nextMiss5zxzr < maxFields5zxzr && (found5zxzr[nextMiss5zxzr] || unmarshalMsgFieldSkip5zxzr[nextMiss5zxzr]) {
								nextMiss5zxzr++
							}
							if nextMiss5zxzr == maxFields5zxzr {
								// filled all the empty fields!
								break doneWithStruct5zxzr
							}
							missingFieldsLeft5zxzr--
							curField5zxzr = unmarshalMsgFieldOrder5zxzr[nextMiss5zxzr]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField5zxzr)
						switch curField5zxzr {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found5zxzr[0] = true
							z.Structs[zgwx].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fields":
							found5zxzr[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zgwx].Fields) = (z.Structs[zgwx].Fields)[:0]
							} else {

								var zukf uint32
								zukf, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zgwx].Fields) >= int(zukf) {
									z.Structs[zgwx].Fields = (z.Structs[zgwx].Fields)[:zukf]
								} else {
									z.Structs[zgwx].Fields = make([]Field, zukf)
								}
								for zbei := range z.Structs[zgwx].Fields {
									bts, err = z.Structs[zgwx].Fields[zbei].UnmarshalMsg(bts)
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
					if nextMiss5zxzr != -1 {
						bts = nbs.PopAlwaysNil()
					}

				}
			}
		case "Imports":
			found4zulh[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zxbn uint32
				zxbn, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Imports) >= int(zxbn) {
					z.Imports = (z.Imports)[:zxbn]
				} else {
					z.Imports = make([]string, zxbn)
				}
				for zzfc := range z.Imports {
					z.Imports[zzfc], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss4zulh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Schema
var unmarshalMsgFieldOrder4zulh = []string{"SourcePath", "SourcePackage", "ZebraSchemaId", "Structs", "Imports"}

var unmarshalMsgFieldSkip4zulh = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zxzr = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip5zxzr = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zgwx := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zgwx].StructName) + 7 + msgp.ArrayHeaderSize
		for zbei := range z.Structs[zgwx].Fields {
			s += z.Structs[zgwx].Fields[zbei].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for zzfc := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zzfc])
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
	const maxFields6zsal = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zsal uint32
	totalEncodedFields6zsal, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zsal := totalEncodedFields6zsal
	missingFieldsLeft6zsal := maxFields6zsal - totalEncodedFields6zsal

	var nextMiss6zsal int32 = -1
	var found6zsal [maxFields6zsal]bool
	var curField6zsal string

doneWithStruct6zsal:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zsal > 0 || missingFieldsLeft6zsal > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zsal, missingFieldsLeft6zsal, msgp.ShowFound(found6zsal[:]), decodeMsgFieldOrder6zsal)
		if encodedFieldsLeft6zsal > 0 {
			encodedFieldsLeft6zsal--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zsal = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zsal < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zsal = 0
			}
			for nextMiss6zsal < maxFields6zsal && (found6zsal[nextMiss6zsal] || decodeMsgFieldSkip6zsal[nextMiss6zsal]) {
				nextMiss6zsal++
			}
			if nextMiss6zsal == maxFields6zsal {
				// filled all the empty fields!
				break doneWithStruct6zsal
			}
			missingFieldsLeft6zsal--
			curField6zsal = decodeMsgFieldOrder6zsal[nextMiss6zsal]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zsal)
		switch curField6zsal {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found6zsal[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fields":
			found6zsal[1] = true
			var zmkl uint32
			zmkl, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fields) >= int(zmkl) {
				z.Fields = (z.Fields)[:zmkl]
			} else {
				z.Fields = make([]Field, zmkl)
			}
			for zibr := range z.Fields {
				err = z.Fields[zibr].DecodeMsg(dc)
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
	if nextMiss6zsal != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Struct
var decodeMsgFieldOrder6zsal = []string{"StructName", "Fields"}

var decodeMsgFieldSkip6zsal = []bool{false, false}

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
	for zibr := range z.Fields {
		err = z.Fields[zibr].EncodeMsg(en)
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
	for zibr := range z.Fields {
		o, err = z.Fields[zibr].MarshalMsg(o)
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
	const maxFields7zyth = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zyth uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zyth, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zyth := totalEncodedFields7zyth
	missingFieldsLeft7zyth := maxFields7zyth - totalEncodedFields7zyth

	var nextMiss7zyth int32 = -1
	var found7zyth [maxFields7zyth]bool
	var curField7zyth string

doneWithStruct7zyth:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zyth > 0 || missingFieldsLeft7zyth > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zyth, missingFieldsLeft7zyth, msgp.ShowFound(found7zyth[:]), unmarshalMsgFieldOrder7zyth)
		if encodedFieldsLeft7zyth > 0 {
			encodedFieldsLeft7zyth--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zyth = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zyth < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zyth = 0
			}
			for nextMiss7zyth < maxFields7zyth && (found7zyth[nextMiss7zyth] || unmarshalMsgFieldSkip7zyth[nextMiss7zyth]) {
				nextMiss7zyth++
			}
			if nextMiss7zyth == maxFields7zyth {
				// filled all the empty fields!
				break doneWithStruct7zyth
			}
			missingFieldsLeft7zyth--
			curField7zyth = unmarshalMsgFieldOrder7zyth[nextMiss7zyth]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zyth)
		switch curField7zyth {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found7zyth[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fields":
			found7zyth[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zjjg uint32
				zjjg, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fields) >= int(zjjg) {
					z.Fields = (z.Fields)[:zjjg]
				} else {
					z.Fields = make([]Field, zjjg)
				}
				for zibr := range z.Fields {
					bts, err = z.Fields[zibr].UnmarshalMsg(bts)
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
	if nextMiss7zyth != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Struct
var unmarshalMsgFieldOrder7zyth = []string{"StructName", "Fields"}

var unmarshalMsgFieldSkip7zyth = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zibr := range z.Fields {
		s += z.Fields[zibr].Msgsize()
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
		var zgve uint64
		zgve, err = dc.ReadUint64()
		(*z) = Zkind(zgve)
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
		var zkvn uint64
		zkvn, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zkvn)
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
	const maxFields8zgci = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zgci uint32
	totalEncodedFields8zgci, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zgci := totalEncodedFields8zgci
	missingFieldsLeft8zgci := maxFields8zgci - totalEncodedFields8zgci

	var nextMiss8zgci int32 = -1
	var found8zgci [maxFields8zgci]bool
	var curField8zgci string

doneWithStruct8zgci:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zgci > 0 || missingFieldsLeft8zgci > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zgci, missingFieldsLeft8zgci, msgp.ShowFound(found8zgci[:]), decodeMsgFieldOrder8zgci)
		if encodedFieldsLeft8zgci > 0 {
			encodedFieldsLeft8zgci--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zgci = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zgci < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zgci = 0
			}
			for nextMiss8zgci < maxFields8zgci && (found8zgci[nextMiss8zgci] || decodeMsgFieldSkip8zgci[nextMiss8zgci]) {
				nextMiss8zgci++
			}
			if nextMiss8zgci == maxFields8zgci {
				// filled all the empty fields!
				break doneWithStruct8zgci
			}
			missingFieldsLeft8zgci--
			curField8zgci = decodeMsgFieldOrder8zgci[nextMiss8zgci]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zgci)
		switch curField8zgci {
		// -- templateDecodeMsg ends here --

		case "Kind":
			found8zgci[0] = true
			{
				var zhxj uint64
				zhxj, err = dc.ReadUint64()
				z.Kind = Zkind(zhxj)
			}
			if err != nil {
				panic(err)
			}
		case "Str":
			found8zgci[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Domain":
			found8zgci[2] = true
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
			found8zgci[3] = true
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
	if nextMiss8zgci != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of Ztype
var decodeMsgFieldOrder8zgci = []string{"Kind", "Str", "Domain", "Range"}

var decodeMsgFieldSkip8zgci = []bool{false, false, false, false}

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
	var empty_zvib [4]bool
	fieldsInUse_zliq := z.fieldsNotEmpty(empty_zvib[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zliq)
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
	if !empty_zvib[1] {
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

	if !empty_zvib[2] {
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

	if !empty_zvib[3] {
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
	const maxFields9zynd = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zynd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zynd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zynd := totalEncodedFields9zynd
	missingFieldsLeft9zynd := maxFields9zynd - totalEncodedFields9zynd

	var nextMiss9zynd int32 = -1
	var found9zynd [maxFields9zynd]bool
	var curField9zynd string

doneWithStruct9zynd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zynd > 0 || missingFieldsLeft9zynd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zynd, missingFieldsLeft9zynd, msgp.ShowFound(found9zynd[:]), unmarshalMsgFieldOrder9zynd)
		if encodedFieldsLeft9zynd > 0 {
			encodedFieldsLeft9zynd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zynd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zynd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zynd = 0
			}
			for nextMiss9zynd < maxFields9zynd && (found9zynd[nextMiss9zynd] || unmarshalMsgFieldSkip9zynd[nextMiss9zynd]) {
				nextMiss9zynd++
			}
			if nextMiss9zynd == maxFields9zynd {
				// filled all the empty fields!
				break doneWithStruct9zynd
			}
			missingFieldsLeft9zynd--
			curField9zynd = unmarshalMsgFieldOrder9zynd[nextMiss9zynd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zynd)
		switch curField9zynd {
		// -- templateUnmarshalMsg ends here --

		case "Kind":
			found9zynd[0] = true
			{
				var zeny uint64
				zeny, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Kind = Zkind(zeny)
			}
		case "Str":
			found9zynd[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Domain":
			found9zynd[2] = true
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
			found9zynd[3] = true
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
	if nextMiss9zynd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of Ztype
var unmarshalMsgFieldOrder9zynd = []string{"Kind", "Str", "Domain", "Range"}

var unmarshalMsgFieldSkip9zynd = []bool{false, false, false, false}

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
