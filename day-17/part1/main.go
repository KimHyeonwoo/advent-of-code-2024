package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-17"
)

func main() {
	computer, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	results := computer.Execute()

	for i := range results {
		fmt.Print(results[i])
		if i < len(results)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println() // 3,5,0,1,5,1,5,1,0
}
