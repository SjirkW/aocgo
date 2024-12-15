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

func GetMemoryKey(number int, blinks int) int {
	return number*1000 + blinks
}

var stoneAmount = make(map[int]int)

func GetStoneAmount(number int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	memoryKey := GetMemoryKey(number, blinks)
	if val, ok := stoneAmount[memoryKey]; ok {
		return val
	}

	value := 0
	if number == 0 {
		value = GetStoneAmount(1, blinks-1)
	} else if GetDigitAmount(number)%2 == 0 {
		// Split the number in half
		mid := utils.DigitCount(number) / 2
		divisor := int(math.Pow(10, float64(mid)))
		left := number / divisor
		right := number % divisor
		value = GetStoneAmount(left, blinks-1) + GetStoneAmount(right, blinks-1)
	} else {
		value = GetStoneAmount(number*2024, blinks-1)
	}
	stoneAmount[memoryKey] = value

	return value
}

func Solve() {
	lines := utils.ReadInputAsLines(11, false)

	numbers := utils.StringToIntArray(lines[0], " ")

	fmt.Println("\nDay 11")
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += GetStoneAmount(numbers[i], 25)
	}
	fmt.Println("Part 1:", total)

	total2 := 0
	for i := 0; i < len(numbers); i++ {
		total2 += GetStoneAmount(numbers[i], 75)
	}
	fmt.Println("Part 2:", total2)
}
