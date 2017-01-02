package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

// Cp implements a portable version of /bin/cp for use in golang
// code that wants to work on many platforms. If the destinationPath
// directory hierarchy does not exist, we will attempt to make it,
// like mkdir -p would.
func Cp(originPath, destinationPath string) (err error) {
	if !FileExists(originPath) {
		return fmt.Errorf("Cp error: source path '%s' does not exist", originPath)
	}

	var dirname, target string

	ostat, err := os.Stat(originPath)
	if err != nil {
		return err
	}
	if !ostat.Mode().IsRegular() {
		return fmt.Errorf("Cp: can only copy regular files, origin is '%s' of mode: '%s'", ostat.Name(), ostat.Mode().String())
	}

	if FileExists(destinationPath) {
		// okay, good to go
		//fmt.Printf("\n okay, destinationPath '%s' exists as a file\n", destinationPath)
		dirname = path.Dir(destinationPath)
		target = path.Base(destinationPath)
	} else {
		// no such file yet
		//fmt.Printf("\n okay, destinationPath '%s': no such file yet.\n", destinationPath)

		// set dirname and target

		// do we end in "/" or "\"?
		if IsDirPath(destinationPath) {
			dirname = destinationPath
			target = path.Base(originPath)
		} else {
			// the final path component is taken to be the target file name, unless
			// it names an existing directory.
			if DirExists(destinationPath) {
				//fmt.Printf("\n okay, destinationPath '%s' exists as a directory\n", destinationPath)
				dirname = path.Dir(destinationPath)
				target = path.Base(originPath)

			} else {
				dirname = path.Dir(destinationPath)
				target = path.Base(destinationPath)
			}
		}

		//fmt.Printf("\n dirname = '%s'   target = '%s'\n", dirname, target)

		// is it a directory that exists?
		if DirExists(dirname) {
			//fmt.Printf("\n okay, dirname '%s' exists as a directory\n", dirname)
		} else {
			//fmt.Printf("\n okay, dirname '%s' is not an existing directory, try to make needed directories\n", dirname)
			// dirname is not an existing directory, try to mkdir -p make the needed dirs.

			err = os.MkdirAll(dirname, 0755)
			if err != nil {
				return fmt.Errorf("could not create destination directory path(s) with MkdirAll on '%s', error: '%s'", dirname, err)
			} else {
				//fmt.Printf("\n made dirname: '%s'\n", dirname)
			}
		}

	}
	// INVAR: we should be good to copy to dirname
	//fmt.Printf("\n INVAR: we should be good to copy to dirname: '%s'\n", dirname)

	fullpath := appendFilename(dirname, target)

	dstat, err := os.Stat(fullpath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		// we have an existing file, don't bother writing atop ourselves.
		if os.SameFile(ostat, dstat) {
			return nil // don't error out here, the copy is actually already done.
		}
	}

	// proceed to copy
	//fmt.Printf("os.Open with originPath = '%s'\n", originPath)
	src, err := os.Open(originPath)
	if err != nil {
		return err
	}
	defer src.Close()

	//fmt.Printf("os.Create with fullpath = '%s'\n", fullpath)
	dest, err := os.Create(fullpath)
	if err != nil {
		//fmt.Printf("returning early after attempt os.Create with fullpath = '%s'\n", fullpath)
		return err
	}
	defer func() {
		closeError := dest.Close()
		// only substitute closeError if there was
		// no previous earlier error
		if err == nil {
			err = closeError
		}
	}()
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	return nil
}

func appendFilename(dir, fn string) string {
	// but don't double // or \\ it
	if IsDirPath(dir) {
		return dir + fn
	}
	return dir + string(os.PathSeparator) + fn
}

func IsDirPath(dir string) bool {
	n := len(dir)
	if n == 0 {
		return false
	}
	if dir[n-1] == os.PathSeparator {
		return true
	}
	return false
}
