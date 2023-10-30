// TODO:
// - flags:
//   -v, --verbose
//   -f, --force
//   -H
//   -i, --interactive
//   -L, --dereference
//   -P, --no-dereference
//   -p
//   -r, --recursive

package main

import (
	"io"
	"os"
	"path/filepath"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "cp"
var usage = `Usage: %s [OPTION]... SOURCE DEST
   or: %s [OPTION]... SOURCE... DIRECTORY
Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.

`

func runFlags() {
	cmd.Init(binary, usage, binary)

	cmd.Parse()
}

func copy(source, dest string) {
	// FIXME: trying to reference an existing file with a trailing slash panics
	exist, err := myio.FileExists(source)
	if err != nil {
		panic(err)
	}
	if !exist {
		cmd.SetErrorPrefix("Cannot find '", source, "'")
		cmd.Error("No such file or directory")
		return
	}

	isDir, err := myio.FileIsDir(source)
	if err != nil {
		panic(err)
	}

	// can't copy directories yet
	if isDir {
		cmd.Log("Recursive copy not supported; ommiting directory '", source, "'")
		return
	}

	destNeedsTrailingSlash := false

	cmd.SetErrorPrefix("Cannot copy to '", dest, "'")

	exist, err = myio.FileExists(dest)
	if err != nil {
		panic(err)
	}
	if exist {
		// is dir
		if dest[len(dest)-1] == '/' {
			isDir = true
		} else {
			isDir, err = myio.FileIsDir(dest)
			if err != nil {
				panic(err)
			}
			destNeedsTrailingSlash = isDir
		}
	} else if dest[len(dest)-1] == '/' {
		// the directory in which the source would be pasted in doesn't exist
		cmd.FatalError("Directory doesn't exist")
	} else {
		// dest is not a directory, it's a file name
		// check if the new file directory exists
		path := filepath.Dir(dest)
		exist, err = myio.FileExists(path)
		if err != nil {
			panic(err)
		}
		if !exist {
			cmd.FatalError("No such file or directory")
		}

		pathIsDir, err := myio.FileIsDir(path)
		if err != nil {
			panic(err)
		}
		if !pathIsDir {
			cmd.FatalError("No such file or directory")
		}
	}

	// if dest exists and is a dir we would like to explicit it with a slash
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
		cmd.FatalHelpError("Missing command-line argument")
	}
	if len(args) == 1 {
		cmd.FatalHelpError("Missing destination file after '", args[0], "'")
	}

	sources := args[:len(args)-1]
	dest := args[len(args)-1]

	for _, src := range sources {
		copy(src, dest)
	}

	os.Exit(0)
}
