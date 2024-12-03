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

func pt2(allowFaulty bool) int {
	lines := utils.ReadInputAsLines(2, false)

	valid := 0
	for _, line := range lines {
		numbers := utils.StringToIntArray(line)
		if lineIsValid(numbers) {
			valid++
		} else if allowFaulty {
			for j := range numbers {
				temp := append(numbers[:j:j], numbers[j+1:]...)

				if lineIsValid(temp) {
					valid++
					break
				}
			}
		}
	}

	return valid
}

func Solve2() {
	fmt.Printf("Part 2 Result: %d\n", pt2(false))
	fmt.Printf("Part 2 Result: %d\n", pt2(true))
}
