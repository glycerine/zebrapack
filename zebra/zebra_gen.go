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
	const maxFields0zanj = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zanj uint32
	totalEncodedFields0zanj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zanj := totalEncodedFields0zanj
	missingFieldsLeft0zanj := maxFields0zanj - totalEncodedFields0zanj

	var nextMiss0zanj int32 = -1
	var found0zanj [maxFields0zanj]bool
	var curField0zanj string

doneWithStruct0zanj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zanj > 0 || missingFieldsLeft0zanj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zanj, missingFieldsLeft0zanj, msgp.ShowFound(found0zanj[:]), decodeMsgFieldOrder0zanj)
		if encodedFieldsLeft0zanj > 0 {
			encodedFieldsLeft0zanj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zanj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zanj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zanj = 0
			}
			for nextMiss0zanj < maxFields0zanj && found0zanj[nextMiss0zanj] {
				nextMiss0zanj++
			}
			if nextMiss0zanj == maxFields0zanj {
				// filled all the empty fields!
				break doneWithStruct0zanj
			}
			missingFieldsLeft0zanj--
			curField0zanj = decodeMsgFieldOrder0zanj[nextMiss0zanj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zanj)
		switch curField0zanj {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zanj[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zanj[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zanj[2] = true
			{
				var zcra uint8
				zcra, err = dc.ReadUint8()
				z.Typ = Zprimitive(zcra)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zanj[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zanj[4] = true
			var zjnd uint32
			zjnd, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zjnd > 0 {
				z.Tag = make(map[string]string, zjnd)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zjnd > 0 {
				zjnd--
				var zjci string
				var zsgo string
				zjci, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zsgo, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zjci] = zsgo
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zanj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zanj = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zqgl [5]bool
	fieldsInUse_zxms := z.fieldsNotEmpty(empty_zqgl[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxms)
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
	if !empty_zqgl[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zjci, zsgo := range z.Tag {
			err = en.WriteString(zjci)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zsgo)
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
		for zjci, zsgo := range z.Tag {
			o = msgp.AppendString(o, zjci)
			o = msgp.AppendString(o, zsgo)
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
	const maxFields1zdbc = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zdbc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zdbc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zdbc := totalEncodedFields1zdbc
	missingFieldsLeft1zdbc := maxFields1zdbc - totalEncodedFields1zdbc

	var nextMiss1zdbc int32 = -1
	var found1zdbc [maxFields1zdbc]bool
	var curField1zdbc string

doneWithStruct1zdbc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zdbc > 0 || missingFieldsLeft1zdbc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zdbc, missingFieldsLeft1zdbc, msgp.ShowFound(found1zdbc[:]), unmarshalMsgFieldOrder1zdbc)
		if encodedFieldsLeft1zdbc > 0 {
			encodedFieldsLeft1zdbc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zdbc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zdbc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zdbc = 0
			}
			for nextMiss1zdbc < maxFields1zdbc && found1zdbc[nextMiss1zdbc] {
				nextMiss1zdbc++
			}
			if nextMiss1zdbc == maxFields1zdbc {
				// filled all the empty fields!
				break doneWithStruct1zdbc
			}
			missingFieldsLeft1zdbc--
			curField1zdbc = unmarshalMsgFieldOrder1zdbc[nextMiss1zdbc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zdbc)
		switch curField1zdbc {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zdbc[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zdbc[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zdbc[2] = true
			{
				var zqfq uint8
				zqfq, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zqfq)
			}
		case "TypStr":
			found1zdbc[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zdbc[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zmoq uint32
				zmoq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zmoq > 0 {
					z.Tag = make(map[string]string, zmoq)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zmoq > 0 {
					var zjci string
					var zsgo string
					zmoq--
					zjci, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zsgo, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zjci] = zsgo
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zdbc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zdbc = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zjci, zsgo := range z.Tag {
			_ = zsgo
			_ = zjci
			s += msgp.StringPrefixSize + len(zjci) + msgp.StringPrefixSize + len(zsgo)
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
	const maxFields2zfqu = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zfqu uint32
	totalEncodedFields2zfqu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zfqu := totalEncodedFields2zfqu
	missingFieldsLeft2zfqu := maxFields2zfqu - totalEncodedFields2zfqu

	var nextMiss2zfqu int32 = -1
	var found2zfqu [maxFields2zfqu]bool
	var curField2zfqu string

doneWithStruct2zfqu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zfqu > 0 || missingFieldsLeft2zfqu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zfqu, missingFieldsLeft2zfqu, msgp.ShowFound(found2zfqu[:]), decodeMsgFieldOrder2zfqu)
		if encodedFieldsLeft2zfqu > 0 {
			encodedFieldsLeft2zfqu--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zfqu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zfqu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zfqu = 0
			}
			for nextMiss2zfqu < maxFields2zfqu && found2zfqu[nextMiss2zfqu] {
				nextMiss2zfqu++
			}
			if nextMiss2zfqu == maxFields2zfqu {
				// filled all the empty fields!
				break doneWithStruct2zfqu
			}
			missingFieldsLeft2zfqu--
			curField2zfqu = decodeMsgFieldOrder2zfqu[nextMiss2zfqu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zfqu)
		switch curField2zfqu {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zfqu[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zfqu[1] = true
			var zorp uint32
			zorp, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zorp) {
				z.Fld = (z.Fld)[:zorp]
			} else {
				z.Fld = make([]FieldT, zorp)
			}
			for znvd := range z.Fld {
				err = z.Fld[znvd].DecodeMsg(dc)
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
	if nextMiss2zfqu != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zfqu = []string{"StructName", "Fld"}

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
	for znvd := range z.Fld {
		err = z.Fld[znvd].EncodeMsg(en)
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
	for znvd := range z.Fld {
		o, err = z.Fld[znvd].MarshalMsg(o)
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
	const maxFields3zhcg = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zhcg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zhcg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zhcg := totalEncodedFields3zhcg
	missingFieldsLeft3zhcg := maxFields3zhcg - totalEncodedFields3zhcg

	var nextMiss3zhcg int32 = -1
	var found3zhcg [maxFields3zhcg]bool
	var curField3zhcg string

doneWithStruct3zhcg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zhcg > 0 || missingFieldsLeft3zhcg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zhcg, missingFieldsLeft3zhcg, msgp.ShowFound(found3zhcg[:]), unmarshalMsgFieldOrder3zhcg)
		if encodedFieldsLeft3zhcg > 0 {
			encodedFieldsLeft3zhcg--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zhcg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zhcg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zhcg = 0
			}
			for nextMiss3zhcg < maxFields3zhcg && found3zhcg[nextMiss3zhcg] {
				nextMiss3zhcg++
			}
			if nextMiss3zhcg == maxFields3zhcg {
				// filled all the empty fields!
				break doneWithStruct3zhcg
			}
			missingFieldsLeft3zhcg--
			curField3zhcg = unmarshalMsgFieldOrder3zhcg[nextMiss3zhcg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zhcg)
		switch curField3zhcg {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zhcg[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zhcg[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zdro uint32
				zdro, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zdro) {
					z.Fld = (z.Fld)[:zdro]
				} else {
					z.Fld = make([]FieldT, zdro)
				}
				for znvd := range z.Fld {
					bts, err = z.Fld[znvd].UnmarshalMsg(bts)
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
	if nextMiss3zhcg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zhcg = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for znvd := range z.Fld {
		s += z.Fld[znvd].Msgsize()
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
	const maxFields4zddj = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zddj uint32
	totalEncodedFields4zddj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zddj := totalEncodedFields4zddj
	missingFieldsLeft4zddj := maxFields4zddj - totalEncodedFields4zddj

	var nextMiss4zddj int32 = -1
	var found4zddj [maxFields4zddj]bool
	var curField4zddj string

doneWithStruct4zddj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zddj > 0 || missingFieldsLeft4zddj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zddj, missingFieldsLeft4zddj, msgp.ShowFound(found4zddj[:]), decodeMsgFieldOrder4zddj)
		if encodedFieldsLeft4zddj > 0 {
			encodedFieldsLeft4zddj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zddj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zddj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zddj = 0
			}
			for nextMiss4zddj < maxFields4zddj && found4zddj[nextMiss4zddj] {
				nextMiss4zddj++
			}
			if nextMiss4zddj == maxFields4zddj {
				// filled all the empty fields!
				break doneWithStruct4zddj
			}
			missingFieldsLeft4zddj--
			curField4zddj = decodeMsgFieldOrder4zddj[nextMiss4zddj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zddj)
		switch curField4zddj {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zddj[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zddj[1] = true
			var zezr uint32
			zezr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zezr) {
				z.Structs = (z.Structs)[:zezr]
			} else {
				z.Structs = make([]StructT, zezr)
			}
			for zgli := range z.Structs {
				const maxFields5zzwl = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zzwl uint32
				totalEncodedFields5zzwl, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zzwl := totalEncodedFields5zzwl
				missingFieldsLeft5zzwl := maxFields5zzwl - totalEncodedFields5zzwl

				var nextMiss5zzwl int32 = -1
				var found5zzwl [maxFields5zzwl]bool
				var curField5zzwl string

			doneWithStruct5zzwl:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zzwl > 0 || missingFieldsLeft5zzwl > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zzwl, missingFieldsLeft5zzwl, msgp.ShowFound(found5zzwl[:]), decodeMsgFieldOrder5zzwl)
					if encodedFieldsLeft5zzwl > 0 {
						encodedFieldsLeft5zzwl--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zzwl = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zzwl < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zzwl = 0
						}
						for nextMiss5zzwl < maxFields5zzwl && found5zzwl[nextMiss5zzwl] {
							nextMiss5zzwl++
						}
						if nextMiss5zzwl == maxFields5zzwl {
							// filled all the empty fields!
							break doneWithStruct5zzwl
						}
						missingFieldsLeft5zzwl--
						curField5zzwl = decodeMsgFieldOrder5zzwl[nextMiss5zzwl]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zzwl)
					switch curField5zzwl {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zzwl[0] = true
						z.Structs[zgli].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zzwl[1] = true
						var zkfj uint32
						zkfj, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zgli].Fld) >= int(zkfj) {
							z.Structs[zgli].Fld = (z.Structs[zgli].Fld)[:zkfj]
						} else {
							z.Structs[zgli].Fld = make([]FieldT, zkfj)
						}
						for zbrc := range z.Structs[zgli].Fld {
							err = z.Structs[zgli].Fld[zbrc].DecodeMsg(dc)
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
				if nextMiss5zzwl != -1 {
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
	if nextMiss4zddj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zddj = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zzwl = []string{"StructName", "Fld"}

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
	for zgli := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zgli].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zgli].Fld)))
		if err != nil {
			panic(err)
		}
		for zbrc := range z.Structs[zgli].Fld {
			err = z.Structs[zgli].Fld[zbrc].EncodeMsg(en)
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
	for zgli := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zgli].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zgli].Fld)))
		for zbrc := range z.Structs[zgli].Fld {
			o, err = z.Structs[zgli].Fld[zbrc].MarshalMsg(o)
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
	const maxFields6zyol = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zyol uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zyol, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zyol := totalEncodedFields6zyol
	missingFieldsLeft6zyol := maxFields6zyol - totalEncodedFields6zyol

	var nextMiss6zyol int32 = -1
	var found6zyol [maxFields6zyol]bool
	var curField6zyol string

doneWithStruct6zyol:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zyol > 0 || missingFieldsLeft6zyol > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zyol, missingFieldsLeft6zyol, msgp.ShowFound(found6zyol[:]), unmarshalMsgFieldOrder6zyol)
		if encodedFieldsLeft6zyol > 0 {
			encodedFieldsLeft6zyol--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zyol = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zyol < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zyol = 0
			}
			for nextMiss6zyol < maxFields6zyol && found6zyol[nextMiss6zyol] {
				nextMiss6zyol++
			}
			if nextMiss6zyol == maxFields6zyol {
				// filled all the empty fields!
				break doneWithStruct6zyol
			}
			missingFieldsLeft6zyol--
			curField6zyol = unmarshalMsgFieldOrder6zyol[nextMiss6zyol]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zyol)
		switch curField6zyol {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zyol[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zyol[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zpzw uint32
				zpzw, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zpzw) {
					z.Structs = (z.Structs)[:zpzw]
				} else {
					z.Structs = make([]StructT, zpzw)
				}
				for zgli := range z.Structs {
					const maxFields7zfyf = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zfyf uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zfyf, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zfyf := totalEncodedFields7zfyf
					missingFieldsLeft7zfyf := maxFields7zfyf - totalEncodedFields7zfyf

					var nextMiss7zfyf int32 = -1
					var found7zfyf [maxFields7zfyf]bool
					var curField7zfyf string

				doneWithStruct7zfyf:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zfyf > 0 || missingFieldsLeft7zfyf > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zfyf, missingFieldsLeft7zfyf, msgp.ShowFound(found7zfyf[:]), unmarshalMsgFieldOrder7zfyf)
						if encodedFieldsLeft7zfyf > 0 {
							encodedFieldsLeft7zfyf--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zfyf = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zfyf < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zfyf = 0
							}
							for nextMiss7zfyf < maxFields7zfyf && found7zfyf[nextMiss7zfyf] {
								nextMiss7zfyf++
							}
							if nextMiss7zfyf == maxFields7zfyf {
								// filled all the empty fields!
								break doneWithStruct7zfyf
							}
							missingFieldsLeft7zfyf--
							curField7zfyf = unmarshalMsgFieldOrder7zfyf[nextMiss7zfyf]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zfyf)
						switch curField7zfyf {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zfyf[0] = true
							z.Structs[zgli].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zfyf[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zgli].Fld) = (z.Structs[zgli].Fld)[:0]
							} else {

								var zujr uint32
								zujr, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zgli].Fld) >= int(zujr) {
									z.Structs[zgli].Fld = (z.Structs[zgli].Fld)[:zujr]
								} else {
									z.Structs[zgli].Fld = make([]FieldT, zujr)
								}
								for zbrc := range z.Structs[zgli].Fld {
									bts, err = z.Structs[zgli].Fld[zbrc].UnmarshalMsg(bts)
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
					if nextMiss7zfyf != -1 {
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
	if nextMiss6zyol != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zyol = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zfyf = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zgli := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zgli].StructName) + 4 + msgp.ArrayHeaderSize
		for zbrc := range z.Structs[zgli].Fld {
			s += z.Structs[zgli].Fld[zbrc].Msgsize()
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
		var zzfk uint8
		zzfk, err = dc.ReadUint8()
		(*z) = Zprimitive(zzfk)
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
		var zytt uint8
		zytt, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zytt)
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
