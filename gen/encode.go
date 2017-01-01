package gen

import (
	"fmt"
	"io"
	"os"

	"github.com/glycerine/zebrapack/cfg"
	"github.com/glycerine/zebrapack/msgp"
)

func encode(w io.Writer, cfg *cfg.ZebraConfig) *encodeGen {
	return &encodeGen{
		p:   printer{w: w},
		cfg: cfg,
	}
}

type encodeGen struct {
	passes
	p    printer
	fuse []byte
	cfg  *cfg.ZebraConfig
}

func (e *encodeGen) Method() Method { return Encode }

func (e *encodeGen) Apply(dirs []string) error {
	return nil
}

func (e *encodeGen) writeAndCheck(typ string, argfmt string, arg interface{}) {
	e.p.printf("\nerr = en.Write%s(%s)", typ, fmt.Sprintf(argfmt, arg))
	e.p.print(errcheck)
}

func (e *encodeGen) fuseHook() {
	if len(e.fuse) > 0 {
		e.appendraw(e.fuse)
		e.fuse = e.fuse[:0]
	}
}

func (e *encodeGen) Fuse(b []byte) {
	if len(e.fuse) > 0 {
		e.fuse = append(e.fuse, b...)
	} else {
		e.fuse = b
	}
}

func (e *encodeGen) Execute(p Elem) error {
	if !e.p.ok() {
		return e.p.err
	}
	p = e.applyall(p)
	if p == nil {
		return nil
	}
	if !IsPrintable(p) {
		return nil
	}

	e.p.comment("EncodeMsg implements msgp.Encodable")

	e.p.printf("\nfunc (%s %s) EncodeMsg(en *msgp.Writer) (err error) {", p.Varname(), imutMethodReceiver(p))
	next(e, p)
	e.p.nakedReturn()
	return e.p.err
}

func (e *encodeGen) gStruct(s *Struct) {
	if !e.p.ok() {
		return
	}
	if s.AsTuple {
		e.tuple(s)
	} else {
		e.structmap(s)
	}
	return
}

func (e *encodeGen) tuple(s *Struct) {
	nfields := len(s.Fields) - s.SkipCount
	data := msgp.AppendArrayHeader(nil, uint32(nfields))
	e.p.printf("\n// array header, size %d", nfields)
	e.Fuse(data)
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if !e.p.ok() {
			return
		}
		next(e, s.Fields[i].FieldElem)
	}
}

func (e *encodeGen) appendraw(bts []byte) {
	e.p.print("\nerr = en.Append(")
	for i, b := range bts {
		if i != 0 {
			e.p.print(", ")
		}
		e.p.printf("0x%x", b)
	}
	e.p.print(")\nif err != nil { return err }")
}

func (e *encodeGen) structmap(s *Struct) {
	nfields := len(s.Fields) - s.SkipCount
	var data []byte
	fast := !e.cfg.UseMsgp2
	empty := "empty_" + randIdent()
	inUse := "fieldsInUse_" + randIdent()
	// if fast, then always omit-empty.
	if fast || s.hasOmitEmptyTags {
		e.p.printf("\n\n// honor the omitempty tags\n")
		e.p.printf("var %s [%d]bool\n", empty, len(s.Fields))
		e.p.printf("%s := %s.fieldsNotEmpty(%s[:])\n",
			inUse, s.vname, empty)
		e.p.printf("\n// map header\n")
		if fast {
			// +1 for the -1:structName type-identifier.
			e.p.printf("	err = en.WriteMapHeader(%s+1)\n", inUse)
		} else {
			e.p.printf("	err = en.WriteMapHeader(%s)\n", inUse)
		}
		e.p.printf("	if err != nil {\n")
		e.p.printf("		return err\n}\n")
	} else {
		data = msgp.AppendMapHeader(nil, uint32(nfields))
		e.p.printf("\n// map header, size %d", nfields)
		e.Fuse(data)
	}

	if fast && !e.cfg.NoRTTI {
		// record the struct name under integer key -1
		recv := s.TypeName() // imutMethodReceiver(s)
		e.p.printf("\n// runtime struct type identification for '%s'\n", recv)
		hexname := ""
		for i := range recv {
			hexname += fmt.Sprintf("0x%x,", recv[i])
		}
		e.p.printf("err = en.Append(0xff)\n")
		e.p.printf("	if err != nil {\n")
		e.p.printf("		return err\n}\n")

		e.p.printf("err = en.WriteStringFromBytes([]byte{%s})\n", hexname)
		e.p.printf("	if err != nil {\n")
		e.p.printf("		return err\n}\n")
	}

	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if !e.p.ok() {
			return
		}

		if fast || (s.hasOmitEmptyTags && s.Fields[i].OmitEmpty) {
			e.p.printf("\n if !%s[%d] {", empty, i)
		}
		if fast {
			// sanity check
			if s.Fields[i].ZebraId < 0 {
				fmt.Fprintf(os.Stderr, "\nzebrapack error: field '%s' is missing a zid number; cannot proceed under -fast\n", s.Fields[i].FieldTag)
				os.Exit(1)
			}
			// proceed
			data = msgp.AppendInt64(nil, s.Fields[i].ZebraId)
			e.p.printf("\n// zid %v for %q", s.Fields[i].ZebraId,
				s.Fields[i].FieldTag)
		} else {
			data = msgp.AppendString(nil, s.Fields[i].FieldTag)
			e.p.printf("\n// write %q", s.Fields[i].FieldTag)
		}
		e.Fuse(data)
		next(e, s.Fields[i].FieldElem)

		if fast || (s.hasOmitEmptyTags && s.Fields[i].OmitEmpty) {
			e.p.printf("\n }\n")
		}
	}
}

func (e *encodeGen) gMap(m *Map) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := m.Varname()
	e.writeAndCheck(mapHeader, lenAsUint32, vname)

	e.p.printf("\nfor %s, %s := range %s {", m.Keyidx, m.Validx, vname)
	e.writeAndCheck(m.KeyTyp, literalFmt, m.Keyidx)
	next(e, m.Value)
	e.p.closeblock()
}

func (e *encodeGen) gPtr(s *Ptr) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.p.printf("\nif %s == nil { err = en.WriteNil(); if err != nil { return; } } else {", s.Varname())
	next(e, s.Value)
	e.p.closeblock()
}

func (e *encodeGen) gSlice(s *Slice) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.writeAndCheck(arrayHeader, lenAsUint32, s.Varname())
	e.p.rangeBlock(s.Index, s.Varname(), e, s.Els)
}

func (e *encodeGen) gArray(a *Array) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	// shortcut for [const]byte
	if be, ok := a.Els.(*BaseElem); ok && (be.Value == Byte || be.Value == Uint8) {
		e.p.printf("\nerr = en.WriteBytes(%s[:])", a.Varname())
		e.p.print(errcheck)
		return
	}

	e.writeAndCheck(arrayHeader, literalFmt, a.SizeResolved)
	e.p.rangeBlock(a.Index, a.Varname(), e, a.Els)
}

func (e *encodeGen) gBase(b *BaseElem) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := b.Varname()
	if b.Convert {
		vname = tobaseConvert(b)
	}

	if b.Value == IDENT { // unknown identity
		e.p.printf("\nerr = %s.EncodeMsg(en)", vname)
		e.p.print(errcheck)
	} else { // typical case
		e.writeAndCheck(b.BaseName(), literalFmt, vname)
	}
}
