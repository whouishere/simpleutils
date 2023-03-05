package cmd

import (
	"fmt"
	"log"
	"os"
)

var errorPrefix string

// Set error prefix string displayed before error messages
func SetErrorPrefix(msg string, a ...any) {
	errorPrefix = fmt.Sprintf("%s%s ", msg, fmt.Sprint(a...))
}

// Print soft message-only log message
func Log(msg string, a ...any) {
	fmt.Printf("%s: %s%s\n", Binary, msg, fmt.Sprint(a...))
}

// Print error message without exiting
func Error(msg string, a ...any) {
	fmt.Printf("%s: %s%s%s\n", Binary, errorPrefix, msg, fmt.Sprint(a...))
}

// Print error message to stderr and exit with code 1
func FatalStderr(msg string, a ...any) {
	fmt.Fprintf(os.Stderr, "%s: %s%s%s\n", Binary, errorPrefix, msg, fmt.Sprint(a...))
	os.Exit(1)
}

// Print error message and exit with code 1
func FatalError(msg string, a ...any) {
	log.Fatalf("%s: %s%s%s\n", Binary, errorPrefix, msg, fmt.Sprint(a...))
}
