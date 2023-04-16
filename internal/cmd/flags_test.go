package cmd

import (
	"os"
	"testing"

	"codeberg.org/whou/simpleutils/internal/test"
)

var args []string

func initFakeFlags() {
	var binary = "flagtest"
	var usage = `Usage: %s [OPTION]... [FILE]...
This is a test for command-line flags functionality.

`
	Init(binary, usage, binary)
	os.Args = append(os.Args, args...)
	Parse()
}

func TestVersionFlag(t *testing.T) {
	args = []string{"--version"}
	out, err := test.GetOutput(initFakeFlags, t.Name())
	if err != nil {
		t.Fatalf("output: %s\nerr: %s", out, err)
	}

	if out != "flagtest version dev\n" {
		t.Fatalf("the flag output is not an exact match.\noutput: '%s'", out)
	}
}

func TestHelpFlag(t *testing.T) {
	args = []string{"--help"}
	out, err := test.GetOutput(initFakeFlags, t.Name())
	if err != nil {
		t.Fatalf("output: %s\nerr: %s", out, err)
	}

	outUsage := `Usage: flagtest [OPTION]... [FILE]...
This is a test for command-line flags functionality.

      --help      display this help text and exit
      --version   print version info and exit
`

	if out != outUsage {
		t.Fatalf("the flag output is not an exact match.\noutput: '%s'", out)
	}
}
