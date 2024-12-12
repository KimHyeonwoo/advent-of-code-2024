package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-12"
)

func main() {
	garden, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	garden.ParseValue()

	countMap := make(map[int]int)
	rowPerimeterMap := make(map[int][]common.Point)
	columnPerimeterMap := make(map[int][]common.Point)

	for col := 0; col < garden.Width+1; col++ {
		for row := 0; row < garden.Height+2; row++ {
			if garden.ValueInt[row][col] != garden.ValueInt[row][col+1] {
				columnPerimeterMap[garden.ValueInt[row][col]] = append(columnPerimeterMap[garden.ValueInt[row][col]], common.Point{Row: row, Col: col * 2})
				columnPerimeterMap[garden.ValueInt[row][col+1]] = append(columnPerimeterMap[garden.ValueInt[row][col+1]], common.Point{Row: row, Col: col*2 + 1})
			}
		}
	}

	for row := 0; row < garden.Height+1; row++ {
		for col := 0; col < garden.Width+2; col++ {
			if garden.ValueInt[row][col] != garden.ValueInt[row+1][col] {
				rowPerimeterMap[garden.ValueInt[row][col]] = append(rowPerimeterMap[garden.ValueInt[row][col]], common.Point{Row: row * 2, Col: col})
				rowPerimeterMap[garden.ValueInt[row+1][col]] = append(rowPerimeterMap[garden.ValueInt[row+1][col]], common.Point{Row: row*2 + 1, Col: col})
			}
		}
	}

	for row := 1; row < garden.Height+1; row++ {
		for col := 1; col < garden.Width+1; col++ {
			countMap[garden.ValueInt[row][col]]++
		}
	}

	costMap := make(map[int]int)

	var previousRow int
	var previousColumn int

	for k, v := range rowPerimeterMap {
		cost := 0

		for idx, point := range v {
			if idx == 0 {
				previousRow = point.Row
				previousColumn = point.Col
				cost++
				continue
			}

			if point.Row == previousRow && point.Col == previousColumn+1 {
				previousColumn = point.Col
				continue
			}

			previousRow = point.Row
			previousColumn = point.Col
			cost++
		}

		costMap[k] = cost
	}

	for k, v := range columnPerimeterMap {
		cost := 0

		for idx, point := range v {
			if idx == 0 {
				previousRow = point.Row
				previousColumn = point.Col
				cost++
				continue
			}

			if point.Row == previousRow+1 && point.Col == previousColumn {
				previousRow = point.Row
				previousColumn = point.Col
				continue
			}

			previousRow = point.Row
			previousColumn = point.Col
			cost++
		}

		costMap[k] += cost
	}

	answer := 0

	for k, v := range countMap {
		if k == 0 {
			continue
		}
		answer += v * costMap[k]
	}

	fmt.Println(answer) // 908042
}
