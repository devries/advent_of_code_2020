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
	boat := utils.Point{0, 0}
	waypoint := utils.Point{10, 1}

	for _, dir := range directions {
		switch dir.Action {
		case 'N':
			waypoint = waypoint.Add(utils.North.Scale(dir.Value))
		case 'S':
			waypoint = waypoint.Add(utils.South.Scale(dir.Value))
		case 'E':
			waypoint = waypoint.Add(utils.East.Scale(dir.Value))
		case 'W':
			waypoint = waypoint.Add(utils.West.Scale(dir.Value))
		case 'L':
			offset := waypoint.Add(boat.Scale(-1))
			for c := 0; c < dir.Value; c += 90 {
				offset = offset.Left()
			}
			waypoint = boat.Add(offset)
		case 'R':
			offset := waypoint.Add(boat.Scale(-1))
			for c := 0; c < dir.Value; c += 90 {
				offset = offset.Right()
			}
			waypoint = boat.Add(offset)
		case 'F':
			offset := waypoint.Add(boat.Scale(-1))
			boat = boat.Add(offset.Scale(dir.Value))
			waypoint = boat.Add(offset)
		}
	}

	return absInt(boat.X) + absInt(boat.Y)
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
