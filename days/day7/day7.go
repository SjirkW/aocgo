package day7

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func part1(lines []string) int {
	return 0
}

func hasResult(result int, current int, numbers []int) bool {
	if len(numbers) == 1 {
		if current+numbers[0] == result {
			return true
		} else if current*numbers[0] == result {
			return true
		} else {
			return false
		}
	}

	left := numbers[0]
	rest := numbers[1:]
	if hasResult(result, current+left, rest) {
		return true
	} else if hasResult(result, current*left, rest) {
		return true
	}

	return false
}

func Solve() {
	lines := utils.ReadInputAsLines(7, false)

	var left []int
	var right [][]int

	// Line looks like "190: 10 19"
	for _, line := range lines {
		split := strings.Split(line, ": ")

		left = append(left, utils.StringToInt(split[0]))
		right = append(right, utils.StringToIntArray(split[1], " "))
	}

	r := 0
	for i := 0; i < len(left); i++ {
		if hasResult(left[i], 0, right[i]) {
			r += left[i]
		}
	}
	fmt.Println(r)

	// fmt.Println("Part 1:", part1(lines))
}
