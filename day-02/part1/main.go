package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-02"
)

func main() {
	reports, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	var safeCount int

	for _, report := range reports {
		if report.IsSafe() {
			safeCount++
		}
	}

	fmt.Println(safeCount) // 218
}
