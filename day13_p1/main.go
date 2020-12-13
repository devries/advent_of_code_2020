package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	timestamp, busses := parseInput(f)

	result := solve(timestamp, busses)
	fmt.Println(result)
}

func solve(timestamp int64, busses []int64) int64 {
	var minInterval int64 = 1<<63 - 1
	var minBusId int64

	for _, busId := range busses {
		lastTime := timestamp % busId
		nextTime := busId - lastTime
		if nextTime < minInterval {
			minInterval = nextTime
			minBusId = busId
		}
	}

	return minBusId * minInterval
}

func parseInput(r io.Reader) (int64, []int64) {
	lines := utils.ReadLines(r)

	timestamp, err := strconv.ParseInt(lines[0], 10, 64)
	utils.Check(err, "error parsing first line")

	busEntries := strings.Split(lines[1], ",")

	busses := []int64{}

	for _, bus := range busEntries {
		if bus == "x" {
			continue
		}
		val, err := strconv.ParseInt(bus, 10, 64)
		utils.Check(err, "Error parsing bus id")
		busses = append(busses, val)
	}

	return timestamp, busses
}
