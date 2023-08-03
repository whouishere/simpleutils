package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "rmdir"
var usage = `Usage: %s [OPTION]... DIRECTORY...
%s removes empty DIRECTORY(ies).

`

var nonemptyFlag *bool
var parentsFlag *bool
var verboseFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	parentsFlag = cmd.NewFlag(false,
		"parents", "p",
		"remove DIRECTORY and its empty parents")

	nonemptyFlag = cmd.NewFlag(false,
		"ignore-fail-on-non-empty", "",
		"ignore any fails solely because of non-empty directories")

	verboseFlag = cmd.NewFlag(false,
		"verbose", "v",
		"ouput verbose messages for every processed directory")

	cmd.Parse()
}

func parentRemove(rmdir string) {
	runedir := []rune(rmdir)
	// remove trailing slash
	if runedir[len(runedir)-1] == '/' {
		runedir = runedir[:len(runedir)-1]
	}

	var parents []string
	for i := range runedir {
		if runedir[i] == '/' {
			parents = append(parents, string(runedir[:i]))
		}
	}

	// remove empty parents
	*parentsFlag = false
	for i := len(parents) - 1; i >= 0; i-- { // reverse iterate
		removeDir(parents[i])
	}
	// the flag is temporarily disabled to avoid recursion
	*parentsFlag = true
}

func removeDir(rmdir string) {
	cmd.SetErrorPrefix("Failed to remove '", rmdir, "'")

	// the verbose print is before the "file exists" check on GNU and BusyBox, for some reason
	if *verboseFlag {
		cmd.Log("Removing directory '", rmdir, "'")
	}

	exist, err := myio.FileExists(rmdir)
	if err != nil {
		panic(err)
	}
	if !exist {
		cmd.FatalError("Directory does not exist")
	}

	// open passed directory
	dir, err := os.Open(rmdir)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// read directory entries
	entries, err := dir.Readdirnames(-1)
	if err != nil {
		panic(err)
	}

	// finally remove it if directory is empty
	if len(entries) == 0 {
		dir.Close()
		err := os.Remove(rmdir)
		if err != nil {
			panic(err)
		}

		if *parentsFlag {
			parentRemove(rmdir)
		}
	} else if !*nonemptyFlag {
		cmd.FatalError("Directory is not empty")
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument")
	}

	for _, arg := range args {
		removeDir(arg)
	}

	os.Exit(0)
}
