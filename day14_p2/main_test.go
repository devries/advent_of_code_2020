package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

var testInput = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestMask(t *testing.T) {
	tests := []struct {
		mask   string
		value  uint64
		result []uint64
	}{
		{"000000000000000000000000000000X1001X", 42, []uint64{26, 27, 58, 59}},
		{"00000000000000000000000000000000X0XX", 26, []uint64{16, 17, 18, 19, 24, 25, 26, 27}},
	}

	for _, test := range tests {
		m := NewPortMask(test.mask)
		r := m.Apply(test.value)

		sort.Slice(test.result, func(i, j int) bool { return test.result[i] < test.result[j] })
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		fmt.Println(r)
		for i := 0; i < len(r); i++ {
			if r[i] != test.result[i] {
				t.Errorf("For %s -> %d, expected %v, got %v", test.mask, test.value, test.result, r)
			}
		}
	}
}

func TestSolution(t *testing.T) {
	r := strings.NewReader(testInput)

	result := solve(r)
	var expected uint64 = 208
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
