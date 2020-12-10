package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	list := parseInput(f)
	a := solve(list)
	fmt.Println(a)
}

func parseInput(r io.Reader) []int {
	lines := utils.ReadLines(r)
	result := []int{}

	for _, line := range lines {
		val, err := strconv.Atoi(line)
		utils.Check(err, "error reading input")
		result = append(result, val)
	}

	return result
}

// Do the whole solution for part b
func solve(input []int) int64 {
	sort.Ints(input)

	max := input[len(input)-1]
	options := make(map[int][]int)

	for i, v := range input {
		if v == max {
			break
		}

		var tests []int
		if i+4 < len(input) {
			tests = input[i+1 : i+4]
		} else {
			tests = input[i+1:]
		}

		for _, v2 := range tests {
			if v2-v <= 3 {
				options[v] = append(options[v], v2)
			}
		}
	}

	for _, v := range input[:4] {
		if v <= 3 {
			options[0] = append(options[0], v)
		}
	}

	known := make(map[int]int64)
	v := countCombinations(0, options, max, known)
	return v
}

func countCombinations(start int, options map[int][]int, max int, known map[int]int64) int64 {
	var total int64 = 0

	if start == max {
		total = 1
		return total
	}
	for _, trial := range options[start] {
		if known[trial] == 0 {
			total += countCombinations(trial, options, max, known)
		} else {
			total += known[trial]
		}
	}

	known[start] = total
	return total
}
