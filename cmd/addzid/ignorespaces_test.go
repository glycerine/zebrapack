package main

import "testing"

func TestShouldStartWithModuloWhiteSpace(t *testing.T) {
	pass(t, so("", ShouldStartWithModuloWhiteSpace, ""))
	pass(t, so("a", ShouldStartWithModuloWhiteSpace, "a"))
	pass(t, so(" a ", ShouldStartWithModuloWhiteSpace, "a"))
	pass(t, so("   a   ", ShouldStartWithModuloWhiteSpace, "a"))
	pass(t, so("a", ShouldStartWithModuloWhiteSpace, " a "))
	pass(t, so("a", ShouldStartWithModuloWhiteSpace, "   a   "))
	pass(t, so("a b c de fgh   ij", ShouldStartWithModuloWhiteSpace, "  abcdefghi    j "))

	fail(t, so("j  ", ShouldMatchModuloSpaces, "j k"), "Expected expected string 'j k' and actual string 'j  ' to match (ignoring {' '}) (but they did not!; first diff at '', pos 2); and Full diff -b: 1c1 < j --- > j k ")

	fail(t, so("asdf", ShouldStartWithModuloWhiteSpace), "This assertion requires exactly 1 comparison values (you provided 0).")
	fail(t, so("asdf", ShouldStartWithModuloWhiteSpace, 1, 2, 3), "This assertion requires exactly 1 comparison values (you provided 3).")

	fail(t, so(123, ShouldStartWithModuloWhiteSpace, 23), "Both arguments to this assertion must be strings (you provided int and int).")

}
