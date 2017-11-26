package gen

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/glycerine/zebrapack/cfg"
)

func unmarshal(w io.Writer, cfg *cfg.ZebraConfig) *unmarshalGen {
	return &unmarshalGen{
		p:   printer{w: w},
		cfg: cfg,
	}
}

type unmarshalGen struct {
	passes
	p        printer
	hasfield bool
	depth    int
	cfg      *cfg.ZebraConfig
	post     postDefs
}

func (u *unmarshalGen) MethodPrefix() string {
	return u.cfg.MethodPrefix
}

func (d *unmarshalGen) postLines() {
	lines := strings.Join(d.post.endlines, "\n")
	d.p.printf("\n%s\n", lines)
	d.post.reset()
}

func (u *unmarshalGen) Method() Method { return Unmarshal }

func (u *unmarshalGen) needsField() {
	if u.hasfield {
		return
	}
	u.p.print("\nvar field []byte; _ = field")
	u.hasfield = true
}

func (u *unmarshalGen) Execute(p Elem) error {
	u.hasfield = false
	if !u.p.ok() {
		return u.p.err
	}
	if !IsPrintable(p) {
		return nil
	}

	u.p.comment(fmt.Sprintf("%sUnmarshalMsg implements msgp.Unmarshaler", u.cfg.MethodPrefix))

	vname := p.Varname()
	methRcvr := methodReceiver(p)
	if u.cfg.ReadStringsFast {
		u.p.printf("\nfunc (%s %s) %sUnmarshalMsg(bts []byte) (o []byte, err error) {\n cfg := &msgp.RuntimeConfig{UnsafeZeroCopy:true}; return %s.%sUnmarshalMsgWithCfg(bts, cfg)\n}", vname, methRcvr, u.cfg.MethodPrefix, vname, u.cfg.MethodPrefix)
	} else {
		u.p.printf("\nfunc (%s %s) %sUnmarshalMsg(bts []byte) (o []byte, err error) {\n return %s.%sUnmarshalMsgWithCfg(bts, nil)\n}", vname, methRcvr, u.cfg.MethodPrefix, vname, u.cfg.MethodPrefix)
	}
	u.p.printf("\nfunc (%s %s) %sUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {", vname, methRcvr, u.cfg.MethodPrefix)
	// u.p.printf("\nvar nbs msgp.NilBitsStack;\nvar sawTopNil bool\n if msgp.IsNil(bts) {\n 	sawTopNil = true\n fmt.Printf(\"len of bts pre push: %%v\\n\", len(bts));	bts = nbs.PushAlwaysNil(bts[1:]);\n	fmt.Printf(\"len of bts post push: %%v\\n\", len(bts));\n   }\n")
	u.p.printf("\nvar nbs msgp.NilBitsStack;\nnbs.Init(cfg)\nvar sawTopNil bool\n if msgp.IsNil(bts) {\n 	sawTopNil = true\n  bts = nbs.PushAlwaysNil(bts[1:]);\n	}\n")
	next(u, p)
	u.p.print("\n	if sawTopNil {bts = nbs.PopAlwaysNil()}\n o = bts")

	u.p.postLoadHook()
	u.p.nakedReturn()
	unsetReceiver(p)
	u.postLines()
	return u.p.err
}

// does assignment to the variable "name" with the type "base"
func (u *unmarshalGen) assignAndCheck(name string, base string) {
	if !u.p.ok() {
		return
	}
	u.p.printf("\n%s, bts, err = nbs.Read%sBytes(bts)", name, base)
	u.p.print(errcheck)
}

func (u *unmarshalGen) gStruct(s *Struct) {
	u.depth++
	defer func() {
		u.depth--
	}()

	if !u.p.ok() {
		return
	}
	if s.AsTuple {
		u.tuple(s)
	} else {
		u.mapstruct(s)
	}
	return
}

func (u *unmarshalGen) tuple(s *Struct) {

	// open block
	sz := randIdent()
	u.p.declare(sz, u32)
	u.assignAndCheck(sz, arrayHeader)
	u.p.arrayCheck(strconv.Itoa(len(s.Fields)-s.SkipCount), sz, "")
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}

		if !u.p.ok() {
			return
		}
		next(u, s.Fields[i].FieldElem)
	}
}

func (u *unmarshalGen) mapstruct(s *Struct) {
	n := len(s.Fields) // - s.SkipCount
	if n == 0 {
		return
	}
	u.needsField()
	k := genSerial()
	fast := !u.cfg.UseMsgp2
	tmpl, nStr := genUnmarshalMsgTemplate(k, fast)

	fieldOrder := fmt.Sprintf("\n var unmarshalMsgFieldOrder%s = []string{", nStr)
	fieldSkip := fmt.Sprintf("\n var unmarshalMsgFieldSkip%s = []bool{", nStr)
	for i := range s.Fields {
		if s.Fields[i].Skip {
			fieldSkip += fmt.Sprintf("true,")
		} else {
			fieldSkip += fmt.Sprintf("false,")
		}
		fieldOrder += fmt.Sprintf("%q,", s.Fields[i].FieldTag)
	}
	fieldOrder += "}\n"
	fieldSkip += "}\n"

	varname := strings.Replace(s.TypeName(), "\n", ";", -1)
	u.post.add(varname, "\n// fields of %s%s%s", varname, fieldOrder, fieldSkip)

	u.p.printf("\n const maxFields%s = %d\n", nStr, n)

	found := "found" + nStr
	u.p.printf(tmpl)

	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if fast {
			u.p.printf("\ncase %v:", s.Fields[i].ZebraId)
			u.p.printf("\n// zid %v for %q", s.Fields[i].ZebraId,
				s.Fields[i].FieldTag)
			u.p.printf("\n%s[%d]=true;", found, i)
		} else {
			u.p.printf("\ncase \"%s\":", s.Fields[i].FieldTag)
			u.p.printf("\n%s[%d]=true;", found, i)
		}
		u.depth++
		next(u, s.Fields[i].FieldElem)
		u.depth--
		if !u.p.ok() {
			return
		}
	}
	u.p.print("\ndefault:\nbts, err = msgp.Skip(bts)")
	u.p.print(errcheck)
	u.p.print("\n}\n}") // close switch and for loop

	u.p.printf("\n if nextMiss%s != -1 { bts = nbs.PopAlwaysNil(); }\n", nStr)
}

func (u *unmarshalGen) gBase(b *BaseElem) {
	if !u.p.ok() {
		return
	}

	refname := b.Varname() // assigned to
	lowered := b.Varname() // passed as argument
	if b.Convert {
		// begin 'tmp' block
		refname = randIdent()
		lowered = b.ToBase() + "(" + lowered + ")"
		u.p.printf("\n{\nvar %s %s", refname, b.BaseType())
	}

	switch b.Value {
	case Bytes:
		u.p.printf("\n if nbs.AlwaysNil || msgp.IsNil(bts) {\n if !nbs.AlwaysNil { bts = bts[1:]  }\n  %s = %s[:0]} else { %s, bts, err = nbs.ReadBytesBytes(bts, %s)\n", refname, refname, refname, lowered)
		u.p.print(errcheck)
		u.p.closeblock()
	case Ext:
		vn := b.Varname()[1:]
		u.p.printf("\n if nbs.AlwaysNil || msgp.IsNil(bts) { if !nbs.AlwaysNil { bts = bts[1:] }\n    %s = msgp.RawExtension{}  \n} else {\n bts, err = nbs.ReadExtensionBytes(bts, %s) \n", vn, lowered)
		u.p.print(errcheck)
		u.p.closeblock()
	case IDENT:
		u.p.printf("\n  bts, err = %s.%sUnmarshalMsg(bts);", lowered, u.cfg.MethodPrefix)
		u.p.print(errcheck)
	default:
		//		u.p.printf("\n if nbs.AlwaysNil || msgp.IsNil(bts) { if !nbs.AlwaysNil { bts=bts[1:]}\n   %s \n} else {  %s, bts, err = nbs.Read%sBytes(bts)\n", b.ZeroLiteral(refname), refname, b.BaseName())
		u.p.printf("\n %s, bts, err = nbs.Read%sBytes(bts)\n", refname, b.BaseName())
	}
	u.p.print(errcheck)
	if b.Convert {
		// close 'tmp' block
		u.p.printf("\n%s = %s(%s)\n}", b.Varname(), b.FromBase(), refname)
	}
}

func (u *unmarshalGen) gArray(a *Array) {
	if !u.p.ok() {
		return
	}

	// special case for [const]byte objects
	// see decode.go for symmetry
	if be, ok := a.Els.(*BaseElem); ok && be.Value == Byte {
		u.p.printf("\nbts, err = nbs.ReadExactBytes(bts, %s[:])", a.Varname())
		u.p.print(errcheck)
		return
	}

	sz := randIdent()
	u.p.declare(sz, u32)
	u.assignAndCheck(sz, arrayHeader)
	u.p.arrayCheck(a.SizeResolved, sz, "!nbs.IsNil(bts) && ")
	u.p.unmarshalRangeBlock(a.Index, a.Varname(), u, a.Els)
}

func (u *unmarshalGen) gSlice(s *Slice) {
	if !u.p.ok() {
		return
	}
	u.p.printf("\n if nbs.AlwaysNil { %s \n} else {\n",
		s.ZeroLiteral(`(`+s.Varname()+`)`))
	sz := randIdent()
	u.p.declare(sz, u32)
	u.assignAndCheck(sz, arrayHeader)
	u.p.resizeSlice(sz, s)
	u.p.rangeBlock(s.Index, s.Varname(), u, s.Els)
	u.p.closeblock()
}

func (u *unmarshalGen) gMap(m *Map) {
	u.depth++
	defer func() {
		u.depth--
	}()

	if !u.p.ok() {
		return
	}
	u.p.printf("\n if nbs.AlwaysNil { %s \n} else {\n",
		m.ZeroLiteral(m.Varname()))
	sz := randIdent()
	u.p.declare(sz, u32)
	u.assignAndCheck(sz, mapHeader)

	// allocate or clear map
	u.p.resizeMap(sz, m)

	// loop and get key,value
	u.p.printf("\nfor %s > 0 {", sz)
	u.p.printf("\nvar %s %s; var %s %s; %s--", m.Keyidx, m.KeyDeclTyp, m.Validx, m.Value.TypeName(), sz)
	u.assignAndCheck(m.Keyidx, m.KeyTyp)
	next(u, m.Value)
	u.p.mapAssign(m)
	u.p.closeblock()
	u.p.closeblock()
}

func (u *unmarshalGen) gPtr(p *Ptr) {
	vname := p.Varname()

	base, isBase := p.Value.(*BaseElem)
	if isBase {
		//u.p.printf("\n // we have a BaseElem: %#v  \n", base)
		switch base.Value {
		case IDENT:
			//u.p.printf("\n // we have an IDENT: \n")

			u.p.printf("\n if nbs.AlwaysNil { ")
			u.p.printf("\n if %s != nil { \n", vname)

			niller := fmt.Sprintf("; %s.%sUnmarshalMsg(msgp.OnlyNilSlice);", vname, u.cfg.MethodPrefix)

			u.p.printf("%s\n}\n } else { \n // not nbs.AlwaysNil \n", niller)
			u.p.printf("if msgp.IsNil(bts) { bts = bts[1:]; if nil != %s { \n %s}", vname, niller)
			u.p.printf("} else { \n // not nbs.AlwaysNil and not IsNil(bts): have something to read \n")
			u.p.initPtr(p)
			next(u, p.Value)
			u.p.closeblock()
			u.p.closeblock()
			return
		case Ext:
			//u.p.printf("\n // we have an base.Value of Ext: \n")

			u.p.printf("\n if (nbs.AlwaysNil || msgp.IsNil(bts)) { \n // don't try to re-use extension pointers\n if !nbs.AlwaysNil { bts=bts[1:]  }\n %s = nil } else {\n // we have data \n", vname)
			u.p.initPtr(p)
			u.p.printf("\n  bts, err = nbs.ReadExtensionBytes(bts, %s)\n", vname)
			u.p.print(errcheck)
			u.p.closeblock()
			return
		}
	}

	u.p.printf("\n // default gPtr logic.")
	u.p.printf("\nif nbs.PeekNil(bts) && %s == nil { \n // consume the nil\n bts, err = nbs.ReadNilBytes(bts) \n if err != nil { return }  } else { \n // read as-if the wire has bytes, letting nbs take care of nils. \n", vname)
	u.p.initPtr(p)
	next(u, p.Value)
	u.p.closeblock()
}
