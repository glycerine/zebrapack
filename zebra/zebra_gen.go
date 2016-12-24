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
	const maxFields0zvfg = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zvfg uint32
	totalEncodedFields0zvfg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zvfg := totalEncodedFields0zvfg
	missingFieldsLeft0zvfg := maxFields0zvfg - totalEncodedFields0zvfg

	var nextMiss0zvfg int32 = -1
	var found0zvfg [maxFields0zvfg]bool
	var curField0zvfg string

doneWithStruct0zvfg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zvfg > 0 || missingFieldsLeft0zvfg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zvfg, missingFieldsLeft0zvfg, msgp.ShowFound(found0zvfg[:]), decodeMsgFieldOrder0zvfg)
		if encodedFieldsLeft0zvfg > 0 {
			encodedFieldsLeft0zvfg--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zvfg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zvfg < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zvfg = 0
			}
			for nextMiss0zvfg < maxFields0zvfg && found0zvfg[nextMiss0zvfg] {
				nextMiss0zvfg++
			}
			if nextMiss0zvfg == maxFields0zvfg {
				// filled all the empty fields!
				break doneWithStruct0zvfg
			}
			missingFieldsLeft0zvfg--
			curField0zvfg = decodeMsgFieldOrder0zvfg[nextMiss0zvfg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zvfg)
		switch curField0zvfg {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zvfg[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zvfg[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Zknd":
			found0zvfg[2] = true
			var zpyd uint32
			zpyd, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Zknd) >= int(zpyd) {
				z.Zknd = (z.Zknd)[:zpyd]
			} else {
				z.Zknd = make([]ZKind, zpyd)
			}
			for zbeg := range z.Zknd {
				{
					var zklq byte
					zklq, err = dc.ReadByte()
					z.Zknd[zbeg] = ZKind(zklq)
				}
				if err != nil {
					panic(err)
				}
			}
		case "Varg":
			found0zvfg[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zvfg[4] = true
			var zbxv uint32
			zbxv, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zbxv > 0 {
				z.Tag = make(map[string]string, zbxv)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zbxv > 0 {
				zbxv--
				var zcaa string
				var zddv string
				zcaa, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zddv, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zcaa] = zddv
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zvfg != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zvfg = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

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
	var empty_zqna [5]bool
	fieldsInUse_zgcv := z.fieldsNotEmpty(empty_zqna[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgcv)
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
	for zbeg := range z.Zknd {
		err = en.WriteByte(byte(z.Zknd[zbeg]))
		if err != nil {
			panic(err)
		}
	}
	if !empty_zqna[3] {
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

	if !empty_zqna[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zcaa, zddv := range z.Tag {
			err = en.WriteString(zcaa)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zddv)
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
	for zbeg := range z.Zknd {
		o = msgp.AppendByte(o, byte(z.Zknd[zbeg]))
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
		for zcaa, zddv := range z.Tag {
			o = msgp.AppendString(o, zcaa)
			o = msgp.AppendString(o, zddv)
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
	const maxFields1zjyq = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zjyq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zjyq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zjyq := totalEncodedFields1zjyq
	missingFieldsLeft1zjyq := maxFields1zjyq - totalEncodedFields1zjyq

	var nextMiss1zjyq int32 = -1
	var found1zjyq [maxFields1zjyq]bool
	var curField1zjyq string

doneWithStruct1zjyq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zjyq > 0 || missingFieldsLeft1zjyq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zjyq, missingFieldsLeft1zjyq, msgp.ShowFound(found1zjyq[:]), unmarshalMsgFieldOrder1zjyq)
		if encodedFieldsLeft1zjyq > 0 {
			encodedFieldsLeft1zjyq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zjyq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zjyq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zjyq = 0
			}
			for nextMiss1zjyq < maxFields1zjyq && found1zjyq[nextMiss1zjyq] {
				nextMiss1zjyq++
			}
			if nextMiss1zjyq == maxFields1zjyq {
				// filled all the empty fields!
				break doneWithStruct1zjyq
			}
			missingFieldsLeft1zjyq--
			curField1zjyq = unmarshalMsgFieldOrder1zjyq[nextMiss1zjyq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zjyq)
		switch curField1zjyq {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zjyq[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zjyq[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Zknd":
			found1zjyq[2] = true
			if nbs.AlwaysNil {
				(z.Zknd) = (z.Zknd)[:0]
			} else {

				var zwbo uint32
				zwbo, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Zknd) >= int(zwbo) {
					z.Zknd = (z.Zknd)[:zwbo]
				} else {
					z.Zknd = make([]ZKind, zwbo)
				}
				for zbeg := range z.Zknd {
					{
						var zbvt byte
						zbvt, bts, err = nbs.ReadByteBytes(bts)

						if err != nil {
							panic(err)
						}
						z.Zknd[zbeg] = ZKind(zbvt)
					}
				}
			}
		case "Varg":
			found1zjyq[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zjyq[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zomk uint32
				zomk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zomk > 0 {
					z.Tag = make(map[string]string, zomk)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zomk > 0 {
					var zcaa string
					var zddv string
					zomk--
					zcaa, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zddv, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zcaa] = zddv
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zjyq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zjyq = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize + (len(z.Zknd) * (msgp.ByteSize)) + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zcaa, zddv := range z.Tag {
			_ = zddv
			_ = zcaa
			s += msgp.StringPrefixSize + len(zcaa) + msgp.StringPrefixSize + len(zddv)
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
		var zmpy int64
		zmpy, err = dc.ReadInt64()
		(*z) = InterfaceId(zmpy)
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
		var zyep int64
		zyep, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = InterfaceId(zyep)
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
	const maxFields2zdje = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zdje uint32
	totalEncodedFields2zdje, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zdje := totalEncodedFields2zdje
	missingFieldsLeft2zdje := maxFields2zdje - totalEncodedFields2zdje

	var nextMiss2zdje int32 = -1
	var found2zdje [maxFields2zdje]bool
	var curField2zdje string

doneWithStruct2zdje:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zdje > 0 || missingFieldsLeft2zdje > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zdje, missingFieldsLeft2zdje, msgp.ShowFound(found2zdje[:]), decodeMsgFieldOrder2zdje)
		if encodedFieldsLeft2zdje > 0 {
			encodedFieldsLeft2zdje--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zdje = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zdje < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zdje = 0
			}
			for nextMiss2zdje < maxFields2zdje && found2zdje[nextMiss2zdje] {
				nextMiss2zdje++
			}
			if nextMiss2zdje == maxFields2zdje {
				// filled all the empty fields!
				break doneWithStruct2zdje
			}
			missingFieldsLeft2zdje--
			curField2zdje = decodeMsgFieldOrder2zdje[nextMiss2zdje]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zdje)
		switch curField2zdje {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found2zdje[0] = true
			var zgax uint32
			zgax, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zgax) {
				z.Methods = (z.Methods)[:zgax]
			} else {
				z.Methods = make([]MethodT, zgax)
			}
			for zogy := range z.Methods {
				err = z.Methods[zogy].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found2zdje[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found2zdje[2] = true
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
	if nextMiss2zdje != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder2zdje = []string{"Methods", "Deprecated", "Comment"}

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
	for zogy := range z.Methods {
		err = z.Methods[zogy].EncodeMsg(en)
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
	for zogy := range z.Methods {
		o, err = z.Methods[zogy].MarshalMsg(o)
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
	const maxFields3zpfz = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zpfz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zpfz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zpfz := totalEncodedFields3zpfz
	missingFieldsLeft3zpfz := maxFields3zpfz - totalEncodedFields3zpfz

	var nextMiss3zpfz int32 = -1
	var found3zpfz [maxFields3zpfz]bool
	var curField3zpfz string

doneWithStruct3zpfz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zpfz > 0 || missingFieldsLeft3zpfz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zpfz, missingFieldsLeft3zpfz, msgp.ShowFound(found3zpfz[:]), unmarshalMsgFieldOrder3zpfz)
		if encodedFieldsLeft3zpfz > 0 {
			encodedFieldsLeft3zpfz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zpfz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zpfz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zpfz = 0
			}
			for nextMiss3zpfz < maxFields3zpfz && found3zpfz[nextMiss3zpfz] {
				nextMiss3zpfz++
			}
			if nextMiss3zpfz == maxFields3zpfz {
				// filled all the empty fields!
				break doneWithStruct3zpfz
			}
			missingFieldsLeft3zpfz--
			curField3zpfz = unmarshalMsgFieldOrder3zpfz[nextMiss3zpfz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zpfz)
		switch curField3zpfz {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found3zpfz[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zvmt uint32
				zvmt, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zvmt) {
					z.Methods = (z.Methods)[:zvmt]
				} else {
					z.Methods = make([]MethodT, zvmt)
				}
				for zogy := range z.Methods {
					bts, err = z.Methods[zogy].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found3zpfz[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found3zpfz[2] = true
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
	if nextMiss3zpfz != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder3zpfz = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zogy := range z.Methods {
		s += z.Methods[zogy].Msgsize()
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
	const maxFields4zbtf = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zbtf uint32
	totalEncodedFields4zbtf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zbtf := totalEncodedFields4zbtf
	missingFieldsLeft4zbtf := maxFields4zbtf - totalEncodedFields4zbtf

	var nextMiss4zbtf int32 = -1
	var found4zbtf [maxFields4zbtf]bool
	var curField4zbtf string

doneWithStruct4zbtf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zbtf > 0 || missingFieldsLeft4zbtf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zbtf, missingFieldsLeft4zbtf, msgp.ShowFound(found4zbtf[:]), decodeMsgFieldOrder4zbtf)
		if encodedFieldsLeft4zbtf > 0 {
			encodedFieldsLeft4zbtf--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zbtf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zbtf < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zbtf = 0
			}
			for nextMiss4zbtf < maxFields4zbtf && found4zbtf[nextMiss4zbtf] {
				nextMiss4zbtf++
			}
			if nextMiss4zbtf == maxFields4zbtf {
				// filled all the empty fields!
				break doneWithStruct4zbtf
			}
			missingFieldsLeft4zbtf--
			curField4zbtf = decodeMsgFieldOrder4zbtf[nextMiss4zbtf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zbtf)
		switch curField4zbtf {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found4zbtf[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found4zbtf[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found4zbtf[2] = true
			const maxFields5zebr = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields5zebr uint32
			totalEncodedFields5zebr, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft5zebr := totalEncodedFields5zebr
			missingFieldsLeft5zebr := maxFields5zebr - totalEncodedFields5zebr

			var nextMiss5zebr int32 = -1
			var found5zebr [maxFields5zebr]bool
			var curField5zebr string

		doneWithStruct5zebr:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft5zebr > 0 || missingFieldsLeft5zebr > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zebr, missingFieldsLeft5zebr, msgp.ShowFound(found5zebr[:]), decodeMsgFieldOrder5zebr)
				if encodedFieldsLeft5zebr > 0 {
					encodedFieldsLeft5zebr--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField5zebr = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss5zebr < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss5zebr = 0
					}
					for nextMiss5zebr < maxFields5zebr && found5zebr[nextMiss5zebr] {
						nextMiss5zebr++
					}
					if nextMiss5zebr == maxFields5zebr {
						// filled all the empty fields!
						break doneWithStruct5zebr
					}
					missingFieldsLeft5zebr--
					curField5zebr = decodeMsgFieldOrder5zebr[nextMiss5zebr]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField5zebr)
				switch curField5zebr {
				// -- templateDecodeMsg ends here --

				case "StructName":
					found5zebr[0] = true
					z.Inputs.StructName, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found5zebr[1] = true
					var ziuz uint32
					ziuz, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(ziuz) {
						z.Inputs.Fld = (z.Inputs.Fld)[:ziuz]
					} else {
						z.Inputs.Fld = make([]FieldT, ziuz)
					}
					for zlks := range z.Inputs.Fld {
						err = z.Inputs.Fld[zlks].DecodeMsg(dc)
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
			if nextMiss5zebr != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found4zbtf[3] = true
			const maxFields6zoom = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields6zoom uint32
			totalEncodedFields6zoom, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft6zoom := totalEncodedFields6zoom
			missingFieldsLeft6zoom := maxFields6zoom - totalEncodedFields6zoom

			var nextMiss6zoom int32 = -1
			var found6zoom [maxFields6zoom]bool
			var curField6zoom string

		doneWithStruct6zoom:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft6zoom > 0 || missingFieldsLeft6zoom > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zoom, missingFieldsLeft6zoom, msgp.ShowFound(found6zoom[:]), decodeMsgFieldOrder6zoom)
				if encodedFieldsLeft6zoom > 0 {
					encodedFieldsLeft6zoom--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField6zoom = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss6zoom < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss6zoom = 0
					}
					for nextMiss6zoom < maxFields6zoom && found6zoom[nextMiss6zoom] {
						nextMiss6zoom++
					}
					if nextMiss6zoom == maxFields6zoom {
						// filled all the empty fields!
						break doneWithStruct6zoom
					}
					missingFieldsLeft6zoom--
					curField6zoom = decodeMsgFieldOrder6zoom[nextMiss6zoom]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField6zoom)
				switch curField6zoom {
				// -- templateDecodeMsg ends here --

				case "StructName":
					found6zoom[0] = true
					z.Returns.StructName, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found6zoom[1] = true
					var ztdv uint32
					ztdv, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(ztdv) {
						z.Returns.Fld = (z.Returns.Fld)[:ztdv]
					} else {
						z.Returns.Fld = make([]FieldT, ztdv)
					}
					for zdib := range z.Returns.Fld {
						err = z.Returns.Fld[zdib].DecodeMsg(dc)
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
			if nextMiss6zoom != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found4zbtf[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zbtf[5] = true
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
	if nextMiss4zbtf != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder4zbtf = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder5zebr = []string{"StructName", "Fld"}

// fields of StructT
var decodeMsgFieldOrder6zoom = []string{"StructName", "Fld"}

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
	for zlks := range z.Inputs.Fld {
		err = z.Inputs.Fld[zlks].EncodeMsg(en)
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
	for zdib := range z.Returns.Fld {
		err = z.Returns.Fld[zdib].EncodeMsg(en)
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
	for zlks := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zlks].MarshalMsg(o)
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
	for zdib := range z.Returns.Fld {
		o, err = z.Returns.Fld[zdib].MarshalMsg(o)
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
	const maxFields7zzvn = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zzvn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zzvn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zzvn := totalEncodedFields7zzvn
	missingFieldsLeft7zzvn := maxFields7zzvn - totalEncodedFields7zzvn

	var nextMiss7zzvn int32 = -1
	var found7zzvn [maxFields7zzvn]bool
	var curField7zzvn string

doneWithStruct7zzvn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zzvn > 0 || missingFieldsLeft7zzvn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zzvn, missingFieldsLeft7zzvn, msgp.ShowFound(found7zzvn[:]), unmarshalMsgFieldOrder7zzvn)
		if encodedFieldsLeft7zzvn > 0 {
			encodedFieldsLeft7zzvn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zzvn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zzvn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zzvn = 0
			}
			for nextMiss7zzvn < maxFields7zzvn && found7zzvn[nextMiss7zzvn] {
				nextMiss7zzvn++
			}
			if nextMiss7zzvn == maxFields7zzvn {
				// filled all the empty fields!
				break doneWithStruct7zzvn
			}
			missingFieldsLeft7zzvn--
			curField7zzvn = unmarshalMsgFieldOrder7zzvn[nextMiss7zzvn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zzvn)
		switch curField7zzvn {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found7zzvn[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found7zzvn[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found7zzvn[2] = true
			const maxFields8zzjd = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields8zzjd uint32
			if !nbs.AlwaysNil {
				totalEncodedFields8zzjd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft8zzjd := totalEncodedFields8zzjd
			missingFieldsLeft8zzjd := maxFields8zzjd - totalEncodedFields8zzjd

			var nextMiss8zzjd int32 = -1
			var found8zzjd [maxFields8zzjd]bool
			var curField8zzjd string

		doneWithStruct8zzjd:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zzjd > 0 || missingFieldsLeft8zzjd > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zzjd, missingFieldsLeft8zzjd, msgp.ShowFound(found8zzjd[:]), unmarshalMsgFieldOrder8zzjd)
				if encodedFieldsLeft8zzjd > 0 {
					encodedFieldsLeft8zzjd--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField8zzjd = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zzjd < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss8zzjd = 0
					}
					for nextMiss8zzjd < maxFields8zzjd && found8zzjd[nextMiss8zzjd] {
						nextMiss8zzjd++
					}
					if nextMiss8zzjd == maxFields8zzjd {
						// filled all the empty fields!
						break doneWithStruct8zzjd
					}
					missingFieldsLeft8zzjd--
					curField8zzjd = unmarshalMsgFieldOrder8zzjd[nextMiss8zzjd]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zzjd)
				switch curField8zzjd {
				// -- templateUnmarshalMsg ends here --

				case "StructName":
					found8zzjd[0] = true
					z.Inputs.StructName, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zzjd[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zgbz uint32
						zgbz, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zgbz) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zgbz]
						} else {
							z.Inputs.Fld = make([]FieldT, zgbz)
						}
						for zlks := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zlks].UnmarshalMsg(bts)
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
			if nextMiss8zzjd != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found7zzvn[3] = true
			const maxFields9zuuw = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields9zuuw uint32
			if !nbs.AlwaysNil {
				totalEncodedFields9zuuw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft9zuuw := totalEncodedFields9zuuw
			missingFieldsLeft9zuuw := maxFields9zuuw - totalEncodedFields9zuuw

			var nextMiss9zuuw int32 = -1
			var found9zuuw [maxFields9zuuw]bool
			var curField9zuuw string

		doneWithStruct9zuuw:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft9zuuw > 0 || missingFieldsLeft9zuuw > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zuuw, missingFieldsLeft9zuuw, msgp.ShowFound(found9zuuw[:]), unmarshalMsgFieldOrder9zuuw)
				if encodedFieldsLeft9zuuw > 0 {
					encodedFieldsLeft9zuuw--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField9zuuw = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss9zuuw < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss9zuuw = 0
					}
					for nextMiss9zuuw < maxFields9zuuw && found9zuuw[nextMiss9zuuw] {
						nextMiss9zuuw++
					}
					if nextMiss9zuuw == maxFields9zuuw {
						// filled all the empty fields!
						break doneWithStruct9zuuw
					}
					missingFieldsLeft9zuuw--
					curField9zuuw = unmarshalMsgFieldOrder9zuuw[nextMiss9zuuw]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField9zuuw)
				switch curField9zuuw {
				// -- templateUnmarshalMsg ends here --

				case "StructName":
					found9zuuw[0] = true
					z.Returns.StructName, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found9zuuw[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zzlb uint32
						zzlb, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zzlb) {
							z.Returns.Fld = (z.Returns.Fld)[:zzlb]
						} else {
							z.Returns.Fld = make([]FieldT, zzlb)
						}
						for zdib := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zdib].UnmarshalMsg(bts)
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
			if nextMiss9zuuw != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found7zzvn[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found7zzvn[5] = true
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
	if nextMiss7zzvn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder7zzvn = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder8zzjd = []string{"StructName", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder9zuuw = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 11 + msgp.StringPrefixSize + len(z.Inputs.StructName) + 4 + msgp.ArrayHeaderSize
	for zlks := range z.Inputs.Fld {
		s += z.Inputs.Fld[zlks].Msgsize()
	}
	s += 8 + 1 + 11 + msgp.StringPrefixSize + len(z.Returns.StructName) + 4 + msgp.ArrayHeaderSize
	for zdib := range z.Returns.Fld {
		s += z.Returns.Fld[zdib].Msgsize()
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
	const maxFields10zakq = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zakq uint32
	totalEncodedFields10zakq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zakq := totalEncodedFields10zakq
	missingFieldsLeft10zakq := maxFields10zakq - totalEncodedFields10zakq

	var nextMiss10zakq int32 = -1
	var found10zakq [maxFields10zakq]bool
	var curField10zakq string

doneWithStruct10zakq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zakq > 0 || missingFieldsLeft10zakq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zakq, missingFieldsLeft10zakq, msgp.ShowFound(found10zakq[:]), decodeMsgFieldOrder10zakq)
		if encodedFieldsLeft10zakq > 0 {
			encodedFieldsLeft10zakq--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zakq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zakq < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zakq = 0
			}
			for nextMiss10zakq < maxFields10zakq && found10zakq[nextMiss10zakq] {
				nextMiss10zakq++
			}
			if nextMiss10zakq == maxFields10zakq {
				// filled all the empty fields!
				break doneWithStruct10zakq
			}
			missingFieldsLeft10zakq--
			curField10zakq = decodeMsgFieldOrder10zakq[nextMiss10zakq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zakq)
		switch curField10zakq {
		// -- templateDecodeMsg ends here --

		case "SchemaId":
			found10zakq[0] = true
			z.SchemaId, err = dc.ReadUint64()
			if err != nil {
				panic(err)
			}
		case "IntToPackageTable":
			found10zakq[1] = true
			var zytz uint32
			zytz, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.IntToPackageTable == nil && zytz > 0 {
				z.IntToPackageTable = make(map[int64]string, zytz)
			} else if len(z.IntToPackageTable) > 0 {
				for key, _ := range z.IntToPackageTable {
					delete(z.IntToPackageTable, key)
				}
			}
			for zytz > 0 {
				zytz--
				var zrrn int64
				var ztwr string
				zrrn, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				ztwr, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.IntToPackageTable[zrrn] = ztwr
			}
		case "PkgPath":
			found10zakq[2] = true
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
	if nextMiss10zakq != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder10zakq = []string{"SchemaId", "IntToPackageTable", "PkgPath"}

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
	for zrrn, ztwr := range z.IntToPackageTable {
		err = en.WriteInt64(zrrn)
		if err != nil {
			panic(err)
		}
		err = en.WriteString(ztwr)
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
	for zrrn, ztwr := range z.IntToPackageTable {
		o = msgp.AppendInt64(o, zrrn)
		o = msgp.AppendString(o, ztwr)
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
	const maxFields11zzks = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields11zzks uint32
	if !nbs.AlwaysNil {
		totalEncodedFields11zzks, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft11zzks := totalEncodedFields11zzks
	missingFieldsLeft11zzks := maxFields11zzks - totalEncodedFields11zzks

	var nextMiss11zzks int32 = -1
	var found11zzks [maxFields11zzks]bool
	var curField11zzks string

doneWithStruct11zzks:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft11zzks > 0 || missingFieldsLeft11zzks > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zzks, missingFieldsLeft11zzks, msgp.ShowFound(found11zzks[:]), unmarshalMsgFieldOrder11zzks)
		if encodedFieldsLeft11zzks > 0 {
			encodedFieldsLeft11zzks--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField11zzks = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss11zzks < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss11zzks = 0
			}
			for nextMiss11zzks < maxFields11zzks && found11zzks[nextMiss11zzks] {
				nextMiss11zzks++
			}
			if nextMiss11zzks == maxFields11zzks {
				// filled all the empty fields!
				break doneWithStruct11zzks
			}
			missingFieldsLeft11zzks--
			curField11zzks = unmarshalMsgFieldOrder11zzks[nextMiss11zzks]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField11zzks)
		switch curField11zzks {
		// -- templateUnmarshalMsg ends here --

		case "SchemaId":
			found11zzks[0] = true
			z.SchemaId, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "IntToPackageTable":
			found11zzks[1] = true
			if nbs.AlwaysNil {
				if len(z.IntToPackageTable) > 0 {
					for key, _ := range z.IntToPackageTable {
						delete(z.IntToPackageTable, key)
					}
				}

			} else {

				var zqgj uint32
				zqgj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.IntToPackageTable == nil && zqgj > 0 {
					z.IntToPackageTable = make(map[int64]string, zqgj)
				} else if len(z.IntToPackageTable) > 0 {
					for key, _ := range z.IntToPackageTable {
						delete(z.IntToPackageTable, key)
					}
				}
				for zqgj > 0 {
					var zrrn int64
					var ztwr string
					zqgj--
					zrrn, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					ztwr, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.IntToPackageTable[zrrn] = ztwr
				}
			}
		case "PkgPath":
			found11zzks[2] = true
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
	if nextMiss11zzks != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder11zzks = []string{"SchemaId", "IntToPackageTable", "PkgPath"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Uint64Size + 18 + msgp.MapHeaderSize
	if z.IntToPackageTable != nil {
		for zrrn, ztwr := range z.IntToPackageTable {
			_ = ztwr
			_ = zrrn
			s += msgp.Int64Size + msgp.StringPrefixSize + len(ztwr)
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
	const maxFields12zrkv = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zrkv uint32
	totalEncodedFields12zrkv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zrkv := totalEncodedFields12zrkv
	missingFieldsLeft12zrkv := maxFields12zrkv - totalEncodedFields12zrkv

	var nextMiss12zrkv int32 = -1
	var found12zrkv [maxFields12zrkv]bool
	var curField12zrkv string

doneWithStruct12zrkv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zrkv > 0 || missingFieldsLeft12zrkv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zrkv, missingFieldsLeft12zrkv, msgp.ShowFound(found12zrkv[:]), decodeMsgFieldOrder12zrkv)
		if encodedFieldsLeft12zrkv > 0 {
			encodedFieldsLeft12zrkv--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zrkv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zrkv < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zrkv = 0
			}
			for nextMiss12zrkv < maxFields12zrkv && found12zrkv[nextMiss12zrkv] {
				nextMiss12zrkv++
			}
			if nextMiss12zrkv == maxFields12zrkv {
				// filled all the empty fields!
				break doneWithStruct12zrkv
			}
			missingFieldsLeft12zrkv--
			curField12zrkv = decodeMsgFieldOrder12zrkv[nextMiss12zrkv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zrkv)
		switch curField12zrkv {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found12zrkv[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found12zrkv[1] = true
			var zdem uint32
			zdem, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zdem) {
				z.Fld = (z.Fld)[:zdem]
			} else {
				z.Fld = make([]FieldT, zdem)
			}
			for zxtp := range z.Fld {
				err = z.Fld[zxtp].DecodeMsg(dc)
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
	if nextMiss12zrkv != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder12zrkv = []string{"StructName", "Fld"}

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
	for zxtp := range z.Fld {
		err = z.Fld[zxtp].EncodeMsg(en)
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
	for zxtp := range z.Fld {
		o, err = z.Fld[zxtp].MarshalMsg(o)
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
	const maxFields13zcbm = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields13zcbm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zcbm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft13zcbm := totalEncodedFields13zcbm
	missingFieldsLeft13zcbm := maxFields13zcbm - totalEncodedFields13zcbm

	var nextMiss13zcbm int32 = -1
	var found13zcbm [maxFields13zcbm]bool
	var curField13zcbm string

doneWithStruct13zcbm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zcbm > 0 || missingFieldsLeft13zcbm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zcbm, missingFieldsLeft13zcbm, msgp.ShowFound(found13zcbm[:]), unmarshalMsgFieldOrder13zcbm)
		if encodedFieldsLeft13zcbm > 0 {
			encodedFieldsLeft13zcbm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField13zcbm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss13zcbm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zcbm = 0
			}
			for nextMiss13zcbm < maxFields13zcbm && found13zcbm[nextMiss13zcbm] {
				nextMiss13zcbm++
			}
			if nextMiss13zcbm == maxFields13zcbm {
				// filled all the empty fields!
				break doneWithStruct13zcbm
			}
			missingFieldsLeft13zcbm--
			curField13zcbm = unmarshalMsgFieldOrder13zcbm[nextMiss13zcbm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zcbm)
		switch curField13zcbm {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found13zcbm[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found13zcbm[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zqyy uint32
				zqyy, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zqyy) {
					z.Fld = (z.Fld)[:zqyy]
				} else {
					z.Fld = make([]FieldT, zqyy)
				}
				for zxtp := range z.Fld {
					bts, err = z.Fld[zxtp].UnmarshalMsg(bts)
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
	if nextMiss13zcbm != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder13zcbm = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zxtp := range z.Fld {
		s += z.Fld[zxtp].Msgsize()
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
		var zugw int64
		zugw, err = dc.ReadInt64()
		(*z) = StructTypeId(zugw)
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
		var zprz int64
		zprz, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zprz)
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
		var zhvf byte
		zhvf, err = dc.ReadByte()
		(*z) = ZKind(zhvf)
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
		var zuty byte
		zuty, bts, err = nbs.ReadByteBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = ZKind(zuty)
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
	const maxFields14zzfs = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields14zzfs uint32
	totalEncodedFields14zzfs, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zzfs := totalEncodedFields14zzfs
	missingFieldsLeft14zzfs := maxFields14zzfs - totalEncodedFields14zzfs

	var nextMiss14zzfs int32 = -1
	var found14zzfs [maxFields14zzfs]bool
	var curField14zzfs string

doneWithStruct14zzfs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zzfs > 0 || missingFieldsLeft14zzfs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zzfs, missingFieldsLeft14zzfs, msgp.ShowFound(found14zzfs[:]), decodeMsgFieldOrder14zzfs)
		if encodedFieldsLeft14zzfs > 0 {
			encodedFieldsLeft14zzfs--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField14zzfs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zzfs < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zzfs = 0
			}
			for nextMiss14zzfs < maxFields14zzfs && found14zzfs[nextMiss14zzfs] {
				nextMiss14zzfs++
			}
			if nextMiss14zzfs == maxFields14zzfs {
				// filled all the empty fields!
				break doneWithStruct14zzfs
			}
			missingFieldsLeft14zzfs--
			curField14zzfs = decodeMsgFieldOrder14zzfs[nextMiss14zzfs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zzfs)
		switch curField14zzfs {
		// -- templateDecodeMsg ends here --

		case "Id":
			found14zzfs[0] = true
			{
				var zltm int64
				zltm, err = dc.ReadInt64()
				z.Id = StructTypeId(zltm)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found14zzfs[1] = true
			var zlrj uint32
			zlrj, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zlrj > 0 {
				z.Da = make(map[int64]msgp.Raw, zlrj)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zlrj > 0 {
				zlrj--
				var zknc int64
				var zkyc msgp.Raw
				zknc, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zkyc.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zknc] = zkyc
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zzfs != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZStructPacket
var decodeMsgFieldOrder14zzfs = []string{"Id", "Da"}

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
	for zknc, zkyc := range z.Da {
		err = en.WriteInt64(zknc)
		if err != nil {
			panic(err)
		}
		err = zkyc.EncodeMsg(en)
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
	for zknc, zkyc := range z.Da {
		o = msgp.AppendInt64(o, zknc)
		o, err = zkyc.MarshalMsg(o)
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
	const maxFields15zxhe = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields15zxhe uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zxhe, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft15zxhe := totalEncodedFields15zxhe
	missingFieldsLeft15zxhe := maxFields15zxhe - totalEncodedFields15zxhe

	var nextMiss15zxhe int32 = -1
	var found15zxhe [maxFields15zxhe]bool
	var curField15zxhe string

doneWithStruct15zxhe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zxhe > 0 || missingFieldsLeft15zxhe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zxhe, missingFieldsLeft15zxhe, msgp.ShowFound(found15zxhe[:]), unmarshalMsgFieldOrder15zxhe)
		if encodedFieldsLeft15zxhe > 0 {
			encodedFieldsLeft15zxhe--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField15zxhe = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss15zxhe < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zxhe = 0
			}
			for nextMiss15zxhe < maxFields15zxhe && found15zxhe[nextMiss15zxhe] {
				nextMiss15zxhe++
			}
			if nextMiss15zxhe == maxFields15zxhe {
				// filled all the empty fields!
				break doneWithStruct15zxhe
			}
			missingFieldsLeft15zxhe--
			curField15zxhe = unmarshalMsgFieldOrder15zxhe[nextMiss15zxhe]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zxhe)
		switch curField15zxhe {
		// -- templateUnmarshalMsg ends here --

		case "Id":
			found15zxhe[0] = true
			{
				var zibj int64
				zibj, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zibj)
			}
		case "Da":
			found15zxhe[1] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zwjp uint32
				zwjp, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zwjp > 0 {
					z.Da = make(map[int64]msgp.Raw, zwjp)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zwjp > 0 {
					var zknc int64
					var zkyc msgp.Raw
					zwjp--
					zknc, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zkyc.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zknc] = zkyc
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss15zxhe != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZStructPacket
var unmarshalMsgFieldOrder15zxhe = []string{"Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZStructPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zknc, zkyc := range z.Da {
			_ = zkyc
			_ = zknc
			s += msgp.Int64Size + zkyc.Msgsize()
		}
	}
	return
}
