package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

const (
	Top int = iota
	Right
	Bottom
	Left
	TopRev
	RightRev
	BottomRev
	LeftRev
)

type Borders map[int][]rune

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening file")
	defer f.Close()

	photos := parseInput(f)

	keys := make([]int, 0, len(photos))
	for k := range photos {
		keys = append(keys, k)
	}

	borderMap := make(map[int]Borders)

	for _, k := range keys {
		borderMap[k] = getBorders(photos[k])
	}

	borderMatches := make(map[int][]int)

	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			b1 := borderMap[keys[i]]
			b2 := borderMap[keys[j]]
		PhotoCompare:
			for _, bOrig := range b1 {
				for _, bCompare := range b2 {
					match := true
					for k := 0; k < 10; k++ {
						if bOrig[k] != bCompare[k] {
							// do next border
							match = false
							break
						}
					}
					// Matching border
					if match {
						borderMatches[keys[i]] = append(borderMatches[keys[i]], keys[j])
						borderMatches[keys[j]] = append(borderMatches[keys[j]], keys[i])
						break PhotoCompare
					}
				}
			}
		}
	}

	var solution int64 = 1
	for k, v := range borderMatches {
		if len(v) == 2 {
			solution *= int64(k)
		}
	}
	fmt.Println(solution)
}

func parseInput(r io.Reader) map[int]map[utils.Point]rune {
	lines := utils.ReadLines(r)

	photos := make(map[int]map[utils.Point]rune)
	var gridTitle int
	var current map[utils.Point]rune
	var err error
	y := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			// Title line
			if current != nil {
				photos[gridTitle] = current
			}

			colonIndex := strings.Index(line, ":")
			gridTitle, err = strconv.Atoi(line[5:colonIndex])
			utils.Check(err, "error parsing tile number")

			current = make(map[utils.Point]rune)
			y = 0
			continue
		}

		for x, r := range line {
			current[utils.Point{x, y}] = r
		}
		y++
	}

	photos[gridTitle] = current

	return photos
}

func getBorders(photo map[utils.Point]rune) Borders {
	result := make(Borders)

	result[Top] = make([]rune, 10)
	result[Right] = make([]rune, 10)
	result[Bottom] = make([]rune, 10)
	result[Left] = make([]rune, 10)
	result[TopRev] = make([]rune, 10)
	result[RightRev] = make([]rune, 10)
	result[BottomRev] = make([]rune, 10)
	result[LeftRev] = make([]rune, 10)

	for i := 0; i < 10; i++ {
		result[Top][i] = photo[utils.Point{i, 0}]
		result[Bottom][i] = photo[utils.Point{i, 9}]
		result[TopRev][9-i] = photo[utils.Point{i, 0}]
		result[BottomRev][9-i] = photo[utils.Point{i, 9}]

		result[Right][i] = photo[utils.Point{9, i}]
		result[Left][i] = photo[utils.Point{0, i}]
		result[RightRev][9-i] = photo[utils.Point{9, i}]
		result[LeftRev][9-i] = photo[utils.Point{0, i}]
	}

	return result
}
