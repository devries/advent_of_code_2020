package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")

	inputBytes, err := ioutil.ReadAll(f)
	utils.Check(err, "error reading input")
	input := string(inputBytes)

	passports := splitPassports(input)
	validCount := 0

	for _, p := range passports {
		v := validatePassport(p)
		if v {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func splitPassports(input string) []string {
	pattern := regexp.MustCompile(`(?m)^\s*$`)
	passports := pattern.Split(input, -1)

	return passports
}

func getAttributes(passport string) map[string]string {
	result := make(map[string]string)
	pattern := regexp.MustCompile(`([^\s:]+):([^\s:]+)`)

	substrings := pattern.FindAllStringSubmatch(passport, -1)
	for _, pair := range substrings {
		result[pair[1]] = pair[2]
	}

	return result
}

func validatePassport(passport string) bool {
	att := getAttributes(passport)

	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, a := range required {
		if _, ok := att[a]; !ok {
			return false
		}
	}

	return true
}
