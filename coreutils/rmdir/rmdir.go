package main

import (
	"fmt"
	"log"
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "rmdir"
var usage = `Usage: %s [OPTION(s)]... DIRECTORY
%s removes empty DIRECTORY(ies).

`

var recursiveFlag *cmd.Flag[bool]

func runFlags() {
	cmd.Init(binary, usage)

	recursiveFlag = cmd.NewFlag(false,
		"recursive", "r",
		"remove empty DIRECTORY(ies) recursively",
		nil)
	cmd.RegisterFlag(recursiveFlag)

	cmd.Parse()
}

func recursiveRemove(rmdir string, entries []string) {
	initialpath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// cd into the passed directory
	err = os.Chdir(rmdir)
	if err != nil {
		panic(err)
	}

	// loop through the directory's found paths
	for _, path := range entries {
		isdir, err := myio.FileIsDir(path)
		if err != nil {
			panic(err)
		}

		if !isdir {
			fmt.Println("A non-empty directory was passed. Can't remove!")
			os.Exit(1)
		}

		isempty, err := myio.IsDirEmpty(path)
		if err != nil {
			panic(err)
		}

		if !isempty {
			fmt.Println("Directory is not empty. Can't remove!")
			os.Exit(1)
		}
	}

	// cd back to the initial directory
	err = os.Chdir(initialpath)
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll(rmdir)
	if err != nil {
		panic(err)
	}
}

func removeDir(rmdir string) {
	exist, err := myio.FileExists(rmdir)
	if err != nil {
		panic(err)
	}
	if !exist {
		log.Fatalln("This directory does not exist!")
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
		err := os.Remove(rmdir)
		if err != nil {
			panic(err)
		}
	} else {
		if *recursiveFlag.Value {
			dir.Close()
			recursiveRemove(rmdir, entries)
		} else {
			log.Fatalln("Directory is not empty. Can't remove!")
		}
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		fmt.Fprintf(os.Stderr, "Missing command-line argument.\nUse '%s --help' for more information.\n", binary)
		os.Exit(1)
	}
	arg := cmd.GetNonFlags()[0]

	removeDir(arg)

	os.Exit(0)
}
