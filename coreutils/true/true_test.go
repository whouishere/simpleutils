package main

import (
	"testing"

	"codeberg.org/whou/simpleutils/internal/test"
)

func TestTrue(t *testing.T) {
	code, err := test.GetExitCode(main, t.Name())

	if code != 0 {
		t.Fatalf("Process ran with err \"%v\", expected exit status 0.", err)
	}
}
