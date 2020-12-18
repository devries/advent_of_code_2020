package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
)

type TokenType int

const (
	NUMBER TokenType = iota
	MULOPER
	ADDOPER
	LPARENS
	RPARENS
)

type Token struct {
	token TokenType
	value int
}

type Statement struct {
	tokens   []Token
	position int
}

func main() {
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening input")
	defer f.Close()

	lines := utils.ReadLines(f)

	var sum int64
	for _, line := range lines {
		stmt := tokenize(line)
		sum += int64(evalExpression(&stmt))
	}

	fmt.Println(sum)
}

func tokenize(input string) Statement {
	toks := []Token{}
	var b strings.Builder

	for _, c := range input {
		if c >= '0' && c <= '9' {
			// Number
			b.WriteRune(c)
		} else {
			if b.Len() > 0 {
				v, err := strconv.Atoi(b.String())
				if err != nil {
					panic(fmt.Errorf("Unable to parse numbers in %s", input))
				}
				toks = append(toks, Token{NUMBER, v})
				b.Reset()
			}
			switch c {
			case '+':
				toks = append(toks, Token{ADDOPER, 0})
			case '*':
				toks = append(toks, Token{MULOPER, 0})
			case '(':
				toks = append(toks, Token{LPARENS, 0})
			case ')':
				toks = append(toks, Token{RPARENS, 0})
			}
		}
	}
	if b.Len() > 0 {
		v, err := strconv.Atoi(b.String())
		if err != nil {
			panic(fmt.Errorf("Unable to parse numbers in %s", input))
		}
		toks = append(toks, Token{NUMBER, v})
	}

	return Statement{toks, 0}
}

func evalExpression(stmt *Statement) int {
	left := evalTerm(stmt)

	for stmt.position < len(stmt.tokens) {
		current := stmt.tokens[stmt.position]
		switch current.token {
		case ADDOPER:
			stmt.position++
			right := evalTerm(stmt)
			left = left + right
		case MULOPER:
			stmt.position++
			right := evalTerm(stmt)
			left = left * right
		default:
			return left
		}
	}

	return left
}

func evalTerm(stmt *Statement) int {
	current := stmt.tokens[stmt.position]

	if current.token == NUMBER {
		stmt.position++
		return current.value
	}

	if current.token == LPARENS {
		stmt.position++
		v := evalExpression(stmt)
		current = stmt.tokens[stmt.position]
		if current.token != RPARENS {
			panic(fmt.Errorf("Unbalanced parenthesis"))
		}
		stmt.position++
		return v
	}

	panic(fmt.Errorf("No term found"))
}
