package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

var GAMEOVER = errors.New("Game Over")

type Deck []int

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	da, db := parseInput(f)
	playGame(da, db)

	result := max(da.Score(), db.Score())

	fmt.Println(result)
}

func (d *Deck) Draw() int {
	v := (*d)[0]
	*d = (*d)[1:]

	return v
}

func (d *Deck) Add(v int) {
	*d = append(*d, v)
}

func (d *Deck) Score() int64 {
	var sum int64
	l := len(*d)

	for i, v := range *d {
		sum += int64((l - i) * v)
	}

	return sum
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func playGame(da *Deck, db *Deck) {
	for {
		e := playRound(da, db)
		if e != nil {
			break
		}
	}
}

func playRound(da *Deck, db *Deck) error {
	if len(*da) == 0 || len(*db) == 0 {
		return GAMEOVER
	}

	a, b := da.Draw(), db.Draw()

	if a > b {
		da.Add(a)
		da.Add(b)
	} else {
		db.Add(b)
		db.Add(a)
	}
	return nil
}

func parseInput(r io.Reader) (*Deck, *Deck) {
	lines := utils.ReadLines(r)

	deck1 := make(Deck, 0, 25)
	deck2 := make(Deck, 0, 25)
	var incr int = 0
	for _, line := range lines {
		if strings.HasPrefix(line, "Player") {
			incr++
			continue
		}

		if line == "" {
			continue
		}

		v, err := strconv.Atoi(line)
		utils.Check(err, "error converting card to integer")

		switch incr {
		case 1:
			deck1 = append(deck1, v)
		case 2:
			deck2 = append(deck2, v)
		default:
			panic(fmt.Errorf("unexpected deck"))
		}
	}

	return &deck1, &deck2
}
