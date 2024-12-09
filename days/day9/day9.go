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

func pt1(blocks []int) {
	firstEmpty := -1
	// Find the first empty space
	for i := range blocks {
		if blocks[i] == -1 {
			firstEmpty = i
			break
		}
	}

	// Loop from the end of the list and add keys to the empty spaces in front
	for i := len(blocks) - 1; i >= 0; i-- {
		blocks[firstEmpty] = blocks[i]
		blocks[i] = -1

		// Find next empty space
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

	fmt.Println("\nDay 9")
	fmt.Println("Part 1:", getResult(blocks))
}

func pt2(pt2Blocks []int, blockSizes map[int]int, emptySpacesMap map[int][]int) {
	keys := make([]int, 0, len(emptySpacesMap))
	for key := range emptySpacesMap {
		keys = append(keys, key)
	}

	// Loop from the end of the list and add keys to the empty spaces in front
	for i := len(pt2Blocks) - 1; i >= 0; i-- {
		id := pt2Blocks[i]
		length := blockSizes[id]

		if id == -1 {
			continue
		}

		// Find the lowest possible index to add
		smallestIndex := 9999999
		size := -1
		for _, key := range keys {
			if key >= length && len(emptySpacesMap[key]) > 0 && emptySpacesMap[key][0] < smallestIndex {
				size = key
				smallestIndex = emptySpacesMap[key][0]
			}
		}

		if size == -1 {
			continue
		}

		// Get the correct index
		indexes := emptySpacesMap[size]
		index := indexes[0]
		if index < i {
			// Overwrite empty blocks and remove old ids
			for x := 0; x < length; x++ {
				pt2Blocks[index+x] = id
				pt2Blocks[i-x] = -1
			}

			// Remove the old size and add a new one
			emptySpacesMap[size] = indexes[1:]
			newSize := size - length
			newIndex := index + length
			if newSize > 0 {
				arrayToAdd := emptySpacesMap[newSize]
				for x := 0; x < len(arrayToAdd); x++ {
					if newIndex < arrayToAdd[x] {
						emptySpacesMap[newSize] = append(arrayToAdd[:x], append([]int{newIndex}, arrayToAdd[x:]...)...)
						break
					}
				}
			}
		}
	}

	fmt.Println("Part 2:", getResult(pt2Blocks))
}

func Solve() {
	lines := utils.ReadInputAsLines(9, false)
	line := lines[0]
	idCounter := 0
	isBlock := true

	var blocks []int
	var blocks2 []int
	blockSizes := make(map[int]int)
	emptySpacesMap := make(map[int][]int)

	for _, char := range line {
		count := utils.StringToInt(string(char))

		for i := 0; i < count; i++ {
			if isBlock {
				blocks = append(blocks, idCounter)
				blocks2 = append(blocks2, idCounter)
			} else {
				if i == 0 {
					emptySpacesMap[count] = append(emptySpacesMap[count], len(blocks))
				}
				blocks = append(blocks, -1)
				blocks2 = append(blocks2, -1)
			}
		}

		if isBlock {
			blockSizes[idCounter] = count
		} else {
			idCounter++
		}
		isBlock = !isBlock
	}

	pt1(blocks)
	pt2(blocks2, blockSizes, emptySpacesMap)
}
