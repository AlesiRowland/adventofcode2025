package day1

import (
	"strconv"
	"strings"
)

type Direction int

const (
	LEFT = iota
	RIGHT
)

func GetSafeCode(start int, input []byte) int {
	var total int
	instructions := parseInput(input)
	dial := NewDial(start, 100)
	for _, instruction := range instructions {
		dial.Turn(instruction)
		if dial.GetValue() == 0 {
			total += 1
		}
	}
	return total

}

func GetNumberOfClicks(start int, input []byte) int {
	var total int
	instructions := parseInput(input)
	dial := NewDial(start, 100)
	for _, instruction := range instructions {
		total += dial.Turn2(instruction)
	}
	return total

}

func parseInput(input []byte) []Instruction {
	var instructions []Instruction
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		var rotation Direction
		switch line[0] {
		case 'R':
			rotation = RIGHT
		case 'L':
			rotation = LEFT
		default:
			panic("Invalid rotation")
		}

		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("Invalid amount")
		}
		instruction := Instruction{
			rotation,
			amount,
		}
		instructions = append(instructions, instruction)
	}
	return instructions
}

type Instruction struct {
	Rotation Direction
	Amount   int
}

type Dial struct {
	value int
	limit int
}

func NewDial(value int, limit int) Dial {
	return Dial{
		value: value,
		limit: limit,
	}
}

func (d *Dial) GetValue() int {
	return d.value
}

func (d *Dial) Turn(instruction Instruction) {
	switch instruction.Rotation {
	case LEFT:
		d.value -= instruction.Amount % d.limit
		if d.value < 0 {
			d.value = d.limit + d.value
		}
	case RIGHT:
		d.value = (d.value + instruction.Amount) % d.limit
	default:
		panic("Invalid direction")
	}
}

func (d *Dial) Turn2(instruction Instruction) int {
	limit := 100
	fullTurns := instruction.Amount / limit

	switch instruction.Rotation {
	case LEFT:

		if d.value != 0 && d.value-(instruction.Amount%limit) < 0 {
			fullTurns += 1
		}

		d.value -= instruction.Amount % d.limit
		if d.value < 0 {
			d.value = d.limit + d.value
		}
	case RIGHT:

		if d.value+(instruction.Amount%limit) > limit {
			fullTurns += 1
		}
		d.value = (d.value + instruction.Amount) % limit

	default:
		panic("Invalid direction")
	}

	if d.value == 0 {
		fullTurns += 1
	}
	return fullTurns
}
