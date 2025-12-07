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


func GetNTimelines(diagram [][]byte) int {
	start := findStart(diagram)
	visited := make(map[Point]int)
	return getFutureTimelines(start, diagram, visited)
}

func getFutureTimelines(point Point, diagram [][]byte, visited map[Point]int) int {
	// To make this run quickly
	if timelines, ok := visited[point]; ok {
		return timelines
	}

	nextPoint := NewPoint(point.row + 1, point.col)
	if nextPoint.row >= len(diagram) {
		return 1
	}

	switch diagram[nextPoint.row][nextPoint.col] {
	case EMPTY:
		timelines := getFutureTimelines(nextPoint, diagram, visited)
		visited[point] = timelines
		return timelines
	case BARRIER:
		var timelines int
		left := NewPoint(nextPoint.row, nextPoint.col-1)
		if left.col >=0 {
			timelines+= getFutureTimelines(left, diagram, visited)

		}
		right := NewPoint(nextPoint.row, nextPoint.col+1)
		
		if right.col < len(diagram[0]) {
			timelines+= getFutureTimelines(right, diagram, visited)
		}
		visited[point]= timelines
		return timelines
	default:
		panic("unknown char")
	}
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
