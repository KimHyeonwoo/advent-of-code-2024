package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-08"
)

func main() {
	cityMap, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	visited := make([][]bool, cityMap.Height)
	for i := range visited {
		visited[i] = make([]bool, cityMap.Width)
	}

	for _, points := range cityMap.Antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				newX := points[i].Row*2 - points[j].Row
				newY := points[i].Col*2 - points[j].Col
				if newX >= 0 && newX < cityMap.Width && newY >= 0 && newY < cityMap.Height {
					visited[newY][newX] = true
				}

				newX = points[j].Row*2 - points[i].Row
				newY = points[j].Col*2 - points[i].Col
				if newX >= 0 && newX < cityMap.Width && newY >= 0 && newY < cityMap.Height {
					visited[newY][newX] = true
				}
			}
		}
	}

	var answer int
	for _, row := range visited {
		for _, cell := range row {
			if cell {
				answer++
			}
		}
	}

	fmt.Println(answer) // 278
}
