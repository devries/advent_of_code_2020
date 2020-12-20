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
	r := matchLengths(message, "0", rules)
	if len(r) > 0 && r[0] == len(message) {
		return true
	}
	return false
}

func matchLengths(message []byte, ruleName string, rules map[string]Rule) []int {
	// Message is empty
	if len(message) == 0 {
		return []int{}
	}

	rule := rules[ruleName]

	// This rule is a byte, see if it matches
	if rule.Value != 0 {
		if message[0] == rule.Value {
			return []int{1}
		} else {
			return []int{}
		}
	}

	allPositions := []int{}
	for _, sequence := range rule.Options {
		positions := matchSequence(message, sequence, rules)
		allPositions = append(allPositions, positions...)
		allPositions = reverseSortAndDedup(allPositions)
	}

	return allPositions
}

func matchSequence(message []byte, sequence []string, rules map[string]Rule) []int {
	positions := []int{0}

	for _, ruleName := range sequence {
		newPositions := []int{}
		for _, pos := range positions {
			lengths := matchLengths(message[pos:], ruleName, rules)
			for _, length := range lengths {
				newPositions = append(newPositions, length+pos)
			}
		}
		if len(newPositions) == 0 {
			return []int{}
		}
		newPositions = reverseSortAndDedup(newPositions)
		positions = newPositions
	}

	return positions
}

func containsInt(s []int, v int) bool {
	for _, t := range s {
		if t == v {
			return true
		}
	}
	return false
}

func reverseSortAndDedup(s []int) []int {
	if len(s) == 0 {
		return s
	}
	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	// Remove duplicates
	prev := s[0]
	dedup := []int{prev}
	for i := 1; i < len(s); i++ {
		v := s[i]
		if v != prev {
			dedup = append(dedup, v)
			prev = v
		}
	}

	return s
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
