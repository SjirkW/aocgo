package day6

import (
	"aoc/utils"
	"fmt"
	"strconv"
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
	loopScore int
	loopCheck bool
}

func indexOfDirection(direction []int, directions [][]int) int {
	for i, dir := range directions {
		if dir[0] == direction[0] && dir[1] == direction[1] {
			return i
		}
	}
	return -1
}

func traverseGrid(gridData GridData, pt1 bool) bool {
	stack := []GridData{gridData}
	visited := make(map[string]bool)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x := current.position[0]
		y := current.position[1]
		grid := current.grid
		direction := current.direction

		// Create a key to track visited positions
		posKey := fmt.Sprintf("%d,%d,%d,%d", x, y, direction[0], direction[1])

		// Skip if this exact position and direction has been visited
		if visited[posKey] {
			continue
		}
		visited[posKey] = true

		// Calculate next position
		horizontalIx := x + direction[0]
		verticalIx := y + direction[1]

		// Check bounds
		if horizontalIx < 0 || horizontalIx >= len(grid[0]) ||
			verticalIx < 0 || verticalIx >= len(grid) {
			fmt.Println("Out of bounds", gridData.score)
			continue
		}

		next := grid[verticalIx][horizontalIx]

		// Different handling based on the cell type
		switch next {
		case "0":
			// Loop detected
			return true
		case ".", "^", "X":
			// Move forward
			newGridData := current
			newGridData.position = []int{horizontalIx, verticalIx}

			if pt1 {
				if next != "X" {
					newGridData.score++
					newGridData.grid[verticalIx][horizontalIx] = "X"
				}
			} else {
				// Loop checking logic
				xs := strconv.Itoa(newGridData.position[0])
				ys := strconv.Itoa(newGridData.position[1])
				key := xs + "," + ys

				if newGridData.obstacles[key] != 1 {
					// Try turning
					for _, newDir := range directions {
						tryGridData := newGridData
						tryGridData.direction = newDir
						stack = append(stack, tryGridData)
					}
				}
			}

			stack = append(stack, newGridData)

		case "#":
			// Turn when hitting an obstacle
			newGridData := current
			newGridData.direction = directions[(indexOfDirection(direction, directions)+1)%len(directions)]
			stack = append(stack, newGridData)
		}
	}

	return false
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

	gridData := GridData{
		grid:      grid,
		position:  start,
		direction: directions[0],
		score:     0,
		obstacles: make(map[string]int),
		loopScore: 0,
		loopCheck: true,
	}

	// traverseGrid(gridData, true)
	traverseGrid(gridData, true)
	fmt.Println(gridData.score)
}
