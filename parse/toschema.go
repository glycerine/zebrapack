package parse

import (
	"fmt"

	"github.com/glycerine/zebrapack/gen"
	"github.com/glycerine/zebrapack/zebra"
)

func TranslateToZebraSchema(path string, fs *FileSet) (*zebra.Schema, error) {

	structs := []zebra.Struct{}

	for _, ele := range fs.Identities {

		switch x := ele.(type) {
		case *gen.Struct:
			tr := zebra.Struct{
				StructName: x.TypeName(),
			}
			for _, f := range x.Fields {
				zc, zp := getCatPrimiType(&f)
				//fmt.Printf("\n on f = %#v\n", f)
				fld := zebra.Field{
					Zid:            f.ZebraId,
					OmitEmpty:      f.OmitEmpty,
					Skip:           f.Skip,
					FieldGoName:    f.FieldName,
					FieldTagName:   f.FieldTag,
					FieldCategory:  zc,
					FieldPrimitive: zp,
				}
				if !fld.Skip {
					fld.FieldTypeStr = f.FieldElem.TypeName()
				}
				//fmt.Printf("\n in %v,  on field %#v ... fld='%#v'\n", tr.StructName, f, fld)
				tr.Fields = append(tr.Fields, fld)
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
		SourcePath:    path,
		ZebraSchemaId: fs.ZebraSchemaId,
		Structs:       structs,
	}
	//fmt.Printf("total number of structs: %v\n", len(structs))
	//fmt.Printf("total number of fields in first struct: %v\n", len(structs[0].Fields))
	return sch, nil

}

func getCatPrimiType(f *gen.StructField) (zc zebra.Zcat, zp zebra.Zprimitive) {

	return
}
