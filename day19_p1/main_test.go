package main

import (
	"strings"
	"testing"
)

var testInput = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

func TestIt(t *testing.T) {
	results := []bool{true, false, true, false, false}

	lines := strings.Split(testInput, "\n")

	rules, messages := parseInput(lines)

	for i, message := range messages {
		r := solve(message, rules)

		if r != results[i] {
			t.Errorf("For %s got %t, expected %t", string(message), r, results[i])
		}
	}
}
