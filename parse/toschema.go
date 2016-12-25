package parse

import (
	"fmt"

	"github.com/glycerine/zebrapack/gen"
	"github.com/glycerine/zebrapack/zebra"
)

func TranslateToZebraSchema(fs *FileSet) (*zebra.Schema, error) {

	structs := []zebra.Struct{}

	for _, ele := range fs.Identities {

		switch x := ele.(type) {
		case *gen.Struct:
			n := len(x.Fields)
			tr := zebra.Struct{
				StructName: "",
				Fld:        make([]zebra.Field, n),
			}
			for _, f := range x.Fields {
				fmt.Printf("\n on field %#v\n", f)
				fld := zebra.Field{
					Zid:       f.ZebraId,
					OmitEmpty: f.OmitEmpty,
					Skip:      f.Skip,
				}
				if !fld.Skip {
					fld.Nam = f.FieldElem.Varname()
					fld.TypStr = f.FieldElem.TypeName()
				}
				tr.Fld = append(tr.Fld, fld)
			}
			structs = append(structs, tr)
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
	sch := &zebra.Schema{
		Structs: structs,
	}
	return sch, nil

}
