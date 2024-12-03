package days

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day3(checkEnable bool) {
	lines := utils.ReadInputAsLines(3, false)

	total := 0
	enabled := true
	for _, line := range lines {
		leftSide := ""
		rightSide := line

		for len(rightSide) > 0 {
			char := rightSide[0]
			rightSide = rightSide[1:]
			leftSide += string(char)

			if strings.HasSuffix(leftSide, "do()") {
				enabled = true
			} else if checkEnable && strings.HasSuffix(leftSide, "don't()") {
				enabled = false
			}

			if strings.HasSuffix(leftSide, "mul") {
				re := regexp.MustCompile(`\((\d+),(\d+)\)`)
				matches := re.FindStringSubmatch(rightSide)
				ix := re.FindStringSubmatchIndex(rightSide)

				if len(matches) > 2 && enabled && ix[0] == 0 {
					x, _ := strconv.Atoi(matches[1])
					y, _ := strconv.Atoi(matches[2])
					total += x * y
				}
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func Solve3() {
	day3(false)
	day3(true)
}
