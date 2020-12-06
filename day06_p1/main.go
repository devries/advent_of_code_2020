package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")

	inputBytes, err := ioutil.ReadAll(f)
	utils.Check(err, "error reading input")
	input := string(inputBytes)

	groups := splitBlankLines(input)

	sum := sumCounts(groups)

	fmt.Println(sum)
}

func splitBlankLines(input string) []string {
	pattern := regexp.MustCompile(`(?m)^\s*$`)
	passports := pattern.Split(input, -1)

	return passports
}

func getGroupSet(group string) map[rune]bool {
	result := make(map[rune]bool)

	for _, char := range group {
		if char >= 'a' && char <= 'z' {
			result[char] = true
		}
	}

	return result
}

func sumCounts(groups []string) int {
	sum := 0
	for _, group := range groups {
		m := getGroupSet(group)
		sum += len(m)
	}

	return sum
}
