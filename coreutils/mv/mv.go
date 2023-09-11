// TODO:
// - flags:
//   -v, --verbose
//   -f, --force
//   -i, --interactive
//   --strip-trailing-slashes

package main

import (
	"os"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "mv"
var usage = `Usage: %s [OPTION]... SOURCE DEST
   or: %s [OPTION]... SOURCE... DIRECTORY
Rename SOURCE to DEST, or move SOURCE(s) to DIRECTORY.

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)

	cmd.Parse()
}

func move(source, dest string) {
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

	destExists, err := myio.FileExists(dest)
	if err != nil {
		panic(err)
	}
	if destExists {
		isDir, err := myio.FileIsDir(dest)
		if err != nil {
			panic(err)
		}

		if isDir {
			var extraSlash string
			if dest[len(dest)-1] != '/' {
				extraSlash = "/"
			}

			dest += extraSlash + myio.GetFilepathBasename(source)
		}
	}

	err = os.Rename(source, dest)
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
		move(src, dest)
	}

	os.Exit(0)
}
