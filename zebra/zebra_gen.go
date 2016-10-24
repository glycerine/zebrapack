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
	const maxFields0zfmh = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zfmh uint32
	totalEncodedFields0zfmh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zfmh := totalEncodedFields0zfmh
	missingFieldsLeft0zfmh := maxFields0zfmh - totalEncodedFields0zfmh

	var nextMiss0zfmh int32 = -1
	var found0zfmh [maxFields0zfmh]bool
	var curField0zfmh string

doneWithStruct0zfmh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zfmh > 0 || missingFieldsLeft0zfmh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zfmh, missingFieldsLeft0zfmh, msgp.ShowFound(found0zfmh[:]), decodeMsgFieldOrder0zfmh)
		if encodedFieldsLeft0zfmh > 0 {
			encodedFieldsLeft0zfmh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zfmh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zfmh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zfmh = 0
			}
			for nextMiss0zfmh < maxFields0zfmh && found0zfmh[nextMiss0zfmh] {
				nextMiss0zfmh++
			}
			if nextMiss0zfmh == maxFields0zfmh {
				// filled all the empty fields!
				break doneWithStruct0zfmh
			}
			missingFieldsLeft0zfmh--
			curField0zfmh = decodeMsgFieldOrder0zfmh[nextMiss0zfmh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zfmh)
		switch curField0zfmh {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zfmh[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zfmh[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found0zfmh[2] = true
			{
				var zusp int32
				zusp, err = dc.ReadInt32()
				z.Ztyp = Ztype(zusp)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0zfmh[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zfmh[4] = true
			var zohb uint32
			zohb, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zohb > 0 {
				z.Tag = make(map[string]string, zohb)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zohb > 0 {
				zohb--
				var zfsa string
				var zbhi string
				zfsa, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zbhi, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zfsa] = zbhi
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zfmh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zfmh = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

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
	var empty_znbv [5]bool
	fieldsInUse_zjnq := z.fieldsNotEmpty(empty_znbv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjnq)
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
	if !empty_znbv[3] {
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

	if !empty_znbv[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zfsa, zbhi := range z.Tag {
			err = en.WriteString(zfsa)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zbhi)
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
		for zfsa, zbhi := range z.Tag {
			o = msgp.AppendString(o, zfsa)
			o = msgp.AppendString(o, zbhi)
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
	const maxFields1zhji = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zhji uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhji, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zhji := totalEncodedFields1zhji
	missingFieldsLeft1zhji := maxFields1zhji - totalEncodedFields1zhji

	var nextMiss1zhji int32 = -1
	var found1zhji [maxFields1zhji]bool
	var curField1zhji string

doneWithStruct1zhji:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhji > 0 || missingFieldsLeft1zhji > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhji, missingFieldsLeft1zhji, msgp.ShowFound(found1zhji[:]), unmarshalMsgFieldOrder1zhji)
		if encodedFieldsLeft1zhji > 0 {
			encodedFieldsLeft1zhji--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zhji = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zhji < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhji = 0
			}
			for nextMiss1zhji < maxFields1zhji && found1zhji[nextMiss1zhji] {
				nextMiss1zhji++
			}
			if nextMiss1zhji == maxFields1zhji {
				// filled all the empty fields!
				break doneWithStruct1zhji
			}
			missingFieldsLeft1zhji--
			curField1zhji = unmarshalMsgFieldOrder1zhji[nextMiss1zhji]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhji)
		switch curField1zhji {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zhji[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zhji[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found1zhji[2] = true
			{
				var zqkw int32
				zqkw, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ztyp = Ztype(zqkw)
			}
		case "Varg":
			found1zhji[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zhji[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zwkz uint32
				zwkz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zwkz > 0 {
					z.Tag = make(map[string]string, zwkz)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zwkz > 0 {
					var zfsa string
					var zbhi string
					zwkz--
					zfsa, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zbhi, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zfsa] = zbhi
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zhji != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zhji = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zfsa, zbhi := range z.Tag {
			_ = zbhi
			s += msgp.StringPrefixSize + len(zfsa) + msgp.StringPrefixSize + len(zbhi)
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
	const maxFields2zadt = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zadt uint32
	totalEncodedFields2zadt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zadt := totalEncodedFields2zadt
	missingFieldsLeft2zadt := maxFields2zadt - totalEncodedFields2zadt

	var nextMiss2zadt int32 = -1
	var found2zadt [maxFields2zadt]bool
	var curField2zadt string

doneWithStruct2zadt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zadt > 0 || missingFieldsLeft2zadt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zadt, missingFieldsLeft2zadt, msgp.ShowFound(found2zadt[:]), decodeMsgFieldOrder2zadt)
		if encodedFieldsLeft2zadt > 0 {
			encodedFieldsLeft2zadt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zadt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zadt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zadt = 0
			}
			for nextMiss2zadt < maxFields2zadt && found2zadt[nextMiss2zadt] {
				nextMiss2zadt++
			}
			if nextMiss2zadt == maxFields2zadt {
				// filled all the empty fields!
				break doneWithStruct2zadt
			}
			missingFieldsLeft2zadt--
			curField2zadt = decodeMsgFieldOrder2zadt[nextMiss2zadt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zadt)
		switch curField2zadt {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zadt[0] = true
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
	if nextMiss2zadt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zadt = []string{"IfaceId"}

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
	const maxFields3ztvh = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3ztvh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3ztvh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3ztvh := totalEncodedFields3ztvh
	missingFieldsLeft3ztvh := maxFields3ztvh - totalEncodedFields3ztvh

	var nextMiss3ztvh int32 = -1
	var found3ztvh [maxFields3ztvh]bool
	var curField3ztvh string

doneWithStruct3ztvh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3ztvh > 0 || missingFieldsLeft3ztvh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3ztvh, missingFieldsLeft3ztvh, msgp.ShowFound(found3ztvh[:]), unmarshalMsgFieldOrder3ztvh)
		if encodedFieldsLeft3ztvh > 0 {
			encodedFieldsLeft3ztvh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3ztvh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3ztvh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3ztvh = 0
			}
			for nextMiss3ztvh < maxFields3ztvh && found3ztvh[nextMiss3ztvh] {
				nextMiss3ztvh++
			}
			if nextMiss3ztvh == maxFields3ztvh {
				// filled all the empty fields!
				break doneWithStruct3ztvh
			}
			missingFieldsLeft3ztvh--
			curField3ztvh = unmarshalMsgFieldOrder3ztvh[nextMiss3ztvh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3ztvh)
		switch curField3ztvh {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3ztvh[0] = true
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
	if nextMiss3ztvh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3ztvh = []string{"IfaceId"}

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
	const maxFields4zpmh = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zpmh uint32
	totalEncodedFields4zpmh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zpmh := totalEncodedFields4zpmh
	missingFieldsLeft4zpmh := maxFields4zpmh - totalEncodedFields4zpmh

	var nextMiss4zpmh int32 = -1
	var found4zpmh [maxFields4zpmh]bool
	var curField4zpmh string

doneWithStruct4zpmh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zpmh > 0 || missingFieldsLeft4zpmh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zpmh, missingFieldsLeft4zpmh, msgp.ShowFound(found4zpmh[:]), decodeMsgFieldOrder4zpmh)
		if encodedFieldsLeft4zpmh > 0 {
			encodedFieldsLeft4zpmh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zpmh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zpmh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zpmh = 0
			}
			for nextMiss4zpmh < maxFields4zpmh && found4zpmh[nextMiss4zpmh] {
				nextMiss4zpmh++
			}
			if nextMiss4zpmh == maxFields4zpmh {
				// filled all the empty fields!
				break doneWithStruct4zpmh
			}
			missingFieldsLeft4zpmh--
			curField4zpmh = decodeMsgFieldOrder4zpmh[nextMiss4zpmh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zpmh)
		switch curField4zpmh {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4zpmh[0] = true
			var zqnb uint32
			zqnb, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zqnb) {
				z.Methods = (z.Methods)[:zqnb]
			} else {
				z.Methods = make([]MethodT, zqnb)
			}
			for zdnk := range z.Methods {
				err = z.Methods[zdnk].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4zpmh[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zpmh[2] = true
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
	if nextMiss4zpmh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4zpmh = []string{"Methods", "Deprecated", "Comment"}

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
	for zdnk := range z.Methods {
		err = z.Methods[zdnk].EncodeMsg(en)
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
	for zdnk := range z.Methods {
		o, err = z.Methods[zdnk].MarshalMsg(o)
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
	const maxFields5zozs = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zozs uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zozs, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zozs := totalEncodedFields5zozs
	missingFieldsLeft5zozs := maxFields5zozs - totalEncodedFields5zozs

	var nextMiss5zozs int32 = -1
	var found5zozs [maxFields5zozs]bool
	var curField5zozs string

doneWithStruct5zozs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zozs > 0 || missingFieldsLeft5zozs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zozs, missingFieldsLeft5zozs, msgp.ShowFound(found5zozs[:]), unmarshalMsgFieldOrder5zozs)
		if encodedFieldsLeft5zozs > 0 {
			encodedFieldsLeft5zozs--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zozs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zozs < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zozs = 0
			}
			for nextMiss5zozs < maxFields5zozs && found5zozs[nextMiss5zozs] {
				nextMiss5zozs++
			}
			if nextMiss5zozs == maxFields5zozs {
				// filled all the empty fields!
				break doneWithStruct5zozs
			}
			missingFieldsLeft5zozs--
			curField5zozs = unmarshalMsgFieldOrder5zozs[nextMiss5zozs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zozs)
		switch curField5zozs {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zozs[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zuvx uint32
				zuvx, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zuvx) {
					z.Methods = (z.Methods)[:zuvx]
				} else {
					z.Methods = make([]MethodT, zuvx)
				}
				for zdnk := range z.Methods {
					bts, err = z.Methods[zdnk].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zozs[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zozs[2] = true
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
	if nextMiss5zozs != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zozs = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zdnk := range z.Methods {
		s += z.Methods[zdnk].Msgsize()
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
	const maxFields6zhbp = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zhbp uint32
	totalEncodedFields6zhbp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zhbp := totalEncodedFields6zhbp
	missingFieldsLeft6zhbp := maxFields6zhbp - totalEncodedFields6zhbp

	var nextMiss6zhbp int32 = -1
	var found6zhbp [maxFields6zhbp]bool
	var curField6zhbp string

doneWithStruct6zhbp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zhbp > 0 || missingFieldsLeft6zhbp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zhbp, missingFieldsLeft6zhbp, msgp.ShowFound(found6zhbp[:]), decodeMsgFieldOrder6zhbp)
		if encodedFieldsLeft6zhbp > 0 {
			encodedFieldsLeft6zhbp--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zhbp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zhbp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zhbp = 0
			}
			for nextMiss6zhbp < maxFields6zhbp && found6zhbp[nextMiss6zhbp] {
				nextMiss6zhbp++
			}
			if nextMiss6zhbp == maxFields6zhbp {
				// filled all the empty fields!
				break doneWithStruct6zhbp
			}
			missingFieldsLeft6zhbp--
			curField6zhbp = decodeMsgFieldOrder6zhbp[nextMiss6zhbp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zhbp)
		switch curField6zhbp {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zhbp[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zhbp[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zhbp[2] = true
			const maxFields7zxay = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7zxay uint32
			totalEncodedFields7zxay, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7zxay := totalEncodedFields7zxay
			missingFieldsLeft7zxay := maxFields7zxay - totalEncodedFields7zxay

			var nextMiss7zxay int32 = -1
			var found7zxay [maxFields7zxay]bool
			var curField7zxay string

		doneWithStruct7zxay:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7zxay > 0 || missingFieldsLeft7zxay > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zxay, missingFieldsLeft7zxay, msgp.ShowFound(found7zxay[:]), decodeMsgFieldOrder7zxay)
				if encodedFieldsLeft7zxay > 0 {
					encodedFieldsLeft7zxay--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7zxay = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7zxay < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7zxay = 0
					}
					for nextMiss7zxay < maxFields7zxay && found7zxay[nextMiss7zxay] {
						nextMiss7zxay++
					}
					if nextMiss7zxay == maxFields7zxay {
						// filled all the empty fields!
						break doneWithStruct7zxay
					}
					missingFieldsLeft7zxay--
					curField7zxay = decodeMsgFieldOrder7zxay[nextMiss7zxay]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7zxay)
				switch curField7zxay {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7zxay[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7zxay[1] = true
					var zrzg uint32
					zrzg, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zrzg) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zrzg]
					} else {
						z.Inputs.Fld = make([]FieldT, zrzg)
					}
					for zhke := range z.Inputs.Fld {
						err = z.Inputs.Fld[zhke].DecodeMsg(dc)
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
			if nextMiss7zxay != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zhbp[3] = true
			const maxFields8zikn = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zikn uint32
			totalEncodedFields8zikn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zikn := totalEncodedFields8zikn
			missingFieldsLeft8zikn := maxFields8zikn - totalEncodedFields8zikn

			var nextMiss8zikn int32 = -1
			var found8zikn [maxFields8zikn]bool
			var curField8zikn string

		doneWithStruct8zikn:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zikn > 0 || missingFieldsLeft8zikn > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zikn, missingFieldsLeft8zikn, msgp.ShowFound(found8zikn[:]), decodeMsgFieldOrder8zikn)
				if encodedFieldsLeft8zikn > 0 {
					encodedFieldsLeft8zikn--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zikn = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zikn < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zikn = 0
					}
					for nextMiss8zikn < maxFields8zikn && found8zikn[nextMiss8zikn] {
						nextMiss8zikn++
					}
					if nextMiss8zikn == maxFields8zikn {
						// filled all the empty fields!
						break doneWithStruct8zikn
					}
					missingFieldsLeft8zikn--
					curField8zikn = decodeMsgFieldOrder8zikn[nextMiss8zikn]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zikn)
				switch curField8zikn {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zikn[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zikn[1] = true
					var zpxb uint32
					zpxb, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zpxb) {
						z.Returns.Fld = (z.Returns.Fld)[:zpxb]
					} else {
						z.Returns.Fld = make([]FieldT, zpxb)
					}
					for zcpf := range z.Returns.Fld {
						err = z.Returns.Fld[zcpf].DecodeMsg(dc)
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
			if nextMiss8zikn != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zhbp[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zhbp[5] = true
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
	if nextMiss6zhbp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zhbp = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7zxay = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zikn = []string{"Nam", "Fld"}

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
	for zhke := range z.Inputs.Fld {
		err = z.Inputs.Fld[zhke].EncodeMsg(en)
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
	for zcpf := range z.Returns.Fld {
		err = z.Returns.Fld[zcpf].EncodeMsg(en)
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
	for zhke := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zhke].MarshalMsg(o)
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
	for zcpf := range z.Returns.Fld {
		o, err = z.Returns.Fld[zcpf].MarshalMsg(o)
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
	const maxFields9zzul = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zzul uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zzul, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zzul := totalEncodedFields9zzul
	missingFieldsLeft9zzul := maxFields9zzul - totalEncodedFields9zzul

	var nextMiss9zzul int32 = -1
	var found9zzul [maxFields9zzul]bool
	var curField9zzul string

doneWithStruct9zzul:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zzul > 0 || missingFieldsLeft9zzul > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zzul, missingFieldsLeft9zzul, msgp.ShowFound(found9zzul[:]), unmarshalMsgFieldOrder9zzul)
		if encodedFieldsLeft9zzul > 0 {
			encodedFieldsLeft9zzul--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zzul = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zzul < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zzul = 0
			}
			for nextMiss9zzul < maxFields9zzul && found9zzul[nextMiss9zzul] {
				nextMiss9zzul++
			}
			if nextMiss9zzul == maxFields9zzul {
				// filled all the empty fields!
				break doneWithStruct9zzul
			}
			missingFieldsLeft9zzul--
			curField9zzul = unmarshalMsgFieldOrder9zzul[nextMiss9zzul]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zzul)
		switch curField9zzul {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zzul[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zzul[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zzul[2] = true
			const maxFields10zjyx = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zjyx uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zjyx, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10zjyx := totalEncodedFields10zjyx
			missingFieldsLeft10zjyx := maxFields10zjyx - totalEncodedFields10zjyx

			var nextMiss10zjyx int32 = -1
			var found10zjyx [maxFields10zjyx]bool
			var curField10zjyx string

		doneWithStruct10zjyx:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10zjyx > 0 || missingFieldsLeft10zjyx > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zjyx, missingFieldsLeft10zjyx, msgp.ShowFound(found10zjyx[:]), unmarshalMsgFieldOrder10zjyx)
				if encodedFieldsLeft10zjyx > 0 {
					encodedFieldsLeft10zjyx--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10zjyx = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zjyx < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zjyx = 0
					}
					for nextMiss10zjyx < maxFields10zjyx && found10zjyx[nextMiss10zjyx] {
						nextMiss10zjyx++
					}
					if nextMiss10zjyx == maxFields10zjyx {
						// filled all the empty fields!
						break doneWithStruct10zjyx
					}
					missingFieldsLeft10zjyx--
					curField10zjyx = unmarshalMsgFieldOrder10zjyx[nextMiss10zjyx]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zjyx)
				switch curField10zjyx {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10zjyx[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10zjyx[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zrul uint32
						zrul, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zrul) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zrul]
						} else {
							z.Inputs.Fld = make([]FieldT, zrul)
						}
						for zhke := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zhke].UnmarshalMsg(bts)
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
			if nextMiss10zjyx != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zzul[3] = true
			const maxFields11zeja = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zeja uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zeja, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zeja := totalEncodedFields11zeja
			missingFieldsLeft11zeja := maxFields11zeja - totalEncodedFields11zeja

			var nextMiss11zeja int32 = -1
			var found11zeja [maxFields11zeja]bool
			var curField11zeja string

		doneWithStruct11zeja:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zeja > 0 || missingFieldsLeft11zeja > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zeja, missingFieldsLeft11zeja, msgp.ShowFound(found11zeja[:]), unmarshalMsgFieldOrder11zeja)
				if encodedFieldsLeft11zeja > 0 {
					encodedFieldsLeft11zeja--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zeja = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zeja < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zeja = 0
					}
					for nextMiss11zeja < maxFields11zeja && found11zeja[nextMiss11zeja] {
						nextMiss11zeja++
					}
					if nextMiss11zeja == maxFields11zeja {
						// filled all the empty fields!
						break doneWithStruct11zeja
					}
					missingFieldsLeft11zeja--
					curField11zeja = unmarshalMsgFieldOrder11zeja[nextMiss11zeja]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zeja)
				switch curField11zeja {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zeja[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zeja[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zijv uint32
						zijv, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zijv) {
							z.Returns.Fld = (z.Returns.Fld)[:zijv]
						} else {
							z.Returns.Fld = make([]FieldT, zijv)
						}
						for zcpf := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zcpf].UnmarshalMsg(bts)
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
			if nextMiss11zeja != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zzul[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zzul[5] = true
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
	if nextMiss9zzul != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zzul = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10zjyx = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zeja = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zhke := range z.Inputs.Fld {
		s += z.Inputs.Fld[zhke].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for zcpf := range z.Returns.Fld {
		s += z.Returns.Fld[zcpf].Msgsize()
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
	const maxFields12zkfr = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zkfr uint32
	totalEncodedFields12zkfr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zkfr := totalEncodedFields12zkfr
	missingFieldsLeft12zkfr := maxFields12zkfr - totalEncodedFields12zkfr

	var nextMiss12zkfr int32 = -1
	var found12zkfr [maxFields12zkfr]bool
	var curField12zkfr string

doneWithStruct12zkfr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zkfr > 0 || missingFieldsLeft12zkfr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zkfr, missingFieldsLeft12zkfr, msgp.ShowFound(found12zkfr[:]), decodeMsgFieldOrder12zkfr)
		if encodedFieldsLeft12zkfr > 0 {
			encodedFieldsLeft12zkfr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zkfr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zkfr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zkfr = 0
			}
			for nextMiss12zkfr < maxFields12zkfr && found12zkfr[nextMiss12zkfr] {
				nextMiss12zkfr++
			}
			if nextMiss12zkfr == maxFields12zkfr {
				// filled all the empty fields!
				break doneWithStruct12zkfr
			}
			missingFieldsLeft12zkfr--
			curField12zkfr = decodeMsgFieldOrder12zkfr[nextMiss12zkfr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zkfr)
		switch curField12zkfr {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zkfr[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found12zkfr[1] = true
			var zqks uint32
			zqks, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zqks > 0 {
				z.Structs = make(map[int64]StructT, zqks)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zqks > 0 {
				zqks--
				var zvtt int64
				var zulr StructT
				zvtt, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				const maxFields13zskr = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields13zskr uint32
				totalEncodedFields13zskr, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft13zskr := totalEncodedFields13zskr
				missingFieldsLeft13zskr := maxFields13zskr - totalEncodedFields13zskr

				var nextMiss13zskr int32 = -1
				var found13zskr [maxFields13zskr]bool
				var curField13zskr string

			doneWithStruct13zskr:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft13zskr > 0 || missingFieldsLeft13zskr > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zskr, missingFieldsLeft13zskr, msgp.ShowFound(found13zskr[:]), decodeMsgFieldOrder13zskr)
					if encodedFieldsLeft13zskr > 0 {
						encodedFieldsLeft13zskr--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField13zskr = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss13zskr < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss13zskr = 0
						}
						for nextMiss13zskr < maxFields13zskr && found13zskr[nextMiss13zskr] {
							nextMiss13zskr++
						}
						if nextMiss13zskr == maxFields13zskr {
							// filled all the empty fields!
							break doneWithStruct13zskr
						}
						missingFieldsLeft13zskr--
						curField13zskr = decodeMsgFieldOrder13zskr[nextMiss13zskr]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField13zskr)
					switch curField13zskr {
					// -- templateDecodeMsg ends here --

					case "Nam":
						found13zskr[0] = true
						zulr.Nam, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found13zskr[1] = true
						var zgqd uint32
						zgqd, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(zulr.Fld) >= int(zgqd) {
							zulr.Fld = (zulr.Fld)[:zgqd]
						} else {
							zulr.Fld = make([]FieldT, zgqd)
						}
						for zdos := range zulr.Fld {
							err = zulr.Fld[zdos].DecodeMsg(dc)
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
				if nextMiss13zskr != -1 {
					dc.PopAlwaysNil()
				}

				z.Structs[zvtt] = zulr
			}
		case "Interfaces":
			found12zkfr[2] = true
			var zwtk uint32
			zwtk, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Interfaces == nil && zwtk > 0 {
				z.Interfaces = make(map[int64]InterfaceT, zwtk)
			} else if len(z.Interfaces) > 0 {
				for key, _ := range z.Interfaces {
					delete(z.Interfaces, key)
				}
			}
			for zwtk > 0 {
				zwtk--
				var zaue int64
				var znss InterfaceT
				zaue, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = znss.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Interfaces[zaue] = znss
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zkfr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zkfr = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var decodeMsgFieldOrder13zskr = []string{"Nam", "Fld"}

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
	for zvtt, zulr := range z.Structs {
		err = en.WriteInt64(zvtt)
		if err != nil {
			panic(err)
		}
		// map header, size 2
		// write "Nam"
		err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteString(zulr.Nam)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(zulr.Fld)))
		if err != nil {
			panic(err)
		}
		for zdos := range zulr.Fld {
			err = zulr.Fld[zdos].EncodeMsg(en)
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
	for zaue, znss := range z.Interfaces {
		err = en.WriteInt64(zaue)
		if err != nil {
			panic(err)
		}
		err = znss.EncodeMsg(en)
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
	for zvtt, zulr := range z.Structs {
		o = msgp.AppendString(o, zvtt)
		// map header, size 2
		// string "Nam"
		o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
		o = msgp.AppendString(o, zulr.Nam)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(zulr.Fld)))
		for zdos := range zulr.Fld {
			o, err = zulr.Fld[zdos].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Interfaces"
	o = append(o, 0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Interfaces)))
	for zaue, znss := range z.Interfaces {
		o = msgp.AppendString(o, zaue)
		o, err = znss.MarshalMsg(o)
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
	const maxFields14zbbl = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields14zbbl uint32
	if !nbs.AlwaysNil {
		totalEncodedFields14zbbl, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft14zbbl := totalEncodedFields14zbbl
	missingFieldsLeft14zbbl := maxFields14zbbl - totalEncodedFields14zbbl

	var nextMiss14zbbl int32 = -1
	var found14zbbl [maxFields14zbbl]bool
	var curField14zbbl string

doneWithStruct14zbbl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zbbl > 0 || missingFieldsLeft14zbbl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zbbl, missingFieldsLeft14zbbl, msgp.ShowFound(found14zbbl[:]), unmarshalMsgFieldOrder14zbbl)
		if encodedFieldsLeft14zbbl > 0 {
			encodedFieldsLeft14zbbl--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField14zbbl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zbbl < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss14zbbl = 0
			}
			for nextMiss14zbbl < maxFields14zbbl && found14zbbl[nextMiss14zbbl] {
				nextMiss14zbbl++
			}
			if nextMiss14zbbl == maxFields14zbbl {
				// filled all the empty fields!
				break doneWithStruct14zbbl
			}
			missingFieldsLeft14zbbl--
			curField14zbbl = unmarshalMsgFieldOrder14zbbl[nextMiss14zbbl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zbbl)
		switch curField14zbbl {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found14zbbl[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found14zbbl[1] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zcnd uint32
				zcnd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zcnd > 0 {
					z.Structs = make(map[int64]StructT, zcnd)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zcnd > 0 {
					var zvtt string
					var zulr StructT
					zcnd--
					zvtt, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					const maxFields15zjmc = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields15zjmc uint32
					if !nbs.AlwaysNil {
						totalEncodedFields15zjmc, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft15zjmc := totalEncodedFields15zjmc
					missingFieldsLeft15zjmc := maxFields15zjmc - totalEncodedFields15zjmc

					var nextMiss15zjmc int32 = -1
					var found15zjmc [maxFields15zjmc]bool
					var curField15zjmc string

				doneWithStruct15zjmc:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft15zjmc > 0 || missingFieldsLeft15zjmc > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zjmc, missingFieldsLeft15zjmc, msgp.ShowFound(found15zjmc[:]), unmarshalMsgFieldOrder15zjmc)
						if encodedFieldsLeft15zjmc > 0 {
							encodedFieldsLeft15zjmc--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField15zjmc = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss15zjmc < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss15zjmc = 0
							}
							for nextMiss15zjmc < maxFields15zjmc && found15zjmc[nextMiss15zjmc] {
								nextMiss15zjmc++
							}
							if nextMiss15zjmc == maxFields15zjmc {
								// filled all the empty fields!
								break doneWithStruct15zjmc
							}
							missingFieldsLeft15zjmc--
							curField15zjmc = unmarshalMsgFieldOrder15zjmc[nextMiss15zjmc]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField15zjmc)
						switch curField15zjmc {
						// -- templateUnmarshalMsg ends here --

						case "Nam":
							found15zjmc[0] = true
							zulr.Nam, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found15zjmc[1] = true
							if nbs.AlwaysNil {
								(zulr.Fld) = (zulr.Fld)[:0]
							} else {

								var zpwp uint32
								zpwp, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(zulr.Fld) >= int(zpwp) {
									zulr.Fld = (zulr.Fld)[:zpwp]
								} else {
									zulr.Fld = make([]FieldT, zpwp)
								}
								for zdos := range zulr.Fld {
									bts, err = zulr.Fld[zdos].UnmarshalMsg(bts)
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
					if nextMiss15zjmc != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Structs[zvtt] = zulr
				}
			}
		case "Interfaces":
			found14zbbl[2] = true
			if nbs.AlwaysNil {
				if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}

			} else {

				var zmby uint32
				zmby, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Interfaces == nil && zmby > 0 {
					z.Interfaces = make(map[int64]InterfaceT, zmby)
				} else if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}
				for zmby > 0 {
					var zaue string
					var znss InterfaceT
					zmby--
					zaue, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = znss.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Interfaces[zaue] = znss
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zbbl != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder14zbbl = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var unmarshalMsgFieldOrder15zjmc = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zvtt, zulr := range z.Structs {
			_ = zulr
			s += msgp.StringPrefixSize + len(zvtt) + 1 + 4 + msgp.StringPrefixSize + len(zulr.Nam) + 4 + msgp.ArrayHeaderSize
			for zdos := range zulr.Fld {
				s += zulr.Fld[zdos].Msgsize()
			}
		}
	}
	s += 11 + msgp.MapHeaderSize
	if z.Interfaces != nil {
		for zaue, znss := range z.Interfaces {
			_ = znss
			s += msgp.StringPrefixSize + len(zaue) + znss.Msgsize()
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
	const maxFields16zqtd = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zqtd uint32
	totalEncodedFields16zqtd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zqtd := totalEncodedFields16zqtd
	missingFieldsLeft16zqtd := maxFields16zqtd - totalEncodedFields16zqtd

	var nextMiss16zqtd int32 = -1
	var found16zqtd [maxFields16zqtd]bool
	var curField16zqtd string

doneWithStruct16zqtd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zqtd > 0 || missingFieldsLeft16zqtd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zqtd, missingFieldsLeft16zqtd, msgp.ShowFound(found16zqtd[:]), decodeMsgFieldOrder16zqtd)
		if encodedFieldsLeft16zqtd > 0 {
			encodedFieldsLeft16zqtd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zqtd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zqtd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zqtd = 0
			}
			for nextMiss16zqtd < maxFields16zqtd && found16zqtd[nextMiss16zqtd] {
				nextMiss16zqtd++
			}
			if nextMiss16zqtd == maxFields16zqtd {
				// filled all the empty fields!
				break doneWithStruct16zqtd
			}
			missingFieldsLeft16zqtd--
			curField16zqtd = decodeMsgFieldOrder16zqtd[nextMiss16zqtd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zqtd)
		switch curField16zqtd {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found16zqtd[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found16zqtd[1] = true
			var zfde uint32
			zfde, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zfde) {
				z.Fld = (z.Fld)[:zfde]
			} else {
				z.Fld = make([]FieldT, zfde)
			}
			for zesg := range z.Fld {
				err = z.Fld[zesg].DecodeMsg(dc)
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
	if nextMiss16zqtd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder16zqtd = []string{"Nam", "Fld"}

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
	for zesg := range z.Fld {
		err = z.Fld[zesg].EncodeMsg(en)
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
	for zesg := range z.Fld {
		o, err = z.Fld[zesg].MarshalMsg(o)
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
	const maxFields17zwpe = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17zwpe uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17zwpe, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17zwpe := totalEncodedFields17zwpe
	missingFieldsLeft17zwpe := maxFields17zwpe - totalEncodedFields17zwpe

	var nextMiss17zwpe int32 = -1
	var found17zwpe [maxFields17zwpe]bool
	var curField17zwpe string

doneWithStruct17zwpe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17zwpe > 0 || missingFieldsLeft17zwpe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zwpe, missingFieldsLeft17zwpe, msgp.ShowFound(found17zwpe[:]), unmarshalMsgFieldOrder17zwpe)
		if encodedFieldsLeft17zwpe > 0 {
			encodedFieldsLeft17zwpe--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17zwpe = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17zwpe < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17zwpe = 0
			}
			for nextMiss17zwpe < maxFields17zwpe && found17zwpe[nextMiss17zwpe] {
				nextMiss17zwpe++
			}
			if nextMiss17zwpe == maxFields17zwpe {
				// filled all the empty fields!
				break doneWithStruct17zwpe
			}
			missingFieldsLeft17zwpe--
			curField17zwpe = unmarshalMsgFieldOrder17zwpe[nextMiss17zwpe]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17zwpe)
		switch curField17zwpe {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found17zwpe[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found17zwpe[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zrty uint32
				zrty, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zrty) {
					z.Fld = (z.Fld)[:zrty]
				} else {
					z.Fld = make([]FieldT, zrty)
				}
				for zesg := range z.Fld {
					bts, err = z.Fld[zesg].UnmarshalMsg(bts)
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
	if nextMiss17zwpe != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder17zwpe = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zesg := range z.Fld {
		s += z.Fld[zesg].Msgsize()
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
		var zuof int64
		zuof, err = dc.ReadInt64()
		(*z) = StructTypeId(zuof)
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
		var zpkr int64
		zpkr, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zpkr)
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
	const maxFields18zmdo = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields18zmdo uint32
	totalEncodedFields18zmdo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft18zmdo := totalEncodedFields18zmdo
	missingFieldsLeft18zmdo := maxFields18zmdo - totalEncodedFields18zmdo

	var nextMiss18zmdo int32 = -1
	var found18zmdo [maxFields18zmdo]bool
	var curField18zmdo string

doneWithStruct18zmdo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft18zmdo > 0 || missingFieldsLeft18zmdo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zmdo, missingFieldsLeft18zmdo, msgp.ShowFound(found18zmdo[:]), decodeMsgFieldOrder18zmdo)
		if encodedFieldsLeft18zmdo > 0 {
			encodedFieldsLeft18zmdo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField18zmdo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss18zmdo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss18zmdo = 0
			}
			for nextMiss18zmdo < maxFields18zmdo && found18zmdo[nextMiss18zmdo] {
				nextMiss18zmdo++
			}
			if nextMiss18zmdo == maxFields18zmdo {
				// filled all the empty fields!
				break doneWithStruct18zmdo
			}
			missingFieldsLeft18zmdo--
			curField18zmdo = decodeMsgFieldOrder18zmdo[nextMiss18zmdo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField18zmdo)
		switch curField18zmdo {
		// -- templateDecodeMsg ends here --

		case "Ty":
			found18zmdo[0] = true
			{
				var zwit int32
				zwit, err = dc.ReadInt32()
				z.Ty = Ztype(zwit)
			}
			if err != nil {
				panic(err)
			}
		case "Id":
			found18zmdo[1] = true
			{
				var zfmr int64
				zfmr, err = dc.ReadInt64()
				z.Id = StructTypeId(zfmr)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found18zmdo[2] = true
			var zcmo uint32
			zcmo, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zcmo > 0 {
				z.Da = make(map[int64]msgp.Raw, zcmo)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zcmo > 0 {
				zcmo--
				var zfee int64
				var znbo msgp.Raw
				zfee, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = znbo.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zfee] = znbo
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss18zmdo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZPacket
var decodeMsgFieldOrder18zmdo = []string{"Ty", "Id", "Da"}

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
	for zfee, znbo := range z.Da {
		err = en.WriteInt64(zfee)
		if err != nil {
			panic(err)
		}
		err = znbo.EncodeMsg(en)
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
	for zfee, znbo := range z.Da {
		o = msgp.AppendString(o, zfee)
		o, err = znbo.MarshalMsg(o)
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
	const maxFields19zond = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19zond uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19zond, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19zond := totalEncodedFields19zond
	missingFieldsLeft19zond := maxFields19zond - totalEncodedFields19zond

	var nextMiss19zond int32 = -1
	var found19zond [maxFields19zond]bool
	var curField19zond string

doneWithStruct19zond:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zond > 0 || missingFieldsLeft19zond > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zond, missingFieldsLeft19zond, msgp.ShowFound(found19zond[:]), unmarshalMsgFieldOrder19zond)
		if encodedFieldsLeft19zond > 0 {
			encodedFieldsLeft19zond--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19zond = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zond < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19zond = 0
			}
			for nextMiss19zond < maxFields19zond && found19zond[nextMiss19zond] {
				nextMiss19zond++
			}
			if nextMiss19zond == maxFields19zond {
				// filled all the empty fields!
				break doneWithStruct19zond
			}
			missingFieldsLeft19zond--
			curField19zond = unmarshalMsgFieldOrder19zond[nextMiss19zond]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zond)
		switch curField19zond {
		// -- templateUnmarshalMsg ends here --

		case "Ty":
			found19zond[0] = true
			{
				var zyfj int32
				zyfj, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ty = Ztype(zyfj)
			}
		case "Id":
			found19zond[1] = true
			{
				var zeih int64
				zeih, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zeih)
			}
		case "Da":
			found19zond[2] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zpdp uint32
				zpdp, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zpdp > 0 {
					z.Da = make(map[int64]msgp.Raw, zpdp)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zpdp > 0 {
					var zfee string
					var znbo msgp.Raw
					zpdp--
					zfee, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = znbo.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zfee] = znbo
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19zond != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZPacket
var unmarshalMsgFieldOrder19zond = []string{"Ty", "Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zfee, znbo := range z.Da {
			_ = znbo
			s += msgp.StringPrefixSize + len(zfee) + znbo.Msgsize()
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
		var zsyg int32
		zsyg, err = dc.ReadInt32()
		(*z) = Ztype(zsyg)
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
		var zzuu int32
		zzuu, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Ztype(zzuu)
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
