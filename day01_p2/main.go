package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	var numbers []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ntxt := scanner.Text()
		n, err := strconv.Atoi(ntxt)
		utils.Check(err, fmt.Sprintf("unable to convert %s to integer", ntxt))
		numbers = append(numbers, n)
	}
	err = scanner.Err()
	utils.Check(err, "error reading input.txt")

	out := findSum(numbers, 2020)
	ans := out[0] * out[1] * out[2]
	fmt.Println(ans)
}

func findSum(numbers []int, total int) []int {
	l := len(numbers)

	for i := 0; i < l-2; i++ {
		a := numbers[i]
		for j := i + 1; j < l-1; j++ {
			b := numbers[j]
			for k := j + 1; k < l; k++ {
				c := numbers[k]

				if a+b+c == total {
					return []int{a, b, c}
				}
			}
		}
	}

	return []int{}
}
