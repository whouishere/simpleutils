package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestTrue(t *testing.T) {
	if os.Getenv("BE_TRUE") == "1" {
		main()
		return
	}

	// invoke go test and return with main() to get the exit code
	cmd := exec.Command(os.Args[0], "-test.run=TestTrue")
	// inject dummy environment variable to avoid infinite loop
	cmd.Env = append(os.Environ(), "BE_TRUE=1")
	err := cmd.Run()

	// if process unexpectedly exited with failure
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Fatalf("Process ran with err %v, expected exit status 0.", err)
	}
}
