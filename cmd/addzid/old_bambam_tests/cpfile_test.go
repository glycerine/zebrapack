package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestCpCopiesFilesIntoDirHier(t *testing.T) {

	cv.Convey("Cp() function should create directories in destinationPath if need be", t, func() {

		origdir, tmpdir := MakeAndMoveToTempDir()
		defer TempDirCleanup(origdir, tmpdir)
		hier := fmt.Sprintf("a%cb%cc%c", os.PathSeparator, os.PathSeparator, os.PathSeparator)
		goal := hier + "cpfile.go"
		err := Cp(origdir+string(os.PathSeparator)+"cpfile.go", tmpdir+string(os.PathSeparator)+hier)
		cv.So(err, cv.ShouldEqual, nil)
		cv.So(FileExists(goal), cv.ShouldEqual, true)
	})
}

func MakeAndMoveToTempDir() (origdir string, tmpdir string) {

	var err error
	origdir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	tmpdir, err = ioutil.TempDir(origdir, "temptestdir")
	if err != nil {
		panic(err)
	}
	err = os.Chdir(tmpdir)
	if err != nil {
		panic(err)
	}

	return origdir, tmpdir
}

func TempDirCleanup(origdir string, tmpdir string) {
	// cleanup
	os.Chdir(origdir)
	err := os.RemoveAll(tmpdir)
	if err != nil {
		panic(err)
	}
}
