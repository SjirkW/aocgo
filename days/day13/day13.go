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
	ButtonA   Point
	ButtonB   Point
	Goal      Point
	Steps     int
	isSwapped bool
}

func parsePoint(line string, delimiter string) Point {
	parts := strings.Split(line, ", ")

	xPart := strings.Split(parts[0], "X"+delimiter)[1]
	x := utils.StringToInt(xPart)

	yPart := strings.Split(parts[1], "Y"+delimiter)[1]
	y := utils.StringToInt(yPart)
	return Point{X: x, Y: y}
}

func hasResultForY(game *Game, aPresses int) bool {
	goalY := game.Goal.Y

	leftOver := goalY - aPresses*game.ButtonA.Y

	return leftOver%game.ButtonB.Y == 0
}

func hasResult(game *Game, startX int) (int, int) {
	do := true
	x := game.Goal.X

	currentA := startX
	leftOver := x - (currentA * game.ButtonA.X)

	for do {
		if (leftOver%game.ButtonB.X) == 0 && hasResultForY(game, currentA) {
			return currentA, leftOver / game.ButtonB.X
		} else {
			leftOver += game.ButtonA.X
			currentA--
		}

		if leftOver > x || currentA < 0 {
			do = false
		}
	}
	return -1, -1
}

func getMinScorePossible(game Game) int {
	resA, resB := hasResult(&game, game.Goal.X/game.ButtonA.X)
	if resA != -1 {
		return resA*3 + resB
	}

	return 0
}

func Solve() {
	lines := utils.ReadInputAsLines(13, true)

	fmt.Println("\nDay 13")

	gamesToPlay := make([]Game, 0)

	game := Game{
		ButtonA: Point{X: 1, Y: 1},
		ButtonB: Point{X: 1, Y: 1},
		Goal:    Point{X: 1, Y: 1},
		Steps:   200,
	}

	for _, line := range lines {
		if line == "" {
			game = Game{
				ButtonA: Point{X: 1, Y: 1},
				ButtonB: Point{X: 1, Y: 1},
				Goal:    Point{X: 1, Y: 1},
				Steps:   200,
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

	total := 0
	for _, game := range gamesToPlay {
		total += getMinScorePossible(game)
	}

	fmt.Println(total)
}
