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
	const maxFields0zhch = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zhch uint32
	totalEncodedFields0zhch, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zhch := totalEncodedFields0zhch
	missingFieldsLeft0zhch := maxFields0zhch - totalEncodedFields0zhch

	var nextMiss0zhch int32 = -1
	var found0zhch [maxFields0zhch]bool
	var curField0zhch string

doneWithStruct0zhch:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zhch > 0 || missingFieldsLeft0zhch > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zhch, missingFieldsLeft0zhch, msgp.ShowFound(found0zhch[:]), decodeMsgFieldOrder0zhch)
		if encodedFieldsLeft0zhch > 0 {
			encodedFieldsLeft0zhch--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zhch = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zhch < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zhch = 0
			}
			for nextMiss0zhch < maxFields0zhch && found0zhch[nextMiss0zhch] {
				nextMiss0zhch++
			}
			if nextMiss0zhch == maxFields0zhch {
				// filled all the empty fields!
				break doneWithStruct0zhch
			}
			missingFieldsLeft0zhch--
			curField0zhch = decodeMsgFieldOrder0zhch[nextMiss0zhch]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zhch)
		switch curField0zhch {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zhch[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zhch[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zhch[2] = true
			{
				var zheb uint8
				zheb, err = dc.ReadUint8()
				z.Typ = Zprimitive(zheb)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zhch[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zhch[4] = true
			var zyne uint32
			zyne, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zyne > 0 {
				z.Tag = make(map[string]string, zyne)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zyne > 0 {
				zyne--
				var zisz string
				var zobf string
				zisz, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zobf, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zisz] = zobf
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zhch != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zhch = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zhpf [5]bool
	fieldsInUse_ztfv := z.fieldsNotEmpty(empty_zhpf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_ztfv)
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
	if !empty_zhpf[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zisz, zobf := range z.Tag {
			err = en.WriteString(zisz)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zobf)
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
		for zisz, zobf := range z.Tag {
			o = msgp.AppendString(o, zisz)
			o = msgp.AppendString(o, zobf)
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
	const maxFields1zsvo = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zsvo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zsvo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zsvo := totalEncodedFields1zsvo
	missingFieldsLeft1zsvo := maxFields1zsvo - totalEncodedFields1zsvo

	var nextMiss1zsvo int32 = -1
	var found1zsvo [maxFields1zsvo]bool
	var curField1zsvo string

doneWithStruct1zsvo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zsvo > 0 || missingFieldsLeft1zsvo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zsvo, missingFieldsLeft1zsvo, msgp.ShowFound(found1zsvo[:]), unmarshalMsgFieldOrder1zsvo)
		if encodedFieldsLeft1zsvo > 0 {
			encodedFieldsLeft1zsvo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zsvo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zsvo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zsvo = 0
			}
			for nextMiss1zsvo < maxFields1zsvo && found1zsvo[nextMiss1zsvo] {
				nextMiss1zsvo++
			}
			if nextMiss1zsvo == maxFields1zsvo {
				// filled all the empty fields!
				break doneWithStruct1zsvo
			}
			missingFieldsLeft1zsvo--
			curField1zsvo = unmarshalMsgFieldOrder1zsvo[nextMiss1zsvo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zsvo)
		switch curField1zsvo {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zsvo[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zsvo[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zsvo[2] = true
			{
				var zmzj uint8
				zmzj, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zmzj)
			}
		case "TypStr":
			found1zsvo[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zsvo[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zkpz uint32
				zkpz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zkpz > 0 {
					z.Tag = make(map[string]string, zkpz)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zkpz > 0 {
					var zisz string
					var zobf string
					zkpz--
					zisz, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zobf, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zisz] = zobf
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zsvo != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zsvo = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zisz, zobf := range z.Tag {
			_ = zobf
			_ = zisz
			s += msgp.StringPrefixSize + len(zisz) + msgp.StringPrefixSize + len(zobf)
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
	const maxFields2zvgn = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zvgn uint32
	totalEncodedFields2zvgn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zvgn := totalEncodedFields2zvgn
	missingFieldsLeft2zvgn := maxFields2zvgn - totalEncodedFields2zvgn

	var nextMiss2zvgn int32 = -1
	var found2zvgn [maxFields2zvgn]bool
	var curField2zvgn string

doneWithStruct2zvgn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zvgn > 0 || missingFieldsLeft2zvgn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zvgn, missingFieldsLeft2zvgn, msgp.ShowFound(found2zvgn[:]), decodeMsgFieldOrder2zvgn)
		if encodedFieldsLeft2zvgn > 0 {
			encodedFieldsLeft2zvgn--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zvgn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zvgn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zvgn = 0
			}
			for nextMiss2zvgn < maxFields2zvgn && found2zvgn[nextMiss2zvgn] {
				nextMiss2zvgn++
			}
			if nextMiss2zvgn == maxFields2zvgn {
				// filled all the empty fields!
				break doneWithStruct2zvgn
			}
			missingFieldsLeft2zvgn--
			curField2zvgn = decodeMsgFieldOrder2zvgn[nextMiss2zvgn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zvgn)
		switch curField2zvgn {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zvgn[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zvgn[1] = true
			var zvty uint32
			zvty, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zvty) {
				z.Fld = (z.Fld)[:zvty]
			} else {
				z.Fld = make([]FieldT, zvty)
			}
			for zpig := range z.Fld {
				err = z.Fld[zpig].DecodeMsg(dc)
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
	if nextMiss2zvgn != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zvgn = []string{"StructName", "Fld"}

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
	for zpig := range z.Fld {
		err = z.Fld[zpig].EncodeMsg(en)
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
	for zpig := range z.Fld {
		o, err = z.Fld[zpig].MarshalMsg(o)
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
	const maxFields3zfdw = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zfdw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zfdw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zfdw := totalEncodedFields3zfdw
	missingFieldsLeft3zfdw := maxFields3zfdw - totalEncodedFields3zfdw

	var nextMiss3zfdw int32 = -1
	var found3zfdw [maxFields3zfdw]bool
	var curField3zfdw string

doneWithStruct3zfdw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zfdw > 0 || missingFieldsLeft3zfdw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zfdw, missingFieldsLeft3zfdw, msgp.ShowFound(found3zfdw[:]), unmarshalMsgFieldOrder3zfdw)
		if encodedFieldsLeft3zfdw > 0 {
			encodedFieldsLeft3zfdw--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zfdw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zfdw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zfdw = 0
			}
			for nextMiss3zfdw < maxFields3zfdw && found3zfdw[nextMiss3zfdw] {
				nextMiss3zfdw++
			}
			if nextMiss3zfdw == maxFields3zfdw {
				// filled all the empty fields!
				break doneWithStruct3zfdw
			}
			missingFieldsLeft3zfdw--
			curField3zfdw = unmarshalMsgFieldOrder3zfdw[nextMiss3zfdw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zfdw)
		switch curField3zfdw {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zfdw[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zfdw[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zzao uint32
				zzao, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zzao) {
					z.Fld = (z.Fld)[:zzao]
				} else {
					z.Fld = make([]FieldT, zzao)
				}
				for zpig := range z.Fld {
					bts, err = z.Fld[zpig].UnmarshalMsg(bts)
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
	if nextMiss3zfdw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zfdw = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zpig := range z.Fld {
		s += z.Fld[zpig].Msgsize()
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
	const maxFields4zuol = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zuol uint32
	totalEncodedFields4zuol, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zuol := totalEncodedFields4zuol
	missingFieldsLeft4zuol := maxFields4zuol - totalEncodedFields4zuol

	var nextMiss4zuol int32 = -1
	var found4zuol [maxFields4zuol]bool
	var curField4zuol string

doneWithStruct4zuol:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zuol > 0 || missingFieldsLeft4zuol > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zuol, missingFieldsLeft4zuol, msgp.ShowFound(found4zuol[:]), decodeMsgFieldOrder4zuol)
		if encodedFieldsLeft4zuol > 0 {
			encodedFieldsLeft4zuol--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zuol = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zuol < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zuol = 0
			}
			for nextMiss4zuol < maxFields4zuol && found4zuol[nextMiss4zuol] {
				nextMiss4zuol++
			}
			if nextMiss4zuol == maxFields4zuol {
				// filled all the empty fields!
				break doneWithStruct4zuol
			}
			missingFieldsLeft4zuol--
			curField4zuol = decodeMsgFieldOrder4zuol[nextMiss4zuol]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zuol)
		switch curField4zuol {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zuol[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zuol[1] = true
			var zuag uint32
			zuag, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zuag) {
				z.Structs = (z.Structs)[:zuag]
			} else {
				z.Structs = make([]StructT, zuag)
			}
			for ziux := range z.Structs {
				const maxFields5zcyj = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zcyj uint32
				totalEncodedFields5zcyj, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zcyj := totalEncodedFields5zcyj
				missingFieldsLeft5zcyj := maxFields5zcyj - totalEncodedFields5zcyj

				var nextMiss5zcyj int32 = -1
				var found5zcyj [maxFields5zcyj]bool
				var curField5zcyj string

			doneWithStruct5zcyj:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zcyj > 0 || missingFieldsLeft5zcyj > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zcyj, missingFieldsLeft5zcyj, msgp.ShowFound(found5zcyj[:]), decodeMsgFieldOrder5zcyj)
					if encodedFieldsLeft5zcyj > 0 {
						encodedFieldsLeft5zcyj--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zcyj = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zcyj < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zcyj = 0
						}
						for nextMiss5zcyj < maxFields5zcyj && found5zcyj[nextMiss5zcyj] {
							nextMiss5zcyj++
						}
						if nextMiss5zcyj == maxFields5zcyj {
							// filled all the empty fields!
							break doneWithStruct5zcyj
						}
						missingFieldsLeft5zcyj--
						curField5zcyj = decodeMsgFieldOrder5zcyj[nextMiss5zcyj]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zcyj)
					switch curField5zcyj {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zcyj[0] = true
						z.Structs[ziux].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zcyj[1] = true
						var zowh uint32
						zowh, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[ziux].Fld) >= int(zowh) {
							z.Structs[ziux].Fld = (z.Structs[ziux].Fld)[:zowh]
						} else {
							z.Structs[ziux].Fld = make([]FieldT, zowh)
						}
						for zwfu := range z.Structs[ziux].Fld {
							err = z.Structs[ziux].Fld[zwfu].DecodeMsg(dc)
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
				if nextMiss5zcyj != -1 {
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
	if nextMiss4zuol != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zuol = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zcyj = []string{"StructName", "Fld"}

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
	for ziux := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[ziux].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[ziux].Fld)))
		if err != nil {
			panic(err)
		}
		for zwfu := range z.Structs[ziux].Fld {
			err = z.Structs[ziux].Fld[zwfu].EncodeMsg(en)
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
	for ziux := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[ziux].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[ziux].Fld)))
		for zwfu := range z.Structs[ziux].Fld {
			o, err = z.Structs[ziux].Fld[zwfu].MarshalMsg(o)
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
	const maxFields6zdco = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zdco uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zdco, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zdco := totalEncodedFields6zdco
	missingFieldsLeft6zdco := maxFields6zdco - totalEncodedFields6zdco

	var nextMiss6zdco int32 = -1
	var found6zdco [maxFields6zdco]bool
	var curField6zdco string

doneWithStruct6zdco:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zdco > 0 || missingFieldsLeft6zdco > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zdco, missingFieldsLeft6zdco, msgp.ShowFound(found6zdco[:]), unmarshalMsgFieldOrder6zdco)
		if encodedFieldsLeft6zdco > 0 {
			encodedFieldsLeft6zdco--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zdco = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zdco < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zdco = 0
			}
			for nextMiss6zdco < maxFields6zdco && found6zdco[nextMiss6zdco] {
				nextMiss6zdco++
			}
			if nextMiss6zdco == maxFields6zdco {
				// filled all the empty fields!
				break doneWithStruct6zdco
			}
			missingFieldsLeft6zdco--
			curField6zdco = unmarshalMsgFieldOrder6zdco[nextMiss6zdco]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zdco)
		switch curField6zdco {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zdco[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zdco[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zxxq uint32
				zxxq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zxxq) {
					z.Structs = (z.Structs)[:zxxq]
				} else {
					z.Structs = make([]StructT, zxxq)
				}
				for ziux := range z.Structs {
					const maxFields7zgge = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zgge uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zgge, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zgge := totalEncodedFields7zgge
					missingFieldsLeft7zgge := maxFields7zgge - totalEncodedFields7zgge

					var nextMiss7zgge int32 = -1
					var found7zgge [maxFields7zgge]bool
					var curField7zgge string

				doneWithStruct7zgge:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zgge > 0 || missingFieldsLeft7zgge > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zgge, missingFieldsLeft7zgge, msgp.ShowFound(found7zgge[:]), unmarshalMsgFieldOrder7zgge)
						if encodedFieldsLeft7zgge > 0 {
							encodedFieldsLeft7zgge--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zgge = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zgge < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zgge = 0
							}
							for nextMiss7zgge < maxFields7zgge && found7zgge[nextMiss7zgge] {
								nextMiss7zgge++
							}
							if nextMiss7zgge == maxFields7zgge {
								// filled all the empty fields!
								break doneWithStruct7zgge
							}
							missingFieldsLeft7zgge--
							curField7zgge = unmarshalMsgFieldOrder7zgge[nextMiss7zgge]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zgge)
						switch curField7zgge {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zgge[0] = true
							z.Structs[ziux].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zgge[1] = true
							if nbs.AlwaysNil {
								(z.Structs[ziux].Fld) = (z.Structs[ziux].Fld)[:0]
							} else {

								var zyir uint32
								zyir, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[ziux].Fld) >= int(zyir) {
									z.Structs[ziux].Fld = (z.Structs[ziux].Fld)[:zyir]
								} else {
									z.Structs[ziux].Fld = make([]FieldT, zyir)
								}
								for zwfu := range z.Structs[ziux].Fld {
									bts, err = z.Structs[ziux].Fld[zwfu].UnmarshalMsg(bts)
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
					if nextMiss7zgge != -1 {
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
	if nextMiss6zdco != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zdco = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zgge = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for ziux := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[ziux].StructName) + 4 + msgp.ArrayHeaderSize
		for zwfu := range z.Structs[ziux].Fld {
			s += z.Structs[ziux].Fld[zwfu].Msgsize()
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
		var znsf uint8
		znsf, err = dc.ReadUint8()
		(*z) = Zprimitive(znsf)
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
		var zprj uint8
		zprj, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zprj)
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
