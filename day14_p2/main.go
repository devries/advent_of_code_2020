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
	memory := make(map[uint64]uint64)
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

		ptr, err := strconv.ParseUint(match[1], 10, 64)
		utils.Check(err, fmt.Sprintf("error converting %s to int", match[1]))
		val, err := strconv.ParseUint(argument, 10, 64)
		utils.Check(err, fmt.Sprintf("error converting %s to uint64", argument))

		for _, subptr := range mask.Apply(ptr) {
			memory[subptr] = val
		}
	}

	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

type PortMask struct {
	mask         uint64
	bitSet       uint64
	combinations uint64
}

func NewPortMask(mask string) PortMask {
	var m uint64
	var b uint64
	var c uint64

	for _, v := range []rune(mask) {
		m <<= 1
		b <<= 1
		c <<= 1
		switch v {
		case 'X':
			c |= 1
		case '1':
			b |= 1
		case '0':
			m |= 1
		}
	}

	return PortMask{m, b, c}
}

func (p PortMask) Apply(x uint64) []uint64 {
	result := []uint64{}
	subv := x & p.mask
	subv |= p.bitSet

	result = append(result, subv)

	changables := p.getChangables()

	for i := 1; i <= len(changables); i++ {
		for cbits := range combinations(i, changables) {
			var sum uint64
			for _, c := range cbits {
				sum += c
			}
			r := subv | sum
			result = append(result, r)
		}
	}

	return result
}

func (p PortMask) getChangables() []uint64 {
	result := []uint64{}

	for i := 0; i < 36; i++ {
		var m uint64 = 1 << i

		v := m & p.combinations
		if v > 0 {
			result = append(result, v)
		}
	}

	return result
}

func combinations(n int, v []uint64) <-chan []uint64 {
	ch := make(chan []uint64)
	go func() {
		prefix := []uint64{}
		combinationsRecursor(prefix, n, v, ch)
		close(ch)
	}()

	return ch
}

func combinationsRecursor(prefix []uint64, n int, v []uint64, ch chan<- []uint64) {
	if n == 1 {
		for _, c := range v {
			l := len(prefix)
			result := make([]uint64, l+1)
			copy(result, prefix)
			result[l] = c
			ch <- result
		}
		return
	}

	if n == len(v) {
		result := make([]uint64, len(prefix)+len(v))
		copy(result, prefix)
		copy(result[len(prefix):], v)
		ch <- result
		return
	}

	l := len(prefix)
	newPrefix := make([]uint64, l+1)

	copy(newPrefix, prefix)
	newPrefix[l] = v[0]

	combinationsRecursor(newPrefix, n-1, v[1:], ch)
	combinationsRecursor(prefix, n, v[1:], ch)
}
