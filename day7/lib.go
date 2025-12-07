package day7

import (
	"strings"
)

const START = 83 
const EMPTY = 46 
const BARRIER = 94 

func SolvePart1(input string) int {
	diagram := parseInput(input)
	return GetNSplits(diagram)

	
}
func SolvePart2(input string) int {
	diagram := parseInput(input)
	return GetNTimelines(diagram)
}


func GetNSplits(manifoldDiagram [][]byte) int {
	var nSplits int 
	var stack []Point
	var visited = make(map[Point]bool)

	start := findStart(manifoldDiagram)	
	stack = append(stack, start)

	visited[start] = true

	for len(stack) != 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		next := NewPoint(point.row + 1, point.col)
		
		if next.row >= len(manifoldDiagram) {
			continue
		}

		if _, ok := visited[next]; ok {
			continue
		}

		switch manifoldDiagram[next.row][next.col] {
		case EMPTY:
			visited[next] = true
			stack = append(stack, next)
			continue
		case BARRIER:
			left := NewPoint(next.row, next.col-1)
			right := NewPoint(next.row, next.col+1)
			
			if _, ok := visited[left]; !ok || left.col >=0 {
				visited[left] = true
				stack = append(stack, left)
			}
			if _, ok := visited[right]; !ok || right.col >=0 {
				visited[right] = true
				stack = append(stack, right)
			}
			nSplits++
			continue
		}
		
	}
	return nSplits


}
// Gets the number of timelines of a single particle
func GetNTimelines(manifoldDiagram [][]byte) int  {
	var stack []Point
	var finalPoints []Point

	start := findStart(manifoldDiagram)	
	stack = append(stack, start)

	for len(stack) != 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nextPoint := NewPoint(point.row + 1, point.col)
		
		// Base case: if we go past the end we have completed one timeline
		if nextPoint.row >= len(manifoldDiagram) {
			finalPoints = append(finalPoints, point)
			continue
		}

		switch manifoldDiagram[nextPoint.row][nextPoint.col] {
		// We can just carry on, the timeline does not need to split
		case EMPTY:
			stack = append(stack, nextPoint)
			continue
		case BARRIER: // The time line splits once
			left := NewPoint(nextPoint.row, nextPoint.col-1)
			right := NewPoint(nextPoint.row, nextPoint.col+1)
			
			if left.col >=0 {
				stack = append(stack, left)
			}
			if right.col < len(manifoldDiagram[0]) {
				stack = append(stack, right)
			}
			continue
		default:
			panic("unrecognised symbol")
		}
	}
	return len(finalPoints)
}

func parseInput(input string) [][]byte {
	var diagram [][]byte

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		diagram = append(diagram, []byte(line))
	}

	return diagram
}

func findStart(manifold [][]byte) Point {
	var row int
	var col int
	
	for col < len(manifold) {
		value := manifold[row][col]
		if value == START {
			break
		} else {
			col++
		}
	}

	if col > len(manifold) {
		panic("No start detected")
	}

	return NewPoint(row, col)
}  


type Point struct {
	row int
	col int
}

func NewPoint(row, col int) Point {
	return Point{row, col}
}
