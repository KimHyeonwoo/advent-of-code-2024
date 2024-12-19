package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-18"
)

func main() {
	memory, err := common.ParseInput("input", 1024)
	if err != nil {
		panic(err)
	}

	memory.BFS()

	fmt.Println(memory.Distances[memory.Height-1][memory.Width-1]) // 312
}
