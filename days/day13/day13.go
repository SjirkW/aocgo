package day13

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Game struct {
	ButtonA Point
	ButtonB Point
	Goal    Point
}

func parsePoint(line string, delimiter string) Point {
	parts := strings.Split(line, ", ")

	xPart := strings.Split(parts[0], "X"+delimiter)[1]
	x := utils.StringToInt(xPart)

	yPart := strings.Split(parts[1], "Y"+delimiter)[1]
	y := utils.StringToInt(yPart)
	return Point{X: x, Y: y}
}

func Solve() {
	lines := utils.ReadInputAsLines(13, true)

	fmt.Println("\nDay 13")

	gamesToPlay := make([]Game, 0)

	game := Game{
		ButtonA: Point{X: 1, Y: 1},
		ButtonB: Point{X: 1, Y: 1},
		Goal:    Point{X: 1, Y: 1},
	}

	for _, line := range lines {
		if line == "" {
			gamesToPlay = append(gamesToPlay, game)
			game = Game{
				ButtonA: Point{X: 1, Y: 1},
				ButtonB: Point{X: 1, Y: 1},
				Goal:    Point{X: 1, Y: 1},
			}
			continue
		}
		if strings.HasPrefix(line, "Button A: ") {
			game.ButtonA = parsePoint(strings.Split(line, "Button A: ")[1], "+")
		} else if strings.HasPrefix(line, "Button B: ") {
			game.ButtonB = parsePoint(strings.Split(line, "Button B: ")[1], "+")
		} else if strings.HasPrefix(line, "Prize: ") {
			game.Goal = parsePoint(strings.Split(line, "Prize: ")[1], "=")
			gamesToPlay = append(gamesToPlay, game)
		}
	}

	fmt.Println("Games to play:", len(gamesToPlay))
}
