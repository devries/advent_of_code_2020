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

	winner := playGame(da, db)
	var score int64
	if winner == 1 {
		score = da.Score()
	} else {
		score = db.Score()
	}
	expected := int64(291)
	if score != expected {
		t.Errorf("expected %d, got %d", expected, score)
	}
}
