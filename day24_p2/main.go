package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/devries/advent_of_code_2020/utils"
)

var directions = map[string]utils.Point{
	"e":  utils.Point{2, 0},
	"w":  utils.Point{-2, 0},
	"ne": utils.Point{1, 1},
	"nw": utils.Point{-1, 1},
	"se": utils.Point{1, -1},
	"sw": utils.Point{-1, -1},
}

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening file")
	defer f.Close()

	r := solve(f)

	fmt.Println(r)
}

func parseLine(line string) []string {
	pattern := regexp.MustCompile(`ne|se|nw|sw|e|w`)

	result := pattern.FindAllString(line, -1)
	return result
}

func findPosition(moves []string) utils.Point {
	p := utils.Point{0, 0}

	for _, move := range moves {
		p = p.Add(directions[move])
	}

	return p
}

func parseInput(r io.Reader) [][]string {
	allMoves := [][]string{}

	for _, line := range utils.ReadLines(r) {
		moves := parseLine(line)
		allMoves = append(allMoves, moves)
	}

	return allMoves
}

func initialMap(r io.Reader) map[utils.Point]bool {
	allMoves := parseInput(r)

	isBlack := make(map[utils.Point]bool)

	for _, moves := range allMoves {
		p := findPosition(moves)
		isBlack[p] = !isBlack[p]
	}

	return isBlack
}

// Find all black tiles and all adjacent white tiles, making blank map
func findRelevantTiles(isBlack map[utils.Point]bool) map[utils.Point]bool {
	next := make(map[utils.Point]bool)

	for k, v := range isBlack {
		if v {
			next[k] = false
			for _, d := range directions {
				next[k.Add(d)] = false
			}
		}
	}

	return next
}

func step(isBlack map[utils.Point]bool) map[utils.Point]bool {
	next := findRelevantTiles(isBlack)

	for k := range next {
		sumBlack := 0
		for _, d := range directions {
			if isBlack[k.Add(d)] {
				sumBlack++
			}
		}
		switch isBlack[k] {
		case true: // black tile
			if sumBlack == 0 || sumBlack > 2 {
				next[k] = false
			} else {
				next[k] = true
			}
		case false:
			if sumBlack == 2 {
				next[k] = true
			}
		}
	}

	return next
}

func solve(r io.Reader) int {
	isBlack := initialMap(r)

	for i := 0; i < 100; i++ {
		isBlack = step(isBlack)
	}

	return countBlack(isBlack)
}

func countBlack(isBlack map[utils.Point]bool) int {
	sum := 0
	for _, v := range isBlack {
		if v {
			sum++
		}
	}

	return sum
}
