package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

type TempDir struct {
	OrigDir string
	DirPath string
	Files   map[string]*os.File
}

func NewTempDir() *TempDir {
	dirname, err := ioutil.TempDir(".", "testdir_")
	if err != nil {
		panic(err)
	}
	origdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// add files needed for capnpc -ogo compilation
	exec.Command("/bin/cp", "go.capnp", dirname).Run()

	return &TempDir{
		OrigDir: origdir,
		DirPath: dirname,
		Files:   make(map[string]*os.File),
	}
}

func (d *TempDir) MoveTo() {
	err := os.Chdir(d.DirPath)
	if err != nil {
		panic(err)
	}
}

func (d *TempDir) Close() {
	for _, f := range d.Files {
		f.Close()
	}
}

func (d *TempDir) Cleanup() {
	d.Close()
	err := os.RemoveAll(d.DirPath)
	if err != nil {
		panic(err)
	}
	err = os.Chdir(d.OrigDir)
	if err != nil {
		panic(err)
	}
}

func (d *TempDir) TempFile() *os.File {

	f, err := ioutil.TempFile(d.DirPath, "testfile.")
	if err != nil {
		panic(err)
	}
	d.Files[f.Name()] = f
	return f
}
