package cmd

import (
	"fmt"
	"os"

	"codeberg.org/whou/simpleutils/coreutils"
)

var errorPrefix string

// Set error prefix string displayed before error messages
func SetErrorPrefix(msg string, a ...any) {
	errorPrefix = fmt.Sprintf("%s%s ", msg, fmt.Sprint(a...))
}

// Print soft message-only log message
func Log(msg string, a ...any) {
	fmt.Printf("%s: %s%s\n", coreutils.Binary, msg, fmt.Sprint(a...))
}

// Print error message without exiting
func Error(msg string, a ...any) {
	fmt.Printf("%s: %s%s%s\n", coreutils.Binary, errorPrefix, msg, fmt.Sprint(a...))
}

// Print error message to stderr and exit with code 1
func FatalStderr(msg string, a ...any) {
	fmt.Fprintf(os.Stderr, "%s: %s%s%s\n", coreutils.Binary, errorPrefix, msg, fmt.Sprint(a...))
	os.Exit(1)
}

// Print error message and exit with code 1
func FatalError(msg string, a ...any) {
	fmt.Printf("%s: %s%s%s\n", coreutils.Binary, errorPrefix, msg, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatally prints an error message to stderr and recommend the help flag
func FatalHelpError(msg string, a ...any) {
	FatalStderr(msg, fmt.Sprint(a...), "\nUse '", coreutils.Binary, " --help' for more information.")
}
