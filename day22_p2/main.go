package main

import (
	"crypto/sha256"
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

func (d *Deck) Hash() string {
	deckRep := fmt.Sprintf("%v", *d)
	h := sha256.New()
	h.Write([]byte(deckRep))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func playGame(da *Deck, db *Deck) int {
	// return winner (1 or 2 for deck a or deck b)
	memory := make(map[string]bool)

	for {
		// Check if decks have been in game before
		if memory[da.Hash()] || memory[db.Hash()] {
			return 1
		}

		// Remember cars
		memory[da.Hash()] = true
		memory[db.Hash()] = true

		// Play a round
		e := playRound(da, db)
		if e != nil {
			// Game ended with empty deck
			if len(*da) > len(*db) {
				return 1
			} else {
				return 2
			}
		}
	}
}

func playRound(da *Deck, db *Deck) error {
	if len(*da) == 0 || len(*db) == 0 {
		return GAMEOVER
	}

	// Draw cards
	a, b := da.Draw(), db.Draw()

	var winner int
	// Check if game will recurse
	if a <= len(*da) && b <= len(*db) {
		// Recursive game, make new subdecks
		dasub := make(Deck, a)
		copy(dasub, *da)
		dbsub := make(Deck, b)
		copy(dbsub, *db)

		// Play recursive game
		winner = playGame(&dasub, &dbsub)
	} else {
		// Play regular game
		if a > b {
			winner = 1
		} else {
			winner = 2
		}
	}

	if winner == 1 {
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
