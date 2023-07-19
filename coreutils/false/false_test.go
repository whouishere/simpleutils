package main

import (
	"testing"

	"codeberg.org/whou/simpleutils/internal/test"
)

func TestFalse(t *testing.T) {
	code, err := test.GetFuncExitCode(main, t.Name())

	if code != 1 {
		t.Fatalf("Process ran with err \"%v\", expected exit status 1.", err)
	}
}
