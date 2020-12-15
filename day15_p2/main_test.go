package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"0,3,6", 175594},
		{"1,3,2", 2578},
		{"2,1,3", 3544142},
		{"1,2,3", 261214},
		{"2,3,1", 6895259},
		{"3,2,1", 18},
		{"3,1,2", 362},
	}

	for _, test := range tests {
		result := solve(test.input)
		if result != test.output {
			t.Errorf("For %s expected %d, got %d", test.input, test.output, result)
		}
	}
}
