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

	answer := 0

	for row := 0; row < len(labMap.Cells); row++ {
		for col := 0; col < len(labMap.Cells[row]); col++ {
			if row == labMap.GuardRow && col == labMap.GuardCol {
				continue
			}
			if labMap.Cells[row][col] == common.CellObstacle {
				continue
			}

			cells := make([][]common.Cell, len(labMap.Cells))
			for i := range cells {
				cells[i] = make([]common.Cell, len(labMap.Cells[i]))
				copy(cells[i], labMap.Cells[i])
			}

			visited := make([][]bool, len(labMap.Cells))
			for i := range visited {
				visited[i] = make([]bool, len(labMap.Cells[i]))
			}
			newMap := common.Map{
				Cells:       cells,
				Visited:     visited,
				GuardRow:    labMap.GuardRow,
				GuardCol:    labMap.GuardCol,
				GuardDirRow: labMap.GuardDirRow,
				GuardDirCol: labMap.GuardDirCol,
			}
			newMap.Cells[row][col] = common.CellObstacle

			success := newMap.Traverse()
			if !success {
				answer++
			}
		}
	}

	fmt.Println(answer) // 1604
}
