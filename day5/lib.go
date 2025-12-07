package day5

import (
	"cmp"
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
	return len(freshIds)
}

func CountFreshIngredientsInRanges(input string) int {
	var total int
	parts := strings.Split(input, "\n\n")
	ranges := parseRanges(parts[0])
	ranges = mergeRanges(ranges)
	for _, r := range ranges {
		total += r.end - r.start
		total++ // because inclusive
	}
	return total
}

func mergeRanges(ranges []Range) []Range {
	var merged []Range
	if len(ranges) == 0 {
		return merged
	}

	var i int
	merged = append(merged, ranges[i])
	i++
	for i < len(ranges) {

		last := merged[len(merged)-1]
		current := ranges[i]

		if current.start <= last.end {
			start := last.start
			end := max(last.end, current.end)
			merged[len(merged)-1] = NewRange(start, end)
		} else {
			merged = append(merged, current)
		}
		i++
	}
	return merged	
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
