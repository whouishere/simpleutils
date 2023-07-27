package main

import (
	"io"
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "cp"
var usage = `Usage: %s [OPTION]... SOURCE... DIRECTORY
Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.

`

func runFlags() {
	cmd.Init(binary, usage, binary)

	cmd.Parse()
}

func copy(source, dest string) {
	exist, err := myio.FileExists(source)
	if err != nil {
		panic(err)
	}
	if !exist {
		cmd.Error("Cannot find '", source, "': No such file or directory")
		return
	}

	isDir, err := myio.FileIsDir(source)
	if err != nil {
		panic(err)
	}
	if isDir {
		cmd.Log("Recursive copy not supported; ommiting directory '", source, "'")
		return
	}

	destNeedsTrailingSlash := false

	exist, err = myio.FileExists(dest)
	if err != nil {
		panic(err)
	}
	if exist {
		if dest[len(dest)-1] == '/' {
			isDir = true
		} else {
			isDir, err = myio.FileIsDir(dest)
			if err != nil {
				panic(err)
			}
			destNeedsTrailingSlash = isDir
		}
	} else if dest[len(dest)-1] == '/' { // directory that doesn't exist
		cmd.FatalError("Cannot copy to '", dest, "': Directory doesn't exist")
	}

	// isDir represents dest
	if isDir {
		var extraSlash string
		if destNeedsTrailingSlash {
			extraSlash = "/"
		}

		dest += extraSlash + myio.GetFilepathBasename(source)
	}

	srcFile, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		panic(err)
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument.")
	}
	if len(args) == 1 {
		cmd.FatalHelpError("Missing destination file after '", args[0], "'.")
	}

	sources := args[:len(args)-1]
	dest := args[len(args)-1]

	for _, src := range sources {
		copy(src, dest)
	}

	os.Exit(0)
}
