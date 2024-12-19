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
		if target.CanConstruct(candidates) {
			count++
		}
	}

	fmt.Println(count) // 327
}
