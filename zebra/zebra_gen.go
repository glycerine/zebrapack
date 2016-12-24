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
	const maxFields0zady = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zady uint32
	totalEncodedFields0zady, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zady := totalEncodedFields0zady
	missingFieldsLeft0zady := maxFields0zady - totalEncodedFields0zady

	var nextMiss0zady int32 = -1
	var found0zady [maxFields0zady]bool
	var curField0zady string

doneWithStruct0zady:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zady > 0 || missingFieldsLeft0zady > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zady, missingFieldsLeft0zady, msgp.ShowFound(found0zady[:]), decodeMsgFieldOrder0zady)
		if encodedFieldsLeft0zady > 0 {
			encodedFieldsLeft0zady--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zady = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zady < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zady = 0
			}
			for nextMiss0zady < maxFields0zady && found0zady[nextMiss0zady] {
				nextMiss0zady++
			}
			if nextMiss0zady == maxFields0zady {
				// filled all the empty fields!
				break doneWithStruct0zady
			}
			missingFieldsLeft0zady--
			curField0zady = decodeMsgFieldOrder0zady[nextMiss0zady]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zady)
		switch curField0zady {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zady[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zady[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zady[2] = true
			{
				var zsat uint8
				zsat, err = dc.ReadUint8()
				z.Typ = Zprimitive(zsat)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zady[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zady[4] = true
			var zfsg uint32
			zfsg, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zfsg > 0 {
				z.Tag = make(map[string]string, zfsg)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zfsg > 0 {
				zfsg--
				var zodc string
				var zrlo string
				zodc, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zrlo, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zodc] = zrlo
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zady != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zady = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_ztqm [5]bool
	fieldsInUse_zfxr := z.fieldsNotEmpty(empty_ztqm[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zfxr)
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
	if !empty_ztqm[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zodc, zrlo := range z.Tag {
			err = en.WriteString(zodc)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zrlo)
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
		for zodc, zrlo := range z.Tag {
			o = msgp.AppendString(o, zodc)
			o = msgp.AppendString(o, zrlo)
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
	const maxFields1zrhd = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zrhd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zrhd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zrhd := totalEncodedFields1zrhd
	missingFieldsLeft1zrhd := maxFields1zrhd - totalEncodedFields1zrhd

	var nextMiss1zrhd int32 = -1
	var found1zrhd [maxFields1zrhd]bool
	var curField1zrhd string

doneWithStruct1zrhd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zrhd > 0 || missingFieldsLeft1zrhd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zrhd, missingFieldsLeft1zrhd, msgp.ShowFound(found1zrhd[:]), unmarshalMsgFieldOrder1zrhd)
		if encodedFieldsLeft1zrhd > 0 {
			encodedFieldsLeft1zrhd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zrhd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zrhd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zrhd = 0
			}
			for nextMiss1zrhd < maxFields1zrhd && found1zrhd[nextMiss1zrhd] {
				nextMiss1zrhd++
			}
			if nextMiss1zrhd == maxFields1zrhd {
				// filled all the empty fields!
				break doneWithStruct1zrhd
			}
			missingFieldsLeft1zrhd--
			curField1zrhd = unmarshalMsgFieldOrder1zrhd[nextMiss1zrhd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zrhd)
		switch curField1zrhd {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zrhd[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zrhd[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zrhd[2] = true
			{
				var zzyb uint8
				zzyb, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zzyb)
			}
		case "TypStr":
			found1zrhd[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zrhd[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zpzv uint32
				zpzv, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zpzv > 0 {
					z.Tag = make(map[string]string, zpzv)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zpzv > 0 {
					var zodc string
					var zrlo string
					zpzv--
					zodc, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zrlo, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zodc] = zrlo
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zrhd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zrhd = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zodc, zrlo := range z.Tag {
			_ = zrlo
			_ = zodc
			s += msgp.StringPrefixSize + len(zodc) + msgp.StringPrefixSize + len(zrlo)
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
	const maxFields2zexj = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zexj uint32
	totalEncodedFields2zexj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zexj := totalEncodedFields2zexj
	missingFieldsLeft2zexj := maxFields2zexj - totalEncodedFields2zexj

	var nextMiss2zexj int32 = -1
	var found2zexj [maxFields2zexj]bool
	var curField2zexj string

doneWithStruct2zexj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zexj > 0 || missingFieldsLeft2zexj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zexj, missingFieldsLeft2zexj, msgp.ShowFound(found2zexj[:]), decodeMsgFieldOrder2zexj)
		if encodedFieldsLeft2zexj > 0 {
			encodedFieldsLeft2zexj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zexj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zexj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zexj = 0
			}
			for nextMiss2zexj < maxFields2zexj && found2zexj[nextMiss2zexj] {
				nextMiss2zexj++
			}
			if nextMiss2zexj == maxFields2zexj {
				// filled all the empty fields!
				break doneWithStruct2zexj
			}
			missingFieldsLeft2zexj--
			curField2zexj = decodeMsgFieldOrder2zexj[nextMiss2zexj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zexj)
		switch curField2zexj {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zexj[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zexj[1] = true
			var zzfr uint32
			zzfr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zzfr) {
				z.Fld = (z.Fld)[:zzfr]
			} else {
				z.Fld = make([]FieldT, zzfr)
			}
			for zosz := range z.Fld {
				err = z.Fld[zosz].DecodeMsg(dc)
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
	if nextMiss2zexj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zexj = []string{"StructName", "Fld"}

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
	for zosz := range z.Fld {
		err = z.Fld[zosz].EncodeMsg(en)
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
	for zosz := range z.Fld {
		o, err = z.Fld[zosz].MarshalMsg(o)
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
	const maxFields3zbwr = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zbwr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zbwr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zbwr := totalEncodedFields3zbwr
	missingFieldsLeft3zbwr := maxFields3zbwr - totalEncodedFields3zbwr

	var nextMiss3zbwr int32 = -1
	var found3zbwr [maxFields3zbwr]bool
	var curField3zbwr string

doneWithStruct3zbwr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zbwr > 0 || missingFieldsLeft3zbwr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zbwr, missingFieldsLeft3zbwr, msgp.ShowFound(found3zbwr[:]), unmarshalMsgFieldOrder3zbwr)
		if encodedFieldsLeft3zbwr > 0 {
			encodedFieldsLeft3zbwr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zbwr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zbwr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zbwr = 0
			}
			for nextMiss3zbwr < maxFields3zbwr && found3zbwr[nextMiss3zbwr] {
				nextMiss3zbwr++
			}
			if nextMiss3zbwr == maxFields3zbwr {
				// filled all the empty fields!
				break doneWithStruct3zbwr
			}
			missingFieldsLeft3zbwr--
			curField3zbwr = unmarshalMsgFieldOrder3zbwr[nextMiss3zbwr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zbwr)
		switch curField3zbwr {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zbwr[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zbwr[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zxrs uint32
				zxrs, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zxrs) {
					z.Fld = (z.Fld)[:zxrs]
				} else {
					z.Fld = make([]FieldT, zxrs)
				}
				for zosz := range z.Fld {
					bts, err = z.Fld[zosz].UnmarshalMsg(bts)
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
	if nextMiss3zbwr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zbwr = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zosz := range z.Fld {
		s += z.Fld[zosz].Msgsize()
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
	const maxFields4zhdr = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zhdr uint32
	totalEncodedFields4zhdr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zhdr := totalEncodedFields4zhdr
	missingFieldsLeft4zhdr := maxFields4zhdr - totalEncodedFields4zhdr

	var nextMiss4zhdr int32 = -1
	var found4zhdr [maxFields4zhdr]bool
	var curField4zhdr string

doneWithStruct4zhdr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zhdr > 0 || missingFieldsLeft4zhdr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zhdr, missingFieldsLeft4zhdr, msgp.ShowFound(found4zhdr[:]), decodeMsgFieldOrder4zhdr)
		if encodedFieldsLeft4zhdr > 0 {
			encodedFieldsLeft4zhdr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zhdr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zhdr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zhdr = 0
			}
			for nextMiss4zhdr < maxFields4zhdr && found4zhdr[nextMiss4zhdr] {
				nextMiss4zhdr++
			}
			if nextMiss4zhdr == maxFields4zhdr {
				// filled all the empty fields!
				break doneWithStruct4zhdr
			}
			missingFieldsLeft4zhdr--
			curField4zhdr = decodeMsgFieldOrder4zhdr[nextMiss4zhdr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zhdr)
		switch curField4zhdr {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zhdr[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zhdr[1] = true
			var zjko uint32
			zjko, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zjko) {
				z.Structs = (z.Structs)[:zjko]
			} else {
				z.Structs = make([]StructT, zjko)
			}
			for zywd := range z.Structs {
				const maxFields5zokg = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zokg uint32
				totalEncodedFields5zokg, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zokg := totalEncodedFields5zokg
				missingFieldsLeft5zokg := maxFields5zokg - totalEncodedFields5zokg

				var nextMiss5zokg int32 = -1
				var found5zokg [maxFields5zokg]bool
				var curField5zokg string

			doneWithStruct5zokg:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zokg > 0 || missingFieldsLeft5zokg > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zokg, missingFieldsLeft5zokg, msgp.ShowFound(found5zokg[:]), decodeMsgFieldOrder5zokg)
					if encodedFieldsLeft5zokg > 0 {
						encodedFieldsLeft5zokg--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zokg = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zokg < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zokg = 0
						}
						for nextMiss5zokg < maxFields5zokg && found5zokg[nextMiss5zokg] {
							nextMiss5zokg++
						}
						if nextMiss5zokg == maxFields5zokg {
							// filled all the empty fields!
							break doneWithStruct5zokg
						}
						missingFieldsLeft5zokg--
						curField5zokg = decodeMsgFieldOrder5zokg[nextMiss5zokg]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zokg)
					switch curField5zokg {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zokg[0] = true
						z.Structs[zywd].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zokg[1] = true
						var zpgb uint32
						zpgb, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zywd].Fld) >= int(zpgb) {
							z.Structs[zywd].Fld = (z.Structs[zywd].Fld)[:zpgb]
						} else {
							z.Structs[zywd].Fld = make([]FieldT, zpgb)
						}
						for zudw := range z.Structs[zywd].Fld {
							err = z.Structs[zywd].Fld[zudw].DecodeMsg(dc)
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
				if nextMiss5zokg != -1 {
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
	if nextMiss4zhdr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zhdr = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zokg = []string{"StructName", "Fld"}

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
	for zywd := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zywd].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zywd].Fld)))
		if err != nil {
			panic(err)
		}
		for zudw := range z.Structs[zywd].Fld {
			err = z.Structs[zywd].Fld[zudw].EncodeMsg(en)
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
	for zywd := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zywd].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zywd].Fld)))
		for zudw := range z.Structs[zywd].Fld {
			o, err = z.Structs[zywd].Fld[zudw].MarshalMsg(o)
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
	const maxFields6ztdr = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6ztdr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6ztdr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6ztdr := totalEncodedFields6ztdr
	missingFieldsLeft6ztdr := maxFields6ztdr - totalEncodedFields6ztdr

	var nextMiss6ztdr int32 = -1
	var found6ztdr [maxFields6ztdr]bool
	var curField6ztdr string

doneWithStruct6ztdr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6ztdr > 0 || missingFieldsLeft6ztdr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6ztdr, missingFieldsLeft6ztdr, msgp.ShowFound(found6ztdr[:]), unmarshalMsgFieldOrder6ztdr)
		if encodedFieldsLeft6ztdr > 0 {
			encodedFieldsLeft6ztdr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6ztdr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6ztdr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6ztdr = 0
			}
			for nextMiss6ztdr < maxFields6ztdr && found6ztdr[nextMiss6ztdr] {
				nextMiss6ztdr++
			}
			if nextMiss6ztdr == maxFields6ztdr {
				// filled all the empty fields!
				break doneWithStruct6ztdr
			}
			missingFieldsLeft6ztdr--
			curField6ztdr = unmarshalMsgFieldOrder6ztdr[nextMiss6ztdr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6ztdr)
		switch curField6ztdr {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6ztdr[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6ztdr[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var zpqz uint32
				zpqz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(zpqz) {
					z.Structs = (z.Structs)[:zpqz]
				} else {
					z.Structs = make([]StructT, zpqz)
				}
				for zywd := range z.Structs {
					const maxFields7zhir = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zhir uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zhir, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zhir := totalEncodedFields7zhir
					missingFieldsLeft7zhir := maxFields7zhir - totalEncodedFields7zhir

					var nextMiss7zhir int32 = -1
					var found7zhir [maxFields7zhir]bool
					var curField7zhir string

				doneWithStruct7zhir:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zhir > 0 || missingFieldsLeft7zhir > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zhir, missingFieldsLeft7zhir, msgp.ShowFound(found7zhir[:]), unmarshalMsgFieldOrder7zhir)
						if encodedFieldsLeft7zhir > 0 {
							encodedFieldsLeft7zhir--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zhir = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zhir < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zhir = 0
							}
							for nextMiss7zhir < maxFields7zhir && found7zhir[nextMiss7zhir] {
								nextMiss7zhir++
							}
							if nextMiss7zhir == maxFields7zhir {
								// filled all the empty fields!
								break doneWithStruct7zhir
							}
							missingFieldsLeft7zhir--
							curField7zhir = unmarshalMsgFieldOrder7zhir[nextMiss7zhir]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zhir)
						switch curField7zhir {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zhir[0] = true
							z.Structs[zywd].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zhir[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zywd].Fld) = (z.Structs[zywd].Fld)[:0]
							} else {

								var zelr uint32
								zelr, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zywd].Fld) >= int(zelr) {
									z.Structs[zywd].Fld = (z.Structs[zywd].Fld)[:zelr]
								} else {
									z.Structs[zywd].Fld = make([]FieldT, zelr)
								}
								for zudw := range z.Structs[zywd].Fld {
									bts, err = z.Structs[zywd].Fld[zudw].UnmarshalMsg(bts)
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
					if nextMiss7zhir != -1 {
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
	if nextMiss6ztdr != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6ztdr = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zhir = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zywd := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zywd].StructName) + 4 + msgp.ArrayHeaderSize
		for zudw := range z.Structs[zywd].Fld {
			s += z.Structs[zywd].Fld[zudw].Msgsize()
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
		var zhxw uint8
		zhxw, err = dc.ReadUint8()
		(*z) = Zprimitive(zhxw)
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
		var zkst uint8
		zkst, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zkst)
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
