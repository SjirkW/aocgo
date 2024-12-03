package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInputAsLines(day int, isTest bool) []string {
	path := ""
	if isTest {
		path = fmt.Sprintf("tests/day%d.txt", day)
	} else {
		path = fmt.Sprintf("inputs/day%d.txt", day)
	}

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x - 0
}
