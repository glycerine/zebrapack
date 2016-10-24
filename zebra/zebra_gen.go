package zebra

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/glycerine/zebrapack)
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
	const maxFields0zaaa = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zaaa uint32
	totalEncodedFields0zaaa, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zaaa := totalEncodedFields0zaaa
	missingFieldsLeft0zaaa := maxFields0zaaa - totalEncodedFields0zaaa

	var nextMiss0zaaa int32 = -1
	var found0zaaa [maxFields0zaaa]bool
	var curField0zaaa string

doneWithStruct0zaaa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaaa > 0 || missingFieldsLeft0zaaa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaaa, missingFieldsLeft0zaaa, msgp.ShowFound(found0zaaa[:]), decodeMsgFieldOrder0zaaa)
		if encodedFieldsLeft0zaaa > 0 {
			encodedFieldsLeft0zaaa--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zaaa = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zaaa < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zaaa = 0
			}
			for nextMiss0zaaa < maxFields0zaaa && found0zaaa[nextMiss0zaaa] {
				nextMiss0zaaa++
			}
			if nextMiss0zaaa == maxFields0zaaa {
				// filled all the empty fields!
				break doneWithStruct0zaaa
			}
			missingFieldsLeft0zaaa--
			curField0zaaa = decodeMsgFieldOrder0zaaa[nextMiss0zaaa]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaaa)
		switch curField0zaaa {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zaaa[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zaaa[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found0zaaa[2] = true
			{
				var zjco int32
				zjco, err = dc.ReadInt32()
				z.Ztyp = Ztype(zjco)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0zaaa[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zaaa[4] = true
			var zkff uint32
			zkff, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zkff > 0 {
				z.Tag = make(map[string]string, zkff)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zkff > 0 {
				zkff--
				var zkht string
				var zaht string
				zkht, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zaht, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zkht] = zaht
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zaaa != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zaaa = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

// fieldsNotEmpty supports omitempty tags
func (z *FieldT) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[3] = (!z.Varg) // bool, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.Tag) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *FieldT) EncodeMsg(en *msgp.Writer) (err error) {

	// honor the omitempty tags
	var empty_zkwz [5]bool
	fieldsInUse_zirk := z.fieldsNotEmpty(empty_zkwz[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zirk)
	if err != nil {
		return err
	}

	// write "FieldId"
	err = en.Append(0xa7, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.FieldId)
	if err != nil {
		panic(err)
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		panic(err)
	}
	// write "Ztyp"
	err = en.Append(0xa4, 0x5a, 0x74, 0x79, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Ztyp))
	if err != nil {
		panic(err)
	}
	if !empty_zkwz[3] {
		// write "Varg"
		err = en.Append(0xa4, 0x56, 0x61, 0x72, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Varg)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zkwz[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zkht, zaht := range z.Tag {
			err = en.WriteString(zkht)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zaht)
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

	// string "FieldId"
	o = append(o, 0xa7, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.FieldId)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Ztyp"
	o = append(o, 0xa4, 0x5a, 0x74, 0x79, 0x70)
	o = msgp.AppendInt32(o, int32(z.Ztyp))
	if !empty[3] {
		// string "Varg"
		o = append(o, 0xa4, 0x56, 0x61, 0x72, 0x67)
		o = msgp.AppendBool(o, z.Varg)
	}

	if !empty[4] {
		// string "Tag"
		o = append(o, 0xa3, 0x54, 0x61, 0x67)
		o = msgp.AppendMapHeader(o, uint32(len(z.Tag)))
		for zkht, zaht := range z.Tag {
			o = msgp.AppendString(o, zkht)
			o = msgp.AppendString(o, zaht)
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
	const maxFields1zuqk = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zuqk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zuqk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zuqk := totalEncodedFields1zuqk
	missingFieldsLeft1zuqk := maxFields1zuqk - totalEncodedFields1zuqk

	var nextMiss1zuqk int32 = -1
	var found1zuqk [maxFields1zuqk]bool
	var curField1zuqk string

doneWithStruct1zuqk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zuqk > 0 || missingFieldsLeft1zuqk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zuqk, missingFieldsLeft1zuqk, msgp.ShowFound(found1zuqk[:]), unmarshalMsgFieldOrder1zuqk)
		if encodedFieldsLeft1zuqk > 0 {
			encodedFieldsLeft1zuqk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zuqk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zuqk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zuqk = 0
			}
			for nextMiss1zuqk < maxFields1zuqk && found1zuqk[nextMiss1zuqk] {
				nextMiss1zuqk++
			}
			if nextMiss1zuqk == maxFields1zuqk {
				// filled all the empty fields!
				break doneWithStruct1zuqk
			}
			missingFieldsLeft1zuqk--
			curField1zuqk = unmarshalMsgFieldOrder1zuqk[nextMiss1zuqk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zuqk)
		switch curField1zuqk {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zuqk[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zuqk[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found1zuqk[2] = true
			{
				var zqdi int32
				zqdi, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ztyp = Ztype(zqdi)
			}
		case "Varg":
			found1zuqk[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zuqk[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zxfj uint32
				zxfj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zxfj > 0 {
					z.Tag = make(map[string]string, zxfj)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zxfj > 0 {
					var zkht string
					var zaht string
					zxfj--
					zkht, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zaht, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zkht] = zaht
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zuqk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zuqk = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zkht, zaht := range z.Tag {
			_ = zaht
			s += msgp.StringPrefixSize + len(zkht) + msgp.StringPrefixSize + len(zaht)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *InterfaceInstance) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zllp = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zllp uint32
	totalEncodedFields2zllp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zllp := totalEncodedFields2zllp
	missingFieldsLeft2zllp := maxFields2zllp - totalEncodedFields2zllp

	var nextMiss2zllp int32 = -1
	var found2zllp [maxFields2zllp]bool
	var curField2zllp string

doneWithStruct2zllp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zllp > 0 || missingFieldsLeft2zllp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zllp, missingFieldsLeft2zllp, msgp.ShowFound(found2zllp[:]), decodeMsgFieldOrder2zllp)
		if encodedFieldsLeft2zllp > 0 {
			encodedFieldsLeft2zllp--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zllp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zllp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zllp = 0
			}
			for nextMiss2zllp < maxFields2zllp && found2zllp[nextMiss2zllp] {
				nextMiss2zllp++
			}
			if nextMiss2zllp == maxFields2zllp {
				// filled all the empty fields!
				break doneWithStruct2zllp
			}
			missingFieldsLeft2zllp--
			curField2zllp = decodeMsgFieldOrder2zllp[nextMiss2zllp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zllp)
		switch curField2zllp {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zllp[0] = true
			z.IfaceId, err = dc.ReadInt64()
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
	if nextMiss2zllp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zllp = []string{"IfaceId"}

// fieldsNotEmpty supports omitempty tags
func (z InterfaceInstance) fieldsNotEmpty(isempty []bool) uint32 {
	return 1
}

// EncodeMsg implements msgp.Encodable
func (z InterfaceInstance) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "IfaceId"
	err = en.Append(0x81, 0xa7, 0x49, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.IfaceId)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z InterfaceInstance) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "IfaceId"
	o = append(o, 0x81, 0xa7, 0x49, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.IfaceId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InterfaceInstance) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *InterfaceInstance) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zlea = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zlea uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zlea, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zlea := totalEncodedFields3zlea
	missingFieldsLeft3zlea := maxFields3zlea - totalEncodedFields3zlea

	var nextMiss3zlea int32 = -1
	var found3zlea [maxFields3zlea]bool
	var curField3zlea string

doneWithStruct3zlea:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zlea > 0 || missingFieldsLeft3zlea > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zlea, missingFieldsLeft3zlea, msgp.ShowFound(found3zlea[:]), unmarshalMsgFieldOrder3zlea)
		if encodedFieldsLeft3zlea > 0 {
			encodedFieldsLeft3zlea--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zlea = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zlea < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zlea = 0
			}
			for nextMiss3zlea < maxFields3zlea && found3zlea[nextMiss3zlea] {
				nextMiss3zlea++
			}
			if nextMiss3zlea == maxFields3zlea {
				// filled all the empty fields!
				break doneWithStruct3zlea
			}
			missingFieldsLeft3zlea--
			curField3zlea = unmarshalMsgFieldOrder3zlea[nextMiss3zlea]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zlea)
		switch curField3zlea {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3zlea[0] = true
			z.IfaceId, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss3zlea != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3zlea = []string{"IfaceId"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z InterfaceInstance) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *InterfaceT) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields4zgdz = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zgdz uint32
	totalEncodedFields4zgdz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zgdz := totalEncodedFields4zgdz
	missingFieldsLeft4zgdz := maxFields4zgdz - totalEncodedFields4zgdz

	var nextMiss4zgdz int32 = -1
	var found4zgdz [maxFields4zgdz]bool
	var curField4zgdz string

doneWithStruct4zgdz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zgdz > 0 || missingFieldsLeft4zgdz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zgdz, missingFieldsLeft4zgdz, msgp.ShowFound(found4zgdz[:]), decodeMsgFieldOrder4zgdz)
		if encodedFieldsLeft4zgdz > 0 {
			encodedFieldsLeft4zgdz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zgdz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zgdz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zgdz = 0
			}
			for nextMiss4zgdz < maxFields4zgdz && found4zgdz[nextMiss4zgdz] {
				nextMiss4zgdz++
			}
			if nextMiss4zgdz == maxFields4zgdz {
				// filled all the empty fields!
				break doneWithStruct4zgdz
			}
			missingFieldsLeft4zgdz--
			curField4zgdz = decodeMsgFieldOrder4zgdz[nextMiss4zgdz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zgdz)
		switch curField4zgdz {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4zgdz[0] = true
			var zziu uint32
			zziu, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zziu) {
				z.Methods = (z.Methods)[:zziu]
			} else {
				z.Methods = make([]MethodT, zziu)
			}
			for zydj := range z.Methods {
				err = z.Methods[zydj].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4zgdz[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zgdz[2] = true
			z.Comment, err = dc.ReadString()
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
	if nextMiss4zgdz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4zgdz = []string{"Methods", "Deprecated", "Comment"}

// fieldsNotEmpty supports omitempty tags
func (z *InterfaceT) fieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// EncodeMsg implements msgp.Encodable
func (z *InterfaceT) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Methods"
	err = en.Append(0x83, 0xa7, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Methods)))
	if err != nil {
		panic(err)
	}
	for zydj := range z.Methods {
		err = z.Methods[zydj].EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	// write "Deprecated"
	err = en.Append(0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Deprecated)
	if err != nil {
		panic(err)
	}
	// write "Comment"
	err = en.Append(0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Comment)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *InterfaceT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Methods"
	o = append(o, 0x83, 0xa7, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Methods)))
	for zydj := range z.Methods {
		o, err = z.Methods[zydj].MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Deprecated"
	o = append(o, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Deprecated)
	// string "Comment"
	o = append(o, 0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Comment)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InterfaceT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *InterfaceT) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields5zcer = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zcer uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zcer, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zcer := totalEncodedFields5zcer
	missingFieldsLeft5zcer := maxFields5zcer - totalEncodedFields5zcer

	var nextMiss5zcer int32 = -1
	var found5zcer [maxFields5zcer]bool
	var curField5zcer string

doneWithStruct5zcer:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zcer > 0 || missingFieldsLeft5zcer > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zcer, missingFieldsLeft5zcer, msgp.ShowFound(found5zcer[:]), unmarshalMsgFieldOrder5zcer)
		if encodedFieldsLeft5zcer > 0 {
			encodedFieldsLeft5zcer--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zcer = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zcer < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zcer = 0
			}
			for nextMiss5zcer < maxFields5zcer && found5zcer[nextMiss5zcer] {
				nextMiss5zcer++
			}
			if nextMiss5zcer == maxFields5zcer {
				// filled all the empty fields!
				break doneWithStruct5zcer
			}
			missingFieldsLeft5zcer--
			curField5zcer = unmarshalMsgFieldOrder5zcer[nextMiss5zcer]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zcer)
		switch curField5zcer {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zcer[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zmjx uint32
				zmjx, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zmjx) {
					z.Methods = (z.Methods)[:zmjx]
				} else {
					z.Methods = make([]MethodT, zmjx)
				}
				for zydj := range z.Methods {
					bts, err = z.Methods[zydj].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zcer[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zcer[2] = true
			z.Comment, bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss5zcer != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zcer = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zydj := range z.Methods {
		s += z.Methods[zydj].Msgsize()
	}
	s += 11 + msgp.BoolSize + 8 + msgp.StringPrefixSize + len(z.Comment)
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *MethodT) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields6zogt = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zogt uint32
	totalEncodedFields6zogt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zogt := totalEncodedFields6zogt
	missingFieldsLeft6zogt := maxFields6zogt - totalEncodedFields6zogt

	var nextMiss6zogt int32 = -1
	var found6zogt [maxFields6zogt]bool
	var curField6zogt string

doneWithStruct6zogt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zogt > 0 || missingFieldsLeft6zogt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zogt, missingFieldsLeft6zogt, msgp.ShowFound(found6zogt[:]), decodeMsgFieldOrder6zogt)
		if encodedFieldsLeft6zogt > 0 {
			encodedFieldsLeft6zogt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zogt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zogt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zogt = 0
			}
			for nextMiss6zogt < maxFields6zogt && found6zogt[nextMiss6zogt] {
				nextMiss6zogt++
			}
			if nextMiss6zogt == maxFields6zogt {
				// filled all the empty fields!
				break doneWithStruct6zogt
			}
			missingFieldsLeft6zogt--
			curField6zogt = decodeMsgFieldOrder6zogt[nextMiss6zogt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zogt)
		switch curField6zogt {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zogt[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zogt[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zogt[2] = true
			const maxFields7zhzg = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7zhzg uint32
			totalEncodedFields7zhzg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7zhzg := totalEncodedFields7zhzg
			missingFieldsLeft7zhzg := maxFields7zhzg - totalEncodedFields7zhzg

			var nextMiss7zhzg int32 = -1
			var found7zhzg [maxFields7zhzg]bool
			var curField7zhzg string

		doneWithStruct7zhzg:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7zhzg > 0 || missingFieldsLeft7zhzg > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zhzg, missingFieldsLeft7zhzg, msgp.ShowFound(found7zhzg[:]), decodeMsgFieldOrder7zhzg)
				if encodedFieldsLeft7zhzg > 0 {
					encodedFieldsLeft7zhzg--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7zhzg = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7zhzg < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7zhzg = 0
					}
					for nextMiss7zhzg < maxFields7zhzg && found7zhzg[nextMiss7zhzg] {
						nextMiss7zhzg++
					}
					if nextMiss7zhzg == maxFields7zhzg {
						// filled all the empty fields!
						break doneWithStruct7zhzg
					}
					missingFieldsLeft7zhzg--
					curField7zhzg = decodeMsgFieldOrder7zhzg[nextMiss7zhzg]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7zhzg)
				switch curField7zhzg {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7zhzg[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7zhzg[1] = true
					var zcvp uint32
					zcvp, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zcvp) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zcvp]
					} else {
						z.Inputs.Fld = make([]FieldT, zcvp)
					}
					for zpgv := range z.Inputs.Fld {
						err = z.Inputs.Fld[zpgv].DecodeMsg(dc)
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
			if nextMiss7zhzg != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zogt[3] = true
			const maxFields8zdfo = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zdfo uint32
			totalEncodedFields8zdfo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zdfo := totalEncodedFields8zdfo
			missingFieldsLeft8zdfo := maxFields8zdfo - totalEncodedFields8zdfo

			var nextMiss8zdfo int32 = -1
			var found8zdfo [maxFields8zdfo]bool
			var curField8zdfo string

		doneWithStruct8zdfo:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zdfo > 0 || missingFieldsLeft8zdfo > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zdfo, missingFieldsLeft8zdfo, msgp.ShowFound(found8zdfo[:]), decodeMsgFieldOrder8zdfo)
				if encodedFieldsLeft8zdfo > 0 {
					encodedFieldsLeft8zdfo--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zdfo = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zdfo < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zdfo = 0
					}
					for nextMiss8zdfo < maxFields8zdfo && found8zdfo[nextMiss8zdfo] {
						nextMiss8zdfo++
					}
					if nextMiss8zdfo == maxFields8zdfo {
						// filled all the empty fields!
						break doneWithStruct8zdfo
					}
					missingFieldsLeft8zdfo--
					curField8zdfo = decodeMsgFieldOrder8zdfo[nextMiss8zdfo]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zdfo)
				switch curField8zdfo {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zdfo[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zdfo[1] = true
					var zbdc uint32
					zbdc, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zbdc) {
						z.Returns.Fld = (z.Returns.Fld)[:zbdc]
					} else {
						z.Returns.Fld = make([]FieldT, zbdc)
					}
					for zoop := range z.Returns.Fld {
						err = z.Returns.Fld[zoop].DecodeMsg(dc)
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
			if nextMiss8zdfo != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zogt[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zogt[5] = true
			z.Comment, err = dc.ReadString()
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
	if nextMiss6zogt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zogt = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7zhzg = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zdfo = []string{"Nam", "Fld"}

// fieldsNotEmpty supports omitempty tags
func (z *MethodT) fieldsNotEmpty(isempty []bool) uint32 {
	return 6
}

// EncodeMsg implements msgp.Encodable
func (z *MethodT) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "MethodId"
	err = en.Append(0x86, 0xa8, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.MethodId)
	if err != nil {
		panic(err)
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		panic(err)
	}
	// write "Inputs"
	// map header, size 2
	// write "Nam"
	err = en.Append(0xa6, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Inputs.Nam)
	if err != nil {
		panic(err)
	}
	// write "Fld"
	err = en.Append(0xa3, 0x46, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Inputs.Fld)))
	if err != nil {
		panic(err)
	}
	for zpgv := range z.Inputs.Fld {
		err = z.Inputs.Fld[zpgv].EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	// write "Returns"
	// map header, size 2
	// write "Nam"
	err = en.Append(0xa7, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Returns.Nam)
	if err != nil {
		panic(err)
	}
	// write "Fld"
	err = en.Append(0xa3, 0x46, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Returns.Fld)))
	if err != nil {
		panic(err)
	}
	for zoop := range z.Returns.Fld {
		err = z.Returns.Fld[zoop].EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	// write "Deprecated"
	err = en.Append(0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Deprecated)
	if err != nil {
		panic(err)
	}
	// write "Comment"
	err = en.Append(0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Comment)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MethodT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "MethodId"
	o = append(o, 0x86, 0xa8, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.MethodId)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Inputs"
	// map header, size 2
	// string "Nam"
	o = append(o, 0xa6, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
	o = msgp.AppendString(o, z.Inputs.Nam)
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Inputs.Fld)))
	for zpgv := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zpgv].MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Returns"
	// map header, size 2
	// string "Nam"
	o = append(o, 0xa7, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
	o = msgp.AppendString(o, z.Returns.Nam)
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Returns.Fld)))
	for zoop := range z.Returns.Fld {
		o, err = z.Returns.Fld[zoop].MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Deprecated"
	o = append(o, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Deprecated)
	// string "Comment"
	o = append(o, 0xa7, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Comment)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MethodT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *MethodT) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields9zwup = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zwup uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zwup, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zwup := totalEncodedFields9zwup
	missingFieldsLeft9zwup := maxFields9zwup - totalEncodedFields9zwup

	var nextMiss9zwup int32 = -1
	var found9zwup [maxFields9zwup]bool
	var curField9zwup string

doneWithStruct9zwup:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zwup > 0 || missingFieldsLeft9zwup > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zwup, missingFieldsLeft9zwup, msgp.ShowFound(found9zwup[:]), unmarshalMsgFieldOrder9zwup)
		if encodedFieldsLeft9zwup > 0 {
			encodedFieldsLeft9zwup--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zwup = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zwup < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zwup = 0
			}
			for nextMiss9zwup < maxFields9zwup && found9zwup[nextMiss9zwup] {
				nextMiss9zwup++
			}
			if nextMiss9zwup == maxFields9zwup {
				// filled all the empty fields!
				break doneWithStruct9zwup
			}
			missingFieldsLeft9zwup--
			curField9zwup = unmarshalMsgFieldOrder9zwup[nextMiss9zwup]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zwup)
		switch curField9zwup {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zwup[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zwup[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zwup[2] = true
			const maxFields10zexe = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zexe uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zexe, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10zexe := totalEncodedFields10zexe
			missingFieldsLeft10zexe := maxFields10zexe - totalEncodedFields10zexe

			var nextMiss10zexe int32 = -1
			var found10zexe [maxFields10zexe]bool
			var curField10zexe string

		doneWithStruct10zexe:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10zexe > 0 || missingFieldsLeft10zexe > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zexe, missingFieldsLeft10zexe, msgp.ShowFound(found10zexe[:]), unmarshalMsgFieldOrder10zexe)
				if encodedFieldsLeft10zexe > 0 {
					encodedFieldsLeft10zexe--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10zexe = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zexe < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zexe = 0
					}
					for nextMiss10zexe < maxFields10zexe && found10zexe[nextMiss10zexe] {
						nextMiss10zexe++
					}
					if nextMiss10zexe == maxFields10zexe {
						// filled all the empty fields!
						break doneWithStruct10zexe
					}
					missingFieldsLeft10zexe--
					curField10zexe = unmarshalMsgFieldOrder10zexe[nextMiss10zexe]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zexe)
				switch curField10zexe {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10zexe[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10zexe[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zgok uint32
						zgok, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zgok) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zgok]
						} else {
							z.Inputs.Fld = make([]FieldT, zgok)
						}
						for zpgv := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zpgv].UnmarshalMsg(bts)
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
			if nextMiss10zexe != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zwup[3] = true
			const maxFields11zuxv = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zuxv uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zuxv, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zuxv := totalEncodedFields11zuxv
			missingFieldsLeft11zuxv := maxFields11zuxv - totalEncodedFields11zuxv

			var nextMiss11zuxv int32 = -1
			var found11zuxv [maxFields11zuxv]bool
			var curField11zuxv string

		doneWithStruct11zuxv:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zuxv > 0 || missingFieldsLeft11zuxv > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zuxv, missingFieldsLeft11zuxv, msgp.ShowFound(found11zuxv[:]), unmarshalMsgFieldOrder11zuxv)
				if encodedFieldsLeft11zuxv > 0 {
					encodedFieldsLeft11zuxv--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zuxv = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zuxv < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zuxv = 0
					}
					for nextMiss11zuxv < maxFields11zuxv && found11zuxv[nextMiss11zuxv] {
						nextMiss11zuxv++
					}
					if nextMiss11zuxv == maxFields11zuxv {
						// filled all the empty fields!
						break doneWithStruct11zuxv
					}
					missingFieldsLeft11zuxv--
					curField11zuxv = unmarshalMsgFieldOrder11zuxv[nextMiss11zuxv]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zuxv)
				switch curField11zuxv {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zuxv[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zuxv[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zpcc uint32
						zpcc, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zpcc) {
							z.Returns.Fld = (z.Returns.Fld)[:zpcc]
						} else {
							z.Returns.Fld = make([]FieldT, zpcc)
						}
						for zoop := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zoop].UnmarshalMsg(bts)
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
			if nextMiss11zuxv != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zwup[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zwup[5] = true
			z.Comment, bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss9zwup != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zwup = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10zexe = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zuxv = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zpgv := range z.Inputs.Fld {
		s += z.Inputs.Fld[zpgv].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for zoop := range z.Returns.Fld {
		s += z.Returns.Fld[zoop].Msgsize()
	}
	s += 11 + msgp.BoolSize + 8 + msgp.StringPrefixSize + len(z.Comment)
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *SchemaT) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields12zwkx = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zwkx uint32
	totalEncodedFields12zwkx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zwkx := totalEncodedFields12zwkx
	missingFieldsLeft12zwkx := maxFields12zwkx - totalEncodedFields12zwkx

	var nextMiss12zwkx int32 = -1
	var found12zwkx [maxFields12zwkx]bool
	var curField12zwkx string

doneWithStruct12zwkx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zwkx > 0 || missingFieldsLeft12zwkx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zwkx, missingFieldsLeft12zwkx, msgp.ShowFound(found12zwkx[:]), decodeMsgFieldOrder12zwkx)
		if encodedFieldsLeft12zwkx > 0 {
			encodedFieldsLeft12zwkx--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zwkx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zwkx < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zwkx = 0
			}
			for nextMiss12zwkx < maxFields12zwkx && found12zwkx[nextMiss12zwkx] {
				nextMiss12zwkx++
			}
			if nextMiss12zwkx == maxFields12zwkx {
				// filled all the empty fields!
				break doneWithStruct12zwkx
			}
			missingFieldsLeft12zwkx--
			curField12zwkx = decodeMsgFieldOrder12zwkx[nextMiss12zwkx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zwkx)
		switch curField12zwkx {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zwkx[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found12zwkx[1] = true
			var zksa uint32
			zksa, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zksa > 0 {
				z.Structs = make(map[int64]StructT, zksa)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zksa > 0 {
				zksa--
				var zwvv int64
				var zkkp StructT
				zwvv, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				const maxFields13zdlm = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields13zdlm uint32
				totalEncodedFields13zdlm, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft13zdlm := totalEncodedFields13zdlm
				missingFieldsLeft13zdlm := maxFields13zdlm - totalEncodedFields13zdlm

				var nextMiss13zdlm int32 = -1
				var found13zdlm [maxFields13zdlm]bool
				var curField13zdlm string

			doneWithStruct13zdlm:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft13zdlm > 0 || missingFieldsLeft13zdlm > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zdlm, missingFieldsLeft13zdlm, msgp.ShowFound(found13zdlm[:]), decodeMsgFieldOrder13zdlm)
					if encodedFieldsLeft13zdlm > 0 {
						encodedFieldsLeft13zdlm--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField13zdlm = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss13zdlm < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss13zdlm = 0
						}
						for nextMiss13zdlm < maxFields13zdlm && found13zdlm[nextMiss13zdlm] {
							nextMiss13zdlm++
						}
						if nextMiss13zdlm == maxFields13zdlm {
							// filled all the empty fields!
							break doneWithStruct13zdlm
						}
						missingFieldsLeft13zdlm--
						curField13zdlm = decodeMsgFieldOrder13zdlm[nextMiss13zdlm]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField13zdlm)
					switch curField13zdlm {
					// -- templateDecodeMsg ends here --

					case "Nam":
						found13zdlm[0] = true
						zkkp.Nam, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found13zdlm[1] = true
						var zdxl uint32
						zdxl, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(zkkp.Fld) >= int(zdxl) {
							zkkp.Fld = (zkkp.Fld)[:zdxl]
						} else {
							zkkp.Fld = make([]FieldT, zdxl)
						}
						for zbit := range zkkp.Fld {
							err = zkkp.Fld[zbit].DecodeMsg(dc)
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
				if nextMiss13zdlm != -1 {
					dc.PopAlwaysNil()
				}

				z.Structs[zwvv] = zkkp
			}
		case "Interfaces":
			found12zwkx[2] = true
			var zbdc uint32
			zbdc, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Interfaces == nil && zbdc > 0 {
				z.Interfaces = make(map[int64]InterfaceT, zbdc)
			} else if len(z.Interfaces) > 0 {
				for key, _ := range z.Interfaces {
					delete(z.Interfaces, key)
				}
			}
			for zbdc > 0 {
				zbdc--
				var zmey int64
				var zscq InterfaceT
				zmey, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zscq.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Interfaces[zmey] = zscq
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zwkx != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zwkx = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var decodeMsgFieldOrder13zdlm = []string{"Nam", "Fld"}

// fieldsNotEmpty supports omitempty tags
func (z *SchemaT) fieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// EncodeMsg implements msgp.Encodable
func (z *SchemaT) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "PkgPath"
	err = en.Append(0x83, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
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
	err = en.WriteMapHeader(uint32(len(z.Structs)))
	if err != nil {
		panic(err)
	}
	for zwvv, zkkp := range z.Structs {
		err = en.WriteInt64(zwvv)
		if err != nil {
			panic(err)
		}
		// map header, size 2
		// write "Nam"
		err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteString(zkkp.Nam)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(zkkp.Fld)))
		if err != nil {
			panic(err)
		}
		for zbit := range zkkp.Fld {
			err = zkkp.Fld[zbit].EncodeMsg(en)
			if err != nil {
				panic(err)
			}
		}
	}
	// write "Interfaces"
	err = en.Append(0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Interfaces)))
	if err != nil {
		panic(err)
	}
	for zmey, zscq := range z.Interfaces {
		err = en.WriteInt64(zmey)
		if err != nil {
			panic(err)
		}
		err = zscq.EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SchemaT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "PkgPath"
	o = append(o, 0x83, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.PkgPath)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Structs)))
	for zwvv, zkkp := range z.Structs {
		o = msgp.AppendString(o, zwvv)
		// map header, size 2
		// string "Nam"
		o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
		o = msgp.AppendString(o, zkkp.Nam)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(zkkp.Fld)))
		for zbit := range zkkp.Fld {
			o, err = zkkp.Fld[zbit].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Interfaces"
	o = append(o, 0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Interfaces)))
	for zmey, zscq := range z.Interfaces {
		o = msgp.AppendString(o, zmey)
		o, err = zscq.MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SchemaT) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *SchemaT) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields14zark = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields14zark uint32
	if !nbs.AlwaysNil {
		totalEncodedFields14zark, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft14zark := totalEncodedFields14zark
	missingFieldsLeft14zark := maxFields14zark - totalEncodedFields14zark

	var nextMiss14zark int32 = -1
	var found14zark [maxFields14zark]bool
	var curField14zark string

doneWithStruct14zark:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zark > 0 || missingFieldsLeft14zark > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zark, missingFieldsLeft14zark, msgp.ShowFound(found14zark[:]), unmarshalMsgFieldOrder14zark)
		if encodedFieldsLeft14zark > 0 {
			encodedFieldsLeft14zark--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField14zark = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zark < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss14zark = 0
			}
			for nextMiss14zark < maxFields14zark && found14zark[nextMiss14zark] {
				nextMiss14zark++
			}
			if nextMiss14zark == maxFields14zark {
				// filled all the empty fields!
				break doneWithStruct14zark
			}
			missingFieldsLeft14zark--
			curField14zark = unmarshalMsgFieldOrder14zark[nextMiss14zark]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zark)
		switch curField14zark {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found14zark[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found14zark[1] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zndg uint32
				zndg, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zndg > 0 {
					z.Structs = make(map[int64]StructT, zndg)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zndg > 0 {
					var zwvv string
					var zkkp StructT
					zndg--
					zwvv, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					const maxFields15zclc = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields15zclc uint32
					if !nbs.AlwaysNil {
						totalEncodedFields15zclc, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft15zclc := totalEncodedFields15zclc
					missingFieldsLeft15zclc := maxFields15zclc - totalEncodedFields15zclc

					var nextMiss15zclc int32 = -1
					var found15zclc [maxFields15zclc]bool
					var curField15zclc string

				doneWithStruct15zclc:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft15zclc > 0 || missingFieldsLeft15zclc > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zclc, missingFieldsLeft15zclc, msgp.ShowFound(found15zclc[:]), unmarshalMsgFieldOrder15zclc)
						if encodedFieldsLeft15zclc > 0 {
							encodedFieldsLeft15zclc--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField15zclc = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss15zclc < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss15zclc = 0
							}
							for nextMiss15zclc < maxFields15zclc && found15zclc[nextMiss15zclc] {
								nextMiss15zclc++
							}
							if nextMiss15zclc == maxFields15zclc {
								// filled all the empty fields!
								break doneWithStruct15zclc
							}
							missingFieldsLeft15zclc--
							curField15zclc = unmarshalMsgFieldOrder15zclc[nextMiss15zclc]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField15zclc)
						switch curField15zclc {
						// -- templateUnmarshalMsg ends here --

						case "Nam":
							found15zclc[0] = true
							zkkp.Nam, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found15zclc[1] = true
							if nbs.AlwaysNil {
								(zkkp.Fld) = (zkkp.Fld)[:0]
							} else {

								var zuyv uint32
								zuyv, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(zkkp.Fld) >= int(zuyv) {
									zkkp.Fld = (zkkp.Fld)[:zuyv]
								} else {
									zkkp.Fld = make([]FieldT, zuyv)
								}
								for zbit := range zkkp.Fld {
									bts, err = zkkp.Fld[zbit].UnmarshalMsg(bts)
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
					if nextMiss15zclc != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Structs[zwvv] = zkkp
				}
			}
		case "Interfaces":
			found14zark[2] = true
			if nbs.AlwaysNil {
				if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}

			} else {

				var zikq uint32
				zikq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Interfaces == nil && zikq > 0 {
					z.Interfaces = make(map[int64]InterfaceT, zikq)
				} else if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}
				for zikq > 0 {
					var zmey string
					var zscq InterfaceT
					zikq--
					zmey, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zscq.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Interfaces[zmey] = zscq
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zark != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder14zark = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var unmarshalMsgFieldOrder15zclc = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zwvv, zkkp := range z.Structs {
			_ = zkkp
			s += msgp.StringPrefixSize + len(zwvv) + 1 + 4 + msgp.StringPrefixSize + len(zkkp.Nam) + 4 + msgp.ArrayHeaderSize
			for zbit := range zkkp.Fld {
				s += zkkp.Fld[zbit].Msgsize()
			}
		}
	}
	s += 11 + msgp.MapHeaderSize
	if z.Interfaces != nil {
		for zmey, zscq := range z.Interfaces {
			_ = zscq
			s += msgp.StringPrefixSize + len(zmey) + zscq.Msgsize()
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
	const maxFields16zmod = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zmod uint32
	totalEncodedFields16zmod, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zmod := totalEncodedFields16zmod
	missingFieldsLeft16zmod := maxFields16zmod - totalEncodedFields16zmod

	var nextMiss16zmod int32 = -1
	var found16zmod [maxFields16zmod]bool
	var curField16zmod string

doneWithStruct16zmod:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zmod > 0 || missingFieldsLeft16zmod > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zmod, missingFieldsLeft16zmod, msgp.ShowFound(found16zmod[:]), decodeMsgFieldOrder16zmod)
		if encodedFieldsLeft16zmod > 0 {
			encodedFieldsLeft16zmod--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zmod = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zmod < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zmod = 0
			}
			for nextMiss16zmod < maxFields16zmod && found16zmod[nextMiss16zmod] {
				nextMiss16zmod++
			}
			if nextMiss16zmod == maxFields16zmod {
				// filled all the empty fields!
				break doneWithStruct16zmod
			}
			missingFieldsLeft16zmod--
			curField16zmod = decodeMsgFieldOrder16zmod[nextMiss16zmod]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zmod)
		switch curField16zmod {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found16zmod[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found16zmod[1] = true
			var zgfw uint32
			zgfw, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zgfw) {
				z.Fld = (z.Fld)[:zgfw]
			} else {
				z.Fld = make([]FieldT, zgfw)
			}
			for zmrd := range z.Fld {
				err = z.Fld[zmrd].DecodeMsg(dc)
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
	if nextMiss16zmod != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder16zmod = []string{"Nam", "Fld"}

// fieldsNotEmpty supports omitempty tags
func (z *StructT) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *StructT) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Nam"
	err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nam)
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
	for zmrd := range z.Fld {
		err = z.Fld[zmrd].EncodeMsg(en)
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
	// string "Nam"
	o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
	o = msgp.AppendString(o, z.Nam)
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fld)))
	for zmrd := range z.Fld {
		o, err = z.Fld[zmrd].MarshalMsg(o)
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
	const maxFields17zgwu = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17zgwu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17zgwu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17zgwu := totalEncodedFields17zgwu
	missingFieldsLeft17zgwu := maxFields17zgwu - totalEncodedFields17zgwu

	var nextMiss17zgwu int32 = -1
	var found17zgwu [maxFields17zgwu]bool
	var curField17zgwu string

doneWithStruct17zgwu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17zgwu > 0 || missingFieldsLeft17zgwu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zgwu, missingFieldsLeft17zgwu, msgp.ShowFound(found17zgwu[:]), unmarshalMsgFieldOrder17zgwu)
		if encodedFieldsLeft17zgwu > 0 {
			encodedFieldsLeft17zgwu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17zgwu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17zgwu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17zgwu = 0
			}
			for nextMiss17zgwu < maxFields17zgwu && found17zgwu[nextMiss17zgwu] {
				nextMiss17zgwu++
			}
			if nextMiss17zgwu == maxFields17zgwu {
				// filled all the empty fields!
				break doneWithStruct17zgwu
			}
			missingFieldsLeft17zgwu--
			curField17zgwu = unmarshalMsgFieldOrder17zgwu[nextMiss17zgwu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17zgwu)
		switch curField17zgwu {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found17zgwu[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found17zgwu[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var ziry uint32
				ziry, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(ziry) {
					z.Fld = (z.Fld)[:ziry]
				} else {
					z.Fld = make([]FieldT, ziry)
				}
				for zmrd := range z.Fld {
					bts, err = z.Fld[zmrd].UnmarshalMsg(bts)
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
	if nextMiss17zgwu != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder17zgwu = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zmrd := range z.Fld {
		s += z.Fld[zmrd].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *StructTypeId) DecodeMsg(dc *msgp.Reader) (err error) {
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
		var zahs int64
		zahs, err = dc.ReadInt64()
		(*z) = StructTypeId(zahs)
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
func (z StructTypeId) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt64(int64(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z StructTypeId) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt64(o, int64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StructTypeId) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *StructTypeId) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zjxo int64
		zjxo, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zjxo)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z StructTypeId) Msgsize() (s int) {
	s = msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *ZPacket) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields18zjfi = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields18zjfi uint32
	totalEncodedFields18zjfi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft18zjfi := totalEncodedFields18zjfi
	missingFieldsLeft18zjfi := maxFields18zjfi - totalEncodedFields18zjfi

	var nextMiss18zjfi int32 = -1
	var found18zjfi [maxFields18zjfi]bool
	var curField18zjfi string

doneWithStruct18zjfi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft18zjfi > 0 || missingFieldsLeft18zjfi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zjfi, missingFieldsLeft18zjfi, msgp.ShowFound(found18zjfi[:]), decodeMsgFieldOrder18zjfi)
		if encodedFieldsLeft18zjfi > 0 {
			encodedFieldsLeft18zjfi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField18zjfi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss18zjfi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss18zjfi = 0
			}
			for nextMiss18zjfi < maxFields18zjfi && found18zjfi[nextMiss18zjfi] {
				nextMiss18zjfi++
			}
			if nextMiss18zjfi == maxFields18zjfi {
				// filled all the empty fields!
				break doneWithStruct18zjfi
			}
			missingFieldsLeft18zjfi--
			curField18zjfi = decodeMsgFieldOrder18zjfi[nextMiss18zjfi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField18zjfi)
		switch curField18zjfi {
		// -- templateDecodeMsg ends here --

		case "Ty":
			found18zjfi[0] = true
			{
				var zqou int32
				zqou, err = dc.ReadInt32()
				z.Ty = Ztype(zqou)
			}
			if err != nil {
				panic(err)
			}
		case "Id":
			found18zjfi[1] = true
			{
				var zkea int64
				zkea, err = dc.ReadInt64()
				z.Id = StructTypeId(zkea)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found18zjfi[2] = true
			var zgai uint32
			zgai, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zgai > 0 {
				z.Da = make(map[int64]msgp.Raw, zgai)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zgai > 0 {
				zgai--
				var zasg int64
				var zqko msgp.Raw
				zasg, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zqko.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zasg] = zqko
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss18zjfi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZPacket
var decodeMsgFieldOrder18zjfi = []string{"Ty", "Id", "Da"}

// fieldsNotEmpty supports omitempty tags
func (z *ZPacket) fieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// EncodeMsg implements msgp.Encodable
func (z *ZPacket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Ty"
	err = en.Append(0x83, 0xa2, 0x54, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteInt32(int32(z.Ty))
	if err != nil {
		panic(err)
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(int64(z.Id))
	if err != nil {
		panic(err)
	}
	// write "Da"
	err = en.Append(0xa2, 0x44, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Da)))
	if err != nil {
		panic(err)
	}
	for zasg, zqko := range z.Da {
		err = en.WriteInt64(zasg)
		if err != nil {
			panic(err)
		}
		err = zqko.EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ZPacket) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Ty"
	o = append(o, 0x83, 0xa2, 0x54, 0x79)
	o = msgp.AppendInt32(o, int32(z.Ty))
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt64(o, int64(z.Id))
	// string "Da"
	o = append(o, 0xa2, 0x44, 0x61)
	o = msgp.AppendMapHeader(o, uint32(len(z.Da)))
	for zasg, zqko := range z.Da {
		o = msgp.AppendString(o, zasg)
		o, err = zqko.MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ZPacket) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *ZPacket) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields19zena = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19zena uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19zena, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19zena := totalEncodedFields19zena
	missingFieldsLeft19zena := maxFields19zena - totalEncodedFields19zena

	var nextMiss19zena int32 = -1
	var found19zena [maxFields19zena]bool
	var curField19zena string

doneWithStruct19zena:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zena > 0 || missingFieldsLeft19zena > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zena, missingFieldsLeft19zena, msgp.ShowFound(found19zena[:]), unmarshalMsgFieldOrder19zena)
		if encodedFieldsLeft19zena > 0 {
			encodedFieldsLeft19zena--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19zena = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zena < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19zena = 0
			}
			for nextMiss19zena < maxFields19zena && found19zena[nextMiss19zena] {
				nextMiss19zena++
			}
			if nextMiss19zena == maxFields19zena {
				// filled all the empty fields!
				break doneWithStruct19zena
			}
			missingFieldsLeft19zena--
			curField19zena = unmarshalMsgFieldOrder19zena[nextMiss19zena]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zena)
		switch curField19zena {
		// -- templateUnmarshalMsg ends here --

		case "Ty":
			found19zena[0] = true
			{
				var zopv int32
				zopv, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ty = Ztype(zopv)
			}
		case "Id":
			found19zena[1] = true
			{
				var zunm int64
				zunm, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zunm)
			}
		case "Da":
			found19zena[2] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zbae uint32
				zbae, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zbae > 0 {
					z.Da = make(map[int64]msgp.Raw, zbae)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zbae > 0 {
					var zasg string
					var zqko msgp.Raw
					zbae--
					zasg, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zqko.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zasg] = zqko
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19zena != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZPacket
var unmarshalMsgFieldOrder19zena = []string{"Ty", "Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zasg, zqko := range z.Da {
			_ = zqko
			s += msgp.StringPrefixSize + len(zasg) + zqko.Msgsize()
		}
	}
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

	{
		var zvif int32
		zvif, err = dc.ReadInt32()
		(*z) = Ztype(zvif)
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
func (z Ztype) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Ztype) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
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

	{
		var zsck int32
		zsck, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Ztype(zsck)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Ztype) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}
