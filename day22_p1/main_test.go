package main

import (
	"strings"
	"testing"
)

var testInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testInput)

	da, db := parseInput(r)

	playGame(da, db)
	result := max(da.Score(), db.Score())
	expected := int64(306)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
