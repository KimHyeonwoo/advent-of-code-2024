package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-09"
)

func main() {
	disk, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	disk.Compact()

	fmt.Println(disk.ComputeChecksum()) // 6344673854800
}
