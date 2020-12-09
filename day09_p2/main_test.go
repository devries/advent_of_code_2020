package main

import (
	"strings"
	"testing"
)

var testInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testInput)
	vals := parseInput(r)

	result := findAnomaly(vals, 5)
	if result != 127 {
		t.Errorf("Expected 127, got %d", result)
	}
	sequence := findSequence(vals, result)
	min, max := findExtrema(sequence)
	if min+max != 62 {
		t.Errorf("Expected 52, got %d", min+max)
	}
}
