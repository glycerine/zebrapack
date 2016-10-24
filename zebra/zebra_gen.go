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
	const maxFields0zdyb = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zdyb uint32
	totalEncodedFields0zdyb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zdyb := totalEncodedFields0zdyb
	missingFieldsLeft0zdyb := maxFields0zdyb - totalEncodedFields0zdyb

	var nextMiss0zdyb int32 = -1
	var found0zdyb [maxFields0zdyb]bool
	var curField0zdyb string

doneWithStruct0zdyb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zdyb > 0 || missingFieldsLeft0zdyb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zdyb, missingFieldsLeft0zdyb, msgp.ShowFound(found0zdyb[:]), decodeMsgFieldOrder0zdyb)
		if encodedFieldsLeft0zdyb > 0 {
			encodedFieldsLeft0zdyb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zdyb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zdyb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zdyb = 0
			}
			for nextMiss0zdyb < maxFields0zdyb && found0zdyb[nextMiss0zdyb] {
				nextMiss0zdyb++
			}
			if nextMiss0zdyb == maxFields0zdyb {
				// filled all the empty fields!
				break doneWithStruct0zdyb
			}
			missingFieldsLeft0zdyb--
			curField0zdyb = decodeMsgFieldOrder0zdyb[nextMiss0zdyb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zdyb)
		switch curField0zdyb {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zdyb[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zdyb[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Zknd":
			found0zdyb[2] = true
			{
				var zugs int32
				zugs, err = dc.ReadInt32()
				z.Zknd = Zkind(zugs)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0zdyb[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zdyb[4] = true
			var zmnq uint32
			zmnq, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zmnq > 0 {
				z.Tag = make(map[string]string, zmnq)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zmnq > 0 {
				zmnq--
				var zvxz string
				var zqbq string
				zvxz, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zqbq, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zvxz] = zqbq
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zdyb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zdyb = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

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
	var empty_zfzv [5]bool
	fieldsInUse_zjvl := z.fieldsNotEmpty(empty_zfzv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjvl)
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
	err = en.WriteInt32(int32(z.Zknd))
	if err != nil {
		panic(err)
	}
	if !empty_zfzv[3] {
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

	if !empty_zfzv[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zvxz, zqbq := range z.Tag {
			err = en.WriteString(zvxz)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zqbq)
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
	o = msgp.AppendInt32(o, int32(z.Zknd))
	if !empty[3] {
		// string "Varg"
		o = append(o, 0xa4, 0x56, 0x61, 0x72, 0x67)
		o = msgp.AppendBool(o, z.Varg)
	}

	if !empty[4] {
		// string "Tag"
		o = append(o, 0xa3, 0x54, 0x61, 0x67)
		o = msgp.AppendMapHeader(o, uint32(len(z.Tag)))
		for zvxz, zqbq := range z.Tag {
			o = msgp.AppendString(o, zvxz)
			o = msgp.AppendString(o, zqbq)
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
	const maxFields1zbvp = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zbvp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zbvp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zbvp := totalEncodedFields1zbvp
	missingFieldsLeft1zbvp := maxFields1zbvp - totalEncodedFields1zbvp

	var nextMiss1zbvp int32 = -1
	var found1zbvp [maxFields1zbvp]bool
	var curField1zbvp string

doneWithStruct1zbvp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zbvp > 0 || missingFieldsLeft1zbvp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zbvp, missingFieldsLeft1zbvp, msgp.ShowFound(found1zbvp[:]), unmarshalMsgFieldOrder1zbvp)
		if encodedFieldsLeft1zbvp > 0 {
			encodedFieldsLeft1zbvp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zbvp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zbvp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zbvp = 0
			}
			for nextMiss1zbvp < maxFields1zbvp && found1zbvp[nextMiss1zbvp] {
				nextMiss1zbvp++
			}
			if nextMiss1zbvp == maxFields1zbvp {
				// filled all the empty fields!
				break doneWithStruct1zbvp
			}
			missingFieldsLeft1zbvp--
			curField1zbvp = unmarshalMsgFieldOrder1zbvp[nextMiss1zbvp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zbvp)
		switch curField1zbvp {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zbvp[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zbvp[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Zknd":
			found1zbvp[2] = true
			{
				var zknh int32
				zknh, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Zknd = Zkind(zknh)
			}
		case "Varg":
			found1zbvp[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zbvp[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zbwf uint32
				zbwf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zbwf > 0 {
					z.Tag = make(map[string]string, zbwf)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zbwf > 0 {
					var zvxz string
					var zqbq string
					zbwf--
					zvxz, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zqbq, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zvxz] = zqbq
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zbvp != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zbvp = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zvxz, zqbq := range z.Tag {
			_ = zqbq
			_ = zvxz
			s += msgp.StringPrefixSize + len(zvxz) + msgp.StringPrefixSize + len(zqbq)
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
	const maxFields2zekm = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zekm uint32
	totalEncodedFields2zekm, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zekm := totalEncodedFields2zekm
	missingFieldsLeft2zekm := maxFields2zekm - totalEncodedFields2zekm

	var nextMiss2zekm int32 = -1
	var found2zekm [maxFields2zekm]bool
	var curField2zekm string

doneWithStruct2zekm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zekm > 0 || missingFieldsLeft2zekm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zekm, missingFieldsLeft2zekm, msgp.ShowFound(found2zekm[:]), decodeMsgFieldOrder2zekm)
		if encodedFieldsLeft2zekm > 0 {
			encodedFieldsLeft2zekm--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zekm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zekm < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zekm = 0
			}
			for nextMiss2zekm < maxFields2zekm && found2zekm[nextMiss2zekm] {
				nextMiss2zekm++
			}
			if nextMiss2zekm == maxFields2zekm {
				// filled all the empty fields!
				break doneWithStruct2zekm
			}
			missingFieldsLeft2zekm--
			curField2zekm = decodeMsgFieldOrder2zekm[nextMiss2zekm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zekm)
		switch curField2zekm {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zekm[0] = true
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
	if nextMiss2zekm != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zekm = []string{"IfaceId"}

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
	const maxFields3zacb = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zacb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zacb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zacb := totalEncodedFields3zacb
	missingFieldsLeft3zacb := maxFields3zacb - totalEncodedFields3zacb

	var nextMiss3zacb int32 = -1
	var found3zacb [maxFields3zacb]bool
	var curField3zacb string

doneWithStruct3zacb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zacb > 0 || missingFieldsLeft3zacb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zacb, missingFieldsLeft3zacb, msgp.ShowFound(found3zacb[:]), unmarshalMsgFieldOrder3zacb)
		if encodedFieldsLeft3zacb > 0 {
			encodedFieldsLeft3zacb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zacb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zacb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zacb = 0
			}
			for nextMiss3zacb < maxFields3zacb && found3zacb[nextMiss3zacb] {
				nextMiss3zacb++
			}
			if nextMiss3zacb == maxFields3zacb {
				// filled all the empty fields!
				break doneWithStruct3zacb
			}
			missingFieldsLeft3zacb--
			curField3zacb = unmarshalMsgFieldOrder3zacb[nextMiss3zacb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zacb)
		switch curField3zacb {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3zacb[0] = true
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
	if nextMiss3zacb != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3zacb = []string{"IfaceId"}

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
	const maxFields4zaua = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zaua uint32
	totalEncodedFields4zaua, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zaua := totalEncodedFields4zaua
	missingFieldsLeft4zaua := maxFields4zaua - totalEncodedFields4zaua

	var nextMiss4zaua int32 = -1
	var found4zaua [maxFields4zaua]bool
	var curField4zaua string

doneWithStruct4zaua:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zaua > 0 || missingFieldsLeft4zaua > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zaua, missingFieldsLeft4zaua, msgp.ShowFound(found4zaua[:]), decodeMsgFieldOrder4zaua)
		if encodedFieldsLeft4zaua > 0 {
			encodedFieldsLeft4zaua--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zaua = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zaua < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zaua = 0
			}
			for nextMiss4zaua < maxFields4zaua && found4zaua[nextMiss4zaua] {
				nextMiss4zaua++
			}
			if nextMiss4zaua == maxFields4zaua {
				// filled all the empty fields!
				break doneWithStruct4zaua
			}
			missingFieldsLeft4zaua--
			curField4zaua = decodeMsgFieldOrder4zaua[nextMiss4zaua]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zaua)
		switch curField4zaua {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4zaua[0] = true
			var zxsw uint32
			zxsw, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(zxsw) {
				z.Methods = (z.Methods)[:zxsw]
			} else {
				z.Methods = make([]MethodT, zxsw)
			}
			for zddc := range z.Methods {
				err = z.Methods[zddc].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4zaua[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zaua[2] = true
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
	if nextMiss4zaua != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4zaua = []string{"Methods", "Deprecated", "Comment"}

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
	for zddc := range z.Methods {
		err = z.Methods[zddc].EncodeMsg(en)
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
	for zddc := range z.Methods {
		o, err = z.Methods[zddc].MarshalMsg(o)
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
	const maxFields5zpnk = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zpnk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zpnk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zpnk := totalEncodedFields5zpnk
	missingFieldsLeft5zpnk := maxFields5zpnk - totalEncodedFields5zpnk

	var nextMiss5zpnk int32 = -1
	var found5zpnk [maxFields5zpnk]bool
	var curField5zpnk string

doneWithStruct5zpnk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zpnk > 0 || missingFieldsLeft5zpnk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zpnk, missingFieldsLeft5zpnk, msgp.ShowFound(found5zpnk[:]), unmarshalMsgFieldOrder5zpnk)
		if encodedFieldsLeft5zpnk > 0 {
			encodedFieldsLeft5zpnk--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zpnk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zpnk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zpnk = 0
			}
			for nextMiss5zpnk < maxFields5zpnk && found5zpnk[nextMiss5zpnk] {
				nextMiss5zpnk++
			}
			if nextMiss5zpnk == maxFields5zpnk {
				// filled all the empty fields!
				break doneWithStruct5zpnk
			}
			missingFieldsLeft5zpnk--
			curField5zpnk = unmarshalMsgFieldOrder5zpnk[nextMiss5zpnk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zpnk)
		switch curField5zpnk {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zpnk[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zmtr uint32
				zmtr, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zmtr) {
					z.Methods = (z.Methods)[:zmtr]
				} else {
					z.Methods = make([]MethodT, zmtr)
				}
				for zddc := range z.Methods {
					bts, err = z.Methods[zddc].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zpnk[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zpnk[2] = true
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
	if nextMiss5zpnk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zpnk = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zddc := range z.Methods {
		s += z.Methods[zddc].Msgsize()
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
	const maxFields6zkuw = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zkuw uint32
	totalEncodedFields6zkuw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zkuw := totalEncodedFields6zkuw
	missingFieldsLeft6zkuw := maxFields6zkuw - totalEncodedFields6zkuw

	var nextMiss6zkuw int32 = -1
	var found6zkuw [maxFields6zkuw]bool
	var curField6zkuw string

doneWithStruct6zkuw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zkuw > 0 || missingFieldsLeft6zkuw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zkuw, missingFieldsLeft6zkuw, msgp.ShowFound(found6zkuw[:]), decodeMsgFieldOrder6zkuw)
		if encodedFieldsLeft6zkuw > 0 {
			encodedFieldsLeft6zkuw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zkuw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zkuw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zkuw = 0
			}
			for nextMiss6zkuw < maxFields6zkuw && found6zkuw[nextMiss6zkuw] {
				nextMiss6zkuw++
			}
			if nextMiss6zkuw == maxFields6zkuw {
				// filled all the empty fields!
				break doneWithStruct6zkuw
			}
			missingFieldsLeft6zkuw--
			curField6zkuw = decodeMsgFieldOrder6zkuw[nextMiss6zkuw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zkuw)
		switch curField6zkuw {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zkuw[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zkuw[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zkuw[2] = true
			const maxFields7zifg = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7zifg uint32
			totalEncodedFields7zifg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7zifg := totalEncodedFields7zifg
			missingFieldsLeft7zifg := maxFields7zifg - totalEncodedFields7zifg

			var nextMiss7zifg int32 = -1
			var found7zifg [maxFields7zifg]bool
			var curField7zifg string

		doneWithStruct7zifg:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7zifg > 0 || missingFieldsLeft7zifg > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zifg, missingFieldsLeft7zifg, msgp.ShowFound(found7zifg[:]), decodeMsgFieldOrder7zifg)
				if encodedFieldsLeft7zifg > 0 {
					encodedFieldsLeft7zifg--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7zifg = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7zifg < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7zifg = 0
					}
					for nextMiss7zifg < maxFields7zifg && found7zifg[nextMiss7zifg] {
						nextMiss7zifg++
					}
					if nextMiss7zifg == maxFields7zifg {
						// filled all the empty fields!
						break doneWithStruct7zifg
					}
					missingFieldsLeft7zifg--
					curField7zifg = decodeMsgFieldOrder7zifg[nextMiss7zifg]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7zifg)
				switch curField7zifg {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7zifg[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7zifg[1] = true
					var zmvb uint32
					zmvb, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zmvb) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zmvb]
					} else {
						z.Inputs.Fld = make([]FieldT, zmvb)
					}
					for zerb := range z.Inputs.Fld {
						err = z.Inputs.Fld[zerb].DecodeMsg(dc)
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
			if nextMiss7zifg != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zkuw[3] = true
			const maxFields8zztp = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zztp uint32
			totalEncodedFields8zztp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zztp := totalEncodedFields8zztp
			missingFieldsLeft8zztp := maxFields8zztp - totalEncodedFields8zztp

			var nextMiss8zztp int32 = -1
			var found8zztp [maxFields8zztp]bool
			var curField8zztp string

		doneWithStruct8zztp:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zztp > 0 || missingFieldsLeft8zztp > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zztp, missingFieldsLeft8zztp, msgp.ShowFound(found8zztp[:]), decodeMsgFieldOrder8zztp)
				if encodedFieldsLeft8zztp > 0 {
					encodedFieldsLeft8zztp--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zztp = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zztp < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zztp = 0
					}
					for nextMiss8zztp < maxFields8zztp && found8zztp[nextMiss8zztp] {
						nextMiss8zztp++
					}
					if nextMiss8zztp == maxFields8zztp {
						// filled all the empty fields!
						break doneWithStruct8zztp
					}
					missingFieldsLeft8zztp--
					curField8zztp = decodeMsgFieldOrder8zztp[nextMiss8zztp]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zztp)
				switch curField8zztp {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zztp[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zztp[1] = true
					var zcqr uint32
					zcqr, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zcqr) {
						z.Returns.Fld = (z.Returns.Fld)[:zcqr]
					} else {
						z.Returns.Fld = make([]FieldT, zcqr)
					}
					for zmec := range z.Returns.Fld {
						err = z.Returns.Fld[zmec].DecodeMsg(dc)
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
			if nextMiss8zztp != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zkuw[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zkuw[5] = true
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
	if nextMiss6zkuw != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zkuw = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7zifg = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zztp = []string{"Nam", "Fld"}

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
	for zerb := range z.Inputs.Fld {
		err = z.Inputs.Fld[zerb].EncodeMsg(en)
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
	for zmec := range z.Returns.Fld {
		err = z.Returns.Fld[zmec].EncodeMsg(en)
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
	for zerb := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zerb].MarshalMsg(o)
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
	for zmec := range z.Returns.Fld {
		o, err = z.Returns.Fld[zmec].MarshalMsg(o)
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
	const maxFields9zlev = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zlev uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zlev, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zlev := totalEncodedFields9zlev
	missingFieldsLeft9zlev := maxFields9zlev - totalEncodedFields9zlev

	var nextMiss9zlev int32 = -1
	var found9zlev [maxFields9zlev]bool
	var curField9zlev string

doneWithStruct9zlev:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zlev > 0 || missingFieldsLeft9zlev > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zlev, missingFieldsLeft9zlev, msgp.ShowFound(found9zlev[:]), unmarshalMsgFieldOrder9zlev)
		if encodedFieldsLeft9zlev > 0 {
			encodedFieldsLeft9zlev--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zlev = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zlev < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zlev = 0
			}
			for nextMiss9zlev < maxFields9zlev && found9zlev[nextMiss9zlev] {
				nextMiss9zlev++
			}
			if nextMiss9zlev == maxFields9zlev {
				// filled all the empty fields!
				break doneWithStruct9zlev
			}
			missingFieldsLeft9zlev--
			curField9zlev = unmarshalMsgFieldOrder9zlev[nextMiss9zlev]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zlev)
		switch curField9zlev {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zlev[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zlev[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zlev[2] = true
			const maxFields10zaqt = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zaqt uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zaqt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10zaqt := totalEncodedFields10zaqt
			missingFieldsLeft10zaqt := maxFields10zaqt - totalEncodedFields10zaqt

			var nextMiss10zaqt int32 = -1
			var found10zaqt [maxFields10zaqt]bool
			var curField10zaqt string

		doneWithStruct10zaqt:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10zaqt > 0 || missingFieldsLeft10zaqt > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zaqt, missingFieldsLeft10zaqt, msgp.ShowFound(found10zaqt[:]), unmarshalMsgFieldOrder10zaqt)
				if encodedFieldsLeft10zaqt > 0 {
					encodedFieldsLeft10zaqt--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10zaqt = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zaqt < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zaqt = 0
					}
					for nextMiss10zaqt < maxFields10zaqt && found10zaqt[nextMiss10zaqt] {
						nextMiss10zaqt++
					}
					if nextMiss10zaqt == maxFields10zaqt {
						// filled all the empty fields!
						break doneWithStruct10zaqt
					}
					missingFieldsLeft10zaqt--
					curField10zaqt = unmarshalMsgFieldOrder10zaqt[nextMiss10zaqt]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zaqt)
				switch curField10zaqt {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10zaqt[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10zaqt[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zrel uint32
						zrel, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zrel) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zrel]
						} else {
							z.Inputs.Fld = make([]FieldT, zrel)
						}
						for zerb := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zerb].UnmarshalMsg(bts)
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
			if nextMiss10zaqt != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zlev[3] = true
			const maxFields11zvua = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zvua uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zvua, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zvua := totalEncodedFields11zvua
			missingFieldsLeft11zvua := maxFields11zvua - totalEncodedFields11zvua

			var nextMiss11zvua int32 = -1
			var found11zvua [maxFields11zvua]bool
			var curField11zvua string

		doneWithStruct11zvua:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zvua > 0 || missingFieldsLeft11zvua > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zvua, missingFieldsLeft11zvua, msgp.ShowFound(found11zvua[:]), unmarshalMsgFieldOrder11zvua)
				if encodedFieldsLeft11zvua > 0 {
					encodedFieldsLeft11zvua--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zvua = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zvua < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zvua = 0
					}
					for nextMiss11zvua < maxFields11zvua && found11zvua[nextMiss11zvua] {
						nextMiss11zvua++
					}
					if nextMiss11zvua == maxFields11zvua {
						// filled all the empty fields!
						break doneWithStruct11zvua
					}
					missingFieldsLeft11zvua--
					curField11zvua = unmarshalMsgFieldOrder11zvua[nextMiss11zvua]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zvua)
				switch curField11zvua {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zvua[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zvua[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var zyrx uint32
						zyrx, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(zyrx) {
							z.Returns.Fld = (z.Returns.Fld)[:zyrx]
						} else {
							z.Returns.Fld = make([]FieldT, zyrx)
						}
						for zmec := range z.Returns.Fld {
							bts, err = z.Returns.Fld[zmec].UnmarshalMsg(bts)
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
			if nextMiss11zvua != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zlev[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zlev[5] = true
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
	if nextMiss9zlev != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zlev = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10zaqt = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zvua = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zerb := range z.Inputs.Fld {
		s += z.Inputs.Fld[zerb].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for zmec := range z.Returns.Fld {
		s += z.Returns.Fld[zmec].Msgsize()
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
	const maxFields12zqln = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zqln uint32
	totalEncodedFields12zqln, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zqln := totalEncodedFields12zqln
	missingFieldsLeft12zqln := maxFields12zqln - totalEncodedFields12zqln

	var nextMiss12zqln int32 = -1
	var found12zqln [maxFields12zqln]bool
	var curField12zqln string

doneWithStruct12zqln:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zqln > 0 || missingFieldsLeft12zqln > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zqln, missingFieldsLeft12zqln, msgp.ShowFound(found12zqln[:]), decodeMsgFieldOrder12zqln)
		if encodedFieldsLeft12zqln > 0 {
			encodedFieldsLeft12zqln--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zqln = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zqln < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zqln = 0
			}
			for nextMiss12zqln < maxFields12zqln && found12zqln[nextMiss12zqln] {
				nextMiss12zqln++
			}
			if nextMiss12zqln == maxFields12zqln {
				// filled all the empty fields!
				break doneWithStruct12zqln
			}
			missingFieldsLeft12zqln--
			curField12zqln = decodeMsgFieldOrder12zqln[nextMiss12zqln]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zqln)
		switch curField12zqln {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zqln[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found12zqln[1] = true
			var zzjp uint32
			zzjp, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zzjp > 0 {
				z.Structs = make(map[int64]StructT, zzjp)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zzjp > 0 {
				zzjp--
				var zvov int64
				var zhek StructT
				zvov, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				const maxFields13zoau = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields13zoau uint32
				totalEncodedFields13zoau, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft13zoau := totalEncodedFields13zoau
				missingFieldsLeft13zoau := maxFields13zoau - totalEncodedFields13zoau

				var nextMiss13zoau int32 = -1
				var found13zoau [maxFields13zoau]bool
				var curField13zoau string

			doneWithStruct13zoau:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft13zoau > 0 || missingFieldsLeft13zoau > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zoau, missingFieldsLeft13zoau, msgp.ShowFound(found13zoau[:]), decodeMsgFieldOrder13zoau)
					if encodedFieldsLeft13zoau > 0 {
						encodedFieldsLeft13zoau--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField13zoau = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss13zoau < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss13zoau = 0
						}
						for nextMiss13zoau < maxFields13zoau && found13zoau[nextMiss13zoau] {
							nextMiss13zoau++
						}
						if nextMiss13zoau == maxFields13zoau {
							// filled all the empty fields!
							break doneWithStruct13zoau
						}
						missingFieldsLeft13zoau--
						curField13zoau = decodeMsgFieldOrder13zoau[nextMiss13zoau]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField13zoau)
					switch curField13zoau {
					// -- templateDecodeMsg ends here --

					case "Nam":
						found13zoau[0] = true
						zhek.Nam, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found13zoau[1] = true
						var zghp uint32
						zghp, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(zhek.Fld) >= int(zghp) {
							zhek.Fld = (zhek.Fld)[:zghp]
						} else {
							zhek.Fld = make([]FieldT, zghp)
						}
						for znld := range zhek.Fld {
							err = zhek.Fld[znld].DecodeMsg(dc)
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
				if nextMiss13zoau != -1 {
					dc.PopAlwaysNil()
				}

				z.Structs[zvov] = zhek
			}
		case "Interfaces":
			found12zqln[2] = true
			var zbjb uint32
			zbjb, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Interfaces == nil && zbjb > 0 {
				z.Interfaces = make(map[int64]InterfaceT, zbjb)
			} else if len(z.Interfaces) > 0 {
				for key, _ := range z.Interfaces {
					delete(z.Interfaces, key)
				}
			}
			for zbjb > 0 {
				zbjb--
				var zwkg int64
				var zbnt InterfaceT
				zwkg, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zbnt.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Interfaces[zwkg] = zbnt
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zqln != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zqln = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var decodeMsgFieldOrder13zoau = []string{"Nam", "Fld"}

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
	for zvov, zhek := range z.Structs {
		err = en.WriteInt64(zvov)
		if err != nil {
			panic(err)
		}
		// map header, size 2
		// write "Nam"
		err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteString(zhek.Nam)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(zhek.Fld)))
		if err != nil {
			panic(err)
		}
		for znld := range zhek.Fld {
			err = zhek.Fld[znld].EncodeMsg(en)
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
	for zwkg, zbnt := range z.Interfaces {
		err = en.WriteInt64(zwkg)
		if err != nil {
			panic(err)
		}
		err = zbnt.EncodeMsg(en)
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
	for zvov, zhek := range z.Structs {
		o = msgp.AppendInt64(o, zvov)
		// map header, size 2
		// string "Nam"
		o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
		o = msgp.AppendString(o, zhek.Nam)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(zhek.Fld)))
		for znld := range zhek.Fld {
			o, err = zhek.Fld[znld].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Interfaces"
	o = append(o, 0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Interfaces)))
	for zwkg, zbnt := range z.Interfaces {
		o = msgp.AppendInt64(o, zwkg)
		o, err = zbnt.MarshalMsg(o)
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
	const maxFields14zzna = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields14zzna uint32
	if !nbs.AlwaysNil {
		totalEncodedFields14zzna, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft14zzna := totalEncodedFields14zzna
	missingFieldsLeft14zzna := maxFields14zzna - totalEncodedFields14zzna

	var nextMiss14zzna int32 = -1
	var found14zzna [maxFields14zzna]bool
	var curField14zzna string

doneWithStruct14zzna:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zzna > 0 || missingFieldsLeft14zzna > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zzna, missingFieldsLeft14zzna, msgp.ShowFound(found14zzna[:]), unmarshalMsgFieldOrder14zzna)
		if encodedFieldsLeft14zzna > 0 {
			encodedFieldsLeft14zzna--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField14zzna = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zzna < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss14zzna = 0
			}
			for nextMiss14zzna < maxFields14zzna && found14zzna[nextMiss14zzna] {
				nextMiss14zzna++
			}
			if nextMiss14zzna == maxFields14zzna {
				// filled all the empty fields!
				break doneWithStruct14zzna
			}
			missingFieldsLeft14zzna--
			curField14zzna = unmarshalMsgFieldOrder14zzna[nextMiss14zzna]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zzna)
		switch curField14zzna {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found14zzna[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found14zzna[1] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zfok uint32
				zfok, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zfok > 0 {
					z.Structs = make(map[int64]StructT, zfok)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zfok > 0 {
					var zvov int64
					var zhek StructT
					zfok--
					zvov, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					const maxFields15zxnp = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields15zxnp uint32
					if !nbs.AlwaysNil {
						totalEncodedFields15zxnp, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft15zxnp := totalEncodedFields15zxnp
					missingFieldsLeft15zxnp := maxFields15zxnp - totalEncodedFields15zxnp

					var nextMiss15zxnp int32 = -1
					var found15zxnp [maxFields15zxnp]bool
					var curField15zxnp string

				doneWithStruct15zxnp:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft15zxnp > 0 || missingFieldsLeft15zxnp > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zxnp, missingFieldsLeft15zxnp, msgp.ShowFound(found15zxnp[:]), unmarshalMsgFieldOrder15zxnp)
						if encodedFieldsLeft15zxnp > 0 {
							encodedFieldsLeft15zxnp--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField15zxnp = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss15zxnp < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss15zxnp = 0
							}
							for nextMiss15zxnp < maxFields15zxnp && found15zxnp[nextMiss15zxnp] {
								nextMiss15zxnp++
							}
							if nextMiss15zxnp == maxFields15zxnp {
								// filled all the empty fields!
								break doneWithStruct15zxnp
							}
							missingFieldsLeft15zxnp--
							curField15zxnp = unmarshalMsgFieldOrder15zxnp[nextMiss15zxnp]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField15zxnp)
						switch curField15zxnp {
						// -- templateUnmarshalMsg ends here --

						case "Nam":
							found15zxnp[0] = true
							zhek.Nam, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found15zxnp[1] = true
							if nbs.AlwaysNil {
								(zhek.Fld) = (zhek.Fld)[:0]
							} else {

								var ztpr uint32
								ztpr, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(zhek.Fld) >= int(ztpr) {
									zhek.Fld = (zhek.Fld)[:ztpr]
								} else {
									zhek.Fld = make([]FieldT, ztpr)
								}
								for znld := range zhek.Fld {
									bts, err = zhek.Fld[znld].UnmarshalMsg(bts)
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
					if nextMiss15zxnp != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Structs[zvov] = zhek
				}
			}
		case "Interfaces":
			found14zzna[2] = true
			if nbs.AlwaysNil {
				if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}

			} else {

				var zpii uint32
				zpii, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Interfaces == nil && zpii > 0 {
					z.Interfaces = make(map[int64]InterfaceT, zpii)
				} else if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}
				for zpii > 0 {
					var zwkg int64
					var zbnt InterfaceT
					zpii--
					zwkg, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zbnt.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Interfaces[zwkg] = zbnt
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zzna != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder14zzna = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var unmarshalMsgFieldOrder15zxnp = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zvov, zhek := range z.Structs {
			_ = zhek
			_ = zvov
			s += msgp.Int64Size + 1 + 4 + msgp.StringPrefixSize + len(zhek.Nam) + 4 + msgp.ArrayHeaderSize
			for znld := range zhek.Fld {
				s += zhek.Fld[znld].Msgsize()
			}
		}
	}
	s += 11 + msgp.MapHeaderSize
	if z.Interfaces != nil {
		for zwkg, zbnt := range z.Interfaces {
			_ = zbnt
			_ = zwkg
			s += msgp.Int64Size + zbnt.Msgsize()
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
	const maxFields16zoqr = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zoqr uint32
	totalEncodedFields16zoqr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zoqr := totalEncodedFields16zoqr
	missingFieldsLeft16zoqr := maxFields16zoqr - totalEncodedFields16zoqr

	var nextMiss16zoqr int32 = -1
	var found16zoqr [maxFields16zoqr]bool
	var curField16zoqr string

doneWithStruct16zoqr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zoqr > 0 || missingFieldsLeft16zoqr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zoqr, missingFieldsLeft16zoqr, msgp.ShowFound(found16zoqr[:]), decodeMsgFieldOrder16zoqr)
		if encodedFieldsLeft16zoqr > 0 {
			encodedFieldsLeft16zoqr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zoqr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zoqr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zoqr = 0
			}
			for nextMiss16zoqr < maxFields16zoqr && found16zoqr[nextMiss16zoqr] {
				nextMiss16zoqr++
			}
			if nextMiss16zoqr == maxFields16zoqr {
				// filled all the empty fields!
				break doneWithStruct16zoqr
			}
			missingFieldsLeft16zoqr--
			curField16zoqr = decodeMsgFieldOrder16zoqr[nextMiss16zoqr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zoqr)
		switch curField16zoqr {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found16zoqr[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found16zoqr[1] = true
			var zuvn uint32
			zuvn, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(zuvn) {
				z.Fld = (z.Fld)[:zuvn]
			} else {
				z.Fld = make([]FieldT, zuvn)
			}
			for zqyr := range z.Fld {
				err = z.Fld[zqyr].DecodeMsg(dc)
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
	if nextMiss16zoqr != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder16zoqr = []string{"Nam", "Fld"}

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
	for zqyr := range z.Fld {
		err = z.Fld[zqyr].EncodeMsg(en)
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
	for zqyr := range z.Fld {
		o, err = z.Fld[zqyr].MarshalMsg(o)
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
	const maxFields17znhj = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17znhj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17znhj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17znhj := totalEncodedFields17znhj
	missingFieldsLeft17znhj := maxFields17znhj - totalEncodedFields17znhj

	var nextMiss17znhj int32 = -1
	var found17znhj [maxFields17znhj]bool
	var curField17znhj string

doneWithStruct17znhj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17znhj > 0 || missingFieldsLeft17znhj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17znhj, missingFieldsLeft17znhj, msgp.ShowFound(found17znhj[:]), unmarshalMsgFieldOrder17znhj)
		if encodedFieldsLeft17znhj > 0 {
			encodedFieldsLeft17znhj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17znhj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17znhj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17znhj = 0
			}
			for nextMiss17znhj < maxFields17znhj && found17znhj[nextMiss17znhj] {
				nextMiss17znhj++
			}
			if nextMiss17znhj == maxFields17znhj {
				// filled all the empty fields!
				break doneWithStruct17znhj
			}
			missingFieldsLeft17znhj--
			curField17znhj = unmarshalMsgFieldOrder17znhj[nextMiss17znhj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17znhj)
		switch curField17znhj {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found17znhj[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found17znhj[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var ztwp uint32
				ztwp, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(ztwp) {
					z.Fld = (z.Fld)[:ztwp]
				} else {
					z.Fld = make([]FieldT, ztwp)
				}
				for zqyr := range z.Fld {
					bts, err = z.Fld[zqyr].UnmarshalMsg(bts)
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
	if nextMiss17znhj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder17znhj = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zqyr := range z.Fld {
		s += z.Fld[zqyr].Msgsize()
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
		var zbvl int64
		zbvl, err = dc.ReadInt64()
		(*z) = StructTypeId(zbvl)
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
		var zkxr int64
		zkxr, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zkxr)
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
	const maxFields18zrqk = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields18zrqk uint32
	totalEncodedFields18zrqk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft18zrqk := totalEncodedFields18zrqk
	missingFieldsLeft18zrqk := maxFields18zrqk - totalEncodedFields18zrqk

	var nextMiss18zrqk int32 = -1
	var found18zrqk [maxFields18zrqk]bool
	var curField18zrqk string

doneWithStruct18zrqk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft18zrqk > 0 || missingFieldsLeft18zrqk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zrqk, missingFieldsLeft18zrqk, msgp.ShowFound(found18zrqk[:]), decodeMsgFieldOrder18zrqk)
		if encodedFieldsLeft18zrqk > 0 {
			encodedFieldsLeft18zrqk--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField18zrqk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss18zrqk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss18zrqk = 0
			}
			for nextMiss18zrqk < maxFields18zrqk && found18zrqk[nextMiss18zrqk] {
				nextMiss18zrqk++
			}
			if nextMiss18zrqk == maxFields18zrqk {
				// filled all the empty fields!
				break doneWithStruct18zrqk
			}
			missingFieldsLeft18zrqk--
			curField18zrqk = decodeMsgFieldOrder18zrqk[nextMiss18zrqk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField18zrqk)
		switch curField18zrqk {
		// -- templateDecodeMsg ends here --

		case "Id":
			found18zrqk[0] = true
			{
				var zymp int64
				zymp, err = dc.ReadInt64()
				z.Id = StructTypeId(zymp)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found18zrqk[1] = true
			var zugq uint32
			zugq, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zugq > 0 {
				z.Da = make(map[int64]msgp.Raw, zugq)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zugq > 0 {
				zugq--
				var zbbj int64
				var zvqo msgp.Raw
				zbbj, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zvqo.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zbbj] = zvqo
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss18zrqk != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZStructPacket
var decodeMsgFieldOrder18zrqk = []string{"Id", "Da"}

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
	for zbbj, zvqo := range z.Da {
		err = en.WriteInt64(zbbj)
		if err != nil {
			panic(err)
		}
		err = zvqo.EncodeMsg(en)
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
	for zbbj, zvqo := range z.Da {
		o = msgp.AppendInt64(o, zbbj)
		o, err = zvqo.MarshalMsg(o)
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
	const maxFields19zfut = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19zfut uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19zfut, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19zfut := totalEncodedFields19zfut
	missingFieldsLeft19zfut := maxFields19zfut - totalEncodedFields19zfut

	var nextMiss19zfut int32 = -1
	var found19zfut [maxFields19zfut]bool
	var curField19zfut string

doneWithStruct19zfut:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zfut > 0 || missingFieldsLeft19zfut > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zfut, missingFieldsLeft19zfut, msgp.ShowFound(found19zfut[:]), unmarshalMsgFieldOrder19zfut)
		if encodedFieldsLeft19zfut > 0 {
			encodedFieldsLeft19zfut--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19zfut = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zfut < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19zfut = 0
			}
			for nextMiss19zfut < maxFields19zfut && found19zfut[nextMiss19zfut] {
				nextMiss19zfut++
			}
			if nextMiss19zfut == maxFields19zfut {
				// filled all the empty fields!
				break doneWithStruct19zfut
			}
			missingFieldsLeft19zfut--
			curField19zfut = unmarshalMsgFieldOrder19zfut[nextMiss19zfut]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zfut)
		switch curField19zfut {
		// -- templateUnmarshalMsg ends here --

		case "Id":
			found19zfut[0] = true
			{
				var zlac int64
				zlac, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zlac)
			}
		case "Da":
			found19zfut[1] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zvri uint32
				zvri, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zvri > 0 {
					z.Da = make(map[int64]msgp.Raw, zvri)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zvri > 0 {
					var zbbj int64
					var zvqo msgp.Raw
					zvri--
					zbbj, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zvqo.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zbbj] = zvqo
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19zfut != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZStructPacket
var unmarshalMsgFieldOrder19zfut = []string{"Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZStructPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zbbj, zvqo := range z.Da {
			_ = zvqo
			_ = zbbj
			s += msgp.Int64Size + zvqo.Msgsize()
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
		var zjkf int32
		zjkf, err = dc.ReadInt32()
		(*z) = Zkind(zjkf)
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
		var zzpn int32
		zzpn, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zzpn)
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
