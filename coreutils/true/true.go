package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "true"
var usage = `Usage: %s [OPTION] [ignore arguments]
%s exits with a success status code (0).

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	cmd.IgnoreUndefinedFlags()

	cmd.Parse()
}

func main() {
	runFlags()
	os.Exit(0)
}
