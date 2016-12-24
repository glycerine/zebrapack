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
	const maxFields0zaqo = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zaqo uint32
	totalEncodedFields0zaqo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zaqo := totalEncodedFields0zaqo
	missingFieldsLeft0zaqo := maxFields0zaqo - totalEncodedFields0zaqo

	var nextMiss0zaqo int32 = -1
	var found0zaqo [maxFields0zaqo]bool
	var curField0zaqo string

doneWithStruct0zaqo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaqo > 0 || missingFieldsLeft0zaqo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaqo, missingFieldsLeft0zaqo, msgp.ShowFound(found0zaqo[:]), decodeMsgFieldOrder0zaqo)
		if encodedFieldsLeft0zaqo > 0 {
			encodedFieldsLeft0zaqo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zaqo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zaqo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zaqo = 0
			}
			for nextMiss0zaqo < maxFields0zaqo && found0zaqo[nextMiss0zaqo] {
				nextMiss0zaqo++
			}
			if nextMiss0zaqo == maxFields0zaqo {
				// filled all the empty fields!
				break doneWithStruct0zaqo
			}
			missingFieldsLeft0zaqo--
			curField0zaqo = decodeMsgFieldOrder0zaqo[nextMiss0zaqo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaqo)
		switch curField0zaqo {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zaqo[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zaqo[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zaqo[2] = true
			{
				var zdzi uint8
				zdzi, err = dc.ReadUint8()
				z.Typ = Zprimitive(zdzi)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zaqo[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zaqo[4] = true
			var zonn uint32
			zonn, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zonn > 0 {
				z.Tag = make(map[string]string, zonn)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zonn > 0 {
				zonn--
				var zcuf string
				var zawm string
				zcuf, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zawm, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zcuf] = zawm
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zaqo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zaqo = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zmgb [5]bool
	fieldsInUse_zqch := z.fieldsNotEmpty(empty_zmgb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqch)
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
	if !empty_zmgb[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zcuf, zawm := range z.Tag {
			err = en.WriteString(zcuf)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zawm)
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
		for zcuf, zawm := range z.Tag {
			o = msgp.AppendString(o, zcuf)
			o = msgp.AppendString(o, zawm)
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
	const maxFields1zxke = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zxke uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zxke, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zxke := totalEncodedFields1zxke
	missingFieldsLeft1zxke := maxFields1zxke - totalEncodedFields1zxke

	var nextMiss1zxke int32 = -1
	var found1zxke [maxFields1zxke]bool
	var curField1zxke string

doneWithStruct1zxke:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zxke > 0 || missingFieldsLeft1zxke > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zxke, missingFieldsLeft1zxke, msgp.ShowFound(found1zxke[:]), unmarshalMsgFieldOrder1zxke)
		if encodedFieldsLeft1zxke > 0 {
			encodedFieldsLeft1zxke--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zxke = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zxke < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zxke = 0
			}
			for nextMiss1zxke < maxFields1zxke && found1zxke[nextMiss1zxke] {
				nextMiss1zxke++
			}
			if nextMiss1zxke == maxFields1zxke {
				// filled all the empty fields!
				break doneWithStruct1zxke
			}
			missingFieldsLeft1zxke--
			curField1zxke = unmarshalMsgFieldOrder1zxke[nextMiss1zxke]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zxke)
		switch curField1zxke {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zxke[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zxke[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zxke[2] = true
			{
				var zdkc uint8
				zdkc, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zdkc)
			}
		case "TypStr":
			found1zxke[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zxke[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zshm uint32
				zshm, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zshm > 0 {
					z.Tag = make(map[string]string, zshm)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zshm > 0 {
					var zcuf string
					var zawm string
					zshm--
					zcuf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zawm, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zcuf] = zawm
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zxke != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zxke = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zcuf, zawm := range z.Tag {
			_ = zawm
			_ = zcuf
			s += msgp.StringPrefixSize + len(zcuf) + msgp.StringPrefixSize + len(zawm)
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
	const maxFields2zvta = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zvta uint32
	totalEncodedFields2zvta, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zvta := totalEncodedFields2zvta
	missingFieldsLeft2zvta := maxFields2zvta - totalEncodedFields2zvta

	var nextMiss2zvta int32 = -1
	var found2zvta [maxFields2zvta]bool
	var curField2zvta string

doneWithStruct2zvta:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zvta > 0 || missingFieldsLeft2zvta > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zvta, missingFieldsLeft2zvta, msgp.ShowFound(found2zvta[:]), decodeMsgFieldOrder2zvta)
		if encodedFieldsLeft2zvta > 0 {
			encodedFieldsLeft2zvta--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zvta = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zvta < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zvta = 0
			}
			for nextMiss2zvta < maxFields2zvta && found2zvta[nextMiss2zvta] {
				nextMiss2zvta++
			}
			if nextMiss2zvta == maxFields2zvta {
				// filled all the empty fields!
				break doneWithStruct2zvta
			}
			missingFieldsLeft2zvta--
			curField2zvta = decodeMsgFieldOrder2zvta[nextMiss2zvta]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zvta)
		switch curField2zvta {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zvta[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zvta[1] = true
			var znkl uint32
			znkl, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(znkl) {
				z.Fld = (z.Fld)[:znkl]
			} else {
				z.Fld = make([]FieldT, znkl)
			}
			for zaoe := range z.Fld {
				err = z.Fld[zaoe].DecodeMsg(dc)
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
	if nextMiss2zvta != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zvta = []string{"StructName", "Fld"}

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
	for zaoe := range z.Fld {
		err = z.Fld[zaoe].EncodeMsg(en)
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
	for zaoe := range z.Fld {
		o, err = z.Fld[zaoe].MarshalMsg(o)
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
	const maxFields3zenu = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zenu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zenu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zenu := totalEncodedFields3zenu
	missingFieldsLeft3zenu := maxFields3zenu - totalEncodedFields3zenu

	var nextMiss3zenu int32 = -1
	var found3zenu [maxFields3zenu]bool
	var curField3zenu string

doneWithStruct3zenu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zenu > 0 || missingFieldsLeft3zenu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zenu, missingFieldsLeft3zenu, msgp.ShowFound(found3zenu[:]), unmarshalMsgFieldOrder3zenu)
		if encodedFieldsLeft3zenu > 0 {
			encodedFieldsLeft3zenu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zenu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zenu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zenu = 0
			}
			for nextMiss3zenu < maxFields3zenu && found3zenu[nextMiss3zenu] {
				nextMiss3zenu++
			}
			if nextMiss3zenu == maxFields3zenu {
				// filled all the empty fields!
				break doneWithStruct3zenu
			}
			missingFieldsLeft3zenu--
			curField3zenu = unmarshalMsgFieldOrder3zenu[nextMiss3zenu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zenu)
		switch curField3zenu {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zenu[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zenu[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zjzf uint32
				zjzf, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zjzf) {
					z.Fld = (z.Fld)[:zjzf]
				} else {
					z.Fld = make([]FieldT, zjzf)
				}
				for zaoe := range z.Fld {
					bts, err = z.Fld[zaoe].UnmarshalMsg(bts)
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
	if nextMiss3zenu != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zenu = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zaoe := range z.Fld {
		s += z.Fld[zaoe].Msgsize()
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
	const maxFields4zesw = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zesw uint32
	totalEncodedFields4zesw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zesw := totalEncodedFields4zesw
	missingFieldsLeft4zesw := maxFields4zesw - totalEncodedFields4zesw

	var nextMiss4zesw int32 = -1
	var found4zesw [maxFields4zesw]bool
	var curField4zesw string

doneWithStruct4zesw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zesw > 0 || missingFieldsLeft4zesw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zesw, missingFieldsLeft4zesw, msgp.ShowFound(found4zesw[:]), decodeMsgFieldOrder4zesw)
		if encodedFieldsLeft4zesw > 0 {
			encodedFieldsLeft4zesw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zesw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zesw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zesw = 0
			}
			for nextMiss4zesw < maxFields4zesw && found4zesw[nextMiss4zesw] {
				nextMiss4zesw++
			}
			if nextMiss4zesw == maxFields4zesw {
				// filled all the empty fields!
				break doneWithStruct4zesw
			}
			missingFieldsLeft4zesw--
			curField4zesw = decodeMsgFieldOrder4zesw[nextMiss4zesw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zesw)
		switch curField4zesw {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zesw[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zesw[1] = true
			var zsmb uint32
			zsmb, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zsmb) {
				z.Structs = (z.Structs)[:zsmb]
			} else {
				z.Structs = make([]StructT, zsmb)
			}
			for zyrx := range z.Structs {
				const maxFields5zwld = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zwld uint32
				totalEncodedFields5zwld, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zwld := totalEncodedFields5zwld
				missingFieldsLeft5zwld := maxFields5zwld - totalEncodedFields5zwld

				var nextMiss5zwld int32 = -1
				var found5zwld [maxFields5zwld]bool
				var curField5zwld string

			doneWithStruct5zwld:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zwld > 0 || missingFieldsLeft5zwld > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zwld, missingFieldsLeft5zwld, msgp.ShowFound(found5zwld[:]), decodeMsgFieldOrder5zwld)
					if encodedFieldsLeft5zwld > 0 {
						encodedFieldsLeft5zwld--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zwld = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zwld < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zwld = 0
						}
						for nextMiss5zwld < maxFields5zwld && found5zwld[nextMiss5zwld] {
							nextMiss5zwld++
						}
						if nextMiss5zwld == maxFields5zwld {
							// filled all the empty fields!
							break doneWithStruct5zwld
						}
						missingFieldsLeft5zwld--
						curField5zwld = decodeMsgFieldOrder5zwld[nextMiss5zwld]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zwld)
					switch curField5zwld {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zwld[0] = true
						z.Structs[zyrx].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zwld[1] = true
						var zlgw uint32
						zlgw, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zyrx].Fld) >= int(zlgw) {
							z.Structs[zyrx].Fld = (z.Structs[zyrx].Fld)[:zlgw]
						} else {
							z.Structs[zyrx].Fld = make([]FieldT, zlgw)
						}
						for zusn := range z.Structs[zyrx].Fld {
							err = z.Structs[zyrx].Fld[zusn].DecodeMsg(dc)
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
				if nextMiss5zwld != -1 {
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
	if nextMiss4zesw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zesw = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zwld = []string{"StructName", "Fld"}

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
	for zyrx := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zyrx].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zyrx].Fld)))
		if err != nil {
			panic(err)
		}
		for zusn := range z.Structs[zyrx].Fld {
			err = z.Structs[zyrx].Fld[zusn].EncodeMsg(en)
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
	for zyrx := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zyrx].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zyrx].Fld)))
		for zusn := range z.Structs[zyrx].Fld {
			o, err = z.Structs[zyrx].Fld[zusn].MarshalMsg(o)
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
	const maxFields6zemv = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zemv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zemv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zemv := totalEncodedFields6zemv
	missingFieldsLeft6zemv := maxFields6zemv - totalEncodedFields6zemv

	var nextMiss6zemv int32 = -1
	var found6zemv [maxFields6zemv]bool
	var curField6zemv string

doneWithStruct6zemv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zemv > 0 || missingFieldsLeft6zemv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zemv, missingFieldsLeft6zemv, msgp.ShowFound(found6zemv[:]), unmarshalMsgFieldOrder6zemv)
		if encodedFieldsLeft6zemv > 0 {
			encodedFieldsLeft6zemv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zemv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zemv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zemv = 0
			}
			for nextMiss6zemv < maxFields6zemv && found6zemv[nextMiss6zemv] {
				nextMiss6zemv++
			}
			if nextMiss6zemv == maxFields6zemv {
				// filled all the empty fields!
				break doneWithStruct6zemv
			}
			missingFieldsLeft6zemv--
			curField6zemv = unmarshalMsgFieldOrder6zemv[nextMiss6zemv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zemv)
		switch curField6zemv {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zemv[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zemv[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zeuy uint32
				zeuy, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zeuy) {
					z.Structs = (z.Structs)[:zeuy]
				} else {
					z.Structs = make([]StructT, zeuy)
				}
				for zyrx := range z.Structs {
					const maxFields7zvla = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zvla uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zvla, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zvla := totalEncodedFields7zvla
					missingFieldsLeft7zvla := maxFields7zvla - totalEncodedFields7zvla

					var nextMiss7zvla int32 = -1
					var found7zvla [maxFields7zvla]bool
					var curField7zvla string

				doneWithStruct7zvla:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zvla > 0 || missingFieldsLeft7zvla > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zvla, missingFieldsLeft7zvla, msgp.ShowFound(found7zvla[:]), unmarshalMsgFieldOrder7zvla)
						if encodedFieldsLeft7zvla > 0 {
							encodedFieldsLeft7zvla--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zvla = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zvla < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zvla = 0
							}
							for nextMiss7zvla < maxFields7zvla && found7zvla[nextMiss7zvla] {
								nextMiss7zvla++
							}
							if nextMiss7zvla == maxFields7zvla {
								// filled all the empty fields!
								break doneWithStruct7zvla
							}
							missingFieldsLeft7zvla--
							curField7zvla = unmarshalMsgFieldOrder7zvla[nextMiss7zvla]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zvla)
						switch curField7zvla {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zvla[0] = true
							z.Structs[zyrx].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zvla[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zyrx].Fld) = (z.Structs[zyrx].Fld)[:0]
							} else {

								var zhqg uint32
								zhqg, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zyrx].Fld) >= int(zhqg) {
									z.Structs[zyrx].Fld = (z.Structs[zyrx].Fld)[:zhqg]
								} else {
									z.Structs[zyrx].Fld = make([]FieldT, zhqg)
								}
								for zusn := range z.Structs[zyrx].Fld {
									bts, err = z.Structs[zyrx].Fld[zusn].UnmarshalMsg(bts)
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
					if nextMiss7zvla != -1 {
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
	if nextMiss6zemv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zemv = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zvla = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zyrx := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zyrx].StructName) + 4 + msgp.ArrayHeaderSize
		for zusn := range z.Structs[zyrx].Fld {
			s += z.Structs[zyrx].Fld[zusn].Msgsize()
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
		var zoxa uint8
		zoxa, err = dc.ReadUint8()
		(*z) = Zprimitive(zoxa)
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
		var zjny uint8
		zjny, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zjny)
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
