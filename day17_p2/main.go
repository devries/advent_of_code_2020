package main

import (
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	grid := parseInput(f)
	n := solve(grid)
	fmt.Println(n)
}

func solve(grid map[Point]bool) int {
	for i := 0; i < 6; i++ {
		grid = step(grid)
	}

	sum := 0
	for _, v := range grid {
		if v {
			sum++
		}
	}

	return sum
}

type Point struct {
	X int
	Y int
	Z int
	W int
}

func (p Point) neighbors() []Point {
	r := make([]Point, 0, 80)

	for i := p.X - 1; i <= p.X+1; i++ {
		for j := p.Y - 1; j <= p.Y+1; j++ {
			for k := p.Z - 1; k <= p.Z+1; k++ {
				for l := p.W - 1; l <= p.W+1; l++ {
					pn := Point{i, j, k, l}
					if pn != p {
						r = append(r, pn)
					}
				}
			}
		}
	}

	return r
}

func parseInput(r io.Reader) map[Point]bool {
	result := make(map[Point]bool)

	for j, line := range utils.ReadLines(r) {
		for i, v := range line {
			if v == '#' {
				result[Point{i, j, 0, 0}] = true
			}
		}
	}

	return result
}

func step(grid map[Point]bool) map[Point]bool {
	nextGrid := make(map[Point]bool)

	for k, v := range grid {
		if v == true {
			nextGrid[k] = false
			for _, p := range k.neighbors() {
				nextGrid[p] = false
			}
		}
	}

	for p := range nextGrid {
		sum := 0
		for _, np := range p.neighbors() {
			if grid[np] == true {
				sum++
			}
		}
		if grid[p] == true && (sum == 2 || sum == 3) {
			nextGrid[p] = true
		} else if grid[p] == false && sum == 3 {
			nextGrid[p] = true
		}
	}

	return nextGrid
}
