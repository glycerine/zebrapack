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
	const maxFields0zipn = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zipn uint32
	totalEncodedFields0zipn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zipn := totalEncodedFields0zipn
	missingFieldsLeft0zipn := maxFields0zipn - totalEncodedFields0zipn

	var nextMiss0zipn int32 = -1
	var found0zipn [maxFields0zipn]bool
	var curField0zipn string

doneWithStruct0zipn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zipn > 0 || missingFieldsLeft0zipn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zipn, missingFieldsLeft0zipn, msgp.ShowFound(found0zipn[:]), decodeMsgFieldOrder0zipn)
		if encodedFieldsLeft0zipn > 0 {
			encodedFieldsLeft0zipn--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zipn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zipn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zipn = 0
			}
			for nextMiss0zipn < maxFields0zipn && found0zipn[nextMiss0zipn] {
				nextMiss0zipn++
			}
			if nextMiss0zipn == maxFields0zipn {
				// filled all the empty fields!
				break doneWithStruct0zipn
			}
			missingFieldsLeft0zipn--
			curField0zipn = decodeMsgFieldOrder0zipn[nextMiss0zipn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zipn)
		switch curField0zipn {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zipn[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zipn[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found0zipn[2] = true
			{
				var zgue int32
				zgue, err = dc.ReadInt32()
				z.Ztyp = Ztype(zgue)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0zipn[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zipn[4] = true
			var zwsb uint32
			zwsb, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zwsb > 0 {
				z.Tag = make(map[string]string, zwsb)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zwsb > 0 {
				zwsb--
				var zibp string
				var zpyn string
				zibp, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zpyn, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zibp] = zpyn
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zipn != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zipn = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

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
	var empty_zrnv [5]bool
	fieldsInUse_zdbt := z.fieldsNotEmpty(empty_zrnv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zdbt)
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
	if !empty_zrnv[3] {
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

	if !empty_zrnv[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zibp, zpyn := range z.Tag {
			err = en.WriteString(zibp)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zpyn)
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
		for zibp, zpyn := range z.Tag {
			o = msgp.AppendString(o, zibp)
			o = msgp.AppendString(o, zpyn)
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
	const maxFields1zxga = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zxga uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zxga, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zxga := totalEncodedFields1zxga
	missingFieldsLeft1zxga := maxFields1zxga - totalEncodedFields1zxga

	var nextMiss1zxga int32 = -1
	var found1zxga [maxFields1zxga]bool
	var curField1zxga string

doneWithStruct1zxga:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zxga > 0 || missingFieldsLeft1zxga > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zxga, missingFieldsLeft1zxga, msgp.ShowFound(found1zxga[:]), unmarshalMsgFieldOrder1zxga)
		if encodedFieldsLeft1zxga > 0 {
			encodedFieldsLeft1zxga--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zxga = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zxga < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zxga = 0
			}
			for nextMiss1zxga < maxFields1zxga && found1zxga[nextMiss1zxga] {
				nextMiss1zxga++
			}
			if nextMiss1zxga == maxFields1zxga {
				// filled all the empty fields!
				break doneWithStruct1zxga
			}
			missingFieldsLeft1zxga--
			curField1zxga = unmarshalMsgFieldOrder1zxga[nextMiss1zxga]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zxga)
		switch curField1zxga {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zxga[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zxga[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found1zxga[2] = true
			{
				var zxws int32
				zxws, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ztyp = Ztype(zxws)
			}
		case "Varg":
			found1zxga[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zxga[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zdra uint32
				zdra, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zdra > 0 {
					z.Tag = make(map[string]string, zdra)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zdra > 0 {
					var zibp string
					var zpyn string
					zdra--
					zibp, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zpyn, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zibp] = zpyn
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zxga != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zxga = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zibp, zpyn := range z.Tag {
			_ = zpyn
			_ = zibp
			s += msgp.StringPrefixSize + len(zibp) + msgp.StringPrefixSize + len(zpyn)
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
	const maxFields2zuve = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zuve uint32
	totalEncodedFields2zuve, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zuve := totalEncodedFields2zuve
	missingFieldsLeft2zuve := maxFields2zuve - totalEncodedFields2zuve

	var nextMiss2zuve int32 = -1
	var found2zuve [maxFields2zuve]bool
	var curField2zuve string

doneWithStruct2zuve:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zuve > 0 || missingFieldsLeft2zuve > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zuve, missingFieldsLeft2zuve, msgp.ShowFound(found2zuve[:]), decodeMsgFieldOrder2zuve)
		if encodedFieldsLeft2zuve > 0 {
			encodedFieldsLeft2zuve--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zuve = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zuve < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zuve = 0
			}
			for nextMiss2zuve < maxFields2zuve && found2zuve[nextMiss2zuve] {
				nextMiss2zuve++
			}
			if nextMiss2zuve == maxFields2zuve {
				// filled all the empty fields!
				break doneWithStruct2zuve
			}
			missingFieldsLeft2zuve--
			curField2zuve = decodeMsgFieldOrder2zuve[nextMiss2zuve]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zuve)
		switch curField2zuve {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zuve[0] = true
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
	if nextMiss2zuve != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zuve = []string{"IfaceId"}

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
	const maxFields3zwhn = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zwhn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zwhn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zwhn := totalEncodedFields3zwhn
	missingFieldsLeft3zwhn := maxFields3zwhn - totalEncodedFields3zwhn

	var nextMiss3zwhn int32 = -1
	var found3zwhn [maxFields3zwhn]bool
	var curField3zwhn string

doneWithStruct3zwhn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zwhn > 0 || missingFieldsLeft3zwhn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zwhn, missingFieldsLeft3zwhn, msgp.ShowFound(found3zwhn[:]), unmarshalMsgFieldOrder3zwhn)
		if encodedFieldsLeft3zwhn > 0 {
			encodedFieldsLeft3zwhn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zwhn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zwhn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zwhn = 0
			}
			for nextMiss3zwhn < maxFields3zwhn && found3zwhn[nextMiss3zwhn] {
				nextMiss3zwhn++
			}
			if nextMiss3zwhn == maxFields3zwhn {
				// filled all the empty fields!
				break doneWithStruct3zwhn
			}
			missingFieldsLeft3zwhn--
			curField3zwhn = unmarshalMsgFieldOrder3zwhn[nextMiss3zwhn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zwhn)
		switch curField3zwhn {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3zwhn[0] = true
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
	if nextMiss3zwhn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3zwhn = []string{"IfaceId"}

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
	const maxFields4zcna = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zcna uint32
	totalEncodedFields4zcna, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zcna := totalEncodedFields4zcna
	missingFieldsLeft4zcna := maxFields4zcna - totalEncodedFields4zcna

	var nextMiss4zcna int32 = -1
	var found4zcna [maxFields4zcna]bool
	var curField4zcna string

doneWithStruct4zcna:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zcna > 0 || missingFieldsLeft4zcna > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zcna, missingFieldsLeft4zcna, msgp.ShowFound(found4zcna[:]), decodeMsgFieldOrder4zcna)
		if encodedFieldsLeft4zcna > 0 {
			encodedFieldsLeft4zcna--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zcna = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zcna < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zcna = 0
			}
			for nextMiss4zcna < maxFields4zcna && found4zcna[nextMiss4zcna] {
				nextMiss4zcna++
			}
			if nextMiss4zcna == maxFields4zcna {
				// filled all the empty fields!
				break doneWithStruct4zcna
			}
			missingFieldsLeft4zcna--
			curField4zcna = decodeMsgFieldOrder4zcna[nextMiss4zcna]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zcna)
		switch curField4zcna {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4zcna[0] = true
			var zkaa uint32
			zkaa, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zkaa) {
				z.Methods = (z.Methods)[:zkaa]
			} else {
				z.Methods = make([]MethodT, zkaa)
			}
			for zjeh := range z.Methods {
				err = z.Methods[zjeh].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4zcna[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zcna[2] = true
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
	if nextMiss4zcna != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4zcna = []string{"Methods", "Deprecated", "Comment"}

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
	for zjeh := range z.Methods {
		err = z.Methods[zjeh].EncodeMsg(en)
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
	for zjeh := range z.Methods {
		o, err = z.Methods[zjeh].MarshalMsg(o)
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
	const maxFields5zxoi = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zxoi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zxoi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zxoi := totalEncodedFields5zxoi
	missingFieldsLeft5zxoi := maxFields5zxoi - totalEncodedFields5zxoi

	var nextMiss5zxoi int32 = -1
	var found5zxoi [maxFields5zxoi]bool
	var curField5zxoi string

doneWithStruct5zxoi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zxoi > 0 || missingFieldsLeft5zxoi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxoi, missingFieldsLeft5zxoi, msgp.ShowFound(found5zxoi[:]), unmarshalMsgFieldOrder5zxoi)
		if encodedFieldsLeft5zxoi > 0 {
			encodedFieldsLeft5zxoi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zxoi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zxoi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zxoi = 0
			}
			for nextMiss5zxoi < maxFields5zxoi && found5zxoi[nextMiss5zxoi] {
				nextMiss5zxoi++
			}
			if nextMiss5zxoi == maxFields5zxoi {
				// filled all the empty fields!
				break doneWithStruct5zxoi
			}
			missingFieldsLeft5zxoi--
			curField5zxoi = unmarshalMsgFieldOrder5zxoi[nextMiss5zxoi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zxoi)
		switch curField5zxoi {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zxoi[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zzqo uint32
				zzqo, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zzqo) {
					z.Methods = (z.Methods)[:zzqo]
				} else {
					z.Methods = make([]MethodT, zzqo)
				}
				for zjeh := range z.Methods {
					bts, err = z.Methods[zjeh].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zxoi[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zxoi[2] = true
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
	if nextMiss5zxoi != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zxoi = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zjeh := range z.Methods {
		s += z.Methods[zjeh].Msgsize()
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
	const maxFields6zppu = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zppu uint32
	totalEncodedFields6zppu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zppu := totalEncodedFields6zppu
	missingFieldsLeft6zppu := maxFields6zppu - totalEncodedFields6zppu

	var nextMiss6zppu int32 = -1
	var found6zppu [maxFields6zppu]bool
	var curField6zppu string

doneWithStruct6zppu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zppu > 0 || missingFieldsLeft6zppu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zppu, missingFieldsLeft6zppu, msgp.ShowFound(found6zppu[:]), decodeMsgFieldOrder6zppu)
		if encodedFieldsLeft6zppu > 0 {
			encodedFieldsLeft6zppu--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zppu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zppu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zppu = 0
			}
			for nextMiss6zppu < maxFields6zppu && found6zppu[nextMiss6zppu] {
				nextMiss6zppu++
			}
			if nextMiss6zppu == maxFields6zppu {
				// filled all the empty fields!
				break doneWithStruct6zppu
			}
			missingFieldsLeft6zppu--
			curField6zppu = decodeMsgFieldOrder6zppu[nextMiss6zppu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zppu)
		switch curField6zppu {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zppu[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zppu[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zppu[2] = true
			const maxFields7ztjz = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7ztjz uint32
			totalEncodedFields7ztjz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7ztjz := totalEncodedFields7ztjz
			missingFieldsLeft7ztjz := maxFields7ztjz - totalEncodedFields7ztjz

			var nextMiss7ztjz int32 = -1
			var found7ztjz [maxFields7ztjz]bool
			var curField7ztjz string

		doneWithStruct7ztjz:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7ztjz > 0 || missingFieldsLeft7ztjz > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7ztjz, missingFieldsLeft7ztjz, msgp.ShowFound(found7ztjz[:]), decodeMsgFieldOrder7ztjz)
				if encodedFieldsLeft7ztjz > 0 {
					encodedFieldsLeft7ztjz--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7ztjz = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7ztjz < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7ztjz = 0
					}
					for nextMiss7ztjz < maxFields7ztjz && found7ztjz[nextMiss7ztjz] {
						nextMiss7ztjz++
					}
					if nextMiss7ztjz == maxFields7ztjz {
						// filled all the empty fields!
						break doneWithStruct7ztjz
					}
					missingFieldsLeft7ztjz--
					curField7ztjz = decodeMsgFieldOrder7ztjz[nextMiss7ztjz]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7ztjz)
				switch curField7ztjz {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7ztjz[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7ztjz[1] = true
					var zgao uint32
					zgao, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zgao) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zgao]
					} else {
						z.Inputs.Fld = make([]FieldT, zgao)
					}
					for zgce := range z.Inputs.Fld {
						err = z.Inputs.Fld[zgce].DecodeMsg(dc)
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
			if nextMiss7ztjz != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zppu[3] = true
			const maxFields8zrvb = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zrvb uint32
			totalEncodedFields8zrvb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zrvb := totalEncodedFields8zrvb
			missingFieldsLeft8zrvb := maxFields8zrvb - totalEncodedFields8zrvb

			var nextMiss8zrvb int32 = -1
			var found8zrvb [maxFields8zrvb]bool
			var curField8zrvb string

		doneWithStruct8zrvb:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zrvb > 0 || missingFieldsLeft8zrvb > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zrvb, missingFieldsLeft8zrvb, msgp.ShowFound(found8zrvb[:]), decodeMsgFieldOrder8zrvb)
				if encodedFieldsLeft8zrvb > 0 {
					encodedFieldsLeft8zrvb--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zrvb = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zrvb < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zrvb = 0
					}
					for nextMiss8zrvb < maxFields8zrvb && found8zrvb[nextMiss8zrvb] {
						nextMiss8zrvb++
					}
					if nextMiss8zrvb == maxFields8zrvb {
						// filled all the empty fields!
						break doneWithStruct8zrvb
					}
					missingFieldsLeft8zrvb--
					curField8zrvb = decodeMsgFieldOrder8zrvb[nextMiss8zrvb]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zrvb)
				switch curField8zrvb {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zrvb[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zrvb[1] = true
					var zczr uint32
					zczr, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zczr) {
						z.Returns.Fld = (z.Returns.Fld)[:zczr]
					} else {
						z.Returns.Fld = make([]FieldT, zczr)
					}
					for zmvc := range z.Returns.Fld {
						err = z.Returns.Fld[zmvc].DecodeMsg(dc)
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
			if nextMiss8zrvb != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zppu[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zppu[5] = true
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
	if nextMiss6zppu != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zppu = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7ztjz = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zrvb = []string{"Nam", "Fld"}

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
	for zgce := range z.Inputs.Fld {
		err = z.Inputs.Fld[zgce].EncodeMsg(en)
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
	for zmvc := range z.Returns.Fld {
		err = z.Returns.Fld[zmvc].EncodeMsg(en)
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
	for zgce := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zgce].MarshalMsg(o)
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
	for zmvc := range z.Returns.Fld {
		o, err = z.Returns.Fld[zmvc].MarshalMsg(o)
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
	const maxFields9zfwk = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zfwk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zfwk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zfwk := totalEncodedFields9zfwk
	missingFieldsLeft9zfwk := maxFields9zfwk - totalEncodedFields9zfwk

	var nextMiss9zfwk int32 = -1
	var found9zfwk [maxFields9zfwk]bool
	var curField9zfwk string

doneWithStruct9zfwk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zfwk > 0 || missingFieldsLeft9zfwk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zfwk, missingFieldsLeft9zfwk, msgp.ShowFound(found9zfwk[:]), unmarshalMsgFieldOrder9zfwk)
		if encodedFieldsLeft9zfwk > 0 {
			encodedFieldsLeft9zfwk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zfwk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zfwk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zfwk = 0
			}
			for nextMiss9zfwk < maxFields9zfwk && found9zfwk[nextMiss9zfwk] {
				nextMiss9zfwk++
			}
			if nextMiss9zfwk == maxFields9zfwk {
				// filled all the empty fields!
				break doneWithStruct9zfwk
			}
			missingFieldsLeft9zfwk--
			curField9zfwk = unmarshalMsgFieldOrder9zfwk[nextMiss9zfwk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zfwk)
		switch curField9zfwk {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zfwk[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zfwk[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zfwk[2] = true
			const maxFields10zeke = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zeke uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zeke, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10zeke := totalEncodedFields10zeke
			missingFieldsLeft10zeke := maxFields10zeke - totalEncodedFields10zeke

			var nextMiss10zeke int32 = -1
			var found10zeke [maxFields10zeke]bool
			var curField10zeke string

		doneWithStruct10zeke:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10zeke > 0 || missingFieldsLeft10zeke > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zeke, missingFieldsLeft10zeke, msgp.ShowFound(found10zeke[:]), unmarshalMsgFieldOrder10zeke)
				if encodedFieldsLeft10zeke > 0 {
					encodedFieldsLeft10zeke--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10zeke = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zeke < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zeke = 0
					}
					for nextMiss10zeke < maxFields10zeke && found10zeke[nextMiss10zeke] {
						nextMiss10zeke++
					}
					if nextMiss10zeke == maxFields10zeke {
						// filled all the empty fields!
						break doneWithStruct10zeke
					}
					missingFieldsLeft10zeke--
					curField10zeke = unmarshalMsgFieldOrder10zeke[nextMiss10zeke]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zeke)
				switch curField10zeke {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10zeke[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10zeke[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zcfk uint32
						zcfk, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zcfk) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zcfk]
						} else {
							z.Inputs.Fld = make([]FieldT, zcfk)
						}
						for zgce := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zgce].UnmarshalMsg(bts)
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
			if nextMiss10zeke != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zfwk[3] = true
			const maxFields11zbxr = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zbxr uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zbxr, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zbxr := totalEncodedFields11zbxr
			missingFieldsLeft11zbxr := maxFields11zbxr - totalEncodedFields11zbxr

			var nextMiss11zbxr int32 = -1
			var found11zbxr [maxFields11zbxr]bool
			var curField11zbxr string

		doneWithStruct11zbxr:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zbxr > 0 || missingFieldsLeft11zbxr > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zbxr, missingFieldsLeft11zbxr, msgp.ShowFound(found11zbxr[:]), unmarshalMsgFieldOrder11zbxr)
				if encodedFieldsLeft11zbxr > 0 {
					encodedFieldsLeft11zbxr--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zbxr = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zbxr < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zbxr = 0
					}
					for nextMiss11zbxr < maxFields11zbxr && found11zbxr[nextMiss11zbxr] {
						nextMiss11zbxr++
					}
					if nextMiss11zbxr == maxFields11zbxr {
						// filled all the empty fields!
						break doneWithStruct11zbxr
					}
					missingFieldsLeft11zbxr--
					curField11zbxr = unmarshalMsgFieldOrder11zbxr[nextMiss11zbxr]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zbxr)
				switch curField11zbxr {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zbxr[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zbxr[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var ztow uint32
						ztow, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(ztow) {
							z.Returns.Fld = (z.Returns.Fld)[:ztow]
						} else {
							z.Returns.Fld = make([]FieldT, ztow)
						}
						for zmvc := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zmvc].UnmarshalMsg(bts)
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
			if nextMiss11zbxr != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zfwk[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zfwk[5] = true
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
	if nextMiss9zfwk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zfwk = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10zeke = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zbxr = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zgce := range z.Inputs.Fld {
		s += z.Inputs.Fld[zgce].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for zmvc := range z.Returns.Fld {
		s += z.Returns.Fld[zmvc].Msgsize()
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
	const maxFields12zija = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zija uint32
	totalEncodedFields12zija, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zija := totalEncodedFields12zija
	missingFieldsLeft12zija := maxFields12zija - totalEncodedFields12zija

	var nextMiss12zija int32 = -1
	var found12zija [maxFields12zija]bool
	var curField12zija string

doneWithStruct12zija:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zija > 0 || missingFieldsLeft12zija > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zija, missingFieldsLeft12zija, msgp.ShowFound(found12zija[:]), decodeMsgFieldOrder12zija)
		if encodedFieldsLeft12zija > 0 {
			encodedFieldsLeft12zija--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zija = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zija < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zija = 0
			}
			for nextMiss12zija < maxFields12zija && found12zija[nextMiss12zija] {
				nextMiss12zija++
			}
			if nextMiss12zija == maxFields12zija {
				// filled all the empty fields!
				break doneWithStruct12zija
			}
			missingFieldsLeft12zija--
			curField12zija = decodeMsgFieldOrder12zija[nextMiss12zija]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zija)
		switch curField12zija {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zija[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found12zija[1] = true
			var zkwc uint32
			zkwc, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zkwc > 0 {
				z.Structs = make(map[int64]StructT, zkwc)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zkwc > 0 {
				zkwc--
				var zkoi int64
				var znlq StructT
				zkoi, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				const maxFields13zuxv = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields13zuxv uint32
				totalEncodedFields13zuxv, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft13zuxv := totalEncodedFields13zuxv
				missingFieldsLeft13zuxv := maxFields13zuxv - totalEncodedFields13zuxv

				var nextMiss13zuxv int32 = -1
				var found13zuxv [maxFields13zuxv]bool
				var curField13zuxv string

			doneWithStruct13zuxv:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft13zuxv > 0 || missingFieldsLeft13zuxv > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zuxv, missingFieldsLeft13zuxv, msgp.ShowFound(found13zuxv[:]), decodeMsgFieldOrder13zuxv)
					if encodedFieldsLeft13zuxv > 0 {
						encodedFieldsLeft13zuxv--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField13zuxv = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss13zuxv < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss13zuxv = 0
						}
						for nextMiss13zuxv < maxFields13zuxv && found13zuxv[nextMiss13zuxv] {
							nextMiss13zuxv++
						}
						if nextMiss13zuxv == maxFields13zuxv {
							// filled all the empty fields!
							break doneWithStruct13zuxv
						}
						missingFieldsLeft13zuxv--
						curField13zuxv = decodeMsgFieldOrder13zuxv[nextMiss13zuxv]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField13zuxv)
					switch curField13zuxv {
					// -- templateDecodeMsg ends here --

					case "Nam":
						found13zuxv[0] = true
						znlq.Nam, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found13zuxv[1] = true
						var zxbl uint32
						zxbl, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(znlq.Fld) >= int(zxbl) {
							znlq.Fld = (znlq.Fld)[:zxbl]
						} else {
							znlq.Fld = make([]FieldT, zxbl)
						}
						for zvwg := range znlq.Fld {
							err = znlq.Fld[zvwg].DecodeMsg(dc)
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
				if nextMiss13zuxv != -1 {
					dc.PopAlwaysNil()
				}

				z.Structs[zkoi] = znlq
			}
		case "Interfaces":
			found12zija[2] = true
			var zvkd uint32
			zvkd, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Interfaces == nil && zvkd > 0 {
				z.Interfaces = make(map[int64]InterfaceT, zvkd)
			} else if len(z.Interfaces) > 0 {
				for key, _ := range z.Interfaces {
					delete(z.Interfaces, key)
				}
			}
			for zvkd > 0 {
				zvkd--
				var zyxj int64
				var zbiv InterfaceT
				zyxj, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zbiv.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Interfaces[zyxj] = zbiv
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zija != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zija = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var decodeMsgFieldOrder13zuxv = []string{"Nam", "Fld"}

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
	for zkoi, znlq := range z.Structs {
		err = en.WriteInt64(zkoi)
		if err != nil {
			panic(err)
		}
		// map header, size 2
		// write "Nam"
		err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteString(znlq.Nam)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(znlq.Fld)))
		if err != nil {
			panic(err)
		}
		for zvwg := range znlq.Fld {
			err = znlq.Fld[zvwg].EncodeMsg(en)
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
	for zyxj, zbiv := range z.Interfaces {
		err = en.WriteInt64(zyxj)
		if err != nil {
			panic(err)
		}
		err = zbiv.EncodeMsg(en)
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
	for zkoi, znlq := range z.Structs {
		o = msgp.AppendInt64(o, zkoi)
		// map header, size 2
		// string "Nam"
		o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
		o = msgp.AppendString(o, znlq.Nam)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(znlq.Fld)))
		for zvwg := range znlq.Fld {
			o, err = znlq.Fld[zvwg].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Interfaces"
	o = append(o, 0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Interfaces)))
	for zyxj, zbiv := range z.Interfaces {
		o = msgp.AppendInt64(o, zyxj)
		o, err = zbiv.MarshalMsg(o)
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
	const maxFields14zirf = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields14zirf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields14zirf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft14zirf := totalEncodedFields14zirf
	missingFieldsLeft14zirf := maxFields14zirf - totalEncodedFields14zirf

	var nextMiss14zirf int32 = -1
	var found14zirf [maxFields14zirf]bool
	var curField14zirf string

doneWithStruct14zirf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zirf > 0 || missingFieldsLeft14zirf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zirf, missingFieldsLeft14zirf, msgp.ShowFound(found14zirf[:]), unmarshalMsgFieldOrder14zirf)
		if encodedFieldsLeft14zirf > 0 {
			encodedFieldsLeft14zirf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField14zirf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zirf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss14zirf = 0
			}
			for nextMiss14zirf < maxFields14zirf && found14zirf[nextMiss14zirf] {
				nextMiss14zirf++
			}
			if nextMiss14zirf == maxFields14zirf {
				// filled all the empty fields!
				break doneWithStruct14zirf
			}
			missingFieldsLeft14zirf--
			curField14zirf = unmarshalMsgFieldOrder14zirf[nextMiss14zirf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zirf)
		switch curField14zirf {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found14zirf[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found14zirf[1] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zxdq uint32
				zxdq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zxdq > 0 {
					z.Structs = make(map[int64]StructT, zxdq)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zxdq > 0 {
					var zkoi int64
					var znlq StructT
					zxdq--
					zkoi, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					const maxFields15zgtj = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields15zgtj uint32
					if !nbs.AlwaysNil {
						totalEncodedFields15zgtj, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft15zgtj := totalEncodedFields15zgtj
					missingFieldsLeft15zgtj := maxFields15zgtj - totalEncodedFields15zgtj

					var nextMiss15zgtj int32 = -1
					var found15zgtj [maxFields15zgtj]bool
					var curField15zgtj string

				doneWithStruct15zgtj:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft15zgtj > 0 || missingFieldsLeft15zgtj > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zgtj, missingFieldsLeft15zgtj, msgp.ShowFound(found15zgtj[:]), unmarshalMsgFieldOrder15zgtj)
						if encodedFieldsLeft15zgtj > 0 {
							encodedFieldsLeft15zgtj--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField15zgtj = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss15zgtj < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss15zgtj = 0
							}
							for nextMiss15zgtj < maxFields15zgtj && found15zgtj[nextMiss15zgtj] {
								nextMiss15zgtj++
							}
							if nextMiss15zgtj == maxFields15zgtj {
								// filled all the empty fields!
								break doneWithStruct15zgtj
							}
							missingFieldsLeft15zgtj--
							curField15zgtj = unmarshalMsgFieldOrder15zgtj[nextMiss15zgtj]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField15zgtj)
						switch curField15zgtj {
						// -- templateUnmarshalMsg ends here --

						case "Nam":
							found15zgtj[0] = true
							znlq.Nam, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found15zgtj[1] = true
							if nbs.AlwaysNil {
								(znlq.Fld) = (znlq.Fld)[:0]
							} else {

								var zeey uint32
								zeey, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(znlq.Fld) >= int(zeey) {
									znlq.Fld = (znlq.Fld)[:zeey]
								} else {
									znlq.Fld = make([]FieldT, zeey)
								}
								for zvwg := range znlq.Fld {
									bts, err = znlq.Fld[zvwg].UnmarshalMsg(bts)
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
					if nextMiss15zgtj != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Structs[zkoi] = znlq
				}
			}
		case "Interfaces":
			found14zirf[2] = true
			if nbs.AlwaysNil {
				if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}

			} else {

				var zuzy uint32
				zuzy, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Interfaces == nil && zuzy > 0 {
					z.Interfaces = make(map[int64]InterfaceT, zuzy)
				} else if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}
				for zuzy > 0 {
					var zyxj int64
					var zbiv InterfaceT
					zuzy--
					zyxj, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zbiv.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Interfaces[zyxj] = zbiv
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zirf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder14zirf = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var unmarshalMsgFieldOrder15zgtj = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zkoi, znlq := range z.Structs {
			_ = znlq
			_ = zkoi
			s += msgp.Int64Size + 1 + 4 + msgp.StringPrefixSize + len(znlq.Nam) + 4 + msgp.ArrayHeaderSize
			for zvwg := range znlq.Fld {
				s += znlq.Fld[zvwg].Msgsize()
			}
		}
	}
	s += 11 + msgp.MapHeaderSize
	if z.Interfaces != nil {
		for zyxj, zbiv := range z.Interfaces {
			_ = zbiv
			_ = zyxj
			s += msgp.Int64Size + zbiv.Msgsize()
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
	const maxFields16znqj = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16znqj uint32
	totalEncodedFields16znqj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16znqj := totalEncodedFields16znqj
	missingFieldsLeft16znqj := maxFields16znqj - totalEncodedFields16znqj

	var nextMiss16znqj int32 = -1
	var found16znqj [maxFields16znqj]bool
	var curField16znqj string

doneWithStruct16znqj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16znqj > 0 || missingFieldsLeft16znqj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16znqj, missingFieldsLeft16znqj, msgp.ShowFound(found16znqj[:]), decodeMsgFieldOrder16znqj)
		if encodedFieldsLeft16znqj > 0 {
			encodedFieldsLeft16znqj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16znqj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16znqj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16znqj = 0
			}
			for nextMiss16znqj < maxFields16znqj && found16znqj[nextMiss16znqj] {
				nextMiss16znqj++
			}
			if nextMiss16znqj == maxFields16znqj {
				// filled all the empty fields!
				break doneWithStruct16znqj
			}
			missingFieldsLeft16znqj--
			curField16znqj = decodeMsgFieldOrder16znqj[nextMiss16znqj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16znqj)
		switch curField16znqj {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found16znqj[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found16znqj[1] = true
			var zdbr uint32
			zdbr, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zdbr) {
				z.Fld = (z.Fld)[:zdbr]
			} else {
				z.Fld = make([]FieldT, zdbr)
			}
			for zktb := range z.Fld {
				err = z.Fld[zktb].DecodeMsg(dc)
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
	if nextMiss16znqj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder16znqj = []string{"Nam", "Fld"}

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
	for zktb := range z.Fld {
		err = z.Fld[zktb].EncodeMsg(en)
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
	for zktb := range z.Fld {
		o, err = z.Fld[zktb].MarshalMsg(o)
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
	const maxFields17zujc = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17zujc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17zujc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17zujc := totalEncodedFields17zujc
	missingFieldsLeft17zujc := maxFields17zujc - totalEncodedFields17zujc

	var nextMiss17zujc int32 = -1
	var found17zujc [maxFields17zujc]bool
	var curField17zujc string

doneWithStruct17zujc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17zujc > 0 || missingFieldsLeft17zujc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zujc, missingFieldsLeft17zujc, msgp.ShowFound(found17zujc[:]), unmarshalMsgFieldOrder17zujc)
		if encodedFieldsLeft17zujc > 0 {
			encodedFieldsLeft17zujc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17zujc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17zujc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17zujc = 0
			}
			for nextMiss17zujc < maxFields17zujc && found17zujc[nextMiss17zujc] {
				nextMiss17zujc++
			}
			if nextMiss17zujc == maxFields17zujc {
				// filled all the empty fields!
				break doneWithStruct17zujc
			}
			missingFieldsLeft17zujc--
			curField17zujc = unmarshalMsgFieldOrder17zujc[nextMiss17zujc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17zujc)
		switch curField17zujc {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found17zujc[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found17zujc[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zcyn uint32
				zcyn, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zcyn) {
					z.Fld = (z.Fld)[:zcyn]
				} else {
					z.Fld = make([]FieldT, zcyn)
				}
				for zktb := range z.Fld {
					bts, err = z.Fld[zktb].UnmarshalMsg(bts)
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
	if nextMiss17zujc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder17zujc = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zktb := range z.Fld {
		s += z.Fld[zktb].Msgsize()
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
		var zrsw int64
		zrsw, err = dc.ReadInt64()
		(*z) = StructTypeId(zrsw)
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
		var zrbi int64
		zrbi, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zrbi)
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
	const maxFields18zzou = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields18zzou uint32
	totalEncodedFields18zzou, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft18zzou := totalEncodedFields18zzou
	missingFieldsLeft18zzou := maxFields18zzou - totalEncodedFields18zzou

	var nextMiss18zzou int32 = -1
	var found18zzou [maxFields18zzou]bool
	var curField18zzou string

doneWithStruct18zzou:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft18zzou > 0 || missingFieldsLeft18zzou > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zzou, missingFieldsLeft18zzou, msgp.ShowFound(found18zzou[:]), decodeMsgFieldOrder18zzou)
		if encodedFieldsLeft18zzou > 0 {
			encodedFieldsLeft18zzou--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField18zzou = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss18zzou < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss18zzou = 0
			}
			for nextMiss18zzou < maxFields18zzou && found18zzou[nextMiss18zzou] {
				nextMiss18zzou++
			}
			if nextMiss18zzou == maxFields18zzou {
				// filled all the empty fields!
				break doneWithStruct18zzou
			}
			missingFieldsLeft18zzou--
			curField18zzou = decodeMsgFieldOrder18zzou[nextMiss18zzou]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField18zzou)
		switch curField18zzou {
		// -- templateDecodeMsg ends here --

		case "Ty":
			found18zzou[0] = true
			{
				var zmha int32
				zmha, err = dc.ReadInt32()
				z.Ty = Ztype(zmha)
			}
			if err != nil {
				panic(err)
			}
		case "Id":
			found18zzou[1] = true
			{
				var zwpg int64
				zwpg, err = dc.ReadInt64()
				z.Id = StructTypeId(zwpg)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found18zzou[2] = true
			var zric uint32
			zric, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zric > 0 {
				z.Da = make(map[int64]msgp.Raw, zric)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zric > 0 {
				zric--
				var zxjj int64
				var zllx msgp.Raw
				zxjj, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zllx.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zxjj] = zllx
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss18zzou != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZPacket
var decodeMsgFieldOrder18zzou = []string{"Ty", "Id", "Da"}

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
	for zxjj, zllx := range z.Da {
		err = en.WriteInt64(zxjj)
		if err != nil {
			panic(err)
		}
		err = zllx.EncodeMsg(en)
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
	for zxjj, zllx := range z.Da {
		o = msgp.AppendInt64(o, zxjj)
		o, err = zllx.MarshalMsg(o)
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
	const maxFields19zafz = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19zafz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19zafz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19zafz := totalEncodedFields19zafz
	missingFieldsLeft19zafz := maxFields19zafz - totalEncodedFields19zafz

	var nextMiss19zafz int32 = -1
	var found19zafz [maxFields19zafz]bool
	var curField19zafz string

doneWithStruct19zafz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zafz > 0 || missingFieldsLeft19zafz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zafz, missingFieldsLeft19zafz, msgp.ShowFound(found19zafz[:]), unmarshalMsgFieldOrder19zafz)
		if encodedFieldsLeft19zafz > 0 {
			encodedFieldsLeft19zafz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19zafz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zafz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19zafz = 0
			}
			for nextMiss19zafz < maxFields19zafz && found19zafz[nextMiss19zafz] {
				nextMiss19zafz++
			}
			if nextMiss19zafz == maxFields19zafz {
				// filled all the empty fields!
				break doneWithStruct19zafz
			}
			missingFieldsLeft19zafz--
			curField19zafz = unmarshalMsgFieldOrder19zafz[nextMiss19zafz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zafz)
		switch curField19zafz {
		// -- templateUnmarshalMsg ends here --

		case "Ty":
			found19zafz[0] = true
			{
				var zjvh int32
				zjvh, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ty = Ztype(zjvh)
			}
		case "Id":
			found19zafz[1] = true
			{
				var zpez int64
				zpez, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zpez)
			}
		case "Da":
			found19zafz[2] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zlof uint32
				zlof, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zlof > 0 {
					z.Da = make(map[int64]msgp.Raw, zlof)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zlof > 0 {
					var zxjj int64
					var zllx msgp.Raw
					zlof--
					zxjj, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zllx.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zxjj] = zllx
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19zafz != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZPacket
var unmarshalMsgFieldOrder19zafz = []string{"Ty", "Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zxjj, zllx := range z.Da {
			_ = zllx
			_ = zxjj
			s += msgp.Int64Size + zllx.Msgsize()
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
		var zmia int32
		zmia, err = dc.ReadInt32()
		(*z) = Ztype(zmia)
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
		var zoxw int32
		zoxw, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Ztype(zoxw)
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
