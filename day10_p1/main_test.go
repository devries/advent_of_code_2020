package main

import (
	"strings"
	"testing"
)

var testInput1 = `16
10
15
5
1
11
7
19
6
12
4`

var testInput2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestProblem(t *testing.T) {
	tests := []struct {
		input string
		total int
	}{
		{testInput1, 7 * 5},
		{testInput2, 22 * 10},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		vals := parseInput(r)

		result := solve(vals)
		if result != test.total {
			t.Errorf("Expected %d, got %d", test.total, result)
		}
	}
}
