package day10

import (
	"aoc/utils"
	"fmt"
)

type Point struct {
	x int
	y int
}

var directions = []Point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func FindTrialAmount(grid [][]int, point Point) int {
	result := 0
	pointsToCheck := make([]Point, 0)
	endPointMap := make(map[Point]bool)

	for _, dir := range directions {
		x := point.x + dir.x
		y := point.y + dir.y

		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			continue
		}

		nextVal := grid[y][x]

		if nextVal == 1 {
			pointsToCheck = append(pointsToCheck, Point{point.x + dir.x, point.y + dir.y})
		}
	}

	for len(pointsToCheck) > 0 {
		point = pointsToCheck[0]
		pointsToCheck = pointsToCheck[1:]
		score := grid[point.y][point.x]
		for _, dir := range directions {
			x := point.x + dir.x
			y := point.y + dir.y

			if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
				continue
			}

			if grid[y][x] == score+1 {
				if grid[y][x] == 9 {
					endPointMap[Point{x, y}] = true
					result++
				} else {
					pointsToCheck = append(pointsToCheck, Point{x, y})
				}
			}
		}
	}

	return result
}

func Part1(grid [][]int) {
	var startingPoints []Point

	for y, row := range grid {
		for x, val := range row {
			if val == 0 {
				startingPoints = append(startingPoints, Point{x, y})
			}
		}
	}

	fmt.Println(startingPoints)

	result := 0
	for _, point := range startingPoints {
		amount := FindTrialAmount(grid, point)
		// fmt.Println(amount, point)
		result += amount
	}

	fmt.Println(result)
}

func Solve() {
	lines := utils.ReadInputAsLines(10, false)

	grid := make([][]int, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
			} else {
				row = append(row, utils.StringToInt(string(char)))
			}
		}
		grid = append(grid, row)

	}

	Part1(grid)

	// utils.PrintIntGrid(grid)
}
