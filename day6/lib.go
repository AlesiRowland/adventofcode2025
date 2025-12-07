package day6

import (
	"strconv"
	"strings"
)

const WHITESPACE uint8 = 32


func SolvePart1(input string) int{
	var total int
	equations := parseInput(input)
	for _, equation := range equations {
		solution := equation.Solve()
		total += solution 
	}
	return total
}

func SolvePart2(input string) int {
	// Step 1, parse input into [][]byte

	var matrix [][]byte 	
	lines := strings.Split(input, "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	ops := strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}


	nCols := len(matrix[0])
	nRows := len(matrix)
	
	var transformed []string
	for col := range nCols {
		var values []byte
		for row := range nRows {
			value := matrix[row][col]
			if value != WHITESPACE {
				values = append(values, matrix[row][col])
			}
		}

		transformed = append(transformed, string(values))
	}
	var total int	
	var index int
	for _, op := range(ops) {
		var current int
		if op == "*" {
			current++
		}

		for index < len(transformed) && transformed[index] != "" {
			num, err := strconv.Atoi(transformed[index])
			if err != nil {
				panic("")
			}

			switch op {
			case "+":
				current += num	
			case "*":
				current *= num 
			default:
				panic("")
			}
			index++
		}
		total += current
		index++
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
		for _, n := range e.numbers {
			total += n
		}
	case MUL:
		total++
		for _, n := range e.numbers {
			total *= n
		}

	default:
		panic("nooo")
	}
	return total
}
