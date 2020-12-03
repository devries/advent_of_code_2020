package main

import (
	"strings"
	"testing"
)

var testInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestTreefinder(t *testing.T) {
	tests := []struct {
		Slope Point
		Trees int
	}{
		{Point{1, 1}, 2},
		{Point{3, 1}, 7},
		{Point{5, 1}, 3},
		{Point{7, 1}, 4},
		{Point{1, 2}, 2},
	}

	r := strings.NewReader(testInput)
	grid, err := parseInput(r)
	if err != nil {
		t.Errorf("Got error: %s parsing input", err)
	}

	for _, test := range tests {
		answer := checkSlope(test.Slope, grid)
		if answer != test.Trees {
			t.Errorf("For slope %v, expected %d got %d", test.Slope, test.Trees, answer)
		}
	}
}
