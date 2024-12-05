package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-04"
)

func main() {
	board, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	directions := [8][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	var answer int

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			for _, direction := range directions {
				if board.HasXMAS(x, y, direction[0], direction[1]) {
					answer++
				}
			}
		}
	}

	fmt.Println(answer) // 2401
}
