package main

import (
	"strings"
	"testing"
)

var testInput = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestSolution(t *testing.T) {
	r := strings.NewReader(testInput)

	labels := parseInput(r)
	result := solve(labels)
	var expected int = 5
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
