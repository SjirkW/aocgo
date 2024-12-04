package day4

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func findStringInDirection(stringToFind string, input [][]string, x int, y int, dx int, dy int) bool {
	strinCharArr := strings.Split(stringToFind, "")

	l := len(strinCharArr)
	ixl := l - 1
	// Check if the string can fit in the direction
	if x+ixl*dx >= len(input) ||
		x+ixl*dx < 0 ||
		y+ixl*dy >= len(input[0]) ||
		y+ixl*dy < 0 {
		return false
	}

	// Check if the string is in the direction
	for i := 0; i < len(strinCharArr); i++ {
		horizontalIx := x + i*dx
		verticalIx := y + i*dy

		if horizontalIx < 0 || horizontalIx >= len(input) {
			fmt.Println("Out of bounds horizontal", x, dx, dy, i, len(input), len(input[0]))
			return false
		}

		if verticalIx < 0 || verticalIx >= len(input[0]) {
			fmt.Println("Out of bounds vertical", y, dy, i, len(input), len(input[0]))
			return false
		}

		if input[horizontalIx][verticalIx] != strinCharArr[i] {
			return false
		}

	}

	return true
}

func findStringInAllDirections(stringToFind string, input [][]string, x int, y int) int {
	// Check all directions
	total := 0
	if findStringInDirection(stringToFind, input, x, y, 0, 1) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, 1, 0) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, 1, 1) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, 1, -1) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, 0, -1) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, -1, 0) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, -1, 1) {
		total++
	}

	if findStringInDirection(stringToFind, input, x, y, -1, -1) {
		total++
	}

	return total
}

func Solve() {
	lines := utils.ReadInputAsLines(4, false)

	// Create a multideimensional array
	var arr [][]string
	for _, line := range lines {
		arr = append(arr, strings.Split(line, ""))
	}

	total := 0
	for i, row := range arr {
		for j, col := range row {
			if col == "X" {
				total += findStringInAllDirections("XMAS", arr, i, j)
			}
			// fmt.Print(col)
		}
		// fmt.Println()
	}

	fmt.Println("Part 1:", total)
	// fmt.Println("\nDay 4", arr)
}
