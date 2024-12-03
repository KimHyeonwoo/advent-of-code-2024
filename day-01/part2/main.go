package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-01"
)

func main() {
	left, right, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	leftOccurrenceCount := make(map[int]int)
	rightOccurrenceCount := make(map[int]int)

	for _, l := range left {
		leftOccurrenceCount[l]++
	}

	for _, r := range right {
		rightOccurrenceCount[r]++
	}

	var answer int

	for l, count := range leftOccurrenceCount {
		answer += l * count * rightOccurrenceCount[l]
	}

	fmt.Println(answer) // 20351745
}
