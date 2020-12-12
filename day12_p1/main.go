package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	directions := parseInput(f)

	result := solve(directions)
	fmt.Println(result)
}

func solve(directions []Move) int {
	boat := Boat{utils.Point{0, 0}, utils.East}

	for _, dir := range directions {
		switch dir.Action {
		case 'N':
			boat.Position = boat.Position.Add(utils.North.Scale(dir.Value))
		case 'S':
			boat.Position = boat.Position.Add(utils.South.Scale(dir.Value))
		case 'E':
			boat.Position = boat.Position.Add(utils.East.Scale(dir.Value))
		case 'W':
			boat.Position = boat.Position.Add(utils.West.Scale(dir.Value))
		case 'L':
			for c := 0; c < dir.Value; c += 90 {
				boat.Direction = boat.Direction.Left()
			}
		case 'R':
			for c := 0; c < dir.Value; c += 90 {
				boat.Direction = boat.Direction.Right()
			}
		case 'F':
			boat.Position = boat.Position.Add(boat.Direction.Scale(dir.Value))
		}
	}

	return absInt(boat.Position.X) + absInt(boat.Position.Y)
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

type Move struct {
	Action rune
	Value  int
}

type Boat struct {
	Position  utils.Point
	Direction utils.Point
}

func parseInput(r io.Reader) []Move {
	lines := utils.ReadLines(r)

	result := []Move{}

	for _, line := range lines {
		ln := []rune(line)
		action := ln[0]
		val, err := strconv.Atoi(string(ln[1:]))
		utils.Check(err, "Error reading integer part of move")
		result = append(result, Move{action, val})
	}

	return result
}
