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

func traverseGrid(gridData GridData, pt1 bool, checkLoop bool, dir []int, start []int, find []int) bool {
	x := start[0]
	y := start[1]
	direction := dir
	grid := gridData.grid
	score := gridData.score
	next := true

	for next {
		nextX := x + direction[0]
		nextY := y + direction[1]

		if nextX < 0 || nextX >= len(grid) {
			// utils.PrintGrid(grid)
			fmt.Println("Out of bounds", score)
			next = false
			return false
		}
		if nextY < 0 || nextY >= len(grid[0]) {
			// utils.PrintGrid(grid)
			next = false
			fmt.Println("Out of bounds", score)
			return false
		}

		next := grid[nextY][nextX]
		if find[0] == nextX && find[1] == nextY {
			fmt.Println("Loop")
			return true
		} else if next == "." || next == "^" || next == "X" {
			x = nextX
			y = nextY

			if pt1 {
				if next != "X" {
					score++
					grid[nextY][nextX] = "X"
				}
			} else if checkLoop {
				xs := strconv.Itoa(x)
				ys := strconv.Itoa(y)
				key := xs + "," + ys

				hasObstacle := gridData.obstacles[key] == 1
				if !hasObstacle {
					direction := directions[(indexOfDirection(direction, directions)+1)%len(directions)]
					find := []int{x, y}
					previous := []int{x, y}

					if traverseGrid(gridData, pt1, false, direction, previous, find) {
						gridData.obstacles[key] = 1
						gridData.loopScore++
					}
				}
			}

		} else if next == "#" {
			direction = directions[(indexOfDirection(direction, directions)+1)%len(directions)]
		}
	}
	return false
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
		loopScore: 0,
		loopCheck: true,
	}

	// traverseGrid(gridData, true)
	find := []int{-999, -999}
	traverseGrid(gridData, true, true, directions[0], start, find)
	fmt.Println(gridData.score)
}
