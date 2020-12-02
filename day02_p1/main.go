package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input.txt")
	defer f.Close()

	validPasswords := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ntxt := scanner.Text()
		valid, err := validPassword(ntxt)
		utils.Check(err, fmt.Sprintf("Error validating entry: %s", ntxt))
		if valid {
			validPasswords++
		}
	}
	err = scanner.Err()
	utils.Check(err, "error reading input.txt")

	fmt.Printf("Answer: %d\n", validPasswords)
}

func validPassword(entry string) (bool, error) {
	parts := strings.Split(entry, ":")
	if len(parts) != 2 {
		return false, fmt.Errorf("Unable to split %s into two parts", entry)
	}

	var min, max int
	var coi rune // character of interest

	_, err := fmt.Sscanf(parts[0], "%d-%d %c", &min, &max, &coi)
	if err != nil {
		return false, err
	}

	password := strings.TrimSpace(parts[1])

	sum := 0
	for _, c := range password {
		if c == coi {
			sum++
		}
	}
	if sum >= min && sum <= max {
		return true, nil
	} else {
		return false, nil
	}
}
