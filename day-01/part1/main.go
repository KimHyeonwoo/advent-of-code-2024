package main

import (
	"fmt"
	"sort"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-01"
)

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	left, right, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	var answer int

	for i := 0; i < len(left); i++ {
		answer += distance(left[i], right[i])
	}

	fmt.Println(answer) // 1579939
}
