// TODO:
// - fix trailing newlines
// - be able to exit STDIN with CTRL + D
// - flags:
//   --show-all
//   --number-nonblank
//   -e
//   --show-ends
//   --squeeze-blank
//   -t
//   --show-tabs
//   --show-nonprinting

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
	"codeberg.org/whou/simpleutils/internal/util"
)

var binary = "cat"
var usage = `Usage: %s [OPTION]... [FILE]...
%s concatenates FILE(s) to the standard output.
If no FILE is given, or if FILE is -, the standard input is read.

`

var numberFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	numberFlag = cmd.NewFlag(false,
		"number", "n",
		"number all output lines")

	// this flag behaviour is the default, thus it is ignored.
	_ = cmd.NewFlag(false,
		"", "u",
		"(ignored)")

	cmd.Parse()
}

func printLineCount(count int) {
	linecount := ""
	digits := util.LenInt(count)

	// the maximum amount of prefix spaces is 5, when there's 1 digit
	if digits <= 6 && digits > -1 {
		linecount = strings.Repeat(" ", 6-digits)
	}

	linecount += strconv.Itoa(count) + "  "

	fmt.Print(linecount)
}

// scans STDIN and return count
func scanStdin(count int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				return count
			}

			panic(err)
		}

		if *numberFlag {
			printLineCount(count)
		}

		fmt.Print(line)
		count++
	}
}

// scan file and return count
func scan(file *os.File, count int, isDone *bool) int {
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	filelen := stat.Size()
	bytes := make([]byte, filelen)

	_, err = file.Read(bytes)
	if err != nil {
		panic(err)
	}

	lines := strings.SplitAfter(string(bytes), "\n")

	for _, line := range lines {
		if *isDone {
			return count
		}

		if *numberFlag {
			printLineCount(count)
		}

		fmt.Print(line)
		count++
	}

	fmt.Print("\n")
	return count
}

// get files from the command-line arguments
func scanFiles(isDone *bool) {
	count := 1

	// if no flags are passed, just read from STDIN
	paths := cmd.GetNonFlags()
	if paths == nil {
		paths = []string{"-"}
	}

	for _, path := range paths {
		if *isDone {
			break
		}

		// if argument is "-", read from STDIN
		if path == "-" {
			count = scanStdin(count)
			if *isDone {
				break
			}
			continue
		}

		exist, err := myio.FileExists(path)
		if err != nil {
			panic(err)
		}

		if !exist {
			cmd.Error(path, ": No such file or directory")
			continue
		}

		file, err := os.Open(path)
		if err != nil {
			panic(nil)
		}

		count = scan(file, count, isDone)
	}
}

func main() {
	var done bool = false

	runFlags()

	// politely stop when termination signal (Ctrl + C) is received
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done = true
	}()

	scanFiles(&done)

	os.Exit(0)
}
