package day6

import (
	"aoc/utils"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
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

var gridWidth = 131

func traverseGrid(gridData *GridData, pt1 bool, dir int, start []int, obstacle []int) bool {
	x := start[0]
	y := start[1]
	direction := dir
	grid := gridData.grid
	score := gridData.score
	next := true
	counter := 0
	visited := make([]bool, 131*131*len(directions))

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
		isObstacle := obstacle[0] == nextX && obstacle[1] == nextY
		if !isObstacle && (next == "." || next == "^" || next == "X") {
			x = nextX
			y = nextY

			if !pt1 {
				visitedKey := (y*131+x)*len(directions) + direction
				if visited[visitedKey] {
					return true
				}
				visited[visitedKey] = true
			} else if next != "X" {
				score++
				grid[nextY][nextX] = "X"
			}
		} else if next == "#" || isObstacle {
			direction = (direction + 1) % len(directions)
		}
	}
	return false
}

func part2(gridData *GridData, start []int) {
	var loop int32
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, runtime.NumCPU())

	for i := 0; i < len(gridData.grid); i++ {
		for j := 0; j < len(gridData.grid[i]); j++ {
			if gridData.grid[i][j] == "X" {
				wg.Add(1)
				semaphore <- struct{}{}

				go func(x, y int) {
					defer wg.Done()
					defer func() { <-semaphore }()

					if traverseGrid(gridData, false, 0, start, []int{x, y}) {
						atomic.AddInt32(&loop, 1)
					}
				}(j, i)
			}
		}
	}

	wg.Wait()
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
		score:     1,
		obstacles: make(map[string]int),
	}

	traverseGrid(&gridData, true, 0, start, []int{-1, -1})
	// Use the result from part 1 to solve part 2
	part2(&gridData, start)
}
