package zebra

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *FieldT) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields0zqgi = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zqgi uint32
	totalEncodedFields0zqgi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zqgi := totalEncodedFields0zqgi
	missingFieldsLeft0zqgi := maxFields0zqgi - totalEncodedFields0zqgi

	var nextMiss0zqgi int32 = -1
	var found0zqgi [maxFields0zqgi]bool
	var curField0zqgi string

doneWithStruct0zqgi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zqgi > 0 || missingFieldsLeft0zqgi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zqgi, missingFieldsLeft0zqgi, msgp.ShowFound(found0zqgi[:]), decodeMsgFieldOrder0zqgi)
		if encodedFieldsLeft0zqgi > 0 {
			encodedFieldsLeft0zqgi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zqgi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zqgi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zqgi = 0
			}
			for nextMiss0zqgi < maxFields0zqgi && found0zqgi[nextMiss0zqgi] {
				nextMiss0zqgi++
			}
			if nextMiss0zqgi == maxFields0zqgi {
				// filled all the empty fields!
				break doneWithStruct0zqgi
			}
			missingFieldsLeft0zqgi--
			curField0zqgi = decodeMsgFieldOrder0zqgi[nextMiss0zqgi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zqgi)
		switch curField0zqgi {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zqgi[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zqgi[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zqgi[2] = true
			{
				var zlls uint8
				zlls, err = dc.ReadUint8()
				z.Typ = Zprimitive(zlls)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zqgi[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zqgi[4] = true
			var zdos uint32
			zdos, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zdos > 0 {
				z.Tag = make(map[string]string, zdos)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zdos > 0 {
				zdos--
				var zonr string
				var zvjt string
				zonr, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zvjt, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zonr] = zvjt
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zqgi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zqgi = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// fieldsNotEmpty supports omitempty tags
func (z *FieldT) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[4] = (len(z.Tag) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *FieldT) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_ztle [5]bool
	fieldsInUse_zpht := z.fieldsNotEmpty(empty_ztle[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpht)
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
	// write "Nam"
	err = en.Append(0xa3, 0x4e, 0x61, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nam)
	if err != nil {
		panic(err)
	}
	// write "Typ"
	err = en.Append(0xa3, 0x54, 0x79, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteUint8(uint8(z.Typ))
	if err != nil {
		panic(err)
	}
	// write "TypStr"
	err = en.Append(0xa6, 0x54, 0x79, 0x70, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TypStr)
	if err != nil {
		panic(err)
	}
	if !empty_ztle[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zonr, zvjt := range z.Tag {
			err = en.WriteString(zonr)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zvjt)
			if err != nil {
				panic(err)
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FieldT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [5]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Zid"
	o = append(o, 0xa3, 0x5a, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Zid)
	// string "Nam"
	o = append(o, 0xa3, 0x4e, 0x61, 0x6d)
	o = msgp.AppendString(o, z.Nam)
	// string "Typ"
	o = append(o, 0xa3, 0x54, 0x79, 0x70)
	o = msgp.AppendUint8(o, uint8(z.Typ))
	// string "TypStr"
	o = append(o, 0xa6, 0x54, 0x79, 0x70, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.TypStr)
	if !empty[4] {
		// string "Tag"
		o = append(o, 0xa3, 0x54, 0x61, 0x67)
		o = msgp.AppendMapHeader(o, uint32(len(z.Tag)))
		for zonr, zvjt := range z.Tag {
			o = msgp.AppendString(o, zonr)
			o = msgp.AppendString(o, zvjt)
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FieldT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *FieldT) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zrlr = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zrlr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zrlr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zrlr := totalEncodedFields1zrlr
	missingFieldsLeft1zrlr := maxFields1zrlr - totalEncodedFields1zrlr

	var nextMiss1zrlr int32 = -1
	var found1zrlr [maxFields1zrlr]bool
	var curField1zrlr string

doneWithStruct1zrlr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zrlr > 0 || missingFieldsLeft1zrlr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zrlr, missingFieldsLeft1zrlr, msgp.ShowFound(found1zrlr[:]), unmarshalMsgFieldOrder1zrlr)
		if encodedFieldsLeft1zrlr > 0 {
			encodedFieldsLeft1zrlr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zrlr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zrlr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zrlr = 0
			}
			for nextMiss1zrlr < maxFields1zrlr && found1zrlr[nextMiss1zrlr] {
				nextMiss1zrlr++
			}
			if nextMiss1zrlr == maxFields1zrlr {
				// filled all the empty fields!
				break doneWithStruct1zrlr
			}
			missingFieldsLeft1zrlr--
			curField1zrlr = unmarshalMsgFieldOrder1zrlr[nextMiss1zrlr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zrlr)
		switch curField1zrlr {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zrlr[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zrlr[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zrlr[2] = true
			{
				var zofg uint8
				zofg, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zofg)
			}
		case "TypStr":
			found1zrlr[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zrlr[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zmbd uint32
				zmbd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zmbd > 0 {
					z.Tag = make(map[string]string, zmbd)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zmbd > 0 {
					var zonr string
					var zvjt string
					zmbd--
					zonr, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zvjt, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zonr] = zvjt
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zrlr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zrlr = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zonr, zvjt := range z.Tag {
			_ = zvjt
			_ = zonr
			s += msgp.StringPrefixSize + len(zonr) + msgp.StringPrefixSize + len(zvjt)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *StructT) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zdtz = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zdtz uint32
	totalEncodedFields2zdtz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zdtz := totalEncodedFields2zdtz
	missingFieldsLeft2zdtz := maxFields2zdtz - totalEncodedFields2zdtz

	var nextMiss2zdtz int32 = -1
	var found2zdtz [maxFields2zdtz]bool
	var curField2zdtz string

doneWithStruct2zdtz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zdtz > 0 || missingFieldsLeft2zdtz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zdtz, missingFieldsLeft2zdtz, msgp.ShowFound(found2zdtz[:]), decodeMsgFieldOrder2zdtz)
		if encodedFieldsLeft2zdtz > 0 {
			encodedFieldsLeft2zdtz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zdtz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zdtz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zdtz = 0
			}
			for nextMiss2zdtz < maxFields2zdtz && found2zdtz[nextMiss2zdtz] {
				nextMiss2zdtz++
			}
			if nextMiss2zdtz == maxFields2zdtz {
				// filled all the empty fields!
				break doneWithStruct2zdtz
			}
			missingFieldsLeft2zdtz--
			curField2zdtz = decodeMsgFieldOrder2zdtz[nextMiss2zdtz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zdtz)
		switch curField2zdtz {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zdtz[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zdtz[1] = true
			var zphu uint32
			zphu, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zphu) {
				z.Fld = (z.Fld)[:zphu]
			} else {
				z.Fld = make([]FieldT, zphu)
			}
			for zqku := range z.Fld {
				err = z.Fld[zqku].DecodeMsg(dc)
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
	if nextMiss2zdtz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zdtz = []string{"StructName", "Fld"}

// fieldsNotEmpty supports omitempty tags
func (z *StructT) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *StructT) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "Fld"
	err = en.Append(0xa3, 0x46, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fld)))
	if err != nil {
		panic(err)
	}
	for zqku := range z.Fld {
		err = z.Fld[zqku].EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StructT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "StructName"
	o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.StructName)
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fld)))
	for zqku := range z.Fld {
		o, err = z.Fld[zqku].MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StructT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *StructT) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zheb = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zheb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zheb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zheb := totalEncodedFields3zheb
	missingFieldsLeft3zheb := maxFields3zheb - totalEncodedFields3zheb

	var nextMiss3zheb int32 = -1
	var found3zheb [maxFields3zheb]bool
	var curField3zheb string

doneWithStruct3zheb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zheb > 0 || missingFieldsLeft3zheb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zheb, missingFieldsLeft3zheb, msgp.ShowFound(found3zheb[:]), unmarshalMsgFieldOrder3zheb)
		if encodedFieldsLeft3zheb > 0 {
			encodedFieldsLeft3zheb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zheb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zheb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zheb = 0
			}
			for nextMiss3zheb < maxFields3zheb && found3zheb[nextMiss3zheb] {
				nextMiss3zheb++
			}
			if nextMiss3zheb == maxFields3zheb {
				// filled all the empty fields!
				break doneWithStruct3zheb
			}
			missingFieldsLeft3zheb--
			curField3zheb = unmarshalMsgFieldOrder3zheb[nextMiss3zheb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zheb)
		switch curField3zheb {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zheb[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zheb[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zbur uint32
				zbur, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zbur) {
					z.Fld = (z.Fld)[:zbur]
				} else {
					z.Fld = make([]FieldT, zbur)
				}
				for zqku := range z.Fld {
					bts, err = z.Fld[zqku].UnmarshalMsg(bts)
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
	if nextMiss3zheb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zheb = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zqku := range z.Fld {
		s += z.Fld[zqku].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *ZebraSchema) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields4zpun = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zpun uint32
	totalEncodedFields4zpun, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zpun := totalEncodedFields4zpun
	missingFieldsLeft4zpun := maxFields4zpun - totalEncodedFields4zpun

	var nextMiss4zpun int32 = -1
	var found4zpun [maxFields4zpun]bool
	var curField4zpun string

doneWithStruct4zpun:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zpun > 0 || missingFieldsLeft4zpun > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zpun, missingFieldsLeft4zpun, msgp.ShowFound(found4zpun[:]), decodeMsgFieldOrder4zpun)
		if encodedFieldsLeft4zpun > 0 {
			encodedFieldsLeft4zpun--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zpun = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zpun < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zpun = 0
			}
			for nextMiss4zpun < maxFields4zpun && found4zpun[nextMiss4zpun] {
				nextMiss4zpun++
			}
			if nextMiss4zpun == maxFields4zpun {
				// filled all the empty fields!
				break doneWithStruct4zpun
			}
			missingFieldsLeft4zpun--
			curField4zpun = decodeMsgFieldOrder4zpun[nextMiss4zpun]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zpun)
		switch curField4zpun {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zpun[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zpun[1] = true
			var zgiz uint32
			zgiz, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zgiz) {
				z.Structs = (z.Structs)[:zgiz]
			} else {
				z.Structs = make([]StructT, zgiz)
			}
			for zpet := range z.Structs {
				const maxFields5zuqe = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zuqe uint32
				totalEncodedFields5zuqe, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zuqe := totalEncodedFields5zuqe
				missingFieldsLeft5zuqe := maxFields5zuqe - totalEncodedFields5zuqe

				var nextMiss5zuqe int32 = -1
				var found5zuqe [maxFields5zuqe]bool
				var curField5zuqe string

			doneWithStruct5zuqe:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zuqe > 0 || missingFieldsLeft5zuqe > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zuqe, missingFieldsLeft5zuqe, msgp.ShowFound(found5zuqe[:]), decodeMsgFieldOrder5zuqe)
					if encodedFieldsLeft5zuqe > 0 {
						encodedFieldsLeft5zuqe--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zuqe = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zuqe < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zuqe = 0
						}
						for nextMiss5zuqe < maxFields5zuqe && found5zuqe[nextMiss5zuqe] {
							nextMiss5zuqe++
						}
						if nextMiss5zuqe == maxFields5zuqe {
							// filled all the empty fields!
							break doneWithStruct5zuqe
						}
						missingFieldsLeft5zuqe--
						curField5zuqe = decodeMsgFieldOrder5zuqe[nextMiss5zuqe]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zuqe)
					switch curField5zuqe {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zuqe[0] = true
						z.Structs[zpet].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zuqe[1] = true
						var zbvu uint32
						zbvu, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zpet].Fld) >= int(zbvu) {
							z.Structs[zpet].Fld = (z.Structs[zpet].Fld)[:zbvu]
						} else {
							z.Structs[zpet].Fld = make([]FieldT, zbvu)
						}
						for zkfc := range z.Structs[zpet].Fld {
							err = z.Structs[zpet].Fld[zkfc].DecodeMsg(dc)
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
				if nextMiss5zuqe != -1 {
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
	if nextMiss4zpun != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zpun = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zuqe = []string{"StructName", "Fld"}

// fieldsNotEmpty supports omitempty tags
func (z *ZebraSchema) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *ZebraSchema) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "PkgPath"
	err = en.Append(0x82, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.PkgPath)
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
	for zpet := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zpet].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zpet].Fld)))
		if err != nil {
			panic(err)
		}
		for zkfc := range z.Structs[zpet].Fld {
			err = z.Structs[zpet].Fld[zkfc].EncodeMsg(en)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ZebraSchema) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "PkgPath"
	o = append(o, 0x82, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.PkgPath)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zpet := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zpet].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zpet].Fld)))
		for zkfc := range z.Structs[zpet].Fld {
			o, err = z.Structs[zpet].Fld[zkfc].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ZebraSchema) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *ZebraSchema) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields6zovb = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zovb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zovb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zovb := totalEncodedFields6zovb
	missingFieldsLeft6zovb := maxFields6zovb - totalEncodedFields6zovb

	var nextMiss6zovb int32 = -1
	var found6zovb [maxFields6zovb]bool
	var curField6zovb string

doneWithStruct6zovb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zovb > 0 || missingFieldsLeft6zovb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zovb, missingFieldsLeft6zovb, msgp.ShowFound(found6zovb[:]), unmarshalMsgFieldOrder6zovb)
		if encodedFieldsLeft6zovb > 0 {
			encodedFieldsLeft6zovb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zovb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zovb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zovb = 0
			}
			for nextMiss6zovb < maxFields6zovb && found6zovb[nextMiss6zovb] {
				nextMiss6zovb++
			}
			if nextMiss6zovb == maxFields6zovb {
				// filled all the empty fields!
				break doneWithStruct6zovb
			}
			missingFieldsLeft6zovb--
			curField6zovb = unmarshalMsgFieldOrder6zovb[nextMiss6zovb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zovb)
		switch curField6zovb {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zovb[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zovb[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zcrk uint32
				zcrk, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zcrk) {
					z.Structs = (z.Structs)[:zcrk]
				} else {
					z.Structs = make([]StructT, zcrk)
				}
				for zpet := range z.Structs {
					const maxFields7zjbg = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zjbg uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zjbg, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zjbg := totalEncodedFields7zjbg
					missingFieldsLeft7zjbg := maxFields7zjbg - totalEncodedFields7zjbg

					var nextMiss7zjbg int32 = -1
					var found7zjbg [maxFields7zjbg]bool
					var curField7zjbg string

				doneWithStruct7zjbg:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zjbg > 0 || missingFieldsLeft7zjbg > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zjbg, missingFieldsLeft7zjbg, msgp.ShowFound(found7zjbg[:]), unmarshalMsgFieldOrder7zjbg)
						if encodedFieldsLeft7zjbg > 0 {
							encodedFieldsLeft7zjbg--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zjbg = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zjbg < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zjbg = 0
							}
							for nextMiss7zjbg < maxFields7zjbg && found7zjbg[nextMiss7zjbg] {
								nextMiss7zjbg++
							}
							if nextMiss7zjbg == maxFields7zjbg {
								// filled all the empty fields!
								break doneWithStruct7zjbg
							}
							missingFieldsLeft7zjbg--
							curField7zjbg = unmarshalMsgFieldOrder7zjbg[nextMiss7zjbg]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zjbg)
						switch curField7zjbg {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zjbg[0] = true
							z.Structs[zpet].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zjbg[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zpet].Fld) = (z.Structs[zpet].Fld)[:0]
							} else {

								var zeuu uint32
								zeuu, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zpet].Fld) >= int(zeuu) {
									z.Structs[zpet].Fld = (z.Structs[zpet].Fld)[:zeuu]
								} else {
									z.Structs[zpet].Fld = make([]FieldT, zeuu)
								}
								for zkfc := range z.Structs[zpet].Fld {
									bts, err = z.Structs[zpet].Fld[zkfc].UnmarshalMsg(bts)
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
					if nextMiss7zjbg != -1 {
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
	if nextMiss6zovb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zovb = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zjbg = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zpet := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zpet].StructName) + 4 + msgp.ArrayHeaderSize
		for zkfc := range z.Structs[zpet].Fld {
			s += z.Structs[zpet].Fld[zkfc].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Zprimitive) DecodeMsg(dc *msgp.Reader) (err error) {
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
		var zkxe uint8
		zkxe, err = dc.ReadUint8()
		(*z) = Zprimitive(zkxe)
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
func (z Zprimitive) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zprimitive) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zprimitive) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Zprimitive) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zbvk uint8
		zbvk, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zbvk)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zprimitive) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}
