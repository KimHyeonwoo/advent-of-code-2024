package common

import (
	"bufio"
	"os"
)

type Cell int

const (
	CellEmpty Cell = iota
	CellObstacle
)

type Map struct {
	Cells       [][]Cell
	Visited     [][]bool
	GuardRow    int
	GuardCol    int
	GuardDirRow int
	GuardDirCol int
}

func (m *Map) Traverse() bool {
	counter := 0
	counterLimit := 4*len(m.Cells)*len(m.Cells[0]) + 1
	for m.MoveGuard() && counter < counterLimit {
		counter++
	}

	return counter < counterLimit
}

func (m *Map) MoveGuard() bool {
	m.Visited[m.GuardRow][m.GuardCol] = true

	nextRow := m.GuardRow + m.GuardDirRow
	nextCol := m.GuardCol + m.GuardDirCol

	if nextRow < 0 || nextRow >= len(m.Cells) || nextCol < 0 || nextCol >= len(m.Cells[nextRow]) {
		return false
	}

	if m.Cells[nextRow][nextCol] == CellObstacle {
		m.changeGuardDirection()
		return true
	}

	m.GuardRow = nextRow
	m.GuardCol = nextCol
	return true
}

func (m *Map) changeGuardDirection() {
	if m.GuardDirRow == -1 && m.GuardDirCol == 0 {
		m.GuardDirRow = 0
		m.GuardDirCol = 1
		return
	}

	if m.GuardDirRow == 0 && m.GuardDirCol == 1 {
		m.GuardDirRow = 1
		m.GuardDirCol = 0
		return
	}

	if m.GuardDirRow == 1 && m.GuardDirCol == 0 {
		m.GuardDirRow = 0
		m.GuardDirCol = -1
		return
	}

	if m.GuardDirRow == 0 && m.GuardDirCol == -1 {
		m.GuardDirRow = -1
		m.GuardDirCol = 0
		return
	}
}

func (m *Map) CountVisited() int {
	var count int
	for _, row := range m.Visited {
		for _, visited := range row {
			if visited {
				count++
			}
		}
	}
	return count
}

func ParseInput(fileName string) (Map, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Map{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cells [][]Cell
	var guardRow, guardCol int
	for scanner.Scan() {
		line := scanner.Text()
		var row []Cell
		for col, char := range line {
			switch char {
			case '.':
				row = append(row, CellEmpty)
			case '#':
				row = append(row, CellObstacle)
			case '^':
				row = append(row, CellEmpty)
				guardRow = len(cells)
				guardCol = col
			}
		}
		cells = append(cells, row)
	}

	visited := make([][]bool, len(cells))
	for i := range visited {
		visited[i] = make([]bool, len(cells[i]))
	}

	return Map{
		Cells:       cells,
		Visited:     visited,
		GuardRow:    guardRow,
		GuardCol:    guardCol,
		GuardDirRow: -1,
		GuardDirCol: 0,
	}, nil
}
