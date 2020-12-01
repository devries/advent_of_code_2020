package main

import (
	"testing"
)

func TestFinder(t *testing.T) {
	testInput := []int{1721, 979, 366, 299, 675, 1456}

	testOut := findSum(testInput, 2020)

	if len(testOut) != 3 {
		t.Errorf("Did not get right number of outputs")
	}

	r := testOut[0] * testOut[1] * testOut[2]
	if r != 241861950 {
		t.Errorf("Got result %d", r)
	}
}
