package day9

import (
	"aoc/utils"
	"fmt"
)

func getResult(blocks []int) int {
	result := 0
	for i, block := range blocks {
		if block != -1 {
			result += block * i
		}
	}

	return result
}

func Solve() {
	lines := utils.ReadInputAsLines(9, false)
	line := lines[0]

	// fmt.Println(line)

	idCounter := 0
	isBlock := true

	var blocks []int
	blockSizes := make(map[int]int)
	emptySizes := make(map[int]int)
	for _, char := range line {
		count := utils.StringToInt(string(char))

		for i := 0; i < count; i++ {
			if isBlock {
				blocks = append(blocks, idCounter)
			} else {
				if i == 0 {
					emptySizes[len(blocks)] = count
				}
				blocks = append(blocks, -1)
			}
		}

		if isBlock {
			blockSizes[idCounter] = count
		} else {
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

	// copy blocks to blocks2
	pt2Blocks := make([]int, len(blocks))
	copy(pt2Blocks, blocks)

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

	fmt.Println("Part 1:", getResult(blocks))

	for i := len(pt2Blocks) - 1; i >= 0; i-- {
		id := pt2Blocks[i]
		length := blockSizes[id]

		if id == -1 {
			continue
		}
		for j := 0; j < len(pt2Blocks); j++ {
			if (i > j && pt2Blocks[j] == -1) &&
				(emptySizes[j] > 0 && emptySizes[j] >= length) {
				for x := 0; x < length; x++ {
					pt2Blocks[j+x] = id
					pt2Blocks[i-x] = -1
				}

				oldSize := emptySizes[j]
				delete(emptySizes, j)
				emptySizes[j+length] = oldSize - length
				break
			}
		}
	}

	fmt.Println("Part 2:", getResult(pt2Blocks))
}
