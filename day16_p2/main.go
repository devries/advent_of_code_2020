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

func solve(r io.Reader) int64 {
	rsmap, myTicket, nearby := parseInput(r)

	allRanges := []Range{}

	for _, rs := range rsmap {
		for _, r := range rs {
			allRanges = append(allRanges, r)
		}
	}
	allRanges = consolidateRanges(allRanges)

	goodTickets := cullInvalid(allRanges, nearby)

	// positions := make(map[string]int)

	// This section looks for possible fields and makes a set of them in
	// possibilities
	possibilities := make([]map[string]bool, len(myTicket))
	for i := range possibilities {
		possibilities[i] = make(map[string]bool)
	}

	for k, v := range rsmap {
		for i := range myTicket {
			valid := true
			for _, tick := range goodTickets {
				if !rangesContain(v, tick[i]) {
					valid = false
					break
				}
			}

			if valid {
				possibilities[i][k] = true
			}
		}
	}

	corresponds := eliminate(possibilities)

	var result int64 = 1
	for i, v := range corresponds {
		if strings.HasPrefix(v, "departure") {
			result *= int64(myTicket[i])
		}
	}

	return result
}

func eliminate(p []map[string]bool) []string {
	done := make([]bool, len(p))
	finals := make([]string, len(p))

	for {
		stillEliminating := false
		for i, isDone := range done {
			if isDone {
				continue
			}
			if len(p[i]) == 1 {
				// We are going to eliminate the index with one possible solution
				stillEliminating = true
				// Get the only key
				for k, _ := range p[i] {
					finals[i] = k
				}
				// Delete key from all other possibilities
				for j, otherDone := range done {
					if j != i && otherDone == false {
						delete(p[j], finals[i])
					}
				}
				// Mark this index as done
				done[i] = true
			}
		}
		if !stillEliminating {
			utils.Check(fmt.Errorf("Unable to eliminate any possibilities"), "error in elimination")
		}

		if loopComplete := containsBool(done, false); !loopComplete {
			break
		}
	}

	return finals
}

func containsBool(s []bool, v bool) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}

	return false
}

func cullInvalid(rs []Range, tickets [][]int) [][]int {
	goodTickets := [][]int{}
	for _, tick := range tickets {
		good := true
		for _, v := range tick {
			if !rangesContain(rs, v) {
				good = false
				break
			}
		}
		if good {
			goodTickets = append(goodTickets, tick)
		}
	}

	return goodTickets
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
