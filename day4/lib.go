package day4

import (
	"strings"
)

const AT_SIGN_BYTE uint8 = 64

func GetNumberOfAccessibleRolls(input string) int {
	var total int
	floor := parseInput(input)
	for i := range floor {
		for j := 0; j < len(floor[i]); j++ {
			tile := floor[i][j]
			if tile == AT_SIGN_BYTE {
				var adjacentRolls int
				adjacentTileIndexes := getAdjacentTiles(floor, i, j)
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
