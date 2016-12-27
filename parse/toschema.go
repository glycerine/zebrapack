package parse

import (
	"fmt"

	"github.com/glycerine/zebrapack/gen"
	"github.com/glycerine/zebrapack/zebra"
	"github.com/shurcooL/go-goon"
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
				fmt.Printf("\n on f = %#v\n", f)
				goon.Dump(f)
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
					zt := f.FieldElem.GetZtype()
					fld.FieldFullType = &zt
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
			//fmt.Fprintf(os.Stderr, "\nwarning: ignoring type '%v'; we only support top-level structs in a schema at present.\n", x.TypeName())
			//return nil, fmt.Errorf("unhandled type %T", x)
		}
	}
	imports := []string{}
	for _, imp := range fs.Imports {
		n := ""
		if imp.Name != nil {
			n = imp.Name.Name // local package name (including "."); or nil
		}
		p := imp.Path.Value // import path
		if len(n) > 0 {
			imports = append(imports, fmt.Sprintf("%s %s", n, p))
		} else {
			imports = append(imports, p)
		}
	}

	sch := &zebra.Schema{
		SourcePath:    path,
		SourcePackage: fs.Package,
		ZebraSchemaId: fs.ZebraSchemaId,
		Structs:       structs,
		Imports:       imports,
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
		zc = zebra.PointerCat
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
