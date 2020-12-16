package main

import (
	"strings"
	"testing"
)

var testInput = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestConsolidation(t *testing.T) {
	tests := []struct {
		input  []Range
		output []Range
	}{
		{[]Range{Range{1, 5}, Range{3, 8}}, []Range{Range{1, 8}}},
		{[]Range{Range{1, 3}, Range{4, 8}, Range{5, 6}}, []Range{Range{1, 3}, Range{4, 8}}},
	}

	for _, test := range tests {
		result := consolidateRanges(test.input)

		if len(result) != len(test.output) {
			t.Errorf("For %v: expected %v, got %v", test.input, test.output, result)
			continue
		}

		for i, r := range result {
			if r != test.output[i] {
				t.Errorf("For %v: expected %v, got %v", test.input, test.output, result)
				break
			}
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		rs     []Range
		val    int
		output bool
	}{
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}}, 11, true},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}}, 9, false},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}}, 2, true},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}}, 20, true},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}}, 25, false},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}}, 0, false},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}, Range{25, 40}}, 42, false},
		{[]Range{Range{1, 3}, Range{5, 8}, Range{10, 15}, Range{18, 20}, Range{25, 40}}, 11, true},
	}

	for _, test := range tests {
		result := rangesContain(test.rs, test.val)

		if result != test.output {
			t.Errorf("For %d in %v expected %t got %t", test.val, test.rs, test.output, result)
		}
	}
}

func TestSolution(t *testing.T) {
	r := strings.NewReader(testInput)

	result := solve(r)
	var expected int64 = 1
	if result != expected {
		t.Errorf("Test input: expected %d, got %d", expected, result)
	}
}
