package zebra

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import "github.com/tinylib/msgp/msgp"

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
				var zbzg uint64
				zbzg, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zbzg)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive":
			{
				var zbai uint64
				zbai, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zbai)
			}
			if err != nil {
				return
			}
		case "FieldFullType":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.FieldFullType = nil
			} else {
				if z.FieldFullType == nil {
					z.FieldFullType = new(Ztype)
				}
				err = z.FieldFullType.DecodeMsg(dc)
				if err != nil {
					return
				}
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
		case "Deprecated":
			z.Deprecated, err = dc.ReadBool()
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
	// map header, size 10
	// write "Zid"
	err = en.Append(0x8a, 0xa3, 0x5a, 0x69, 0x64)
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
	err = en.WriteUint64(uint64(z.FieldCategory))
	if err != nil {
		return
	}
	// write "FieldPrimitive"
	err = en.Append(0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldPrimitive))
	if err != nil {
		return
	}
	// write "FieldFullType"
	err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	if z.FieldFullType == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.FieldFullType.EncodeMsg(en)
		if err != nil {
			return
		}
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
	// write "Deprecated"
	err = en.Append(0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Deprecated)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Field) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "Zid"
	o = append(o, 0x8a, 0xa3, 0x5a, 0x69, 0x64)
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
	o = msgp.AppendUint64(o, uint64(z.FieldCategory))
	// string "FieldPrimitive"
	o = append(o, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65)
	o = msgp.AppendUint64(o, uint64(z.FieldPrimitive))
	// string "FieldFullType"
	o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65)
	if z.FieldFullType == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.FieldFullType.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "OmitEmpty"
	o = append(o, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
	o = msgp.AppendBool(o, z.OmitEmpty)
	// string "Skip"
	o = append(o, 0xa4, 0x53, 0x6b, 0x69, 0x70)
	o = msgp.AppendBool(o, z.Skip)
	// string "Deprecated"
	o = append(o, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Deprecated)
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
				var zajw uint64
				zajw, bts, err = msgp.ReadUint64Bytes(bts)
				z.FieldCategory = Zkind(zajw)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive":
			{
				var zwht uint64
				zwht, bts, err = msgp.ReadUint64Bytes(bts)
				z.FieldPrimitive = Zkind(zwht)
			}
			if err != nil {
				return
			}
		case "FieldFullType":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.FieldFullType = nil
			} else {
				if z.FieldFullType == nil {
					z.FieldFullType = new(Ztype)
				}
				bts, err = z.FieldFullType.UnmarshalMsg(bts)
				if err != nil {
					return
				}
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
		case "Deprecated":
			z.Deprecated, bts, err = msgp.ReadBoolBytes(bts)
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
	s = 1 + 4 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.FieldGoName) + 13 + msgp.StringPrefixSize + len(z.FieldTagName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 14 + msgp.Uint64Size + 15 + msgp.Uint64Size + 14
	if z.FieldFullType == nil {
		s += msgp.NilSize
	} else {
		s += z.FieldFullType.Msgsize()
	}
	s += 10 + msgp.BoolSize + 5 + msgp.BoolSize + 11 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *KS) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
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
		case "Kind":
			{
				var zcua uint64
				zcua, err = dc.ReadUint64()
				z.Kind = Zkind(zcua)
			}
			if err != nil {
				return
			}
		case "Str":
			z.Str, err = dc.ReadString()
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
func (z KS) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Kind"
	err = en.Append(0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Str)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z KS) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Kind"
	o = append(o, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KS) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Kind":
			{
				var zlqf uint64
				zlqf, bts, err = msgp.ReadUint64Bytes(bts)
				z.Kind = Zkind(zlqf)
			}
			if err != nil {
				return
			}
		case "Str":
			z.Str, bts, err = msgp.ReadStringBytes(bts)
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
func (z KS) Msgsize() (s int) {
	s = 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Str)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Schema) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
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
		case "SourcePackage":
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ZebraSchemaId":
			z.ZebraSchemaId, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Structs":
			var zcxo uint32
			zcxo, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zcxo) {
				z.Structs = (z.Structs)[:zcxo]
			} else {
				z.Structs = make([]Struct, zcxo)
			}
			for zdaf := range z.Structs {
				var zeff uint32
				zeff, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zeff > 0 {
					zeff--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zdaf].StructName, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Fields":
						var zrsw uint32
						zrsw, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Structs[zdaf].Fields) >= int(zrsw) {
							z.Structs[zdaf].Fields = (z.Structs[zdaf].Fields)[:zrsw]
						} else {
							z.Structs[zdaf].Fields = make([]Field, zrsw)
						}
						for zpks := range z.Structs[zdaf].Fields {
							err = z.Structs[zdaf].Fields[zpks].DecodeMsg(dc)
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
	// map header, size 4
	// write "SourcePath"
	err = en.Append(0x84, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePath)
	if err != nil {
		return
	}
	// write "SourcePackage"
	err = en.Append(0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePackage)
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
	for zdaf := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zdaf].StructName)
		if err != nil {
			return
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zdaf].Fields)))
		if err != nil {
			return
		}
		for zpks := range z.Structs[zdaf].Fields {
			err = z.Structs[zdaf].Fields[zpks].EncodeMsg(en)
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
	// map header, size 4
	// string "SourcePath"
	o = append(o, 0x84, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.SourcePath)
	// string "SourcePackage"
	o = append(o, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.SourcePackage)
	// string "ZebraSchemaId"
	o = append(o, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64)
	o = msgp.AppendInt64(o, z.ZebraSchemaId)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zdaf := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zdaf].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zdaf].Fields)))
		for zpks := range z.Structs[zdaf].Fields {
			o, err = z.Structs[zdaf].Fields[zpks].MarshalMsg(o)
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
	var zxpk uint32
	zxpk, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zxpk > 0 {
		zxpk--
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
		case "SourcePackage":
			z.SourcePackage, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ZebraSchemaId":
			z.ZebraSchemaId, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Structs":
			var zdnj uint32
			zdnj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zdnj) {
				z.Structs = (z.Structs)[:zdnj]
			} else {
				z.Structs = make([]Struct, zdnj)
			}
			for zdaf := range z.Structs {
				var zobc uint32
				zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zobc > 0 {
					zobc--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zdaf].StructName, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Fields":
						var zsnv uint32
						zsnv, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Structs[zdaf].Fields) >= int(zsnv) {
							z.Structs[zdaf].Fields = (z.Structs[zdaf].Fields)[:zsnv]
						} else {
							z.Structs[zdaf].Fields = make([]Field, zsnv)
						}
						for zpks := range z.Structs[zdaf].Fields {
							bts, err = z.Structs[zdaf].Fields[zpks].UnmarshalMsg(bts)
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
	s = 1 + 11 + msgp.StringPrefixSize + len(z.SourcePath) + 14 + msgp.StringPrefixSize + len(z.SourcePackage) + 14 + msgp.Int64Size + 8 + msgp.ArrayHeaderSize
	for zdaf := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zdaf].StructName) + 7 + msgp.ArrayHeaderSize
		for zpks := range z.Structs[zdaf].Fields {
			s += z.Structs[zdaf].Fields[zpks].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
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
			var zpez uint32
			zpez, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zpez) {
				z.Fields = (z.Fields)[:zpez]
			} else {
				z.Fields = make([]Field, zpez)
			}
			for zkgt := range z.Fields {
				err = z.Fields[zkgt].DecodeMsg(dc)
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
	for zkgt := range z.Fields {
		err = z.Fields[zkgt].EncodeMsg(en)
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
	for zkgt := range z.Fields {
		o, err = z.Fields[zkgt].MarshalMsg(o)
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
	var zqke uint32
	zqke, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zqke > 0 {
		zqke--
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
			var zqyh uint32
			zqyh, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zqyh) {
				z.Fields = (z.Fields)[:zqyh]
			} else {
				z.Fields = make([]Field, zqyh)
			}
			for zkgt := range z.Fields {
				bts, err = z.Fields[zkgt].UnmarshalMsg(bts)
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
	for zkgt := range z.Fields {
		s += z.Fields[zkgt].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Zkind) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zyzr uint64
		zyzr, err = dc.ReadUint64()
		(*z) = Zkind(zyzr)
	}
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Zkind) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint64(uint64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zkind) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zkind) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zywj uint64
		zywj, bts, err = msgp.ReadUint64Bytes(bts)
		(*z) = Zkind(zywj)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zkind) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Ztype) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjpj uint32
	zjpj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjpj > 0 {
		zjpj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			var zzpf uint32
			zzpf, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zzpf > 0 {
				zzpf--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var zrfe uint64
						zrfe, err = dc.ReadUint64()
						z.Name.Kind = Zkind(zrfe)
					}
					if err != nil {
						return
					}
				case "Str":
					z.Name.Str, err = dc.ReadString()
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
		case "Domain":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Domain = nil
			} else {
				if z.Domain == nil {
					z.Domain = new(Ztype)
				}
				err = z.Domain.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Range":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.Range = nil
			} else {
				if z.Range == nil {
					z.Range = new(Ztype)
				}
				err = z.Range.DecodeMsg(dc)
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
func (z *Ztype) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Name"
	// map header, size 2
	// write "Kind"
	err = en.Append(0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Name.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name.Str)
	if err != nil {
		return
	}
	// write "Domain"
	err = en.Append(0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
	if err != nil {
		return err
	}
	if z.Domain == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Domain.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "Range"
	err = en.Append(0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65)
	if err != nil {
		return err
	}
	if z.Range == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Range.EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Ztype) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Name"
	// map header, size 2
	// string "Kind"
	o = append(o, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Name.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Name.Str)
	// string "Domain"
	o = append(o, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e)
	if z.Domain == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Domain.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Range"
	o = append(o, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65)
	if z.Range == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Range.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Ztype) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zgmo uint32
	zgmo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zgmo > 0 {
		zgmo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			var ztaf uint32
			ztaf, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for ztaf > 0 {
				ztaf--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var zeth uint64
						zeth, bts, err = msgp.ReadUint64Bytes(bts)
						z.Name.Kind = Zkind(zeth)
					}
					if err != nil {
						return
					}
				case "Str":
					z.Name.Str, bts, err = msgp.ReadStringBytes(bts)
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
		case "Domain":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Domain = nil
			} else {
				if z.Domain == nil {
					z.Domain = new(Ztype)
				}
				bts, err = z.Domain.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Range":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Range = nil
			} else {
				if z.Range == nil {
					z.Range = new(Ztype)
				}
				bts, err = z.Range.UnmarshalMsg(bts)
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
func (z *Ztype) Msgsize() (s int) {
	s = 1 + 5 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Name.Str) + 7
	if z.Domain == nil {
		s += msgp.NilSize
	} else {
		s += z.Domain.Msgsize()
	}
	s += 6
	if z.Range == nil {
		s += msgp.NilSize
	} else {
		s += z.Range.Msgsize()
	}
	return
}
