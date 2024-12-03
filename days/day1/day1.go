package day1

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve() int {
	fmt.Println("Day 1")
	lines := utils.ReadInputAsLines(1, false)

	var left, right []int
	amountMap := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		numbers := strings.Fields(lines[i])

		num1, _ := strconv.Atoi(numbers[0])
		left = append(left, num1)
		num2, _ := strconv.Atoi(numbers[1])
		right = append(right, num2)

		val, ok := amountMap[num2]
		if ok {
			amountMap[num2] = val + 1
		} else {
			amountMap[num2] = 1
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i := 0; i < len(left); i++ {
		total += utils.Abs(left[i] - right[i])
	}

	fmt.Printf("Part 1: %d\n", total)

	total2 := 0
	for i := 0; i < len(left); i++ {
		val, ok := amountMap[left[i]]
		if ok {
			total2 += left[i] * val
		}
	}

	fmt.Printf("Part 2: %d\n", total2)

	return total
}
