package main

import (
	"fmt"
	"os"
	"path/filepath"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "pwd"
var usage = `Usage: %s [OPTION]...
%s prints the full name of the working directory.

`

var logicalFlag *bool
var physicalFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	logicalFlag = cmd.NewFlag(false,
		"logical", "L",
		"output absolute working directory, keeping symbolic links")

	physicalFlag = cmd.NewFlag(false,
		"physical", "P",
		"output the physical working path, with resolved symbolic links")

	cmd.Parse()
}

func main() {
	runFlags()

	if cmd.GetNonFlags() != nil {
		cmd.Log("Ignoring non-option arguments.")
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if *logicalFlag {
		cwd, err = filepath.Abs(cwd)
		if err != nil {
			panic(err)
		}
	}

	if *physicalFlag {
		cwd, err = filepath.EvalSymlinks(cwd)
		if err != nil {
			panic(err)
		}
	}

	cwd = filepath.ToSlash(cwd)

	fmt.Println(cwd)

	os.Exit(0)
}
