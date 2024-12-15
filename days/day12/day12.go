package day12

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type FloodFillResult struct {
	Area      int
	Perimeter int
	Corners   int
}

func CountCorners(grid [][]string, x int, y int, result *FloodFillResult) {
	current := strings.ToLower(grid[y][x])

	hasTopLeftCorner := (y == 0 || strings.ToLower(grid[y-1][x]) != current) && (x == 0 || strings.ToLower(grid[y][x-1]) != current)
	if hasTopLeftCorner {
		result.Corners++
	}
	hasTopRightCorner := (y == 0 || strings.ToLower(grid[y-1][x]) != current) && (x == len(grid)-1 || strings.ToLower(grid[y][x+1]) != current)
	if hasTopRightCorner {
		result.Corners++
	}
	hasBottomLeftCorner := (y == len(grid[0])-1 || strings.ToLower(grid[y+1][x]) != current) && (x == 0 || strings.ToLower(grid[y][x-1]) != current)
	if hasBottomLeftCorner {
		result.Corners++
	}
	hasBottomRightCorner := (y == len(grid[0])-1 || strings.ToLower(grid[y+1][x]) != current) && (x == len(grid)-1 || strings.ToLower(grid[y][x+1]) != current)
	if hasBottomRightCorner {
		result.Corners++
	}
}

func FloodFill(grid [][]string, x int, y int, fill string, result *FloodFillResult) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		result.Perimeter++
		return
	}

	if grid[y][x] == strings.ToLower(fill) {
		return
	} else if grid[y][x] != fill {
		result.Perimeter++
		return
	}

	// Mark the cell as visited
	grid[y][x] = strings.ToLower(fill)
	CountCorners(grid, x, y, result)

	// Increment the count
	result.Area++

	// Explore neighbors
	FloodFill(grid, x+1, y, fill, result)
	FloodFill(grid, x-1, y, fill, result)
	FloodFill(grid, x, y+1, fill, result)
	FloodFill(grid, x, y-1, fill, result)
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

	total := 0
	for y, row := range grid {
		for x, _ := range row {
			result := &FloodFillResult{Area: 0, Perimeter: 0, Corners: 0}
			FloodFill(grid, x, y, grid[y][x], result)
			if result.Area > 0 {
				fmt.Println(grid[y][x], result.Area, result.Perimeter, result.Corners)
				total += result.Area * result.Perimeter
			}
		}
	}

	fmt.Println("\nDay 12")
	fmt.Println("Part 1:", total)
}
