package main

import (
	"strings"
	"testing"
)

var testInput = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestLineParse(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"nwwswee", []string{"nw", "w", "sw", "e", "e"}},
		{"eneenew", []string{"e", "ne", "e", "ne", "w"}},
	}

	for _, test := range tests {
		r := parseLine(test.input)

		if len(r) != len(test.expected) {
			t.Errorf("For %s, expected %v, got %v", test.input, test.expected, r)
		}
		for i, p := range r {
			if p != test.expected[i] {
				t.Errorf("For %s, expected %v, got %v", test.input, test.expected, r)
			}
		}
	}
}

func TestAll(t *testing.T) {
	r := strings.NewReader(testInput)

	result := solve(r)
	expected := 10

	if result != expected {
		t.Errorf("For test input expected %d, got %d", expected, result)
	}
}
