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
	IgnoreUndefinedFlags()
	os.Args = append(os.Args, args...)

	var _ = NewFlag(false, "", "s", "shorthand-only flag")
	var _ = NewFlag(false, "long", "", "longform-only flag")
	var _ = NewFlag(false, "both", "b", "long flag with a shorthand")

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

  -b, --both      long flag with a shorthand
      --help      display this help text and exit
      --long      longform-only flag
  -s              shorthand-only flag
      --version   print version info and exit
`

	if out != outUsage {
		t.Fatalf("the flag output is not an exact match.\noutput: '%s'", out)
	}
}
