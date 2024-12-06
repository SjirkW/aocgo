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

type GridData struct {
	grid      [][]string
	position  []int
	direction []int
	score     int
	obstacles map[string]int
}

func indexOfDirection(direction []int, directions [][]int) int {
	for i, dir := range directions {
		if dir[0] == direction[0] && dir[1] == direction[1] {
			return i
		}
	}
	return -1
}

func traverseGrid(gridData GridData) {
	// utils.PrintGrid(grid)
	x := gridData.position[0]
	y := gridData.position[1]
	grid := gridData.grid
	score := gridData.score
	direction := gridData.direction
	previous := []int{x, y}

	horizontalIx := x + gridData.direction[0]
	verticalIx := y + gridData.direction[1]
	// Bounds check
	if horizontalIx < 0 || horizontalIx >= len(grid) {
		utils.PrintGrid(grid)
		fmt.Print("Out of bounds", score)
		return
	}
	if verticalIx < 0 || verticalIx >= len(grid[0]) {
		utils.PrintGrid(grid)
		fmt.Print("Out of bounds", score)
		return
	}

	next := grid[verticalIx][horizontalIx]
	if next == "." || next == "^" {
		grid[verticalIx][horizontalIx] = "X"
		gridData.position = []int{horizontalIx, verticalIx}
		gridData.score++
		traverseGrid(gridData)
	} else if next == "X" {
		gridData.position = []int{horizontalIx, verticalIx}
		traverseGrid(gridData)
	} else if next == "#" {
		nextDirection := directions[(indexOfDirection(direction, directions)+1)%len(directions)]
		gridData.direction = nextDirection
		gridData.position = []int{previous[0], previous[1]}
		traverseGrid(gridData)
	}
}

func Solve() {
	lines := utils.ReadInputAsLines(6, true)

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

	gridData := GridData{
		grid:      grid,
		position:  start,
		direction: directions[0],
		score:     0,
		obstacles: make(map[string]int),
	}

	traverseGrid(gridData)
}
