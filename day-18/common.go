package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	Row int
	Col int
}

type Memory struct {
	Width     int
	Height    int
	Cells     [][]bool
	Distances [][]int
}

func (m *Memory) BFS() {
	var queue []Point
	queue = append(queue, Point{Row: 0, Col: 0})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, delta := range []Point{{Row: -1, Col: 0}, {Row: 1, Col: 0}, {Row: 0, Col: -1}, {Row: 0, Col: 1}} {
			newRow := current.Row + delta.Row
			newCol := current.Col + delta.Col
			if newRow < 0 || newRow >= m.Height || newCol < 0 || newCol >= m.Width {
				continue
			}
			if m.Cells[newRow][newCol] {
				continue
			}
			if m.Distances[newRow][newCol] != -1 {
				continue
			}
			m.Distances[newRow][newCol] = m.Distances[current.Row][current.Col] + 1
			queue = append(queue, Point{Row: newRow, Col: newCol})
		}
	}
}

func ParseInput(fileName string, limit int) (Memory, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Memory{}, err
	}
	defer file.Close()

	width := 71
	height := 71
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	distances := make([][]int, height)
	for i := range distances {
		distances[i] = make([]int, width)
		for j := range distances[i] {
			distances[i][j] = -1
		}
	}
	distances[0][0] = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && limit != 0 {
		line := scanner.Text()
		segments := strings.Split(line, ",")
		row, _ := strconv.Atoi(segments[0])
		col, _ := strconv.Atoi(segments[1])
		cells[row][col] = true
		limit--
	}

	return Memory{Width: width, Height: height, Cells: cells, Distances: distances}, nil
}
