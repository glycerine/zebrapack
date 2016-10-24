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
	const maxFields0ztxl = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0ztxl uint32
	totalEncodedFields0ztxl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0ztxl := totalEncodedFields0ztxl
	missingFieldsLeft0ztxl := maxFields0ztxl - totalEncodedFields0ztxl

	var nextMiss0ztxl int32 = -1
	var found0ztxl [maxFields0ztxl]bool
	var curField0ztxl string

doneWithStruct0ztxl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ztxl > 0 || missingFieldsLeft0ztxl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ztxl, missingFieldsLeft0ztxl, msgp.ShowFound(found0ztxl[:]), decodeMsgFieldOrder0ztxl)
		if encodedFieldsLeft0ztxl > 0 {
			encodedFieldsLeft0ztxl--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0ztxl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ztxl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0ztxl = 0
			}
			for nextMiss0ztxl < maxFields0ztxl && found0ztxl[nextMiss0ztxl] {
				nextMiss0ztxl++
			}
			if nextMiss0ztxl == maxFields0ztxl {
				// filled all the empty fields!
				break doneWithStruct0ztxl
			}
			missingFieldsLeft0ztxl--
			curField0ztxl = decodeMsgFieldOrder0ztxl[nextMiss0ztxl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ztxl)
		switch curField0ztxl {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0ztxl[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0ztxl[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found0ztxl[2] = true
			{
				var zgli int32
				zgli, err = dc.ReadInt32()
				z.Ztyp = Zkind(zgli)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0ztxl[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0ztxl[4] = true
			var zres uint32
			zres, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zres > 0 {
				z.Tag = make(map[string]string, zres)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zres > 0 {
				zres--
				var ztxg string
				var zaui string
				ztxg, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zaui, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[ztxg] = zaui
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0ztxl != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0ztxl = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

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
	var empty_zdzr [5]bool
	fieldsInUse_znzq := z.fieldsNotEmpty(empty_zdzr[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_znzq)
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
	if !empty_zdzr[3] {
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

	if !empty_zdzr[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for ztxg, zaui := range z.Tag {
			err = en.WriteString(ztxg)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zaui)
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
		for ztxg, zaui := range z.Tag {
			o = msgp.AppendString(o, ztxg)
			o = msgp.AppendString(o, zaui)
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
	const maxFields1zuue = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zuue uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zuue, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zuue := totalEncodedFields1zuue
	missingFieldsLeft1zuue := maxFields1zuue - totalEncodedFields1zuue

	var nextMiss1zuue int32 = -1
	var found1zuue [maxFields1zuue]bool
	var curField1zuue string

doneWithStruct1zuue:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zuue > 0 || missingFieldsLeft1zuue > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zuue, missingFieldsLeft1zuue, msgp.ShowFound(found1zuue[:]), unmarshalMsgFieldOrder1zuue)
		if encodedFieldsLeft1zuue > 0 {
			encodedFieldsLeft1zuue--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zuue = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zuue < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zuue = 0
			}
			for nextMiss1zuue < maxFields1zuue && found1zuue[nextMiss1zuue] {
				nextMiss1zuue++
			}
			if nextMiss1zuue == maxFields1zuue {
				// filled all the empty fields!
				break doneWithStruct1zuue
			}
			missingFieldsLeft1zuue--
			curField1zuue = unmarshalMsgFieldOrder1zuue[nextMiss1zuue]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zuue)
		switch curField1zuue {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zuue[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zuue[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Ztyp":
			found1zuue[2] = true
			{
				var zcbw int32
				zcbw, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ztyp = Zkind(zcbw)
			}
		case "Varg":
			found1zuue[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zuue[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zieq uint32
				zieq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zieq > 0 {
					z.Tag = make(map[string]string, zieq)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zieq > 0 {
					var ztxg string
					var zaui string
					zieq--
					ztxg, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zaui, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[ztxg] = zaui
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zuue != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zuue = []string{"FieldId", "Name", "Ztyp", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for ztxg, zaui := range z.Tag {
			_ = zaui
			_ = ztxg
			s += msgp.StringPrefixSize + len(ztxg) + msgp.StringPrefixSize + len(zaui)
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
	const maxFields2zqdi = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zqdi uint32
	totalEncodedFields2zqdi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zqdi := totalEncodedFields2zqdi
	missingFieldsLeft2zqdi := maxFields2zqdi - totalEncodedFields2zqdi

	var nextMiss2zqdi int32 = -1
	var found2zqdi [maxFields2zqdi]bool
	var curField2zqdi string

doneWithStruct2zqdi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zqdi > 0 || missingFieldsLeft2zqdi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zqdi, missingFieldsLeft2zqdi, msgp.ShowFound(found2zqdi[:]), decodeMsgFieldOrder2zqdi)
		if encodedFieldsLeft2zqdi > 0 {
			encodedFieldsLeft2zqdi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zqdi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zqdi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zqdi = 0
			}
			for nextMiss2zqdi < maxFields2zqdi && found2zqdi[nextMiss2zqdi] {
				nextMiss2zqdi++
			}
			if nextMiss2zqdi == maxFields2zqdi {
				// filled all the empty fields!
				break doneWithStruct2zqdi
			}
			missingFieldsLeft2zqdi--
			curField2zqdi = decodeMsgFieldOrder2zqdi[nextMiss2zqdi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zqdi)
		switch curField2zqdi {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zqdi[0] = true
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
	if nextMiss2zqdi != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zqdi = []string{"IfaceId"}

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
	const maxFields3zqnv = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zqnv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zqnv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zqnv := totalEncodedFields3zqnv
	missingFieldsLeft3zqnv := maxFields3zqnv - totalEncodedFields3zqnv

	var nextMiss3zqnv int32 = -1
	var found3zqnv [maxFields3zqnv]bool
	var curField3zqnv string

doneWithStruct3zqnv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zqnv > 0 || missingFieldsLeft3zqnv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zqnv, missingFieldsLeft3zqnv, msgp.ShowFound(found3zqnv[:]), unmarshalMsgFieldOrder3zqnv)
		if encodedFieldsLeft3zqnv > 0 {
			encodedFieldsLeft3zqnv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zqnv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zqnv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zqnv = 0
			}
			for nextMiss3zqnv < maxFields3zqnv && found3zqnv[nextMiss3zqnv] {
				nextMiss3zqnv++
			}
			if nextMiss3zqnv == maxFields3zqnv {
				// filled all the empty fields!
				break doneWithStruct3zqnv
			}
			missingFieldsLeft3zqnv--
			curField3zqnv = unmarshalMsgFieldOrder3zqnv[nextMiss3zqnv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zqnv)
		switch curField3zqnv {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3zqnv[0] = true
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
	if nextMiss3zqnv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3zqnv = []string{"IfaceId"}

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
	const maxFields4zrmh = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zrmh uint32
	totalEncodedFields4zrmh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zrmh := totalEncodedFields4zrmh
	missingFieldsLeft4zrmh := maxFields4zrmh - totalEncodedFields4zrmh

	var nextMiss4zrmh int32 = -1
	var found4zrmh [maxFields4zrmh]bool
	var curField4zrmh string

doneWithStruct4zrmh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zrmh > 0 || missingFieldsLeft4zrmh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zrmh, missingFieldsLeft4zrmh, msgp.ShowFound(found4zrmh[:]), decodeMsgFieldOrder4zrmh)
		if encodedFieldsLeft4zrmh > 0 {
			encodedFieldsLeft4zrmh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zrmh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zrmh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zrmh = 0
			}
			for nextMiss4zrmh < maxFields4zrmh && found4zrmh[nextMiss4zrmh] {
				nextMiss4zrmh++
			}
			if nextMiss4zrmh == maxFields4zrmh {
				// filled all the empty fields!
				break doneWithStruct4zrmh
			}
			missingFieldsLeft4zrmh--
			curField4zrmh = decodeMsgFieldOrder4zrmh[nextMiss4zrmh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zrmh)
		switch curField4zrmh {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4zrmh[0] = true
			var zzxj uint32
			zzxj, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zzxj) {
				z.Methods = (z.Methods)[:zzxj]
			} else {
				z.Methods = make([]MethodT, zzxj)
			}
			for zhtl := range z.Methods {
				err = z.Methods[zhtl].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4zrmh[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zrmh[2] = true
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
	if nextMiss4zrmh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4zrmh = []string{"Methods", "Deprecated", "Comment"}

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
	for zhtl := range z.Methods {
		err = z.Methods[zhtl].EncodeMsg(en)
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
	for zhtl := range z.Methods {
		o, err = z.Methods[zhtl].MarshalMsg(o)
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
	const maxFields5zpmx = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zpmx uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zpmx, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zpmx := totalEncodedFields5zpmx
	missingFieldsLeft5zpmx := maxFields5zpmx - totalEncodedFields5zpmx

	var nextMiss5zpmx int32 = -1
	var found5zpmx [maxFields5zpmx]bool
	var curField5zpmx string

doneWithStruct5zpmx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zpmx > 0 || missingFieldsLeft5zpmx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zpmx, missingFieldsLeft5zpmx, msgp.ShowFound(found5zpmx[:]), unmarshalMsgFieldOrder5zpmx)
		if encodedFieldsLeft5zpmx > 0 {
			encodedFieldsLeft5zpmx--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zpmx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zpmx < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zpmx = 0
			}
			for nextMiss5zpmx < maxFields5zpmx && found5zpmx[nextMiss5zpmx] {
				nextMiss5zpmx++
			}
			if nextMiss5zpmx == maxFields5zpmx {
				// filled all the empty fields!
				break doneWithStruct5zpmx
			}
			missingFieldsLeft5zpmx--
			curField5zpmx = unmarshalMsgFieldOrder5zpmx[nextMiss5zpmx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zpmx)
		switch curField5zpmx {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zpmx[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zsao uint32
				zsao, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zsao) {
					z.Methods = (z.Methods)[:zsao]
				} else {
					z.Methods = make([]MethodT, zsao)
				}
				for zhtl := range z.Methods {
					bts, err = z.Methods[zhtl].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zpmx[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zpmx[2] = true
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
	if nextMiss5zpmx != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zpmx = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zhtl := range z.Methods {
		s += z.Methods[zhtl].Msgsize()
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
	const maxFields6zhrb = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zhrb uint32
	totalEncodedFields6zhrb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zhrb := totalEncodedFields6zhrb
	missingFieldsLeft6zhrb := maxFields6zhrb - totalEncodedFields6zhrb

	var nextMiss6zhrb int32 = -1
	var found6zhrb [maxFields6zhrb]bool
	var curField6zhrb string

doneWithStruct6zhrb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zhrb > 0 || missingFieldsLeft6zhrb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zhrb, missingFieldsLeft6zhrb, msgp.ShowFound(found6zhrb[:]), decodeMsgFieldOrder6zhrb)
		if encodedFieldsLeft6zhrb > 0 {
			encodedFieldsLeft6zhrb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zhrb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zhrb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zhrb = 0
			}
			for nextMiss6zhrb < maxFields6zhrb && found6zhrb[nextMiss6zhrb] {
				nextMiss6zhrb++
			}
			if nextMiss6zhrb == maxFields6zhrb {
				// filled all the empty fields!
				break doneWithStruct6zhrb
			}
			missingFieldsLeft6zhrb--
			curField6zhrb = decodeMsgFieldOrder6zhrb[nextMiss6zhrb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zhrb)
		switch curField6zhrb {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zhrb[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zhrb[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zhrb[2] = true
			const maxFields7zneo = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7zneo uint32
			totalEncodedFields7zneo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7zneo := totalEncodedFields7zneo
			missingFieldsLeft7zneo := maxFields7zneo - totalEncodedFields7zneo

			var nextMiss7zneo int32 = -1
			var found7zneo [maxFields7zneo]bool
			var curField7zneo string

		doneWithStruct7zneo:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7zneo > 0 || missingFieldsLeft7zneo > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zneo, missingFieldsLeft7zneo, msgp.ShowFound(found7zneo[:]), decodeMsgFieldOrder7zneo)
				if encodedFieldsLeft7zneo > 0 {
					encodedFieldsLeft7zneo--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7zneo = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7zneo < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7zneo = 0
					}
					for nextMiss7zneo < maxFields7zneo && found7zneo[nextMiss7zneo] {
						nextMiss7zneo++
					}
					if nextMiss7zneo == maxFields7zneo {
						// filled all the empty fields!
						break doneWithStruct7zneo
					}
					missingFieldsLeft7zneo--
					curField7zneo = decodeMsgFieldOrder7zneo[nextMiss7zneo]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7zneo)
				switch curField7zneo {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7zneo[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7zneo[1] = true
					var zfcb uint32
					zfcb, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zfcb) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zfcb]
					} else {
						z.Inputs.Fld = make([]FieldT, zfcb)
					}
					for zeky := range z.Inputs.Fld {
						err = z.Inputs.Fld[zeky].DecodeMsg(dc)
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
			if nextMiss7zneo != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zhrb[3] = true
			const maxFields8zorc = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zorc uint32
			totalEncodedFields8zorc, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zorc := totalEncodedFields8zorc
			missingFieldsLeft8zorc := maxFields8zorc - totalEncodedFields8zorc

			var nextMiss8zorc int32 = -1
			var found8zorc [maxFields8zorc]bool
			var curField8zorc string

		doneWithStruct8zorc:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zorc > 0 || missingFieldsLeft8zorc > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zorc, missingFieldsLeft8zorc, msgp.ShowFound(found8zorc[:]), decodeMsgFieldOrder8zorc)
				if encodedFieldsLeft8zorc > 0 {
					encodedFieldsLeft8zorc--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zorc = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zorc < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zorc = 0
					}
					for nextMiss8zorc < maxFields8zorc && found8zorc[nextMiss8zorc] {
						nextMiss8zorc++
					}
					if nextMiss8zorc == maxFields8zorc {
						// filled all the empty fields!
						break doneWithStruct8zorc
					}
					missingFieldsLeft8zorc--
					curField8zorc = decodeMsgFieldOrder8zorc[nextMiss8zorc]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zorc)
				switch curField8zorc {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zorc[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zorc[1] = true
					var zwaf uint32
					zwaf, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zwaf) {
						z.Returns.Fld = (z.Returns.Fld)[:zwaf]
					} else {
						z.Returns.Fld = make([]FieldT, zwaf)
					}
					for zoby := range z.Returns.Fld {
						err = z.Returns.Fld[zoby].DecodeMsg(dc)
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
			if nextMiss8zorc != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zhrb[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zhrb[5] = true
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
	if nextMiss6zhrb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zhrb = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7zneo = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zorc = []string{"Nam", "Fld"}

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
	for zeky := range z.Inputs.Fld {
		err = z.Inputs.Fld[zeky].EncodeMsg(en)
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
	for zoby := range z.Returns.Fld {
		err = z.Returns.Fld[zoby].EncodeMsg(en)
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
	for zeky := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zeky].MarshalMsg(o)
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
	for zoby := range z.Returns.Fld {
		o, err = z.Returns.Fld[zoby].MarshalMsg(o)
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
	const maxFields9zfen = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zfen uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zfen, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zfen := totalEncodedFields9zfen
	missingFieldsLeft9zfen := maxFields9zfen - totalEncodedFields9zfen

	var nextMiss9zfen int32 = -1
	var found9zfen [maxFields9zfen]bool
	var curField9zfen string

doneWithStruct9zfen:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zfen > 0 || missingFieldsLeft9zfen > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zfen, missingFieldsLeft9zfen, msgp.ShowFound(found9zfen[:]), unmarshalMsgFieldOrder9zfen)
		if encodedFieldsLeft9zfen > 0 {
			encodedFieldsLeft9zfen--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zfen = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zfen < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zfen = 0
			}
			for nextMiss9zfen < maxFields9zfen && found9zfen[nextMiss9zfen] {
				nextMiss9zfen++
			}
			if nextMiss9zfen == maxFields9zfen {
				// filled all the empty fields!
				break doneWithStruct9zfen
			}
			missingFieldsLeft9zfen--
			curField9zfen = unmarshalMsgFieldOrder9zfen[nextMiss9zfen]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zfen)
		switch curField9zfen {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zfen[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zfen[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zfen[2] = true
			const maxFields10zsld = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zsld uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zsld, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10zsld := totalEncodedFields10zsld
			missingFieldsLeft10zsld := maxFields10zsld - totalEncodedFields10zsld

			var nextMiss10zsld int32 = -1
			var found10zsld [maxFields10zsld]bool
			var curField10zsld string

		doneWithStruct10zsld:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10zsld > 0 || missingFieldsLeft10zsld > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zsld, missingFieldsLeft10zsld, msgp.ShowFound(found10zsld[:]), unmarshalMsgFieldOrder10zsld)
				if encodedFieldsLeft10zsld > 0 {
					encodedFieldsLeft10zsld--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10zsld = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zsld < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zsld = 0
					}
					for nextMiss10zsld < maxFields10zsld && found10zsld[nextMiss10zsld] {
						nextMiss10zsld++
					}
					if nextMiss10zsld == maxFields10zsld {
						// filled all the empty fields!
						break doneWithStruct10zsld
					}
					missingFieldsLeft10zsld--
					curField10zsld = unmarshalMsgFieldOrder10zsld[nextMiss10zsld]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zsld)
				switch curField10zsld {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10zsld[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10zsld[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var ztqq uint32
						ztqq, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(ztqq) {
							z.Inputs.Fld = (z.Inputs.Fld)[:ztqq]
						} else {
							z.Inputs.Fld = make([]FieldT, ztqq)
						}
						for zeky := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zeky].UnmarshalMsg(bts)
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
			if nextMiss10zsld != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zfen[3] = true
			const maxFields11zxup = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zxup uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zxup, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zxup := totalEncodedFields11zxup
			missingFieldsLeft11zxup := maxFields11zxup - totalEncodedFields11zxup

			var nextMiss11zxup int32 = -1
			var found11zxup [maxFields11zxup]bool
			var curField11zxup string

		doneWithStruct11zxup:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zxup > 0 || missingFieldsLeft11zxup > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zxup, missingFieldsLeft11zxup, msgp.ShowFound(found11zxup[:]), unmarshalMsgFieldOrder11zxup)
				if encodedFieldsLeft11zxup > 0 {
					encodedFieldsLeft11zxup--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zxup = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zxup < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zxup = 0
					}
					for nextMiss11zxup < maxFields11zxup && found11zxup[nextMiss11zxup] {
						nextMiss11zxup++
					}
					if nextMiss11zxup == maxFields11zxup {
						// filled all the empty fields!
						break doneWithStruct11zxup
					}
					missingFieldsLeft11zxup--
					curField11zxup = unmarshalMsgFieldOrder11zxup[nextMiss11zxup]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zxup)
				switch curField11zxup {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zxup[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zxup[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zuen uint32
						zuen, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zuen) {
							z.Returns.Fld = (z.Returns.Fld)[:zuen]
						} else {
							z.Returns.Fld = make([]FieldT, zuen)
						}
						for zoby := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zoby].UnmarshalMsg(bts)
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
			if nextMiss11zxup != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zfen[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zfen[5] = true
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
	if nextMiss9zfen != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zfen = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10zsld = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zxup = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zeky := range z.Inputs.Fld {
		s += z.Inputs.Fld[zeky].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for zoby := range z.Returns.Fld {
		s += z.Returns.Fld[zoby].Msgsize()
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
	const maxFields12zeqf = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zeqf uint32
	totalEncodedFields12zeqf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zeqf := totalEncodedFields12zeqf
	missingFieldsLeft12zeqf := maxFields12zeqf - totalEncodedFields12zeqf

	var nextMiss12zeqf int32 = -1
	var found12zeqf [maxFields12zeqf]bool
	var curField12zeqf string

doneWithStruct12zeqf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zeqf > 0 || missingFieldsLeft12zeqf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zeqf, missingFieldsLeft12zeqf, msgp.ShowFound(found12zeqf[:]), decodeMsgFieldOrder12zeqf)
		if encodedFieldsLeft12zeqf > 0 {
			encodedFieldsLeft12zeqf--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zeqf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zeqf < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zeqf = 0
			}
			for nextMiss12zeqf < maxFields12zeqf && found12zeqf[nextMiss12zeqf] {
				nextMiss12zeqf++
			}
			if nextMiss12zeqf == maxFields12zeqf {
				// filled all the empty fields!
				break doneWithStruct12zeqf
			}
			missingFieldsLeft12zeqf--
			curField12zeqf = decodeMsgFieldOrder12zeqf[nextMiss12zeqf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zeqf)
		switch curField12zeqf {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zeqf[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found12zeqf[1] = true
			var zxzs uint32
			zxzs, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zxzs > 0 {
				z.Structs = make(map[int64]StructT, zxzs)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zxzs > 0 {
				zxzs--
				var zbvm int64
				var zkfs StructT
				zbvm, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				const maxFields13zolf = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields13zolf uint32
				totalEncodedFields13zolf, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft13zolf := totalEncodedFields13zolf
				missingFieldsLeft13zolf := maxFields13zolf - totalEncodedFields13zolf

				var nextMiss13zolf int32 = -1
				var found13zolf [maxFields13zolf]bool
				var curField13zolf string

			doneWithStruct13zolf:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft13zolf > 0 || missingFieldsLeft13zolf > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zolf, missingFieldsLeft13zolf, msgp.ShowFound(found13zolf[:]), decodeMsgFieldOrder13zolf)
					if encodedFieldsLeft13zolf > 0 {
						encodedFieldsLeft13zolf--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField13zolf = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss13zolf < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss13zolf = 0
						}
						for nextMiss13zolf < maxFields13zolf && found13zolf[nextMiss13zolf] {
							nextMiss13zolf++
						}
						if nextMiss13zolf == maxFields13zolf {
							// filled all the empty fields!
							break doneWithStruct13zolf
						}
						missingFieldsLeft13zolf--
						curField13zolf = decodeMsgFieldOrder13zolf[nextMiss13zolf]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField13zolf)
					switch curField13zolf {
					// -- templateDecodeMsg ends here --

					case "Nam":
						found13zolf[0] = true
						zkfs.Nam, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found13zolf[1] = true
						var zcnz uint32
						zcnz, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(zkfs.Fld) >= int(zcnz) {
							zkfs.Fld = (zkfs.Fld)[:zcnz]
						} else {
							zkfs.Fld = make([]FieldT, zcnz)
						}
						for zwjq := range zkfs.Fld {
							err = zkfs.Fld[zwjq].DecodeMsg(dc)
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
				if nextMiss13zolf != -1 {
					dc.PopAlwaysNil()
				}

				z.Structs[zbvm] = zkfs
			}
		case "Interfaces":
			found12zeqf[2] = true
			var zgle uint32
			zgle, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Interfaces == nil && zgle > 0 {
				z.Interfaces = make(map[int64]InterfaceT, zgle)
			} else if len(z.Interfaces) > 0 {
				for key, _ := range z.Interfaces {
					delete(z.Interfaces, key)
				}
			}
			for zgle > 0 {
				zgle--
				var znbu int64
				var znap InterfaceT
				znbu, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = znap.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Interfaces[znbu] = znap
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zeqf != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zeqf = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var decodeMsgFieldOrder13zolf = []string{"Nam", "Fld"}

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
	for zbvm, zkfs := range z.Structs {
		err = en.WriteInt64(zbvm)
		if err != nil {
			panic(err)
		}
		// map header, size 2
		// write "Nam"
		err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteString(zkfs.Nam)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(zkfs.Fld)))
		if err != nil {
			panic(err)
		}
		for zwjq := range zkfs.Fld {
			err = zkfs.Fld[zwjq].EncodeMsg(en)
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
	for znbu, znap := range z.Interfaces {
		err = en.WriteInt64(znbu)
		if err != nil {
			panic(err)
		}
		err = znap.EncodeMsg(en)
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
	for zbvm, zkfs := range z.Structs {
		o = msgp.AppendInt64(o, zbvm)
		// map header, size 2
		// string "Nam"
		o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
		o = msgp.AppendString(o, zkfs.Nam)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(zkfs.Fld)))
		for zwjq := range zkfs.Fld {
			o, err = zkfs.Fld[zwjq].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Interfaces"
	o = append(o, 0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Interfaces)))
	for znbu, znap := range z.Interfaces {
		o = msgp.AppendInt64(o, znbu)
		o, err = znap.MarshalMsg(o)
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
	const maxFields14znuh = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields14znuh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields14znuh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft14znuh := totalEncodedFields14znuh
	missingFieldsLeft14znuh := maxFields14znuh - totalEncodedFields14znuh

	var nextMiss14znuh int32 = -1
	var found14znuh [maxFields14znuh]bool
	var curField14znuh string

doneWithStruct14znuh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14znuh > 0 || missingFieldsLeft14znuh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14znuh, missingFieldsLeft14znuh, msgp.ShowFound(found14znuh[:]), unmarshalMsgFieldOrder14znuh)
		if encodedFieldsLeft14znuh > 0 {
			encodedFieldsLeft14znuh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField14znuh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14znuh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss14znuh = 0
			}
			for nextMiss14znuh < maxFields14znuh && found14znuh[nextMiss14znuh] {
				nextMiss14znuh++
			}
			if nextMiss14znuh == maxFields14znuh {
				// filled all the empty fields!
				break doneWithStruct14znuh
			}
			missingFieldsLeft14znuh--
			curField14znuh = unmarshalMsgFieldOrder14znuh[nextMiss14znuh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14znuh)
		switch curField14znuh {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found14znuh[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found14znuh[1] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zvqn uint32
				zvqn, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zvqn > 0 {
					z.Structs = make(map[int64]StructT, zvqn)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zvqn > 0 {
					var zbvm int64
					var zkfs StructT
					zvqn--
					zbvm, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					const maxFields15zbrj = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields15zbrj uint32
					if !nbs.AlwaysNil {
						totalEncodedFields15zbrj, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft15zbrj := totalEncodedFields15zbrj
					missingFieldsLeft15zbrj := maxFields15zbrj - totalEncodedFields15zbrj

					var nextMiss15zbrj int32 = -1
					var found15zbrj [maxFields15zbrj]bool
					var curField15zbrj string

				doneWithStruct15zbrj:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft15zbrj > 0 || missingFieldsLeft15zbrj > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zbrj, missingFieldsLeft15zbrj, msgp.ShowFound(found15zbrj[:]), unmarshalMsgFieldOrder15zbrj)
						if encodedFieldsLeft15zbrj > 0 {
							encodedFieldsLeft15zbrj--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField15zbrj = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss15zbrj < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss15zbrj = 0
							}
							for nextMiss15zbrj < maxFields15zbrj && found15zbrj[nextMiss15zbrj] {
								nextMiss15zbrj++
							}
							if nextMiss15zbrj == maxFields15zbrj {
								// filled all the empty fields!
								break doneWithStruct15zbrj
							}
							missingFieldsLeft15zbrj--
							curField15zbrj = unmarshalMsgFieldOrder15zbrj[nextMiss15zbrj]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField15zbrj)
						switch curField15zbrj {
						// -- templateUnmarshalMsg ends here --

						case "Nam":
							found15zbrj[0] = true
							zkfs.Nam, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found15zbrj[1] = true
							if nbs.AlwaysNil {
								(zkfs.Fld) = (zkfs.Fld)[:0]
							} else {

								var zbgf uint32
								zbgf, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(zkfs.Fld) >= int(zbgf) {
									zkfs.Fld = (zkfs.Fld)[:zbgf]
								} else {
									zkfs.Fld = make([]FieldT, zbgf)
								}
								for zwjq := range zkfs.Fld {
									bts, err = zkfs.Fld[zwjq].UnmarshalMsg(bts)
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
					if nextMiss15zbrj != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Structs[zbvm] = zkfs
				}
			}
		case "Interfaces":
			found14znuh[2] = true
			if nbs.AlwaysNil {
				if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}

			} else {

				var zlqj uint32
				zlqj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Interfaces == nil && zlqj > 0 {
					z.Interfaces = make(map[int64]InterfaceT, zlqj)
				} else if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}
				for zlqj > 0 {
					var znbu int64
					var znap InterfaceT
					zlqj--
					znbu, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = znap.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Interfaces[znbu] = znap
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14znuh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder14znuh = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var unmarshalMsgFieldOrder15zbrj = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zbvm, zkfs := range z.Structs {
			_ = zkfs
			_ = zbvm
			s += msgp.Int64Size + 1 + 4 + msgp.StringPrefixSize + len(zkfs.Nam) + 4 + msgp.ArrayHeaderSize
			for zwjq := range zkfs.Fld {
				s += zkfs.Fld[zwjq].Msgsize()
			}
		}
	}
	s += 11 + msgp.MapHeaderSize
	if z.Interfaces != nil {
		for znbu, znap := range z.Interfaces {
			_ = znap
			_ = znbu
			s += msgp.Int64Size + znap.Msgsize()
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
	const maxFields16zsrp = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zsrp uint32
	totalEncodedFields16zsrp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zsrp := totalEncodedFields16zsrp
	missingFieldsLeft16zsrp := maxFields16zsrp - totalEncodedFields16zsrp

	var nextMiss16zsrp int32 = -1
	var found16zsrp [maxFields16zsrp]bool
	var curField16zsrp string

doneWithStruct16zsrp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zsrp > 0 || missingFieldsLeft16zsrp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zsrp, missingFieldsLeft16zsrp, msgp.ShowFound(found16zsrp[:]), decodeMsgFieldOrder16zsrp)
		if encodedFieldsLeft16zsrp > 0 {
			encodedFieldsLeft16zsrp--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zsrp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zsrp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zsrp = 0
			}
			for nextMiss16zsrp < maxFields16zsrp && found16zsrp[nextMiss16zsrp] {
				nextMiss16zsrp++
			}
			if nextMiss16zsrp == maxFields16zsrp {
				// filled all the empty fields!
				break doneWithStruct16zsrp
			}
			missingFieldsLeft16zsrp--
			curField16zsrp = decodeMsgFieldOrder16zsrp[nextMiss16zsrp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zsrp)
		switch curField16zsrp {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found16zsrp[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found16zsrp[1] = true
			var zcue uint32
			zcue, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zcue) {
				z.Fld = (z.Fld)[:zcue]
			} else {
				z.Fld = make([]FieldT, zcue)
			}
			for zgun := range z.Fld {
				err = z.Fld[zgun].DecodeMsg(dc)
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
	if nextMiss16zsrp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder16zsrp = []string{"Nam", "Fld"}

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
	for zgun := range z.Fld {
		err = z.Fld[zgun].EncodeMsg(en)
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
	for zgun := range z.Fld {
		o, err = z.Fld[zgun].MarshalMsg(o)
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
	const maxFields17zjxl = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17zjxl uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17zjxl, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17zjxl := totalEncodedFields17zjxl
	missingFieldsLeft17zjxl := maxFields17zjxl - totalEncodedFields17zjxl

	var nextMiss17zjxl int32 = -1
	var found17zjxl [maxFields17zjxl]bool
	var curField17zjxl string

doneWithStruct17zjxl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17zjxl > 0 || missingFieldsLeft17zjxl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zjxl, missingFieldsLeft17zjxl, msgp.ShowFound(found17zjxl[:]), unmarshalMsgFieldOrder17zjxl)
		if encodedFieldsLeft17zjxl > 0 {
			encodedFieldsLeft17zjxl--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17zjxl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17zjxl < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17zjxl = 0
			}
			for nextMiss17zjxl < maxFields17zjxl && found17zjxl[nextMiss17zjxl] {
				nextMiss17zjxl++
			}
			if nextMiss17zjxl == maxFields17zjxl {
				// filled all the empty fields!
				break doneWithStruct17zjxl
			}
			missingFieldsLeft17zjxl--
			curField17zjxl = unmarshalMsgFieldOrder17zjxl[nextMiss17zjxl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17zjxl)
		switch curField17zjxl {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found17zjxl[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found17zjxl[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zdhz uint32
				zdhz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zdhz) {
					z.Fld = (z.Fld)[:zdhz]
				} else {
					z.Fld = make([]FieldT, zdhz)
				}
				for zgun := range z.Fld {
					bts, err = z.Fld[zgun].UnmarshalMsg(bts)
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
	if nextMiss17zjxl != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder17zjxl = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zgun := range z.Fld {
		s += z.Fld[zgun].Msgsize()
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
		var zuan int64
		zuan, err = dc.ReadInt64()
		(*z) = StructTypeId(zuan)
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
		var zxkz int64
		zxkz, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zxkz)
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
	const maxFields18zlrj = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields18zlrj uint32
	totalEncodedFields18zlrj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft18zlrj := totalEncodedFields18zlrj
	missingFieldsLeft18zlrj := maxFields18zlrj - totalEncodedFields18zlrj

	var nextMiss18zlrj int32 = -1
	var found18zlrj [maxFields18zlrj]bool
	var curField18zlrj string

doneWithStruct18zlrj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft18zlrj > 0 || missingFieldsLeft18zlrj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zlrj, missingFieldsLeft18zlrj, msgp.ShowFound(found18zlrj[:]), decodeMsgFieldOrder18zlrj)
		if encodedFieldsLeft18zlrj > 0 {
			encodedFieldsLeft18zlrj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField18zlrj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss18zlrj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss18zlrj = 0
			}
			for nextMiss18zlrj < maxFields18zlrj && found18zlrj[nextMiss18zlrj] {
				nextMiss18zlrj++
			}
			if nextMiss18zlrj == maxFields18zlrj {
				// filled all the empty fields!
				break doneWithStruct18zlrj
			}
			missingFieldsLeft18zlrj--
			curField18zlrj = decodeMsgFieldOrder18zlrj[nextMiss18zlrj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField18zlrj)
		switch curField18zlrj {
		// -- templateDecodeMsg ends here --

		case "Ty":
			found18zlrj[0] = true
			{
				var zckj int32
				zckj, err = dc.ReadInt32()
				z.Ty = Zkind(zckj)
			}
			if err != nil {
				panic(err)
			}
		case "Id":
			found18zlrj[1] = true
			{
				var zizk int64
				zizk, err = dc.ReadInt64()
				z.Id = StructTypeId(zizk)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found18zlrj[2] = true
			var zmfq uint32
			zmfq, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zmfq > 0 {
				z.Da = make(map[int64]msgp.Raw, zmfq)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zmfq > 0 {
				zmfq--
				var zewp int64
				var zdgd msgp.Raw
				zewp, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zdgd.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zewp] = zdgd
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss18zlrj != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZPacket
var decodeMsgFieldOrder18zlrj = []string{"Ty", "Id", "Da"}

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
	for zewp, zdgd := range z.Da {
		err = en.WriteInt64(zewp)
		if err != nil {
			panic(err)
		}
		err = zdgd.EncodeMsg(en)
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
	for zewp, zdgd := range z.Da {
		o = msgp.AppendInt64(o, zewp)
		o, err = zdgd.MarshalMsg(o)
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
	const maxFields19ztob = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19ztob uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19ztob, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19ztob := totalEncodedFields19ztob
	missingFieldsLeft19ztob := maxFields19ztob - totalEncodedFields19ztob

	var nextMiss19ztob int32 = -1
	var found19ztob [maxFields19ztob]bool
	var curField19ztob string

doneWithStruct19ztob:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19ztob > 0 || missingFieldsLeft19ztob > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19ztob, missingFieldsLeft19ztob, msgp.ShowFound(found19ztob[:]), unmarshalMsgFieldOrder19ztob)
		if encodedFieldsLeft19ztob > 0 {
			encodedFieldsLeft19ztob--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19ztob = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19ztob < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19ztob = 0
			}
			for nextMiss19ztob < maxFields19ztob && found19ztob[nextMiss19ztob] {
				nextMiss19ztob++
			}
			if nextMiss19ztob == maxFields19ztob {
				// filled all the empty fields!
				break doneWithStruct19ztob
			}
			missingFieldsLeft19ztob--
			curField19ztob = unmarshalMsgFieldOrder19ztob[nextMiss19ztob]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19ztob)
		switch curField19ztob {
		// -- templateUnmarshalMsg ends here --

		case "Ty":
			found19ztob[0] = true
			{
				var zpir int32
				zpir, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Ty = Zkind(zpir)
			}
		case "Id":
			found19ztob[1] = true
			{
				var ztqi int64
				ztqi, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(ztqi)
			}
		case "Da":
			found19ztob[2] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zptb uint32
				zptb, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zptb > 0 {
					z.Da = make(map[int64]msgp.Raw, zptb)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zptb > 0 {
					var zewp int64
					var zdgd msgp.Raw
					zptb--
					zewp, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zdgd.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zewp] = zdgd
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19ztob != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZPacket
var unmarshalMsgFieldOrder19ztob = []string{"Ty", "Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int32Size + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zewp, zdgd := range z.Da {
			_ = zdgd
			_ = zewp
			s += msgp.Int64Size + zdgd.Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Zkind) DecodeMsg(dc *msgp.Reader) (err error) {
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
		var zjdx int32
		zjdx, err = dc.ReadInt32()
		(*z) = Zkind(zjdx)
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
func (z Zkind) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt32(int32(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zkind) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt32(o, int32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zkind) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Zkind) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zsou int32
		zsou, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zsou)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zkind) Msgsize() (s int) {
	s = msgp.Int32Size
	return
}
