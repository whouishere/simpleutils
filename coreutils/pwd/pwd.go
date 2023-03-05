package main

import (
	"fmt"
	"os"
	"path/filepath"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "pwd"
var usage = `Usage: %s [OPTION(s)]...
%s prints the full name of the working directory.

`

var logicalFlag *cmd.Flag[bool]
var physicalFlag *cmd.Flag[bool]

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	logicalFlag = cmd.NewFlag(false,
		"logical", "L",
		"output absolute working directory, keeping symbolic links",
		nil)
	cmd.RegisterFlag(logicalFlag)

	physicalFlag = cmd.NewFlag(false,
		"physical", "P",
		"output the physical working path, with resolved symbolic links",
		nil)
	cmd.RegisterFlag(physicalFlag)

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

	if *logicalFlag.Value {
		cwd, err = filepath.Abs(cwd)
		if err != nil {
			panic(err)
		}
	}

	if *physicalFlag.Value {
		cwd, err = filepath.EvalSymlinks(cwd)
		if err != nil {
			panic(err)
		}
	}

	cwd = filepath.ToSlash(cwd)

	fmt.Println(cwd)

	os.Exit(0)
}
