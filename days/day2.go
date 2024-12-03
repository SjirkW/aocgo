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

func Solve2() int {
	lines := utils.ReadInputAsLines(2, false)

	valid := len(lines)
	for i := 0; i < len(lines); i++ {

		numbers := utils.StringToIntArray(lines[i])
		increasing := isIncreasing(numbers)
		for j := 0; j < len(numbers)-1; j++ {
			diff := numbers[j+1] - numbers[j]
			if increasing {
				if diff < 1 || diff > 3 {
					valid--
					break
				}
			} else {
				if diff > -1 || diff < -3 {
					valid--
					break
				}
			}
		}
	}

	fmt.Printf("Part 1 Result: %d\n", valid)

	return 1
}
