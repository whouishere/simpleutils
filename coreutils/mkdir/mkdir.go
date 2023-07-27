package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "mkdir"
var usage = `Usage: %s [OPTION]... DIRECTORY...
%s creates DIRECTORY(ies), if they do not already exist.

`

var parentsFlag *bool
var verboseFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	parentsFlag = cmd.NewFlag(false,
		"parents", "p",
		"no error if DIRECTORY exists and create DIRECTORY and its parents")

	verboseFlag = cmd.NewFlag(false,
		"verbose", "v",
		"output verbose messages for every created directory")

	cmd.Parse()
}

func makeDirs(dir string) {
	cmd.SetErrorPrefix("Failed to create '", dir, "'.")

	exist, err := myio.FileExists(dir)
	if err != nil {
		panic(err)
	}
	if exist && !*parentsFlag {
		cmd.FatalError("File/directory already exists.")
	}

	// the verbose print is after the "file exists" check on GNU
	if *verboseFlag {
		cmd.Log("Creating directory '", dir, "'.")
	}

	if *parentsFlag {
		err = os.MkdirAll(dir, 0755)

		if *verboseFlag {
			runedir := []rune(dir)
			// remove trailing slash
			if runedir[len(runedir)-1] == '/' {
				runedir = runedir[:len(runedir)-1]
			}

			var parents []string
			for i := range runedir {
				if runedir[i] == '/' {
					parents = append(parents, string(runedir[:i]))
				}
			}

			for i := len(parents) - 1; i >= 0; i-- { // reverse iterate
				cmd.Log("Creating directory '", parents[i], "'.")
			}
		}
	} else {
		err = os.Mkdir(dir, 0755)
	}

	if err != nil {
		panic(err)
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument.")
	}

	for _, arg := range args {
		makeDirs(arg)
	}

	os.Exit(0)
}
