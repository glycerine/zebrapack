package zebra

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Field) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Zid":
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "FieldGoName":
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTagName":
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTypeStr":
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldCategory":
			{
				var zbzg uint8
				zbzg, err = dc.ReadUint8()
				z.FieldCategory = Zcat(zbzg)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive":
			{
				var zbai uint8
				zbai, err = dc.ReadUint8()
				z.FieldPrimitive = Zprimitive(zbai)
			}
			if err != nil {
				return
			}
		case "OmitEmpty":
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Skip":
			z.Skip, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Field) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "Zid"
	err = en.Append(0x88, 0xa3, 0x5a, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Zid)
	if err != nil {
		return
	}
	// write "FieldGoName"
	err = en.Append(0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldGoName)
	if err != nil {
		return
	}
	// write "FieldTagName"
	err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldTagName)
	if err != nil {
		return
	}
	// write "FieldTypeStr"
	err = en.Append(0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldTypeStr)
	if err != nil {
		return
	}
	// write "FieldCategory"
	err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteUint8(uint8(z.FieldCategory))
	if err != nil {
		return
	}
	// write "FieldPrimitive"
	err = en.Append(0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint8(uint8(z.FieldPrimitive))
	if err != nil {
		return
	}
	// write "OmitEmpty"
	err = en.Append(0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.OmitEmpty)
	if err != nil {
		return
	}
	// write "Skip"
	err = en.Append(0xa4, 0x53, 0x6b, 0x69, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Skip)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Field) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "Zid"
	o = append(o, 0x88, 0xa3, 0x5a, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Zid)
	// string "FieldGoName"
	o = append(o, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FieldGoName)
	// string "FieldTagName"
	o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FieldTagName)
	// string "FieldTypeStr"
	o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldTypeStr)
	// string "FieldCategory"
	o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	o = msgp.AppendUint8(o, uint8(z.FieldCategory))
	// string "FieldPrimitive"
	o = append(o, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendUint8(o, uint8(z.FieldPrimitive))
	// string "OmitEmpty"
	o = append(o, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
	o = msgp.AppendBool(o, z.OmitEmpty)
	// string "Skip"
	o = append(o, 0xa4, 0x53, 0x6b, 0x69, 0x70)
	o = msgp.AppendBool(o, z.Skip)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Field) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Zid":
			z.Zid, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "FieldGoName":
			z.FieldGoName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FieldTagName":
			z.FieldTagName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FieldTypeStr":
			z.FieldTypeStr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FieldCategory":
			{
				var zajw uint8
				zajw, bts, err = msgp.ReadUint8Bytes(bts)
				z.FieldCategory = Zcat(zajw)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive":
			{
				var zwht uint8
				zwht, bts, err = msgp.ReadUint8Bytes(bts)
				z.FieldPrimitive = Zprimitive(zwht)
			}
			if err != nil {
				return
			}
		case "OmitEmpty":
			z.OmitEmpty, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Skip":
			z.Skip, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Field) Msgsize() (s int) {
	s = 1 + 4 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.FieldGoName) + 13 + msgp.StringPrefixSize + len(z.FieldTagName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 14 + msgp.Uint8Size + 15 + msgp.Uint8Size + 10 + msgp.BoolSize + 5 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Schema) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "SourcePath":
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ZebraSchemaId":
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Structs":
			var zlqf uint32
			zlqf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zlqf) {
				z.Structs = (z.Structs)[:zlqf]
			} else {
				z.Structs = make([]Struct, zlqf)
			}
			for zhct := range z.Structs {
				var zdaf uint32
				zdaf, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zdaf > 0 {
					zdaf--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zhct].StructName, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Fields":
						var zpks uint32
						zpks, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Structs[zhct].Fields) >= int(zpks) {
							z.Structs[zhct].Fields = (z.Structs[zhct].Fields)[:zpks]
						} else {
							z.Structs[zhct].Fields = make([]Field, zpks)
						}
						for zcua := range z.Structs[zhct].Fields {
							err = z.Structs[zhct].Fields[zcua].DecodeMsg(dc)
							if err != nil {
								return
							}
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Schema) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "SourcePath"
	err = en.Append(0x83, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePath)
	if err != nil {
		return
	}
	// write "ZebraSchemaId"
	err = en.Append(0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.ZebraSchemaId)
	if err != nil {
		return
	}
	// write "Structs"
	err = en.Append(0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Structs)))
	if err != nil {
		return
	}
	for zhct := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zhct].StructName)
		if err != nil {
			return
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zhct].Fields)))
		if err != nil {
			return
		}
		for zcua := range z.Structs[zhct].Fields {
			err = z.Structs[zhct].Fields[zcua].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Schema) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "SourcePath"
	o = append(o, 0x83, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.SourcePath)
	// string "ZebraSchemaId"
	o = append(o, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.ZebraSchemaId)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zhct := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zhct].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zhct].Fields)))
		for zcua := range z.Structs[zhct].Fields {
			o, err = z.Structs[zhct].Fields[zcua].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Schema) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "SourcePath":
			z.SourcePath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ZebraSchemaId":
			z.ZebraSchemaId, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Structs":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zcxo) {
				z.Structs = (z.Structs)[:zcxo]
			} else {
				z.Structs = make([]Struct, zcxo)
			}
			for zhct := range z.Structs {
				var zeff uint32
				zeff, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zeff > 0 {
					zeff--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zhct].StructName, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Fields":
						var zrsw uint32
						zrsw, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Structs[zhct].Fields) >= int(zrsw) {
							z.Structs[zhct].Fields = (z.Structs[zhct].Fields)[:zrsw]
						} else {
							z.Structs[zhct].Fields = make([]Field, zrsw)
						}
						for zcua := range z.Structs[zhct].Fields {
							bts, err = z.Structs[zhct].Fields[zcua].UnmarshalMsg(bts)
							if err != nil {
								return
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zhct := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zhct].StructName) + 7 + msgp.ArrayHeaderSize
		for zcua := range z.Structs[zhct].Fields {
			s += z.Structs[zhct].Fields[zcua].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdnj uint32
	zdnj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdnj > 0 {
		zdnj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "StructName":
			z.StructName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Fields":
			var zobc uint32
			zobc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zobc) {
				z.Fields = (z.Fields)[:zobc]
			} else {
				z.Fields = make([]Field, zobc)
			}
			for zxpk := range z.Fields {
				err = z.Fields[zxpk].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Struct) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "StructName"
	err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StructName)
	if err != nil {
		return
	}
	// write "Fields"
	err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fields)))
	if err != nil {
		return
	}
	for zxpk := range z.Fields {
		err = z.Fields[zxpk].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Struct) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "StructName"
	o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.StructName)
	// string "Fields"
	o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fields)))
	for zxpk := range z.Fields {
		o, err = z.Fields[zxpk].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Struct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsnv uint32
	zsnv, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsnv > 0 {
		zsnv--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "StructName":
			z.StructName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Fields":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zkgt) {
				z.Fields = (z.Fields)[:zkgt]
			} else {
				z.Fields = make([]Field, zkgt)
			}
			for zxpk := range z.Fields {
				bts, err = z.Fields[zxpk].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 7 + msgp.ArrayHeaderSize
	for zxpk := range z.Fields {
		s += z.Fields[zxpk].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Zcat) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zema uint8
		zema, err = dc.ReadUint8()
		(*z) = Zcat(zema)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Zcat) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zcat) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zcat) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zpez uint8
		zpez, bts, err = msgp.ReadUint8Bytes(bts)
		(*z) = Zcat(zpez)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zcat) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Zprimitive) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zqke uint8
		zqke, err = dc.ReadUint8()
		(*z) = Zprimitive(zqke)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Zprimitive) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zprimitive) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zprimitive) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zqyh uint8
		zqyh, bts, err = msgp.ReadUint8Bytes(bts)
		(*z) = Zprimitive(zqyh)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zprimitive) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}
