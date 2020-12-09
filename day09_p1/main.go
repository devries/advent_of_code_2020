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

	list := parseInput(f)
	a := findAnomaly(list, 25)
	fmt.Println(a)
}

func parseInput(r io.Reader) []int64 {
	lines := utils.ReadLines(r)
	result := []int64{}

	for _, line := range lines {
		val, err := strconv.ParseInt(line, 10, 64)
		utils.Check(err, "error reading input")
		result = append(result, val)
	}

	return result
}

// find first number which is not a sum of the previous
// n numbers, where n is the preambleLength
func findAnomaly(vals []int64, preambleLength int) int64 {
	previousSet := make(map[int64]int)

	// Load up preamble
	for i := 0; i < preambleLength; i++ {
		previousSet[vals[i]] = previousSet[vals[i]] + 1
	}

	for i := preambleLength; i < len(vals); i++ {
		current := vals[i]
		found := false
		for j := i - preambleLength; j < i; j++ {
			previousSet[vals[j]]--
			desired := current - vals[j]
			if previousSet[desired] > 0 {
				found = true
				previousSet[vals[j]]++
				break
			}
			previousSet[vals[j]]++
		}
		if found == false {
			return vals[i]
		}
		previousSet[vals[i-preambleLength]]--
		previousSet[vals[i]]++
	}
	return -1
}
