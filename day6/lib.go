package day6

import (
	"fmt"
	"strconv"
	"strings"
)

const WHITESPACE uint8 = 32

func SumCephalodHomeWork(homework string) int {
	lines := strings.Split(homework, "\n")
	ops, lines := strings.Fields(lines[len(lines)-1]), lines[:len(lines)-1]	
	var inEquation bool 

	for col := range len(lines[0]) {
		var digitsFound []byte

		for row := range len(lines) {
			b := lines[row][col]
			if b != WHITESPACE {
				digitsFound = append(digitsFound, b)
			}
		}

	}


}

func SumCephalodHomework(homework string) int {
	var total int

	equations := parseInput(homework)
	for _, equation := range equations {
		fmt.Printf("%v\n", equation)
		solution := equation.Solve()
		total += solution 
		fmt.Printf("%v\n",solution)
	}
	return total
}

func parseInput(homework string) []Equation {
	var equations []Equation
	var numberMatrix [][]int
	lines := strings.Split(homework, "\n")
	ops, lines := strings.Fields(lines[len(lines)-1]), lines[:len(lines)-1]	
	for _, line := range lines {
		row, err := parseLine(line)
		if err != nil {
			panic("")
		}
		numberMatrix = append(numberMatrix,row)
	}
	for col := range len(numberMatrix[0]) {
		var op Op
		switch ops[col] {
		case "+": 
			op = ADD
		case "*": 
			op = MUL
		default:
			panic("")
		}

		var numbers []int
		for row := range len(numberMatrix) {

			number := numberMatrix[row][col]
			numbers = append(numbers, number)
		}
		equation := Equation{op, numbers}
		equations = append(equations, equation)
	}
	return equations
} 


func parseLine(line string) ([]int, error) {
	var parsed []int
	fields := strings.Fields(line)
	for _, value := range fields {
		number, err := strconv.Atoi(value)
		if err != nil {
			return parsed, err
		}
		parsed = append(parsed, number)
	}
	return parsed, nil
}

type Op int

const (
	ADD Op = iota
	MUL
)

type Equation struct {
	op Op
	numbers []int
}

func (e Equation) Solve() int {
	var total int

	switch e.op {
	case ADD: 
		fmt.Println("Adding")
		for _, n := range e.numbers {
			total += n
		}
	case MUL:
		fmt.Println("Multiplying")
		total++
		for _, n := range e.numbers {
			total *= n
		}

	default:
		panic("nooo")
	}
	return total
}
