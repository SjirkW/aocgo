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
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

type GridData struct {
	grid       [][]string
	position   []int
	direction  []int
	score      int
	obstacles  map[string]int
	directions map[int]int
}

var gridWidth = 131

func getKey(x, y int) int {
	return x<<16 | y
}

func traverseGrid(gridData *GridData, findLoop bool, dir int, start []int, obstacle []int) bool {
	x, y := start[0], start[1]
	direction := dir
	grid := gridData.grid
	score := gridData.score
	gridWidth, gridHeight := len(grid[0]), len(grid)
	visited := make([]bool, gridWidth*gridHeight*4)

	for {
		dx, dy := directions[direction][0], directions[direction][1]
		nextX, nextY := x+dx, y+dy

		if nextX < 0 || nextX >= gridWidth || nextY < 0 || nextY >= gridHeight {
			if !findLoop {
				fmt.Println("Part 1:", score)
			}
			return false
		}

		next := grid[nextY][nextX]
		isObstacle := obstacle[0] == nextX && obstacle[1] == nextY

		if !isObstacle && next != "#" {
			x, y = nextX, nextY

			if findLoop {
				key := (y*gridWidth+x)*4 + direction
				if visited[key] {
					return true
				}
				visited[key] = true
			} else if next != "X" {
				score++
				grid[nextY][nextX] = "X"
				gridData.directions[getKey(x, y)] = direction
			}
		} else {
			direction = (direction + 1) % 4
		}
	}
}

func part2(gridData *GridData, start []int) {
	var loop int32
	var wg sync.WaitGroup
	workerPool := make(chan struct{}, runtime.NumCPU())

	for i := 0; i < len(gridData.grid); i++ {
		for j := 0; j < len(gridData.grid[i]); j++ {
			if gridData.grid[i][j] == "X" {
				wg.Add(1)
				workerPool <- struct{}{}

				go func(x, y int) {
					defer wg.Done()
					defer func() { <-workerPool }()

					dir := gridData.directions[getKey(x, y)]
					start := []int{x, y}
					if dir == 0 {
						start[1]++
					} else if dir == 1 {
						start[0]--
					} else if dir == 2 {
						start[1]--
					} else if dir == 3 {
						start[0]++
					}
					if traverseGrid(gridData, true, dir, start, []int{x, y}) {
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
		grid:       grid,
		position:   start,
		direction:  directions[0],
		score:      1,
		obstacles:  make(map[string]int),
		directions: make(map[int]int),
	}

	traverseGrid(&gridData, false, 0, start, []int{-1, -1})
	part2(&gridData, start)
}
