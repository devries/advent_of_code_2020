package main

import (
	"testing"
)

var testInput = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestCount(t *testing.T) {
	groups := splitBlankLines(testInput)
	sum := sumCounts(groups)

	if sum != 11 {
		t.Errorf("Got sum of %d, expected 11", sum)
	}
}
