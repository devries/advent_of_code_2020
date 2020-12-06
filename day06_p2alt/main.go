package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")

	declarations, err := parseCustoms(f)
	utils.Check(err, "error parsing input")

	sum := sumCounts(declarations)

	fmt.Println(sum)
}

func parseCustoms(r io.Reader) ([][]string, error) {
	result := [][]string{}
	current := []string{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error scanning input: %s", err)
	}
	if len(current) != 0 {
		result = append(result, current)
	}

	return result, nil
}

func getYesAll(declarations []string) []rune {
	set := make(map[rune]int)

	for _, decl := range declarations {
		for _, char := range decl {
			set[char] = set[char] + 1
		}
	}

	result := []rune{}
	ndecl := len(declarations)

	for k, v := range set {
		if v == ndecl {
			result = append(result, k)
		}
	}

	return result
}

func sumCounts(groups [][]string) int {
	sum := 0
	for _, group := range groups {
		m := getYesAll(group)
		sum += len(m)
	}

	return sum
}
