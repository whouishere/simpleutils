package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestFalse(t *testing.T) {
	if os.Getenv("BE_FALSE") == "1" {
		main()
		return
	}

	// invoke go test and return with main() to get the exit code
	cmd := exec.Command(os.Args[0], "-test.run=TestFalse")
	// inject dummy environment variable to avoid infinite loop
	cmd.Env = append(os.Environ(), "BE_FALSE=1")
	err := cmd.Run()

	// if process exited with code 1, as expected
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("Process ran with err %v, expected exit status 1.", err)
}
