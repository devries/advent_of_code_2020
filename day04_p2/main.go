package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

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

	if !checkYear(att["byr"], 1920, 2002) {
		return false
	}
	if !checkYear(att["iyr"], 2010, 2020) {
		return false
	}
	if !checkYear(att["eyr"], 2020, 2030) {
		return false
	}
	if !checkHeight(att["hgt"]) {
		return false
	}
	if !checkHairColor(att["hcl"]) {
		return false
	}
	if !checkEyeColor(att["ecl"]) {
		return false
	}
	if !checkPassportId(att["pid"]) {
		return false
	}

	return true
}

func checkYear(input string, min int, max int) bool {
	v, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	if v >= min && v <= max {
		return true
	}
	return false
}

func checkHeight(input string) bool {
	if len(input) < 4 {
		return false
	}
	ending := input[len(input)-2:]
	numeric := input[:len(input)-2]
	number, err := strconv.Atoi(numeric)
	if err != nil {
		return false
	}

	switch ending {
	case "cm":
		if number >= 150 && number <= 193 {
			return true
		} else {
			return false
		}
	case "in":
		if number >= 59 && number <= 76 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func checkHairColor(input string) bool {
	pattern := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return pattern.MatchString(input)
}

func checkEyeColor(input string) bool {
	values := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	return values[input]
}

func checkPassportId(input string) bool {
	pattern := regexp.MustCompile(`^[0-9]{9}$`)
	return pattern.MatchString(input)
}
