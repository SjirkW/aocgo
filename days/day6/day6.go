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

func traverseGrid(gridData GridData, pt1 bool, dir int, start []int) bool {
	x := start[0]
	y := start[1]
	direction := dir
	grid := gridData.grid
	score := gridData.score
	next := true
	counter := 0
	visited := make(map[string]bool)

	for next {
		counter++
		dx := directions[direction][0]
		dy := directions[direction][1]
		nextX := x + dx
		nextY := y + dy

		if (nextX < 0 || nextX >= len(grid)) ||
			(nextY < 0 || nextY >= len(grid[0])) {
			if pt1 {
				fmt.Println("Part 1:", score)
			}
			return false
		}

		next := grid[nextY][nextX]
		if next == "." || next == "^" || next == "X" {
			x = nextX
			y = nextY

			if !pt1 {
				visitedKey := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(dx) + "," + strconv.Itoa(dy)
				if visited[visitedKey] {
					return true
				}
				visited[visitedKey] = true
			} else if next != "X" {
				score++
				grid[nextY][nextX] = "X"
			}
		} else if next == "#" {
			direction = (direction + 1) % len(directions)
		}
	}
	return false
}

func part2(gridData GridData, start []int) {
	loop := 0
	for i := 0; i < len(gridData.grid); i++ {
		for j := 0; j < len(gridData.grid[i]); j++ {
			if gridData.grid[i][j] == "X" {
				gridData.grid[i][j] = "#"
				if traverseGrid(gridData, false, 0, start) {
					loop++
				}
				gridData.grid[i][j] = "X"
			}
		}
	}

	fmt.Println("Part 2:", loop)
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

	traverseGrid(gridData, true, 0, start)
	part2(gridData, start)
}
