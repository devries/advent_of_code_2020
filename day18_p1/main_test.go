package main

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, test := range tests {
		stmt := tokenize(test.input)
		r := evalExpression(&stmt)

		if r != test.output {
			t.Errorf("For %s: expected %d, got %d", test.input, test.output, r)
		}
	}
}
