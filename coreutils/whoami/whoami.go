package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"codeberg.org/whou/simpleutils/internal/cmd"
)

var binary = "whoami"
var usage = `Usage: %s [OPTION]...
%s prints the current user name.

`

func runFlags() {
	cmd.Init(binary, usage, binary, binary)
	cmd.Parse()
}

func main() {
	runFlags()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	uname := user.Username

	// in order to match unix's whoami
	if runtime.GOOS == "windows" {
		uname = strings.Split(uname, "\\")[1]
	}

	fmt.Println(uname)

	os.Exit(0)
}
