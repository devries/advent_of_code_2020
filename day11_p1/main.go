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
			} else if val == '#' && surroundingSum >= 4 {
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

	for pos := range seats {
		list := make([]utils.Point, 0, 8)
		for i := pos.X - 1; i <= pos.X+1; i++ {
			for j := pos.Y - 1; j <= pos.Y+1; j++ {
				adjPos := utils.Point{i, j}
				if seats[adjPos] != 0 && adjPos != pos {
					list = append(list, adjPos)
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
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			v := seats[utils.Point{i, j}]
			if v == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", v)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n--\n")
}
