package main

import (
	"testing"
)

func TestTerm(t *testing.T) {
	term := NewTerm()

	term.Normal("Normal\n")
	term.Bold("Bold\n")
	term.Success("Success\n")
	term.Info("Info\n")
	term.Warn("Warning\n")
	term.Error("Error\n")
}
