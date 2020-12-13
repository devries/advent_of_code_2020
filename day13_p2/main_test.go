package main

import (
	"testing"
)

func TestProblem(t *testing.T) {
	tests := []struct {
		input  string
		result int64
	}{
		{"7,13,x,x,59,x,31,19", 1068781},
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
		{"2,x,4,3", 6},
	}

	for _, test := range tests {
		busses := parseBusses(test.input)

		result := solve(busses)
		if result != test.result {
			t.Errorf("Expected %d, got %d", test.result, result)
		}
	}
}
