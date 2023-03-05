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

func runFlags() {
	cmd.Init(binary, usage, binary, binary)
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

	cwd = filepath.ToSlash(cwd)

	fmt.Println(cwd)

	os.Exit(0)
}
