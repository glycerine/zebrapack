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
		case "FieldName":
			z.FieldName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTypeStr":
			z.FieldTypeStr, err = dc.ReadString()
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
	// map header, size 5
	// write "Zid"
	err = en.Append(0x85, 0xa3, 0x5a, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Zid)
	if err != nil {
		return
	}
	// write "FieldName"
	err = en.Append(0xa9, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldName)
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
	// map header, size 5
	// string "Zid"
	o = append(o, 0x85, 0xa3, 0x5a, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Zid)
	// string "FieldName"
	o = append(o, 0xa9, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FieldName)
	// string "FieldTypeStr"
	o = append(o, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldTypeStr)
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
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
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
		case "FieldName":
			z.FieldName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FieldTypeStr":
			z.FieldTypeStr, bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 4 + msgp.Int64Size + 10 + msgp.StringPrefixSize + len(z.FieldName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 10 + msgp.BoolSize + 5 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Schema) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
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
		case "ZebraSchemaId64":
			z.ZebraSchemaId64, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Structs":
			var zwht uint32
			zwht, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zwht) {
				z.Structs = (z.Structs)[:zwht]
			} else {
				z.Structs = make([]Struct, zwht)
			}
			for zbai := range z.Structs {
				var zhct uint32
				zhct, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zhct > 0 {
					zhct--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zbai].StructName, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Fields":
						var zcua uint32
						zcua, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Structs[zbai].Fields) >= int(zcua) {
							z.Structs[zbai].Fields = (z.Structs[zbai].Fields)[:zcua]
						} else {
							z.Structs[zbai].Fields = make([]Field, zcua)
						}
						for zcmr := range z.Structs[zbai].Fields {
							err = z.Structs[zbai].Fields[zcmr].DecodeMsg(dc)
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
	// write "ZebraSchemaId64"
	err = en.Append(0xaf, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x36, 0x34)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.ZebraSchemaId64)
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
	for zbai := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zbai].StructName)
		if err != nil {
			return
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zbai].Fields)))
		if err != nil {
			return
		}
		for zcmr := range z.Structs[zbai].Fields {
			err = z.Structs[zbai].Fields[zcmr].EncodeMsg(en)
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
	// string "ZebraSchemaId64"
	o = append(o, 0xaf, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.ZebraSchemaId64)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zbai := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zbai].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zbai].Fields)))
		for zcmr := range z.Structs[zbai].Fields {
			o, err = z.Structs[zbai].Fields[zcmr].MarshalMsg(o)
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
	var zxhx uint32
	zxhx, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
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
		case "ZebraSchemaId64":
			z.ZebraSchemaId64, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Structs":
			var zlqf uint32
			zlqf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zlqf) {
				z.Structs = (z.Structs)[:zlqf]
			} else {
				z.Structs = make([]Struct, zlqf)
			}
			for zbai := range z.Structs {
				var zdaf uint32
				zdaf, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zdaf > 0 {
					zdaf--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zbai].StructName, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Fields":
						var zpks uint32
						zpks, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Structs[zbai].Fields) >= int(zpks) {
							z.Structs[zbai].Fields = (z.Structs[zbai].Fields)[:zpks]
						} else {
							z.Structs[zbai].Fields = make([]Field, zpks)
						}
						for zcmr := range z.Structs[zbai].Fields {
							bts, err = z.Structs[zbai].Fields[zcmr].UnmarshalMsg(bts)
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
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 16 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zbai := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zbai].StructName) + 7 + msgp.ArrayHeaderSize
		for zcmr := range z.Structs[zbai].Fields {
			s += z.Structs[zbai].Fields[zcmr].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
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
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zeff) {
				z.Fields = (z.Fields)[:zeff]
			} else {
				z.Fields = make([]Field, zeff)
			}
			for zjfb := range z.Fields {
				err = z.Fields[zjfb].DecodeMsg(dc)
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
	for zjfb := range z.Fields {
		err = z.Fields[zjfb].EncodeMsg(en)
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
	for zjfb := range z.Fields {
		o, err = z.Fields[zjfb].MarshalMsg(o)
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
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
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
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zxpk) {
				z.Fields = (z.Fields)[:zxpk]
			} else {
				z.Fields = make([]Field, zxpk)
			}
			for zjfb := range z.Fields {
				bts, err = z.Fields[zjfb].UnmarshalMsg(bts)
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
	for zjfb := range z.Fields {
		s += z.Fields[zjfb].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Zprimitive) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zdnj uint8
		zdnj, err = dc.ReadUint8()
		(*z) = Zprimitive(zdnj)
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
		var zobc uint8
		zobc, bts, err = msgp.ReadUint8Bytes(bts)
		(*z) = Zprimitive(zobc)
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
