package common

import (
	"bufio"
	"os"
	"strconv"
)

type Map struct {
	Heights [][]int
}

type point struct {
	row int
	col int
}

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func (m Map) Traverse(row, col int) int {
	if row < 0 || row >= len(m.Heights) || col < 0 || col >= len(m.Heights[0]) {
		return 0
	}
	if m.Heights[row][col] != 0 {
		return 0
	}

	var queue []point
	var visited [][]bool
	for i := 0; i < len(m.Heights); i++ {
		visited = append(visited, make([]bool, len(m.Heights[0])))
	}
	queue = append(queue, point{row, col})
	visited[row][col] = true

	sum := 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		currRow := p.row
		currCol := p.col
		currHeight := m.Heights[currRow][currCol]

		for _, d := range dir {
			nextRow := currRow + d[0]
			nextCol := currCol + d[1]

			if nextRow < 0 || nextRow >= len(m.Heights) || nextCol < 0 || nextCol >= len(m.Heights[0]) {
				continue
			}

			if visited[nextRow][nextCol] {
				continue
			}

			if currHeight == 8 && m.Heights[nextRow][nextCol] == 9 {
				sum++
				visited[nextRow][nextCol] = true
				continue
			}

			if m.Heights[nextRow][nextCol] == currHeight+1 {
				queue = append(queue, point{nextRow, nextCol})
				visited[nextRow][nextCol] = true
			}
		}
	}

	return sum
}

func (m Map) TraverseDFS(row, col, currHeight int, visited [][]bool) int {
	if row < 0 || row >= len(m.Heights) || col < 0 || col >= len(m.Heights[0]) {
		return 0
	}
	if m.Heights[row][col] != currHeight {
		return 0
	}
	if visited[row][col] {
		return 0
	}
	if currHeight == 9 {
		return 1
	}

	visited[row][col] = true

	sum := 0
	for _, d := range dir {
		sum += m.TraverseDFS(row+d[0], col+d[1], currHeight+1, visited)
	}

	visited[row][col] = false

	return sum
}

func ParseInput(fileName string) (Map, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Map{}, err
	}
	defer file.Close()

	var heights [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // Reads the entire line

		var row []int
		for _, char := range line {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				row = append(row, -1)
				continue
			}

			row = append(row, height)
		}

		heights = append(heights, row)
	}

	return Map{heights}, nil
}
