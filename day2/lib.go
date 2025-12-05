package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func GetInvalidProductIds(input []byte) int {
	var count int
	productIDRanges := parseInput(input)
	for _, record := range productIDRanges {
		count += getTotalOfInvalidIDsInRange(record[0], record[1])
	}
	return count
}

func GetInvalidProductIds2(input []byte) int {
	var count int
	productIDRanges := parseInput(input)
	for _, record := range productIDRanges {
		count += getTotalOfInvalidIDsInRange2(record[0], record[1])
	}
	return count
}

func getTotalOfInvalidIDsInRange(start int, end int) int {
	fmt.Printf("Start: %v, End: %v\n", start, end)
	var total int

	startChars := strconv.Itoa(start)
	length := len(startChars)

	endIndex := length / 2

	base, err := strconv.Atoi(startChars[0:endIndex])
	if err != nil {
		base = 0
	}
	candidate, err := doubleInteger(base)
	if err != nil {
		panic("invalid number")
	}
	for candidate < start {
		base += 1
		candidate, err = doubleInteger(base)
		if err != nil {
			panic("Invalid integer")
		}
	}
	for candidate <= end {
		total += candidate
		base += 1
		candidate, err = doubleInteger(base)
		if err != nil {
			panic("Invalid integer")
		}
	}
	fmt.Printf("Count: %v\n", total)
	return total
}

func getTotalOfInvalidIDsInRange2(start int, end int) int {
	fmt.Printf("Start: %v, End: %v\n", start, end)
	var total int
	visited := make(map[int]bool)
	base := 1
	repeat := 2
	doubled := repeatingInteger(base, repeat)
	for doubled <= end {
		candidate := doubled
		for candidate < start {
			repeat += 1
			candidate = repeatingInteger(base, repeat)
		}
		for candidate <= end {
			if _, exists := visited[candidate]; !exists {
				total += candidate
				visited[candidate] = true
			}
			fmt.Printf("%v\n", candidate)
			repeat += 1
			candidate = repeatingInteger(base, repeat)
		}
		// reset
		base += 1
		repeat = 2
		doubled = repeatingInteger(base, repeat)
	}
	return total

}

func doubleInteger(integer int) (int, error) {
	asString := strconv.Itoa(integer)
	doubled, err := strconv.Atoi(asString + asString)
	return doubled, err
}

func repeatingInteger(integer int, n int) int {
	asString := strconv.Itoa(integer)
	doubled, err := strconv.Atoi(strings.Repeat(asString, n))
	if err != nil {
		panic("should not be possible")
	}
	return doubled
}

func parseInput(input []byte) [][2]int {
	var ranges [][2]int

	lines := strings.Split(string(input), ",")

	for _, line := range lines {
		before, after, found := strings.Cut(line, "-")
		if !found {
			panic("Invalid Range")
		}

		start, err := strconv.Atoi(before)
		if err != nil {
			panic(fmt.Sprintf("first value not an integer: %v", before))
		}

		end, err := strconv.Atoi(after)
		if err != nil {
			panic("second value not an integer")
		}

		identifierRange := [2]int{start, end}
		ranges = append(ranges, identifierRange)
	}

	return ranges
}
