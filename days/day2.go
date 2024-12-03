package days

import (
	"aoc/utils"
	"fmt"
)

func isIncreasing(numbers []int) bool {
	increaseCount := 0

	for i := 0; i < 3; i++ {
		if numbers[i] < numbers[i+1] {
			increaseCount++
		}
	}

	return increaseCount > 1
}

func lineIsValid(numbers []int) bool {
	increasing := isIncreasing(numbers)
	for j := 0; j < len(numbers)-1; j++ {
		diff := numbers[j+1] - numbers[j]
		if increasing {
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			if diff > -1 || diff < -3 {
				return false
			}
		}
	}

	return true
}

func pt1() {
	lines := utils.ReadInputAsLines(2, false)

	valid := len(lines)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		numbers := utils.StringToIntArray(line)
		if !lineIsValid(numbers) {
			valid--
		}
	}

	fmt.Printf("Part 1 Result: %d\n", valid)
}

func pt2() {
	lines := utils.ReadInputAsLines(2, false)

	valid := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		numbers := utils.StringToIntArray(line)
		if lineIsValid(numbers) {
			valid++
		} else {
			for j := 0; j < len(numbers); j++ {
				temp := append(numbers[:j:j], numbers[j+1:]...)

				if lineIsValid(temp) {
					valid++
					break
				}
			}
		}
	}

	fmt.Printf("Part 2 Result: %d\n", valid)
}

func Solve2() {
	pt1()
	pt2()
}
