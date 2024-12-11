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

	stonesMap := make(map[common.Stone]int)
	for _, stone := range stones {
		stonesMap[stone]++
	}

	for range 75 {
		newStonesMap := make(map[common.Stone]int)
		for stone, count := range stonesMap {
			processedStones := stone.Process()
			for _, processedStone := range processedStones {
				newStonesMap[processedStone] += count
			}
		}
		stonesMap = newStonesMap
	}

	var sum int
	for _, count := range stonesMap {
		sum += count
	}

	fmt.Println(sum) // 252442982856820
}
