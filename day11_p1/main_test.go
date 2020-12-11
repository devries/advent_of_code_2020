package main

import (
	"strings"
	"testing"
)

var testInput = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testInput)
	seating := parseInput(r)

	result := solve(seating)
	expected := 37
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
