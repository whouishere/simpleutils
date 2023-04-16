package cmd

import (
	"fmt"
	"os"

	"codeberg.org/whou/simpleutils/coreutils"
	flag "github.com/cornfeedhobo/pflag"
)

var versionFlag *bool
var helpFlag *bool

func Init(binary, usage string, format ...any) {
	// modify default usage text
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, format...)
		flag.PrintDefaults()
	}

	coreutils.Binary = binary

	// universal flags
	versionFlag = NewFlag(false, "version", "", "print version info and exit")
	helpFlag = NewFlag(false, "help", "", "display this help text and exit")
}

func Parse() {
	// FIXME?: internally error is printed twice, which is weird
	// this bug might be unfixable on the user side (check library repo)
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s version %s\n", coreutils.Binary, coreutils.Version)
		os.Exit(0)
	}

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}
}

// whitelist unknown flags
func IgnoreUndefinedFlags() {
	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
}

// return a new boolean pointer
func NewFlag(value bool, name, shorthand, usage string) *bool {
	if name == "" {
		return flag.BoolS(shorthand, shorthand, value, usage)
	}

	return flag.BoolP(name, shorthand, value, usage)
}
