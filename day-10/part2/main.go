package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-10"
)

func main() {
	heightMap, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	answer := 0

	for row := 0; row < len(heightMap.Heights); row++ {
		for col := 0; col < len(heightMap.Heights[row]); col++ {
			visited := make([][]bool, len(heightMap.Heights))
			for i := 0; i < len(visited); i++ {
				visited[i] = make([]bool, len(heightMap.Heights[i]))
			}
			subAnswer := heightMap.TraverseDFS(row, col, 0, visited)
			if subAnswer != 0 {
				fmt.Println(row, col, subAnswer)
			}
			answer += subAnswer
		}
	}

	fmt.Println(answer) // 1875
}
