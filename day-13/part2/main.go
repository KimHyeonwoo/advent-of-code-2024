package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-13"
)

func main() {
	machines, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	var totalCost int64

	for _, machine := range machines {
		cost, ok := machine.GetActualPrize()
		if ok {
			totalCost += cost
		}
	}

	fmt.Println(totalCost) //83232379451012
}
