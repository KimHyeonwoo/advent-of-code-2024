package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-19"
)

func main() {
	candidates, targets, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	count := 0
	for _, target := range targets {
		if c := target.Construct(candidates); c != 0 {
			count++
		}
	}

	fmt.Println(count) // 327
}
