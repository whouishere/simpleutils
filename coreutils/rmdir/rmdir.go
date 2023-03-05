package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "rmdir"
var usage = `Usage: %s [OPTION(s)]... DIRECTORY
%s removes empty DIRECTORY(ies).

`

var parentsFlag *cmd.Flag[bool]

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	parentsFlag = cmd.NewFlag(false,
		"parents", "p",
		"remove DIRECTORY and its empty parents",
		nil)
	cmd.RegisterFlag(parentsFlag)

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

	// disable parents flag and remove empty parents
	*parentsFlag.Value = false
	for i := len(parents) - 1; i >= 0; i-- { // reverse iterate
		removeDir(parents[i])
	}
}

func removeDir(rmdir string) {
	cmd.SetErrorPrefix("Failed to remove '", rmdir, "'.")

	exist, err := myio.FileExists(rmdir)
	if err != nil {
		panic(err)
	}
	if !exist {
		cmd.FatalError("Directory does not exist.")
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

		if *parentsFlag.Value {
			parentRemove(rmdir)
		}
	} else {
		cmd.FatalError("Directory is not empty.")
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalStderr("Missing command-line argument.\nUse '", binary, " --help' for more information.")
	}
	arg := cmd.GetNonFlags()[0]

	removeDir(arg)

	os.Exit(0)
}
