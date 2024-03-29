package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "rm"
var usage = `Usage %s: [OPTION] [FILE]...
Remove (unlink) the FILE(s).

`

var forceFlag *bool
var recursiveFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary)

	forceFlag = cmd.NewFlag(false,
		"force", "f",
		"ignore nonexistent files")

	recursiveFlag = cmd.NewFlag(false,
		"recursive", "r",
		"remove directories and their contents recursively")

	cmd.Parse()
}

func removeFile(filename string) {
	cmd.SetErrorPrefix("Cannot remove '", filename, "'")

	exist, err := myio.FileExists(filename)
	if err != nil {
		panic(err)
	}
	if !exist {
		if !*forceFlag {
			cmd.Error("No such file or directory")
		}
		return
	}

	if !*recursiveFlag {
		isDir, err := myio.FileIsDir(filename)
		if err != nil {
			panic(err)
		}

		if isDir {
			cmd.Error("Is a directory")
			return
		}
	}

	// os.RemoveAll calls os.Remove first anyway
	err = os.RemoveAll(filename)
	if err != nil {
		panic(err)
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument")
	}

	for _, arg := range args {
		removeFile(arg)
	}

	os.Exit(0)
}
