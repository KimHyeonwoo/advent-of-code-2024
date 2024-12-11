package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-11"
)

func main() {
	stones, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	for range 25 {
		newStones := make([]common.Stone, 0, len(stones)*2)
		for _, stone := range stones {
			newStones = append(newStones, stone.Process()...)
		}
		stones = newStones
	}

	fmt.Println(len(stones)) // 213625
}
