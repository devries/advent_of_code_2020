package main

import (
	"strings"
	"testing"
)

var testInput = `F10
N3
F7
R90
F11`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testInput)
	directions := parseInput(r)

	result := solve(directions)
	expected := 286
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
