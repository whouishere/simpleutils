package test

import (
	"fmt"
	"os"
	"os/exec"
)

// run `go test` again and return the function exit code
//
// returns -1 and an error if the function process hasn't exited, was terminated by a signal, triggered and error, or if the function didn't exit by itself
func GetExitCode(function func(), testName string) (int, error) {
	// avoid invocation infinite loop
	if os.Getenv("IS_TEST") == "1" {
		function()
		// this will only be reached if function didn't use os.Exit
		return -1, fmt.Errorf("The passed in function hasn't exited by itself.")
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
