package cmd

import (
	flag "github.com/erikjuhani/miniflag"
)

type Flag[T any] struct {
	Value        *T
	defaultValue T
	name         string
	shorthand    string
	usage        string
	Function     func()
}

func NewFlag[T any](value T, name, shorthand, usage string, function func()) *Flag[T] {
	return &Flag[T]{
		defaultValue: value,
		name:         name,
		shorthand:    shorthand,
		usage:        usage,
		Function:     function,
	}
}

// registers every passed command-line flag
func RegisterFlag[T any](flagVar *Flag[T]) {
	flagVar.Value = flag.Flag(flagVar.name, flagVar.shorthand, flagVar.defaultValue, flagVar.usage)
}
