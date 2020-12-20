package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	rules, messages := parseInput(utils.ReadLines(f))

	sum := 0
	for _, message := range messages {
		if solve(message, rules) {
			sum++
		}
	}
	fmt.Println(sum)
}

func solve(message []byte, rules map[string]Rule) bool {
	r := match(message, rules["0"], rules)

	if r[0] == len(message) {
		return true
	}
	return false
}

func match(message []byte, rule Rule, rules map[string]Rule) []int {
	// Single byte matches value
	if message[0] == rule.Value {
		return []int{1}
	}

	if rule.Value == 0 {
		// Match subrules

		// Loop through possible matches {
		allPositions := []int{}
		for _, sequence := range rule.Options {
			// For each ruleset loop through sequence of rules recording
			// starting position of each
			positions := []int{0}
			for _, ruleName := range sequence {
				newPositions := []int{}
				for _, pos := range positions {
					subpos := match(message[pos:], rules[ruleName], rules)
					for _, newpos := range subpos {
						newPositions = append(newPositions, newpos+pos)
					}
				}
				if len(newPositions) == 0 {
					break
				}
				positions = newPositions
			}
			allPositions = append(allPositions, positions...)
		}

		if len(allPositions) == 0 {
			return []int{}
		}
		// Sort in descending order
		sort.Sort(sort.Reverse(sort.IntSlice(allPositions)))

		// Remove duplicates
		prev := allPositions[0]
		dedup := []int{prev}
		for i := 1; i < len(allPositions); i++ {
			v := allPositions[i]
			if v != prev {
				dedup = append(dedup, v)
				prev = v
			}
		}

		return dedup
	}

	return []int{}
}

type Rule struct {
	Value   byte
	Options [][]string
}

func parseInput(lines []string) (map[string]Rule, [][]byte) {
	rules := make(map[string]Rule)
	var messages [][]byte

	for _, line := range lines {
		if parts := strings.Split(line, ": "); len(parts) == 2 {
			if strings.Contains(parts[1], "\"") {
				rules[parts[0]] = Rule{Value: []byte(parts[1])[1]}
			} else {
				cr := childRules(parts[1])
				rules[parts[0]] = Rule{Options: cr}
			}
		} else {
			message := []byte(line)
			if len(message) > 0 {
				messages = append(messages, message)
			}
		}
	}
	return rules, messages
}

func childRules(s string) [][]string {
	sb := strings.Split(s, " | ")

	var res [][]string

	for _, sc := range sb {
		res = append(res, strings.Split(sc, " "))
	}
	return res
}
