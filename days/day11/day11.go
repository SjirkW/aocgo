package day11

import (
	"aoc/utils"
	"fmt"
	"math"
)

func GetDigitAmount(number int) int {
	amount := 0
	for number > 0 {
		number = number / 10
		amount++
	}

	return amount
}

func DoLoop(numbers []int) {
	blinks := 25

	for i := 0; i < blinks; i++ {
		// empty array
		newNumbers := make([]int, 0)
		for i := 0; i < len(numbers); i++ {
			number := numbers[i]

			if number == 0 {
				newNumbers = append(newNumbers, 1)
			} else {
				digitAmount := GetDigitAmount(number)
				if digitAmount%2 == 0 {
					// Split the number in half
					mid := utils.DigitCount(number) / 2
					divisor := int(math.Pow(10, float64(mid)))
					left := number / divisor
					right := number % divisor
					newNumbers = append(newNumbers, left, right)
				} else {
					newNumbers = append(newNumbers, number*2024)
				}
			}
		}

		numbers = newNumbers
	}

	fmt.Println("\nDay 11")
	fmt.Println("Part 1:", len(numbers))
}

func Solve() {
	lines := utils.ReadInputAsLines(11, false)

	numbers := utils.StringToIntArray(lines[0], " ")

	DoLoop(numbers)
}
