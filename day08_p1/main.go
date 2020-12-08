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
	a := runUntilRepeat(instructions)
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
	ci := c.inst[c.ptr]
	return ci.Seen
}

// Run instructions until one is seen again. Return accumulator value.
func runUntilRepeat(instructions []Instruction) int {
	c := NewComputer(instructions)
	for c.nextSeen() == false {
		c.step()
	}
	return c.acc
}
