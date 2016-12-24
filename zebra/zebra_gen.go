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
	const maxFields0zphd = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zphd uint32
	totalEncodedFields0zphd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zphd := totalEncodedFields0zphd
	missingFieldsLeft0zphd := maxFields0zphd - totalEncodedFields0zphd

	var nextMiss0zphd int32 = -1
	var found0zphd [maxFields0zphd]bool
	var curField0zphd string

doneWithStruct0zphd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zphd > 0 || missingFieldsLeft0zphd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zphd, missingFieldsLeft0zphd, msgp.ShowFound(found0zphd[:]), decodeMsgFieldOrder0zphd)
		if encodedFieldsLeft0zphd > 0 {
			encodedFieldsLeft0zphd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zphd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zphd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zphd = 0
			}
			for nextMiss0zphd < maxFields0zphd && found0zphd[nextMiss0zphd] {
				nextMiss0zphd++
			}
			if nextMiss0zphd == maxFields0zphd {
				// filled all the empty fields!
				break doneWithStruct0zphd
			}
			missingFieldsLeft0zphd--
			curField0zphd = decodeMsgFieldOrder0zphd[nextMiss0zphd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zphd)
		switch curField0zphd {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zphd[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zphd[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zphd[2] = true
			{
				var zlza uint8
				zlza, err = dc.ReadUint8()
				z.Typ = Zprimitive(zlza)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zphd[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zphd[4] = true
			var zqmn uint32
			zqmn, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zqmn > 0 {
				z.Tag = make(map[string]string, zqmn)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zqmn > 0 {
				zqmn--
				var zpws string
				var zrxm string
				zpws, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zrxm, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zpws] = zrxm
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zphd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zphd = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zobm [5]bool
	fieldsInUse_zkni := z.fieldsNotEmpty(empty_zobm[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zkni)
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
	if !empty_zobm[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zpws, zrxm := range z.Tag {
			err = en.WriteString(zpws)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zrxm)
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
		for zpws, zrxm := range z.Tag {
			o = msgp.AppendString(o, zpws)
			o = msgp.AppendString(o, zrxm)
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
	const maxFields1zzzh = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zzzh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zzzh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zzzh := totalEncodedFields1zzzh
	missingFieldsLeft1zzzh := maxFields1zzzh - totalEncodedFields1zzzh

	var nextMiss1zzzh int32 = -1
	var found1zzzh [maxFields1zzzh]bool
	var curField1zzzh string

doneWithStruct1zzzh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zzzh > 0 || missingFieldsLeft1zzzh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zzzh, missingFieldsLeft1zzzh, msgp.ShowFound(found1zzzh[:]), unmarshalMsgFieldOrder1zzzh)
		if encodedFieldsLeft1zzzh > 0 {
			encodedFieldsLeft1zzzh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zzzh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zzzh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zzzh = 0
			}
			for nextMiss1zzzh < maxFields1zzzh && found1zzzh[nextMiss1zzzh] {
				nextMiss1zzzh++
			}
			if nextMiss1zzzh == maxFields1zzzh {
				// filled all the empty fields!
				break doneWithStruct1zzzh
			}
			missingFieldsLeft1zzzh--
			curField1zzzh = unmarshalMsgFieldOrder1zzzh[nextMiss1zzzh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zzzh)
		switch curField1zzzh {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zzzh[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zzzh[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zzzh[2] = true
			{
				var zstt uint8
				zstt, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zstt)
			}
		case "TypStr":
			found1zzzh[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zzzh[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zevt uint32
				zevt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zevt > 0 {
					z.Tag = make(map[string]string, zevt)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zevt > 0 {
					var zpws string
					var zrxm string
					zevt--
					zpws, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zrxm, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zpws] = zrxm
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zzzh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zzzh = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zpws, zrxm := range z.Tag {
			_ = zrxm
			_ = zpws
			s += msgp.StringPrefixSize + len(zpws) + msgp.StringPrefixSize + len(zrxm)
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
	const maxFields2zlbh = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zlbh uint32
	totalEncodedFields2zlbh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zlbh := totalEncodedFields2zlbh
	missingFieldsLeft2zlbh := maxFields2zlbh - totalEncodedFields2zlbh

	var nextMiss2zlbh int32 = -1
	var found2zlbh [maxFields2zlbh]bool
	var curField2zlbh string

doneWithStruct2zlbh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zlbh > 0 || missingFieldsLeft2zlbh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zlbh, missingFieldsLeft2zlbh, msgp.ShowFound(found2zlbh[:]), decodeMsgFieldOrder2zlbh)
		if encodedFieldsLeft2zlbh > 0 {
			encodedFieldsLeft2zlbh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zlbh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zlbh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zlbh = 0
			}
			for nextMiss2zlbh < maxFields2zlbh && found2zlbh[nextMiss2zlbh] {
				nextMiss2zlbh++
			}
			if nextMiss2zlbh == maxFields2zlbh {
				// filled all the empty fields!
				break doneWithStruct2zlbh
			}
			missingFieldsLeft2zlbh--
			curField2zlbh = decodeMsgFieldOrder2zlbh[nextMiss2zlbh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zlbh)
		switch curField2zlbh {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zlbh[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zlbh[1] = true
			var zekr uint32
			zekr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zekr) {
				z.Fld = (z.Fld)[:zekr]
			} else {
				z.Fld = make([]FieldT, zekr)
			}
			for zvze := range z.Fld {
				err = z.Fld[zvze].DecodeMsg(dc)
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
	if nextMiss2zlbh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zlbh = []string{"StructName", "Fld"}

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
	for zvze := range z.Fld {
		err = z.Fld[zvze].EncodeMsg(en)
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
	for zvze := range z.Fld {
		o, err = z.Fld[zvze].MarshalMsg(o)
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
	const maxFields3ztzf = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3ztzf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3ztzf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3ztzf := totalEncodedFields3ztzf
	missingFieldsLeft3ztzf := maxFields3ztzf - totalEncodedFields3ztzf

	var nextMiss3ztzf int32 = -1
	var found3ztzf [maxFields3ztzf]bool
	var curField3ztzf string

doneWithStruct3ztzf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3ztzf > 0 || missingFieldsLeft3ztzf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3ztzf, missingFieldsLeft3ztzf, msgp.ShowFound(found3ztzf[:]), unmarshalMsgFieldOrder3ztzf)
		if encodedFieldsLeft3ztzf > 0 {
			encodedFieldsLeft3ztzf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3ztzf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3ztzf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3ztzf = 0
			}
			for nextMiss3ztzf < maxFields3ztzf && found3ztzf[nextMiss3ztzf] {
				nextMiss3ztzf++
			}
			if nextMiss3ztzf == maxFields3ztzf {
				// filled all the empty fields!
				break doneWithStruct3ztzf
			}
			missingFieldsLeft3ztzf--
			curField3ztzf = unmarshalMsgFieldOrder3ztzf[nextMiss3ztzf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3ztzf)
		switch curField3ztzf {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3ztzf[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3ztzf[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zfyq uint32
				zfyq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zfyq) {
					z.Fld = (z.Fld)[:zfyq]
				} else {
					z.Fld = make([]FieldT, zfyq)
				}
				for zvze := range z.Fld {
					bts, err = z.Fld[zvze].UnmarshalMsg(bts)
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
	if nextMiss3ztzf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3ztzf = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zvze := range z.Fld {
		s += z.Fld[zvze].Msgsize()
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
	const maxFields4zvyj = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zvyj uint32
	totalEncodedFields4zvyj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zvyj := totalEncodedFields4zvyj
	missingFieldsLeft4zvyj := maxFields4zvyj - totalEncodedFields4zvyj

	var nextMiss4zvyj int32 = -1
	var found4zvyj [maxFields4zvyj]bool
	var curField4zvyj string

doneWithStruct4zvyj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zvyj > 0 || missingFieldsLeft4zvyj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zvyj, missingFieldsLeft4zvyj, msgp.ShowFound(found4zvyj[:]), decodeMsgFieldOrder4zvyj)
		if encodedFieldsLeft4zvyj > 0 {
			encodedFieldsLeft4zvyj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zvyj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zvyj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zvyj = 0
			}
			for nextMiss4zvyj < maxFields4zvyj && found4zvyj[nextMiss4zvyj] {
				nextMiss4zvyj++
			}
			if nextMiss4zvyj == maxFields4zvyj {
				// filled all the empty fields!
				break doneWithStruct4zvyj
			}
			missingFieldsLeft4zvyj--
			curField4zvyj = decodeMsgFieldOrder4zvyj[nextMiss4zvyj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zvyj)
		switch curField4zvyj {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zvyj[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zvyj[1] = true
			var zqpv uint32
			zqpv, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(zqpv) {
				z.Structs = (z.Structs)[:zqpv]
			} else {
				z.Structs = make([]StructT, zqpv)
			}
			for zgqh := range z.Structs {
				const maxFields5zntn = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zntn uint32
				totalEncodedFields5zntn, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zntn := totalEncodedFields5zntn
				missingFieldsLeft5zntn := maxFields5zntn - totalEncodedFields5zntn

				var nextMiss5zntn int32 = -1
				var found5zntn [maxFields5zntn]bool
				var curField5zntn string

			doneWithStruct5zntn:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zntn > 0 || missingFieldsLeft5zntn > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zntn, missingFieldsLeft5zntn, msgp.ShowFound(found5zntn[:]), decodeMsgFieldOrder5zntn)
					if encodedFieldsLeft5zntn > 0 {
						encodedFieldsLeft5zntn--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zntn = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zntn < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zntn = 0
						}
						for nextMiss5zntn < maxFields5zntn && found5zntn[nextMiss5zntn] {
							nextMiss5zntn++
						}
						if nextMiss5zntn == maxFields5zntn {
							// filled all the empty fields!
							break doneWithStruct5zntn
						}
						missingFieldsLeft5zntn--
						curField5zntn = decodeMsgFieldOrder5zntn[nextMiss5zntn]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zntn)
					switch curField5zntn {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zntn[0] = true
						z.Structs[zgqh].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zntn[1] = true
						var zxja uint32
						zxja, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zgqh].Fld) >= int(zxja) {
							z.Structs[zgqh].Fld = (z.Structs[zgqh].Fld)[:zxja]
						} else {
							z.Structs[zgqh].Fld = make([]FieldT, zxja)
						}
						for zdgw := range z.Structs[zgqh].Fld {
							err = z.Structs[zgqh].Fld[zdgw].DecodeMsg(dc)
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
				if nextMiss5zntn != -1 {
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
	if nextMiss4zvyj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zvyj = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zntn = []string{"StructName", "Fld"}

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
	for zgqh := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zgqh].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zgqh].Fld)))
		if err != nil {
			panic(err)
		}
		for zdgw := range z.Structs[zgqh].Fld {
			err = z.Structs[zgqh].Fld[zdgw].EncodeMsg(en)
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
	for zgqh := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zgqh].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zgqh].Fld)))
		for zdgw := range z.Structs[zgqh].Fld {
			o, err = z.Structs[zgqh].Fld[zdgw].MarshalMsg(o)
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
	const maxFields6zfed = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zfed uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zfed, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zfed := totalEncodedFields6zfed
	missingFieldsLeft6zfed := maxFields6zfed - totalEncodedFields6zfed

	var nextMiss6zfed int32 = -1
	var found6zfed [maxFields6zfed]bool
	var curField6zfed string

doneWithStruct6zfed:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zfed > 0 || missingFieldsLeft6zfed > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zfed, missingFieldsLeft6zfed, msgp.ShowFound(found6zfed[:]), unmarshalMsgFieldOrder6zfed)
		if encodedFieldsLeft6zfed > 0 {
			encodedFieldsLeft6zfed--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zfed = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zfed < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zfed = 0
			}
			for nextMiss6zfed < maxFields6zfed && found6zfed[nextMiss6zfed] {
				nextMiss6zfed++
			}
			if nextMiss6zfed == maxFields6zfed {
				// filled all the empty fields!
				break doneWithStruct6zfed
			}
			missingFieldsLeft6zfed--
			curField6zfed = unmarshalMsgFieldOrder6zfed[nextMiss6zfed]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zfed)
		switch curField6zfed {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zfed[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zfed[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var ztpq uint32
				ztpq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(ztpq) {
					z.Structs = (z.Structs)[:ztpq]
				} else {
					z.Structs = make([]StructT, ztpq)
				}
				for zgqh := range z.Structs {
					const maxFields7zrox = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zrox uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zrox, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zrox := totalEncodedFields7zrox
					missingFieldsLeft7zrox := maxFields7zrox - totalEncodedFields7zrox

					var nextMiss7zrox int32 = -1
					var found7zrox [maxFields7zrox]bool
					var curField7zrox string

				doneWithStruct7zrox:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zrox > 0 || missingFieldsLeft7zrox > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zrox, missingFieldsLeft7zrox, msgp.ShowFound(found7zrox[:]), unmarshalMsgFieldOrder7zrox)
						if encodedFieldsLeft7zrox > 0 {
							encodedFieldsLeft7zrox--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zrox = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zrox < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zrox = 0
							}
							for nextMiss7zrox < maxFields7zrox && found7zrox[nextMiss7zrox] {
								nextMiss7zrox++
							}
							if nextMiss7zrox == maxFields7zrox {
								// filled all the empty fields!
								break doneWithStruct7zrox
							}
							missingFieldsLeft7zrox--
							curField7zrox = unmarshalMsgFieldOrder7zrox[nextMiss7zrox]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zrox)
						switch curField7zrox {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zrox[0] = true
							z.Structs[zgqh].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zrox[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zgqh].Fld) = (z.Structs[zgqh].Fld)[:0]
							} else {

								var zepd uint32
								zepd, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zgqh].Fld) >= int(zepd) {
									z.Structs[zgqh].Fld = (z.Structs[zgqh].Fld)[:zepd]
								} else {
									z.Structs[zgqh].Fld = make([]FieldT, zepd)
								}
								for zdgw := range z.Structs[zgqh].Fld {
									bts, err = z.Structs[zgqh].Fld[zdgw].UnmarshalMsg(bts)
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
					if nextMiss7zrox != -1 {
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
	if nextMiss6zfed != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zfed = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zrox = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zgqh := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zgqh].StructName) + 4 + msgp.ArrayHeaderSize
		for zdgw := range z.Structs[zgqh].Fld {
			s += z.Structs[zgqh].Fld[zdgw].Msgsize()
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
		var zuiz uint8
		zuiz, err = dc.ReadUint8()
		(*z) = Zprimitive(zuiz)
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
		var znkj uint8
		znkj, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(znkj)
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
