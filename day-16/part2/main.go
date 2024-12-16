package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-16"
)

func main() {
	maze, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	fmt.Println(maze.Solve()) // 479
}
