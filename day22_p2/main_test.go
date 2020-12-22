package main

import (
	"os"
	"strings"
	"testing"

	"github.com/devries/advent_of_code_2020/utils"
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

func BenchmarkProgram(b *testing.B) {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	daorig, dborig := parseInput(f)

	for n := 0; n < b.N; n++ {
		da := make(Deck, len(*daorig))
		copy(da, *daorig)
		db := make(Deck, len(*dborig))
		copy(db, *dborig)

		playGame(&da, &db)
	}
}
