package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestDiffB(t *testing.T) {

	cv.Convey("Given two different (possibly long, differing in whitespace) strings a and b", t, func() {
		cv.Convey("then DiffB(a, b) should return the output of diff -b to consisely summarize their differences", func() {

			a := `
1
2
3
4
5
type   s1    struct {
  MyInts    []int
}`
			b := `
1
2
3
4
type s1 struct {
  MyInts []int
}`
			cv.So(string(Diffb(a, b)), cv.ShouldEqual, `6d5
< 5
`)

		})
	})
}
