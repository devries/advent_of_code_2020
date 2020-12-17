package main

import (
	"strings"
	"testing"
)

var testInput = `.#.
..#
###`

func TestSolution(t *testing.T) {
	r := strings.NewReader(testInput)

	grid := parseInput(r)
	result := solve(grid)
	expected := 112
	if result != expected {
		t.Errorf("Test input: expected %d, got %d", expected, result)
	}
}
