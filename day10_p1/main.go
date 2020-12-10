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

// Do the whole solution for part a
func solve(input []int) int {
	sort.Ints(input)

	start := 0
	diffcount := make(map[int]int)

	for _, v := range input {
		diffcount[v-start]++
		start = v
	}

	// add for device
	diffcount[3]++

	return diffcount[1] * diffcount[3]
}
