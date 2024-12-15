package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-15"
)

func main() {
	warehouse, movements, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	for _, movement := range movements {
		switch movement {
		case common.MovementUp:
			warehouse.Up()
		case common.MovementDown:
			warehouse.Down()
		case common.MovementLeft:
			warehouse.Left()
		case common.MovementRight:
			warehouse.Right()
		}
	}

	fmt.Println(warehouse.GPS()) // 1487337
}
