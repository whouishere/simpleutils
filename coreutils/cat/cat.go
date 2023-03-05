package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "cat"
var usage = `Usage: %s [OPTION]... [FILE]...
%s concatenates FILE(s) to the standard output.
If no FILE is given, or if FILE is -, the standard input is read.

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)
	cmd.Parse()
}

// scan given file list
func scan(files []*os.File, isDone *bool) {
	for _, file := range files {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())

			if *isDone {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
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
