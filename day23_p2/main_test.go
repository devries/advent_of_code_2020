package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	r := solve(testInput)
	var expected int64 = 149245887792

	if r != expected {
		t.Errorf("Expected %d, got %d", expected, r)
	}
}

func BenchmarkSolution(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = solve(input)
	}
}
