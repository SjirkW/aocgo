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

func getNodeCoords(grid [][]string, antiNodePositions map[string]int, node1 []int, node2 []int, doSingle bool, key string) {
	n1X := node1[0]
	n1Y := node1[1]
	n2X := node2[0]
	n2Y := node2[1]

	dx := utils.Abs(n2X - n1X)
	dy := utils.Abs(n2Y - n1Y)

	do := true
	for do {
		do = false

		a1X, a1Y, a2X, a2Y := -1, -1, -1, -1
		if n1X < n2X {
			a1X = n1X - dx
			a2X = n2X + dx
		} else {
			a1X = n1X + dx
			a2X = n2X - dx
		}

		n1X = a1X
		n2X = a2X

		if n1Y < n2Y {
			a1Y = n1Y - dy
			a2Y = n2Y + dy
		} else {
			a1Y = n1Y + dy
			a2Y = n2Y - dy
		}

		n1Y = a1Y
		n2Y = a2Y

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
	for key, letterValues := range nodeMap {
		for i := 0; i < len(letterValues); i++ {
			for j := 0; j < len(letterValues); j++ {
				if i != j {
					getNodeCoords(grid, antiNodePositions, letterValues[i], letterValues[j], true, key)
				}
			}
		}
	}

	fmt.Println("Pt1:", len(antiNodePositions))

	antiNodePositions2 := make(map[string]int)
	for key, letterValues := range nodeMap {
		for i := 0; i < len(letterValues); i++ {
			for j := 0; j < len(letterValues); j++ {
				if i != j {
					getNodeCoords(grid, antiNodePositions2, letterValues[i], letterValues[j], false, key)
				}
			}
		}
	}

	nonEmpty := countNonEmptyCells(grid)
	fmt.Println("Pt2:", len(antiNodePositions2)+nonEmpty)
}

func Solve() {
	lines := utils.ReadInputAsLines(8, false)

	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	createAntiNodes(grid)
}
