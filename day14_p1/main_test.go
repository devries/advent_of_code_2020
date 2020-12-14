package main

import (
	"strings"
	"testing"
)

var testInput = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestMask(t *testing.T) {
	tests := []struct {
		mask   string
		value  uint64
		result uint64
	}{
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 11, 73},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 101, 101},
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", 0, 64},
	}

	for _, test := range tests {
		m := NewPortMask(test.mask)
		r := m.Apply(test.value)

		if r != test.result {
			t.Errorf("For %s -> %d expected %d, got %d", test.mask, test.value, test.result, r)
		}
	}
}

func TestSolution(t *testing.T) {
	r := strings.NewReader(testInput)

	result := solve(r)
	var expected uint64 = 165
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
