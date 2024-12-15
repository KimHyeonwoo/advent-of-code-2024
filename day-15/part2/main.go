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

	widerWarehouse := warehouse.GetWiderWarehouse()

	for _, movement := range movements {
		switch movement {
		case common.MovementUp:
			widerWarehouse.Up()
		case common.MovementDown:
			widerWarehouse.Down()
		case common.MovementLeft:
			widerWarehouse.Left()
		case common.MovementRight:
			widerWarehouse.Right()
		}
	}

	fmt.Println(widerWarehouse.GPS()) // 1521952
}
