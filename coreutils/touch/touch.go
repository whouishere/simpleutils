package main

import (
	"os"
	"time"

	"github.com/djherbis/atime"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "touch"
var usage = `Usage: %s [OPTION]... FILE...
%s updates the access and modification times of each FILE to the current time.

If a FILE doesn't exist, it is created empty.

`

var noCreateFlag *bool
var accessFlag *bool
var modificationFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	noCreateFlag = cmd.NewFlag(false,
		"no-create", "c",
		"do not create any files")

	accessFlag = cmd.NewFlag(false,
		"", "a",
		"change only the access time")

	modificationFlag = cmd.NewFlag(false,
		"", "m",
		"change only the modification time")

	cmd.Parse()
}

func touchFiles(changeTime time.Time, files []*os.File) {
	var aTime = changeTime
	var mTime = changeTime

	for _, file := range files {
		// if either -a or -m were used, keep the time of the one that wasn't selected
		if *accessFlag || *modificationFlag {
			info, err := file.Stat()
			if err != nil {
				panic(err)
			}

			if !*accessFlag {
				aTime = atime.Get(info)
			}

			if !*modificationFlag {
				mTime = info.ModTime()
			}
		}

		err := os.Chtimes(file.Name(), aTime, mTime)
		if err != nil {
			panic(err)
		}
	}
}

// get existent and non-existent files
func getFiles(paths []string) []*os.File {
	var files []*os.File

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
		} else if !*noCreateFlag {
			file, err = os.Create(path)
			if err != nil {
				panic(err)
			}
		} else {
			continue
		}

		files = append(files, file)
	}

	return files
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument")
	}

	now := time.Now().Local()
	files := getFiles(args)
	touchFiles(now, files)

	os.Exit(0)
}
