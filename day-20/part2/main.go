package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-20"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	maze, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	_ = maze.Solve()

	answer := 0

	multipliers := [4][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	type Points struct {
		StartRow int
		StartCol int
		EndRow   int
		EndCol   int
	}

	visited := make(map[Points]struct{})

	for currRow := 1; currRow < maze.Height-1; currRow++ {
		for currCol := 1; currCol < maze.Width-1; currCol++ {
			if maze.Cells[currRow][currCol] == common.CellTypeWall {
				continue
			}
			for cheatLength := 2; cheatLength <= 20; cheatLength++ {
				for dx := 0; dx <= cheatLength; dx++ {
					dy := cheatLength - abs(dx)
					for _, multiplier := range multipliers {
						row := currRow + dy*multiplier[0]
						col := currCol + dx*multiplier[1]
						if row < 0 || row >= maze.Height || col < 0 || col >= maze.Width {
							continue
						}
						if maze.Cells[row][col] == common.CellTypeWall {
							continue
						}
						if _, ok := visited[Points{StartRow: currRow, StartCol: currCol, EndRow: row, EndCol: col}]; ok {
							continue
						}
						cheatGain := maze.Distances[common.Point{Row: row, Col: col}] - maze.Distances[common.Point{Row: currRow, Col: currCol}] - cheatLength
						if cheatGain >= 100 {
							answer++
							visited[Points{StartRow: currRow, StartCol: currCol, EndRow: row, EndCol: col}] = struct{}{}
						}
					}
				}
			}
		}
	}

	fmt.Println(answer) // 1007186
}
