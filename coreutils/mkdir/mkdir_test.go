package main

import (
	"os"
	"testing"

	"codeberg.org/whou/simpleutils/internal/test"
)

func TestParentsNoErrIfDirExists(t *testing.T) {
	binPrefix := os.Getenv("PREFIX")
	if binPrefix == "" {
		t.Fatal("Required PREFIX environment variable not found.")
	}

	binDir := os.Getenv("BUILD_DIR")
	if binDir == "" {
		t.Fatal("Required BUILD_DIR environment variable not found.")
	}

	binPath := binDir + "/" + binPrefix + "mkdir"

	// create directory
	testDir := t.TempDir()

	// try use mkdir in the already created directory with `-p` flag
	out, code, err := test.GetCmdOutput(binPath, "-p", testDir)
	t.Log(out)

	if code != 0 {
		t.Fatalf("Process ran with err \"%v\", expected exit status 0.", err)
	}
}
