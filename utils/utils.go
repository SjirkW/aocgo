package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputAsLines(day int, isTest bool) []string {
	path := ""
	if isTest {
		path = fmt.Sprintf("test/day%d.txt", day)
	} else {
		path = fmt.Sprintf("input/day%d.txt", day)
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

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x - 0
}

func PrintGrid(grid [][]string) {
	fmt.Println()
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func StringToIntArray(input string, splitChar string) []int {
	// Split the string on spaces
	strValues := strings.Split(input, splitChar)

	// Convert each string to an integer
	intValues := make([]int, len(strValues))
	for i, str := range strValues {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error converting to int:", err)
		}
		intValues[i] = num
	}

	return intValues
}
