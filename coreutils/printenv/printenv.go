package main

import (
	"fmt"
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "printenv"
var usage = `Usage: %s [OPTION] [VARIABLE]...
%s prints the value of the specified environment VARIABLE(s). If none are given, it prints them all.

`

var nullFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	nullFlag = cmd.NewFlag(false,
		"null", "0",
		"end each output line with a NUL byte rather than a newline")

	cmd.Parse()
}

func printAllEnv() {
	variables := os.Environ()

	for _, env := range variables {
		if *nullFlag {
			fmt.Print(env)
		} else {
			fmt.Println(env)
		}
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()

	if args == nil {
		printAllEnv()
		os.Exit(0)
	}

	for _, arg := range args {
		if *nullFlag {
			fmt.Print(os.Getenv(arg))
		} else {
			fmt.Println(os.Getenv(arg))
		}
	}

	os.Exit(0)
}
