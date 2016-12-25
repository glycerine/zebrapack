package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tinylib/msgp/msgp"
)

func (fs *FileSet) SaveMsgpackFile(parsedPath, path string) error {
	//fmt.Printf("\n saveMsgpackFile saving to: '%v'", path)

	var f *os.File
	var err error
	if path == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	w := msgp.NewWriter(f)
	defer w.Flush()

	tr, err := TranslateToZebraSchema(parsedPath, fs)
	if err != nil {
		return err
	}
	err = tr.EncodeMsg(w)
	if err != nil {
		return err
	}
	if path != "-" {
		by, err := tr.MarshalMsg(nil)
		if err != nil {
			return err
		}
		fjson, err := os.Create(path + ".json")
		if err != nil {
			return err
		}
		defer fjson.Close()

		// and write out pretty json version too.
		buf := bytes.NewBuffer(by)
		var compactJson, pretty bytes.Buffer
		_, err = msgp.CopyToJSON(&compactJson, buf)
		if err != nil {
			return err
		}

		err = json.Indent(&pretty, compactJson.Bytes(), "", "    ")
		if err != nil {
			return err
		}
		_, err = io.Copy(fjson, &pretty)
		fmt.Fprintf(fjson, "\n")
		return err
	}
	return nil
}
