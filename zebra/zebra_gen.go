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
	const maxFields0zbwg = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zbwg uint32
	totalEncodedFields0zbwg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zbwg := totalEncodedFields0zbwg
	missingFieldsLeft0zbwg := maxFields0zbwg - totalEncodedFields0zbwg

	var nextMiss0zbwg int32 = -1
	var found0zbwg [maxFields0zbwg]bool
	var curField0zbwg string

doneWithStruct0zbwg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zbwg > 0 || missingFieldsLeft0zbwg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zbwg, missingFieldsLeft0zbwg, msgp.ShowFound(found0zbwg[:]), decodeMsgFieldOrder0zbwg)
		if encodedFieldsLeft0zbwg > 0 {
			encodedFieldsLeft0zbwg--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zbwg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zbwg < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zbwg = 0
			}
			for nextMiss0zbwg < maxFields0zbwg && found0zbwg[nextMiss0zbwg] {
				nextMiss0zbwg++
			}
			if nextMiss0zbwg == maxFields0zbwg {
				// filled all the empty fields!
				break doneWithStruct0zbwg
			}
			missingFieldsLeft0zbwg--
			curField0zbwg = decodeMsgFieldOrder0zbwg[nextMiss0zbwg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zbwg)
		switch curField0zbwg {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zbwg[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zbwg[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zbwg[2] = true
			{
				var zamk uint8
				zamk, err = dc.ReadUint8()
				z.Typ = Zprimitive(zamk)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zbwg[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zbwg[4] = true
			var zlpm uint32
			zlpm, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zlpm > 0 {
				z.Tag = make(map[string]string, zlpm)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zlpm > 0 {
				zlpm--
				var zqzy string
				var znyw string
				zqzy, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				znyw, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zqzy] = znyw
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zbwg != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zbwg = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zsmp [5]bool
	fieldsInUse_zuob := z.fieldsNotEmpty(empty_zsmp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zuob)
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
	if !empty_zsmp[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zqzy, znyw := range z.Tag {
			err = en.WriteString(zqzy)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(znyw)
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
		for zqzy, znyw := range z.Tag {
			o = msgp.AppendString(o, zqzy)
			o = msgp.AppendString(o, znyw)
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
	const maxFields1zels = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zels uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zels, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zels := totalEncodedFields1zels
	missingFieldsLeft1zels := maxFields1zels - totalEncodedFields1zels

	var nextMiss1zels int32 = -1
	var found1zels [maxFields1zels]bool
	var curField1zels string

doneWithStruct1zels:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zels > 0 || missingFieldsLeft1zels > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zels, missingFieldsLeft1zels, msgp.ShowFound(found1zels[:]), unmarshalMsgFieldOrder1zels)
		if encodedFieldsLeft1zels > 0 {
			encodedFieldsLeft1zels--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zels = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zels < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zels = 0
			}
			for nextMiss1zels < maxFields1zels && found1zels[nextMiss1zels] {
				nextMiss1zels++
			}
			if nextMiss1zels == maxFields1zels {
				// filled all the empty fields!
				break doneWithStruct1zels
			}
			missingFieldsLeft1zels--
			curField1zels = unmarshalMsgFieldOrder1zels[nextMiss1zels]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zels)
		switch curField1zels {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zels[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zels[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zels[2] = true
			{
				var zflr uint8
				zflr, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zflr)
			}
		case "TypStr":
			found1zels[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zels[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zjeh uint32
				zjeh, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zjeh > 0 {
					z.Tag = make(map[string]string, zjeh)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zjeh > 0 {
					var zqzy string
					var znyw string
					zjeh--
					zqzy, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					znyw, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zqzy] = znyw
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zels != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zels = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zqzy, znyw := range z.Tag {
			_ = znyw
			_ = zqzy
			s += msgp.StringPrefixSize + len(zqzy) + msgp.StringPrefixSize + len(znyw)
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
	const maxFields2zqad = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zqad uint32
	totalEncodedFields2zqad, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zqad := totalEncodedFields2zqad
	missingFieldsLeft2zqad := maxFields2zqad - totalEncodedFields2zqad

	var nextMiss2zqad int32 = -1
	var found2zqad [maxFields2zqad]bool
	var curField2zqad string

doneWithStruct2zqad:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zqad > 0 || missingFieldsLeft2zqad > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zqad, missingFieldsLeft2zqad, msgp.ShowFound(found2zqad[:]), decodeMsgFieldOrder2zqad)
		if encodedFieldsLeft2zqad > 0 {
			encodedFieldsLeft2zqad--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zqad = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zqad < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zqad = 0
			}
			for nextMiss2zqad < maxFields2zqad && found2zqad[nextMiss2zqad] {
				nextMiss2zqad++
			}
			if nextMiss2zqad == maxFields2zqad {
				// filled all the empty fields!
				break doneWithStruct2zqad
			}
			missingFieldsLeft2zqad--
			curField2zqad = decodeMsgFieldOrder2zqad[nextMiss2zqad]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zqad)
		switch curField2zqad {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zqad[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zqad[1] = true
			var zkcj uint32
			zkcj, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zkcj) {
				z.Fld = (z.Fld)[:zkcj]
			} else {
				z.Fld = make([]FieldT, zkcj)
			}
			for zmsc := range z.Fld {
				err = z.Fld[zmsc].DecodeMsg(dc)
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
	if nextMiss2zqad != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zqad = []string{"StructName", "Fld"}

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
	for zmsc := range z.Fld {
		err = z.Fld[zmsc].EncodeMsg(en)
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
	for zmsc := range z.Fld {
		o, err = z.Fld[zmsc].MarshalMsg(o)
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
	const maxFields3zdem = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zdem uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zdem, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zdem := totalEncodedFields3zdem
	missingFieldsLeft3zdem := maxFields3zdem - totalEncodedFields3zdem

	var nextMiss3zdem int32 = -1
	var found3zdem [maxFields3zdem]bool
	var curField3zdem string

doneWithStruct3zdem:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zdem > 0 || missingFieldsLeft3zdem > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zdem, missingFieldsLeft3zdem, msgp.ShowFound(found3zdem[:]), unmarshalMsgFieldOrder3zdem)
		if encodedFieldsLeft3zdem > 0 {
			encodedFieldsLeft3zdem--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zdem = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zdem < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zdem = 0
			}
			for nextMiss3zdem < maxFields3zdem && found3zdem[nextMiss3zdem] {
				nextMiss3zdem++
			}
			if nextMiss3zdem == maxFields3zdem {
				// filled all the empty fields!
				break doneWithStruct3zdem
			}
			missingFieldsLeft3zdem--
			curField3zdem = unmarshalMsgFieldOrder3zdem[nextMiss3zdem]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zdem)
		switch curField3zdem {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zdem[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zdem[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zqdm uint32
				zqdm, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zqdm) {
					z.Fld = (z.Fld)[:zqdm]
				} else {
					z.Fld = make([]FieldT, zqdm)
				}
				for zmsc := range z.Fld {
					bts, err = z.Fld[zmsc].UnmarshalMsg(bts)
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
	if nextMiss3zdem != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zdem = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zmsc := range z.Fld {
		s += z.Fld[zmsc].Msgsize()
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
	const maxFields4zdwd = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zdwd uint32
	totalEncodedFields4zdwd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zdwd := totalEncodedFields4zdwd
	missingFieldsLeft4zdwd := maxFields4zdwd - totalEncodedFields4zdwd

	var nextMiss4zdwd int32 = -1
	var found4zdwd [maxFields4zdwd]bool
	var curField4zdwd string

doneWithStruct4zdwd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zdwd > 0 || missingFieldsLeft4zdwd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zdwd, missingFieldsLeft4zdwd, msgp.ShowFound(found4zdwd[:]), decodeMsgFieldOrder4zdwd)
		if encodedFieldsLeft4zdwd > 0 {
			encodedFieldsLeft4zdwd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zdwd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zdwd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zdwd = 0
			}
			for nextMiss4zdwd < maxFields4zdwd && found4zdwd[nextMiss4zdwd] {
				nextMiss4zdwd++
			}
			if nextMiss4zdwd == maxFields4zdwd {
				// filled all the empty fields!
				break doneWithStruct4zdwd
			}
			missingFieldsLeft4zdwd--
			curField4zdwd = decodeMsgFieldOrder4zdwd[nextMiss4zdwd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zdwd)
		switch curField4zdwd {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zdwd[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zdwd[1] = true
			var zpej uint32
			zpej, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zpej) {
				z.Structs = (z.Structs)[:zpej]
			} else {
				z.Structs = make([]StructT, zpej)
			}
			for zjav := range z.Structs {
				const maxFields5zpni = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zpni uint32
				totalEncodedFields5zpni, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zpni := totalEncodedFields5zpni
				missingFieldsLeft5zpni := maxFields5zpni - totalEncodedFields5zpni

				var nextMiss5zpni int32 = -1
				var found5zpni [maxFields5zpni]bool
				var curField5zpni string

			doneWithStruct5zpni:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zpni > 0 || missingFieldsLeft5zpni > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zpni, missingFieldsLeft5zpni, msgp.ShowFound(found5zpni[:]), decodeMsgFieldOrder5zpni)
					if encodedFieldsLeft5zpni > 0 {
						encodedFieldsLeft5zpni--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zpni = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zpni < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zpni = 0
						}
						for nextMiss5zpni < maxFields5zpni && found5zpni[nextMiss5zpni] {
							nextMiss5zpni++
						}
						if nextMiss5zpni == maxFields5zpni {
							// filled all the empty fields!
							break doneWithStruct5zpni
						}
						missingFieldsLeft5zpni--
						curField5zpni = decodeMsgFieldOrder5zpni[nextMiss5zpni]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zpni)
					switch curField5zpni {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zpni[0] = true
						z.Structs[zjav].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zpni[1] = true
						var zouq uint32
						zouq, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zjav].Fld) >= int(zouq) {
							z.Structs[zjav].Fld = (z.Structs[zjav].Fld)[:zouq]
						} else {
							z.Structs[zjav].Fld = make([]FieldT, zouq)
						}
						for zbdj := range z.Structs[zjav].Fld {
							err = z.Structs[zjav].Fld[zbdj].DecodeMsg(dc)
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
				if nextMiss5zpni != -1 {
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
	if nextMiss4zdwd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zdwd = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zpni = []string{"StructName", "Fld"}

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
	for zjav := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zjav].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zjav].Fld)))
		if err != nil {
			panic(err)
		}
		for zbdj := range z.Structs[zjav].Fld {
			err = z.Structs[zjav].Fld[zbdj].EncodeMsg(en)
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
	for zjav := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zjav].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zjav].Fld)))
		for zbdj := range z.Structs[zjav].Fld {
			o, err = z.Structs[zjav].Fld[zbdj].MarshalMsg(o)
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
	const maxFields6zyja = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zyja uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zyja, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zyja := totalEncodedFields6zyja
	missingFieldsLeft6zyja := maxFields6zyja - totalEncodedFields6zyja

	var nextMiss6zyja int32 = -1
	var found6zyja [maxFields6zyja]bool
	var curField6zyja string

doneWithStruct6zyja:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zyja > 0 || missingFieldsLeft6zyja > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zyja, missingFieldsLeft6zyja, msgp.ShowFound(found6zyja[:]), unmarshalMsgFieldOrder6zyja)
		if encodedFieldsLeft6zyja > 0 {
			encodedFieldsLeft6zyja--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zyja = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zyja < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zyja = 0
			}
			for nextMiss6zyja < maxFields6zyja && found6zyja[nextMiss6zyja] {
				nextMiss6zyja++
			}
			if nextMiss6zyja == maxFields6zyja {
				// filled all the empty fields!
				break doneWithStruct6zyja
			}
			missingFieldsLeft6zyja--
			curField6zyja = unmarshalMsgFieldOrder6zyja[nextMiss6zyja]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zyja)
		switch curField6zyja {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zyja[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zyja[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zphs uint32
				zphs, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zphs) {
					z.Structs = (z.Structs)[:zphs]
				} else {
					z.Structs = make([]StructT, zphs)
				}
				for zjav := range z.Structs {
					const maxFields7zizq = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zizq uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zizq, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zizq := totalEncodedFields7zizq
					missingFieldsLeft7zizq := maxFields7zizq - totalEncodedFields7zizq

					var nextMiss7zizq int32 = -1
					var found7zizq [maxFields7zizq]bool
					var curField7zizq string

				doneWithStruct7zizq:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zizq > 0 || missingFieldsLeft7zizq > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zizq, missingFieldsLeft7zizq, msgp.ShowFound(found7zizq[:]), unmarshalMsgFieldOrder7zizq)
						if encodedFieldsLeft7zizq > 0 {
							encodedFieldsLeft7zizq--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zizq = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zizq < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zizq = 0
							}
							for nextMiss7zizq < maxFields7zizq && found7zizq[nextMiss7zizq] {
								nextMiss7zizq++
							}
							if nextMiss7zizq == maxFields7zizq {
								// filled all the empty fields!
								break doneWithStruct7zizq
							}
							missingFieldsLeft7zizq--
							curField7zizq = unmarshalMsgFieldOrder7zizq[nextMiss7zizq]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zizq)
						switch curField7zizq {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zizq[0] = true
							z.Structs[zjav].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zizq[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zjav].Fld) = (z.Structs[zjav].Fld)[:0]
							} else {

								var zbcy uint32
								zbcy, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zjav].Fld) >= int(zbcy) {
									z.Structs[zjav].Fld = (z.Structs[zjav].Fld)[:zbcy]
								} else {
									z.Structs[zjav].Fld = make([]FieldT, zbcy)
								}
								for zbdj := range z.Structs[zjav].Fld {
									bts, err = z.Structs[zjav].Fld[zbdj].UnmarshalMsg(bts)
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
					if nextMiss7zizq != -1 {
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
	if nextMiss6zyja != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zyja = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zizq = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zjav := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zjav].StructName) + 4 + msgp.ArrayHeaderSize
		for zbdj := range z.Structs[zjav].Fld {
			s += z.Structs[zjav].Fld[zbdj].Msgsize()
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
		var zhvx uint8
		zhvx, err = dc.ReadUint8()
		(*z) = Zprimitive(zhvx)
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
		var zksl uint8
		zksl, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zksl)
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
