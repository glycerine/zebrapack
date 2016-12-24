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
	const maxFields0zwai = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zwai uint32
	totalEncodedFields0zwai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zwai := totalEncodedFields0zwai
	missingFieldsLeft0zwai := maxFields0zwai - totalEncodedFields0zwai

	var nextMiss0zwai int32 = -1
	var found0zwai [maxFields0zwai]bool
	var curField0zwai string

doneWithStruct0zwai:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zwai > 0 || missingFieldsLeft0zwai > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zwai, missingFieldsLeft0zwai, msgp.ShowFound(found0zwai[:]), decodeMsgFieldOrder0zwai)
		if encodedFieldsLeft0zwai > 0 {
			encodedFieldsLeft0zwai--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zwai = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zwai < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zwai = 0
			}
			for nextMiss0zwai < maxFields0zwai && found0zwai[nextMiss0zwai] {
				nextMiss0zwai++
			}
			if nextMiss0zwai == maxFields0zwai {
				// filled all the empty fields!
				break doneWithStruct0zwai
			}
			missingFieldsLeft0zwai--
			curField0zwai = decodeMsgFieldOrder0zwai[nextMiss0zwai]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zwai)
		switch curField0zwai {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zwai[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zwai[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Zknd":
			found0zwai[2] = true
			var zppc uint32
			zppc, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Zknd) >= int(zppc) {
				z.Zknd = (z.Zknd)[:zppc]
			} else {
				z.Zknd = make([]ZKind, zppc)
			}
			for zfoz := range z.Zknd {
				{
					var zszb byte
					zszb, err = dc.ReadByte()
					z.Zknd[zfoz] = ZKind(zszb)
				}
				if err != nil {
					panic(err)
				}
			}
		case "Varg":
			found0zwai[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zwai[4] = true
			var ziar uint32
			ziar, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && ziar > 0 {
				z.Tag = make(map[string]string, ziar)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for ziar > 0 {
				ziar--
				var zaoe string
				var zulw string
				zaoe, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zulw, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zaoe] = zulw
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zwai != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zwai = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

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
	var empty_zczh [5]bool
	fieldsInUse_zqdl := z.fieldsNotEmpty(empty_zczh[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqdl)
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
	// write "Zknd"
	err = en.Append(0xa4, 0x5a, 0x6b, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Zknd)))
	if err != nil {
		panic(err)
	}
	for zfoz := range z.Zknd {
		err = en.WriteByte(byte(z.Zknd[zfoz]))
		if err != nil {
			panic(err)
		}
	}
	if !empty_zczh[3] {
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

	if !empty_zczh[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zaoe, zulw := range z.Tag {
			err = en.WriteString(zaoe)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zulw)
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
	// string "Zknd"
	o = append(o, 0xa4, 0x5a, 0x6b, 0x6e, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Zknd)))
	for zfoz := range z.Zknd {
		o = msgp.AppendByte(o, byte(z.Zknd[zfoz]))
	}
	if !empty[3] {
		// string "Varg"
		o = append(o, 0xa4, 0x56, 0x61, 0x72, 0x67)
		o = msgp.AppendBool(o, z.Varg)
	}

	if !empty[4] {
		// string "Tag"
		o = append(o, 0xa3, 0x54, 0x61, 0x67)
		o = msgp.AppendMapHeader(o, uint32(len(z.Tag)))
		for zaoe, zulw := range z.Tag {
			o = msgp.AppendString(o, zaoe)
			o = msgp.AppendString(o, zulw)
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
	const maxFields1ztuh = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1ztuh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1ztuh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1ztuh := totalEncodedFields1ztuh
	missingFieldsLeft1ztuh := maxFields1ztuh - totalEncodedFields1ztuh

	var nextMiss1ztuh int32 = -1
	var found1ztuh [maxFields1ztuh]bool
	var curField1ztuh string

doneWithStruct1ztuh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1ztuh > 0 || missingFieldsLeft1ztuh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1ztuh, missingFieldsLeft1ztuh, msgp.ShowFound(found1ztuh[:]), unmarshalMsgFieldOrder1ztuh)
		if encodedFieldsLeft1ztuh > 0 {
			encodedFieldsLeft1ztuh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1ztuh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1ztuh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1ztuh = 0
			}
			for nextMiss1ztuh < maxFields1ztuh && found1ztuh[nextMiss1ztuh] {
				nextMiss1ztuh++
			}
			if nextMiss1ztuh == maxFields1ztuh {
				// filled all the empty fields!
				break doneWithStruct1ztuh
			}
			missingFieldsLeft1ztuh--
			curField1ztuh = unmarshalMsgFieldOrder1ztuh[nextMiss1ztuh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1ztuh)
		switch curField1ztuh {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1ztuh[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1ztuh[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Zknd":
			found1ztuh[2] = true
			if nbs.AlwaysNil {
				(z.Zknd) = (z.Zknd)[:0]
			} else {

				var ztwk uint32
				ztwk, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Zknd) >= int(ztwk) {
					z.Zknd = (z.Zknd)[:ztwk]
				} else {
					z.Zknd = make([]ZKind, ztwk)
				}
				for zfoz := range z.Zknd {
					{
						var zhuf byte
						zhuf, bts, err = nbs.ReadByteBytes(bts)

						if err != nil {
							panic(err)
						}
						z.Zknd[zfoz] = ZKind(zhuf)
					}
				}
			}
		case "Varg":
			found1ztuh[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1ztuh[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zejb uint32
				zejb, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zejb > 0 {
					z.Tag = make(map[string]string, zejb)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zejb > 0 {
					var zaoe string
					var zulw string
					zejb--
					zaoe, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zulw, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zaoe] = zulw
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1ztuh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1ztuh = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize + (len(z.Zknd) * (msgp.ByteSize)) + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zaoe, zulw := range z.Tag {
			_ = zulw
			_ = zaoe
			s += msgp.StringPrefixSize + len(zaoe) + msgp.StringPrefixSize + len(zulw)
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *InterfaceId) DecodeMsg(dc *msgp.Reader) (err error) {
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
		var zccc int64
		zccc, err = dc.ReadInt64()
		(*z) = InterfaceId(zccc)
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
func (z InterfaceId) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt64(int64(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z InterfaceId) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt64(o, int64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InterfaceId) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *InterfaceId) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zxqp int64
		zxqp, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = InterfaceId(zxqp)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z InterfaceId) Msgsize() (s int) {
	s = msgp.Int64Size
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
	const maxFields2zwlz = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zwlz uint32
	totalEncodedFields2zwlz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zwlz := totalEncodedFields2zwlz
	missingFieldsLeft2zwlz := maxFields2zwlz - totalEncodedFields2zwlz

	var nextMiss2zwlz int32 = -1
	var found2zwlz [maxFields2zwlz]bool
	var curField2zwlz string

doneWithStruct2zwlz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zwlz > 0 || missingFieldsLeft2zwlz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zwlz, missingFieldsLeft2zwlz, msgp.ShowFound(found2zwlz[:]), decodeMsgFieldOrder2zwlz)
		if encodedFieldsLeft2zwlz > 0 {
			encodedFieldsLeft2zwlz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zwlz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zwlz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zwlz = 0
			}
			for nextMiss2zwlz < maxFields2zwlz && found2zwlz[nextMiss2zwlz] {
				nextMiss2zwlz++
			}
			if nextMiss2zwlz == maxFields2zwlz {
				// filled all the empty fields!
				break doneWithStruct2zwlz
			}
			missingFieldsLeft2zwlz--
			curField2zwlz = decodeMsgFieldOrder2zwlz[nextMiss2zwlz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zwlz)
		switch curField2zwlz {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found2zwlz[0] = true
			var zefl uint32
			zefl, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zefl) {
				z.Methods = (z.Methods)[:zefl]
			} else {
				z.Methods = make([]MethodT, zefl)
			}
			for zhtm := range z.Methods {
				err = z.Methods[zhtm].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found2zwlz[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found2zwlz[2] = true
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
	if nextMiss2zwlz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder2zwlz = []string{"Methods", "Deprecated", "Comment"}

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
	for zhtm := range z.Methods {
		err = z.Methods[zhtm].EncodeMsg(en)
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
	for zhtm := range z.Methods {
		o, err = z.Methods[zhtm].MarshalMsg(o)
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
	const maxFields3zxay = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zxay uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zxay, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zxay := totalEncodedFields3zxay
	missingFieldsLeft3zxay := maxFields3zxay - totalEncodedFields3zxay

	var nextMiss3zxay int32 = -1
	var found3zxay [maxFields3zxay]bool
	var curField3zxay string

doneWithStruct3zxay:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zxay > 0 || missingFieldsLeft3zxay > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zxay, missingFieldsLeft3zxay, msgp.ShowFound(found3zxay[:]), unmarshalMsgFieldOrder3zxay)
		if encodedFieldsLeft3zxay > 0 {
			encodedFieldsLeft3zxay--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zxay = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zxay < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zxay = 0
			}
			for nextMiss3zxay < maxFields3zxay && found3zxay[nextMiss3zxay] {
				nextMiss3zxay++
			}
			if nextMiss3zxay == maxFields3zxay {
				// filled all the empty fields!
				break doneWithStruct3zxay
			}
			missingFieldsLeft3zxay--
			curField3zxay = unmarshalMsgFieldOrder3zxay[nextMiss3zxay]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zxay)
		switch curField3zxay {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found3zxay[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zazz uint32
				zazz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zazz) {
					z.Methods = (z.Methods)[:zazz]
				} else {
					z.Methods = make([]MethodT, zazz)
				}
				for zhtm := range z.Methods {
					bts, err = z.Methods[zhtm].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found3zxay[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found3zxay[2] = true
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
	if nextMiss3zxay != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder3zxay = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zhtm := range z.Methods {
		s += z.Methods[zhtm].Msgsize()
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
	const maxFields4zlgu = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zlgu uint32
	totalEncodedFields4zlgu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zlgu := totalEncodedFields4zlgu
	missingFieldsLeft4zlgu := maxFields4zlgu - totalEncodedFields4zlgu

	var nextMiss4zlgu int32 = -1
	var found4zlgu [maxFields4zlgu]bool
	var curField4zlgu string

doneWithStruct4zlgu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zlgu > 0 || missingFieldsLeft4zlgu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zlgu, missingFieldsLeft4zlgu, msgp.ShowFound(found4zlgu[:]), decodeMsgFieldOrder4zlgu)
		if encodedFieldsLeft4zlgu > 0 {
			encodedFieldsLeft4zlgu--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zlgu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zlgu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zlgu = 0
			}
			for nextMiss4zlgu < maxFields4zlgu && found4zlgu[nextMiss4zlgu] {
				nextMiss4zlgu++
			}
			if nextMiss4zlgu == maxFields4zlgu {
				// filled all the empty fields!
				break doneWithStruct4zlgu
			}
			missingFieldsLeft4zlgu--
			curField4zlgu = decodeMsgFieldOrder4zlgu[nextMiss4zlgu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zlgu)
		switch curField4zlgu {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found4zlgu[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found4zlgu[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found4zlgu[2] = true
			const maxFields5zvwj = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields5zvwj uint32
			totalEncodedFields5zvwj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft5zvwj := totalEncodedFields5zvwj
			missingFieldsLeft5zvwj := maxFields5zvwj - totalEncodedFields5zvwj

			var nextMiss5zvwj int32 = -1
			var found5zvwj [maxFields5zvwj]bool
			var curField5zvwj string

		doneWithStruct5zvwj:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft5zvwj > 0 || missingFieldsLeft5zvwj > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zvwj, missingFieldsLeft5zvwj, msgp.ShowFound(found5zvwj[:]), decodeMsgFieldOrder5zvwj)
				if encodedFieldsLeft5zvwj > 0 {
					encodedFieldsLeft5zvwj--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField5zvwj = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss5zvwj < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss5zvwj = 0
					}
					for nextMiss5zvwj < maxFields5zvwj && found5zvwj[nextMiss5zvwj] {
						nextMiss5zvwj++
					}
					if nextMiss5zvwj == maxFields5zvwj {
						// filled all the empty fields!
						break doneWithStruct5zvwj
					}
					missingFieldsLeft5zvwj--
					curField5zvwj = decodeMsgFieldOrder5zvwj[nextMiss5zvwj]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField5zvwj)
				switch curField5zvwj {
				// -- templateDecodeMsg ends here --

				case "StructName":
					found5zvwj[0] = true
					z.Inputs.StructName, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found5zvwj[1] = true
					var zetb uint32
					zetb, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zetb) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zetb]
					} else {
						z.Inputs.Fld = make([]FieldT, zetb)
					}
					for zvbb := range z.Inputs.Fld {
						err = z.Inputs.Fld[zvbb].DecodeMsg(dc)
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
			if nextMiss5zvwj != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found4zlgu[3] = true
			const maxFields6zajm = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields6zajm uint32
			totalEncodedFields6zajm, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft6zajm := totalEncodedFields6zajm
			missingFieldsLeft6zajm := maxFields6zajm - totalEncodedFields6zajm

			var nextMiss6zajm int32 = -1
			var found6zajm [maxFields6zajm]bool
			var curField6zajm string

		doneWithStruct6zajm:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft6zajm > 0 || missingFieldsLeft6zajm > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zajm, missingFieldsLeft6zajm, msgp.ShowFound(found6zajm[:]), decodeMsgFieldOrder6zajm)
				if encodedFieldsLeft6zajm > 0 {
					encodedFieldsLeft6zajm--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField6zajm = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss6zajm < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss6zajm = 0
					}
					for nextMiss6zajm < maxFields6zajm && found6zajm[nextMiss6zajm] {
						nextMiss6zajm++
					}
					if nextMiss6zajm == maxFields6zajm {
						// filled all the empty fields!
						break doneWithStruct6zajm
					}
					missingFieldsLeft6zajm--
					curField6zajm = decodeMsgFieldOrder6zajm[nextMiss6zajm]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField6zajm)
				switch curField6zajm {
				// -- templateDecodeMsg ends here --

				case "StructName":
					found6zajm[0] = true
					z.Returns.StructName, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found6zajm[1] = true
					var zlrt uint32
					zlrt, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zlrt) {
						z.Returns.Fld = (z.Returns.Fld)[:zlrt]
					} else {
						z.Returns.Fld = make([]FieldT, zlrt)
					}
					for ztnn := range z.Returns.Fld {
						err = z.Returns.Fld[ztnn].DecodeMsg(dc)
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
			if nextMiss6zajm != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found4zlgu[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zlgu[5] = true
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
	if nextMiss4zlgu != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder4zlgu = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder5zvwj = []string{"StructName", "Fld"}

// fields of StructT
var decodeMsgFieldOrder6zajm = []string{"StructName", "Fld"}

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
	// write "StructName"
	err = en.Append(0xa6, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Inputs.StructName)
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
	for zvbb := range z.Inputs.Fld {
		err = z.Inputs.Fld[zvbb].EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	// write "Returns"
	// map header, size 2
	// write "StructName"
	err = en.Append(0xa7, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Returns.StructName)
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
	for ztnn := range z.Returns.Fld {
		err = z.Returns.Fld[ztnn].EncodeMsg(en)
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
	// string "StructName"
	o = append(o, 0xa6, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Inputs.StructName)
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Inputs.Fld)))
	for zvbb := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zvbb].MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	// string "Returns"
	// map header, size 2
	// string "StructName"
	o = append(o, 0xa7, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Returns.StructName)
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Returns.Fld)))
	for ztnn := range z.Returns.Fld {
		o, err = z.Returns.Fld[ztnn].MarshalMsg(o)
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
	const maxFields7ztbd = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7ztbd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7ztbd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7ztbd := totalEncodedFields7ztbd
	missingFieldsLeft7ztbd := maxFields7ztbd - totalEncodedFields7ztbd

	var nextMiss7ztbd int32 = -1
	var found7ztbd [maxFields7ztbd]bool
	var curField7ztbd string

doneWithStruct7ztbd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7ztbd > 0 || missingFieldsLeft7ztbd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7ztbd, missingFieldsLeft7ztbd, msgp.ShowFound(found7ztbd[:]), unmarshalMsgFieldOrder7ztbd)
		if encodedFieldsLeft7ztbd > 0 {
			encodedFieldsLeft7ztbd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7ztbd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7ztbd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7ztbd = 0
			}
			for nextMiss7ztbd < maxFields7ztbd && found7ztbd[nextMiss7ztbd] {
				nextMiss7ztbd++
			}
			if nextMiss7ztbd == maxFields7ztbd {
				// filled all the empty fields!
				break doneWithStruct7ztbd
			}
			missingFieldsLeft7ztbd--
			curField7ztbd = unmarshalMsgFieldOrder7ztbd[nextMiss7ztbd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7ztbd)
		switch curField7ztbd {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found7ztbd[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found7ztbd[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found7ztbd[2] = true
			const maxFields8zfic = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields8zfic uint32
			if !nbs.AlwaysNil {
				totalEncodedFields8zfic, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft8zfic := totalEncodedFields8zfic
			missingFieldsLeft8zfic := maxFields8zfic - totalEncodedFields8zfic

			var nextMiss8zfic int32 = -1
			var found8zfic [maxFields8zfic]bool
			var curField8zfic string

		doneWithStruct8zfic:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zfic > 0 || missingFieldsLeft8zfic > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zfic, missingFieldsLeft8zfic, msgp.ShowFound(found8zfic[:]), unmarshalMsgFieldOrder8zfic)
				if encodedFieldsLeft8zfic > 0 {
					encodedFieldsLeft8zfic--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField8zfic = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zfic < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss8zfic = 0
					}
					for nextMiss8zfic < maxFields8zfic && found8zfic[nextMiss8zfic] {
						nextMiss8zfic++
					}
					if nextMiss8zfic == maxFields8zfic {
						// filled all the empty fields!
						break doneWithStruct8zfic
					}
					missingFieldsLeft8zfic--
					curField8zfic = unmarshalMsgFieldOrder8zfic[nextMiss8zfic]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zfic)
				switch curField8zfic {
				// -- templateUnmarshalMsg ends here --

				case "StructName":
					found8zfic[0] = true
					z.Inputs.StructName, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zfic[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zurj uint32
						zurj, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zurj) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zurj]
						} else {
							z.Inputs.Fld = make([]FieldT, zurj)
						}
						for zvbb := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zvbb].UnmarshalMsg(bts)
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
			if nextMiss8zfic != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found7ztbd[3] = true
			const maxFields9zxhz = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields9zxhz uint32
			if !nbs.AlwaysNil {
				totalEncodedFields9zxhz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft9zxhz := totalEncodedFields9zxhz
			missingFieldsLeft9zxhz := maxFields9zxhz - totalEncodedFields9zxhz

			var nextMiss9zxhz int32 = -1
			var found9zxhz [maxFields9zxhz]bool
			var curField9zxhz string

		doneWithStruct9zxhz:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft9zxhz > 0 || missingFieldsLeft9zxhz > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zxhz, missingFieldsLeft9zxhz, msgp.ShowFound(found9zxhz[:]), unmarshalMsgFieldOrder9zxhz)
				if encodedFieldsLeft9zxhz > 0 {
					encodedFieldsLeft9zxhz--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField9zxhz = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss9zxhz < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss9zxhz = 0
					}
					for nextMiss9zxhz < maxFields9zxhz && found9zxhz[nextMiss9zxhz] {
						nextMiss9zxhz++
					}
					if nextMiss9zxhz == maxFields9zxhz {
						// filled all the empty fields!
						break doneWithStruct9zxhz
					}
					missingFieldsLeft9zxhz--
					curField9zxhz = unmarshalMsgFieldOrder9zxhz[nextMiss9zxhz]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField9zxhz)
				switch curField9zxhz {
				// -- templateUnmarshalMsg ends here --

				case "StructName":
					found9zxhz[0] = true
					z.Returns.StructName, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found9zxhz[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var ztgd uint32
						ztgd, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(ztgd) {
							z.Returns.Fld = (z.Returns.Fld)[:ztgd]
						} else {
							z.Returns.Fld = make([]FieldT, ztgd)
						}
						for ztnn := range z.Returns.Fld {
							bts, err = z.Returns.Fld[ztnn].UnmarshalMsg(bts)
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
			if nextMiss9zxhz != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found7ztbd[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found7ztbd[5] = true
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
	if nextMiss7ztbd != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder7ztbd = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder8zfic = []string{"StructName", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder9zxhz = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 11 + msgp.StringPrefixSize + len(z.Inputs.StructName) + 4 + msgp.ArrayHeaderSize
	for zvbb := range z.Inputs.Fld {
		s += z.Inputs.Fld[zvbb].Msgsize()
	}
	s += 8 + 1 + 11 + msgp.StringPrefixSize + len(z.Returns.StructName) + 4 + msgp.ArrayHeaderSize
	for ztnn := range z.Returns.Fld {
		s += z.Returns.Fld[ztnn].Msgsize()
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
	const maxFields10zizh = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zizh uint32
	totalEncodedFields10zizh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zizh := totalEncodedFields10zizh
	missingFieldsLeft10zizh := maxFields10zizh - totalEncodedFields10zizh

	var nextMiss10zizh int32 = -1
	var found10zizh [maxFields10zizh]bool
	var curField10zizh string

doneWithStruct10zizh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zizh > 0 || missingFieldsLeft10zizh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zizh, missingFieldsLeft10zizh, msgp.ShowFound(found10zizh[:]), decodeMsgFieldOrder10zizh)
		if encodedFieldsLeft10zizh > 0 {
			encodedFieldsLeft10zizh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zizh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zizh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zizh = 0
			}
			for nextMiss10zizh < maxFields10zizh && found10zizh[nextMiss10zizh] {
				nextMiss10zizh++
			}
			if nextMiss10zizh == maxFields10zizh {
				// filled all the empty fields!
				break doneWithStruct10zizh
			}
			missingFieldsLeft10zizh--
			curField10zizh = decodeMsgFieldOrder10zizh[nextMiss10zizh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zizh)
		switch curField10zizh {
		// -- templateDecodeMsg ends here --

		case "SchemaId":
			found10zizh[0] = true
			z.SchemaId, err = dc.ReadUint64()
			if err != nil {
				panic(err)
			}
		case "IntToPackageTable":
			found10zizh[1] = true
			var znhj uint32
			znhj, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.IntToPackageTable == nil && znhj > 0 {
				z.IntToPackageTable = make(map[int64]string, znhj)
			} else if len(z.IntToPackageTable) > 0 {
				for key, _ := range z.IntToPackageTable {
					delete(z.IntToPackageTable, key)
				}
			}
			for znhj > 0 {
				znhj--
				var zkxy int64
				var zveh string
				zkxy, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				zveh, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.IntToPackageTable[zkxy] = zveh
			}
		case "PkgPath":
			found10zizh[2] = true
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
	if nextMiss10zizh != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder10zizh = []string{"SchemaId", "IntToPackageTable", "PkgPath"}

// fieldsNotEmpty supports omitempty tags
func (z *SchemaT) fieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// EncodeMsg implements msgp.Encodable
func (z *SchemaT) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "SchemaId"
	err = en.Append(0x83, 0xa8, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.SchemaId)
	if err != nil {
		panic(err)
	}
	// write "IntToPackageTable"
	err = en.Append(0xb1, 0x49, 0x6e, 0x74, 0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.IntToPackageTable)))
	if err != nil {
		panic(err)
	}
	for zkxy, zveh := range z.IntToPackageTable {
		err = en.WriteInt64(zkxy)
		if err != nil {
			panic(err)
		}
		err = en.WriteString(zveh)
		if err != nil {
			panic(err)
		}
	}
	// write "PkgPath"
	err = en.Append(0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
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
func (z *SchemaT) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "SchemaId"
	o = append(o, 0x83, 0xa8, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.SchemaId)
	// string "IntToPackageTable"
	o = append(o, 0xb1, 0x49, 0x6e, 0x74, 0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65)
	o = msgp.AppendMapHeader(o, uint32(len(z.IntToPackageTable)))
	for zkxy, zveh := range z.IntToPackageTable {
		o = msgp.AppendInt64(o, zkxy)
		o = msgp.AppendString(o, zveh)
	}
	// string "PkgPath"
	o = append(o, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
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
	const maxFields11zmqz = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields11zmqz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields11zmqz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft11zmqz := totalEncodedFields11zmqz
	missingFieldsLeft11zmqz := maxFields11zmqz - totalEncodedFields11zmqz

	var nextMiss11zmqz int32 = -1
	var found11zmqz [maxFields11zmqz]bool
	var curField11zmqz string

doneWithStruct11zmqz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft11zmqz > 0 || missingFieldsLeft11zmqz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zmqz, missingFieldsLeft11zmqz, msgp.ShowFound(found11zmqz[:]), unmarshalMsgFieldOrder11zmqz)
		if encodedFieldsLeft11zmqz > 0 {
			encodedFieldsLeft11zmqz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField11zmqz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss11zmqz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss11zmqz = 0
			}
			for nextMiss11zmqz < maxFields11zmqz && found11zmqz[nextMiss11zmqz] {
				nextMiss11zmqz++
			}
			if nextMiss11zmqz == maxFields11zmqz {
				// filled all the empty fields!
				break doneWithStruct11zmqz
			}
			missingFieldsLeft11zmqz--
			curField11zmqz = unmarshalMsgFieldOrder11zmqz[nextMiss11zmqz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField11zmqz)
		switch curField11zmqz {
		// -- templateUnmarshalMsg ends here --

		case "SchemaId":
			found11zmqz[0] = true
			z.SchemaId, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "IntToPackageTable":
			found11zmqz[1] = true
			if nbs.AlwaysNil {
				if len(z.IntToPackageTable) > 0 {
					for key, _ := range z.IntToPackageTable {
						delete(z.IntToPackageTable, key)
					}
				}

			} else {

				var zjip uint32
				zjip, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.IntToPackageTable == nil && zjip > 0 {
					z.IntToPackageTable = make(map[int64]string, zjip)
				} else if len(z.IntToPackageTable) > 0 {
					for key, _ := range z.IntToPackageTable {
						delete(z.IntToPackageTable, key)
					}
				}
				for zjip > 0 {
					var zkxy int64
					var zveh string
					zjip--
					zkxy, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					zveh, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.IntToPackageTable[zkxy] = zveh
				}
			}
		case "PkgPath":
			found11zmqz[2] = true
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
	if nextMiss11zmqz != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder11zmqz = []string{"SchemaId", "IntToPackageTable", "PkgPath"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Uint64Size + 18 + msgp.MapHeaderSize
	if z.IntToPackageTable != nil {
		for zkxy, zveh := range z.IntToPackageTable {
			_ = zveh
			_ = zkxy
			s += msgp.Int64Size + msgp.StringPrefixSize + len(zveh)
		}
	}
	s += 8 + msgp.StringPrefixSize + len(z.PkgPath)
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
	const maxFields12ziiw = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12ziiw uint32
	totalEncodedFields12ziiw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12ziiw := totalEncodedFields12ziiw
	missingFieldsLeft12ziiw := maxFields12ziiw - totalEncodedFields12ziiw

	var nextMiss12ziiw int32 = -1
	var found12ziiw [maxFields12ziiw]bool
	var curField12ziiw string

doneWithStruct12ziiw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12ziiw > 0 || missingFieldsLeft12ziiw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12ziiw, missingFieldsLeft12ziiw, msgp.ShowFound(found12ziiw[:]), decodeMsgFieldOrder12ziiw)
		if encodedFieldsLeft12ziiw > 0 {
			encodedFieldsLeft12ziiw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12ziiw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12ziiw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12ziiw = 0
			}
			for nextMiss12ziiw < maxFields12ziiw && found12ziiw[nextMiss12ziiw] {
				nextMiss12ziiw++
			}
			if nextMiss12ziiw == maxFields12ziiw {
				// filled all the empty fields!
				break doneWithStruct12ziiw
			}
			missingFieldsLeft12ziiw--
			curField12ziiw = decodeMsgFieldOrder12ziiw[nextMiss12ziiw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12ziiw)
		switch curField12ziiw {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found12ziiw[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found12ziiw[1] = true
			var zfai uint32
			zfai, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zfai) {
				z.Fld = (z.Fld)[:zfai]
			} else {
				z.Fld = make([]FieldT, zfai)
			}
			for zirg := range z.Fld {
				err = z.Fld[zirg].DecodeMsg(dc)
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
	if nextMiss12ziiw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder12ziiw = []string{"StructName", "Fld"}

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
	for zirg := range z.Fld {
		err = z.Fld[zirg].EncodeMsg(en)
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
	for zirg := range z.Fld {
		o, err = z.Fld[zirg].MarshalMsg(o)
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
	const maxFields13zyib = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields13zyib uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zyib, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft13zyib := totalEncodedFields13zyib
	missingFieldsLeft13zyib := maxFields13zyib - totalEncodedFields13zyib

	var nextMiss13zyib int32 = -1
	var found13zyib [maxFields13zyib]bool
	var curField13zyib string

doneWithStruct13zyib:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zyib > 0 || missingFieldsLeft13zyib > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zyib, missingFieldsLeft13zyib, msgp.ShowFound(found13zyib[:]), unmarshalMsgFieldOrder13zyib)
		if encodedFieldsLeft13zyib > 0 {
			encodedFieldsLeft13zyib--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField13zyib = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss13zyib < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zyib = 0
			}
			for nextMiss13zyib < maxFields13zyib && found13zyib[nextMiss13zyib] {
				nextMiss13zyib++
			}
			if nextMiss13zyib == maxFields13zyib {
				// filled all the empty fields!
				break doneWithStruct13zyib
			}
			missingFieldsLeft13zyib--
			curField13zyib = unmarshalMsgFieldOrder13zyib[nextMiss13zyib]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zyib)
		switch curField13zyib {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found13zyib[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found13zyib[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zihu uint32
				zihu, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zihu) {
					z.Fld = (z.Fld)[:zihu]
				} else {
					z.Fld = make([]FieldT, zihu)
				}
				for zirg := range z.Fld {
					bts, err = z.Fld[zirg].UnmarshalMsg(bts)
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
	if nextMiss13zyib != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder13zyib = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zirg := range z.Fld {
		s += z.Fld[zirg].Msgsize()
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
		var zxsl int64
		zxsl, err = dc.ReadInt64()
		(*z) = StructTypeId(zxsl)
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
		var zsjq int64
		zsjq, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zsjq)
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
func (z *ZKind) DecodeMsg(dc *msgp.Reader) (err error) {
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
		var zofs byte
		zofs, err = dc.ReadByte()
		(*z) = ZKind(zofs)
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
func (z ZKind) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteByte(byte(z))
	if err != nil {
		panic(err)
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ZKind) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendByte(o, byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ZKind) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *ZKind) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zywl byte
		zywl, bts, err = nbs.ReadByteBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = ZKind(zywl)
	}
	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ZKind) Msgsize() (s int) {
	s = msgp.ByteSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *ZStructPacket) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields14zabo = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields14zabo uint32
	totalEncodedFields14zabo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zabo := totalEncodedFields14zabo
	missingFieldsLeft14zabo := maxFields14zabo - totalEncodedFields14zabo

	var nextMiss14zabo int32 = -1
	var found14zabo [maxFields14zabo]bool
	var curField14zabo string

doneWithStruct14zabo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zabo > 0 || missingFieldsLeft14zabo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zabo, missingFieldsLeft14zabo, msgp.ShowFound(found14zabo[:]), decodeMsgFieldOrder14zabo)
		if encodedFieldsLeft14zabo > 0 {
			encodedFieldsLeft14zabo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField14zabo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zabo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zabo = 0
			}
			for nextMiss14zabo < maxFields14zabo && found14zabo[nextMiss14zabo] {
				nextMiss14zabo++
			}
			if nextMiss14zabo == maxFields14zabo {
				// filled all the empty fields!
				break doneWithStruct14zabo
			}
			missingFieldsLeft14zabo--
			curField14zabo = decodeMsgFieldOrder14zabo[nextMiss14zabo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zabo)
		switch curField14zabo {
		// -- templateDecodeMsg ends here --

		case "Id":
			found14zabo[0] = true
			{
				var zolw int64
				zolw, err = dc.ReadInt64()
				z.Id = StructTypeId(zolw)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found14zabo[1] = true
			var zkqr uint32
			zkqr, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zkqr > 0 {
				z.Da = make(map[int64]msgp.Raw, zkqr)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zkqr > 0 {
				zkqr--
				var zdvs int64
				var zull msgp.Raw
				zdvs, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zull.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zdvs] = zull
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zabo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZStructPacket
var decodeMsgFieldOrder14zabo = []string{"Id", "Da"}

// fieldsNotEmpty supports omitempty tags
func (z *ZStructPacket) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *ZStructPacket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Id"
	err = en.Append(0x82, 0xa2, 0x49, 0x64)
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
	for zdvs, zull := range z.Da {
		err = en.WriteInt64(zdvs)
		if err != nil {
			panic(err)
		}
		err = zull.EncodeMsg(en)
		if err != nil {
			panic(err)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ZStructPacket) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Id"
	o = append(o, 0x82, 0xa2, 0x49, 0x64)
	o = msgp.AppendInt64(o, int64(z.Id))
	// string "Da"
	o = append(o, 0xa2, 0x44, 0x61)
	o = msgp.AppendMapHeader(o, uint32(len(z.Da)))
	for zdvs, zull := range z.Da {
		o = msgp.AppendInt64(o, zdvs)
		o, err = zull.MarshalMsg(o)
		if err != nil {
			panic(err)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ZStructPacket) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *ZStructPacket) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields15zhcc = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields15zhcc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zhcc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft15zhcc := totalEncodedFields15zhcc
	missingFieldsLeft15zhcc := maxFields15zhcc - totalEncodedFields15zhcc

	var nextMiss15zhcc int32 = -1
	var found15zhcc [maxFields15zhcc]bool
	var curField15zhcc string

doneWithStruct15zhcc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zhcc > 0 || missingFieldsLeft15zhcc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zhcc, missingFieldsLeft15zhcc, msgp.ShowFound(found15zhcc[:]), unmarshalMsgFieldOrder15zhcc)
		if encodedFieldsLeft15zhcc > 0 {
			encodedFieldsLeft15zhcc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField15zhcc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss15zhcc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zhcc = 0
			}
			for nextMiss15zhcc < maxFields15zhcc && found15zhcc[nextMiss15zhcc] {
				nextMiss15zhcc++
			}
			if nextMiss15zhcc == maxFields15zhcc {
				// filled all the empty fields!
				break doneWithStruct15zhcc
			}
			missingFieldsLeft15zhcc--
			curField15zhcc = unmarshalMsgFieldOrder15zhcc[nextMiss15zhcc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zhcc)
		switch curField15zhcc {
		// -- templateUnmarshalMsg ends here --

		case "Id":
			found15zhcc[0] = true
			{
				var zyvb int64
				zyvb, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zyvb)
			}
		case "Da":
			found15zhcc[1] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zlmc uint32
				zlmc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zlmc > 0 {
					z.Da = make(map[int64]msgp.Raw, zlmc)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zlmc > 0 {
					var zdvs int64
					var zull msgp.Raw
					zlmc--
					zdvs, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zull.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zdvs] = zull
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss15zhcc != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZStructPacket
var unmarshalMsgFieldOrder15zhcc = []string{"Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZStructPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zdvs, zull := range z.Da {
			_ = zull
			_ = zdvs
			s += msgp.Int64Size + zull.Msgsize()
		}
	}
	return
}
