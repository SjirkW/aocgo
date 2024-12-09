package day9

import (
	"aoc/utils"
	"fmt"
)

func Solve() {
	lines := utils.ReadInputAsLines(9, true)
	line := lines[0]

	fmt.Println(line)

	idCounter := 0
	isBlock := true

	var blocks [][]int
	for _, char := range line {
		count := utils.StringToInt(string(char))
		if isBlock {
			blocks = append(blocks, make([]int, 2))
			blocks[idCounter][0] = count
		} else {
			blocks[idCounter][1] = count
		}

		if !isBlock {
			idCounter++
		}
		isBlock = !isBlock
	}

	fmt.Println(blocks)

	// Loop backwards
	firstEmptySpaceIndex := -1
	for index, block := range blocks {
		if block[1] > 0 {
			firstEmptySpaceIndex = index
			break
		}
	}

	fmt.Println("First empty space index: ", firstEmptySpaceIndex)

	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]
		blockAmount := block[0]
		for j := 0; j < blockAmount; j++ {
			blocks[firstEmptySpaceIndex][= block[1]
		}
	}
}
