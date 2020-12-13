package main

import (
	"strings"
	"testing"
)

var testInput = `939
7,13,x,x,59,x,31,19`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testInput)
	timestamp, busses := parseInput(r)

	result := solve(timestamp, busses)
	var expected int64 = 295
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
