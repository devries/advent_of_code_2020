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
	n := recurseQuantity("shiny gold", rules)
	fmt.Println(n - 1)
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
func recurseQuantity(color string, rules map[string]map[string]int) int64 {
	var result int64 = 0

	contents := rules[color]
	if len(contents) == 0 {
		return 1
	}

	for k, v := range contents {
		result += int64(v) * recurseQuantity(k, rules)
	}

	return result + 1
}
