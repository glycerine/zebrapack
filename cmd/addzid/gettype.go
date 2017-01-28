package main

import "go/ast"

// recursively extract the go type as a string
func GetTypeAsString(ty ast.Expr, sofar string, goTypeSeq []string) (string, string, []string) {
	//p("debug: in GetTypeAsSTring, ty = %#v", ty)
	switch x := ty.(type) {

	case (*ast.StarExpr):
		return GetTypeAsString(ty.(*ast.StarExpr).X, sofar+"*", append(goTypeSeq, "*"))

	case (*ast.Ident):
		return sofar, ty.(*ast.Ident).Name, append(goTypeSeq, ty.(*ast.Ident).Name)

	case (*ast.ArrayType):
		// slice or array
		return GetTypeAsString(ty.(*ast.ArrayType).Elt, sofar+"[]", append(goTypeSeq, "[]"))
	case (*ast.SelectorExpr):
		//p("debug: SelectorExpr case! x.X=%#v, x.Sel=%#v", x.X, x.Sel)
		if ident, ok := x.X.(*ast.Ident); ok {
			//p("debug: constructing '%s'", ident.Name+"."+x.Sel.Name)
			return sofar, ident.Name + "." + x.Sel.Name, goTypeSeq
		}
	}

	return sofar, "", goTypeSeq
}
