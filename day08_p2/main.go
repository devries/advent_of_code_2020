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

	instructions := parseInput(f)
	a, err := mutateAndRun(instructions)
	utils.Check(err, "did not find a successful solution")
	fmt.Println(a)
}

type Instruction struct {
	Op   string
	Arg  int
	Seen bool
}

func parseInput(r io.Reader) []Instruction {
	lines := utils.ReadLines(r)
	result := []Instruction{}

	for _, line := range lines {
		fields := strings.Fields(line)
		arg, err := strconv.Atoi(fields[1])
		utils.Check(err, "error reading argument")
		ins := Instruction{fields[0], arg, false}
		result = append(result, ins)
	}

	return result
}

type Computer struct {
	ptr  int           // pointer
	acc  int           // accumulator
	inst []Instruction // program
}

func NewComputer(instructions []Instruction) *Computer {
	c := Computer{0, 0, instructions}
	return &c
}

func (c *Computer) step() {
	c.inst[c.ptr].Seen = true
	ci := c.inst[c.ptr]
	switch ci.Op {
	case "nop":
		c.ptr++
	case "acc":
		c.acc += ci.Arg
		c.ptr++
	case "jmp":
		c.ptr += ci.Arg
	}
}

func (c *Computer) nextSeen() bool {
	if c.ptr >= len(c.inst) {
		return true
	}
	ci := c.inst[c.ptr]
	return ci.Seen
}

// Run instructions until one is seen again. Return boolean for terminating at
// end and accumulator value.
func runUntilRepeat(instructions []Instruction) (bool, int) {
	c := NewComputer(instructions)
	for c.nextSeen() == false {
		c.step()
	}

	success := false
	if c.ptr == len(c.inst) {
		success = true
	}

	return success, c.acc
}

// Mutate instructions and run until success
func mutateAndRun(instructions []Instruction) (int, error) {
	for i := 0; i < len(instructions); i++ {
		if instructions[i].Op != "acc" {
			test := make([]Instruction, len(instructions))
			copy(test, instructions)

			if test[i].Op == "jmp" {
				test[i].Op = "nop"
			} else {
				test[i].Op = "jmp"
			}

			success, acc := runUntilRepeat(test)
			if success {
				return acc, nil
			}
		}
	}
	return 0, fmt.Errorf("Did not find a solution")
}
