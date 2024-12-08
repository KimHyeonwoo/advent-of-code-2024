package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-07"
)

func main() {
	equations, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	answer := 0

	for _, eq := range equations {
		if eq.IsValid(common.OptionActivateConcatenation) {
			answer += eq.Target
		}
	}

	fmt.Println(answer) // 149956401519484
}
