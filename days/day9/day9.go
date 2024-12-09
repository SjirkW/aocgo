package day9

import (
	"aoc/utils"
	"fmt"
)

func Solve() {
	lines := utils.ReadInputAsLines(9, false)
	line := lines[0]

	// fmt.Println(line)

	idCounter := 0
	isBlock := true

	var blocks []int
	for _, char := range line {
		count := utils.StringToInt(string(char))

		for i := 0; i < count; i++ {
			if isBlock {
				blocks = append(blocks, idCounter)
			} else {
				blocks = append(blocks, -1)
			}
		}

		if !isBlock {
			idCounter++
		}
		isBlock = !isBlock
	}

	// fmt.Println(blocks)

	// Loop backwards
	firstEmpty := -1
	for i, _ := range blocks {
		if blocks[i] == -1 {
			firstEmpty = i
			break
		}
	}
	for i := len(blocks) - 1; i >= 0; i-- {
		blocks[firstEmpty] = blocks[i]
		blocks[i] = -1

		// Find next empty
		for x := firstEmpty; x < len(blocks); x++ {
			if blocks[x] == -1 {
				firstEmpty = x
				break
			}
		}

		if firstEmpty >= i {
			break
		}
	}

	result := 0
	for i, block := range blocks {
		if block != -1 {
			result += block * i
		}
	}

	fmt.Println(result)
}
