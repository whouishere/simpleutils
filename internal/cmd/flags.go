package cmd

import (
	"fmt"
	"os"

	"codeberg.org/whou/simpleutils/coreutils"
	flag "github.com/erikjuhani/miniflag"
)

var Binary string
var versionFlag *Flag[bool]

type Flag[T any] struct {
	Value        *T
	defaultValue T
	name         string
	shorthand    string
	usage        string
	Function     func()
}

func Init(binary, usage string, format ...any) {
	// modify default usage text
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, format...)
		flag.CommandLine.PrintDefaults()
	}

	Binary = binary

	// universal version flag
	versionFlag = NewFlag(false, "version", "version", "print version info and exit", func() {
		fmt.Printf("%s version %s\n", binary, coreutils.Version)
		os.Exit(0)
	})
	RegisterFlag(versionFlag)
}

func Parse() {
	flag.Parse()

	if *versionFlag.Value {
		versionFlag.Function()
	}
}

func NewFlag[T any](value T, name, shorthand, usage string, function func()) *Flag[T] {
	return &Flag[T]{
		defaultValue: value,
		name:         name,
		shorthand:    shorthand,
		usage:        usage,
		Function:     function,
	}
}

// registers every passed command-line flag
func RegisterFlag[T any](flagVar *Flag[T]) {
	flagVar.Value = flag.Flag(flagVar.name, flagVar.shorthand, flagVar.defaultValue, flagVar.usage)
}
