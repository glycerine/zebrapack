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
	const maxFields0zlmd = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zlmd uint32
	totalEncodedFields0zlmd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zlmd := totalEncodedFields0zlmd
	missingFieldsLeft0zlmd := maxFields0zlmd - totalEncodedFields0zlmd

	var nextMiss0zlmd int32 = -1
	var found0zlmd [maxFields0zlmd]bool
	var curField0zlmd string

doneWithStruct0zlmd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zlmd > 0 || missingFieldsLeft0zlmd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zlmd, missingFieldsLeft0zlmd, msgp.ShowFound(found0zlmd[:]), decodeMsgFieldOrder0zlmd)
		if encodedFieldsLeft0zlmd > 0 {
			encodedFieldsLeft0zlmd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zlmd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zlmd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zlmd = 0
			}
			for nextMiss0zlmd < maxFields0zlmd && found0zlmd[nextMiss0zlmd] {
				nextMiss0zlmd++
			}
			if nextMiss0zlmd == maxFields0zlmd {
				// filled all the empty fields!
				break doneWithStruct0zlmd
			}
			missingFieldsLeft0zlmd--
			curField0zlmd = decodeMsgFieldOrder0zlmd[nextMiss0zlmd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zlmd)
		switch curField0zlmd {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zlmd[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zlmd[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zlmd[2] = true
			{
				var zgqc uint8
				zgqc, err = dc.ReadUint8()
				z.Typ = Zprimitive(zgqc)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zlmd[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zlmd[4] = true
			var zgkn uint32
			zgkn, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zgkn > 0 {
				z.Tag = make(map[string]string, zgkn)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zgkn > 0 {
				zgkn--
				var zddr string
				var ztwu string
				zddr, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				ztwu, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zddr] = ztwu
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zlmd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zlmd = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zisb [5]bool
	fieldsInUse_zzjf := z.fieldsNotEmpty(empty_zisb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zzjf)
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
	if !empty_zisb[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zddr, ztwu := range z.Tag {
			err = en.WriteString(zddr)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(ztwu)
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
		for zddr, ztwu := range z.Tag {
			o = msgp.AppendString(o, zddr)
			o = msgp.AppendString(o, ztwu)
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
	const maxFields1zfgx = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zfgx uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zfgx, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zfgx := totalEncodedFields1zfgx
	missingFieldsLeft1zfgx := maxFields1zfgx - totalEncodedFields1zfgx

	var nextMiss1zfgx int32 = -1
	var found1zfgx [maxFields1zfgx]bool
	var curField1zfgx string

doneWithStruct1zfgx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zfgx > 0 || missingFieldsLeft1zfgx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zfgx, missingFieldsLeft1zfgx, msgp.ShowFound(found1zfgx[:]), unmarshalMsgFieldOrder1zfgx)
		if encodedFieldsLeft1zfgx > 0 {
			encodedFieldsLeft1zfgx--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zfgx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zfgx < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zfgx = 0
			}
			for nextMiss1zfgx < maxFields1zfgx && found1zfgx[nextMiss1zfgx] {
				nextMiss1zfgx++
			}
			if nextMiss1zfgx == maxFields1zfgx {
				// filled all the empty fields!
				break doneWithStruct1zfgx
			}
			missingFieldsLeft1zfgx--
			curField1zfgx = unmarshalMsgFieldOrder1zfgx[nextMiss1zfgx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zfgx)
		switch curField1zfgx {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zfgx[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zfgx[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zfgx[2] = true
			{
				var zpjf uint8
				zpjf, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zpjf)
			}
		case "TypStr":
			found1zfgx[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zfgx[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var ztuu uint32
				ztuu, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && ztuu > 0 {
					z.Tag = make(map[string]string, ztuu)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for ztuu > 0 {
					var zddr string
					var ztwu string
					ztuu--
					zddr, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					ztwu, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zddr] = ztwu
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zfgx != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zfgx = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zddr, ztwu := range z.Tag {
			_ = ztwu
			_ = zddr
			s += msgp.StringPrefixSize + len(zddr) + msgp.StringPrefixSize + len(ztwu)
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
	const maxFields2zihk = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zihk uint32
	totalEncodedFields2zihk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zihk := totalEncodedFields2zihk
	missingFieldsLeft2zihk := maxFields2zihk - totalEncodedFields2zihk

	var nextMiss2zihk int32 = -1
	var found2zihk [maxFields2zihk]bool
	var curField2zihk string

doneWithStruct2zihk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zihk > 0 || missingFieldsLeft2zihk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zihk, missingFieldsLeft2zihk, msgp.ShowFound(found2zihk[:]), decodeMsgFieldOrder2zihk)
		if encodedFieldsLeft2zihk > 0 {
			encodedFieldsLeft2zihk--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zihk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zihk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zihk = 0
			}
			for nextMiss2zihk < maxFields2zihk && found2zihk[nextMiss2zihk] {
				nextMiss2zihk++
			}
			if nextMiss2zihk == maxFields2zihk {
				// filled all the empty fields!
				break doneWithStruct2zihk
			}
			missingFieldsLeft2zihk--
			curField2zihk = decodeMsgFieldOrder2zihk[nextMiss2zihk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zihk)
		switch curField2zihk {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zihk[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zihk[1] = true
			var zvge uint32
			zvge, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zvge) {
				z.Fld = (z.Fld)[:zvge]
			} else {
				z.Fld = make([]FieldT, zvge)
			}
			for zihu := range z.Fld {
				err = z.Fld[zihu].DecodeMsg(dc)
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
	if nextMiss2zihk != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zihk = []string{"StructName", "Fld"}

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
	for zihu := range z.Fld {
		err = z.Fld[zihu].EncodeMsg(en)
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
	for zihu := range z.Fld {
		o, err = z.Fld[zihu].MarshalMsg(o)
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
	const maxFields3zlvj = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zlvj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zlvj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zlvj := totalEncodedFields3zlvj
	missingFieldsLeft3zlvj := maxFields3zlvj - totalEncodedFields3zlvj

	var nextMiss3zlvj int32 = -1
	var found3zlvj [maxFields3zlvj]bool
	var curField3zlvj string

doneWithStruct3zlvj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zlvj > 0 || missingFieldsLeft3zlvj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zlvj, missingFieldsLeft3zlvj, msgp.ShowFound(found3zlvj[:]), unmarshalMsgFieldOrder3zlvj)
		if encodedFieldsLeft3zlvj > 0 {
			encodedFieldsLeft3zlvj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zlvj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zlvj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zlvj = 0
			}
			for nextMiss3zlvj < maxFields3zlvj && found3zlvj[nextMiss3zlvj] {
				nextMiss3zlvj++
			}
			if nextMiss3zlvj == maxFields3zlvj {
				// filled all the empty fields!
				break doneWithStruct3zlvj
			}
			missingFieldsLeft3zlvj--
			curField3zlvj = unmarshalMsgFieldOrder3zlvj[nextMiss3zlvj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zlvj)
		switch curField3zlvj {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zlvj[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zlvj[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zsfl uint32
				zsfl, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zsfl) {
					z.Fld = (z.Fld)[:zsfl]
				} else {
					z.Fld = make([]FieldT, zsfl)
				}
				for zihu := range z.Fld {
					bts, err = z.Fld[zihu].UnmarshalMsg(bts)
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
	if nextMiss3zlvj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zlvj = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zihu := range z.Fld {
		s += z.Fld[zihu].Msgsize()
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
	const maxFields4zbuy = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zbuy uint32
	totalEncodedFields4zbuy, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zbuy := totalEncodedFields4zbuy
	missingFieldsLeft4zbuy := maxFields4zbuy - totalEncodedFields4zbuy

	var nextMiss4zbuy int32 = -1
	var found4zbuy [maxFields4zbuy]bool
	var curField4zbuy string

doneWithStruct4zbuy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zbuy > 0 || missingFieldsLeft4zbuy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zbuy, missingFieldsLeft4zbuy, msgp.ShowFound(found4zbuy[:]), decodeMsgFieldOrder4zbuy)
		if encodedFieldsLeft4zbuy > 0 {
			encodedFieldsLeft4zbuy--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zbuy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zbuy < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zbuy = 0
			}
			for nextMiss4zbuy < maxFields4zbuy && found4zbuy[nextMiss4zbuy] {
				nextMiss4zbuy++
			}
			if nextMiss4zbuy == maxFields4zbuy {
				// filled all the empty fields!
				break doneWithStruct4zbuy
			}
			missingFieldsLeft4zbuy--
			curField4zbuy = decodeMsgFieldOrder4zbuy[nextMiss4zbuy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zbuy)
		switch curField4zbuy {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zbuy[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zbuy[1] = true
			var zebb uint32
			zebb, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zebb) {
				z.Structs = (z.Structs)[:zebb]
			} else {
				z.Structs = make([]StructT, zebb)
			}
			for zvax := range z.Structs {
				const maxFields5zxam = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zxam uint32
				totalEncodedFields5zxam, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zxam := totalEncodedFields5zxam
				missingFieldsLeft5zxam := maxFields5zxam - totalEncodedFields5zxam

				var nextMiss5zxam int32 = -1
				var found5zxam [maxFields5zxam]bool
				var curField5zxam string

			doneWithStruct5zxam:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zxam > 0 || missingFieldsLeft5zxam > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxam, missingFieldsLeft5zxam, msgp.ShowFound(found5zxam[:]), decodeMsgFieldOrder5zxam)
					if encodedFieldsLeft5zxam > 0 {
						encodedFieldsLeft5zxam--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zxam = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zxam < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zxam = 0
						}
						for nextMiss5zxam < maxFields5zxam && found5zxam[nextMiss5zxam] {
							nextMiss5zxam++
						}
						if nextMiss5zxam == maxFields5zxam {
							// filled all the empty fields!
							break doneWithStruct5zxam
						}
						missingFieldsLeft5zxam--
						curField5zxam = decodeMsgFieldOrder5zxam[nextMiss5zxam]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zxam)
					switch curField5zxam {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zxam[0] = true
						z.Structs[zvax].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zxam[1] = true
						var zjpl uint32
						zjpl, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zvax].Fld) >= int(zjpl) {
							z.Structs[zvax].Fld = (z.Structs[zvax].Fld)[:zjpl]
						} else {
							z.Structs[zvax].Fld = make([]FieldT, zjpl)
						}
						for zjfr := range z.Structs[zvax].Fld {
							err = z.Structs[zvax].Fld[zjfr].DecodeMsg(dc)
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
				if nextMiss5zxam != -1 {
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
	if nextMiss4zbuy != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zbuy = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zxam = []string{"StructName", "Fld"}

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
	for zvax := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zvax].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zvax].Fld)))
		if err != nil {
			panic(err)
		}
		for zjfr := range z.Structs[zvax].Fld {
			err = z.Structs[zvax].Fld[zjfr].EncodeMsg(en)
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
	for zvax := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zvax].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zvax].Fld)))
		for zjfr := range z.Structs[zvax].Fld {
			o, err = z.Structs[zvax].Fld[zjfr].MarshalMsg(o)
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
	const maxFields6zyqk = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zyqk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zyqk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zyqk := totalEncodedFields6zyqk
	missingFieldsLeft6zyqk := maxFields6zyqk - totalEncodedFields6zyqk

	var nextMiss6zyqk int32 = -1
	var found6zyqk [maxFields6zyqk]bool
	var curField6zyqk string

doneWithStruct6zyqk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zyqk > 0 || missingFieldsLeft6zyqk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zyqk, missingFieldsLeft6zyqk, msgp.ShowFound(found6zyqk[:]), unmarshalMsgFieldOrder6zyqk)
		if encodedFieldsLeft6zyqk > 0 {
			encodedFieldsLeft6zyqk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zyqk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zyqk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zyqk = 0
			}
			for nextMiss6zyqk < maxFields6zyqk && found6zyqk[nextMiss6zyqk] {
				nextMiss6zyqk++
			}
			if nextMiss6zyqk == maxFields6zyqk {
				// filled all the empty fields!
				break doneWithStruct6zyqk
			}
			missingFieldsLeft6zyqk--
			curField6zyqk = unmarshalMsgFieldOrder6zyqk[nextMiss6zyqk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zyqk)
		switch curField6zyqk {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zyqk[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zyqk[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zkgi uint32
				zkgi, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zkgi) {
					z.Structs = (z.Structs)[:zkgi]
				} else {
					z.Structs = make([]StructT, zkgi)
				}
				for zvax := range z.Structs {
					const maxFields7zktc = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zktc uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zktc, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zktc := totalEncodedFields7zktc
					missingFieldsLeft7zktc := maxFields7zktc - totalEncodedFields7zktc

					var nextMiss7zktc int32 = -1
					var found7zktc [maxFields7zktc]bool
					var curField7zktc string

				doneWithStruct7zktc:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zktc > 0 || missingFieldsLeft7zktc > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zktc, missingFieldsLeft7zktc, msgp.ShowFound(found7zktc[:]), unmarshalMsgFieldOrder7zktc)
						if encodedFieldsLeft7zktc > 0 {
							encodedFieldsLeft7zktc--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zktc = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zktc < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zktc = 0
							}
							for nextMiss7zktc < maxFields7zktc && found7zktc[nextMiss7zktc] {
								nextMiss7zktc++
							}
							if nextMiss7zktc == maxFields7zktc {
								// filled all the empty fields!
								break doneWithStruct7zktc
							}
							missingFieldsLeft7zktc--
							curField7zktc = unmarshalMsgFieldOrder7zktc[nextMiss7zktc]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zktc)
						switch curField7zktc {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zktc[0] = true
							z.Structs[zvax].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zktc[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zvax].Fld) = (z.Structs[zvax].Fld)[:0]
							} else {

								var zumn uint32
								zumn, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zvax].Fld) >= int(zumn) {
									z.Structs[zvax].Fld = (z.Structs[zvax].Fld)[:zumn]
								} else {
									z.Structs[zvax].Fld = make([]FieldT, zumn)
								}
								for zjfr := range z.Structs[zvax].Fld {
									bts, err = z.Structs[zvax].Fld[zjfr].UnmarshalMsg(bts)
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
					if nextMiss7zktc != -1 {
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
	if nextMiss6zyqk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zyqk = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zktc = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zvax := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zvax].StructName) + 4 + msgp.ArrayHeaderSize
		for zjfr := range z.Structs[zvax].Fld {
			s += z.Structs[zvax].Fld[zjfr].Msgsize()
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
		var zpcd uint8
		zpcd, err = dc.ReadUint8()
		(*z) = Zprimitive(zpcd)
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
		var zeej uint8
		zeej, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zeej)
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
