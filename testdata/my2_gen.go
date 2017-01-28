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
	const maxFields0zqst = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zqst uint32
	totalEncodedFields0zqst, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zqst := totalEncodedFields0zqst
	missingFieldsLeft0zqst := maxFields0zqst - totalEncodedFields0zqst

	var nextMiss0zqst int = -1
	var found0zqst [maxFields0zqst]bool
	var curField0zqst int

doneWithStruct0zqst:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zqst > 0 || missingFieldsLeft0zqst > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zqst, missingFieldsLeft0zqst, msgp.ShowFound(found0zqst[:]), decodeMsgFieldOrder0zqst)
		if encodedFieldsLeft0zqst > 0 {
			encodedFieldsLeft0zqst--
			curField0zqst, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zqst < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zqst = 0
			}
			for nextMiss0zqst < maxFields0zqst && (found0zqst[nextMiss0zqst] || decodeMsgFieldSkip0zqst[nextMiss0zqst]) {
				nextMiss0zqst++
			}
			if nextMiss0zqst == maxFields0zqst {
				// filled all the empty fields!
				break doneWithStruct0zqst
			}
			missingFieldsLeft0zqst--
			curField0zqst = nextMiss0zqst
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zqst)
		switch curField0zqst {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found0zqst[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "Nt"
			found0zqst[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "Snot"
			found0zqst[3] = true
			var zvwq uint32
			zvwq, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Snot == nil && zvwq > 0 {
				z.Snot = make(map[string]bool, zvwq)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zvwq > 0 {
				zvwq--
				var zlgj string
				var zibz bool
				zlgj, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zibz, err = dc.ReadBool()
				if err != nil {
					panic(err)
				}
				z.Snot[zlgj] = zibz
			}
		case 3:
			// zid 3 for "Notyet"
			found0zqst[4] = true
			var zunz uint32
			zunz, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.Notyet == nil && zunz > 0 {
				z.Notyet = make(map[string]bool, zunz)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zunz > 0 {
				zunz--
				var zork string
				var zdng bool
				zork, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zdng, err = dc.ReadBool()
				if err != nil {
					panic(err)
				}
				z.Notyet[zork] = zdng
			}
		case 4:
			// zid 4 for "Setm"
			found0zqst[5] = true
			var zmhf uint32
			zmhf, err = dc.ReadArrayHeader()
			if err != nil {
				panic(err)
			}
			if cap(z.Setm) >= int(zmhf) {
				z.Setm = (z.Setm)[:zmhf]
			} else {
				z.Setm = make([]*inn, zmhf)
			}
			for zdys := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zdys] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zdys].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zdys] == nil {
						z.Setm[zdys] = new(inn)
					}
					err = z.Setm[zdys].DecodeMsg(dc)
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
	if nextMiss0zqst != -1 {
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
var decodeMsgFieldOrder0zqst = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var decodeMsgFieldSkip0zqst = []bool{false, true, false, false, false, false}

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
	var empty_zymf [6]bool
	fieldsInUse_zqji := z.fieldsNotEmpty(empty_zymf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqji)
	if err != nil {
		return err
	}

	if !empty_zymf[0] {
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

	if !empty_zymf[2] {
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

	if !empty_zymf[3] {
		// zid 2 for "Snot"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			panic(err)
		}
		for zlgj, zibz := range z.Snot {
			err = en.WriteString(zlgj)
			if err != nil {
				panic(err)
			}
			err = en.WriteBool(zibz)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zymf[4] {
		// zid 3 for "Notyet"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			panic(err)
		}
		for zork, zdng := range z.Notyet {
			err = en.WriteString(zork)
			if err != nil {
				panic(err)
			}
			err = en.WriteBool(zdng)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zymf[5] {
		// zid 4 for "Setm"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			panic(err)
		}
		for zdys := range z.Setm {
			if z.Setm[zdys] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zdys].EncodeMsg(en)
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
		for zlgj, zibz := range z.Snot {
			o = msgp.AppendString(o, zlgj)
			o = msgp.AppendBool(o, zibz)
		}
	}

	if !empty[4] {
		// zid 3 for "Notyet"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zork, zdng := range z.Notyet {
			o = msgp.AppendString(o, zork)
			o = msgp.AppendBool(o, zdng)
		}
	}

	if !empty[5] {
		// zid 4 for "Setm"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zdys := range z.Setm {
			if z.Setm[zdys] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zdys].MarshalMsg(o)
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
	const maxFields1zejf = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zejf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zejf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft1zejf := totalEncodedFields1zejf
	missingFieldsLeft1zejf := maxFields1zejf - totalEncodedFields1zejf

	var nextMiss1zejf int = -1
	var found1zejf [maxFields1zejf]bool
	var curField1zejf int

doneWithStruct1zejf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zejf > 0 || missingFieldsLeft1zejf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zejf, missingFieldsLeft1zejf, msgp.ShowFound(found1zejf[:]), unmarshalMsgFieldOrder1zejf)
		if encodedFieldsLeft1zejf > 0 {
			encodedFieldsLeft1zejf--
			curField1zejf, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zejf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zejf = 0
			}
			for nextMiss1zejf < maxFields1zejf && (found1zejf[nextMiss1zejf] || unmarshalMsgFieldSkip1zejf[nextMiss1zejf]) {
				nextMiss1zejf++
			}
			if nextMiss1zejf == maxFields1zejf {
				// filled all the empty fields!
				break doneWithStruct1zejf
			}
			missingFieldsLeft1zejf--
			curField1zejf = nextMiss1zejf
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zejf)
		switch curField1zejf {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found1zejf[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 1:
			// zid 1 for "Nt"
			found1zejf[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "Snot"
			found1zejf[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zncb uint32
				zncb, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Snot == nil && zncb > 0 {
					z.Snot = make(map[string]bool, zncb)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zncb > 0 {
					var zlgj string
					var zibz bool
					zncb--
					zlgj, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zibz, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Snot[zlgj] = zibz
				}
			}
		case 3:
			// zid 3 for "Notyet"
			found1zejf[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zgdu uint32
				zgdu, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.Notyet == nil && zgdu > 0 {
					z.Notyet = make(map[string]bool, zgdu)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zgdu > 0 {
					var zork string
					var zdng bool
					zgdu--
					zork, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zdng, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						panic(err)
					}
					z.Notyet[zork] = zdng
				}
			}
		case 4:
			// zid 4 for "Setm"
			found1zejf[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zdch uint32
				zdch, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if cap(z.Setm) >= int(zdch) {
					z.Setm = (z.Setm)[:zdch]
				} else {
					z.Setm = make([]*inn, zdch)
				}
				for zdys := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zdys] != nil {
							z.Setm[zdys].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zdys] {
								z.Setm[zdys].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zdys] == nil {
								z.Setm[zdys] = new(inn)
							}
							bts, err = z.Setm[zdys].UnmarshalMsg(bts)
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
	if nextMiss1zejf != -1 {
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
var unmarshalMsgFieldOrder1zejf = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var unmarshalMsgFieldSkip1zejf = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.U) + 3 + msgp.IntSize + 5 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zlgj, zibz := range z.Snot {
			_ = zibz
			_ = zlgj
			s += msgp.StringPrefixSize + len(zlgj) + msgp.BoolSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zork, zdng := range z.Notyet {
			_ = zdng
			_ = zork
			s += msgp.StringPrefixSize + len(zork) + msgp.BoolSize
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zdys := range z.Setm {
		if z.Setm[zdys] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zdys].Msgsize()
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
	const maxFields2zwtk = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zwtk uint32
	totalEncodedFields2zwtk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zwtk := totalEncodedFields2zwtk
	missingFieldsLeft2zwtk := maxFields2zwtk - totalEncodedFields2zwtk

	var nextMiss2zwtk int = -1
	var found2zwtk [maxFields2zwtk]bool
	var curField2zwtk int

doneWithStruct2zwtk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zwtk > 0 || missingFieldsLeft2zwtk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zwtk, missingFieldsLeft2zwtk, msgp.ShowFound(found2zwtk[:]), decodeMsgFieldOrder2zwtk)
		if encodedFieldsLeft2zwtk > 0 {
			encodedFieldsLeft2zwtk--
			curField2zwtk, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zwtk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zwtk = 0
			}
			for nextMiss2zwtk < maxFields2zwtk && (found2zwtk[nextMiss2zwtk] || decodeMsgFieldSkip2zwtk[nextMiss2zwtk]) {
				nextMiss2zwtk++
			}
			if nextMiss2zwtk == maxFields2zwtk {
				// filled all the empty fields!
				break doneWithStruct2zwtk
			}
			missingFieldsLeft2zwtk--
			curField2zwtk = nextMiss2zwtk
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zwtk)
		switch curField2zwtk {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found2zwtk[0] = true
			var zcpu uint32
			zcpu, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.j == nil && zcpu > 0 {
				z.j = make(map[string]bool, zcpu)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zcpu > 0 {
				zcpu--
				var zljy string
				var zkea bool
				zljy, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				zkea, err = dc.ReadBool()
				if err != nil {
					panic(err)
				}
				z.j[zljy] = zkea
			}
		case 1:
			// zid 1 for "e"
			found2zwtk[1] = true
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
	if nextMiss2zwtk != -1 {
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
var decodeMsgFieldOrder2zwtk = []string{"j", "e"}

var decodeMsgFieldSkip2zwtk = []bool{false, false}

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
	var empty_zzoj [2]bool
	fieldsInUse_zthv := z.fieldsNotEmpty(empty_zzoj[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zthv)
	if err != nil {
		return err
	}

	if !empty_zzoj[0] {
		// zid 0 for "j"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			panic(err)
		}
		for zljy, zkea := range z.j {
			err = en.WriteString(zljy)
			if err != nil {
				panic(err)
			}
			err = en.WriteBool(zkea)
			if err != nil {
				panic(err)
			}
		}
	}

	if !empty_zzoj[1] {
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
		for zljy, zkea := range z.j {
			o = msgp.AppendString(o, zljy)
			o = msgp.AppendBool(o, zkea)
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
	const maxFields3zmhy = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zmhy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zmhy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft3zmhy := totalEncodedFields3zmhy
	missingFieldsLeft3zmhy := maxFields3zmhy - totalEncodedFields3zmhy

	var nextMiss3zmhy int = -1
	var found3zmhy [maxFields3zmhy]bool
	var curField3zmhy int

doneWithStruct3zmhy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zmhy > 0 || missingFieldsLeft3zmhy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zmhy, missingFieldsLeft3zmhy, msgp.ShowFound(found3zmhy[:]), unmarshalMsgFieldOrder3zmhy)
		if encodedFieldsLeft3zmhy > 0 {
			encodedFieldsLeft3zmhy--
			curField3zmhy, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zmhy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zmhy = 0
			}
			for nextMiss3zmhy < maxFields3zmhy && (found3zmhy[nextMiss3zmhy] || unmarshalMsgFieldSkip3zmhy[nextMiss3zmhy]) {
				nextMiss3zmhy++
			}
			if nextMiss3zmhy == maxFields3zmhy {
				// filled all the empty fields!
				break doneWithStruct3zmhy
			}
			missingFieldsLeft3zmhy--
			curField3zmhy = nextMiss3zmhy
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zmhy)
		switch curField3zmhy {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found3zmhy[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zcjy uint32
				zcjy, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.j == nil && zcjy > 0 {
					z.j = make(map[string]bool, zcjy)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zcjy > 0 {
					var zljy string
					var zkea bool
					zcjy--
					zljy, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					zkea, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						panic(err)
					}
					z.j[zljy] = zkea
				}
			}
		case 1:
			// zid 1 for "e"
			found3zmhy[1] = true
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
	if nextMiss3zmhy != -1 {
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
var unmarshalMsgFieldOrder3zmhy = []string{"j", "e"}

var unmarshalMsgFieldSkip3zmhy = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.j != nil {
		for zljy, zkea := range z.j {
			_ = zkea
			_ = zljy
			s += msgp.StringPrefixSize + len(zljy) + msgp.BoolSize
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
	const maxFields4znxd = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4znxd uint32
	totalEncodedFields4znxd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4znxd := totalEncodedFields4znxd
	missingFieldsLeft4znxd := maxFields4znxd - totalEncodedFields4znxd

	var nextMiss4znxd int = -1
	var found4znxd [maxFields4znxd]bool
	var curField4znxd int

doneWithStruct4znxd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4znxd > 0 || missingFieldsLeft4znxd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4znxd, missingFieldsLeft4znxd, msgp.ShowFound(found4znxd[:]), decodeMsgFieldOrder4znxd)
		if encodedFieldsLeft4znxd > 0 {
			encodedFieldsLeft4znxd--
			curField4znxd, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4znxd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4znxd = 0
			}
			for nextMiss4znxd < maxFields4znxd && (found4znxd[nextMiss4znxd] || decodeMsgFieldSkip4znxd[nextMiss4znxd]) {
				nextMiss4znxd++
			}
			if nextMiss4znxd == maxFields4znxd {
				// filled all the empty fields!
				break doneWithStruct4znxd
			}
			missingFieldsLeft4znxd--
			curField4znxd = nextMiss4znxd
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4znxd)
		switch curField4znxd {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found4znxd[0] = true
			var zdub uint32
			zdub, err = dc.ReadMapHeader()
			if err != nil {
				panic(err)
			}
			if z.m == nil && zdub > 0 {
				z.m = make(map[string]*Tr, zdub)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zdub > 0 {
				zdub--
				var zhkg string
				var zmcf *Tr
				zhkg, err = dc.ReadString()
				if err != nil {
					panic(err)
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zmcf != nil {
						dc.PushAlwaysNil()
						err = zmcf.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zmcf == nil {
						zmcf = new(Tr)
					}
					err = zmcf.DecodeMsg(dc)
					if err != nil {
						panic(err)
					}
				}
				z.m[zhkg] = zmcf
			}
		case 1:
			// zid 1 for "s"
			found4znxd[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "n"
			found4znxd[2] = true
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
	if nextMiss4znxd != -1 {
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
var decodeMsgFieldOrder4znxd = []string{"m", "s", "n"}

var decodeMsgFieldSkip4znxd = []bool{false, false, false}

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
	var empty_zvmk [3]bool
	fieldsInUse_zjfp := z.fieldsNotEmpty(empty_zvmk[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjfp)
	if err != nil {
		return err
	}

	if !empty_zvmk[0] {
		// zid 0 for "m"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			panic(err)
		}
		for zhkg, zmcf := range z.m {
			err = en.WriteString(zhkg)
			if err != nil {
				panic(err)
			}
			if zmcf == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zmcf.EncodeMsg(en)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if !empty_zvmk[1] {
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

	if !empty_zvmk[2] {
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
		for zhkg, zmcf := range z.m {
			o = msgp.AppendString(o, zhkg)
			if zmcf == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zmcf.MarshalMsg(o)
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
	const maxFields5znig = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5znig uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5znig, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			panic(err)
			return
		}
	}
	encodedFieldsLeft5znig := totalEncodedFields5znig
	missingFieldsLeft5znig := maxFields5znig - totalEncodedFields5znig

	var nextMiss5znig int = -1
	var found5znig [maxFields5znig]bool
	var curField5znig int

doneWithStruct5znig:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5znig > 0 || missingFieldsLeft5znig > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5znig, missingFieldsLeft5znig, msgp.ShowFound(found5znig[:]), unmarshalMsgFieldOrder5znig)
		if encodedFieldsLeft5znig > 0 {
			encodedFieldsLeft5znig--
			curField5znig, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				panic(err)
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5znig < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5znig = 0
			}
			for nextMiss5znig < maxFields5znig && (found5znig[nextMiss5znig] || unmarshalMsgFieldSkip5znig[nextMiss5znig]) {
				nextMiss5znig++
			}
			if nextMiss5znig == maxFields5znig {
				// filled all the empty fields!
				break doneWithStruct5znig
			}
			missingFieldsLeft5znig--
			curField5znig = nextMiss5znig
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5znig)
		switch curField5znig {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found5znig[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zfst uint32
				zfst, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					panic(err)
				}
				if z.m == nil && zfst > 0 {
					z.m = make(map[string]*Tr, zfst)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zfst > 0 {
					var zhkg string
					var zmcf *Tr
					zfst--
					zhkg, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						panic(err)
					}
					if nbs.AlwaysNil {
						if zmcf != nil {
							zmcf.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zmcf {
								zmcf.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zmcf == nil {
								zmcf = new(Tr)
							}
							bts, err = zmcf.UnmarshalMsg(bts)
							if err != nil {
								panic(err)
							}
							if err != nil {
								panic(err)
							}
						}
					}
					z.m[zhkg] = zmcf
				}
			}
		case 1:
			// zid 1 for "s"
			found5znig[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				panic(err)
			}
		case 2:
			// zid 2 for "n"
			found5znig[2] = true
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
	if nextMiss5znig != -1 {
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
var unmarshalMsgFieldOrder5znig = []string{"m", "s", "n"}

var unmarshalMsgFieldSkip5znig = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.m != nil {
		for zhkg, zmcf := range z.m {
			_ = zmcf
			_ = zhkg
			s += msgp.StringPrefixSize + len(zhkg)
			if zmcf == nil {
				s += msgp.NilSize
			} else {
				s += zmcf.Msgsize()
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
		0x72, 0x75, 0x63, 0x74, 0x73, 0x83, 0xa3, 0x69, 0x6e, 0x6e,
		0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0xa3, 0x69, 0x6e, 0x6e, 0xa6, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x73, 0x92, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0xa1, 0x6a, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6a, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d,
		0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61,
		0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12,
		0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x87,
		0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x65, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa1, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xa1, 0x75, 0x82, 0xaa, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x75, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x93, 0x86,
		0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6d, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa1, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xae, 0x6d, 0x61,
		0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a,
		0x54, 0x72, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53,
		0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02,
		0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50,
		0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16,
		0xa3, 0x53, 0x74, 0x72, 0xa2, 0x54, 0x72, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x73, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6e,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0xa2, 0x54, 0x72, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x54, 0x72, 0xa6,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x55, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x55, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x85, 0xa3, 0x5a, 0x69, 0x64,
		0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa2, 0x65, 0x74, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x2d, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xa4, 0x53, 0x6b, 0x69,
		0x70, 0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa2, 0x4e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x4e, 0x74, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0d, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0d, 0xa3, 0x53,
		0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xac, 0x46, 0x69,
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
		0x69, 0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74, 0x79,
		0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74, 0x79,
		0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
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
		0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa4, 0x53, 0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53,
		0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x5b, 0x5d, 0x2a,
		0x69, 0x6e, 0x6e, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xa3,
		0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xa6,
		0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x83, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f,
		0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61,
		0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x6e, 0xa7, 0x49, 0x6d,
		0x70, 0x6f, 0x72, 0x74, 0x73, 0x91, 0xbd, 0x22, 0x67, 0x69,
		0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
		0x6c, 0x79, 0x63, 0x65, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x72,
		0x62, 0x74, 0x72, 0x65, 0x65, 0x22,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 2153 bytes
func (FileMy2) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"my2.go","SourcePackage":"testdata","ZebraSchemaId":0,"Structs":{"inn":{"StructName":"inn","Fields":[{"Zid":0,"FieldGoName":"j","FieldTagName":"j","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":1,"FieldGoName":"e","FieldTagName":"e","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"u":{"StructName":"u","Fields":[{"Zid":0,"FieldGoName":"m","FieldTagName":"m","FieldTypeStr":"map[string]*Tr","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Tr"}}}},{"Zid":1,"FieldGoName":"s","FieldTagName":"s","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"n","FieldTagName":"n","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Tr":{"StructName":"Tr","Fields":[{"Zid":0,"FieldGoName":"U","FieldTagName":"U","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"et","FieldTagName":"-","FieldCategory":28,"Skip":true},{"Zid":1,"FieldGoName":"Nt","FieldTagName":"Nt","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}},{"Zid":2,"FieldGoName":"Snot","FieldTagName":"Snot","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":3,"FieldGoName":"Notyet","FieldTagName":"Notyet","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":4,"FieldGoName":"Setm","FieldTagName":"Setm","FieldTypeStr":"[]*inn","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"inn"}}}}]}},"Imports":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 6280 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "my2.go",
    "SourcePackage": "testdata",
    "ZebraSchemaId": 0,
    "Structs": {
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
        }
    },
    "Imports": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
