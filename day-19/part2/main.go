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

	var answer int

	for _, target := range targets {
		answer += target.Construct(candidates)
	}

	fmt.Println(answer) // 772696486795255
}
