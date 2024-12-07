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
}

func indexOfDirection(direction []int, directions [][]int) int {
	for i, dir := range directions {
		if dir[0] == direction[0] && dir[1] == direction[1] {
			return i
		}
	}
	return -1
}

func checkLoop(gridData GridData, dir []int, start []int) bool {
	x := start[0]
	y := start[1]
	direction := dir
	grid := gridData.grid
	next := true
	visited := make(map[string]bool)

	for next {

		nextX := x + direction[0]
		nextY := y + direction[1]

		if (nextX < 0 || nextX >= len(grid)) ||
			(nextY < 0 || nextY >= len(grid[0])) {
			return false
		}

		next := grid[nextY][nextX]
		if next == "#" || next == "0" {
			direction = directions[(indexOfDirection(direction, directions)+1)%len(directions)]
		} else if next == "." || next == "^" {
			visitedKey := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(direction[0]) + "," + strconv.Itoa(direction[1])
			if visited[visitedKey] {
				return true
			}
			visited[visitedKey] = true

			x = nextX
			y = nextY
		}
	}
	return false
}

func traverseGrid(gridData GridData, pt1 bool, dir []int, start []int) bool {
	x := start[0]
	y := start[1]
	direction := dir
	grid := gridData.grid
	score := gridData.score
	next := true
	counter := 0

	for next {
		counter++
		nextX := x + direction[0]
		nextY := y + direction[1]

		if (nextX < 0 || nextX >= len(grid)) ||
			(nextY < 0 || nextY >= len(grid[0])) {
			if pt1 {
				fmt.Println("Part 1:", score)
			}
			return false
		}

		next := grid[nextY][nextX]
		if next == "." || next == "^" || next == "X" {
			if pt1 {
				x = nextX
				y = nextY
				if next != "X" {
					score++
					grid[nextY][nextX] = "X"
				}
			} else if counter > 0 {
				obstacleKey := strconv.Itoa(nextX) + "," + strconv.Itoa(nextY)
				// If the obstacle is not already in the map
				if gridData.obstacles[obstacleKey] != 1 {
					gridData.grid[nextY][nextX] = "0"
					if checkLoop(gridData, direction, []int{x, y}) {
						gridData.obstacles[obstacleKey] = 1
					}
					gridData.grid[nextY][nextX] = "."
				}

				x = nextX
				y = nextY
			}

		} else if next == "#" {
			direction = directions[(indexOfDirection(direction, directions)+1)%len(directions)]
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
	}

	// traverseGrid(gridData, true, directions[0], start)
	traverseGrid(gridData, false, directions[0], start)
	fmt.Println("Part 2:", len(gridData.obstacles))
}
