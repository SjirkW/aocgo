package days

import (
	"aoc/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

// SolvePart1 solves the first part of the problem
func Solve() int {
	lines := utils.ReadInputAsLines(1, false)

	var left []int
	var right []int
	amountMap := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		// Split on space
		numbers := strings.Fields(lines[i])

		num1, err1 := strconv.Atoi(numbers[0])
		if err1 != nil {
			log.Fatalf("Failed to convert %s to int: %v", numbers[0], err1)
		}

		num2, err2 := strconv.Atoi(numbers[1])
		if err2 != nil {
			log.Fatalf("Failed to convert %s to int: %v", numbers[1], err2)
		}

		left = append(left, num1)

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

	fmt.Printf("Part 1 Result: %d\n", total)

	total2 := 0
	for i := 0; i < len(left); i++ {
		val, ok := amountMap[left[i]]
		if ok {
			total2 += left[i] * val
		}
	}

	fmt.Printf("Part 2 Result: %d\n", total2)

	return total
}
