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
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
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
		case "Nam":
			z.Nam, err = dc.ReadString()
			if err != nil {
				return
			}
		case "TypStr":
			z.TypStr, err = dc.ReadString()
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
		case "Tag":
			var zcmr uint32
			zcmr, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Tag == nil && zcmr > 0 {
				z.Tag = make(map[string]string, zcmr)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zcmr > 0 {
				zcmr--
				var zxvk string
				var zbzg string
				zxvk, err = dc.ReadString()
				if err != nil {
					return
				}
				zbzg, err = dc.ReadString()
				if err != nil {
					return
				}
				z.Tag[zxvk] = zbzg
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
	// map header, size 6
	// write "Zid"
	err = en.Append(0x86, 0xa3, 0x5a, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Zid)
	if err != nil {
		return
	}
	// write "Nam"
	err = en.Append(0xa3, 0x4e, 0x61, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nam)
	if err != nil {
		return
	}
	// write "TypStr"
	err = en.Append(0xa6, 0x54, 0x79, 0x70, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TypStr)
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
	// write "Tag"
	err = en.Append(0xa3, 0x54, 0x61, 0x67)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Tag)))
	if err != nil {
		return
	}
	for zxvk, zbzg := range z.Tag {
		err = en.WriteString(zxvk)
		if err != nil {
			return
		}
		err = en.WriteString(zbzg)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Field) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Zid"
	o = append(o, 0x86, 0xa3, 0x5a, 0x69, 0x64)
	o = msgp.AppendInt64(o, z.Zid)
	// string "Nam"
	o = append(o, 0xa3, 0x4e, 0x61, 0x6d)
	o = msgp.AppendString(o, z.Nam)
	// string "TypStr"
	o = append(o, 0xa6, 0x54, 0x79, 0x70, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.TypStr)
	// string "OmitEmpty"
	o = append(o, 0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79)
	o = msgp.AppendBool(o, z.OmitEmpty)
	// string "Skip"
	o = append(o, 0xa4, 0x53, 0x6b, 0x69, 0x70)
	o = msgp.AppendBool(o, z.Skip)
	// string "Tag"
	o = append(o, 0xa3, 0x54, 0x61, 0x67)
	o = msgp.AppendMapHeader(o, uint32(len(z.Tag)))
	for zxvk, zbzg := range z.Tag {
		o = msgp.AppendString(o, zxvk)
		o = msgp.AppendString(o, zbzg)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Field) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
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
		case "Nam":
			z.Nam, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TypStr":
			z.TypStr, bts, err = msgp.ReadStringBytes(bts)
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
		case "Tag":
			var zwht uint32
			zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Tag == nil && zwht > 0 {
				z.Tag = make(map[string]string, zwht)
			} else if len(z.Tag) > 0 {
				for key, _ := range z.Tag {
					delete(z.Tag, key)
				}
			}
			for zwht > 0 {
				var zxvk string
				var zbzg string
				zwht--
				zxvk, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				zbzg, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				z.Tag[zxvk] = zbzg
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
	s = 1 + 4 + msgp.Int64Size + 4 + msgp.StringPrefixSize + len(z.Nam) + 7 + msgp.StringPrefixSize + len(z.TypStr) + 10 + msgp.BoolSize + 5 + msgp.BoolSize + 4 + msgp.MapHeaderSize
	if z.Tag != nil {
		for zxvk, zbzg := range z.Tag {
			_ = zbzg
			s += msgp.StringPrefixSize + len(zxvk) + msgp.StringPrefixSize + len(zbzg)
		}
	}
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
		case "PkgPath":
			z.PkgPath, err = dc.ReadString()
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
					case "Fld":
						var zpks uint32
						zpks, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Structs[zhct].Fld) >= int(zpks) {
							z.Structs[zhct].Fld = (z.Structs[zhct].Fld)[:zpks]
						} else {
							z.Structs[zhct].Fld = make([]Field, zpks)
						}
						for zcua := range z.Structs[zhct].Fld {
							err = z.Structs[zhct].Fld[zcua].DecodeMsg(dc)
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
	// map header, size 2
	// write "PkgPath"
	err = en.Append(0x82, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.PkgPath)
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
		// write "Fld"
		err = en.Append(0xa3, 0x46, 0x6c, 0x64)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zhct].Fld)))
		if err != nil {
			return
		}
		for zcua := range z.Structs[zhct].Fld {
			err = z.Structs[zhct].Fld[zcua].EncodeMsg(en)
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
	// map header, size 2
	// string "PkgPath"
	o = append(o, 0x82, 0xa7, 0x50, 0x6b, 0x67, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.PkgPath)
	// string "Structs"
	o = append(o, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Structs)))
	for zhct := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zhct].StructName)
		// string "Fld"
		o = append(o, 0xa3, 0x46, 0x6c, 0x64)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zhct].Fld)))
		for zcua := range z.Structs[zhct].Fld {
			o, err = z.Structs[zhct].Fld[zcua].MarshalMsg(o)
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
		case "PkgPath":
			z.PkgPath, bts, err = msgp.ReadStringBytes(bts)
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
					case "Fld":
						var zrsw uint32
						zrsw, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Structs[zhct].Fld) >= int(zrsw) {
							z.Structs[zhct].Fld = (z.Structs[zhct].Fld)[:zrsw]
						} else {
							z.Structs[zhct].Fld = make([]Field, zrsw)
						}
						for zcua := range z.Structs[zhct].Fld {
							bts, err = z.Structs[zhct].Fld[zcua].UnmarshalMsg(bts)
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
	s = 1 + 8 + msgp.StringPrefixSize + len(z.PkgPath) + 8 + msgp.ArrayHeaderSize
	for zhct := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zhct].StructName) + 4 + msgp.ArrayHeaderSize
		for zcua := range z.Structs[zhct].Fld {
			s += z.Structs[zhct].Fld[zcua].Msgsize()
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
		case "Fld":
			var zobc uint32
			zobc, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fld) >= int(zobc) {
				z.Fld = (z.Fld)[:zobc]
			} else {
				z.Fld = make([]Field, zobc)
			}
			for zxpk := range z.Fld {
				err = z.Fld[zxpk].DecodeMsg(dc)
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
	// write "Fld"
	err = en.Append(0xa3, 0x46, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fld)))
	if err != nil {
		return
	}
	for zxpk := range z.Fld {
		err = z.Fld[zxpk].EncodeMsg(en)
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
	// string "Fld"
	o = append(o, 0xa3, 0x46, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fld)))
	for zxpk := range z.Fld {
		o, err = z.Fld[zxpk].MarshalMsg(o)
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
		case "Fld":
			var zkgt uint32
			zkgt, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fld) >= int(zkgt) {
				z.Fld = (z.Fld)[:zkgt]
			} else {
				z.Fld = make([]Field, zkgt)
			}
			for zxpk := range z.Fld {
				bts, err = z.Fld[zxpk].UnmarshalMsg(bts)
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
	s = 1 + 11 + msgp.StringPrefixSize + len(z.StructName) + 4 + msgp.ArrayHeaderSize
	for zxpk := range z.Fld {
		s += z.Fld[zxpk].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Zprimitive) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zema uint8
		zema, err = dc.ReadUint8()
		(*z) = Zprimitive(zema)
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
		var zpez uint8
		zpez, bts, err = msgp.ReadUint8Bytes(bts)
		(*z) = Zprimitive(zpez)
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
