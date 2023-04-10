package main

import (
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
var unbufferedFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	numberFlag = cmd.NewFlag(false,
		"number", "n",
		"number all output lines")

	// this flag behaviour is the default, thus it is ignored.
	unbufferedFlag = cmd.NewFlag(false,
		"u", "u",
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

// scan given file list
func scan(files []*os.File, isDone *bool) {
	for _, file := range files {
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

		lines := strings.Split(string(bytes), "\n")
		lines = lines[:len(lines)-1]

		count := 1
		for _, line := range lines {
			if *numberFlag {
				printLineCount(count)
			}

			fmt.Println(line)
			count++
		}

		if *isDone {
			break
		}

		continue
	}
}

// get files from the command-line arguments
func getFiles() []*os.File {
	var files []*os.File

	// if no flags are passed, just read from STDIN
	// we don't use flag.Args() because it doesn't detect '-' only args
	paths := cmd.GetNonFlags()
	if paths == nil {
		return []*os.File{os.Stdin}
	}

	for _, path := range paths {
		// if argument is "-" add STDIN to read
		if path == "-" {
			files = append(files, os.Stdin)
			continue
		}

		exist, err := myio.FileExists(path)
		if err != nil {
			panic(err)
		}

		if exist {
			file, err := os.Open(path)
			if err != nil {
				panic(nil)
			}

			files = append(files, file)
		}
	}

	return files
}

func main() {
	var done bool = false

	runFlags()

	files := getFiles()

	// politely stop when termination signal (Ctrl + C) is received
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done = true
	}()

	scan(files, &done)

	os.Exit(0)
}
