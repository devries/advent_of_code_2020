package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	fmt.Println(solve(f))
}
func solve(r io.Reader) int {
	rsmap, _, nearby := parseInput(r)

	allRanges := []Range{}

	for _, rs := range rsmap {
		for _, r := range rs {
			allRanges = append(allRanges, r)
		}
	}
	allRanges = consolidateRanges(allRanges)

	invalidSum := 0

	for _, tick := range nearby {
		for _, v := range tick {
			if !rangesContain(allRanges, v) {
				invalidSum += v
			}
		}
	}
	return invalidSum
}

type Range struct {
	Min int
	Max int
}

func consolidateRanges(rs []Range) []Range {
	result := []Range{}

	sort.Slice(rs, func(i, j int) bool { return rs[i].Min < rs[j].Min })

	var current Range

	for i, r := range rs {
		if i == 0 {
			current = r
			continue
		}

		if r.Min <= current.Max {
			if r.Max > current.Max {
				current.Max = r.Max
			}
		} else {
			result = append(result, current)
			current = r
		}
	}

	result = append(result, current)

	return result
}

// Is value in ranges
func rangesContain(rs []Range, val int) bool {
	max := len(rs)
	min := 0

	for {
		var i int
		if (max - min) > 1 {
			i = min + ((max - min) / 2)
		} else {
			i = min
		}

		if val >= rs[i].Min {
			if val <= rs[i].Max {
				return true
			}
			min = i + 1
		} else {
			max = i
		}
		if max <= min {
			return false
		}
	}
}

func parseInput(r io.Reader) (map[string][]Range, []int, [][]int) {
	fulltext, err := ioutil.ReadAll(r)
	utils.Check(err, "error reading input")

	rsMap := make(map[string][]Range)
	yourTicket := []int{}
	nearbyTickets := [][]int{}

	parts := strings.Split(string(fulltext), "\n\n")

	// Parse ranges
	pattern := regexp.MustCompile(`(?m)^([^:]+):\s+([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)

	for _, v := range pattern.FindAllStringSubmatch(parts[0], -1) {
		ramin, err := strconv.Atoi(v[2])
		utils.Check(err, "unable to parse integer")
		ramax, err := strconv.Atoi(v[3])
		utils.Check(err, "unable to parse integer")
		rbmin, err := strconv.Atoi(v[4])
		utils.Check(err, "unable to parse integer")
		rbmax, err := strconv.Atoi(v[5])
		utils.Check(err, "unable to parse integer")

		rsMap[v[1]] = []Range{Range{ramin, ramax}, Range{rbmin, rbmax}}
	}

	// fmt.Printf("%+v\n", rsMap)

	// Parse my ticket
	subparts := strings.Split(parts[1], "\n")
	for _, number := range strings.Split(subparts[1], ",") {
		val, err := strconv.Atoi(number)
		utils.Check(err, "unable to parse integer in ticket")
		yourTicket = append(yourTicket, val)
	}

	// Parse other tickets
	subparts = strings.Split(strings.TrimSpace(parts[2]), "\n")
	for _, line := range subparts[1:] {
		thisTicket := []int{}
		for _, number := range strings.Split(line, ",") {
			val, err := strconv.Atoi(number)
			utils.Check(err, "unable to parse integer in ticket")
			thisTicket = append(thisTicket, val)
		}
		nearbyTickets = append(nearbyTickets, thisTicket)
	}

	return rsMap, yourTicket, nearbyTickets
}
