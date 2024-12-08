package day8

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func countNonEmptyCells(grid [][]string) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell != "." {
				count++
			}
		}
	}

	return count
}

func nodeInBounds(node []int, grid [][]string) bool {
	return node[0] >= 0 && node[0] < len(grid[0]) && node[1] >= 0 && node[1] < len(grid)
}

func getNodeCoords(grid [][]string, node1 []int, node2 []int, doSingle bool) map[string]int {
	n1X, n1Y := node1[0], node1[1]
	n2X, n2Y := node2[0], node2[1]

	antiNodePositions := make(map[string]int)

	dx := utils.Abs(n2X - n1X)
	dy := utils.Abs(n2Y - n1Y)

	do := true
	for do {
		do = false
		var a1X, a1Y, a2X, a2Y int

		if n1X < n2X {
			a1X = n1X - dx
			a2X = n2X + dx
		} else {
			a1X = n1X + dx
			a2X = n2X - dx
		}

		if n1Y < n2Y {
			a1Y = n1Y - dy
			a2Y = n2Y + dy
		} else {
			a1Y = n1Y + dy
			a2Y = n2Y - dy
		}

		n1X, n2X = a1X, a2X
		n1Y, n2Y = a1Y, a2Y

		n1 := []int{a1X, a1Y}
		if nodeInBounds(n1, grid) {
			positionKey := fmt.Sprintf("%d,%d", n1[0], n1[1])
			if antiNodePositions[positionKey] != 1 {
				antiNodePositions[positionKey] = 1
				if !doSingle {
					do = true
				}
			}
		}

		n2 := []int{a2X, a2Y}
		if nodeInBounds(n2, grid) {
			positionKey := fmt.Sprintf("%d,%d", n2[0], n2[1])
			if antiNodePositions[positionKey] != 1 {
				antiNodePositions[positionKey] = 1
				if !doSingle {
					do = true
				}
			}
		}
	}

	return antiNodePositions
}

func createAntiNodes(grid [][]string) {
	nodeMap := make(map[string][][]int)
	for y, row := range grid {
		for x, cell := range row {
			if cell != "." {
				nodeMap[cell] = append(nodeMap[cell], []int{x, y})
			}
		}
	}

	antiNodePositions := make(map[string]int)
	for _, letterValues := range nodeMap {
		for i := 0; i < len(letterValues); i++ {
			for j := 0; j < len(letterValues); j++ {
				if i != j {
					for k, v := range getNodeCoords(grid, letterValues[i], letterValues[j], true) {
						antiNodePositions[k] = v
					}
				}
			}
		}
	}

	fmt.Println("Pt1:", len(antiNodePositions))

	antiNodePositions2 := make(map[string]int)
	for _, letterValues := range nodeMap {
		for i := 0; i < len(letterValues); i++ {
			for j := 0; j < len(letterValues); j++ {
				if i != j {
					for k, v := range getNodeCoords(grid, letterValues[i], letterValues[j], false) {
						antiNodePositions2[k] = v
					}
				}
			}
		}
	}

	for key := range antiNodePositions2 {
		coords := strings.Split(key, ",")
		x, y := utils.StringToInt(coords[0]), utils.StringToInt(coords[1])
		grid[y][x] = "#"
	}
	nonEmpty := countNonEmptyCells(grid)
	fmt.Println("Pt2:", nonEmpty)
	// utils.PrintGrid(grid)
}

func Solve() {
	lines := utils.ReadInputAsLines(8, false)

	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	createAntiNodes(grid)
}
