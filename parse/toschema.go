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
					Deprecated:     f.Deprecated,
				}
				if !fld.Skip {
					fld.FieldFullType = f.FieldElem.GetZtype()
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
		SourcePackage: fs.Package,
		ZebraSchemaId: fs.ZebraSchemaId,
		Structs:       structs,
	}
	//fmt.Printf("total number of structs: %v\n", len(structs))
	//fmt.Printf("total number of fields in first struct: %v\n", len(structs[0].Fields))
	return sch, nil

}

func getCatPrimiType(f *gen.StructField) (zc zebra.Zkind, zp zebra.Zkind) {
	switch e := f.FieldElem.(type) {
	case *gen.Map:
		zc = zebra.MapCat
	case *gen.Struct:
		zc = zebra.StructCat
	case *gen.Slice:
		zc = zebra.SliceCat
	case *gen.Array:
		zc = zebra.ArrayCat
	case *gen.Ptr:
		zc = zebra.PtrCat
	case *gen.BaseElem:
		zc = zebra.BaseElemCat
		zp = zebra.Zkind(e.Value)
	case nil:
		// struct{} or other skippable, default 0, 0 is fine.
	default:
		panic(fmt.Errorf("bad element type %T", e))
	}

	return
}
