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
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		boardingPass := scanner.Text()
		id, err := getSeatId(boardingPass)
		utils.Check(err, "error parsing boarding pass")
		if id > max {
			max = id
		}
	}
	err = scanner.Err()
	utils.Check(err, "error reading input.txt")

	fmt.Println(max)
}

func getSeatId(boardingPass string) (int, error) {
	pass := []rune(boardingPass)

	if len(pass) != 10 {
		return 0, fmt.Errorf("the boarding pass %s is not the proper length", boardingPass)
	}

	rowCode := pass[:7]
	seatCode := pass[7:]

	row := 0

	for _, code := range rowCode {
		row <<= 1
		switch code {
		case 'B':
			row |= 1
		case 'F':
			row |= 0
		default:
			return 0, fmt.Errorf("unexpected row code %c found in %s", code, boardingPass)
		}
	}

	seat := 0

	for _, code := range seatCode {
		seat <<= 1
		switch code {
		case 'R':
			seat |= 1
		case 'L':
			seat |= 0
		default:
			return 0, fmt.Errorf("unexpected seat code %c found in %s", code, boardingPass)
		}
	}

	return row*8 + seat, nil
}
