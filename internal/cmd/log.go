package cmd

import (
	"fmt"
	"log"
	"os"
)

var errorPrefix string

func SetErrorPrefix(msg string, a ...any) {
	errorPrefix = fmt.Sprintf("%s%s ", msg, fmt.Sprint(a...))
}

func Error(msg string, a ...any) {
	fmt.Printf("%s: %s%s%s\n", Binary, errorPrefix, msg, fmt.Sprint(a...))
}

func FatalStderr(msg string, a ...any) {
	fmt.Fprintf(os.Stderr, "%s: %s%s%s\n", Binary, errorPrefix, msg, fmt.Sprint(a...))
	os.Exit(1)
}

func FatalError(msg string, a ...any) {
	log.Fatalf("%s: %s%s%s\n", Binary, errorPrefix, msg, fmt.Sprint(a...))
}
