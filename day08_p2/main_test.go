package main

import (
	"strings"
	"testing"
)

var testProgram = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testProgram)
	instructions := parseInput(r)

	a, err := mutateAndRun(instructions)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if a != 8 {
		t.Errorf("Expected accumulator of 8, got $d")
	}
}
