package main

var capnpKeywords map[string]bool = map[string]bool{
	"Void": true, "Bool": true, "Int8": true, "Int16": true, "Int32": true, "Int64": true, "UInt8": true, "UInt16": true, "UInt32": true, "UInt64": true, "Float32": true, "Float64": true, "Text": true, "Data": true, "List": true, "struct": true, "union": true, "group": true, "enum": true, "AnyPointer": true, "interface": true, "extends": true, "const": true, "using": true, "import": true, "annotation": true}

func isCapnpKeyword(w string) bool {
	return capnpKeywords[w] // not found will return false, the zero value for bool.
}
