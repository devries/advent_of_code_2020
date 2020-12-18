package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, test := range tests {
		stmt := tokenize(test.input)
		r := evalExpression(&stmt)

		if r != test.output {
			t.Errorf("For %s: expected %d, got %d", test.input, test.output, r)
		}
	}
}
