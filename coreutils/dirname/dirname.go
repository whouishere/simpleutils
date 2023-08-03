package main

import (
	"fmt"
	"os"
	"path/filepath"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "dirname"
var usage = `Usage: %s [OPTION] NAME...
%s strips non-directory suffixes from NAME(s).

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)
	cmd.Parse()
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument")
	}

	for _, arg := range args {
		fmt.Println(filepath.ToSlash(filepath.Dir(arg)))
	}

	os.Exit(0)
}
