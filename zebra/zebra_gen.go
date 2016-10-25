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
	const maxFields0zxpb = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zxpb uint32
	totalEncodedFields0zxpb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zxpb := totalEncodedFields0zxpb
	missingFieldsLeft0zxpb := maxFields0zxpb - totalEncodedFields0zxpb

	var nextMiss0zxpb int32 = -1
	var found0zxpb [maxFields0zxpb]bool
	var curField0zxpb string

doneWithStruct0zxpb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zxpb > 0 || missingFieldsLeft0zxpb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zxpb, missingFieldsLeft0zxpb, msgp.ShowFound(found0zxpb[:]), decodeMsgFieldOrder0zxpb)
		if encodedFieldsLeft0zxpb > 0 {
			encodedFieldsLeft0zxpb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zxpb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zxpb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zxpb = 0
			}
			for nextMiss0zxpb < maxFields0zxpb && found0zxpb[nextMiss0zxpb] {
				nextMiss0zxpb++
			}
			if nextMiss0zxpb == maxFields0zxpb {
				// filled all the empty fields!
				break doneWithStruct0zxpb
			}
			missingFieldsLeft0zxpb--
			curField0zxpb = decodeMsgFieldOrder0zxpb[nextMiss0zxpb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zxpb)
		switch curField0zxpb {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zxpb[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zxpb[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Zknd":
			found0zxpb[2] = true
			var zdyn uint32
			zdyn, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Zknd) >= int(zdyn) {
				z.Zknd = (z.Zknd)[:zdyn]
			} else {
				z.Zknd = make([]ZKind, zdyn)
			}
			for zpri := range z.Zknd {
				{
					var zfvv byte
					zfvv, err = dc.ReadByte()
					z.Zknd[zpri] = ZKind(zfvv)
				}
				if err != nil {
					panic(err)
				}
			}
		case "Varg":
			found0zxpb[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zxpb[4] = true
			var ztmg uint32
			ztmg, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && ztmg > 0 {
				z.Tag = make(map[string]string, ztmg)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for ztmg > 0 {
				ztmg--
				var zajo string
				var zbtr string
				zajo, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zbtr, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zajo] = zbtr
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zxpb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zxpb = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

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
	var empty_zdnb [5]bool
	fieldsInUse_ztpa := z.fieldsNotEmpty(empty_zdnb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_ztpa)
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
	for zpri := range z.Zknd {
		err = en.WriteByte(byte(z.Zknd[zpri]))
		if err != nil {
			panic(err)
		}
	}
	if !empty_zdnb[3] {
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

	if !empty_zdnb[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zajo, zbtr := range z.Tag {
			err = en.WriteString(zajo)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zbtr)
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
	for zpri := range z.Zknd {
		o = msgp.AppendByte(o, byte(z.Zknd[zpri]))
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
		for zajo, zbtr := range z.Tag {
			o = msgp.AppendString(o, zajo)
			o = msgp.AppendString(o, zbtr)
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
	const maxFields1zjrp = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zjrp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zjrp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zjrp := totalEncodedFields1zjrp
	missingFieldsLeft1zjrp := maxFields1zjrp - totalEncodedFields1zjrp

	var nextMiss1zjrp int32 = -1
	var found1zjrp [maxFields1zjrp]bool
	var curField1zjrp string

doneWithStruct1zjrp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zjrp > 0 || missingFieldsLeft1zjrp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zjrp, missingFieldsLeft1zjrp, msgp.ShowFound(found1zjrp[:]), unmarshalMsgFieldOrder1zjrp)
		if encodedFieldsLeft1zjrp > 0 {
			encodedFieldsLeft1zjrp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zjrp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zjrp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zjrp = 0
			}
			for nextMiss1zjrp < maxFields1zjrp && found1zjrp[nextMiss1zjrp] {
				nextMiss1zjrp++
			}
			if nextMiss1zjrp == maxFields1zjrp {
				// filled all the empty fields!
				break doneWithStruct1zjrp
			}
			missingFieldsLeft1zjrp--
			curField1zjrp = unmarshalMsgFieldOrder1zjrp[nextMiss1zjrp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zjrp)
		switch curField1zjrp {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zjrp[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zjrp[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Zknd":
			found1zjrp[2] = true
			if nbs.AlwaysNil {
				(z.Zknd) = (z.Zknd)[:0]
			} else {

				var zmmq uint32
				zmmq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Zknd) >= int(zmmq) {
					z.Zknd = (z.Zknd)[:zmmq]
				} else {
					z.Zknd = make([]ZKind, zmmq)
				}
				for zpri := range z.Zknd {
					{
						var zvdf byte
						zvdf, bts, err = nbs.ReadByteBytes(bts)

						if err != nil {
							panic(err)
						}
						z.Zknd[zpri] = ZKind(zvdf)
					}
				}
			}
		case "Varg":
			found1zjrp[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zjrp[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zbec uint32
				zbec, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zbec > 0 {
					z.Tag = make(map[string]string, zbec)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zbec > 0 {
					var zajo string
					var zbtr string
					zbec--
					zajo, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zbtr, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zajo] = zbtr
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zjrp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zjrp = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize + (len(z.Zknd) * (msgp.ByteSize)) + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zajo, zbtr := range z.Tag {
			_ = zbtr
			_ = zajo
			s += msgp.StringPrefixSize + len(zajo) + msgp.StringPrefixSize + len(zbtr)
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
		var zngc int64
		zngc, err = dc.ReadInt64()
		(*z) = InterfaceId(zngc)
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
		var zrlk int64
		zrlk, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = InterfaceId(zrlk)
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
	const maxFields2zmuo = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zmuo uint32
	totalEncodedFields2zmuo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zmuo := totalEncodedFields2zmuo
	missingFieldsLeft2zmuo := maxFields2zmuo - totalEncodedFields2zmuo

	var nextMiss2zmuo int32 = -1
	var found2zmuo [maxFields2zmuo]bool
	var curField2zmuo string

doneWithStruct2zmuo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zmuo > 0 || missingFieldsLeft2zmuo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zmuo, missingFieldsLeft2zmuo, msgp.ShowFound(found2zmuo[:]), decodeMsgFieldOrder2zmuo)
		if encodedFieldsLeft2zmuo > 0 {
			encodedFieldsLeft2zmuo--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zmuo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zmuo < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zmuo = 0
			}
			for nextMiss2zmuo < maxFields2zmuo && found2zmuo[nextMiss2zmuo] {
				nextMiss2zmuo++
			}
			if nextMiss2zmuo == maxFields2zmuo {
				// filled all the empty fields!
				break doneWithStruct2zmuo
			}
			missingFieldsLeft2zmuo--
			curField2zmuo = decodeMsgFieldOrder2zmuo[nextMiss2zmuo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zmuo)
		switch curField2zmuo {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found2zmuo[0] = true
			var zntf uint32
			zntf, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zntf) {
				z.Methods = (z.Methods)[:zntf]
			} else {
				z.Methods = make([]MethodT, zntf)
			}
			for zubs := range z.Methods {
				err = z.Methods[zubs].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found2zmuo[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found2zmuo[2] = true
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
	if nextMiss2zmuo != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder2zmuo = []string{"Methods", "Deprecated", "Comment"}

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
	for zubs := range z.Methods {
		err = z.Methods[zubs].EncodeMsg(en)
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
	for zubs := range z.Methods {
		o, err = z.Methods[zubs].MarshalMsg(o)
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
	const maxFields3zwtf = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zwtf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zwtf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zwtf := totalEncodedFields3zwtf
	missingFieldsLeft3zwtf := maxFields3zwtf - totalEncodedFields3zwtf

	var nextMiss3zwtf int32 = -1
	var found3zwtf [maxFields3zwtf]bool
	var curField3zwtf string

doneWithStruct3zwtf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zwtf > 0 || missingFieldsLeft3zwtf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zwtf, missingFieldsLeft3zwtf, msgp.ShowFound(found3zwtf[:]), unmarshalMsgFieldOrder3zwtf)
		if encodedFieldsLeft3zwtf > 0 {
			encodedFieldsLeft3zwtf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zwtf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zwtf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zwtf = 0
			}
			for nextMiss3zwtf < maxFields3zwtf && found3zwtf[nextMiss3zwtf] {
				nextMiss3zwtf++
			}
			if nextMiss3zwtf == maxFields3zwtf {
				// filled all the empty fields!
				break doneWithStruct3zwtf
			}
			missingFieldsLeft3zwtf--
			curField3zwtf = unmarshalMsgFieldOrder3zwtf[nextMiss3zwtf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zwtf)
		switch curField3zwtf {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found3zwtf[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zcue uint32
				zcue, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zcue) {
					z.Methods = (z.Methods)[:zcue]
				} else {
					z.Methods = make([]MethodT, zcue)
				}
				for zubs := range z.Methods {
					bts, err = z.Methods[zubs].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found3zwtf[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found3zwtf[2] = true
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
	if nextMiss3zwtf != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder3zwtf = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zubs := range z.Methods {
		s += z.Methods[zubs].Msgsize()
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
	const maxFields4zppm = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zppm uint32
	totalEncodedFields4zppm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zppm := totalEncodedFields4zppm
	missingFieldsLeft4zppm := maxFields4zppm - totalEncodedFields4zppm

	var nextMiss4zppm int32 = -1
	var found4zppm [maxFields4zppm]bool
	var curField4zppm string

doneWithStruct4zppm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zppm > 0 || missingFieldsLeft4zppm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zppm, missingFieldsLeft4zppm, msgp.ShowFound(found4zppm[:]), decodeMsgFieldOrder4zppm)
		if encodedFieldsLeft4zppm > 0 {
			encodedFieldsLeft4zppm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zppm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zppm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zppm = 0
			}
			for nextMiss4zppm < maxFields4zppm && found4zppm[nextMiss4zppm] {
				nextMiss4zppm++
			}
			if nextMiss4zppm == maxFields4zppm {
				// filled all the empty fields!
				break doneWithStruct4zppm
			}
			missingFieldsLeft4zppm--
			curField4zppm = decodeMsgFieldOrder4zppm[nextMiss4zppm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zppm)
		switch curField4zppm {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found4zppm[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found4zppm[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found4zppm[2] = true
			const maxFields5zxro = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields5zxro uint32
			totalEncodedFields5zxro, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft5zxro := totalEncodedFields5zxro
			missingFieldsLeft5zxro := maxFields5zxro - totalEncodedFields5zxro

			var nextMiss5zxro int32 = -1
			var found5zxro [maxFields5zxro]bool
			var curField5zxro string

		doneWithStruct5zxro:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft5zxro > 0 || missingFieldsLeft5zxro > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zxro, missingFieldsLeft5zxro, msgp.ShowFound(found5zxro[:]), decodeMsgFieldOrder5zxro)
				if encodedFieldsLeft5zxro > 0 {
					encodedFieldsLeft5zxro--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField5zxro = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss5zxro < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss5zxro = 0
					}
					for nextMiss5zxro < maxFields5zxro && found5zxro[nextMiss5zxro] {
						nextMiss5zxro++
					}
					if nextMiss5zxro == maxFields5zxro {
						// filled all the empty fields!
						break doneWithStruct5zxro
					}
					missingFieldsLeft5zxro--
					curField5zxro = decodeMsgFieldOrder5zxro[nextMiss5zxro]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField5zxro)
				switch curField5zxro {
				// -- templateDecodeMsg ends here --

				case "StructName":
					found5zxro[0] = true
					z.Inputs.StructName, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found5zxro[1] = true
					var zoyh uint32
					zoyh, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zoyh) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zoyh]
					} else {
						z.Inputs.Fld = make([]FieldT, zoyh)
					}
					for zqxp := range z.Inputs.Fld {
						err = z.Inputs.Fld[zqxp].DecodeMsg(dc)
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
			if nextMiss5zxro != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found4zppm[3] = true
			const maxFields6zmum = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields6zmum uint32
			totalEncodedFields6zmum, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft6zmum := totalEncodedFields6zmum
			missingFieldsLeft6zmum := maxFields6zmum - totalEncodedFields6zmum

			var nextMiss6zmum int32 = -1
			var found6zmum [maxFields6zmum]bool
			var curField6zmum string

		doneWithStruct6zmum:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft6zmum > 0 || missingFieldsLeft6zmum > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zmum, missingFieldsLeft6zmum, msgp.ShowFound(found6zmum[:]), decodeMsgFieldOrder6zmum)
				if encodedFieldsLeft6zmum > 0 {
					encodedFieldsLeft6zmum--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField6zmum = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss6zmum < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss6zmum = 0
					}
					for nextMiss6zmum < maxFields6zmum && found6zmum[nextMiss6zmum] {
						nextMiss6zmum++
					}
					if nextMiss6zmum == maxFields6zmum {
						// filled all the empty fields!
						break doneWithStruct6zmum
					}
					missingFieldsLeft6zmum--
					curField6zmum = decodeMsgFieldOrder6zmum[nextMiss6zmum]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField6zmum)
				switch curField6zmum {
				// -- templateDecodeMsg ends here --

				case "StructName":
					found6zmum[0] = true
					z.Returns.StructName, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found6zmum[1] = true
					var zqqy uint32
					zqqy, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zqqy) {
						z.Returns.Fld = (z.Returns.Fld)[:zqqy]
					} else {
						z.Returns.Fld = make([]FieldT, zqqy)
					}
					for zjtz := range z.Returns.Fld {
						err = z.Returns.Fld[zjtz].DecodeMsg(dc)
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
			if nextMiss6zmum != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found4zppm[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zppm[5] = true
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
	if nextMiss4zppm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder4zppm = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder5zxro = []string{"StructName", "Fld"}

// fields of StructT
var decodeMsgFieldOrder6zmum = []string{"StructName", "Fld"}

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
	for zqxp := range z.Inputs.Fld {
		err = z.Inputs.Fld[zqxp].EncodeMsg(en)
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
	for zjtz := range z.Returns.Fld {
		err = z.Returns.Fld[zjtz].EncodeMsg(en)
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
	for zqxp := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zqxp].MarshalMsg(o)
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
	for zjtz := range z.Returns.Fld {
		o, err = z.Returns.Fld[zjtz].MarshalMsg(o)
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
	const maxFields7zuub = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zuub uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zuub, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft7zuub := totalEncodedFields7zuub
	missingFieldsLeft7zuub := maxFields7zuub - totalEncodedFields7zuub

	var nextMiss7zuub int32 = -1
	var found7zuub [maxFields7zuub]bool
	var curField7zuub string

doneWithStruct7zuub:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zuub > 0 || missingFieldsLeft7zuub > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zuub, missingFieldsLeft7zuub, msgp.ShowFound(found7zuub[:]), unmarshalMsgFieldOrder7zuub)
		if encodedFieldsLeft7zuub > 0 {
			encodedFieldsLeft7zuub--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField7zuub = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zuub < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zuub = 0
			}
			for nextMiss7zuub < maxFields7zuub && found7zuub[nextMiss7zuub] {
				nextMiss7zuub++
			}
			if nextMiss7zuub == maxFields7zuub {
				// filled all the empty fields!
				break doneWithStruct7zuub
			}
			missingFieldsLeft7zuub--
			curField7zuub = unmarshalMsgFieldOrder7zuub[nextMiss7zuub]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zuub)
		switch curField7zuub {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found7zuub[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found7zuub[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found7zuub[2] = true
			const maxFields8zdpj = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields8zdpj uint32
			if !nbs.AlwaysNil {
				totalEncodedFields8zdpj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft8zdpj := totalEncodedFields8zdpj
			missingFieldsLeft8zdpj := maxFields8zdpj - totalEncodedFields8zdpj

			var nextMiss8zdpj int32 = -1
			var found8zdpj [maxFields8zdpj]bool
			var curField8zdpj string

		doneWithStruct8zdpj:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zdpj > 0 || missingFieldsLeft8zdpj > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zdpj, missingFieldsLeft8zdpj, msgp.ShowFound(found8zdpj[:]), unmarshalMsgFieldOrder8zdpj)
				if encodedFieldsLeft8zdpj > 0 {
					encodedFieldsLeft8zdpj--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField8zdpj = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zdpj < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss8zdpj = 0
					}
					for nextMiss8zdpj < maxFields8zdpj && found8zdpj[nextMiss8zdpj] {
						nextMiss8zdpj++
					}
					if nextMiss8zdpj == maxFields8zdpj {
						// filled all the empty fields!
						break doneWithStruct8zdpj
					}
					missingFieldsLeft8zdpj--
					curField8zdpj = unmarshalMsgFieldOrder8zdpj[nextMiss8zdpj]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zdpj)
				switch curField8zdpj {
				// -- templateUnmarshalMsg ends here --

				case "StructName":
					found8zdpj[0] = true
					z.Inputs.StructName, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zdpj[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zckd uint32
						zckd, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zckd) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zckd]
						} else {
							z.Inputs.Fld = make([]FieldT, zckd)
						}
						for zqxp := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zqxp].UnmarshalMsg(bts)
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
			if nextMiss8zdpj != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found7zuub[3] = true
			const maxFields9zaie = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields9zaie uint32
			if !nbs.AlwaysNil {
				totalEncodedFields9zaie, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft9zaie := totalEncodedFields9zaie
			missingFieldsLeft9zaie := maxFields9zaie - totalEncodedFields9zaie

			var nextMiss9zaie int32 = -1
			var found9zaie [maxFields9zaie]bool
			var curField9zaie string

		doneWithStruct9zaie:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft9zaie > 0 || missingFieldsLeft9zaie > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zaie, missingFieldsLeft9zaie, msgp.ShowFound(found9zaie[:]), unmarshalMsgFieldOrder9zaie)
				if encodedFieldsLeft9zaie > 0 {
					encodedFieldsLeft9zaie--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField9zaie = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss9zaie < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss9zaie = 0
					}
					for nextMiss9zaie < maxFields9zaie && found9zaie[nextMiss9zaie] {
						nextMiss9zaie++
					}
					if nextMiss9zaie == maxFields9zaie {
						// filled all the empty fields!
						break doneWithStruct9zaie
					}
					missingFieldsLeft9zaie--
					curField9zaie = unmarshalMsgFieldOrder9zaie[nextMiss9zaie]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField9zaie)
				switch curField9zaie {
				// -- templateUnmarshalMsg ends here --

				case "StructName":
					found9zaie[0] = true
					z.Returns.StructName, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found9zaie[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zuwx uint32
						zuwx, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zuwx) {
							z.Returns.Fld = (z.Returns.Fld)[:zuwx]
						} else {
							z.Returns.Fld = make([]FieldT, zuwx)
						}
						for zjtz := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zjtz].UnmarshalMsg(bts)
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
			if nextMiss9zaie != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found7zuub[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found7zuub[5] = true
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
	if nextMiss7zuub != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder7zuub = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder8zdpj = []string{"StructName", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder9zaie = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 11 + msgp.StringPrefixSize + len(z.Inputs.StructName) + 4 + msgp.ArrayHeaderSize
	for zqxp := range z.Inputs.Fld {
		s += z.Inputs.Fld[zqxp].Msgsize()
	}
	s += 8 + 1 + 11 + msgp.StringPrefixSize + len(z.Returns.StructName) + 4 + msgp.ArrayHeaderSize
	for zjtz := range z.Returns.Fld {
		s += z.Returns.Fld[zjtz].Msgsize()
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
	const maxFields10zewn = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields10zewn uint32
	totalEncodedFields10zewn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft10zewn := totalEncodedFields10zewn
	missingFieldsLeft10zewn := maxFields10zewn - totalEncodedFields10zewn

	var nextMiss10zewn int32 = -1
	var found10zewn [maxFields10zewn]bool
	var curField10zewn string

doneWithStruct10zewn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zewn > 0 || missingFieldsLeft10zewn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zewn, missingFieldsLeft10zewn, msgp.ShowFound(found10zewn[:]), decodeMsgFieldOrder10zewn)
		if encodedFieldsLeft10zewn > 0 {
			encodedFieldsLeft10zewn--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField10zewn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss10zewn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss10zewn = 0
			}
			for nextMiss10zewn < maxFields10zewn && found10zewn[nextMiss10zewn] {
				nextMiss10zewn++
			}
			if nextMiss10zewn == maxFields10zewn {
				// filled all the empty fields!
				break doneWithStruct10zewn
			}
			missingFieldsLeft10zewn--
			curField10zewn = decodeMsgFieldOrder10zewn[nextMiss10zewn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zewn)
		switch curField10zewn {
		// -- templateDecodeMsg ends here --

		case "SchemaId":
			found10zewn[0] = true
			z.SchemaId, err = dc.ReadUint64()
			if err != nil {
				panic(err)
			}
		case "IntToPackageTable":
			found10zewn[1] = true
			var zyyl uint32
			zyyl, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.IntToPackageTable == nil && zyyl > 0 {
				z.IntToPackageTable = make(map[int64]string, zyyl)
			} else if len(z.IntToPackageTable) > 0 {
				for key, _ := range z.IntToPackageTable {
					delete(z.IntToPackageTable, key)
				}
			}
			for zyyl > 0 {
				zyyl--
				var zzrg int64
				var zwco string
				zzrg, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				zwco, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.IntToPackageTable[zzrg] = zwco
			}
		case "PkgPath":
			found10zewn[2] = true
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
	if nextMiss10zewn != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder10zewn = []string{"SchemaId", "IntToPackageTable", "PkgPath"}

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
	for zzrg, zwco := range z.IntToPackageTable {
		err = en.WriteInt64(zzrg)
		if err != nil {
			panic(err)
		}
		err = en.WriteString(zwco)
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
	for zzrg, zwco := range z.IntToPackageTable {
		o = msgp.AppendInt64(o, zzrg)
		o = msgp.AppendString(o, zwco)
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
	const maxFields11zyes = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields11zyes uint32
	if !nbs.AlwaysNil {
		totalEncodedFields11zyes, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft11zyes := totalEncodedFields11zyes
	missingFieldsLeft11zyes := maxFields11zyes - totalEncodedFields11zyes

	var nextMiss11zyes int32 = -1
	var found11zyes [maxFields11zyes]bool
	var curField11zyes string

doneWithStruct11zyes:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft11zyes > 0 || missingFieldsLeft11zyes > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zyes, missingFieldsLeft11zyes, msgp.ShowFound(found11zyes[:]), unmarshalMsgFieldOrder11zyes)
		if encodedFieldsLeft11zyes > 0 {
			encodedFieldsLeft11zyes--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField11zyes = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss11zyes < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss11zyes = 0
			}
			for nextMiss11zyes < maxFields11zyes && found11zyes[nextMiss11zyes] {
				nextMiss11zyes++
			}
			if nextMiss11zyes == maxFields11zyes {
				// filled all the empty fields!
				break doneWithStruct11zyes
			}
			missingFieldsLeft11zyes--
			curField11zyes = unmarshalMsgFieldOrder11zyes[nextMiss11zyes]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField11zyes)
		switch curField11zyes {
		// -- templateUnmarshalMsg ends here --

		case "SchemaId":
			found11zyes[0] = true
			z.SchemaId, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "IntToPackageTable":
			found11zyes[1] = true
			if nbs.AlwaysNil {
				if len(z.IntToPackageTable) > 0 {
					for key, _ := range z.IntToPackageTable {
						delete(z.IntToPackageTable, key)
					}
				}

			} else {

				var zsey uint32
				zsey, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.IntToPackageTable == nil && zsey > 0 {
					z.IntToPackageTable = make(map[int64]string, zsey)
				} else if len(z.IntToPackageTable) > 0 {
					for key, _ := range z.IntToPackageTable {
						delete(z.IntToPackageTable, key)
					}
				}
				for zsey > 0 {
					var zzrg int64
					var zwco string
					zsey--
					zzrg, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					zwco, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.IntToPackageTable[zzrg] = zwco
				}
			}
		case "PkgPath":
			found11zyes[2] = true
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
	if nextMiss11zyes != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder11zyes = []string{"SchemaId", "IntToPackageTable", "PkgPath"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Uint64Size + 18 + msgp.MapHeaderSize
	if z.IntToPackageTable != nil {
		for zzrg, zwco := range z.IntToPackageTable {
			_ = zwco
			_ = zzrg
			s += msgp.Int64Size + msgp.StringPrefixSize + len(zwco)
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
	const maxFields12zfcn = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zfcn uint32
	totalEncodedFields12zfcn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zfcn := totalEncodedFields12zfcn
	missingFieldsLeft12zfcn := maxFields12zfcn - totalEncodedFields12zfcn

	var nextMiss12zfcn int32 = -1
	var found12zfcn [maxFields12zfcn]bool
	var curField12zfcn string

doneWithStruct12zfcn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zfcn > 0 || missingFieldsLeft12zfcn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zfcn, missingFieldsLeft12zfcn, msgp.ShowFound(found12zfcn[:]), decodeMsgFieldOrder12zfcn)
		if encodedFieldsLeft12zfcn > 0 {
			encodedFieldsLeft12zfcn--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zfcn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zfcn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zfcn = 0
			}
			for nextMiss12zfcn < maxFields12zfcn && found12zfcn[nextMiss12zfcn] {
				nextMiss12zfcn++
			}
			if nextMiss12zfcn == maxFields12zfcn {
				// filled all the empty fields!
				break doneWithStruct12zfcn
			}
			missingFieldsLeft12zfcn--
			curField12zfcn = decodeMsgFieldOrder12zfcn[nextMiss12zfcn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zfcn)
		switch curField12zfcn {
		// -- templateDecodeMsg ends here --

		case "StructName":
			found12zfcn[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found12zfcn[1] = true
			var zmco uint32
			zmco, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zmco) {
				z.Fld = (z.Fld)[:zmco]
			} else {
				z.Fld = make([]FieldT, zmco)
			}
			for zkhr := range z.Fld {
				err = z.Fld[zkhr].DecodeMsg(dc)
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
	if nextMiss12zfcn != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder12zfcn = []string{"StructName", "Fld"}

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
	for zkhr := range z.Fld {
		err = z.Fld[zkhr].EncodeMsg(en)
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
	for zkhr := range z.Fld {
		o, err = z.Fld[zkhr].MarshalMsg(o)
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
	const maxFields13zqdq = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields13zqdq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zqdq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft13zqdq := totalEncodedFields13zqdq
	missingFieldsLeft13zqdq := maxFields13zqdq - totalEncodedFields13zqdq

	var nextMiss13zqdq int32 = -1
	var found13zqdq [maxFields13zqdq]bool
	var curField13zqdq string

doneWithStruct13zqdq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zqdq > 0 || missingFieldsLeft13zqdq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zqdq, missingFieldsLeft13zqdq, msgp.ShowFound(found13zqdq[:]), unmarshalMsgFieldOrder13zqdq)
		if encodedFieldsLeft13zqdq > 0 {
			encodedFieldsLeft13zqdq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField13zqdq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss13zqdq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zqdq = 0
			}
			for nextMiss13zqdq < maxFields13zqdq && found13zqdq[nextMiss13zqdq] {
				nextMiss13zqdq++
			}
			if nextMiss13zqdq == maxFields13zqdq {
				// filled all the empty fields!
				break doneWithStruct13zqdq
			}
			missingFieldsLeft13zqdq--
			curField13zqdq = unmarshalMsgFieldOrder13zqdq[nextMiss13zqdq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zqdq)
		switch curField13zqdq {
		// -- templateUnmarshalMsg ends here --

		case "StructName":
			found13zqdq[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found13zqdq[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zsei uint32
				zsei, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zsei) {
					z.Fld = (z.Fld)[:zsei]
				} else {
					z.Fld = make([]FieldT, zsei)
				}
				for zkhr := range z.Fld {
					bts, err = z.Fld[zkhr].UnmarshalMsg(bts)
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
	if nextMiss13zqdq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder13zqdq = []string{"StructName", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zkhr := range z.Fld {
		s += z.Fld[zkhr].Msgsize()
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
		var zdqn int64
		zdqn, err = dc.ReadInt64()
		(*z) = StructTypeId(zdqn)
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
		var zjme int64
		zjme, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zjme)
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
		var zvir byte
		zvir, err = dc.ReadByte()
		(*z) = ZKind(zvir)
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
		var zdfn byte
		zdfn, bts, err = nbs.ReadByteBytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = ZKind(zdfn)
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
	const maxFields14zvoe = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields14zvoe uint32
	totalEncodedFields14zvoe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zvoe := totalEncodedFields14zvoe
	missingFieldsLeft14zvoe := maxFields14zvoe - totalEncodedFields14zvoe

	var nextMiss14zvoe int32 = -1
	var found14zvoe [maxFields14zvoe]bool
	var curField14zvoe string

doneWithStruct14zvoe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zvoe > 0 || missingFieldsLeft14zvoe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zvoe, missingFieldsLeft14zvoe, msgp.ShowFound(found14zvoe[:]), decodeMsgFieldOrder14zvoe)
		if encodedFieldsLeft14zvoe > 0 {
			encodedFieldsLeft14zvoe--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField14zvoe = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zvoe < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zvoe = 0
			}
			for nextMiss14zvoe < maxFields14zvoe && found14zvoe[nextMiss14zvoe] {
				nextMiss14zvoe++
			}
			if nextMiss14zvoe == maxFields14zvoe {
				// filled all the empty fields!
				break doneWithStruct14zvoe
			}
			missingFieldsLeft14zvoe--
			curField14zvoe = decodeMsgFieldOrder14zvoe[nextMiss14zvoe]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zvoe)
		switch curField14zvoe {
		// -- templateDecodeMsg ends here --

		case "Id":
			found14zvoe[0] = true
			{
				var zjgq int64
				zjgq, err = dc.ReadInt64()
				z.Id = StructTypeId(zjgq)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found14zvoe[1] = true
			var zsvm uint32
			zsvm, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zsvm > 0 {
				z.Da = make(map[int64]msgp.Raw, zsvm)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zsvm > 0 {
				zsvm--
				var zjuh int64
				var zvqy msgp.Raw
				zjuh, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zvqy.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zjuh] = zvqy
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zvoe != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZStructPacket
var decodeMsgFieldOrder14zvoe = []string{"Id", "Da"}

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
	for zjuh, zvqy := range z.Da {
		err = en.WriteInt64(zjuh)
		if err != nil {
			panic(err)
		}
		err = zvqy.EncodeMsg(en)
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
	for zjuh, zvqy := range z.Da {
		o = msgp.AppendInt64(o, zjuh)
		o, err = zvqy.MarshalMsg(o)
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
	const maxFields15zlkb = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields15zlkb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zlkb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft15zlkb := totalEncodedFields15zlkb
	missingFieldsLeft15zlkb := maxFields15zlkb - totalEncodedFields15zlkb

	var nextMiss15zlkb int32 = -1
	var found15zlkb [maxFields15zlkb]bool
	var curField15zlkb string

doneWithStruct15zlkb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zlkb > 0 || missingFieldsLeft15zlkb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zlkb, missingFieldsLeft15zlkb, msgp.ShowFound(found15zlkb[:]), unmarshalMsgFieldOrder15zlkb)
		if encodedFieldsLeft15zlkb > 0 {
			encodedFieldsLeft15zlkb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField15zlkb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss15zlkb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zlkb = 0
			}
			for nextMiss15zlkb < maxFields15zlkb && found15zlkb[nextMiss15zlkb] {
				nextMiss15zlkb++
			}
			if nextMiss15zlkb == maxFields15zlkb {
				// filled all the empty fields!
				break doneWithStruct15zlkb
			}
			missingFieldsLeft15zlkb--
			curField15zlkb = unmarshalMsgFieldOrder15zlkb[nextMiss15zlkb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zlkb)
		switch curField15zlkb {
		// -- templateUnmarshalMsg ends here --

		case "Id":
			found15zlkb[0] = true
			{
				var zpit int64
				zpit, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zpit)
			}
		case "Da":
			found15zlkb[1] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zcwl uint32
				zcwl, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zcwl > 0 {
					z.Da = make(map[int64]msgp.Raw, zcwl)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zcwl > 0 {
					var zjuh int64
					var zvqy msgp.Raw
					zcwl--
					zjuh, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zvqy.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zjuh] = zvqy
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss15zlkb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZStructPacket
var unmarshalMsgFieldOrder15zlkb = []string{"Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZStructPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zjuh, zvqy := range z.Da {
			_ = zvqy
			_ = zjuh
			s += msgp.Int64Size + zvqy.Msgsize()
		}
	}
	return
}
