package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"0,3,6", 436},
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for _, test := range tests {
		result := solve(test.input)
		if result != test.output {
			t.Errorf("For %s expected %d, got %d", test.input, test.output, result)
		}
	}
}
