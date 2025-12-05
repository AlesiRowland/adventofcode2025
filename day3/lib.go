package day3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetTotalOutputJoltage(input string, n int) int {
	var total int

	for _, bank := range parseInput(input) {
		joltage, err := getJoltage(bank, n)
		if err != nil {
			panic("Invalid input")
		}
		total += joltage
	}
	return total
}

func getJoltage(bank []byte, n int) (int, error) {
	if len(bank) < n {
		return 0, errors.New("Bank is smaller than n")
	}

	var bytes []byte
	var index int
	digitsLeft := n

	for range digitsLeft {
		var b byte
		for i := index; i < len(bank)-(digitsLeft-1); i++ {
			if bank[i] > b {
				b = bank[i]
				index = i + 1
			}
		}
		bytes = append(bytes, b)
		digitsLeft -= 1
	}

	joltage, err := strconv.Atoi(string(bytes))
	return joltage, err

}

// func getJoltage(bank []byte, n int) (int, error) {
// 	// Find largest number not at the end.
// 	var firstByte byte
// 	var firstByteIndex int
// 	var secondByte byte
//
//
//
// 	for i, digit := range bank[:len(bank)-1] {
// 		if digit > firstByte {
// 			firstByte = digit
// 			firstByteIndex = i
// 		}
// 	}
// 	// Find the large number between that number and the end
// 	for _, digit := range bank[firstByteIndex + 1:] {
// 		if digit > secondByte {
// 			secondByte = digit
// 		}
// 	}
// 	joltage, err := strconv.Atoi(string([]byte{firstByte, secondByte}))
// 	return joltage, err
// }

func parseInput(input string) [][]byte {
	var banks [][]byte
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		bank := []byte(line)
		banks = append(banks, bank)
	}
	return banks
}
