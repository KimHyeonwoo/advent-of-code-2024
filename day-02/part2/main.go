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
		reportLength := report.Length()

		for idx := 0; idx < reportLength; idx++ {
			newReport := report.RemoveByIndex(idx)
			if newReport.IsSafe() {
				safeCount++
				break
			}
		}
	}

	fmt.Println(safeCount) // 290
}
