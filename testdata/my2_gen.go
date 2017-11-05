package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
)

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
	const maxFields0zkhz = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zkhz uint32
	totalEncodedFields0zkhz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zkhz := totalEncodedFields0zkhz
	missingFieldsLeft0zkhz := maxFields0zkhz - totalEncodedFields0zkhz

	var nextMiss0zkhz int = -1
	var found0zkhz [maxFields0zkhz]bool
	var curField0zkhz int

doneWithStruct0zkhz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zkhz > 0 || missingFieldsLeft0zkhz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zkhz, missingFieldsLeft0zkhz, msgp.ShowFound(found0zkhz[:]), decodeMsgFieldOrder0zkhz)
		if encodedFieldsLeft0zkhz > 0 {
			encodedFieldsLeft0zkhz--
			curField0zkhz, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zkhz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zkhz = 0
			}
			for nextMiss0zkhz < maxFields0zkhz && (found0zkhz[nextMiss0zkhz] || decodeMsgFieldSkip0zkhz[nextMiss0zkhz]) {
				nextMiss0zkhz++
			}
			if nextMiss0zkhz == maxFields0zkhz {
				// filled all the empty fields!
				break doneWithStruct0zkhz
			}
			missingFieldsLeft0zkhz--
			curField0zkhz = nextMiss0zkhz
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zkhz)
		switch curField0zkhz {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found0zkhz[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found0zkhz[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found0zkhz[3] = true
			var zffc uint32
			zffc, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zffc > 0 {
				z.Snot = make(map[string]bool, zffc)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zffc > 0 {
				zffc--
				var zlzp string
				var zxbo bool
				zlzp, err = dc.ReadString()
				if err != nil {
					return
				}
				zxbo, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zlzp] = zxbo
			}
		case 3:
			// zid 3 for "Notyet"
			found0zkhz[4] = true
			var zptl uint32
			zptl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zptl > 0 {
				z.Notyet = make(map[string]bool, zptl)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zptl > 0 {
				zptl--
				var zlqt string
				var zfjd bool
				zlqt, err = dc.ReadString()
				if err != nil {
					return
				}
				zfjd, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zlqt] = zfjd
			}
		case 4:
			// zid 4 for "Setm"
			found0zkhz[5] = true
			var zaix uint32
			zaix, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zaix) {
				z.Setm = (z.Setm)[:zaix]
			} else {
				z.Setm = make([]*inn, zaix)
			}
			for zxqq := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zxqq] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zxqq].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zxqq] == nil {
						z.Setm[zxqq] = new(inn)
					}
					err = z.Setm[zxqq].DecodeMsg(dc)
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
	if nextMiss0zkhz != -1 {
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
var decodeMsgFieldOrder0zkhz = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var decodeMsgFieldSkip0zkhz = []bool{false, true, false, false, false, false}

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
	var empty_zsze [6]bool
	fieldsInUse_zcuh := z.fieldsNotEmpty(empty_zsze[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zcuh)
	if err != nil {
		return err
	}

	if !empty_zsze[0] {
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

	if !empty_zsze[2] {
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

	if !empty_zsze[3] {
		// zid 2 for "Snot"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zlzp, zxbo := range z.Snot {
			err = en.WriteString(zlzp)
			if err != nil {
				return
			}
			err = en.WriteBool(zxbo)
			if err != nil {
				return
			}
		}
	}

	if !empty_zsze[4] {
		// zid 3 for "Notyet"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zlqt, zfjd := range z.Notyet {
			err = en.WriteString(zlqt)
			if err != nil {
				return
			}
			err = en.WriteBool(zfjd)
			if err != nil {
				return
			}
		}
	}

	if !empty_zsze[5] {
		// zid 4 for "Setm"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zxqq := range z.Setm {
			if z.Setm[zxqq] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zxqq].EncodeMsg(en)
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
		for zlzp, zxbo := range z.Snot {
			o = msgp.AppendString(o, zlzp)
			o = msgp.AppendBool(o, zxbo)
		}
	}

	if !empty[4] {
		// zid 3 for "Notyet"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zlqt, zfjd := range z.Notyet {
			o = msgp.AppendString(o, zlqt)
			o = msgp.AppendBool(o, zfjd)
		}
	}

	if !empty[5] {
		// zid 4 for "Setm"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zxqq := range z.Setm {
			if z.Setm[zxqq] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zxqq].MarshalMsg(o)
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
	const maxFields1zvok = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zvok uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zvok, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zvok := totalEncodedFields1zvok
	missingFieldsLeft1zvok := maxFields1zvok - totalEncodedFields1zvok

	var nextMiss1zvok int = -1
	var found1zvok [maxFields1zvok]bool
	var curField1zvok int

doneWithStruct1zvok:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zvok > 0 || missingFieldsLeft1zvok > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zvok, missingFieldsLeft1zvok, msgp.ShowFound(found1zvok[:]), unmarshalMsgFieldOrder1zvok)
		if encodedFieldsLeft1zvok > 0 {
			encodedFieldsLeft1zvok--
			curField1zvok, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zvok < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zvok = 0
			}
			for nextMiss1zvok < maxFields1zvok && (found1zvok[nextMiss1zvok] || unmarshalMsgFieldSkip1zvok[nextMiss1zvok]) {
				nextMiss1zvok++
			}
			if nextMiss1zvok == maxFields1zvok {
				// filled all the empty fields!
				break doneWithStruct1zvok
			}
			missingFieldsLeft1zvok--
			curField1zvok = nextMiss1zvok
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zvok)
		switch curField1zvok {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found1zvok[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found1zvok[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found1zvok[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var znej uint32
				znej, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && znej > 0 {
					z.Snot = make(map[string]bool, znej)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for znej > 0 {
					var zlzp string
					var zxbo bool
					znej--
					zlzp, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zxbo, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zlzp] = zxbo
				}
			}
		case 3:
			// zid 3 for "Notyet"
			found1zvok[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zfsn uint32
				zfsn, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zfsn > 0 {
					z.Notyet = make(map[string]bool, zfsn)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zfsn > 0 {
					var zlqt string
					var zfjd bool
					zfsn--
					zlqt, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zfjd, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zlqt] = zfjd
				}
			}
		case 4:
			// zid 4 for "Setm"
			found1zvok[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zltg uint32
				zltg, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zltg) {
					z.Setm = (z.Setm)[:zltg]
				} else {
					z.Setm = make([]*inn, zltg)
				}
				for zxqq := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zxqq] != nil {
							z.Setm[zxqq].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zxqq] {
								z.Setm[zxqq].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zxqq] == nil {
								z.Setm[zxqq] = new(inn)
							}
							bts, err = z.Setm[zxqq].UnmarshalMsg(bts)
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
	if nextMiss1zvok != -1 {
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
var unmarshalMsgFieldOrder1zvok = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var unmarshalMsgFieldSkip1zvok = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.U) + 3 + msgp.IntSize + 5 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zlzp, zxbo := range z.Snot {
			_ = zxbo
			_ = zlzp
			s += msgp.StringPrefixSize + len(zlzp) + msgp.BoolSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zlqt, zfjd := range z.Notyet {
			_ = zfjd
			_ = zlqt
			s += msgp.StringPrefixSize + len(zlqt) + msgp.BoolSize
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zxqq := range z.Setm {
		if z.Setm[zxqq] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zxqq].Msgsize()
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
	const maxFields2zbdb = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zbdb uint32
	totalEncodedFields2zbdb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zbdb := totalEncodedFields2zbdb
	missingFieldsLeft2zbdb := maxFields2zbdb - totalEncodedFields2zbdb

	var nextMiss2zbdb int = -1
	var found2zbdb [maxFields2zbdb]bool
	var curField2zbdb int

doneWithStruct2zbdb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zbdb > 0 || missingFieldsLeft2zbdb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zbdb, missingFieldsLeft2zbdb, msgp.ShowFound(found2zbdb[:]), decodeMsgFieldOrder2zbdb)
		if encodedFieldsLeft2zbdb > 0 {
			encodedFieldsLeft2zbdb--
			curField2zbdb, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zbdb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zbdb = 0
			}
			for nextMiss2zbdb < maxFields2zbdb && (found2zbdb[nextMiss2zbdb] || decodeMsgFieldSkip2zbdb[nextMiss2zbdb]) {
				nextMiss2zbdb++
			}
			if nextMiss2zbdb == maxFields2zbdb {
				// filled all the empty fields!
				break doneWithStruct2zbdb
			}
			missingFieldsLeft2zbdb--
			curField2zbdb = nextMiss2zbdb
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zbdb)
		switch curField2zbdb {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found2zbdb[0] = true
			var zddz uint32
			zddz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zddz > 0 {
				z.j = make(map[string]bool, zddz)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zddz > 0 {
				zddz--
				var ztny string
				var zduh bool
				ztny, err = dc.ReadString()
				if err != nil {
					return
				}
				zduh, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[ztny] = zduh
			}
		case 1:
			// zid 1 for "e"
			found2zbdb[1] = true
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
	if nextMiss2zbdb != -1 {
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
var decodeMsgFieldOrder2zbdb = []string{"j", "e"}

var decodeMsgFieldSkip2zbdb = []bool{false, false}

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
	var empty_zxop [2]bool
	fieldsInUse_zfxb := z.fieldsNotEmpty(empty_zxop[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zfxb)
	if err != nil {
		return err
	}

	if !empty_zxop[0] {
		// zid 0 for "j"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for ztny, zduh := range z.j {
			err = en.WriteString(ztny)
			if err != nil {
				return
			}
			err = en.WriteBool(zduh)
			if err != nil {
				return
			}
		}
	}

	if !empty_zxop[1] {
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
		for ztny, zduh := range z.j {
			o = msgp.AppendString(o, ztny)
			o = msgp.AppendBool(o, zduh)
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
	const maxFields3zmyf = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zmyf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zmyf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zmyf := totalEncodedFields3zmyf
	missingFieldsLeft3zmyf := maxFields3zmyf - totalEncodedFields3zmyf

	var nextMiss3zmyf int = -1
	var found3zmyf [maxFields3zmyf]bool
	var curField3zmyf int

doneWithStruct3zmyf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zmyf > 0 || missingFieldsLeft3zmyf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zmyf, missingFieldsLeft3zmyf, msgp.ShowFound(found3zmyf[:]), unmarshalMsgFieldOrder3zmyf)
		if encodedFieldsLeft3zmyf > 0 {
			encodedFieldsLeft3zmyf--
			curField3zmyf, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zmyf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zmyf = 0
			}
			for nextMiss3zmyf < maxFields3zmyf && (found3zmyf[nextMiss3zmyf] || unmarshalMsgFieldSkip3zmyf[nextMiss3zmyf]) {
				nextMiss3zmyf++
			}
			if nextMiss3zmyf == maxFields3zmyf {
				// filled all the empty fields!
				break doneWithStruct3zmyf
			}
			missingFieldsLeft3zmyf--
			curField3zmyf = nextMiss3zmyf
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zmyf)
		switch curField3zmyf {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found3zmyf[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zpmt uint32
				zpmt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zpmt > 0 {
					z.j = make(map[string]bool, zpmt)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zpmt > 0 {
					var ztny string
					var zduh bool
					zpmt--
					ztny, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zduh, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[ztny] = zduh
				}
			}
		case 1:
			// zid 1 for "e"
			found3zmyf[1] = true
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
	if nextMiss3zmyf != -1 {
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
var unmarshalMsgFieldOrder3zmyf = []string{"j", "e"}

var unmarshalMsgFieldSkip3zmyf = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.j != nil {
		for ztny, zduh := range z.j {
			_ = zduh
			_ = ztny
			s += msgp.StringPrefixSize + len(ztny) + msgp.BoolSize
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
	const maxFields4zefb = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zefb uint32
	totalEncodedFields4zefb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zefb := totalEncodedFields4zefb
	missingFieldsLeft4zefb := maxFields4zefb - totalEncodedFields4zefb

	var nextMiss4zefb int = -1
	var found4zefb [maxFields4zefb]bool
	var curField4zefb int

doneWithStruct4zefb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zefb > 0 || missingFieldsLeft4zefb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zefb, missingFieldsLeft4zefb, msgp.ShowFound(found4zefb[:]), decodeMsgFieldOrder4zefb)
		if encodedFieldsLeft4zefb > 0 {
			encodedFieldsLeft4zefb--
			curField4zefb, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zefb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zefb = 0
			}
			for nextMiss4zefb < maxFields4zefb && (found4zefb[nextMiss4zefb] || decodeMsgFieldSkip4zefb[nextMiss4zefb]) {
				nextMiss4zefb++
			}
			if nextMiss4zefb == maxFields4zefb {
				// filled all the empty fields!
				break doneWithStruct4zefb
			}
			missingFieldsLeft4zefb--
			curField4zefb = nextMiss4zefb
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zefb)
		switch curField4zefb {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found4zefb[0] = true
			var zozp uint32
			zozp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zozp > 0 {
				z.m = make(map[string]*Tr, zozp)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zozp > 0 {
				zozp--
				var zrwf string
				var zgbq *Tr
				zrwf, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zgbq != nil {
						dc.PushAlwaysNil()
						err = zgbq.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zgbq == nil {
						zgbq = new(Tr)
					}
					err = zgbq.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zrwf] = zgbq
			}
		case 1:
			// zid 1 for "s"
			found4zefb[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found4zefb[2] = true
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
	if nextMiss4zefb != -1 {
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
var decodeMsgFieldOrder4zefb = []string{"m", "s", "n"}

var decodeMsgFieldSkip4zefb = []bool{false, false, false}

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
	var empty_zefu [3]bool
	fieldsInUse_zebt := z.fieldsNotEmpty(empty_zefu[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zebt)
	if err != nil {
		return err
	}

	if !empty_zefu[0] {
		// zid 0 for "m"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zrwf, zgbq := range z.m {
			err = en.WriteString(zrwf)
			if err != nil {
				return
			}
			if zgbq == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zgbq.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zefu[1] {
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

	if !empty_zefu[2] {
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
		for zrwf, zgbq := range z.m {
			o = msgp.AppendString(o, zrwf)
			if zgbq == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zgbq.MarshalMsg(o)
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
	const maxFields5ztgj = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5ztgj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5ztgj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5ztgj := totalEncodedFields5ztgj
	missingFieldsLeft5ztgj := maxFields5ztgj - totalEncodedFields5ztgj

	var nextMiss5ztgj int = -1
	var found5ztgj [maxFields5ztgj]bool
	var curField5ztgj int

doneWithStruct5ztgj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5ztgj > 0 || missingFieldsLeft5ztgj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5ztgj, missingFieldsLeft5ztgj, msgp.ShowFound(found5ztgj[:]), unmarshalMsgFieldOrder5ztgj)
		if encodedFieldsLeft5ztgj > 0 {
			encodedFieldsLeft5ztgj--
			curField5ztgj, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5ztgj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5ztgj = 0
			}
			for nextMiss5ztgj < maxFields5ztgj && (found5ztgj[nextMiss5ztgj] || unmarshalMsgFieldSkip5ztgj[nextMiss5ztgj]) {
				nextMiss5ztgj++
			}
			if nextMiss5ztgj == maxFields5ztgj {
				// filled all the empty fields!
				break doneWithStruct5ztgj
			}
			missingFieldsLeft5ztgj--
			curField5ztgj = nextMiss5ztgj
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5ztgj)
		switch curField5ztgj {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found5ztgj[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zqqi uint32
				zqqi, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zqqi > 0 {
					z.m = make(map[string]*Tr, zqqi)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zqqi > 0 {
					var zrwf string
					var zgbq *Tr
					zqqi--
					zrwf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zgbq != nil {
							zgbq.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zgbq {
								zgbq.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zgbq == nil {
								zgbq = new(Tr)
							}
							bts, err = zgbq.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zrwf] = zgbq
				}
			}
		case 1:
			// zid 1 for "s"
			found5ztgj[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found5ztgj[2] = true
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
	if nextMiss5ztgj != -1 {
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
var unmarshalMsgFieldOrder5ztgj = []string{"m", "s", "n"}

var unmarshalMsgFieldSkip5ztgj = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.m != nil {
		for zrwf, zgbq := range z.m {
			_ = zgbq
			_ = zrwf
			s += msgp.StringPrefixSize + len(zrwf)
			if zgbq == nil {
				s += msgp.NilSize
			} else {
				s += zgbq.Msgsize()
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
		0x72, 0x75, 0x63, 0x74, 0x73, 0x83, 0xa2, 0x54, 0x72, 0x82,
		0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
		0x65, 0xa2, 0x54, 0x72, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x73, 0x96, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x55, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x55, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x02, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0x85, 0xa3, 0x5a, 0x69, 0x64, 0xff, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x65,
		0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x2d, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x1c, 0xa4, 0x53, 0x6b, 0x69, 0x70, 0xc3, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x4e, 0x74, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa2, 0x4e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e,
		0x74, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
		0x65, 0x0d, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x0d, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e,
		0x74, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x02, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
		0x53, 0x6e, 0x6f, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e,
		0x6f, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f,
		0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74,
		0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61,
		0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f,
		0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x03, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa6, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa6, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53,
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa3, 0x5a,
		0x69, 0x64, 0x04, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53, 0x65, 0x74, 0x6d,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa4, 0x53, 0x65, 0x74, 0x6d, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xa6, 0x5b, 0x5d, 0x2a, 0x69, 0x6e, 0x6e, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x1a, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x1a, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53,
		0x6c, 0x69, 0x63, 0x65, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53,
		0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69,
		0x6e, 0x6e, 0xa3, 0x69, 0x6e, 0x6e, 0x82, 0xaa, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x69,
		0x6e, 0x6e, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x92,
		0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6a,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x6a, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d,
		0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d,
		0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18,
		0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44,
		0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72,
		0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x65,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0xa1, 0x75, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63,
		0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x75, 0xa6, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x73, 0x93, 0x86, 0xa3, 0x5a, 0x69, 0x64,
		0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6d,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xae, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a, 0x54, 0x72, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d,
		0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61,
		0x6e, 0x67, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c,
		0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74,
		0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72,
		0xa2, 0x54, 0x72, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa1, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x73, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
		0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x02, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02,
		0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x02, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6e, 0xac, 0x46, 0x69, 0x65,
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
	return []byte(`{"SourcePath":"my2.go","SourcePackage":"testdata","ZebraSchemaId":0,"Structs":{"Tr":{"StructName":"Tr","Fields":[{"Zid":0,"FieldGoName":"U","FieldTagName":"U","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"et","FieldTagName":"-","FieldCategory":28,"Skip":true},{"Zid":1,"FieldGoName":"Nt","FieldTagName":"Nt","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}},{"Zid":2,"FieldGoName":"Snot","FieldTagName":"Snot","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":3,"FieldGoName":"Notyet","FieldTagName":"Notyet","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":4,"FieldGoName":"Setm","FieldTagName":"Setm","FieldTypeStr":"[]*inn","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"inn"}}}}]},"inn":{"StructName":"inn","Fields":[{"Zid":0,"FieldGoName":"j","FieldTagName":"j","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":1,"FieldGoName":"e","FieldTagName":"e","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"u":{"StructName":"u","Fields":[{"Zid":0,"FieldGoName":"m","FieldTagName":"m","FieldTypeStr":"map[string]*Tr","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Tr"}}}},{"Zid":1,"FieldGoName":"s","FieldTagName":"s","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"n","FieldTagName":"n","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]}},"Imports":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 6280 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "my2.go",
    "SourcePackage": "testdata",
    "ZebraSchemaId": 0,
    "Structs": {
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
        },
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
        }
    },
    "Imports": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
