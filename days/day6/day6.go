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
	x := gridData.position[0]
	y := gridData.position[1]
	direction := gridData.direction
	previous := []int{x, y}
	grid := gridData.grid
	checkLoop := gridData.loopCheck
	score := gridData.score
	next := true

	for next {
		previous = []int{x, y}
		horizontalIx := x + direction[0]
		verticalIx := y + direction[1]
		if horizontalIx < 0 || horizontalIx >= len(grid) {
			// utils.PrintGrid(grid)
			fmt.Println("Out of bounds", score)
			next = false
			return false
		}
		if verticalIx < 0 || verticalIx >= len(grid[0]) {
			// utils.PrintGrid(grid)
			next = false
			fmt.Println("Out of bounds", score)
			return false
		}

		next := grid[verticalIx][horizontalIx]
		if next == "0" {
			// fmt.Println("Loop")
			return true
		} else if next == "." || next == "^" || next == "X" {
			x = horizontalIx
			y = verticalIx

			if pt1 {
				if next != "X" {
					score++
					grid[verticalIx][horizontalIx] = "X"
				}
			} else if checkLoop {
				xs := strconv.Itoa(gridData.position[0])
				ys := strconv.Itoa(gridData.position[1])
				key := xs + "," + ys

				hasObstacle := gridData.obstacles[key] == 1
				if !hasObstacle {
					prev := grid[verticalIx][horizontalIx]
					prevDirection := []int{gridData.direction[0], gridData.direction[1]}

					gridData.loopCheck = false
					gridData.grid[verticalIx][horizontalIx] = "0"
					gridData.position = []int{x, y}
					gridData.direction = directions[(indexOfDirection(direction, directions)+1)%len(directions)]

					if traverseGrid(gridData, pt1) {
						gridData.obstacles[key] = 1
						gridData.loopScore++
						// fmt.Print(gridData.loopScore)
						// utils.PrintGrid(grid)
					}

					gridData.loopCheck = true
					gridData.grid[verticalIx][horizontalIx] = prev
					gridData.position = []int{horizontalIx, verticalIx}
					gridData.direction = prevDirection
				}
			}

			// return traverseGrid(gridData, pt1)
		} else if next == "#" {
			direction = directions[(indexOfDirection(direction, directions)+1)%len(directions)]
			x = previous[0]
			y = previous[1]
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
