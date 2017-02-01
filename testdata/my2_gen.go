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
	const maxFields0zkmu = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zkmu uint32
	totalEncodedFields0zkmu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zkmu := totalEncodedFields0zkmu
	missingFieldsLeft0zkmu := maxFields0zkmu - totalEncodedFields0zkmu

	var nextMiss0zkmu int = -1
	var found0zkmu [maxFields0zkmu]bool
	var curField0zkmu int

doneWithStruct0zkmu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zkmu > 0 || missingFieldsLeft0zkmu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zkmu, missingFieldsLeft0zkmu, msgp.ShowFound(found0zkmu[:]), decodeMsgFieldOrder0zkmu)
		if encodedFieldsLeft0zkmu > 0 {
			encodedFieldsLeft0zkmu--
			curField0zkmu, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zkmu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zkmu = 0
			}
			for nextMiss0zkmu < maxFields0zkmu && (found0zkmu[nextMiss0zkmu] || decodeMsgFieldSkip0zkmu[nextMiss0zkmu]) {
				nextMiss0zkmu++
			}
			if nextMiss0zkmu == maxFields0zkmu {
				// filled all the empty fields!
				break doneWithStruct0zkmu
			}
			missingFieldsLeft0zkmu--
			curField0zkmu = nextMiss0zkmu
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zkmu)
		switch curField0zkmu {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found0zkmu[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "Nt"
			found0zkmu[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "Snot"
			found0zkmu[3] = true
			var ziju uint32
			ziju, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Snot == nil && ziju > 0 {
				z.Snot = make(map[string]bool, ziju)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for ziju > 0 {
				ziju--
				var zsih string
				var zbep bool
				zsih, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zbep, err = dc.ReadBool()
				if err != nil {
					panic(err)
				}
				z.Snot[zsih] = zbep
			}
		case 3:
			// zid 3 for "Notyet"
			found0zkmu[4] = true
			var ztix uint32
			ztix, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Notyet == nil && ztix > 0 {
				z.Notyet = make(map[string]bool, ztix)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for ztix > 0 {
				ztix--
				var zwhg string
				var zghx bool
				zwhg, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zghx, err = dc.ReadBool()
				if err != nil {
					panic(err)
				}
				z.Notyet[zwhg] = zghx
			}
		case 4:
			// zid 4 for "Setm"
			found0zkmu[5] = true
			var zjgm uint32
			zjgm, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Setm) >= int(zjgm) {
				z.Setm = (z.Setm)[:zjgm]
			} else {
				z.Setm = make([]*inn, zjgm)
			}
			for zbce := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zbce] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zbce].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zbce] == nil {
						z.Setm[zbce] = new(inn)
					}
					err = z.Setm[zbce].DecodeMsg(dc)
					if err != nil {
						panic(err)
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				panic(err)
			}
		}
	}
	if nextMiss0zkmu != -1 {
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
var decodeMsgFieldOrder0zkmu = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var decodeMsgFieldSkip0zkmu = []bool{false, true, false, false, false, false}

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
	var empty_zwob [6]bool
	fieldsInUse_zjjd := z.fieldsNotEmpty(empty_zwob[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjjd)
	if err != nil {
		return err
	}

	if !empty_zwob[0] {
		// zid 0 for "U"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteString(z.U)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zwob[2] {
		// zid 1 for "Nt"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Nt)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zwob[3] {
		// zid 2 for "Snot"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			panic(err)
		}
		for zsih, zbep := range z.Snot {
			err = en.WriteString(zsih)
			if err != nil {
				panic(err)
			}
			err = en.WriteBool(zbep)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zwob[4] {
		// zid 3 for "Notyet"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			panic(err)
		}
		for zwhg, zghx := range z.Notyet {
			err = en.WriteString(zwhg)
			if err != nil {
				panic(err)
			}
			err = en.WriteBool(zghx)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zwob[5] {
		// zid 4 for "Setm"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			panic(err)
		}
		for zbce := range z.Setm {
			if z.Setm[zbce] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zbce].EncodeMsg(en)
				if err != nil {
					panic(err)
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
		for zsih, zbep := range z.Snot {
			o = msgp.AppendString(o, zsih)
			o = msgp.AppendBool(o, zbep)
		}
	}

	if !empty[4] {
		// zid 3 for "Notyet"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zwhg, zghx := range z.Notyet {
			o = msgp.AppendString(o, zwhg)
			o = msgp.AppendBool(o, zghx)
		}
	}

	if !empty[5] {
		// zid 4 for "Setm"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zbce := range z.Setm {
			if z.Setm[zbce] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zbce].MarshalMsg(o)
				if err != nil {
					panic(err)
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
	const maxFields1zfmm = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zfmm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zfmm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			//return
		}
	}
	encodedFieldsLeft1zfmm := totalEncodedFields1zfmm
	missingFieldsLeft1zfmm := maxFields1zfmm - totalEncodedFields1zfmm

	var nextMiss1zfmm int = -1
	var found1zfmm [maxFields1zfmm]bool
	var curField1zfmm int

doneWithStruct1zfmm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zfmm > 0 || missingFieldsLeft1zfmm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zfmm, missingFieldsLeft1zfmm, msgp.ShowFound(found1zfmm[:]), unmarshalMsgFieldOrder1zfmm)
		if encodedFieldsLeft1zfmm > 0 {
			encodedFieldsLeft1zfmm--
			curField1zfmm, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				//return
			}
		} else {
			//missing fields need handling
			if nextMiss1zfmm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zfmm = 0
			}
			for nextMiss1zfmm < maxFields1zfmm && (found1zfmm[nextMiss1zfmm] || unmarshalMsgFieldSkip1zfmm[nextMiss1zfmm]) {
				nextMiss1zfmm++
			}
			if nextMiss1zfmm == maxFields1zfmm {
				// filled all the empty fields!
				break doneWithStruct1zfmm
			}
			missingFieldsLeft1zfmm--
			curField1zfmm = nextMiss1zfmm
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zfmm)
		switch curField1zfmm {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found1zfmm[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "Nt"
			found1zfmm[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "Snot"
			found1zfmm[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zxnx uint32
				zxnx, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Snot == nil && zxnx > 0 {
					z.Snot = make(map[string]bool, zxnx)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zxnx > 0 {
					var zsih string
					var zbep bool
					zxnx--
					zsih, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zbep, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Snot[zsih] = zbep
				}
			}
		case 3:
			// zid 3 for "Notyet"
			found1zfmm[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zndo uint32
				zndo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Notyet == nil && zndo > 0 {
					z.Notyet = make(map[string]bool, zndo)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zndo > 0 {
					var zwhg string
					var zghx bool
					zndo--
					zwhg, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zghx, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Notyet[zwhg] = zghx
				}
			}
		case 4:
			// zid 4 for "Setm"
			found1zfmm[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zuwr uint32
				zuwr, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Setm) >= int(zuwr) {
					z.Setm = (z.Setm)[:zuwr]
				} else {
					z.Setm = make([]*inn, zuwr)
				}
				for zbce := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zbce] != nil {
							z.Setm[zbce].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zbce] {
								z.Setm[zbce].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zbce] == nil {
								z.Setm[zbce] = new(inn)
							}
							bts, err = z.Setm[zbce].UnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
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
	if nextMiss1zfmm != -1 {
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
var unmarshalMsgFieldOrder1zfmm = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var unmarshalMsgFieldSkip1zfmm = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.U) + 3 + msgp.IntSize + 5 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zsih, zbep := range z.Snot {
			_ = zbep
			_ = zsih
			s += msgp.StringPrefixSize + len(zsih) + msgp.BoolSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zwhg, zghx := range z.Notyet {
			_ = zghx
			_ = zwhg
			s += msgp.StringPrefixSize + len(zwhg) + msgp.BoolSize
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zbce := range z.Setm {
		if z.Setm[zbce] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zbce].Msgsize()
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
	const maxFields2zbar = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zbar uint32
	totalEncodedFields2zbar, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zbar := totalEncodedFields2zbar
	missingFieldsLeft2zbar := maxFields2zbar - totalEncodedFields2zbar

	var nextMiss2zbar int = -1
	var found2zbar [maxFields2zbar]bool
	var curField2zbar int

doneWithStruct2zbar:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zbar > 0 || missingFieldsLeft2zbar > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zbar, missingFieldsLeft2zbar, msgp.ShowFound(found2zbar[:]), decodeMsgFieldOrder2zbar)
		if encodedFieldsLeft2zbar > 0 {
			encodedFieldsLeft2zbar--
			curField2zbar, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zbar < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zbar = 0
			}
			for nextMiss2zbar < maxFields2zbar && (found2zbar[nextMiss2zbar] || decodeMsgFieldSkip2zbar[nextMiss2zbar]) {
				nextMiss2zbar++
			}
			if nextMiss2zbar == maxFields2zbar {
				// filled all the empty fields!
				break doneWithStruct2zbar
			}
			missingFieldsLeft2zbar--
			curField2zbar = nextMiss2zbar
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zbar)
		switch curField2zbar {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found2zbar[0] = true
			var zbfj uint32
			zbfj, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.j == nil && zbfj > 0 {
				z.j = make(map[string]bool, zbfj)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zbfj > 0 {
				zbfj--
				var zwnl string
				var zpnc bool
				zwnl, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zpnc, err = dc.ReadBool()
				if err != nil {
					panic(err)
				}
				z.j[zwnl] = zpnc
			}
		case 1:
			// zid 1 for "e"
			found2zbar[1] = true
			z.e, err = dc.ReadInt64()
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
	if nextMiss2zbar != -1 {
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
var decodeMsgFieldOrder2zbar = []string{"j", "e"}

var decodeMsgFieldSkip2zbar = []bool{false, false}

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
	var empty_zobi [2]bool
	fieldsInUse_zybt := z.fieldsNotEmpty(empty_zobi[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zybt)
	if err != nil {
		return err
	}

	if !empty_zobi[0] {
		// zid 0 for "j"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			panic(err)
		}
		for zwnl, zpnc := range z.j {
			err = en.WriteString(zwnl)
			if err != nil {
				panic(err)
			}
			err = en.WriteBool(zpnc)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zobi[1] {
		// zid 1 for "e"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.e)
		if err != nil {
			panic(err)
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
		for zwnl, zpnc := range z.j {
			o = msgp.AppendString(o, zwnl)
			o = msgp.AppendBool(o, zpnc)
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
	const maxFields3zhns = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zhns uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zhns, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			//return
		}
	}
	encodedFieldsLeft3zhns := totalEncodedFields3zhns
	missingFieldsLeft3zhns := maxFields3zhns - totalEncodedFields3zhns

	var nextMiss3zhns int = -1
	var found3zhns [maxFields3zhns]bool
	var curField3zhns int

doneWithStruct3zhns:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zhns > 0 || missingFieldsLeft3zhns > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zhns, missingFieldsLeft3zhns, msgp.ShowFound(found3zhns[:]), unmarshalMsgFieldOrder3zhns)
		if encodedFieldsLeft3zhns > 0 {
			encodedFieldsLeft3zhns--
			curField3zhns, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				//return
			}
		} else {
			//missing fields need handling
			if nextMiss3zhns < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zhns = 0
			}
			for nextMiss3zhns < maxFields3zhns && (found3zhns[nextMiss3zhns] || unmarshalMsgFieldSkip3zhns[nextMiss3zhns]) {
				nextMiss3zhns++
			}
			if nextMiss3zhns == maxFields3zhns {
				// filled all the empty fields!
				break doneWithStruct3zhns
			}
			missingFieldsLeft3zhns--
			curField3zhns = nextMiss3zhns
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zhns)
		switch curField3zhns {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found3zhns[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zbvs uint32
				zbvs, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.j == nil && zbvs > 0 {
					z.j = make(map[string]bool, zbvs)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zbvs > 0 {
					var zwnl string
					var zpnc bool
					zbvs--
					zwnl, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zpnc, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						panic(err)
					}
					z.j[zwnl] = zpnc
				}
			}
		case 1:
			// zid 1 for "e"
			found3zhns[1] = true
			z.e, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss3zhns != -1 {
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
var unmarshalMsgFieldOrder3zhns = []string{"j", "e"}

var unmarshalMsgFieldSkip3zhns = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.j != nil {
		for zwnl, zpnc := range z.j {
			_ = zpnc
			_ = zwnl
			s += msgp.StringPrefixSize + len(zwnl) + msgp.BoolSize
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
	const maxFields4zzyl = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zzyl uint32
	totalEncodedFields4zzyl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zzyl := totalEncodedFields4zzyl
	missingFieldsLeft4zzyl := maxFields4zzyl - totalEncodedFields4zzyl

	var nextMiss4zzyl int = -1
	var found4zzyl [maxFields4zzyl]bool
	var curField4zzyl int

doneWithStruct4zzyl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zzyl > 0 || missingFieldsLeft4zzyl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zzyl, missingFieldsLeft4zzyl, msgp.ShowFound(found4zzyl[:]), decodeMsgFieldOrder4zzyl)
		if encodedFieldsLeft4zzyl > 0 {
			encodedFieldsLeft4zzyl--
			curField4zzyl, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zzyl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zzyl = 0
			}
			for nextMiss4zzyl < maxFields4zzyl && (found4zzyl[nextMiss4zzyl] || decodeMsgFieldSkip4zzyl[nextMiss4zzyl]) {
				nextMiss4zzyl++
			}
			if nextMiss4zzyl == maxFields4zzyl {
				// filled all the empty fields!
				break doneWithStruct4zzyl
			}
			missingFieldsLeft4zzyl--
			curField4zzyl = nextMiss4zzyl
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zzyl)
		switch curField4zzyl {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found4zzyl[0] = true
			var zyaa uint32
			zyaa, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.m == nil && zyaa > 0 {
				z.m = make(map[string]*Tr, zyaa)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zyaa > 0 {
				zyaa--
				var zggc string
				var zbgh *Tr
				zggc, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zbgh != nil {
						dc.PushAlwaysNil()
						err = zbgh.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zbgh == nil {
						zbgh = new(Tr)
					}
					err = zbgh.DecodeMsg(dc)
					if err != nil {
						panic(err)
					}
				}
				z.m[zggc] = zbgh
			}
		case 1:
			// zid 1 for "s"
			found4zzyl[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "n"
			found4zzyl[2] = true
			z.n, err = dc.ReadInt64()
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
	if nextMiss4zzyl != -1 {
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
var decodeMsgFieldOrder4zzyl = []string{"m", "s", "n"}

var decodeMsgFieldSkip4zzyl = []bool{false, false, false}

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
	var empty_zddf [3]bool
	fieldsInUse_zatv := z.fieldsNotEmpty(empty_zddf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zatv)
	if err != nil {
		return err
	}

	if !empty_zddf[0] {
		// zid 0 for "m"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			panic(err)
		}
		for zggc, zbgh := range z.m {
			err = en.WriteString(zggc)
			if err != nil {
				panic(err)
			}
			if zbgh == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zbgh.EncodeMsg(en)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if !empty_zddf[1] {
		// zid 1 for "s"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.s)
		if err != nil {
			panic(err)
		}
	}

	if !empty_zddf[2] {
		// zid 2 for "n"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.n)
		if err != nil {
			panic(err)
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
		for zggc, zbgh := range z.m {
			o = msgp.AppendString(o, zggc)
			if zbgh == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zbgh.MarshalMsg(o)
				if err != nil {
					panic(err)
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
	const maxFields5zroc = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zroc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zroc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			//return
		}
	}
	encodedFieldsLeft5zroc := totalEncodedFields5zroc
	missingFieldsLeft5zroc := maxFields5zroc - totalEncodedFields5zroc

	var nextMiss5zroc int = -1
	var found5zroc [maxFields5zroc]bool
	var curField5zroc int

doneWithStruct5zroc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zroc > 0 || missingFieldsLeft5zroc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zroc, missingFieldsLeft5zroc, msgp.ShowFound(found5zroc[:]), unmarshalMsgFieldOrder5zroc)
		if encodedFieldsLeft5zroc > 0 {
			encodedFieldsLeft5zroc--
			curField5zroc, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				//return
			}
		} else {
			//missing fields need handling
			if nextMiss5zroc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zroc = 0
			}
			for nextMiss5zroc < maxFields5zroc && (found5zroc[nextMiss5zroc] || unmarshalMsgFieldSkip5zroc[nextMiss5zroc]) {
				nextMiss5zroc++
			}
			if nextMiss5zroc == maxFields5zroc {
				// filled all the empty fields!
				break doneWithStruct5zroc
			}
			missingFieldsLeft5zroc--
			curField5zroc = nextMiss5zroc
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zroc)
		switch curField5zroc {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found5zroc[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zhwq uint32
				zhwq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.m == nil && zhwq > 0 {
					z.m = make(map[string]*Tr, zhwq)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zhwq > 0 {
					var zggc string
					var zbgh *Tr
					zhwq--
					zggc, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zbgh != nil {
							zbgh.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zbgh {
								zbgh.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zbgh == nil {
								zbgh = new(Tr)
							}
							bts, err = zbgh.UnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.m[zggc] = zbgh
				}
			}
		case 1:
			// zid 1 for "s"
			found5zroc[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "n"
			found5zroc[2] = true
			z.n, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss5zroc != -1 {
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
var unmarshalMsgFieldOrder5zroc = []string{"m", "s", "n"}

var unmarshalMsgFieldSkip5zroc = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.m != nil {
		for zggc, zbgh := range z.m {
			_ = zbgh
			_ = zggc
			s += msgp.StringPrefixSize + len(zggc)
			if zbgh == nil {
				s += msgp.NilSize
			} else {
				s += zbgh.Msgsize()
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
