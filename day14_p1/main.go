package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	v := solve(f)
	fmt.Println(v)
}

func solve(r io.Reader) uint64 {
	memory := make(map[int]uint64)
	lines := utils.ReadLines(r)
	pattern := regexp.MustCompile(`mem\[(\d+)\]`)
	var mask PortMask

	for _, line := range lines {
		parts := strings.Split(line, "=")
		inst := strings.TrimSpace(parts[0])
		argument := strings.TrimSpace(parts[1])

		if inst == "mask" {
			mask = NewPortMask(argument)
			continue
		}

		match := pattern.FindStringSubmatch(inst)
		if len(match) == 0 {
			utils.Check(fmt.Errorf("%s not parsed", inst), "error parsing instruction")
		}

		ptr, err := strconv.Atoi(match[1])
		utils.Check(err, fmt.Sprintf("error converting %s to int", match[1]))
		val, err := strconv.ParseUint(argument, 10, 64)
		utils.Check(err, fmt.Sprintf("error converting %s to uint64", argument))

		memory[ptr] = mask.Apply(val)
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

type PortMask struct {
	mask   uint64
	bitSet uint64
}

func NewPortMask(mask string) PortMask {
	var m uint64
	var b uint64

	for _, v := range []rune(mask) {
		m <<= 1
		b <<= 1
		switch v {
		case 'X':
			m |= 1
		case '1':
			b |= 1
		case '0':
			m |= 0
		}
	}

	return PortMask{m, b}
}

func (p PortMask) Apply(x uint64) uint64 {
	result := x & p.mask
	result |= p.bitSet

	return result
}
