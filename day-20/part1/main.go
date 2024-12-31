package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-20"
)

func main() {
	maze, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	originalDistance := maze.Solve()
	answer := 0

	for row := 1; row < maze.Height-1; row++ {
		for col := 1; col < maze.Width-1; col++ {
			if maze.Cells[row][col] != common.CellTypeWall {
				continue
			}

			maze.Cells[row][col] = common.CellTypeEmpty
			distance := maze.Solve()
			if originalDistance-distance >= 100 {
				answer++
			}
			maze.Cells[row][col] = common.CellTypeWall
		}
	}

	fmt.Println(answer) // 1363
}
