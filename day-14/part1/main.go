package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-14"
)

const (
	TopLeft = iota
	TopRight
	BottomLeft
	BottomRight
	Undefined
)

func positionToSector(row, col, width, height int) int {
	centerRow := (height - 1) / 2
	centerCol := (width - 1) / 2

	if row < centerRow && col < centerCol {
		return TopLeft
	}

	if row < centerRow && col > centerCol {
		return TopRight
	}

	if row > centerRow && col < centerCol {
		return BottomLeft
	}

	if row > centerRow && col > centerCol {
		return BottomRight
	}

	return Undefined
}

func main() {
	robots, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	width := 101
	height := 103

	for range 100 {
		for r := range robots {
			robots[r].Move(width, height)
		}
	}

	robotCount := make(map[int]int)

	for r := range robots {
		sector := positionToSector(robots[r].PositionRow, robots[r].PositionCol, width, height)
		robotCount[sector]++
	}

	var answer = 1

	for key, value := range robotCount {
		if key == Undefined {
			continue
		}

		answer *= value
	}

	fmt.Println(answer) // 217328832
}
