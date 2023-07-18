package io

import (
	"os"
	"strings"
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

// Returns the basename of a file from a filepath string
func GetFilepathBasename(filepath string) string {
	paths := strings.Split(filepath, "/")
	return paths[len(paths)-1]
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
