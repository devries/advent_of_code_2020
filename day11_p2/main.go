package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/devries/advent_of_code_2020/utils"
	"github.com/spf13/pflag"
)

var visualizer = pflag.BoolP("visualize", "v", false, "enable visualization")

func main() {
	pflag.Parse()
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	seats := parseInput(f)

	result := solve(seats)
	fmt.Println(result)
}

func solve(seats map[utils.Point]rune) int {
	/*
	   If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
	   If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
	   Otherwise, the seat's state does not change.
	*/
	adjacents := getAdjacents(seats)

	generations := 0
	for {
		if *visualizer {
			printMap(seats)
			time.Sleep(200 * time.Millisecond)
		}
		changes := 0
		nextGeneration := make(map[utils.Point]rune)
		for pos, val := range seats {
			surroundingSum := 0
			for _, adjPos := range adjacents[pos] {
				v := seats[adjPos]
				if v == '#' {
					surroundingSum++
				}
			}
			if val == 'L' && surroundingSum == 0 {
				changes++
				nextGeneration[pos] = '#'
			} else if val == '#' && surroundingSum >= 5 {
				changes++
				nextGeneration[pos] = 'L'
			} else {
				nextGeneration[pos] = val
			}
		}
		generations++
		if changes == 0 {
			break
		}
		seats = nextGeneration
	}

	occupied := 0
	for _, v := range seats {
		if v == '#' {
			occupied++
		}
	}

	return occupied
}

func getAdjacents(seats map[utils.Point]rune) map[utils.Point][]utils.Point {
	// Do this once, find adjacent points
	result := make(map[utils.Point][]utils.Point)

	xmax := 0
	ymax := 0
	for p, _ := range seats {
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}

	for pos := range seats {
		list := make([]utils.Point, 0, 8)
		for slopex := -1; slopex <= 1; slopex++ {
			for slopey := -1; slopey <= 1; slopey++ {
				if slopex == 0 && slopey == 0 {
					continue
				}
				for m := 1; ; m++ {
					t := utils.Point{pos.X + slopex*m, pos.Y + slopey*m}
					if seats[t] != 0 {
						list = append(list, t)
						break
					}
					if t.X < 0 || t.X > xmax {
						break
					}
					if t.Y < 0 || t.Y > ymax {
						break
					}
				}
			}
		}
		result[pos] = list
	}

	return result
}

func parseInput(r io.Reader) map[utils.Point]rune {
	lines := utils.ReadLines(r)

	result := make(map[utils.Point]rune)

	for j, line := range lines {
		squares := []rune(line)
		for i, v := range squares {
			if v != '.' {
				result[utils.Point{i, j}] = v
			}
		}
	}

	return result
}

func printMap(seats map[utils.Point]rune) {
	xmax := 0
	ymax := 0
	for p, _ := range seats {
		if p.X > xmax {
			xmax = p.X
		}
		if p.Y > ymax {
			ymax = p.Y
		}
	}

	fmt.Printf("\033[3J\033[H")
	for j := 0; j <= ymax; j++ {
		for i := 0; i <= xmax; i++ {
			v := seats[utils.Point{i, j}]
			if v == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", v)
			}
		}
		fmt.Printf("\n")
	}
}
