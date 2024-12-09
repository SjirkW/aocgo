package day9

import (
	"aoc/utils"
	"fmt"
	"sort"
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
	idCounter := 0
	isBlock := true

	var blocks []int
	blockSizes := make(map[int]int)

	var emptySpaces []struct {
		index int
		size  int
	}
	for _, char := range line {
		count := utils.StringToInt(string(char))

		for i := 0; i < count; i++ {
			if isBlock {
				blocks = append(blocks, idCounter)
			} else {
				if i == 0 {
					emptySpaces = append(emptySpaces, struct {
						index int
						size  int
					}{len(blocks), count})
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

	sort.Slice(emptySpaces, func(i, j int) bool {
		return emptySpaces[i].index < emptySpaces[j].index
	})

	// Loop backwards
	firstEmpty := -1
	for i := range blocks {
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

		for spaceIx, emptySpace := range emptySpaces {
			if emptySpace.index < i && emptySpace.size >= length {
				// Move block to the empty space
				for x := 0; x < length; x++ {
					pt2Blocks[emptySpace.index+x] = id
					pt2Blocks[i-x] = -1
				}
				i -= length - 1

				emptySpaces[spaceIx].size -= length
				emptySpaces[spaceIx].index += length
				break
			}
		}
	}

	fmt.Println("Part 2:", getResult(pt2Blocks))
}
