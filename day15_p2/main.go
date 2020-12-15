package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	input := "7,12,1,0,16,2"
	result := solve(input)
	fmt.Println(result)
}

func solve(input string) int {
	parts := strings.Split(input, ",")

	sequence := []int{}
	for _, v := range parts {
		intv, err := strconv.Atoi(v)
		utils.Check(err, "error converting to integer")
		sequence = append(sequence, intv)
	}

	lastUse := make(map[int]int)
	previous := 0
	current := 0

	for i := 1; i <= 30000000; i++ {
		if i <= len(sequence) {
			current = sequence[i-1]
		} else {
			seen, ok := lastUse[previous]
			if ok {
				current = i - seen - 1
			} else {
				current = 0
			}
		}
		if i > 1 {
			lastUse[previous] = i - 1
		}
		previous = current
		// fmt.Printf("%d: %d\n", i, previous)
	}

	return previous
}
