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

func FindTrialAmount(grid [][]int, point Point) []int {
	result1 := 0
	result2 := 0
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
					if (!endPointMap[Point{x, y}]) {
						result1++
						endPointMap[Point{x, y}] = true
					}
					result2++
				} else {
					pointsToCheck = append(pointsToCheck, Point{x, y})
				}
			}
		}
	}

	return []int{result1, result2}
}

func Part1And2(grid [][]int) {
	var startingPoints []Point

	for y, row := range grid {
		for x, val := range row {
			if val == 0 {
				startingPoints = append(startingPoints, Point{x, y})
			}
		}
	}

	result1 := 0
	result2 := 0
	for _, point := range startingPoints {
		amounts := FindTrialAmount(grid, point)
		// fmt.Println(amount, point)
		result1 += amounts[0]
		result2 += amounts[1]
	}

	fmt.Println("\nDay10")
	fmt.Println("Part 1", result1)
	fmt.Println("Part 2", result2)
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

	Part1And2(grid)
}
