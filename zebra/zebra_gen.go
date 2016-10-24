package msgp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

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
	const maxFields0zpha = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zpha uint32
	totalEncodedFields0zpha, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zpha := totalEncodedFields0zpha
	missingFieldsLeft0zpha := maxFields0zpha - totalEncodedFields0zpha

	var nextMiss0zpha int32 = -1
	var found0zpha [maxFields0zpha]bool
	var curField0zpha string

doneWithStruct0zpha:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zpha > 0 || missingFieldsLeft0zpha > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zpha, missingFieldsLeft0zpha, msgp.ShowFound(found0zpha[:]), decodeMsgFieldOrder0zpha)
		if encodedFieldsLeft0zpha > 0 {
			encodedFieldsLeft0zpha--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zpha = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zpha < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zpha = 0
			}
			for nextMiss0zpha < maxFields0zpha && found0zpha[nextMiss0zpha] {
				nextMiss0zpha++
			}
			if nextMiss0zpha == maxFields0zpha {
				// filled all the empty fields!
				break doneWithStruct0zpha
			}
			missingFieldsLeft0zpha--
			curField0zpha = decodeMsgFieldOrder0zpha[nextMiss0zpha]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zpha)
		switch curField0zpha {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zpha[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zpha[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found0zpha[2] = true
			{
				var zgha int32
				zgha, err = dc.ReadInt32()
				z.Ztyp = Ztype(zgha)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0zpha[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zpha[4] = true
			var zgzp uint32
			zgzp, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zgzp > 0 {
				z.Tag = make(map[string]string, zgzp)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zgzp > 0 {
				zgzp--
				var zkji string
				var zsfp string
				zkji, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zsfp, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zkji] = zsfp
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zpha != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zpha = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

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
	var empty_ztdl [5]bool
	fieldsInUse_zuzk := z.fieldsNotEmpty(empty_ztdl[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zuzk)
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
	if !empty_ztdl[3] {
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

	if !empty_ztdl[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zkji, zsfp := range z.Tag {
			err = en.WriteString(zkji)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zsfp)
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
		for zkji, zsfp := range z.Tag {
			o = msgp.AppendString(o, zkji)
			o = msgp.AppendString(o, zsfp)
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
	const maxFields1zcnh = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zcnh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zcnh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zcnh := totalEncodedFields1zcnh
	missingFieldsLeft1zcnh := maxFields1zcnh - totalEncodedFields1zcnh

	var nextMiss1zcnh int32 = -1
	var found1zcnh [maxFields1zcnh]bool
	var curField1zcnh string

doneWithStruct1zcnh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zcnh > 0 || missingFieldsLeft1zcnh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zcnh, missingFieldsLeft1zcnh, msgp.ShowFound(found1zcnh[:]), unmarshalMsgFieldOrder1zcnh)
		if encodedFieldsLeft1zcnh > 0 {
			encodedFieldsLeft1zcnh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zcnh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zcnh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zcnh = 0
			}
			for nextMiss1zcnh < maxFields1zcnh && found1zcnh[nextMiss1zcnh] {
				nextMiss1zcnh++
			}
			if nextMiss1zcnh == maxFields1zcnh {
				// filled all the empty fields!
				break doneWithStruct1zcnh
			}
			missingFieldsLeft1zcnh--
			curField1zcnh = unmarshalMsgFieldOrder1zcnh[nextMiss1zcnh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zcnh)
		switch curField1zcnh {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zcnh[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zcnh[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found1zcnh[2] = true
			{
				var zsve int32
				zsve, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ztyp = Ztype(zsve)
			}
		case "Varg":
			found1zcnh[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zcnh[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zptc uint32
				zptc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zptc > 0 {
					z.Tag = make(map[string]string, zptc)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zptc > 0 {
					var zkji string
					var zsfp string
					zptc--
					zkji, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zsfp, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zkji] = zsfp
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zcnh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zcnh = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zkji, zsfp := range z.Tag {
			_ = zsfp
			s += msgp.StringPrefixSize + len(zkji) + msgp.StringPrefixSize + len(zsfp)
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
	const maxFields2zand = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zand uint32
	totalEncodedFields2zand, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zand := totalEncodedFields2zand
	missingFieldsLeft2zand := maxFields2zand - totalEncodedFields2zand

	var nextMiss2zand int32 = -1
	var found2zand [maxFields2zand]bool
	var curField2zand string

doneWithStruct2zand:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zand > 0 || missingFieldsLeft2zand > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zand, missingFieldsLeft2zand, msgp.ShowFound(found2zand[:]), decodeMsgFieldOrder2zand)
		if encodedFieldsLeft2zand > 0 {
			encodedFieldsLeft2zand--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zand = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zand < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zand = 0
			}
			for nextMiss2zand < maxFields2zand && found2zand[nextMiss2zand] {
				nextMiss2zand++
			}
			if nextMiss2zand == maxFields2zand {
				// filled all the empty fields!
				break doneWithStruct2zand
			}
			missingFieldsLeft2zand--
			curField2zand = decodeMsgFieldOrder2zand[nextMiss2zand]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zand)
		switch curField2zand {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zand[0] = true
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
	if nextMiss2zand != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zand = []string{"IfaceId"}

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
	const maxFields3ztum = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3ztum uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3ztum, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3ztum := totalEncodedFields3ztum
	missingFieldsLeft3ztum := maxFields3ztum - totalEncodedFields3ztum

	var nextMiss3ztum int32 = -1
	var found3ztum [maxFields3ztum]bool
	var curField3ztum string

doneWithStruct3ztum:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3ztum > 0 || missingFieldsLeft3ztum > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3ztum, missingFieldsLeft3ztum, msgp.ShowFound(found3ztum[:]), unmarshalMsgFieldOrder3ztum)
		if encodedFieldsLeft3ztum > 0 {
			encodedFieldsLeft3ztum--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3ztum = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3ztum < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3ztum = 0
			}
			for nextMiss3ztum < maxFields3ztum && found3ztum[nextMiss3ztum] {
				nextMiss3ztum++
			}
			if nextMiss3ztum == maxFields3ztum {
				// filled all the empty fields!
				break doneWithStruct3ztum
			}
			missingFieldsLeft3ztum--
			curField3ztum = unmarshalMsgFieldOrder3ztum[nextMiss3ztum]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3ztum)
		switch curField3ztum {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3ztum[0] = true
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
	if nextMiss3ztum != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3ztum = []string{"IfaceId"}

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
	const maxFields4ziuc = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4ziuc uint32
	totalEncodedFields4ziuc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4ziuc := totalEncodedFields4ziuc
	missingFieldsLeft4ziuc := maxFields4ziuc - totalEncodedFields4ziuc

	var nextMiss4ziuc int32 = -1
	var found4ziuc [maxFields4ziuc]bool
	var curField4ziuc string

doneWithStruct4ziuc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4ziuc > 0 || missingFieldsLeft4ziuc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4ziuc, missingFieldsLeft4ziuc, msgp.ShowFound(found4ziuc[:]), decodeMsgFieldOrder4ziuc)
		if encodedFieldsLeft4ziuc > 0 {
			encodedFieldsLeft4ziuc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4ziuc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4ziuc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4ziuc = 0
			}
			for nextMiss4ziuc < maxFields4ziuc && found4ziuc[nextMiss4ziuc] {
				nextMiss4ziuc++
			}
			if nextMiss4ziuc == maxFields4ziuc {
				// filled all the empty fields!
				break doneWithStruct4ziuc
			}
			missingFieldsLeft4ziuc--
			curField4ziuc = decodeMsgFieldOrder4ziuc[nextMiss4ziuc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4ziuc)
		switch curField4ziuc {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4ziuc[0] = true
			var zlan uint32
			zlan, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zlan) {
				z.Methods = (z.Methods)[:zlan]
			} else {
				z.Methods = make([]MethodT, zlan)
			}
			for zdky := range z.Methods {
				err = z.Methods[zdky].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4ziuc[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4ziuc[2] = true
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
	if nextMiss4ziuc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4ziuc = []string{"Methods", "Deprecated", "Comment"}

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
	for zdky := range z.Methods {
		err = z.Methods[zdky].EncodeMsg(en)
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
	for zdky := range z.Methods {
		o, err = z.Methods[zdky].MarshalMsg(o)
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
	const maxFields5zqke = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zqke uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zqke, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zqke := totalEncodedFields5zqke
	missingFieldsLeft5zqke := maxFields5zqke - totalEncodedFields5zqke

	var nextMiss5zqke int32 = -1
	var found5zqke [maxFields5zqke]bool
	var curField5zqke string

doneWithStruct5zqke:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zqke > 0 || missingFieldsLeft5zqke > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zqke, missingFieldsLeft5zqke, msgp.ShowFound(found5zqke[:]), unmarshalMsgFieldOrder5zqke)
		if encodedFieldsLeft5zqke > 0 {
			encodedFieldsLeft5zqke--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zqke = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zqke < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zqke = 0
			}
			for nextMiss5zqke < maxFields5zqke && found5zqke[nextMiss5zqke] {
				nextMiss5zqke++
			}
			if nextMiss5zqke == maxFields5zqke {
				// filled all the empty fields!
				break doneWithStruct5zqke
			}
			missingFieldsLeft5zqke--
			curField5zqke = unmarshalMsgFieldOrder5zqke[nextMiss5zqke]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zqke)
		switch curField5zqke {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zqke[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zqjd uint32
				zqjd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zqjd) {
					z.Methods = (z.Methods)[:zqjd]
				} else {
					z.Methods = make([]MethodT, zqjd)
				}
				for zdky := range z.Methods {
					bts, err = z.Methods[zdky].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zqke[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zqke[2] = true
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
	if nextMiss5zqke != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zqke = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zdky := range z.Methods {
		s += z.Methods[zdky].Msgsize()
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
	const maxFields6zrcm = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zrcm uint32
	totalEncodedFields6zrcm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zrcm := totalEncodedFields6zrcm
	missingFieldsLeft6zrcm := maxFields6zrcm - totalEncodedFields6zrcm

	var nextMiss6zrcm int32 = -1
	var found6zrcm [maxFields6zrcm]bool
	var curField6zrcm string

doneWithStruct6zrcm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zrcm > 0 || missingFieldsLeft6zrcm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zrcm, missingFieldsLeft6zrcm, msgp.ShowFound(found6zrcm[:]), decodeMsgFieldOrder6zrcm)
		if encodedFieldsLeft6zrcm > 0 {
			encodedFieldsLeft6zrcm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zrcm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zrcm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zrcm = 0
			}
			for nextMiss6zrcm < maxFields6zrcm && found6zrcm[nextMiss6zrcm] {
				nextMiss6zrcm++
			}
			if nextMiss6zrcm == maxFields6zrcm {
				// filled all the empty fields!
				break doneWithStruct6zrcm
			}
			missingFieldsLeft6zrcm--
			curField6zrcm = decodeMsgFieldOrder6zrcm[nextMiss6zrcm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zrcm)
		switch curField6zrcm {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zrcm[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zrcm[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zrcm[2] = true
			const maxFields7zdij = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7zdij uint32
			totalEncodedFields7zdij, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7zdij := totalEncodedFields7zdij
			missingFieldsLeft7zdij := maxFields7zdij - totalEncodedFields7zdij

			var nextMiss7zdij int32 = -1
			var found7zdij [maxFields7zdij]bool
			var curField7zdij string

		doneWithStruct7zdij:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7zdij > 0 || missingFieldsLeft7zdij > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zdij, missingFieldsLeft7zdij, msgp.ShowFound(found7zdij[:]), decodeMsgFieldOrder7zdij)
				if encodedFieldsLeft7zdij > 0 {
					encodedFieldsLeft7zdij--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7zdij = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7zdij < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7zdij = 0
					}
					for nextMiss7zdij < maxFields7zdij && found7zdij[nextMiss7zdij] {
						nextMiss7zdij++
					}
					if nextMiss7zdij == maxFields7zdij {
						// filled all the empty fields!
						break doneWithStruct7zdij
					}
					missingFieldsLeft7zdij--
					curField7zdij = decodeMsgFieldOrder7zdij[nextMiss7zdij]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7zdij)
				switch curField7zdij {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7zdij[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7zdij[1] = true
					var zkai uint32
					zkai, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zkai) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zkai]
					} else {
						z.Inputs.Fld = make([]FieldT, zkai)
					}
					for zlyi := range z.Inputs.Fld {
						err = z.Inputs.Fld[zlyi].DecodeMsg(dc)
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
			if nextMiss7zdij != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zrcm[3] = true
			const maxFields8zrqb = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zrqb uint32
			totalEncodedFields8zrqb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zrqb := totalEncodedFields8zrqb
			missingFieldsLeft8zrqb := maxFields8zrqb - totalEncodedFields8zrqb

			var nextMiss8zrqb int32 = -1
			var found8zrqb [maxFields8zrqb]bool
			var curField8zrqb string

		doneWithStruct8zrqb:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zrqb > 0 || missingFieldsLeft8zrqb > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zrqb, missingFieldsLeft8zrqb, msgp.ShowFound(found8zrqb[:]), decodeMsgFieldOrder8zrqb)
				if encodedFieldsLeft8zrqb > 0 {
					encodedFieldsLeft8zrqb--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zrqb = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zrqb < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zrqb = 0
					}
					for nextMiss8zrqb < maxFields8zrqb && found8zrqb[nextMiss8zrqb] {
						nextMiss8zrqb++
					}
					if nextMiss8zrqb == maxFields8zrqb {
						// filled all the empty fields!
						break doneWithStruct8zrqb
					}
					missingFieldsLeft8zrqb--
					curField8zrqb = decodeMsgFieldOrder8zrqb[nextMiss8zrqb]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zrqb)
				switch curField8zrqb {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zrqb[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zrqb[1] = true
					var zqun uint32
					zqun, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zqun) {
						z.Returns.Fld = (z.Returns.Fld)[:zqun]
					} else {
						z.Returns.Fld = make([]FieldT, zqun)
					}
					for zizf := range z.Returns.Fld {
						err = z.Returns.Fld[zizf].DecodeMsg(dc)
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
			if nextMiss8zrqb != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zrcm[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zrcm[5] = true
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
	if nextMiss6zrcm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zrcm = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7zdij = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zrqb = []string{"Nam", "Fld"}

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
	for zlyi := range z.Inputs.Fld {
		err = z.Inputs.Fld[zlyi].EncodeMsg(en)
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
	for zizf := range z.Returns.Fld {
		err = z.Returns.Fld[zizf].EncodeMsg(en)
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
	for zlyi := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zlyi].MarshalMsg(o)
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
	for zizf := range z.Returns.Fld {
		o, err = z.Returns.Fld[zizf].MarshalMsg(o)
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
	const maxFields9zteg = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zteg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zteg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zteg := totalEncodedFields9zteg
	missingFieldsLeft9zteg := maxFields9zteg - totalEncodedFields9zteg

	var nextMiss9zteg int32 = -1
	var found9zteg [maxFields9zteg]bool
	var curField9zteg string

doneWithStruct9zteg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zteg > 0 || missingFieldsLeft9zteg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zteg, missingFieldsLeft9zteg, msgp.ShowFound(found9zteg[:]), unmarshalMsgFieldOrder9zteg)
		if encodedFieldsLeft9zteg > 0 {
			encodedFieldsLeft9zteg--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zteg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zteg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zteg = 0
			}
			for nextMiss9zteg < maxFields9zteg && found9zteg[nextMiss9zteg] {
				nextMiss9zteg++
			}
			if nextMiss9zteg == maxFields9zteg {
				// filled all the empty fields!
				break doneWithStruct9zteg
			}
			missingFieldsLeft9zteg--
			curField9zteg = unmarshalMsgFieldOrder9zteg[nextMiss9zteg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zteg)
		switch curField9zteg {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zteg[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zteg[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zteg[2] = true
			const maxFields10ztjj = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10ztjj uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10ztjj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10ztjj := totalEncodedFields10ztjj
			missingFieldsLeft10ztjj := maxFields10ztjj - totalEncodedFields10ztjj

			var nextMiss10ztjj int32 = -1
			var found10ztjj [maxFields10ztjj]bool
			var curField10ztjj string

		doneWithStruct10ztjj:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10ztjj > 0 || missingFieldsLeft10ztjj > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10ztjj, missingFieldsLeft10ztjj, msgp.ShowFound(found10ztjj[:]), unmarshalMsgFieldOrder10ztjj)
				if encodedFieldsLeft10ztjj > 0 {
					encodedFieldsLeft10ztjj--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10ztjj = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10ztjj < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10ztjj = 0
					}
					for nextMiss10ztjj < maxFields10ztjj && found10ztjj[nextMiss10ztjj] {
						nextMiss10ztjj++
					}
					if nextMiss10ztjj == maxFields10ztjj {
						// filled all the empty fields!
						break doneWithStruct10ztjj
					}
					missingFieldsLeft10ztjj--
					curField10ztjj = unmarshalMsgFieldOrder10ztjj[nextMiss10ztjj]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10ztjj)
				switch curField10ztjj {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10ztjj[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10ztjj[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zafh uint32
						zafh, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zafh) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zafh]
						} else {
							z.Inputs.Fld = make([]FieldT, zafh)
						}
						for zlyi := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zlyi].UnmarshalMsg(bts)
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
			if nextMiss10ztjj != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zteg[3] = true
			const maxFields11zcxw = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zcxw uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zcxw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zcxw := totalEncodedFields11zcxw
			missingFieldsLeft11zcxw := maxFields11zcxw - totalEncodedFields11zcxw

			var nextMiss11zcxw int32 = -1
			var found11zcxw [maxFields11zcxw]bool
			var curField11zcxw string

		doneWithStruct11zcxw:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zcxw > 0 || missingFieldsLeft11zcxw > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zcxw, missingFieldsLeft11zcxw, msgp.ShowFound(found11zcxw[:]), unmarshalMsgFieldOrder11zcxw)
				if encodedFieldsLeft11zcxw > 0 {
					encodedFieldsLeft11zcxw--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zcxw = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zcxw < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zcxw = 0
					}
					for nextMiss11zcxw < maxFields11zcxw && found11zcxw[nextMiss11zcxw] {
						nextMiss11zcxw++
					}
					if nextMiss11zcxw == maxFields11zcxw {
						// filled all the empty fields!
						break doneWithStruct11zcxw
					}
					missingFieldsLeft11zcxw--
					curField11zcxw = unmarshalMsgFieldOrder11zcxw[nextMiss11zcxw]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zcxw)
				switch curField11zcxw {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zcxw[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zcxw[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zksa uint32
						zksa, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zksa) {
							z.Returns.Fld = (z.Returns.Fld)[:zksa]
						} else {
							z.Returns.Fld = make([]FieldT, zksa)
						}
						for zizf := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zizf].UnmarshalMsg(bts)
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
			if nextMiss11zcxw != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zteg[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zteg[5] = true
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
	if nextMiss9zteg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zteg = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10ztjj = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zcxw = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zlyi := range z.Inputs.Fld {
		s += z.Inputs.Fld[zlyi].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for zizf := range z.Returns.Fld {
		s += z.Returns.Fld[zizf].Msgsize()
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
	const maxFields12zvke = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zvke uint32
	totalEncodedFields12zvke, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zvke := totalEncodedFields12zvke
	missingFieldsLeft12zvke := maxFields12zvke - totalEncodedFields12zvke

	var nextMiss12zvke int32 = -1
	var found12zvke [maxFields12zvke]bool
	var curField12zvke string

doneWithStruct12zvke:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zvke > 0 || missingFieldsLeft12zvke > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zvke, missingFieldsLeft12zvke, msgp.ShowFound(found12zvke[:]), decodeMsgFieldOrder12zvke)
		if encodedFieldsLeft12zvke > 0 {
			encodedFieldsLeft12zvke--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zvke = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zvke < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zvke = 0
			}
			for nextMiss12zvke < maxFields12zvke && found12zvke[nextMiss12zvke] {
				nextMiss12zvke++
			}
			if nextMiss12zvke == maxFields12zvke {
				// filled all the empty fields!
				break doneWithStruct12zvke
			}
			missingFieldsLeft12zvke--
			curField12zvke = decodeMsgFieldOrder12zvke[nextMiss12zvke]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zvke)
		switch curField12zvke {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zvke[0] = true
			z.PkgPath, err = dc.ReadString()
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
	if nextMiss12zvke != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zvke = []string{"PkgPath"}

// fieldsNotEmpty supports omitempty tags
func (z SchemaT) fieldsNotEmpty(isempty []bool) uint32 {
	return 1
}

// EncodeMsg implements msgp.Encodable
func (z SchemaT) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "PkgPath"
	err = en.Append(0x81, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.PkgPath)
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SchemaT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "PkgPath"
	o = append(o, 0x81, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.PkgPath)
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
	const maxFields13zyec = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields13zyec uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zyec, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft13zyec := totalEncodedFields13zyec
	missingFieldsLeft13zyec := maxFields13zyec - totalEncodedFields13zyec

	var nextMiss13zyec int32 = -1
	var found13zyec [maxFields13zyec]bool
	var curField13zyec string

doneWithStruct13zyec:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zyec > 0 || missingFieldsLeft13zyec > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zyec, missingFieldsLeft13zyec, msgp.ShowFound(found13zyec[:]), unmarshalMsgFieldOrder13zyec)
		if encodedFieldsLeft13zyec > 0 {
			encodedFieldsLeft13zyec--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField13zyec = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss13zyec < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zyec = 0
			}
			for nextMiss13zyec < maxFields13zyec && found13zyec[nextMiss13zyec] {
				nextMiss13zyec++
			}
			if nextMiss13zyec == maxFields13zyec {
				// filled all the empty fields!
				break doneWithStruct13zyec
			}
			missingFieldsLeft13zyec--
			curField13zyec = unmarshalMsgFieldOrder13zyec[nextMiss13zyec]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zyec)
		switch curField13zyec {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found13zyec[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss13zyec != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder13zyec = []string{"PkgPath"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath)
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
	const maxFields14zlln = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields14zlln uint32
	totalEncodedFields14zlln, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zlln := totalEncodedFields14zlln
	missingFieldsLeft14zlln := maxFields14zlln - totalEncodedFields14zlln

	var nextMiss14zlln int32 = -1
	var found14zlln [maxFields14zlln]bool
	var curField14zlln string

doneWithStruct14zlln:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zlln > 0 || missingFieldsLeft14zlln > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zlln, missingFieldsLeft14zlln, msgp.ShowFound(found14zlln[:]), decodeMsgFieldOrder14zlln)
		if encodedFieldsLeft14zlln > 0 {
			encodedFieldsLeft14zlln--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField14zlln = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zlln < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zlln = 0
			}
			for nextMiss14zlln < maxFields14zlln && found14zlln[nextMiss14zlln] {
				nextMiss14zlln++
			}
			if nextMiss14zlln == maxFields14zlln {
				// filled all the empty fields!
				break doneWithStruct14zlln
			}
			missingFieldsLeft14zlln--
			curField14zlln = decodeMsgFieldOrder14zlln[nextMiss14zlln]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zlln)
		switch curField14zlln {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found14zlln[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found14zlln[1] = true
			var zeqs uint32
			zeqs, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zeqs) {
				z.Fld = (z.Fld)[:zeqs]
			} else {
				z.Fld = make([]FieldT, zeqs)
			}
			for zwvc := range z.Fld {
				err = z.Fld[zwvc].DecodeMsg(dc)
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
	if nextMiss14zlln != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder14zlln = []string{"Nam", "Fld"}

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
	for zwvc := range z.Fld {
		err = z.Fld[zwvc].EncodeMsg(en)
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
	for zwvc := range z.Fld {
		o, err = z.Fld[zwvc].MarshalMsg(o)
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
	const maxFields15zkrx = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields15zkrx uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zkrx, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft15zkrx := totalEncodedFields15zkrx
	missingFieldsLeft15zkrx := maxFields15zkrx - totalEncodedFields15zkrx

	var nextMiss15zkrx int32 = -1
	var found15zkrx [maxFields15zkrx]bool
	var curField15zkrx string

doneWithStruct15zkrx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zkrx > 0 || missingFieldsLeft15zkrx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zkrx, missingFieldsLeft15zkrx, msgp.ShowFound(found15zkrx[:]), unmarshalMsgFieldOrder15zkrx)
		if encodedFieldsLeft15zkrx > 0 {
			encodedFieldsLeft15zkrx--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField15zkrx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss15zkrx < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zkrx = 0
			}
			for nextMiss15zkrx < maxFields15zkrx && found15zkrx[nextMiss15zkrx] {
				nextMiss15zkrx++
			}
			if nextMiss15zkrx == maxFields15zkrx {
				// filled all the empty fields!
				break doneWithStruct15zkrx
			}
			missingFieldsLeft15zkrx--
			curField15zkrx = unmarshalMsgFieldOrder15zkrx[nextMiss15zkrx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zkrx)
		switch curField15zkrx {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found15zkrx[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found15zkrx[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zkpk uint32
				zkpk, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zkpk) {
					z.Fld = (z.Fld)[:zkpk]
				} else {
					z.Fld = make([]FieldT, zkpk)
				}
				for zwvc := range z.Fld {
					bts, err = z.Fld[zwvc].UnmarshalMsg(bts)
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
	if nextMiss15zkrx != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder15zkrx = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zwvc := range z.Fld {
		s += z.Fld[zwvc].Msgsize()
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
		var zmzx int64
		zmzx, err = dc.ReadInt64()
		(*z) = StructTypeId(zmzx)
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
		var zold int64
		zold, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zold)
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
	const maxFields16zvar = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zvar uint32
	totalEncodedFields16zvar, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zvar := totalEncodedFields16zvar
	missingFieldsLeft16zvar := maxFields16zvar - totalEncodedFields16zvar

	var nextMiss16zvar int32 = -1
	var found16zvar [maxFields16zvar]bool
	var curField16zvar string

doneWithStruct16zvar:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zvar > 0 || missingFieldsLeft16zvar > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zvar, missingFieldsLeft16zvar, msgp.ShowFound(found16zvar[:]), decodeMsgFieldOrder16zvar)
		if encodedFieldsLeft16zvar > 0 {
			encodedFieldsLeft16zvar--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zvar = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zvar < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zvar = 0
			}
			for nextMiss16zvar < maxFields16zvar && found16zvar[nextMiss16zvar] {
				nextMiss16zvar++
			}
			if nextMiss16zvar == maxFields16zvar {
				// filled all the empty fields!
				break doneWithStruct16zvar
			}
			missingFieldsLeft16zvar--
			curField16zvar = decodeMsgFieldOrder16zvar[nextMiss16zvar]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zvar)
		switch curField16zvar {
		// -- templateDecodeMsg ends here --

		case "Ty":
			found16zvar[0] = true
			{
				var zsvf int32
				zsvf, err = dc.ReadInt32()
				z.Ty = Ztype(zsvf)
			}
			if err != nil {
				panic(err)
			}
		case "Id":
			found16zvar[1] = true
			{
				var zazg int64
				zazg, err = dc.ReadInt64()
				z.Id = StructTypeId(zazg)
			}
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
	if nextMiss16zvar != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZPacket
var decodeMsgFieldOrder16zvar = []string{"Ty", "Id"}

// fieldsNotEmpty supports omitempty tags
func (z ZPacket) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z ZPacket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Ty"
	err = en.Append(0x82, 0xa2, 0x54, 0x79)
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
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ZPacket) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Ty"
	o = append(o, 0x82, 0xa2, 0x54, 0x79)
	o = msgp.AppendInt32(o, int32(z.Ty))
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt64(o, int64(z.Id))
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
	const maxFields17zaeo = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17zaeo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17zaeo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17zaeo := totalEncodedFields17zaeo
	missingFieldsLeft17zaeo := maxFields17zaeo - totalEncodedFields17zaeo

	var nextMiss17zaeo int32 = -1
	var found17zaeo [maxFields17zaeo]bool
	var curField17zaeo string

doneWithStruct17zaeo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17zaeo > 0 || missingFieldsLeft17zaeo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zaeo, missingFieldsLeft17zaeo, msgp.ShowFound(found17zaeo[:]), unmarshalMsgFieldOrder17zaeo)
		if encodedFieldsLeft17zaeo > 0 {
			encodedFieldsLeft17zaeo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17zaeo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17zaeo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17zaeo = 0
			}
			for nextMiss17zaeo < maxFields17zaeo && found17zaeo[nextMiss17zaeo] {
				nextMiss17zaeo++
			}
			if nextMiss17zaeo == maxFields17zaeo {
				// filled all the empty fields!
				break doneWithStruct17zaeo
			}
			missingFieldsLeft17zaeo--
			curField17zaeo = unmarshalMsgFieldOrder17zaeo[nextMiss17zaeo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17zaeo)
		switch curField17zaeo {
		// -- templateUnmarshalMsg ends here --

		case "Ty":
			found17zaeo[0] = true
			{
				var zghn int32
				zghn, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ty = Ztype(zghn)
			}
		case "Id":
			found17zaeo[1] = true
			{
				var zijd int64
				zijd, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zijd)
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss17zaeo != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZPacket
var unmarshalMsgFieldOrder17zaeo = []string{"Ty", "Id"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ZPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 3 + msgp.Int64Size
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
		var zulu int32
		zulu, err = dc.ReadInt32()
		(*z) = Ztype(zulu)
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
		var zkis int32
		zkis, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Ztype(zkis)
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
