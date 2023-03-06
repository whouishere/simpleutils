package cmd

import (
	"fmt"
	"os"
	"strings"

	"codeberg.org/whou/simpleutils/coreutils"
	flag "github.com/erikjuhani/miniflag"
)

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

	coreutils.Binary = binary

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

// ignore undefined flags error by removing them from the argument list
func IgnoreUndefinedFlags() {
	args := GetArgs()
	if args == nil {
		return
	}

	// returns the index of a slice element
	indexOf := func(slice []string, element string) int {
		for k, v := range slice {
			if element == v {
				return k
			}
		}
		return -1 // not found.
	}

	// returns the slice with a removed element
	remove := func(slice []string, s string) []string {
		return append(slice[:indexOf(slice, s)], slice[indexOf(slice, s)+1:]...)
	}

	for _, arg := range args {
		if arg[0] == '-' {
			splitindex := 1
			if len(arg) > 1 && arg[1] == '-' {
				splitindex = 2
			}

			search := flag.CommandLine.Lookup(strings.Split(arg, "-")[splitindex])

			if arg == "-h" || arg == "--help" || arg == "-help" || search != nil {
				continue
			}

			os.Args = remove(os.Args, arg)
		}
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
