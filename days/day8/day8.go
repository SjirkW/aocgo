package day8

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func getNodeCoords(node1 []int, node2 []int) [][]int {
	n1X := node1[0]
	n1Y := node1[1]
	n2X := node2[0]
	n2Y := node2[1]

	dx := utils.Abs(n2X - n1X)
	dy := utils.Abs(n2Y - n1Y)

	a1X, a1Y, a2X, a2Y := -1, -1, -1, -1
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

	return [][]int{{a1X, a1Y}, {a2X, a2Y}}
}

func nodeInBounds(node []int, grid [][]string) bool {
	return node[0] >= 0 && node[0] < len(grid[0]) && node[1] >= 0 && node[1] < len(grid)
}

func addNodeToGrid(grid [][]string, nodes [][]int, key string) int {
	counter := 0
	if nodeInBounds(nodes[0], grid) {
		node := nodes[0]
		if grid[node[1]][node[0]] != "#" && grid[node[1]][node[0]] != key {
			grid[node[1]][node[0]] = "#"
			counter++
		}
	}

	if nodeInBounds(nodes[1], grid) {
		node := nodes[1]
		if grid[node[1]][node[0]] != "#" && grid[node[1]][node[0]] != key {
			grid[node[1]][node[0]] = "#"
			counter++
		}
	}

	return counter
}

func resetGrid(grid [][]string, original [][]string) [][]string {
	gridCopy := make([][]string, len(grid))
	for i := range grid {
		gridCopy[i] = make([]string, len(grid[i]))
		copy(gridCopy[i], original[i])
	}

	// Add all "#" from grid to gridCopy
	for y, row := range grid {
		for x, cell := range row {
			if cell == "#" {
				gridCopy[y][x] = "#"
			}
		}
	}

	return gridCopy
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

	originalGrid := resetGrid(grid, grid)

	// Loop nodemap
	gridWithNodes := resetGrid(grid, originalGrid)
	count := 0
	for key, letterValues := range nodeMap {
		gridWithNodes = resetGrid(gridWithNodes, originalGrid)
		// fmt.Println("Key:", key)

		for i := 0; i < len(letterValues); i++ {
			for j := 0; j < len(letterValues); j++ {
				if i != j {
					coords := getNodeCoords(letterValues[i], letterValues[j])

					newNodes := addNodeToGrid(gridWithNodes, coords, key)
					count += newNodes
				}
			}
		}

		// utils.PrintGrid(gridWithNodes)
	}

	// fmt.Println(nodeMap)
	fmt.Println("Count:", count)
	utils.PrintGrid(gridWithNodes)
}

func Solve() {
	lines := utils.ReadInputAsLines(8, true)

	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	createAntiNodes(grid)

	// utils.PrintGrid(grid)
	// fmt.Println("Day 8:", lines)
}
