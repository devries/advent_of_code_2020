package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	max := 0
	occupied := make(map[int]bool)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		boardingPass := scanner.Text()
		id, err := getSeatId(boardingPass)
		utils.Check(err, "error parsing boarding pass")
		occupied[id] = true
		if id > max {
			max = id
		}
	}
	err = scanner.Err()
	utils.Check(err, "error reading input.txt")

	for i := 1; i < max; i++ {
		if occupied[i] == false && occupied[i-1] == true && occupied[i+1] == true {
			fmt.Println(i)
		}
	}
}

func getSeatId(boardingPass string) (int, error) {
	pass := []rune(boardingPass)

	if len(pass) != 10 {
		return 0, fmt.Errorf("the boarding pass %s is not the proper length", boardingPass)
	}

	// Originally I calculated the row and column separately, but that is not necessary
	id := 0

	for _, code := range pass {
		id <<= 1
		switch code {
		case 'B', 'R':
			id |= 1
		case 'F', 'L':
			id |= 0
		default:
			return 0, fmt.Errorf("unexpected character %c found in %s", code, boardingPass)
		}
	}
	return id, nil
}
