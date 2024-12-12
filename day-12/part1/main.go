package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-12"
)

func main() {
	garden, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	garden.ParseValue()

	countMap := make(map[int]int)
	perimeterMap := make(map[int]int)

	for row := 0; row < garden.Height+2; row++ {
		for col := 0; col < garden.Width+1; col++ {
			if garden.ValueInt[row][col] != garden.ValueInt[row][col+1] {
				perimeterMap[garden.ValueInt[row][col]]++
				perimeterMap[garden.ValueInt[row][col+1]]++
			}
		}
	}

	for row := 0; row < garden.Height+1; row++ {
		for col := 0; col < garden.Width+2; col++ {
			if garden.ValueInt[row][col] != garden.ValueInt[row+1][col] {
				perimeterMap[garden.ValueInt[row][col]]++
				perimeterMap[garden.ValueInt[row+1][col]]++
			}
		}
	}

	for row := 1; row < garden.Height+1; row++ {
		for col := 1; col < garden.Width+1; col++ {
			countMap[garden.ValueInt[row][col]]++
		}
	}

	answer := 0
	for k, v := range countMap {
		answer += v * perimeterMap[k]
	}

	fmt.Println(answer) // 1449902
}
