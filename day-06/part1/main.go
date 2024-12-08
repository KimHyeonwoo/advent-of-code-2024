package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-06"
)

func main() {
	labMap, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	labMap.Traverse()

	fmt.Println(labMap.CountVisited()) //4559
}
