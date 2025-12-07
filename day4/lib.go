package day4

import (
	"fmt"
	"strings"
)


const AT_SIGN_BYTE uint8 = 64

type Point struct {
	x, y int
}

func (p Point) equals(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func CountAccessibleRolls(input string) int {
	return len(GetAccessibleRolls(parseInput(input)))
}

func GetAccessibleRolls(floor [][]byte) []Point {

	var points []Point
	for row := range floor {
		for col := 0; col < len(floor[row]); col++ {
			point := Point {x: row, y: col}
			if hasRoll(point, floor) && len(getAdjacentRolls(floor, point)) < 4 {	
				points = append(points, point) 
			}
		}
	}

	return points 

}


// What if i flipped this on its head? If I know all the starting positions, why can't i then check all their surrounding nodes until i find what they have

func CountAllAccessibleRolls(input string) int {
	var total int
	var floor = parseInput(input)
	for {
		accessibleRollPoints := GetAccessibleRolls(floor)
		if len(accessibleRollPoints) == 0 {
			return total
		}
		total += len(accessibleRollPoints)

		for _, point := range accessibleRollPoints {
			floor[point.x][point.y] = []byte("x")[0]
		}
	}
}
func CountAllAccessibleRollsOld(input string) int {
	floor := parseInput(input)
	visited := make(map[Point]bool)

	for row := range floor {
		for col := 0; col < len(floor[row]); col++ {
			point := Point {x: row, y: col}
			if hasRoll(point, floor) {
				walk(point, floor, visited)
			}
		}
	}
	

	printResult(visited, floor)
	var total int
	for _, accessible := range visited {
		if accessible {
			total ++
		}
	}
	return total 
}


func hasRoll(point Point, floor [][]byte) bool {
	value := floor[point.x][point.y]
	return value == AT_SIGN_BYTE 
}

func walk(
	point Point,
	floor[][]byte,
	visited map[Point]bool,
) bool {
	if rem, ok := visited[point]; ok {
		return rem
	}
	
	adjacentRollPoints := getAdjacentRolls(floor, point)
	
	if len(adjacentRollPoints) < 4 {
		visited[point] = true 
		return visited[point]

	} else{
		visited[point] = false
	}

	var total int
	for _, adjacentPoint := range adjacentRollPoints {
		if !walk(adjacentPoint, floor, visited) {
			total += 1
		}	
	}
	visited[point] = total < 4
	return visited[point]
}


func printResult(removed map[Point]bool, floor [][]byte) {
	for row := range floor {
		for col := 0; col < len(floor[row]); col++ {
			point := Point {x: row, y: col}
			value := string(floor[point.x][point.y])
			if rem, ok := removed[point]; ok && rem {
				value = "x"
			}
			fmt.Printf("%v", value)
			
		}
		fmt.Printf("\n")
	}
	fmt.Print("\n")
}

func getAdjacentRolls(floor [][]byte, point Point) []Point {
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
		
		if !hasRoll(p, floor) {
			continue
		}
		adjacentRolls = append(adjacentRolls, p)
	}
	return adjacentRolls

}


func parseInput(input string) [][]byte {
	var parsed [][]byte
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parsed = append(parsed, []byte(line))
	}
	return parsed
}
