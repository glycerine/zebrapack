package parse

import (
	"testing"

	"github.com/glycerine/diff"
	cv "github.com/glycerine/goconvey/convey"
)

func Test100AddZidBasic(t *testing.T) {
	cv.Convey("addzid: addzid should add zid tags to structs with public fields.", t, func() {
		ex := "package a; type A struct { A int }"
		expected := "package a; type A struct { A int `zid:\"0\"`}"
		out, err := addZid(ex)
		cv.So(err, cv.ShouldBeNil)
		cv.So(string(out), diff.ShouldStartWithModuloWhiteSpace, expected)
	})
}

func Test101AddZidExistingSequence(t *testing.T) {
	cv.Convey("addzid: the zid numbering should follow the existing field sequence and be sequential.", t, func() {
	})
}

func Test102AddZidDetectExistingZid(t *testing.T) {
	cv.Convey("addzid: if there is already a zid field on the struct, then we don't add any new zids automatically", t, func() {
	})
}

func Test103AddZidRespectIgnore(t *testing.T) {
	cv.Convey("addzid: if the struct is designated `ignore` then we don't add any zids", t, func() {})
}

func Test104AddZidUnexportedPrivateBothRespected(t *testing.T) {
	cv.Convey("addzid: private fields don't get zids added, unless unexported is added.", t, func() {})
}

func Test105AddZidFieldsThatShouldBeSkippedAreSkipped(t *testing.T) {
	cv.Convey("addzid: fields marked msg:\"-\" or of type struct{} are skipped and don't get zids.", t, func() {})
}

func Test106AddZidPreserveExistingFieldTags(t *testing.T) {
	cv.Convey("addzid: we preserve existing field tags.", t, func() {})
}

func Test107AddZidDoNotChangeOriginalGoSourceFile(t *testing.T) {
	cv.Convey("addzid: write new file to .addzid version of input .go file don't blow away an existing file, however.", t, func() {})
}
