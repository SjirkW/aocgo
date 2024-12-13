package day7

import (
	"aoc/utils"
	"fmt"
	"math"
	"runtime"
	"strings"
	"sync"
)

func concatInts(a int, b int) int {
	if b < 10 {
		return a*10 + b
	} else if b < 100 {
		return a*100 + b
	} else if b < 1000 {
		return a*1000 + b
	} else {
		return a*int(math.Pow(10, math.Floor(math.Log10(float64(b))+1))) + b
	}
}

func hasResult(result int, current int, numbers []int, concat bool) bool {
	if len(numbers) == 0 || current > result {
		return current == result
	}

	left := numbers[0]
	rest := numbers[1:]
	if hasResult(result, current+left, rest, concat) {
		return true
	} else if hasResult(result, current*left, rest, concat) {
		return true
	} else if concat && hasResult(result, concatInts(current, left), rest, concat) {
		return true
	}

	return false
}

func Solve() {
	lines := utils.ReadInputAsLines(7, false)

	var left []int
	var right [][]int
	for _, line := range lines {
		split := strings.Split(line, ": ")

		left = append(left, utils.StringToInt(split[0]))
		right = append(right, utils.StringToIntArray(split[1], " "))
	}

	numWorkers := runtime.NumCPU()
	chunkSize := (len(left) + numWorkers - 1) / numWorkers

	var mu sync.Mutex
	var wg sync.WaitGroup

	var pt1 int
	var pt2 int
	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		end := start + chunkSize
		if end > len(left) {
			end = len(left)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			localPt1 := 0
			localPt2 := 0

			for i := start; i < end; i++ {
				if hasResult(left[i], 0, right[i], false) {
					localPt1 += left[i]
				}
				if hasResult(left[i], 0, right[i], true) {
					localPt2 += left[i]
				}
			}

			mu.Lock()
			pt1 += localPt1
			pt2 += localPt2
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()

	fmt.Println("\nDay 7")
	fmt.Println("Part 1:", pt1)
	fmt.Println("Part 2:", pt2)
}
