package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "false"
var usage = `Usage: %s [OPTION] [ignore arguments]
%s exits with a failure status code (1).

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	cmd.IgnoreUndefinedFlags()

	cmd.Parse()
}

func main() {
	runFlags()
	os.Exit(1)
}
