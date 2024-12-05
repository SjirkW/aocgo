package day3

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strings"
)

func run(doEnableCheck bool) int {
	lines := utils.ReadInputAsLines(3, false)

	total := 0
	enabled := true
	for _, line := range lines {
		leftSide := ""
		rightSide := line

		for len(rightSide) > 0 {
			leftSide += string(rightSide[0])
			rightSide = rightSide[1:]

			if doEnableCheck {
				if strings.HasSuffix(leftSide, "do()") {
					enabled = true
				} else if strings.HasSuffix(leftSide, "don't()") {
					enabled = false
				}
			}

			if enabled && strings.HasSuffix(leftSide, "mul") {
				re := regexp.MustCompile(`^\((\d+),(\d+)\)`)
				matches := re.FindStringSubmatch(rightSide)

				if len(matches) > 2 {
					x := utils.StringToInt(matches[1])
					y := utils.StringToInt(matches[2])
					total += x * y
				}
			}
		}
	}

	return total
}

func Solve() {
	fmt.Println("\nDay 3")
	fmt.Printf("Part 1: %d\n", run(false))
	fmt.Printf("Part 2: %d\n", run(true))
}
