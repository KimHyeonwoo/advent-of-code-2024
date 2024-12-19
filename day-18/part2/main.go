package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-18"
)

func main() {
	limit := 1024

	for {
		memory, err := common.ParseInput("input", limit)
		if err != nil {
			panic(err)
		}

		memory.BFS()

		if memory.Distances[memory.Height-1][memory.Width-1] == -1 {
			fmt.Println(limit) // 3038 -> 28,26
			break
		}

		limit++
	}
}
