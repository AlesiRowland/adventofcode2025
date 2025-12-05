package day5

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)
func CountFreshIngredients(input string ) int{
	var freshIds []int
	var rangeIndex, idIndex int

	ranges, ids := parseInput(input)
	nIds := len(ids)
	nRanges := len(ranges)

	for idIndex < nIds && rangeIndex < nRanges {
		var isFresh bool
		id := ids[idIndex] // The ID we shall interrogate

		for rangeIndex < nRanges {
			if id < ranges[rangeIndex].start {
				break
			}

			if id > ranges[rangeIndex].end {
				rangeIndex += 1
				continue
			}
			isFresh = true
			break
		}

		if isFresh {
			freshIds = append(freshIds, id)
		}
		idIndex++
	}
	fmt.Printf("%v\n", freshIds)
	return len(freshIds)
}

func CountFreshIngredientsInRanges(input string) int {
	parts := strings.Split(input, "\n\n")
	ranges := parseRanges[parts[0]
	
	// Because we have ranges that are clearly going to be fucking massive, we can't use the naive solution of looping through all the numbers.
	// Ideally, we need to:
	// Merge ranges together
	// then sum their end-start diffs

	
}

func parseInput(input string) ([]Range, []int) {

	parts := strings.Split(input, "\n\n")
	ranges := parseRanges(parts[0])

	ids := parseIds(parts[1])
	return ranges, ids
}

func parseRanges(text string) []Range {
	var ranges []Range
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		rangeValues := strings.SplitN(line, "-", 2)	
		start, err := strconv.Atoi(rangeValues[0])
		if err != nil {
			panic("start couldn't be converted to int")
		}
	 	end, err := strconv.Atoi(rangeValues[1])
		if err != nil {
			panic("end couldn't be converted to int")
		}
		r := NewRange(start, end)
		ranges = append(ranges, r)
	}
	slices.SortFunc(ranges, func(a, b Range) int {
	 	return cmp.Compare(a.start, b.start)	
	})
	return ranges
}

func mergeRanges(ranges[]Range) []Range) {
	// assume the ranges are already start 
}

func parseIds(text string) []int {
	var ids []int
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		id, err := strconv.Atoi(line)
		if err != nil {
			panic("Could not convert to identifier")
		}
		ids = append(ids, id)
	}
	slices.Sort(ids)
	return ids
}

type Range struct {
	start int
	end int
}

func NewRange(start, end int) Range {
	return Range {
		start,
		end,
	}
}
