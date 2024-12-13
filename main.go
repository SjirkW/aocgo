package main

import (
	"aoc/days/day1"
	"aoc/days/day10"
	"aoc/days/day2"
	"aoc/days/day3"
	"aoc/days/day4"
	"aoc/days/day5"
	"aoc/days/day6"
	"aoc/days/day7"
	"aoc/days/day8"
	"aoc/days/day9"
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
	day6.Solve()
	day7.Solve()
	day8.Solve()
	day9.Solve()
	day10.Solve()
	// day11.Solve()

	fmt.Printf("Execution time: %v\n", time.Since(start))
}
