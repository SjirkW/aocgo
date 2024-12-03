package days

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type day3 struct {
	result int
}

func (d *day3) run(checkEnable bool) int {
	lines := utils.ReadInputAsLines(3, false)

	total := 0
	enabled := true
	for _, line := range lines {
		leftSide := ""
		rightSide := line

		for len(rightSide) > 0 {
			leftSide += string(rightSide[0])
			rightSide = rightSide[1:]

			if checkEnable {
				if strings.HasSuffix(leftSide, "do()") {
					enabled = true
				} else if strings.HasSuffix(leftSide, "don't()") {
					enabled = false
				}
			}

			if enabled && strings.HasSuffix(leftSide, "mul") {
				re := regexp.MustCompile(`\((\d+),(\d+)\)`)
				matches := re.FindStringSubmatch(rightSide)
				matchIsAtStart := re.FindStringSubmatchIndex(rightSide)[0] == 0

				if matchIsAtStart && len(matches) > 2 {
					x, _ := strconv.Atoi(matches[1])
					y, _ := strconv.Atoi(matches[2])
					total += x * y
				}
			}
		}
	}

	return total
}

func Solve3() {
	pt1 := day3{}
	pt1.run(false)
	fmt.Printf("Part 1 Result: %d\n", pt1.result)

	pt2 := day3{}
	pt2.run(true)
	fmt.Printf("Part 2 Result: %d\n", pt2.result)
}
