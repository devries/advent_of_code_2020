package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2020/utils"
)

type Point struct {
	X int
	Y int
}

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

type TreeMap struct {
	Grid   map[Point]rune
	Width  int
	Height int
}

func (t TreeMap) Get(p Point) rune {
	return t.Grid[Point{p.X % t.Width, p.Y}]
}

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	grid, err := parseInput(f)
	utils.Check(err, "error parsing input")

	slopes := []Point{Point{1, 1}, Point{3, 1}, Point{5, 1}, Point{7, 1}, Point{1, 2}}
	treecounts := make([]int, len(slopes))

	for i, s := range slopes {
		treecounts[i] = checkSlope(s, grid)
	}
	// fmt.Println(treecounts)

	ans := 1
	for _, c := range treecounts {
		ans *= c
	}
	fmt.Println(ans)
}

func checkSlope(slope Point, grid TreeMap) int {
	trees := 0
	pos := Point{0, 0}
	for {
		pos = pos.Add(slope)
		if pos.Y >= grid.Height {
			break
		}
		if grid.Get(pos) == '#' {
			trees++
		}
	}
	return trees
}

func parseInput(r io.Reader) (TreeMap, error) {
	result := TreeMap{make(map[Point]rune), 0, 0}

	scanner := bufio.NewScanner(r)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		squares := []rune(line)

		result.Width = len(squares)

		for i, v := range squares {
			result.Grid[Point{i, y}] = v
		}
		y++
	}
	err := scanner.Err()
	if err != nil {
		return result, fmt.Errorf("error scanning input: %s", err)
	}
	result.Height = y

	return result, nil
}
