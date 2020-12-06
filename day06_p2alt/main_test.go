package main

import (
	"strings"
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
	r := strings.NewReader(testInput)
	groups, err := parseCustoms(r)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	sum := sumCounts(groups)

	if sum != 6 {
		t.Errorf("Got sum of %d, expected 11", sum)
	}
}
