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
	const maxFields0zptk = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zptk uint32
	totalEncodedFields0zptk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zptk := totalEncodedFields0zptk
	missingFieldsLeft0zptk := maxFields0zptk - totalEncodedFields0zptk

	var nextMiss0zptk int32 = -1
	var found0zptk [maxFields0zptk]bool
	var curField0zptk string

doneWithStruct0zptk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zptk > 0 || missingFieldsLeft0zptk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zptk, missingFieldsLeft0zptk, msgp.ShowFound(found0zptk[:]), decodeMsgFieldOrder0zptk)
		if encodedFieldsLeft0zptk > 0 {
			encodedFieldsLeft0zptk--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zptk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zptk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zptk = 0
			}
			for nextMiss0zptk < maxFields0zptk && found0zptk[nextMiss0zptk] {
				nextMiss0zptk++
			}
			if nextMiss0zptk == maxFields0zptk {
				// filled all the empty fields!
				break doneWithStruct0zptk
			}
			missingFieldsLeft0zptk--
			curField0zptk = decodeMsgFieldOrder0zptk[nextMiss0zptk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zptk)
		switch curField0zptk {
		// -- templateDecodeMsg ends here --

		case "Zid":
			found0zptk[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Nam":
			found0zptk[1] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Typ":
			found0zptk[2] = true
			{
				var zmal uint8
				zmal, err = dc.ReadUint8()
				z.Typ = Zprimitive(zmal)
			}
			if err != nil {
				panic(err)
			}
		case "TypStr":
			found0zptk[3] = true
			z.TypStr, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zptk[4] = true
			var zknq uint32
			zknq, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zknq > 0 {
				z.Tag = make(map[string]string, zknq)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zknq > 0 {
				zknq--
				var zonw string
				var zmrt string
				zonw, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zmrt, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zonw] = zmrt
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zptk != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zptk = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

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
	var empty_zeer [5]bool
	fieldsInUse_zyjk := z.fieldsNotEmpty(empty_zeer[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zyjk)
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
	if !empty_zeer[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zonw, zmrt := range z.Tag {
			err = en.WriteString(zonw)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zmrt)
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
		for zonw, zmrt := range z.Tag {
			o = msgp.AppendString(o, zonw)
			o = msgp.AppendString(o, zmrt)
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
	const maxFields1zmjn = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zmjn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zmjn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zmjn := totalEncodedFields1zmjn
	missingFieldsLeft1zmjn := maxFields1zmjn - totalEncodedFields1zmjn

	var nextMiss1zmjn int32 = -1
	var found1zmjn [maxFields1zmjn]bool
	var curField1zmjn string

doneWithStruct1zmjn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zmjn > 0 || missingFieldsLeft1zmjn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zmjn, missingFieldsLeft1zmjn, msgp.ShowFound(found1zmjn[:]), unmarshalMsgFieldOrder1zmjn)
		if encodedFieldsLeft1zmjn > 0 {
			encodedFieldsLeft1zmjn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zmjn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zmjn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zmjn = 0
			}
			for nextMiss1zmjn < maxFields1zmjn && found1zmjn[nextMiss1zmjn] {
				nextMiss1zmjn++
			}
			if nextMiss1zmjn == maxFields1zmjn {
				// filled all the empty fields!
				break doneWithStruct1zmjn
			}
			missingFieldsLeft1zmjn--
			curField1zmjn = unmarshalMsgFieldOrder1zmjn[nextMiss1zmjn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zmjn)
		switch curField1zmjn {
		// -- templateUnmarshalMsg ends here --

		case "Zid":
			found1zmjn[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Nam":
			found1zmjn[1] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Typ":
			found1zmjn[2] = true
			{
				var zcgt uint8
				zcgt, bts, err = nbs.ReadUint8Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Typ = Zprimitive(zcgt)
			}
		case "TypStr":
			found1zmjn[3] = true
			z.TypStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zmjn[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zbzf uint32
				zbzf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zbzf > 0 {
					z.Tag = make(map[string]string, zbzf)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zbzf > 0 {
					var zonw string
					var zmrt string
					zbzf--
					zonw, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zmrt, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zonw] = zmrt
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zmjn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zmjn = []string{"Zid", "Nam", "Typ", "TypStr", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.Uint8Size + 7 + msgp.StringPrefixSize + len(z.TypStr) + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zonw, zmrt := range z.Tag {
			_ = zmrt
			_ = zonw
			s += msgp.StringPrefixSize + len(zonw) + msgp.StringPrefixSize + len(zmrt)
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
	const maxFields2zlvx = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zlvx uint32
	totalEncodedFields2zlvx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zlvx := totalEncodedFields2zlvx
	missingFieldsLeft2zlvx := maxFields2zlvx - totalEncodedFields2zlvx

	var nextMiss2zlvx int32 = -1
	var found2zlvx [maxFields2zlvx]bool
	var curField2zlvx string

doneWithStruct2zlvx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zlvx > 0 || missingFieldsLeft2zlvx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zlvx, missingFieldsLeft2zlvx, msgp.ShowFound(found2zlvx[:]), decodeMsgFieldOrder2zlvx)
		if encodedFieldsLeft2zlvx > 0 {
			encodedFieldsLeft2zlvx--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zlvx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zlvx < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zlvx = 0
			}
			for nextMiss2zlvx < maxFields2zlvx && found2zlvx[nextMiss2zlvx] {
				nextMiss2zlvx++
			}
			if nextMiss2zlvx == maxFields2zlvx {
				// filled all the empty fields!
				break doneWithStruct2zlvx
			}
			missingFieldsLeft2zlvx--
			curField2zlvx = decodeMsgFieldOrder2zlvx[nextMiss2zlvx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zlvx)
		switch curField2zlvx {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found2zlvx[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found2zlvx[1] = true
			var zqin uint32
			zqin, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zqin) {
				z.Fld = (z.Fld)[:zqin]
			} else {
				z.Fld = make([]FieldT, zqin)
			}
			for zyej := range z.Fld {
				err = z.Fld[zyej].DecodeMsg(dc)
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
	if nextMiss2zlvx != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder2zlvx = []string{"StructName", "Fld"}

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
	for zyej := range z.Fld {
		err = z.Fld[zyej].EncodeMsg(en)
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
	for zyej := range z.Fld {
		o, err = z.Fld[zyej].MarshalMsg(o)
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
	const maxFields3zbrv = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zbrv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zbrv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zbrv := totalEncodedFields3zbrv
	missingFieldsLeft3zbrv := maxFields3zbrv - totalEncodedFields3zbrv

	var nextMiss3zbrv int32 = -1
	var found3zbrv [maxFields3zbrv]bool
	var curField3zbrv string

doneWithStruct3zbrv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zbrv > 0 || missingFieldsLeft3zbrv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zbrv, missingFieldsLeft3zbrv, msgp.ShowFound(found3zbrv[:]), unmarshalMsgFieldOrder3zbrv)
		if encodedFieldsLeft3zbrv > 0 {
			encodedFieldsLeft3zbrv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zbrv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zbrv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zbrv = 0
			}
			for nextMiss3zbrv < maxFields3zbrv && found3zbrv[nextMiss3zbrv] {
				nextMiss3zbrv++
			}
			if nextMiss3zbrv == maxFields3zbrv {
				// filled all the empty fields!
				break doneWithStruct3zbrv
			}
			missingFieldsLeft3zbrv--
			curField3zbrv = unmarshalMsgFieldOrder3zbrv[nextMiss3zbrv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zbrv)
		switch curField3zbrv {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found3zbrv[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found3zbrv[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zhpa uint32
				zhpa, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zhpa) {
					z.Fld = (z.Fld)[:zhpa]
				} else {
					z.Fld = make([]FieldT, zhpa)
				}
				for zyej := range z.Fld {
					bts, err = z.Fld[zyej].UnmarshalMsg(bts)
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
	if nextMiss3zbrv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder3zbrv = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zyej := range z.Fld {
		s += z.Fld[zyej].Msgsize()
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
	const maxFields4zgrs = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zgrs uint32
	totalEncodedFields4zgrs, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zgrs := totalEncodedFields4zgrs
	missingFieldsLeft4zgrs := maxFields4zgrs - totalEncodedFields4zgrs

	var nextMiss4zgrs int32 = -1
	var found4zgrs [maxFields4zgrs]bool
	var curField4zgrs string

doneWithStruct4zgrs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zgrs > 0 || missingFieldsLeft4zgrs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zgrs, missingFieldsLeft4zgrs, msgp.ShowFound(found4zgrs[:]), decodeMsgFieldOrder4zgrs)
		if encodedFieldsLeft4zgrs > 0 {
			encodedFieldsLeft4zgrs--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zgrs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zgrs < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zgrs = 0
			}
			for nextMiss4zgrs < maxFields4zgrs && found4zgrs[nextMiss4zgrs] {
				nextMiss4zgrs++
			}
			if nextMiss4zgrs == maxFields4zgrs {
				// filled all the empty fields!
				break doneWithStruct4zgrs
			}
			missingFieldsLeft4zgrs--
			curField4zgrs = decodeMsgFieldOrder4zgrs[nextMiss4zgrs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zgrs)
		switch curField4zgrs {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found4zgrs[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found4zgrs[1] = true
			var ztzd uint32
			ztzd, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Structs) >= int(ztzd) {
				z.Structs = (z.Structs)[:ztzd]
			} else {
				z.Structs = make([]StructT, ztzd)
			}
			for zzbo := range z.Structs {
				const maxFields5zxtr = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields5zxtr uint32
				totalEncodedFields5zxtr, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft5zxtr := totalEncodedFields5zxtr
				missingFieldsLeft5zxtr := maxFields5zxtr - totalEncodedFields5zxtr

				var nextMiss5zxtr int32 = -1
				var found5zxtr [maxFields5zxtr]bool
				var curField5zxtr string

			doneWithStruct5zxtr:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft5zxtr > 0 || missingFieldsLeft5zxtr > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxtr, missingFieldsLeft5zxtr, msgp.ShowFound(found5zxtr[:]), decodeMsgFieldOrder5zxtr)
					if encodedFieldsLeft5zxtr > 0 {
						encodedFieldsLeft5zxtr--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField5zxtr = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss5zxtr < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss5zxtr = 0
						}
						for nextMiss5zxtr < maxFields5zxtr && found5zxtr[nextMiss5zxtr] {
							nextMiss5zxtr++
						}
						if nextMiss5zxtr == maxFields5zxtr {
							// filled all the empty fields!
							break doneWithStruct5zxtr
						}
						missingFieldsLeft5zxtr--
						curField5zxtr = decodeMsgFieldOrder5zxtr[nextMiss5zxtr]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField5zxtr)
					switch curField5zxtr {
					// -- templateDecodeMsg ends here --

					case "StructName":
						found5zxtr[0] = true
						z.Structs[zzbo].StructName, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found5zxtr[1] = true
						var zrev uint32
						zrev, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(z.Structs[zzbo].Fld) >= int(zrev) {
							z.Structs[zzbo].Fld = (z.Structs[zzbo].Fld)[:zrev]
						} else {
							z.Structs[zzbo].Fld = make([]FieldT, zrev)
						}
						for zqsv := range z.Structs[zzbo].Fld {
							err = z.Structs[zzbo].Fld[zqsv].DecodeMsg(dc)
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
				if nextMiss5zxtr != -1 {
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
	if nextMiss4zgrs != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZebraSchema
var decodeMsgFieldOrder4zgrs = []string{"PkgPath", "Structs"}

// fields of StructT
var decodeMsgFieldOrder5zxtr = []string{"StructName", "Fld"}

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
	for zzbo := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zzbo].StructName)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zzbo].Fld)))
		if err != nil {
			panic(err)
		}
		for zqsv := range z.Structs[zzbo].Fld {
			err = z.Structs[zzbo].Fld[zqsv].EncodeMsg(en)
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
	for zzbo := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zzbo].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zzbo].Fld)))
		for zqsv := range z.Structs[zzbo].Fld {
			o, err = z.Structs[zzbo].Fld[zqsv].MarshalMsg(o)
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
	const maxFields6zoqp = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields6zoqp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields6zoqp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft6zoqp := totalEncodedFields6zoqp
	missingFieldsLeft6zoqp := maxFields6zoqp - totalEncodedFields6zoqp

	var nextMiss6zoqp int32 = -1
	var found6zoqp [maxFields6zoqp]bool
	var curField6zoqp string

doneWithStruct6zoqp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zoqp > 0 || missingFieldsLeft6zoqp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zoqp, missingFieldsLeft6zoqp, msgp.ShowFound(found6zoqp[:]), unmarshalMsgFieldOrder6zoqp)
		if encodedFieldsLeft6zoqp > 0 {
			encodedFieldsLeft6zoqp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField6zoqp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zoqp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss6zoqp = 0
			}
			for nextMiss6zoqp < maxFields6zoqp && found6zoqp[nextMiss6zoqp] {
				nextMiss6zoqp++
			}
			if nextMiss6zoqp == maxFields6zoqp {
				// filled all the empty fields!
				break doneWithStruct6zoqp
			}
			missingFieldsLeft6zoqp--
			curField6zoqp = unmarshalMsgFieldOrder6zoqp[nextMiss6zoqp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zoqp)
		switch curField6zoqp {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found6zoqp[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found6zoqp[1] = true
			if nbs.AlwaysNil {
				(z.Structs) = (z.Structs)[:0]
			} else {

				var ziry uint32
				ziry, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Structs) >= int(ziry) {
					z.Structs = (z.Structs)[:ziry]
				} else {
					z.Structs = make([]StructT, ziry)
				}
				for zzbo := range z.Structs {
					const maxFields7zwml = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields7zwml uint32
					if !nbs.AlwaysNil {
						totalEncodedFields7zwml, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft7zwml := totalEncodedFields7zwml
					missingFieldsLeft7zwml := maxFields7zwml - totalEncodedFields7zwml

					var nextMiss7zwml int32 = -1
					var found7zwml [maxFields7zwml]bool
					var curField7zwml string

				doneWithStruct7zwml:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft7zwml > 0 || missingFieldsLeft7zwml > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zwml, missingFieldsLeft7zwml, msgp.ShowFound(found7zwml[:]), unmarshalMsgFieldOrder7zwml)
						if encodedFieldsLeft7zwml > 0 {
							encodedFieldsLeft7zwml--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField7zwml = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss7zwml < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss7zwml = 0
							}
							for nextMiss7zwml < maxFields7zwml && found7zwml[nextMiss7zwml] {
								nextMiss7zwml++
							}
							if nextMiss7zwml == maxFields7zwml {
								// filled all the empty fields!
								break doneWithStruct7zwml
							}
							missingFieldsLeft7zwml--
							curField7zwml = unmarshalMsgFieldOrder7zwml[nextMiss7zwml]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField7zwml)
						switch curField7zwml {
						// -- templateUnmarshalMsg ends here --

						case "StructName":
							found7zwml[0] = true
							z.Structs[zzbo].StructName, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found7zwml[1] = true
							if nbs.AlwaysNil {
								(z.Structs[zzbo].Fld) = (z.Structs[zzbo].Fld)[:0]
							} else {

								var zuuo uint32
								zuuo, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(z.Structs[zzbo].Fld) >= int(zuuo) {
									z.Structs[zzbo].Fld = (z.Structs[zzbo].Fld)[:zuuo]
								} else {
									z.Structs[zzbo].Fld = make([]FieldT, zuuo)
								}
								for zqsv := range z.Structs[zzbo].Fld {
									bts, err = z.Structs[zzbo].Fld[zqsv].UnmarshalMsg(bts)
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
					if nextMiss7zwml != -1 {
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
	if nextMiss6zoqp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZebraSchema
var unmarshalMsgFieldOrder6zoqp = []string{"PkgPath", "Structs"}

// fields of StructT
var unmarshalMsgFieldOrder7zwml = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZebraSchema) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zzbo := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zzbo].StructName) + 4 + msgp.ArrayHeaderSize
		for zqsv := range z.Structs[zzbo].Fld {
			s += z.Structs[zzbo].Fld[zqsv].Msgsize()
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
		var zuvv uint8
		zuvv, err = dc.ReadUint8()
		(*z) = Zprimitive(zuvv)
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
		var zowh uint8
		zowh, bts, err = nbs.ReadUint8Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zprimitive(zowh)
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
