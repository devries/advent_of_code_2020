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

func solve(r io.Reader) int {
	allMoves := parseInput(r)

	isBlack := make(map[utils.Point]bool)

	for _, moves := range allMoves {
		p := findPosition(moves)
		isBlack[p] = !isBlack[p]
	}

	sum := 0
	for _, v := range isBlack {
		if v {
			sum++
		}
	}

	return sum
}
