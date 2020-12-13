package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	busses := parseInput(f)

	result := solve(busses)
	fmt.Println(result)
}

func solve(busses map[int64]int64) int64 {
	var skips int64
	var offset int64

	keys := make([]int64, 0, len(busses))
	for k := range busses {
		keys = append(keys, k)
	}

	// Not sure I need to sort, but I had a bug and thought this might help.
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	// For any group, the pattern repeats every LCM of the elements involved.
	// The offset is the offset from that LCM at which the pattern starts.
	// That offset becomes the new remainder for the next step.
	for _, k := range keys {
		v := busses[k]
		if skips == 0 {
			skips = k
			offset = v
			continue
		}
		lcm := utils.Lcm(skips, k)
		for i := offset % lcm; i <= lcm; i += skips {
			if i%k == v {
				skips = lcm
				offset = i
				break
			}
		}
	}
	return offset
}

func parseInput(r io.Reader) map[int64]int64 {
	lines := utils.ReadLines(r)

	return parseBusses(lines[1])
}

func parseBusses(schedule string) map[int64]int64 {
	busEntries := strings.Split(schedule, ",")

	busses := make(map[int64]int64)

	for i, bus := range busEntries {
		if bus == "x" {
			continue
		}
		val, err := strconv.ParseInt(bus, 10, 64)
		utils.Check(err, "Error parsing bus id")
		if int64(i)%val == 0 {
			busses[val] = 0
		} else {
			busses[val] = val - int64(i)%val
		}
	}

	return busses
}
