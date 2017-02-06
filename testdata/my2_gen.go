package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import "github.com/glycerine/zebrapack/msgp"

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Tr) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields0zesn = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zesn uint32
	totalEncodedFields0zesn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zesn := totalEncodedFields0zesn
	missingFieldsLeft0zesn := maxFields0zesn - totalEncodedFields0zesn

	var nextMiss0zesn int = -1
	var found0zesn [maxFields0zesn]bool
	var curField0zesn int

doneWithStruct0zesn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zesn > 0 || missingFieldsLeft0zesn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zesn, missingFieldsLeft0zesn, msgp.ShowFound(found0zesn[:]), decodeMsgFieldOrder0zesn)
		if encodedFieldsLeft0zesn > 0 {
			encodedFieldsLeft0zesn--
			curField0zesn, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zesn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zesn = 0
			}
			for nextMiss0zesn < maxFields0zesn && (found0zesn[nextMiss0zesn] || decodeMsgFieldSkip0zesn[nextMiss0zesn]) {
				nextMiss0zesn++
			}
			if nextMiss0zesn == maxFields0zesn {
				// filled all the empty fields!
				break doneWithStruct0zesn
			}
			missingFieldsLeft0zesn--
			curField0zesn = nextMiss0zesn
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zesn)
		switch curField0zesn {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found0zesn[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found0zesn[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found0zesn[3] = true
			var zwba uint32
			zwba, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zwba > 0 {
				z.Snot = make(map[string]bool, zwba)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zwba > 0 {
				zwba--
				var zijs string
				var zbiw bool
				zijs, err = dc.ReadString()
				if err != nil {
					return
				}
				zbiw, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zijs] = zbiw
			}
		case 3:
			// zid 3 for "Notyet"
			found0zesn[4] = true
			var zchv uint32
			zchv, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zchv > 0 {
				z.Notyet = make(map[string]bool, zchv)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zchv > 0 {
				zchv--
				var zsor string
				var zioo bool
				zsor, err = dc.ReadString()
				if err != nil {
					return
				}
				zioo, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zsor] = zioo
			}
		case 4:
			// zid 4 for "Setm"
			found0zesn[5] = true
			var zjmv uint32
			zjmv, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zjmv) {
				z.Setm = (z.Setm)[:zjmv]
			} else {
				z.Setm = make([]*inn, zjmv)
			}
			for zeob := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zeob] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zeob].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zeob] == nil {
						z.Setm[zeob] = new(inn)
					}
					err = z.Setm[zeob].DecodeMsg(dc)
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
	if nextMiss0zesn != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Tr
var decodeMsgFieldOrder0zesn = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var decodeMsgFieldSkip0zesn = []bool{false, true, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Tr) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.U) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[2] = (z.Nt == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.Snot) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.Notyet) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (len(z.Setm) == 0) // string, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Tr) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zety [6]bool
	fieldsInUse_zjvt := z.fieldsNotEmpty(empty_zety[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjvt)
	if err != nil {
		return err
	}

	if !empty_zety[0] {
		// zid 0 for "U"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteString(z.U)
		if err != nil {
			return
		}
	}

	if !empty_zety[2] {
		// zid 1 for "Nt"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Nt)
		if err != nil {
			return
		}
	}

	if !empty_zety[3] {
		// zid 2 for "Snot"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zijs, zbiw := range z.Snot {
			err = en.WriteString(zijs)
			if err != nil {
				return
			}
			err = en.WriteBool(zbiw)
			if err != nil {
				return
			}
		}
	}

	if !empty_zety[4] {
		// zid 3 for "Notyet"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zsor, zioo := range z.Notyet {
			err = en.WriteString(zsor)
			if err != nil {
				return
			}
			err = en.WriteBool(zioo)
			if err != nil {
				return
			}
		}
	}

	if !empty_zety[5] {
		// zid 4 for "Setm"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zeob := range z.Setm {
			if z.Setm[zeob] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zeob].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tr) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "U"
		o = append(o, 0x0)
		o = msgp.AppendString(o, z.U)
	}

	if !empty[2] {
		// zid 1 for "Nt"
		o = append(o, 0x1)
		o = msgp.AppendInt(o, z.Nt)
	}

	if !empty[3] {
		// zid 2 for "Snot"
		o = append(o, 0x2)
		o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
		for zijs, zbiw := range z.Snot {
			o = msgp.AppendString(o, zijs)
			o = msgp.AppendBool(o, zbiw)
		}
	}

	if !empty[4] {
		// zid 3 for "Notyet"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zsor, zioo := range z.Notyet {
			o = msgp.AppendString(o, zsor)
			o = msgp.AppendBool(o, zioo)
		}
	}

	if !empty[5] {
		// zid 4 for "Setm"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zeob := range z.Setm {
			if z.Setm[zeob] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zeob].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Tr) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zhrg = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zhrg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhrg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zhrg := totalEncodedFields1zhrg
	missingFieldsLeft1zhrg := maxFields1zhrg - totalEncodedFields1zhrg

	var nextMiss1zhrg int = -1
	var found1zhrg [maxFields1zhrg]bool
	var curField1zhrg int

doneWithStruct1zhrg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhrg > 0 || missingFieldsLeft1zhrg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhrg, missingFieldsLeft1zhrg, msgp.ShowFound(found1zhrg[:]), unmarshalMsgFieldOrder1zhrg)
		if encodedFieldsLeft1zhrg > 0 {
			encodedFieldsLeft1zhrg--
			curField1zhrg, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zhrg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhrg = 0
			}
			for nextMiss1zhrg < maxFields1zhrg && (found1zhrg[nextMiss1zhrg] || unmarshalMsgFieldSkip1zhrg[nextMiss1zhrg]) {
				nextMiss1zhrg++
			}
			if nextMiss1zhrg == maxFields1zhrg {
				// filled all the empty fields!
				break doneWithStruct1zhrg
			}
			missingFieldsLeft1zhrg--
			curField1zhrg = nextMiss1zhrg
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhrg)
		switch curField1zhrg {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found1zhrg[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found1zhrg[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found1zhrg[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zaxw uint32
				zaxw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zaxw > 0 {
					z.Snot = make(map[string]bool, zaxw)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zaxw > 0 {
					var zijs string
					var zbiw bool
					zaxw--
					zijs, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zbiw, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zijs] = zbiw
				}
			}
		case 3:
			// zid 3 for "Notyet"
			found1zhrg[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var ziun uint32
				ziun, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && ziun > 0 {
					z.Notyet = make(map[string]bool, ziun)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for ziun > 0 {
					var zsor string
					var zioo bool
					ziun--
					zsor, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zioo, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zsor] = zioo
				}
			}
		case 4:
			// zid 4 for "Setm"
			found1zhrg[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zktf uint32
				zktf, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zktf) {
					z.Setm = (z.Setm)[:zktf]
				} else {
					z.Setm = make([]*inn, zktf)
				}
				for zeob := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zeob] != nil {
							z.Setm[zeob].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zeob] {
								z.Setm[zeob].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zeob] == nil {
								z.Setm[zeob] = new(inn)
							}
							bts, err = z.Setm[zeob].UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
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
	if nextMiss1zhrg != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Tr
var unmarshalMsgFieldOrder1zhrg = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var unmarshalMsgFieldSkip1zhrg = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.U) + 3 + msgp.IntSize + 5 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zijs, zbiw := range z.Snot {
			_ = zbiw
			_ = zijs
			s += msgp.StringPrefixSize + len(zijs) + msgp.BoolSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zsor, zioo := range z.Notyet {
			_ = zioo
			_ = zsor
			s += msgp.StringPrefixSize + len(zsor) + msgp.BoolSize
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zeob := range z.Setm {
		if z.Setm[zeob] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zeob].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *inn) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zjyx = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zjyx uint32
	totalEncodedFields2zjyx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zjyx := totalEncodedFields2zjyx
	missingFieldsLeft2zjyx := maxFields2zjyx - totalEncodedFields2zjyx

	var nextMiss2zjyx int = -1
	var found2zjyx [maxFields2zjyx]bool
	var curField2zjyx int

doneWithStruct2zjyx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zjyx > 0 || missingFieldsLeft2zjyx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zjyx, missingFieldsLeft2zjyx, msgp.ShowFound(found2zjyx[:]), decodeMsgFieldOrder2zjyx)
		if encodedFieldsLeft2zjyx > 0 {
			encodedFieldsLeft2zjyx--
			curField2zjyx, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zjyx < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zjyx = 0
			}
			for nextMiss2zjyx < maxFields2zjyx && (found2zjyx[nextMiss2zjyx] || decodeMsgFieldSkip2zjyx[nextMiss2zjyx]) {
				nextMiss2zjyx++
			}
			if nextMiss2zjyx == maxFields2zjyx {
				// filled all the empty fields!
				break doneWithStruct2zjyx
			}
			missingFieldsLeft2zjyx--
			curField2zjyx = nextMiss2zjyx
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zjyx)
		switch curField2zjyx {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found2zjyx[0] = true
			var zhrr uint32
			zhrr, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zhrr > 0 {
				z.j = make(map[string]bool, zhrr)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zhrr > 0 {
				zhrr--
				var zfkn string
				var zial bool
				zfkn, err = dc.ReadString()
				if err != nil {
					return
				}
				zial, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zfkn] = zial
			}
		case 1:
			// zid 1 for "e"
			found2zjyx[1] = true
			z.e, err = dc.ReadInt64()
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
	if nextMiss2zjyx != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of inn
var decodeMsgFieldOrder2zjyx = []string{"j", "e"}

var decodeMsgFieldSkip2zjyx = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *inn) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.j) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.e == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *inn) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zmcs [2]bool
	fieldsInUse_zxjj := z.fieldsNotEmpty(empty_zmcs[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxjj)
	if err != nil {
		return err
	}

	if !empty_zmcs[0] {
		// zid 0 for "j"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zfkn, zial := range z.j {
			err = en.WriteString(zfkn)
			if err != nil {
				return
			}
			err = en.WriteBool(zial)
			if err != nil {
				return
			}
		}
	}

	if !empty_zmcs[1] {
		// zid 1 for "e"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.e)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *inn) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "j"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.j)))
		for zfkn, zial := range z.j {
			o = msgp.AppendString(o, zfkn)
			o = msgp.AppendBool(o, zial)
		}
	}

	if !empty[1] {
		// zid 1 for "e"
		o = append(o, 0x1)
		o = msgp.AppendInt64(o, z.e)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *inn) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *inn) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zhdv = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zhdv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zhdv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zhdv := totalEncodedFields3zhdv
	missingFieldsLeft3zhdv := maxFields3zhdv - totalEncodedFields3zhdv

	var nextMiss3zhdv int = -1
	var found3zhdv [maxFields3zhdv]bool
	var curField3zhdv int

doneWithStruct3zhdv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zhdv > 0 || missingFieldsLeft3zhdv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zhdv, missingFieldsLeft3zhdv, msgp.ShowFound(found3zhdv[:]), unmarshalMsgFieldOrder3zhdv)
		if encodedFieldsLeft3zhdv > 0 {
			encodedFieldsLeft3zhdv--
			curField3zhdv, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zhdv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zhdv = 0
			}
			for nextMiss3zhdv < maxFields3zhdv && (found3zhdv[nextMiss3zhdv] || unmarshalMsgFieldSkip3zhdv[nextMiss3zhdv]) {
				nextMiss3zhdv++
			}
			if nextMiss3zhdv == maxFields3zhdv {
				// filled all the empty fields!
				break doneWithStruct3zhdv
			}
			missingFieldsLeft3zhdv--
			curField3zhdv = nextMiss3zhdv
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zhdv)
		switch curField3zhdv {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found3zhdv[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var znen uint32
				znen, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && znen > 0 {
					z.j = make(map[string]bool, znen)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for znen > 0 {
					var zfkn string
					var zial bool
					znen--
					zfkn, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zial, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zfkn] = zial
				}
			}
		case 1:
			// zid 1 for "e"
			found3zhdv[1] = true
			z.e, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss3zhdv != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of inn
var unmarshalMsgFieldOrder3zhdv = []string{"j", "e"}

var unmarshalMsgFieldSkip3zhdv = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.j != nil {
		for zfkn, zial := range z.j {
			_ = zial
			_ = zfkn
			s += msgp.StringPrefixSize + len(zfkn) + msgp.BoolSize
		}
	}
	s += 2 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *u) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields4zmqd = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zmqd uint32
	totalEncodedFields4zmqd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zmqd := totalEncodedFields4zmqd
	missingFieldsLeft4zmqd := maxFields4zmqd - totalEncodedFields4zmqd

	var nextMiss4zmqd int = -1
	var found4zmqd [maxFields4zmqd]bool
	var curField4zmqd int

doneWithStruct4zmqd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zmqd > 0 || missingFieldsLeft4zmqd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zmqd, missingFieldsLeft4zmqd, msgp.ShowFound(found4zmqd[:]), decodeMsgFieldOrder4zmqd)
		if encodedFieldsLeft4zmqd > 0 {
			encodedFieldsLeft4zmqd--
			curField4zmqd, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zmqd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zmqd = 0
			}
			for nextMiss4zmqd < maxFields4zmqd && (found4zmqd[nextMiss4zmqd] || decodeMsgFieldSkip4zmqd[nextMiss4zmqd]) {
				nextMiss4zmqd++
			}
			if nextMiss4zmqd == maxFields4zmqd {
				// filled all the empty fields!
				break doneWithStruct4zmqd
			}
			missingFieldsLeft4zmqd--
			curField4zmqd = nextMiss4zmqd
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zmqd)
		switch curField4zmqd {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found4zmqd[0] = true
			var zkqe uint32
			zkqe, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zkqe > 0 {
				z.m = make(map[string]*Tr, zkqe)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zkqe > 0 {
				zkqe--
				var zhfq string
				var zabq *Tr
				zhfq, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zabq != nil {
						dc.PushAlwaysNil()
						err = zabq.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zabq == nil {
						zabq = new(Tr)
					}
					err = zabq.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zhfq] = zabq
			}
		case 1:
			// zid 1 for "s"
			found4zmqd[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found4zmqd[2] = true
			z.n, err = dc.ReadInt64()
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
	if nextMiss4zmqd != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of u
var decodeMsgFieldOrder4zmqd = []string{"m", "s", "n"}

var decodeMsgFieldSkip4zmqd = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *u) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (len(z.m) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.s) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.n == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *u) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zzuu [3]bool
	fieldsInUse_zpoq := z.fieldsNotEmpty(empty_zzuu[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpoq)
	if err != nil {
		return err
	}

	if !empty_zzuu[0] {
		// zid 0 for "m"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zhfq, zabq := range z.m {
			err = en.WriteString(zhfq)
			if err != nil {
				return
			}
			if zabq == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zabq.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zzuu[1] {
		// zid 1 for "s"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.s)
		if err != nil {
			return
		}
	}

	if !empty_zzuu[2] {
		// zid 2 for "n"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.n)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *u) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// zid 0 for "m"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.m)))
		for zhfq, zabq := range z.m {
			o = msgp.AppendString(o, zhfq)
			if zabq == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zabq.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[1] {
		// zid 1 for "s"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.s)
	}

	if !empty[2] {
		// zid 2 for "n"
		o = append(o, 0x2)
		o = msgp.AppendInt64(o, z.n)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *u) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *u) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields5zbbw = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zbbw uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zbbw, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zbbw := totalEncodedFields5zbbw
	missingFieldsLeft5zbbw := maxFields5zbbw - totalEncodedFields5zbbw

	var nextMiss5zbbw int = -1
	var found5zbbw [maxFields5zbbw]bool
	var curField5zbbw int

doneWithStruct5zbbw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zbbw > 0 || missingFieldsLeft5zbbw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zbbw, missingFieldsLeft5zbbw, msgp.ShowFound(found5zbbw[:]), unmarshalMsgFieldOrder5zbbw)
		if encodedFieldsLeft5zbbw > 0 {
			encodedFieldsLeft5zbbw--
			curField5zbbw, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zbbw < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zbbw = 0
			}
			for nextMiss5zbbw < maxFields5zbbw && (found5zbbw[nextMiss5zbbw] || unmarshalMsgFieldSkip5zbbw[nextMiss5zbbw]) {
				nextMiss5zbbw++
			}
			if nextMiss5zbbw == maxFields5zbbw {
				// filled all the empty fields!
				break doneWithStruct5zbbw
			}
			missingFieldsLeft5zbbw--
			curField5zbbw = nextMiss5zbbw
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zbbw)
		switch curField5zbbw {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found5zbbw[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zbyt uint32
				zbyt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zbyt > 0 {
					z.m = make(map[string]*Tr, zbyt)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zbyt > 0 {
					var zhfq string
					var zabq *Tr
					zbyt--
					zhfq, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zabq != nil {
							zabq.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zabq {
								zabq.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zabq == nil {
								zabq = new(Tr)
							}
							bts, err = zabq.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zhfq] = zabq
				}
			}
		case 1:
			// zid 1 for "s"
			found5zbbw[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found5zbbw[2] = true
			z.n, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss5zbbw != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of u
var unmarshalMsgFieldOrder5zbbw = []string{"m", "s", "n"}

var unmarshalMsgFieldSkip5zbbw = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.m != nil {
		for zhfq, zabq := range z.m {
			_ = zabq
			_ = zhfq
			s += msgp.StringPrefixSize + len(zhfq)
			if zabq == nil {
				s += msgp.NilSize
			} else {
				s += zabq.Msgsize()
			}
		}
	}
	s += 2 + msgp.StringPrefixSize + len(z.s) + 2 + msgp.Int64Size
	return
}

// FileMy2 holds ZebraPack schema from file 'my2.go'
type FileMy2 struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 1636 bytes
func (FileMy2) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xa6, 0x6d, 0x79, 0x32, 0x2e, 0x67, 0x6f, 0xad,
		0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
		0x61, 0x67, 0x65, 0xa8, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
		0x74, 0x61, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x00, 0xa7, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x73, 0x83, 0xa1, 0x75, 0x82, 0xaa,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x75, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x93,
		0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6d,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xae, 0x6d,
		0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d,
		0x2a, 0x54, 0x72, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7,
		0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x16, 0xa3, 0x53, 0x74, 0x72, 0xa2, 0x54, 0x72, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x73, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa1, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69,
		0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6e, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74,
		0x36, 0x34, 0xa2, 0x54, 0x72, 0x82, 0xaa, 0x53, 0x74, 0x72,
		0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x54, 0x72,
		0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x55, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa1, 0x55, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x85, 0xa3, 0x5a, 0x69,
		0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x65, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x2d, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xa4, 0x53, 0x6b,
		0x69, 0x70, 0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa2, 0x4e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x4e, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0d, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0d, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0x86, 0xa3, 0x5a,
		0x69, 0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e,
		0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3,
		0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa3,
		0x5a, 0x69, 0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74,
		0x79, 0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74,
		0x79, 0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70,
		0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f,
		0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53,
		0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02,
		0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62,
		0x6f, 0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa4, 0x53, 0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
		0x53, 0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x5b, 0x5d,
		0x2a, 0x69, 0x6e, 0x6e, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a,
		0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x83, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50,
		0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16,
		0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x6e, 0xa3, 0x69,
		0x6e, 0x6e, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
		0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x69, 0x6e, 0x6e, 0xa6, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x73, 0x92, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6a, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x6a, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72,
		0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53,
		0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5,
		0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f,
		0x6c, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x65, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17,
		0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74,
		0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa7, 0x49, 0x6d,
		0x70, 0x6f, 0x72, 0x74, 0x73, 0x91, 0xbd, 0x22, 0x67, 0x69,
		0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
		0x6c, 0x79, 0x63, 0x65, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x72,
		0x62, 0x74, 0x72, 0x65, 0x65, 0x22,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 2153 bytes
func (FileMy2) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"my2.go","SourcePackage":"testdata","ZebraSchemaId":0,"Structs":{"u":{"StructName":"u","Fields":[{"Zid":0,"FieldGoName":"m","FieldTagName":"m","FieldTypeStr":"map[string]*Tr","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Tr"}}}},{"Zid":1,"FieldGoName":"s","FieldTagName":"s","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"n","FieldTagName":"n","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Tr":{"StructName":"Tr","Fields":[{"Zid":0,"FieldGoName":"U","FieldTagName":"U","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"et","FieldTagName":"-","FieldCategory":28,"Skip":true},{"Zid":1,"FieldGoName":"Nt","FieldTagName":"Nt","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}},{"Zid":2,"FieldGoName":"Snot","FieldTagName":"Snot","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":3,"FieldGoName":"Notyet","FieldTagName":"Notyet","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":4,"FieldGoName":"Setm","FieldTagName":"Setm","FieldTypeStr":"[]*inn","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"inn"}}}}]},"inn":{"StructName":"inn","Fields":[{"Zid":0,"FieldGoName":"j","FieldTagName":"j","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":1,"FieldGoName":"e","FieldTagName":"e","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]}},"Imports":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 6280 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "my2.go",
    "SourcePackage": "testdata",
    "ZebraSchemaId": 0,
    "Structs": {
        "u": {
            "StructName": "u",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "m",
                    "FieldTagName": "m",
                    "FieldTypeStr": "map[string]*Tr",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "Tr"
                            }
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "s",
                    "FieldTagName": "s",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "n",
                    "FieldTagName": "n",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        },
        "Tr": {
            "StructName": "Tr",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "U",
                    "FieldTagName": "U",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "et",
                    "FieldTagName": "-",
                    "FieldCategory": 28,
                    "Skip": true
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Nt",
                    "FieldTagName": "Nt",
                    "FieldTypeStr": "int",
                    "FieldCategory": 23,
                    "FieldPrimitive": 13,
                    "FieldFullType": {
                        "Kind": 13,
                        "Str": "int"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Snot",
                    "FieldTagName": "Snot",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 3,
                    "FieldGoName": "Notyet",
                    "FieldTagName": "Notyet",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 4,
                    "FieldGoName": "Setm",
                    "FieldTagName": "Setm",
                    "FieldTypeStr": "[]*inn",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "inn"
                            }
                        }
                    }
                }
            ]
        },
        "inn": {
            "StructName": "inn",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "j",
                    "FieldTagName": "j",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "e",
                    "FieldTagName": "e",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        }
    },
    "Imports": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
