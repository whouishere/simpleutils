package cmd

import (
	"errors"
	"os"
)

// Returns arguments passed to the command-line as an array. Returns `nil` if no arguments were passed.
func GetArgs() []string {
	args := os.Args[1:]
	if len(args) < 1 {
		return nil
	}

	return args
}

// Returns the specified zero-indexed argument. Returns an error if index is out of range.
func GetArg(argIndex int) (string, error) {
	args := GetArgs()
	if argIndex < 0 || argIndex > len(args)-1 {
		return os.Args[0], errors.New("cmd error: tried to access non-existence or out of range command-line argument")
	}

	return args[argIndex], nil
}

func GetNonFlags() []string {
	var nonFlags []string

	args := GetArgs()
	if args == nil {
		return nil
	}

	for _, arg := range args {
		if arg == "-" {
			nonFlags = append(nonFlags, arg)
		}

		if []rune(arg)[0] == '-' {
			continue
		}

		nonFlags = append(nonFlags, arg)
	}

	return nonFlags
}
