package io

import (
	"os"
)

// Returns true if a file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}

		return false, err
	}

	return true, nil
}

// Returns true if a given path is a directory
func FileIsDir(path string) (bool, error) {
	file, err := os.Stat(path)
	return file.IsDir(), err
}

// Recursively checks if a directory does not have any files.
func IsDirEmpty(path string) (bool, error) {
	dir, err := os.ReadDir(path)

	// if dir is not empty
	if len(dir) != 0 {
		// check if there are anything other than folders
		for _, dirfile := range dir {
			// if any file returns false to IsDir(), return false
			if !dirfile.IsDir() {
				return false, err
			}

			// check if inner dir has something
			innerdirisempty, err := IsDirEmpty(dirfile.Name())
			if !innerdirisempty {
				return false, err
			}
		}

		// return true when looped through empty dirs
		return true, err
	}

	return true, err
}
