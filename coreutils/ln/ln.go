// TODO:
// - flags:
//   -v, --verbose
//   -L, --logical
//   -t, --target-directory
//   -T, --no-target-directory

package main

import (
	"os"
	"path/filepath"
	"runtime"

	"codeberg.org/whou/simpleutils/internal/cmd"
	myio "codeberg.org/whou/simpleutils/internal/io"
)

var binary = "ln"
var usage = `Usage: %s [OPTION]... [-T] TARGET LINK_NAME
or:  %s [OPTION]... TARGET
or:  %s [OPTION]... TARGET... DIRECTORY
or:  %s [OPTION]... -t DIRECTORY TARGET...
In the 1st form, create a link to TARGET with the name LINK_NAME.
In the 2nd form, create a link to TARGET in the current directory.
In the 3rd and 4th forms, create links to each TARGET in DIRECTORY.
Create hard links by default, symbolic links with --symbolic.
By default, each destination (name of new link) should not already exist.
When creating hard links, each TARGET must exist.  Symbolic links
can hold arbitrary text; if later resolved, a relative link is
interpreted in relation to its parent directory.

`

var symbolicFlag *bool

func runFlags() {
	cmd.Init(binary, usage, binary, binary, binary, binary)

	symbolicFlag = cmd.NewFlag(false,
		"symbolic", "s",
		"make symbolic links instead of hard links")

	cmd.Parse()
}

func link(target string, dest string) {
	cmd.SetErrorPrefix("Failed to link '", target, "'")

	targetExists, err := myio.FileExists(target)
	if err != nil {
		panic(err)
	}

	if !targetExists {
		cmd.SetErrorPrefix("Cannot find '", target, "'")
		cmd.Error("No such file or directory")
		return
	}

	if !*symbolicFlag {
		isDir, err := myio.FileIsDir(target)
		if err != nil {
			panic(err)
		}

		if isDir {
			cmd.Error("Hard links not allowed for directories")
			return
		}
	}

	destExists, err := myio.FileExists(dest)
	if err != nil {
		panic(err)
	}

	if !destExists && dest[len(dest)-1] == '/' {
		cmd.SetErrorPrefix("Failed to link to '", dest, "'")
		cmd.FatalError("No such file or directory")
	}

	if destExists {
		isDir, err := myio.FileIsDir(dest)
		if err != nil {
			panic(err)
		}

		if !isDir {
			cmd.FatalError("File exists")
		}

		// ln foo bar/<dir>
		dest = filepath.Join(dest, myio.GetFilepathBasename(target))
	}

	if *symbolicFlag {
		err := os.Symlink(target, dest)
		if err != nil {
			switch err.(type) {
			case *os.LinkError:
				if runtime.GOOS == "windows" {
					cmd.FatalError("You need administrator rights to make a symbolic link on Windows")
				}
			}

			panic(err)
		}
	} else {
		err := os.Link(target, dest)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	runFlags()

	args := cmd.GetNonFlags()
	if args == nil {
		cmd.FatalHelpError("Missing command-line argument")
	}

	if len(args) == 1 {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		link(args[0], pwd)
		os.Exit(0)
	}

	targets := args[:len(args)-1]
	dest := args[len(args)-1]

	for _, target := range targets {
		link(target, dest)
	}

	os.Exit(0)
}
