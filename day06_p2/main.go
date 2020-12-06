package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

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
	groupings := pattern.Split(input, -1)
	for i, p := range groupings {
		groupings[i] = strings.TrimSpace(p)
	}
	return groupings
}

func getGroupAll(group string) []rune {
	set := make(map[rune]int)
	pattern := regexp.MustCompile(`(?m)^\s*([a-z]+)\s*$`)

	declarations := pattern.FindAllStringSubmatch(group, -1)
	people := len(declarations)

	for _, subm := range declarations {
		for _, char := range subm[1] {
			if char >= 'a' && char <= 'z' {
				set[char] = set[char] + 1
			}
		}
	}

	result := []rune{}

	for k, v := range set {
		if v == people {
			result = append(result, k)
		}
	}

	return result
}

func sumCounts(groups []string) int {
	sum := 0
	for _, group := range groups {
		m := getGroupAll(group)
		sum += len(m)
	}

	return sum
}
