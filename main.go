package main

import (
	"aoc/days/day1"
	"aoc/days/day2"
	"aoc/days/day3"
	"aoc/days/day4"
	"aoc/days/day5"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day1.Solve()
	day2.Solve()
	day3.Solve()
	day4.Solve()
	day5.Solve()

	fmt.Printf("Execution time: %v\n", time.Since(start))
}
