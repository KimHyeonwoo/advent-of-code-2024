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

	answer := 0

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			if board.HasTwoMASInXShape(x, y) {
				answer++
			}
		}
	}

	fmt.Println(answer) // 1822
}
