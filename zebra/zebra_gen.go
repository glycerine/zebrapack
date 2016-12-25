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
			var zcmr uint32
			zcmr, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zcmr > 0 {
				zcmr--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Name":
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
						case "Kind":
							{
								var zwht uint64
								zwht, err = dc.ReadUint64()
								z.FieldFullType.Name.Kind = Zkind(zwht)
							}
							if err != nil {
								return
							}
						case "Str":
							z.FieldFullType.Name.Str, err = dc.ReadString()
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
								z.FieldFullType.Domain.Kind = Zkind(zcua)
							}
							if err != nil {
								return
							}
						case "Str":
							z.FieldFullType.Domain.Str, err = dc.ReadString()
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
				case "Range":
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
						case "Kind":
							{
								var zlqf uint64
								zlqf, err = dc.ReadUint64()
								z.FieldFullType.Range.Kind = Zkind(zlqf)
							}
							if err != nil {
								return
							}
						case "Str":
							z.FieldFullType.Range.Str, err = dc.ReadString()
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
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
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
	// map header, size 3
	// write "Name"
	// map header, size 2
	// write "Kind"
	err = en.Append(0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldFullType.Name.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldFullType.Name.Str)
	if err != nil {
		return
	}
	// write "Domain"
	// map header, size 2
	// write "Kind"
	err = en.Append(0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldFullType.Domain.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldFullType.Domain.Str)
	if err != nil {
		return
	}
	// write "Range"
	// map header, size 2
	// write "Kind"
	err = en.Append(0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.FieldFullType.Range.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldFullType.Range.Str)
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
	// map header, size 3
	// string "Name"
	// map header, size 2
	// string "Kind"
	o = append(o, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.FieldFullType.Name.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldFullType.Name.Str)
	// string "Domain"
	// map header, size 2
	// string "Kind"
	o = append(o, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.FieldFullType.Domain.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldFullType.Domain.Str)
	// string "Range"
	// map header, size 2
	// string "Kind"
	o = append(o, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.FieldFullType.Range.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldFullType.Range.Str)
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
				var zpks uint64
				zpks, bts, err = msgp.ReadUint64Bytes(bts)
				z.FieldCategory = Zkind(zpks)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive":
			{
				var zjfb uint64
				zjfb, bts, err = msgp.ReadUint64Bytes(bts)
				z.FieldPrimitive = Zkind(zjfb)
			}
			if err != nil {
				return
			}
		case "FieldFullType":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zcxo > 0 {
				zcxo--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Name":
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
						case "Kind":
							{
								var zrsw uint64
								zrsw, bts, err = msgp.ReadUint64Bytes(bts)
								z.FieldFullType.Name.Kind = Zkind(zrsw)
							}
							if err != nil {
								return
							}
						case "Str":
							z.FieldFullType.Name.Str, bts, err = msgp.ReadStringBytes(bts)
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
						case "Kind":
							{
								var zdnj uint64
								zdnj, bts, err = msgp.ReadUint64Bytes(bts)
								z.FieldFullType.Domain.Kind = Zkind(zdnj)
							}
							if err != nil {
								return
							}
						case "Str":
							z.FieldFullType.Domain.Str, bts, err = msgp.ReadStringBytes(bts)
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
				case "Range":
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
						case "Kind":
							{
								var zsnv uint64
								zsnv, bts, err = msgp.ReadUint64Bytes(bts)
								z.FieldFullType.Range.Kind = Zkind(zsnv)
							}
							if err != nil {
								return
							}
						case "Str":
							z.FieldFullType.Range.Str, bts, err = msgp.ReadStringBytes(bts)
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
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
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
	s = 1 + 4 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.FieldGoName) + 13 + msgp.StringPrefixSize + len(z.FieldTagName) + 13 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 14 + msgp.Uint64Size + 15 + msgp.Uint64Size + 14 + 1 + 5 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.FieldFullType.Name.Str) + 7 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.FieldFullType.Domain.Str) + 6 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.FieldFullType.Range.Str) + 10 + msgp.BoolSize + 5 + msgp.BoolSize + 11 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *KS) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zkgt uint32
	zkgt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkgt > 0 {
		zkgt--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Kind":
			{
				var zema uint64
				zema, err = dc.ReadUint64()
				z.Kind = Zkind(zema)
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
	var zpez uint32
	zpez, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Kind":
			{
				var zqke uint64
				zqke, bts, err = msgp.ReadUint64Bytes(bts)
				z.Kind = Zkind(zqke)
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
	var zywj uint32
	zywj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zywj > 0 {
		zywj--
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
			var zjpj uint32
			zjpj, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(zjpj) {
				z.Structs = (z.Structs)[:zjpj]
			} else {
				z.Structs = make([]Struct, zjpj)
			}
			for zqyh := range z.Structs {
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
					case "StructName":
						z.Structs[zqyh].StructName, err = dc.ReadString()
						if err != nil {
							return
						}
					case "Fields":
						var zrfe uint32
						zrfe, err = dc.ReadArrayHeader()
						if err != nil {
							return
						}
						if cap(z.Structs[zqyh].Fields) >= int(zrfe) {
							z.Structs[zqyh].Fields = (z.Structs[zqyh].Fields)[:zrfe]
						} else {
							z.Structs[zqyh].Fields = make([]Field, zrfe)
						}
						for zyzr := range z.Structs[zqyh].Fields {
							err = z.Structs[zqyh].Fields[zyzr].DecodeMsg(dc)
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
	for zqyh := range z.Structs {
		// map header, size 2
		// write "StructName"
		err = en.Append(0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Structs[zqyh].StructName)
		if err != nil {
			return
		}
		// write "Fields"
		err = en.Append(0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Structs[zqyh].Fields)))
		if err != nil {
			return
		}
		for zyzr := range z.Structs[zqyh].Fields {
			err = z.Structs[zqyh].Fields[zyzr].EncodeMsg(en)
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
	for zqyh := range z.Structs {
		// map header, size 2
		// string "StructName"
		o = append(o, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Structs[zqyh].StructName)
		// string "Fields"
		o = append(o, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Structs[zqyh].Fields)))
		for zyzr := range z.Structs[zqyh].Fields {
			o, err = z.Structs[zqyh].Fields[zyzr].MarshalMsg(o)
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
			var ztaf uint32
			ztaf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Structs) >= int(ztaf) {
				z.Structs = (z.Structs)[:ztaf]
			} else {
				z.Structs = make([]Struct, ztaf)
			}
			for zqyh := range z.Structs {
				var zeth uint32
				zeth, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zeth > 0 {
					zeth--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "StructName":
						z.Structs[zqyh].StructName, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "Fields":
						var zsbz uint32
						zsbz, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Structs[zqyh].Fields) >= int(zsbz) {
							z.Structs[zqyh].Fields = (z.Structs[zqyh].Fields)[:zsbz]
						} else {
							z.Structs[zqyh].Fields = make([]Field, zsbz)
						}
						for zyzr := range z.Structs[zqyh].Fields {
							bts, err = z.Structs[zqyh].Fields[zyzr].UnmarshalMsg(bts)
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
	for zqyh := range z.Structs {
		s += 1 + 11 + msgp.StringPrefixSize + len(z.Structs[zqyh].StructName) + 7 + msgp.ArrayHeaderSize
		for zyzr := range z.Structs[zqyh].Fields {
			s += z.Structs[zqyh].Fields[zyzr].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zawn uint32
	zawn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zawn > 0 {
		zawn--
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
			var zwel uint32
			zwel, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zwel) {
				z.Fields = (z.Fields)[:zwel]
			} else {
				z.Fields = make([]Field, zwel)
			}
			for zrjx := range z.Fields {
				err = z.Fields[zrjx].DecodeMsg(dc)
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
	for zrjx := range z.Fields {
		err = z.Fields[zrjx].EncodeMsg(en)
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
	for zrjx := range z.Fields {
		o, err = z.Fields[zrjx].MarshalMsg(o)
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
	var zrbe uint32
	zrbe, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrbe > 0 {
		zrbe--
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
			var zmfd uint32
			zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zmfd) {
				z.Fields = (z.Fields)[:zmfd]
			} else {
				z.Fields = make([]Field, zmfd)
			}
			for zrjx := range z.Fields {
				bts, err = z.Fields[zrjx].UnmarshalMsg(bts)
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
	for zrjx := range z.Fields {
		s += z.Fields[zrjx].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Zkind) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zzdc uint64
		zzdc, err = dc.ReadUint64()
		(*z) = Zkind(zzdc)
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
		var zelx uint64
		zelx, bts, err = msgp.ReadUint64Bytes(bts)
		(*z) = Zkind(zelx)
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
	var zbal uint32
	zbal, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			var zjqz uint32
			zjqz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zjqz > 0 {
				zjqz--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var zkct uint64
						zkct, err = dc.ReadUint64()
						z.Name.Kind = Zkind(zkct)
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
			var ztmt uint32
			ztmt, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for ztmt > 0 {
				ztmt--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var ztco uint64
						ztco, err = dc.ReadUint64()
						z.Domain.Kind = Zkind(ztco)
					}
					if err != nil {
						return
					}
				case "Str":
					z.Domain.Str, err = dc.ReadString()
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
		case "Range":
			var zana uint32
			zana, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zana > 0 {
				zana--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var ztyy uint64
						ztyy, err = dc.ReadUint64()
						z.Range.Kind = Zkind(ztyy)
					}
					if err != nil {
						return
					}
				case "Str":
					z.Range.Str, err = dc.ReadString()
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
	// map header, size 2
	// write "Kind"
	err = en.Append(0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Domain.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Domain.Str)
	if err != nil {
		return
	}
	// write "Range"
	// map header, size 2
	// write "Kind"
	err = en.Append(0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Range.Kind))
	if err != nil {
		return
	}
	// write "Str"
	err = en.Append(0xa3, 0x53, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Range.Str)
	if err != nil {
		return
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
	// map header, size 2
	// string "Kind"
	o = append(o, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Domain.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Domain.Str)
	// string "Range"
	// map header, size 2
	// string "Kind"
	o = append(o, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64)
	o = msgp.AppendUint64(o, uint64(z.Range.Kind))
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Range.Str)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Ztype) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zinl uint32
	zinl, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zinl > 0 {
		zinl--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			var zare uint32
			zare, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zare > 0 {
				zare--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var zljy uint64
						zljy, bts, err = msgp.ReadUint64Bytes(bts)
						z.Name.Kind = Zkind(zljy)
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
			var zixj uint32
			zixj, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zixj > 0 {
				zixj--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var zrsc uint64
						zrsc, bts, err = msgp.ReadUint64Bytes(bts)
						z.Domain.Kind = Zkind(zrsc)
					}
					if err != nil {
						return
					}
				case "Str":
					z.Domain.Str, bts, err = msgp.ReadStringBytes(bts)
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
		case "Range":
			var zctn uint32
			zctn, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zctn > 0 {
				zctn--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "Kind":
					{
						var zswy uint64
						zswy, bts, err = msgp.ReadUint64Bytes(bts)
						z.Range.Kind = Zkind(zswy)
					}
					if err != nil {
						return
					}
				case "Str":
					z.Range.Str, bts, err = msgp.ReadStringBytes(bts)
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
	s = 1 + 5 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Name.Str) + 7 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Domain.Str) + 6 + 1 + 5 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Range.Str)
	return
}
