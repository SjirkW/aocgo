package day4

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func find(stringToFind string, input [][]string, x int, y int, dx int, dy int) bool {
	strinCharArr := strings.Split(stringToFind, "")

	for i := 0; i < len(strinCharArr); i++ {
		horizontalIx := x + i*dx
		verticalIx := y + i*dy
		// Bounds check
		if horizontalIx < 0 || horizontalIx >= len(input) {
			return false
		}
		if verticalIx < 0 || verticalIx >= len(input[0]) {
			return false
		}

		if input[horizontalIx][verticalIx] != strinCharArr[i] {
			return false
		}
	}

	return true
}

func findStringInAllDirections(stringToFind string, input [][]string, x int, y int) int {
	directions := []int{-1, 0, 1}

	total := 0
	for _, dx := range directions {
		for _, dy := range directions {
			if dx == 0 && dy == 0 {
				continue
			}

			if find(stringToFind, input, x, y, dx, dy) {
				total++
			}
		}
	}

	return total
}

func findStringInAllDirections2(input [][]string, x int, y int) int {
	if (find("MAS", input, x-1, y-1, 1, 1) ||
		find("SAM", input, x-1, y-1, 1, 1)) &&
		(find("MAS", input, x-1, y+1, 1, -1) ||
			find("SAM", input, x-1, y+1, 1, -1)) {
		return 1
	}

	return 0
}

func Solve() {
	lines := utils.ReadInputAsLines(4, false)

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
		}
	}

	fmt.Println("Part 1:", total)

	total2 := 0
	for i, row := range arr {
		for j, col := range row {
			if col == "A" {
				total2 += findStringInAllDirections2(arr, i, j)
			}
		}
	}

	fmt.Println("Part 2:", total2)
}
