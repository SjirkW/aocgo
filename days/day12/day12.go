package day12

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var fillMap = make(map[int]bool)

func FloodFill(grid [][]string, x int, y int, fill string) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 0
	}

	if grid[y][x] != fill || grid[y][x] == strings.ToLower(fill) {
		return 0
	}

	grid[y][x] = strings.ToLower(fill)

	result := 1

	result += FloodFill(grid, x+1, y, fill)
	result += FloodFill(grid, x-1, y, fill)
	result += FloodFill(grid, x, y+1, fill)
	result += FloodFill(grid, x, y-1, fill)

	return result
}

func Solve() {
	lines := utils.ReadInputAsLines(12, true)

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}

	for y, row := range grid {
		for x, _ := range row {
			res := FloodFill(grid, x, y, grid[y][x])
			if res > 0 {
				fmt.Println(res)
			}
		}
	}

	// fmt.Println("\nDay 12")
	// utils.PrintGrid(grid)
	// fmt.Println("Part 1:")
}
