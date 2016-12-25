package gen

import (
	"fmt"

	"github.com/glycerine/zebrapack/zebra"
)

func TranslateToZebraSchema(elem Elem) (*zebra.Schema, error) {
	switch x := elem.(type) {
	case *Struct:
		n := len(x.Fields)
		tr := zebra.Struct{
			StructName: "",
			Fld:        make([]zebra.Field, n),
		}
		for _, f := range x.Fields {
			fld := zebra.Field{
				Zid:       f.ZebraId,
				Nam:       f.FieldElem.Varname(),
				TypStr:    f.FieldElem.TypeName(),
				OmitEmpty: f.OmitEmpty,
				Skip:      f.Skip,
			}
			tr.Fld = append(tr.Fld, fld)
		}
		sch := &zebra.Schema{
			Structs: []zebra.Struct{tr},
		}
		return sch, nil
		/*
			case *Ptr:
			case *Array:
			case *Slice:
			case *Map:
			case *BaseElem:
		*/
	default:
		return nil, fmt.Errorf("unhandled type %T", x)
	}
}
