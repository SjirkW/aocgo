package day6

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var directions = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func indexOfDirection(direction []int, directions [][]int) int {
	for i, dir := range directions {
		if dir[0] == direction[0] && dir[1] == direction[1] {
			return i
		}
	}
	return -1
}

func traverseGrid(grid [][]string, x int, y int, direction []int, steps int) {
	// utils.PrintGrid(grid)
	previous := []int{x, y}

	dx := direction[0]
	dy := direction[1]
	horizontalIx := x + dx
	verticalIx := y + dy
	// Bounds check
	if horizontalIx < 0 || horizontalIx >= len(grid) {
		utils.PrintGrid(grid)
		fmt.Print("Out of bounds", steps)
		return
	}
	if verticalIx < 0 || verticalIx >= len(grid[0]) {
		utils.PrintGrid(grid)
		fmt.Print("Out of bounds", steps)
		return
	}

	next := grid[verticalIx][horizontalIx]
	if next == "." || next == "^" {
		grid[verticalIx][horizontalIx] = "X"
		traverseGrid(grid, horizontalIx, verticalIx, direction, steps+1)
	} else if next == "X" {
		traverseGrid(grid, horizontalIx, verticalIx, direction, steps)
	} else if next == "#" {
		nextDirection := directions[(indexOfDirection(direction, directions)+1)%len(directions)]
		traverseGrid(grid, previous[0], previous[1], nextDirection, steps)
	}

	// fmt.Print(grid[horizontalIx][verticalIx])
}

func Solve() {
	lines := utils.ReadInputAsLines(6, false)

	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	start := []int{0, 0}
	for i, row := range grid {
		for j, col := range row {
			if col == "^" {
				start = []int{j, i}
				break
			}
		}
	}

	traverseGrid(grid, start[0], start[1], directions[0], 0)
}
