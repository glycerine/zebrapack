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
	const maxFields0zaes = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zaes uint32
	totalEncodedFields0zaes, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zaes := totalEncodedFields0zaes
	missingFieldsLeft0zaes := maxFields0zaes - totalEncodedFields0zaes

	var nextMiss0zaes int32 = -1
	var found0zaes [maxFields0zaes]bool
	var curField0zaes string

doneWithStruct0zaes:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaes > 0 || missingFieldsLeft0zaes > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaes, missingFieldsLeft0zaes, msgp.ShowFound(found0zaes[:]), decodeMsgFieldOrder0zaes)
		if encodedFieldsLeft0zaes > 0 {
			encodedFieldsLeft0zaes--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zaes = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zaes < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zaes = 0
			}
			for nextMiss0zaes < maxFields0zaes && found0zaes[nextMiss0zaes] {
				nextMiss0zaes++
			}
			if nextMiss0zaes == maxFields0zaes {
				// filled all the empty fields!
				break doneWithStruct0zaes
			}
			missingFieldsLeft0zaes--
			curField0zaes = decodeMsgFieldOrder0zaes[nextMiss0zaes]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaes)
		switch curField0zaes {
		// -- templateDecodeMsg ends here --

		case "FieldId":
			found0zaes[0] = true
			z.FieldId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found0zaes[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Zknd":
			found0zaes[2] = true
			{
				var zlge int32
				zlge, err = dc.ReadInt32()
				z.Zknd = Zkind(zlge)
			}
			if err != nil {
				panic(err)
			}
		case "Varg":
			found0zaes[3] = true
			z.Varg, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Tag":
			found0zaes[4] = true
			var zrrg uint32
			zrrg, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Tag == nil && zrrg > 0 {
				z.Tag = make(map[string]string, zrrg)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zrrg > 0 {
				zrrg--
				var zltp string
				var zguf string
				zltp, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zguf, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				z.Tag[zltp] = zguf
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zaes != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of FieldT
var decodeMsgFieldOrder0zaes = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

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
	var empty_zdfc [5]bool
	fieldsInUse_zljm := z.fieldsNotEmpty(empty_zdfc[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zljm)
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
	if !empty_zdfc[3] {
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

	if !empty_zdfc[4] {
		// write "Tag"
		err = en.Append(0xa3, 0x54, 0x61, 0x67)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Tag)))
		if err != nil {
			panic(err)
		}
		for zltp, zguf := range z.Tag {
			err = en.WriteString(zltp)
			if err != nil {
				panic(err)
			}
			err = en.WriteString(zguf)
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
		for zltp, zguf := range z.Tag {
			o = msgp.AppendString(o, zltp)
			o = msgp.AppendString(o, zguf)
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
	const maxFields1zptj = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zptj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zptj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zptj := totalEncodedFields1zptj
	missingFieldsLeft1zptj := maxFields1zptj - totalEncodedFields1zptj

	var nextMiss1zptj int32 = -1
	var found1zptj [maxFields1zptj]bool
	var curField1zptj string

doneWithStruct1zptj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zptj > 0 || missingFieldsLeft1zptj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zptj, missingFieldsLeft1zptj, msgp.ShowFound(found1zptj[:]), unmarshalMsgFieldOrder1zptj)
		if encodedFieldsLeft1zptj > 0 {
			encodedFieldsLeft1zptj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField1zptj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zptj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zptj = 0
			}
			for nextMiss1zptj < maxFields1zptj && found1zptj[nextMiss1zptj] {
				nextMiss1zptj++
			}
			if nextMiss1zptj == maxFields1zptj {
				// filled all the empty fields!
				break doneWithStruct1zptj
			}
			missingFieldsLeft1zptj--
			curField1zptj = unmarshalMsgFieldOrder1zptj[nextMiss1zptj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zptj)
		switch curField1zptj {
		// -- templateUnmarshalMsg ends here --

		case "FieldId":
			found1zptj[0] = true
			z.FieldId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found1zptj[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Zknd":
			found1zptj[2] = true
			{
				var zexd int32
				zexd, bts, err = nbs.ReadInt32Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Zknd = Zkind(zexd)
			}
		case "Varg":
			found1zptj[3] = true
			z.Varg, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Tag":
			found1zptj[4] = true
			if nbs.AlwaysNil {
				if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}

			} else {

				var zjfw uint32
				zjfw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Tag == nil && zjfw > 0 {
					z.Tag = make(map[string]string, zjfw)
				} else if len(z.Tag) > 0 {
					for key, _ := range z.Tag {
						delete(z.Tag, key)
					}
				}
				for zjfw > 0 {
					var zltp string
					var zguf string
					zjfw--
					zltp, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zguf, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Tag[zltp] = zguf
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss1zptj != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of FieldT
var unmarshalMsgFieldOrder1zptj = []string{"FieldId", "Name", "Zknd", "Varg", "Tag"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FieldT) Msgsize() (s int) {
	s = 1 + 8 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.Int32Size + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zltp, zguf := range z.Tag {
			_ = zguf
			_ = zltp
			s += msgp.StringPrefixSize + len(zltp) + msgp.StringPrefixSize + len(zguf)
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
	const maxFields2zrvc = 1

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zrvc uint32
	totalEncodedFields2zrvc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zrvc := totalEncodedFields2zrvc
	missingFieldsLeft2zrvc := maxFields2zrvc - totalEncodedFields2zrvc

	var nextMiss2zrvc int32 = -1
	var found2zrvc [maxFields2zrvc]bool
	var curField2zrvc string

doneWithStruct2zrvc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zrvc > 0 || missingFieldsLeft2zrvc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zrvc, missingFieldsLeft2zrvc, msgp.ShowFound(found2zrvc[:]), decodeMsgFieldOrder2zrvc)
		if encodedFieldsLeft2zrvc > 0 {
			encodedFieldsLeft2zrvc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zrvc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zrvc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zrvc = 0
			}
			for nextMiss2zrvc < maxFields2zrvc && found2zrvc[nextMiss2zrvc] {
				nextMiss2zrvc++
			}
			if nextMiss2zrvc == maxFields2zrvc {
				// filled all the empty fields!
				break doneWithStruct2zrvc
			}
			missingFieldsLeft2zrvc--
			curField2zrvc = decodeMsgFieldOrder2zrvc[nextMiss2zrvc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zrvc)
		switch curField2zrvc {
		// -- templateDecodeMsg ends here --

		case "IfaceId":
			found2zrvc[0] = true
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
	if nextMiss2zrvc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceInstance
var decodeMsgFieldOrder2zrvc = []string{"IfaceId"}

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
	const maxFields3zquw = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zquw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zquw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zquw := totalEncodedFields3zquw
	missingFieldsLeft3zquw := maxFields3zquw - totalEncodedFields3zquw

	var nextMiss3zquw int32 = -1
	var found3zquw [maxFields3zquw]bool
	var curField3zquw string

doneWithStruct3zquw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zquw > 0 || missingFieldsLeft3zquw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zquw, missingFieldsLeft3zquw, msgp.ShowFound(found3zquw[:]), unmarshalMsgFieldOrder3zquw)
		if encodedFieldsLeft3zquw > 0 {
			encodedFieldsLeft3zquw--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField3zquw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zquw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zquw = 0
			}
			for nextMiss3zquw < maxFields3zquw && found3zquw[nextMiss3zquw] {
				nextMiss3zquw++
			}
			if nextMiss3zquw == maxFields3zquw {
				// filled all the empty fields!
				break doneWithStruct3zquw
			}
			missingFieldsLeft3zquw--
			curField3zquw = unmarshalMsgFieldOrder3zquw[nextMiss3zquw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zquw)
		switch curField3zquw {
		// -- templateUnmarshalMsg ends here --

		case "IfaceId":
			found3zquw[0] = true
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
	if nextMiss3zquw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceInstance
var unmarshalMsgFieldOrder3zquw = []string{"IfaceId"}

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
	const maxFields4zhnc = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zhnc uint32
	totalEncodedFields4zhnc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zhnc := totalEncodedFields4zhnc
	missingFieldsLeft4zhnc := maxFields4zhnc - totalEncodedFields4zhnc

	var nextMiss4zhnc int32 = -1
	var found4zhnc [maxFields4zhnc]bool
	var curField4zhnc string

doneWithStruct4zhnc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zhnc > 0 || missingFieldsLeft4zhnc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zhnc, missingFieldsLeft4zhnc, msgp.ShowFound(found4zhnc[:]), decodeMsgFieldOrder4zhnc)
		if encodedFieldsLeft4zhnc > 0 {
			encodedFieldsLeft4zhnc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zhnc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zhnc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zhnc = 0
			}
			for nextMiss4zhnc < maxFields4zhnc && found4zhnc[nextMiss4zhnc] {
				nextMiss4zhnc++
			}
			if nextMiss4zhnc == maxFields4zhnc {
				// filled all the empty fields!
				break doneWithStruct4zhnc
			}
			missingFieldsLeft4zhnc--
			curField4zhnc = decodeMsgFieldOrder4zhnc[nextMiss4zhnc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zhnc)
		switch curField4zhnc {
		// -- templateDecodeMsg ends here --

		case "Methods":
			found4zhnc[0] = true
			var ztax uint32
			ztax, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Methods) >= int(ztax) {
				z.Methods = (z.Methods)[:ztax]
			} else {
				z.Methods = make([]MethodT, ztax)
			}
			for zmxq := range z.Methods {
				err = z.Methods[zmxq].DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
			}
		case "Deprecated":
			found4zhnc[1] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found4zhnc[2] = true
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
	if nextMiss4zhnc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of InterfaceT
var decodeMsgFieldOrder4zhnc = []string{"Methods", "Deprecated", "Comment"}

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
	for zmxq := range z.Methods {
		err = z.Methods[zmxq].EncodeMsg(en)
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
	for zmxq := range z.Methods {
		o, err = z.Methods[zmxq].MarshalMsg(o)
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
	const maxFields5zrco = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zrco uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zrco, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5zrco := totalEncodedFields5zrco
	missingFieldsLeft5zrco := maxFields5zrco - totalEncodedFields5zrco

	var nextMiss5zrco int32 = -1
	var found5zrco [maxFields5zrco]bool
	var curField5zrco string

doneWithStruct5zrco:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zrco > 0 || missingFieldsLeft5zrco > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zrco, missingFieldsLeft5zrco, msgp.ShowFound(found5zrco[:]), unmarshalMsgFieldOrder5zrco)
		if encodedFieldsLeft5zrco > 0 {
			encodedFieldsLeft5zrco--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField5zrco = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zrco < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zrco = 0
			}
			for nextMiss5zrco < maxFields5zrco && found5zrco[nextMiss5zrco] {
				nextMiss5zrco++
			}
			if nextMiss5zrco == maxFields5zrco {
				// filled all the empty fields!
				break doneWithStruct5zrco
			}
			missingFieldsLeft5zrco--
			curField5zrco = unmarshalMsgFieldOrder5zrco[nextMiss5zrco]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zrco)
		switch curField5zrco {
		// -- templateUnmarshalMsg ends here --

		case "Methods":
			found5zrco[0] = true
			if nbs.AlwaysNil {
				(z.Methods) = (z.Methods)[:0]
			} else {

				var zmna uint32
				zmna, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Methods) >= int(zmna) {
					z.Methods = (z.Methods)[:zmna]
				} else {
					z.Methods = make([]MethodT, zmna)
				}
				for zmxq := range z.Methods {
					bts, err = z.Methods[zmxq].UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
				}
			}
		case "Deprecated":
			found5zrco[1] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found5zrco[2] = true
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
	if nextMiss5zrco != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of InterfaceT
var unmarshalMsgFieldOrder5zrco = []string{"Methods", "Deprecated", "Comment"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InterfaceT) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for zmxq := range z.Methods {
		s += z.Methods[zmxq].Msgsize()
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
	const maxFields6zdft = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zdft uint32
	totalEncodedFields6zdft, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zdft := totalEncodedFields6zdft
	missingFieldsLeft6zdft := maxFields6zdft - totalEncodedFields6zdft

	var nextMiss6zdft int32 = -1
	var found6zdft [maxFields6zdft]bool
	var curField6zdft string

doneWithStruct6zdft:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zdft > 0 || missingFieldsLeft6zdft > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zdft, missingFieldsLeft6zdft, msgp.ShowFound(found6zdft[:]), decodeMsgFieldOrder6zdft)
		if encodedFieldsLeft6zdft > 0 {
			encodedFieldsLeft6zdft--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zdft = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zdft < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zdft = 0
			}
			for nextMiss6zdft < maxFields6zdft && found6zdft[nextMiss6zdft] {
				nextMiss6zdft++
			}
			if nextMiss6zdft == maxFields6zdft {
				// filled all the empty fields!
				break doneWithStruct6zdft
			}
			missingFieldsLeft6zdft--
			curField6zdft = decodeMsgFieldOrder6zdft[nextMiss6zdft]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zdft)
		switch curField6zdft {
		// -- templateDecodeMsg ends here --

		case "MethodId":
			found6zdft[0] = true
			z.MethodId, err = dc.ReadInt64()
			if err != nil {
				panic(err)
			}
		case "Name":
			found6zdft[1] = true
			z.Name, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Inputs":
			found6zdft[2] = true
			const maxFields7zklu = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields7zklu uint32
			totalEncodedFields7zklu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft7zklu := totalEncodedFields7zklu
			missingFieldsLeft7zklu := maxFields7zklu - totalEncodedFields7zklu

			var nextMiss7zklu int32 = -1
			var found7zklu [maxFields7zklu]bool
			var curField7zklu string

		doneWithStruct7zklu:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft7zklu > 0 || missingFieldsLeft7zklu > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zklu, missingFieldsLeft7zklu, msgp.ShowFound(found7zklu[:]), decodeMsgFieldOrder7zklu)
				if encodedFieldsLeft7zklu > 0 {
					encodedFieldsLeft7zklu--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField7zklu = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss7zklu < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss7zklu = 0
					}
					for nextMiss7zklu < maxFields7zklu && found7zklu[nextMiss7zklu] {
						nextMiss7zklu++
					}
					if nextMiss7zklu == maxFields7zklu {
						// filled all the empty fields!
						break doneWithStruct7zklu
					}
					missingFieldsLeft7zklu--
					curField7zklu = decodeMsgFieldOrder7zklu[nextMiss7zklu]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField7zklu)
				switch curField7zklu {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found7zklu[0] = true
					z.Inputs.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found7zklu[1] = true
					var zyrg uint32
					zyrg, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Inputs.Fld) >= int(zyrg) {
						z.Inputs.Fld = (z.Inputs.Fld)[:zyrg]
					} else {
						z.Inputs.Fld = make([]FieldT, zyrg)
					}
					for zrlk := range z.Inputs.Fld {
						err = z.Inputs.Fld[zrlk].DecodeMsg(dc)
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
			if nextMiss7zklu != -1 {
				dc.PopAlwaysNil()
			}

		case "Returns":
			found6zdft[3] = true
			const maxFields8zzvn = 2

			// -- templateDecodeMsg starts here--
			var totalEncodedFields8zzvn uint32
			totalEncodedFields8zzvn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			encodedFieldsLeft8zzvn := totalEncodedFields8zzvn
			missingFieldsLeft8zzvn := maxFields8zzvn - totalEncodedFields8zzvn

			var nextMiss8zzvn int32 = -1
			var found8zzvn [maxFields8zzvn]bool
			var curField8zzvn string

		doneWithStruct8zzvn:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft8zzvn > 0 || missingFieldsLeft8zzvn > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zzvn, missingFieldsLeft8zzvn, msgp.ShowFound(found8zzvn[:]), decodeMsgFieldOrder8zzvn)
				if encodedFieldsLeft8zzvn > 0 {
					encodedFieldsLeft8zzvn--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					curField8zzvn = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss8zzvn < 0 {
						// tell the reader to only give us Nils
						// until further notice.
						dc.PushAlwaysNil()
						nextMiss8zzvn = 0
					}
					for nextMiss8zzvn < maxFields8zzvn && found8zzvn[nextMiss8zzvn] {
						nextMiss8zzvn++
					}
					if nextMiss8zzvn == maxFields8zzvn {
						// filled all the empty fields!
						break doneWithStruct8zzvn
					}
					missingFieldsLeft8zzvn--
					curField8zzvn = decodeMsgFieldOrder8zzvn[nextMiss8zzvn]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField8zzvn)
				switch curField8zzvn {
				// -- templateDecodeMsg ends here --

				case "Nam":
					found8zzvn[0] = true
					z.Returns.Nam, err = dc.ReadString()
					if err != nil {
						panic(err)
					}
				case "Fld":
					found8zzvn[1] = true
					var zbqg uint32
					zbqg, err = dc.ReadArrayHeader()
					if err != nil {
						panic(err)
					}
					if cap(z.Returns.Fld) >= int(zbqg) {
						z.Returns.Fld = (z.Returns.Fld)[:zbqg]
					} else {
						z.Returns.Fld = make([]FieldT, zbqg)
					}
					for ztoo := range z.Returns.Fld {
						err = z.Returns.Fld[ztoo].DecodeMsg(dc)
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
			if nextMiss8zzvn != -1 {
				dc.PopAlwaysNil()
			}

		case "Deprecated":
			found6zdft[4] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				panic(err)
			}
		case "Comment":
			found6zdft[5] = true
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
	if nextMiss6zdft != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of MethodT
var decodeMsgFieldOrder6zdft = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var decodeMsgFieldOrder7zklu = []string{"Nam", "Fld"}

// fields of StructT
var decodeMsgFieldOrder8zzvn = []string{"Nam", "Fld"}

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
	for zrlk := range z.Inputs.Fld {
		err = z.Inputs.Fld[zrlk].EncodeMsg(en)
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
	for ztoo := range z.Returns.Fld {
		err = z.Returns.Fld[ztoo].EncodeMsg(en)
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
	for zrlk := range z.Inputs.Fld {
		o, err = z.Inputs.Fld[zrlk].MarshalMsg(o)
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
	for ztoo := range z.Returns.Fld {
		o, err = z.Returns.Fld[ztoo].MarshalMsg(o)
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
	const maxFields9zbap = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zbap uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zbap, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft9zbap := totalEncodedFields9zbap
	missingFieldsLeft9zbap := maxFields9zbap - totalEncodedFields9zbap

	var nextMiss9zbap int32 = -1
	var found9zbap [maxFields9zbap]bool
	var curField9zbap string

doneWithStruct9zbap:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zbap > 0 || missingFieldsLeft9zbap > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zbap, missingFieldsLeft9zbap, msgp.ShowFound(found9zbap[:]), unmarshalMsgFieldOrder9zbap)
		if encodedFieldsLeft9zbap > 0 {
			encodedFieldsLeft9zbap--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField9zbap = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zbap < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zbap = 0
			}
			for nextMiss9zbap < maxFields9zbap && found9zbap[nextMiss9zbap] {
				nextMiss9zbap++
			}
			if nextMiss9zbap == maxFields9zbap {
				// filled all the empty fields!
				break doneWithStruct9zbap
			}
			missingFieldsLeft9zbap--
			curField9zbap = unmarshalMsgFieldOrder9zbap[nextMiss9zbap]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zbap)
		switch curField9zbap {
		// -- templateUnmarshalMsg ends here --

		case "MethodId":
			found9zbap[0] = true
			z.MethodId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				panic(err)
			}
		case "Name":
			found9zbap[1] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Inputs":
			found9zbap[2] = true
			const maxFields10zdgw = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields10zdgw uint32
			if !nbs.AlwaysNil {
				totalEncodedFields10zdgw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft10zdgw := totalEncodedFields10zdgw
			missingFieldsLeft10zdgw := maxFields10zdgw - totalEncodedFields10zdgw

			var nextMiss10zdgw int32 = -1
			var found10zdgw [maxFields10zdgw]bool
			var curField10zdgw string

		doneWithStruct10zdgw:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft10zdgw > 0 || missingFieldsLeft10zdgw > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zdgw, missingFieldsLeft10zdgw, msgp.ShowFound(found10zdgw[:]), unmarshalMsgFieldOrder10zdgw)
				if encodedFieldsLeft10zdgw > 0 {
					encodedFieldsLeft10zdgw--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField10zdgw = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss10zdgw < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss10zdgw = 0
					}
					for nextMiss10zdgw < maxFields10zdgw && found10zdgw[nextMiss10zdgw] {
						nextMiss10zdgw++
					}
					if nextMiss10zdgw == maxFields10zdgw {
						// filled all the empty fields!
						break doneWithStruct10zdgw
					}
					missingFieldsLeft10zdgw--
					curField10zdgw = unmarshalMsgFieldOrder10zdgw[nextMiss10zdgw]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField10zdgw)
				switch curField10zdgw {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found10zdgw[0] = true
					z.Inputs.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found10zdgw[1] = true
					if nbs.AlwaysNil {
						(z.Inputs.Fld) = (z.Inputs.Fld)[:0]
					} else {

						var zyvt uint32
						zyvt, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Inputs.Fld) >= int(zyvt) {
							z.Inputs.Fld = (z.Inputs.Fld)[:zyvt]
						} else {
							z.Inputs.Fld = make([]FieldT, zyvt)
						}
						for zrlk := range z.Inputs.Fld {
							bts, err = z.Inputs.Fld[zrlk].UnmarshalMsg(bts)
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
			if nextMiss10zdgw != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Returns":
			found9zbap[3] = true
			const maxFields11zfoe = 2

			// -- templateUnmarshalMsg starts here--
			var totalEncodedFields11zfoe uint32
			if !nbs.AlwaysNil {
				totalEncodedFields11zfoe, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
					return
				}
			}
			encodedFieldsLeft11zfoe := totalEncodedFields11zfoe
			missingFieldsLeft11zfoe := maxFields11zfoe - totalEncodedFields11zfoe

			var nextMiss11zfoe int32 = -1
			var found11zfoe [maxFields11zfoe]bool
			var curField11zfoe string

		doneWithStruct11zfoe:
			// First fill all the encoded fields, then
			// treat the remaining, missing fields, as Nil.
			for encodedFieldsLeft11zfoe > 0 || missingFieldsLeft11zfoe > 0 {
				//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zfoe, missingFieldsLeft11zfoe, msgp.ShowFound(found11zfoe[:]), unmarshalMsgFieldOrder11zfoe)
				if encodedFieldsLeft11zfoe > 0 {
					encodedFieldsLeft11zfoe--
					field, bts, err = nbs.ReadMapKeyZC(bts)
					if err != nil {
						panic(err)
						return
					}
					curField11zfoe = msgp.UnsafeString(field)
				} else {
					//missing fields need handling
					if nextMiss11zfoe < 0 {
						// set bts to contain just mnil (0xc0)
						bts = nbs.PushAlwaysNil(bts)
						nextMiss11zfoe = 0
					}
					for nextMiss11zfoe < maxFields11zfoe && found11zfoe[nextMiss11zfoe] {
						nextMiss11zfoe++
					}
					if nextMiss11zfoe == maxFields11zfoe {
						// filled all the empty fields!
						break doneWithStruct11zfoe
					}
					missingFieldsLeft11zfoe--
					curField11zfoe = unmarshalMsgFieldOrder11zfoe[nextMiss11zfoe]
				}
				//fmt.Printf("switching on curField: '%v'\n", curField11zfoe)
				switch curField11zfoe {
				// -- templateUnmarshalMsg ends here --

				case "Nam":
					found11zfoe[0] = true
					z.Returns.Nam, bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						panic(err)
					}
				case "Fld":
					found11zfoe[1] = true
					if nbs.AlwaysNil {
						(z.Returns.Fld) = (z.Returns.Fld)[:0]
					} else {

						var ztvy uint32
						ztvy, bts, err = nbs.ReadArrayHeaderBytes(bts)
						if err != nil {
							panic(err)
						}
						if cap(z.Returns.Fld) >= int(ztvy) {
							z.Returns.Fld = (z.Returns.Fld)[:ztvy]
						} else {
							z.Returns.Fld = make([]FieldT, ztvy)
						}
						for ztoo := range z.Returns.Fld {
							bts, err = z.Returns.Fld[ztoo].UnmarshalMsg(bts)
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
			if nextMiss11zfoe != -1 {
				bts = nbs.PopAlwaysNil()
			}

		case "Deprecated":
			found9zbap[4] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Comment":
			found9zbap[5] = true
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
	if nextMiss9zbap != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of MethodT
var unmarshalMsgFieldOrder9zbap = []string{"MethodId", "Name", "Inputs", "Returns", "Deprecated", "Comment"}

// fields of StructT
var unmarshalMsgFieldOrder10zdgw = []string{"Nam", "Fld"}

// fields of StructT
var unmarshalMsgFieldOrder11zfoe = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MethodT) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + 1 + 4 + msgp.StringPrefixSize + len(z.Inputs.Nam) + 4 + msgp.ArrayHeaderSize
	for zrlk := range z.Inputs.Fld {
		s += z.Inputs.Fld[zrlk].Msgsize()
	}
	s += 8 + 1 + 4 + msgp.StringPrefixSize + len(z.Returns.Nam) + 4 + msgp.ArrayHeaderSize
	for ztoo := range z.Returns.Fld {
		s += z.Returns.Fld[ztoo].Msgsize()
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
	const maxFields12zhjt = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields12zhjt uint32
	totalEncodedFields12zhjt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zhjt := totalEncodedFields12zhjt
	missingFieldsLeft12zhjt := maxFields12zhjt - totalEncodedFields12zhjt

	var nextMiss12zhjt int32 = -1
	var found12zhjt [maxFields12zhjt]bool
	var curField12zhjt string

doneWithStruct12zhjt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zhjt > 0 || missingFieldsLeft12zhjt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zhjt, missingFieldsLeft12zhjt, msgp.ShowFound(found12zhjt[:]), decodeMsgFieldOrder12zhjt)
		if encodedFieldsLeft12zhjt > 0 {
			encodedFieldsLeft12zhjt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField12zhjt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss12zhjt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zhjt = 0
			}
			for nextMiss12zhjt < maxFields12zhjt && found12zhjt[nextMiss12zhjt] {
				nextMiss12zhjt++
			}
			if nextMiss12zhjt == maxFields12zhjt {
				// filled all the empty fields!
				break doneWithStruct12zhjt
			}
			missingFieldsLeft12zhjt--
			curField12zhjt = decodeMsgFieldOrder12zhjt[nextMiss12zhjt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zhjt)
		switch curField12zhjt {
		// -- templateDecodeMsg ends here --

		case "PkgPath":
			found12zhjt[0] = true
			z.PkgPath, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Structs":
			found12zhjt[1] = true
			var zkxe uint32
			zkxe, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Structs == nil && zkxe > 0 {
				z.Structs = make(map[int64]StructT, zkxe)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zkxe > 0 {
				zkxe--
				var zpcm int64
				var zgmi StructT
				zpcm, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				const maxFields13zxpp = 2

				// -- templateDecodeMsg starts here--
				var totalEncodedFields13zxpp uint32
				totalEncodedFields13zxpp, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft13zxpp := totalEncodedFields13zxpp
				missingFieldsLeft13zxpp := maxFields13zxpp - totalEncodedFields13zxpp

				var nextMiss13zxpp int32 = -1
				var found13zxpp [maxFields13zxpp]bool
				var curField13zxpp string

			doneWithStruct13zxpp:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft13zxpp > 0 || missingFieldsLeft13zxpp > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zxpp, missingFieldsLeft13zxpp, msgp.ShowFound(found13zxpp[:]), decodeMsgFieldOrder13zxpp)
					if encodedFieldsLeft13zxpp > 0 {
						encodedFieldsLeft13zxpp--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						curField13zxpp = msgp.UnsafeString(field)
					} else {
						//missing fields need handling
						if nextMiss13zxpp < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss13zxpp = 0
						}
						for nextMiss13zxpp < maxFields13zxpp && found13zxpp[nextMiss13zxpp] {
							nextMiss13zxpp++
						}
						if nextMiss13zxpp == maxFields13zxpp {
							// filled all the empty fields!
							break doneWithStruct13zxpp
						}
						missingFieldsLeft13zxpp--
						curField13zxpp = decodeMsgFieldOrder13zxpp[nextMiss13zxpp]
					}
					//fmt.Printf("switching on curField: '%v'\n", curField13zxpp)
					switch curField13zxpp {
					// -- templateDecodeMsg ends here --

					case "Nam":
						found13zxpp[0] = true
						zgmi.Nam, err = dc.ReadString()
						if err != nil {
							panic(err)
						}
					case "Fld":
						found13zxpp[1] = true
						var zrdn uint32
						zrdn, err = dc.ReadArrayHeader()
						if err != nil {
							panic(err)
						}
						if cap(zgmi.Fld) >= int(zrdn) {
							zgmi.Fld = (zgmi.Fld)[:zrdn]
						} else {
							zgmi.Fld = make([]FieldT, zrdn)
						}
						for zgkl := range zgmi.Fld {
							err = zgmi.Fld[zgkl].DecodeMsg(dc)
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
				if nextMiss13zxpp != -1 {
					dc.PopAlwaysNil()
				}

				z.Structs[zpcm] = zgmi
			}
		case "Interfaces":
			found12zhjt[2] = true
			var zvhd uint32
			zvhd, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Interfaces == nil && zvhd > 0 {
				z.Interfaces = make(map[int64]InterfaceT, zvhd)
			} else if len(z.Interfaces) > 0 {
				for key, _ := range z.Interfaces {
					delete(z.Interfaces, key)
				}
			}
			for zvhd > 0 {
				zvhd--
				var zxmj int64
				var zzif InterfaceT
				zxmj, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zzif.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Interfaces[zxmj] = zzif
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss12zhjt != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of SchemaT
var decodeMsgFieldOrder12zhjt = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var decodeMsgFieldOrder13zxpp = []string{"Nam", "Fld"}

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
	for zpcm, zgmi := range z.Structs {
		err = en.WriteInt64(zpcm)
		if err != nil {
			panic(err)
		}
		// map header, size 2
		// write "Nam"
		err = en.Append(0x82, 0xa3, 0x4e, 0x61, 0x6d)
		if err != nil {
			return err
		}
		err = en.WriteString(zgmi.Nam)
		if err != nil {
			panic(err)
		}
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(zgmi.Fld)))
		if err != nil {
			panic(err)
		}
		for zgkl := range zgmi.Fld {
			err = zgmi.Fld[zgkl].EncodeMsg(en)
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
	for zxmj, zzif := range z.Interfaces {
		err = en.WriteInt64(zxmj)
		if err != nil {
			panic(err)
		}
		err = zzif.EncodeMsg(en)
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
	for zpcm, zgmi := range z.Structs {
		o = msgp.AppendInt64(o, zpcm)
		// map header, size 2
		// string "Nam"
		o = append(o, 0x82, 0xa3, 0x4e, 0x61, 0x6d)
		o = msgp.AppendString(o, zgmi.Nam)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(zgmi.Fld)))
		for zgkl := range zgmi.Fld {
			o, err = zgmi.Fld[zgkl].MarshalMsg(o)
			if err != nil {
				panic(err)
			}
		}
	}
	// string "Interfaces"
	o = append(o, 0xaa, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Interfaces)))
	for zxmj, zzif := range z.Interfaces {
		o = msgp.AppendInt64(o, zxmj)
		o, err = zzif.MarshalMsg(o)
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
	const maxFields14zotv = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields14zotv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields14zotv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft14zotv := totalEncodedFields14zotv
	missingFieldsLeft14zotv := maxFields14zotv - totalEncodedFields14zotv

	var nextMiss14zotv int32 = -1
	var found14zotv [maxFields14zotv]bool
	var curField14zotv string

doneWithStruct14zotv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zotv > 0 || missingFieldsLeft14zotv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zotv, missingFieldsLeft14zotv, msgp.ShowFound(found14zotv[:]), unmarshalMsgFieldOrder14zotv)
		if encodedFieldsLeft14zotv > 0 {
			encodedFieldsLeft14zotv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField14zotv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss14zotv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss14zotv = 0
			}
			for nextMiss14zotv < maxFields14zotv && found14zotv[nextMiss14zotv] {
				nextMiss14zotv++
			}
			if nextMiss14zotv == maxFields14zotv {
				// filled all the empty fields!
				break doneWithStruct14zotv
			}
			missingFieldsLeft14zotv--
			curField14zotv = unmarshalMsgFieldOrder14zotv[nextMiss14zotv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zotv)
		switch curField14zotv {
		// -- templateUnmarshalMsg ends here --

		case "PkgPath":
			found14zotv[0] = true
			z.PkgPath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Structs":
			found14zotv[1] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zmst uint32
				zmst, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Structs == nil && zmst > 0 {
					z.Structs = make(map[int64]StructT, zmst)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zmst > 0 {
					var zpcm int64
					var zgmi StructT
					zmst--
					zpcm, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					const maxFields15zrks = 2

					// -- templateUnmarshalMsg starts here--
					var totalEncodedFields15zrks uint32
					if !nbs.AlwaysNil {
						totalEncodedFields15zrks, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							panic(err)
							return
						}
					}
					encodedFieldsLeft15zrks := totalEncodedFields15zrks
					missingFieldsLeft15zrks := maxFields15zrks - totalEncodedFields15zrks

					var nextMiss15zrks int32 = -1
					var found15zrks [maxFields15zrks]bool
					var curField15zrks string

				doneWithStruct15zrks:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft15zrks > 0 || missingFieldsLeft15zrks > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zrks, missingFieldsLeft15zrks, msgp.ShowFound(found15zrks[:]), unmarshalMsgFieldOrder15zrks)
						if encodedFieldsLeft15zrks > 0 {
							encodedFieldsLeft15zrks--
							field, bts, err = nbs.ReadMapKeyZC(bts)
							if err != nil {
								panic(err)
								return
							}
							curField15zrks = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss15zrks < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss15zrks = 0
							}
							for nextMiss15zrks < maxFields15zrks && found15zrks[nextMiss15zrks] {
								nextMiss15zrks++
							}
							if nextMiss15zrks == maxFields15zrks {
								// filled all the empty fields!
								break doneWithStruct15zrks
							}
							missingFieldsLeft15zrks--
							curField15zrks = unmarshalMsgFieldOrder15zrks[nextMiss15zrks]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField15zrks)
						switch curField15zrks {
						// -- templateUnmarshalMsg ends here --

						case "Nam":
							found15zrks[0] = true
							zgmi.Nam, bts, err = nbs.ReadStringBytes(bts)

							if err != nil {
								panic(err)
							}
						case "Fld":
							found15zrks[1] = true
							if nbs.AlwaysNil {
								(zgmi.Fld) = (zgmi.Fld)[:0]
							} else {

								var zolm uint32
								zolm, bts, err = nbs.ReadArrayHeaderBytes(bts)
								if err != nil {
									panic(err)
								}
								if cap(zgmi.Fld) >= int(zolm) {
									zgmi.Fld = (zgmi.Fld)[:zolm]
								} else {
									zgmi.Fld = make([]FieldT, zolm)
								}
								for zgkl := range zgmi.Fld {
									bts, err = zgmi.Fld[zgkl].UnmarshalMsg(bts)
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
					if nextMiss15zrks != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Structs[zpcm] = zgmi
				}
			}
		case "Interfaces":
			found14zotv[2] = true
			if nbs.AlwaysNil {
				if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}

			} else {

				var zaaf uint32
				zaaf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Interfaces == nil && zaaf > 0 {
					z.Interfaces = make(map[int64]InterfaceT, zaaf)
				} else if len(z.Interfaces) > 0 {
					for key, _ := range z.Interfaces {
						delete(z.Interfaces, key)
					}
				}
				for zaaf > 0 {
					var zxmj int64
					var zzif InterfaceT
					zaaf--
					zxmj, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zzif.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Interfaces[zxmj] = zzif
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss14zotv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of SchemaT
var unmarshalMsgFieldOrder14zotv = []string{"PkgPath", "Structs", "Interfaces"}

// fields of StructT
var unmarshalMsgFieldOrder15zrks = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SchemaT) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.MapHeaderSize
	if z.Structs != nil {
		for zpcm, zgmi := range z.Structs {
			_ = zgmi
			_ = zpcm
			s += msgp.Int64Size + 1 + 4 + msgp.StringPrefixSize + len(zgmi.Nam) + 4 + msgp.ArrayHeaderSize
			for zgkl := range zgmi.Fld {
				s += zgmi.Fld[zgkl].Msgsize()
			}
		}
	}
	s += 11 + msgp.MapHeaderSize
	if z.Interfaces != nil {
		for zxmj, zzif := range z.Interfaces {
			_ = zzif
			_ = zxmj
			s += msgp.Int64Size + zzif.Msgsize()
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
	const maxFields16zsac = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields16zsac uint32
	totalEncodedFields16zsac, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft16zsac := totalEncodedFields16zsac
	missingFieldsLeft16zsac := maxFields16zsac - totalEncodedFields16zsac

	var nextMiss16zsac int32 = -1
	var found16zsac [maxFields16zsac]bool
	var curField16zsac string

doneWithStruct16zsac:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft16zsac > 0 || missingFieldsLeft16zsac > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft16zsac, missingFieldsLeft16zsac, msgp.ShowFound(found16zsac[:]), decodeMsgFieldOrder16zsac)
		if encodedFieldsLeft16zsac > 0 {
			encodedFieldsLeft16zsac--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField16zsac = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss16zsac < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss16zsac = 0
			}
			for nextMiss16zsac < maxFields16zsac && found16zsac[nextMiss16zsac] {
				nextMiss16zsac++
			}
			if nextMiss16zsac == maxFields16zsac {
				// filled all the empty fields!
				break doneWithStruct16zsac
			}
			missingFieldsLeft16zsac--
			curField16zsac = decodeMsgFieldOrder16zsac[nextMiss16zsac]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField16zsac)
		switch curField16zsac {
		// -- templateDecodeMsg ends here --

		case "Nam":
			found16zsac[0] = true
			z.Nam, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case "Fld":
			found16zsac[1] = true
			var ziwo uint32
			ziwo, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Fld) >= int(ziwo) {
				z.Fld = (z.Fld)[:ziwo]
			} else {
				z.Fld = make([]FieldT, ziwo)
			}
			for zvhf := range z.Fld {
				err = z.Fld[zvhf].DecodeMsg(dc)
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
	if nextMiss16zsac != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of StructT
var decodeMsgFieldOrder16zsac = []string{"Nam", "Fld"}

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
	for zvhf := range z.Fld {
		err = z.Fld[zvhf].EncodeMsg(en)
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
	for zvhf := range z.Fld {
		o, err = z.Fld[zvhf].MarshalMsg(o)
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
	const maxFields17zykg = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields17zykg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields17zykg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft17zykg := totalEncodedFields17zykg
	missingFieldsLeft17zykg := maxFields17zykg - totalEncodedFields17zykg

	var nextMiss17zykg int32 = -1
	var found17zykg [maxFields17zykg]bool
	var curField17zykg string

doneWithStruct17zykg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft17zykg > 0 || missingFieldsLeft17zykg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft17zykg, missingFieldsLeft17zykg, msgp.ShowFound(found17zykg[:]), unmarshalMsgFieldOrder17zykg)
		if encodedFieldsLeft17zykg > 0 {
			encodedFieldsLeft17zykg--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField17zykg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss17zykg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss17zykg = 0
			}
			for nextMiss17zykg < maxFields17zykg && found17zykg[nextMiss17zykg] {
				nextMiss17zykg++
			}
			if nextMiss17zykg == maxFields17zykg {
				// filled all the empty fields!
				break doneWithStruct17zykg
			}
			missingFieldsLeft17zykg--
			curField17zykg = unmarshalMsgFieldOrder17zykg[nextMiss17zykg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField17zykg)
		switch curField17zykg {
		// -- templateUnmarshalMsg ends here --

		case "Nam":
			found17zykg[0] = true
			z.Nam, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case "Fld":
			found17zykg[1] = true
			if nbs.AlwaysNil {
				(z.Fld) = (z.Fld)[:0]
			} else {

				var zchd uint32
				zchd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Fld) >= int(zchd) {
					z.Fld = (z.Fld)[:zchd]
				} else {
					z.Fld = make([]FieldT, zchd)
				}
				for zvhf := range z.Fld {
					bts, err = z.Fld[zvhf].UnmarshalMsg(bts)
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
	if nextMiss17zykg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of StructT
var unmarshalMsgFieldOrder17zykg = []string{"Nam", "Fld"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructT) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Nam) + 4 + msgp.ArrayHeaderSize
	for zvhf := range z.Fld {
		s += z.Fld[zvhf].Msgsize()
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
		var zbac int64
		zbac, err = dc.ReadInt64()
		(*z) = StructTypeId(zbac)
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
		var zypj int64
		zypj, bts, err = nbs.ReadInt64Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = StructTypeId(zypj)
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
	const maxFields18zqfb = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields18zqfb uint32
	totalEncodedFields18zqfb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft18zqfb := totalEncodedFields18zqfb
	missingFieldsLeft18zqfb := maxFields18zqfb - totalEncodedFields18zqfb

	var nextMiss18zqfb int32 = -1
	var found18zqfb [maxFields18zqfb]bool
	var curField18zqfb string

doneWithStruct18zqfb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft18zqfb > 0 || missingFieldsLeft18zqfb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft18zqfb, missingFieldsLeft18zqfb, msgp.ShowFound(found18zqfb[:]), decodeMsgFieldOrder18zqfb)
		if encodedFieldsLeft18zqfb > 0 {
			encodedFieldsLeft18zqfb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField18zqfb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss18zqfb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss18zqfb = 0
			}
			for nextMiss18zqfb < maxFields18zqfb && found18zqfb[nextMiss18zqfb] {
				nextMiss18zqfb++
			}
			if nextMiss18zqfb == maxFields18zqfb {
				// filled all the empty fields!
				break doneWithStruct18zqfb
			}
			missingFieldsLeft18zqfb--
			curField18zqfb = decodeMsgFieldOrder18zqfb[nextMiss18zqfb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField18zqfb)
		switch curField18zqfb {
		// -- templateDecodeMsg ends here --

		case "Id":
			found18zqfb[0] = true
			{
				var zzce int64
				zzce, err = dc.ReadInt64()
				z.Id = StructTypeId(zzce)
			}
			if err != nil {
				panic(err)
			}
		case "Da":
			found18zqfb[1] = true
			var zoec uint32
			zoec, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Da == nil && zoec > 0 {
				z.Da = make(map[int64]msgp.Raw, zoec)
			} else if len(z.Da) > 0 {
				for key, _ := range z.Da {
					delete(z.Da, key)
				}
			}
			for zoec > 0 {
				zoec--
				var zkia int64
				var zpjo msgp.Raw
				zkia, err = dc.ReadInt64()
				if err != nil {
					panic(err)
				}
				err = zpjo.DecodeMsg(dc)
				if err != nil {
					panic(err)
				}
				z.Da[zkia] = zpjo
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss18zqfb != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	return
}

// fields of ZStructPacket
var decodeMsgFieldOrder18zqfb = []string{"Id", "Da"}

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
	for zkia, zpjo := range z.Da {
		err = en.WriteInt64(zkia)
		if err != nil {
			panic(err)
		}
		err = zpjo.EncodeMsg(en)
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
	for zkia, zpjo := range z.Da {
		o = msgp.AppendInt64(o, zkia)
		o, err = zpjo.MarshalMsg(o)
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
	const maxFields19zifq = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields19zifq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields19zifq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft19zifq := totalEncodedFields19zifq
	missingFieldsLeft19zifq := maxFields19zifq - totalEncodedFields19zifq

	var nextMiss19zifq int32 = -1
	var found19zifq [maxFields19zifq]bool
	var curField19zifq string

doneWithStruct19zifq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft19zifq > 0 || missingFieldsLeft19zifq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft19zifq, missingFieldsLeft19zifq, msgp.ShowFound(found19zifq[:]), unmarshalMsgFieldOrder19zifq)
		if encodedFieldsLeft19zifq > 0 {
			encodedFieldsLeft19zifq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				panic(err)
				return
			}
			curField19zifq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss19zifq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss19zifq = 0
			}
			for nextMiss19zifq < maxFields19zifq && found19zifq[nextMiss19zifq] {
				nextMiss19zifq++
			}
			if nextMiss19zifq == maxFields19zifq {
				// filled all the empty fields!
				break doneWithStruct19zifq
			}
			missingFieldsLeft19zifq--
			curField19zifq = unmarshalMsgFieldOrder19zifq[nextMiss19zifq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField19zifq)
		switch curField19zifq {
		// -- templateUnmarshalMsg ends here --

		case "Id":
			found19zifq[0] = true
			{
				var zjkj int64
				zjkj, bts, err = nbs.ReadInt64Bytes(bts)

				if err != nil {
					panic(err)
				}
				z.Id = StructTypeId(zjkj)
			}
		case "Da":
			found19zifq[1] = true
			if nbs.AlwaysNil {
				if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}

			} else {

				var zquf uint32
				zquf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Da == nil && zquf > 0 {
					z.Da = make(map[int64]msgp.Raw, zquf)
				} else if len(z.Da) > 0 {
					for key, _ := range z.Da {
						delete(z.Da, key)
					}
				}
				for zquf > 0 {
					var zkia int64
					var zpjo msgp.Raw
					zquf--
					zkia, bts, err = nbs.ReadInt64Bytes(bts)
					if err != nil {
						panic(err)
					}
					bts, err = zpjo.UnmarshalMsg(bts)
					if err != nil {
						panic(err)
					}
					if err != nil {
						panic(err)
					}
					z.Da[zkia] = zpjo
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss19zifq != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	return
}

// fields of ZStructPacket
var unmarshalMsgFieldOrder19zifq = []string{"Id", "Da"}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ZStructPacket) Msgsize() (s int) {
	s = 1 + 3 + msgp.Int64Size + 3 + msgp.MapHeaderSize
	if z.Da != nil {
		for zkia, zpjo := range z.Da {
			_ = zpjo
			_ = zkia
			s += msgp.Int64Size + zpjo.Msgsize()
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
		var zgkv int32
		zgkv, err = dc.ReadInt32()
		(*z) = Zkind(zgkv)
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
		var zrwd int32
		zrwd, bts, err = nbs.ReadInt32Bytes(bts)

		if err != nil {
			panic(err)
		}
		(*z) = Zkind(zrwd)
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
