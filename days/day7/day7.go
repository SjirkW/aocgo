package day7

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

func concatInts(a int, b int) int {
	if b < 10 {
		return a*10 + b
	} else if b < 100 {
		return a*100 + b
	} else if b < 1000 {
		return a*1000 + b
	} else {
		return a*int(math.Pow(10, math.Floor(math.Log10(float64(b))+1))) + b
	}
}

func hasResult(result int, current int, numbers []int, concat bool) bool {
	if len(numbers) == 0 || current > result {
		return current == result
	}

	left := numbers[0]
	rest := numbers[1:]
	if hasResult(result, current+left, rest, concat) {
		return true
	} else if hasResult(result, current*left, rest, concat) {
		return true
	} else if concat && hasResult(result, concatInts(current, left), rest, concat) {
		return true
	}

	return false
}

func Solve() {
	lines := utils.ReadInputAsLines(7, false)

	var left []int
	var right [][]int
	for _, line := range lines {
		split := strings.Split(line, ": ")

		left = append(left, utils.StringToInt(split[0]))
		right = append(right, utils.StringToIntArray(split[1], " "))
	}

	pt1 := 0
	pt2 := 0
	for i := 0; i < len(left); i++ {
		if hasResult(left[i], 0, right[i], false) {
			pt1 += left[i]
		}
		if hasResult(left[i], 0, right[i], true) {
			pt2 += left[i]
		}
	}

	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}
