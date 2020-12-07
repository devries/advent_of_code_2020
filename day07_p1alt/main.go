package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening file")
	rules := parseInput(f)
	bags := recurseInsides("shiny gold", rules)

	fmt.Println(len(bags))
}

// Rules are returned as a map of bag colors containing a map of the contained
// colors to the quantity within.
func parseInput(r io.Reader) map[string]map[string]int {
	lines := utils.ReadLines(r)
	rules := make(map[string]map[string]int)
	pattern := regexp.MustCompile(`(\d+)\s+([a-z]+ [a-z]+)\s+bags?`)

	for _, line := range lines {
		parts := strings.Split(line, "bags contain")

		container := strings.TrimSpace(parts[0])
		submatches := pattern.FindAllStringSubmatch(parts[1], -1)
		subrule := make(map[string]int)
		if submatches == nil {
			rules[container] = subrule
			continue
		}
		for _, subparts := range submatches {
			q, err := strconv.Atoi(subparts[1])
			utils.Check(err, "error parsing rules")
			subrule[subparts[2]] = q
		}
		rules[container] = subrule
	}

	return rules
}

// Find all bags which recursively contain a bag of color
func recurseInsides(color string, rules map[string]map[string]int) []string {
	known := make(map[string]bool)
	result := []string{}

	for k, _ := range rules {
		if contains(k, color, rules, known) {
			result = append(result, k)
		}
	}

	return result
}

// Recursively search container for bag of color using rules, and memoize with known if we know already
func contains(container string, color string, rules map[string]map[string]int, known map[string]bool) bool {
	if r, ok := known[container]; ok == true {
		return r
	}

	if rules[container][color] != 0 {
		known[container] = true
		return true
	}
	if len(rules[container]) == 0 {
		known[container] = false
		return false
	}

	for subcontainer, _ := range rules[container] {
		if contains(subcontainer, color, rules, known) {
			known[container] = true
			return true
		}
	}

	known[container] = false
	return false
}
