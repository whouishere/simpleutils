package main

import (
	"os"
	"time"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "touch"
var usage = `Usage: %s [OPTION]... FILE...
%s updates the access and modification times of each FILE to the current time.

If a FILE doesn't exist, it is created empty.

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	cmd.Parse()
}

func touchFiles(time time.Time, files []*os.File) {
	for _, file := range files {
		err := os.Chtimes(file.Name(), time, time)
		if err != nil {
			panic(err)
		}
	}
}

// get existent and non-existent files
func getFiles() []*os.File {
	var files []*os.File

	paths := cmd.GetNonFlags()
	for _, path := range paths {
		exists, err := myio.FileExists(path)
		if err != nil {
			panic(err)
		}

		var file *os.File
		if exists {
			file, err = os.Open(path)
			if err != nil {
				panic(err)
			}
		} else {
			file, err = os.Create(path)
			if err != nil {
				panic(err)
			}
		}

		files = append(files, file)
	}

	return files
}

func main() {
	now := time.Now().Local()
	runFlags()

	files := getFiles()
	touchFiles(now, files)

	os.Exit(0)
}
