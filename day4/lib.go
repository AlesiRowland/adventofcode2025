package day4

import (
	"fmt"
	"strings"
)


const AT_SIGN_BYTE uint8 = 64

type Point struct {
	x, y int
}


func CountAccessibleRolls(input string) int {
	// This is essentially a depth first search where we need to remember where we've been and if the roll has already been removed.
	floor := parseInput(input)

	var total int

	nRows := len(floor)
	if nRows == 0 {
		return 0
	}
	nCols := len(floor[0])
	if nCols == 0 {
		panic("For now..") 
	}

	for x := range nRows {
		for y := range nCols {
			point := Point {x, y}
			if hasRoll(point, floor) && len(getAdjacentRolls(floor, point, nil)) < 4 {
				total += 1
			}
		}
	}
	return total
}

func GetNumberOfAccessibleRolls(input string) int {
	floor := parseInput(input)

	var total int
	for row := range floor {
		for col := 0; col < len(floor[row]); col++ {
			point := Point {x: row, y: col}
			if hasRoll(point, floor) {	
				var adjacentRolls int
				adjacentTileIndexes := getAdjacentTiles(floor, row, col)
				for _, tile := range adjacentTileIndexes {
					if tile == AT_SIGN_BYTE {
						adjacentRolls += 1
					}
				}
				if adjacentRolls < 4 {
					total += 1
				}
			}
		}
	}
	return total
}

func CountAllAccessibleRolls(input string) int {
	floor := parseInput(input)

	var total int

	visited := make(map[Point]bool) // remember where we've been
	
	nRows := len(floor)
	if nRows == 0 {
		return 0
	}
	nCols := len(floor[0])
	if nCols == 0 {
		panic("For now..") 
	}

	for x := range nRows {
		for y := range nCols {
			point := Point {x, y}
			if hasRoll(point, floor) && isRemoveable(point, nil, floor, visited) {
				fmt.Printf("%v\n", point)
				total += 1
			}
		}
	}
	return total
}


func hasRoll(point Point, floor [][]byte) bool {
	value := floor[point.x][point.y]
	return value == AT_SIGN_BYTE 
}

// We are starting recursively because this is confusing
func isRemoveable(point Point, last *Point, floor[][]byte, visited map[Point]bool) bool {
	if rem, ok := visited[point]; ok {
		return rem
	}
	
	visited[point] = false
	adjacentPoints := getAdjacentRolls(floor, point, last)
	if len(adjacentPoints) < 4 {
		visited[point] = true
		return visited[point]
	}

	var total int
	for _, adjacentPoint := range adjacentPoints {
		if !isRemoveable(adjacentPoint, &point, floor, visited) {
			total += 1
		}	
	}

	if total < 4 {
		visited[point] = true
	}

	return visited[point]
}


func getAdjacentRolls(floor [][]byte, point Point, last *Point) []Point {
	var adjacentRolls []Point
	adjacentPoints := [8]Point{
		Point{x: point.x - 1, y:point.y - 1},
		Point{x: point.x - 1, y:point.y},
		Point{x: point.x - 1, y: point.y + 1},
		Point{x: point.x, y: point.y - 1},
		Point{x:point.x, y:point.y + 1},
		Point{x:point.x + 1, y:point.y - 1},
		Point{x:point.x + 1, y:point.y},
		Point{x:point.x + 1, y:point.y + 1},
	}
	
	for _, p := range adjacentPoints {
		if p.x < 0 || p.y < 0 || p.x >= len(floor) || p.y >= len(floor[0]) {
			continue
		}
		
		if last != nil && p.x == last.y && p.y == last.y {
			continue
		}
		if !hasRoll(p, floor) {
			continue
		}
		adjacentRolls = append(adjacentRolls, p)
	}
	return adjacentRolls
}




func getAdjacentTiles(floor[][]byte , row int, column int) []byte {
	tileIndexes := [8][2]int{
		{row - 1, column - 1},
		{row - 1, column},
		{row - 1, column + 1},
		{row, column - 1},
		{row, column + 1},
		{row + 1, column - 1},
		{row + 1, column},
		{row + 1, column + 1},
	}
	
	var adjacentTiles []byte

	for _, index := range tileIndexes {
		if index[0] < 0 || index[1] < 0 || index[0] >= len(floor) || index[1] >= len(floor[0]) {
			continue
		}
		adjacentTiles = append(adjacentTiles, floor[index[0]][index[1]])
	}
	


	return adjacentTiles 

}


func parseInput(input string) [][]byte {
	var parsed [][]byte
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parsed = append(parsed, []byte(line))
	}
	return parsed
}
