package day5

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
)

func Solve() {
	lines := utils.ReadInputAsLines(5, false)

	orderMap := make(map[string]int)
	var toSortLists [][]int

	isParsingNums := true
	for _, line := range lines {
		if line == "" {
			isParsingNums = false
			continue
		}

		if isParsingNums {
			orderMap[line] = 1
		} else {
			parsedRow := utils.StringToIntArray(line, ",")
			toSortLists = append(toSortLists, parsedRow)
		}
	}

	totalPt1 := 0
	totalPt2 := 0
	for _, list := range toSortLists {
		correct := true
		sort.Slice(list, func(i, j int) bool {
			a, b := strconv.Itoa(list[i]), strconv.Itoa(list[j])
			key := a + "|" + b

			if orderMap[key] == 1 {
				correct = false
				return true
			}

			return false
		})

		if correct {
			totalPt1 += list[len(list)/2]
		} else {
			totalPt2 += list[len(list)/2]
		}
	}

	fmt.Println("Day 5")
	fmt.Printf("Part 1: %d\n", totalPt1)
	fmt.Printf("Part 2: %d\n", totalPt2)
}
