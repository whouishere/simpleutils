package test

import (
	"fmt"
	"os"
	"os/exec"
)

// return the function exit code (from os.Exit)
//
// returns -1 and an error if the function process hasn't exited, was terminated by a signal, triggered and error, or if the function didn't exit by itself
func GetExitCode(function func(), testName string) (int, error) {
	// avoid process loop
	if os.Getenv("IS_TEST") == "1" {
		function()
		// this will only be reached if function didn't use os.Exit
		return -1, fmt.Errorf("the passed in function hasn't exited by itself")
	}

	// invoke itself to retrieve the error code
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%s", testName))
	cmd.Env = append(os.Environ(), "IS_TEST=1")
	err := cmd.Run()

	var code int

	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	} else if err != nil {
		code = -1
	}

	return code, err
}

// get the stdout and stderr from a function
func GetOutput(function func(), testName string) (string, error) {
	// avoid process loop
	if os.Getenv("IS_TEST") == "1" {
		function()
		// this will only be reached if function didn't use os.Exit
		return "", fmt.Errorf("the passed in function hasn't exited by itself")
	}

	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%s", testName))
	cmd.Env = append(os.Environ(), "IS_TEST=1")
	out, err := cmd.CombinedOutput()

	return string(out), err
}
