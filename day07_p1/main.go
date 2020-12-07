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
	insides := findInside(rules)
	bags := recurseInsides("shiny gold", insides)

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

// Find out what bags are inside what other bags. Returns a map from the bag
// color to the bag colors it is found inside.
func findInside(rules map[string]map[string]int) map[string][]string {
	result := make(map[string][]string)
	for container, contained := range rules {
		for bag, _ := range contained {
			result[bag] = append(result[bag], container)
		}
	}

	return result
}

// Find all bags which recursively contain a bag of color
func recurseInsides(color string, insides map[string][]string) []string {
	bagset := make(map[string]bool)
	bagqueue := []string{color}

	for len(bagqueue) > 0 {
		c := bagqueue[0]
		bagqueue = bagqueue[1:]
		values := insides[c]
		if values == nil {
			continue
		}
		for _, v := range values {
			if bagset[v] == false {
				bagset[v] = true
				bagqueue = append(bagqueue, v)
			}
		}
	}

	result := []string{}
	for k, _ := range bagset {
		result = append(result, k)
	}

	return result
}
