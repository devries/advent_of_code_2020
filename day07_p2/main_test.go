package main

import (
	"strings"
	"testing"
)

var testInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestProblem(t *testing.T) {
	r := strings.NewReader(testInput)
	rules := parseInput(r)
	n := recurseQuantity("shiny gold", rules)
	if n-1 != 32 {
		t.Errorf("Expected 32 bags, got %d", n-1)
	}
}
