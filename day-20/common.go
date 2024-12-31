package common

import (
	"bufio"
	"os"
)

type Point struct {
	Row int
	Col int
}

type CellType int

const (
	CellTypeEmpty CellType = iota
	CellTypeWall
)

type Maze struct {
	Width    int
	Height   int
	Cells    [][]CellType
	StartRow int
	StartCol int
	EndRow   int
	EndCol   int
}

func (m *Maze) Solve() int {
	// Find distance from start to end
	// Return the distance
	visited := make(map[Point]struct{})
	distances := make(map[Point]int)

	queue := []Point{{m.StartRow, m.StartCol}}
	visited[Point{m.StartRow, m.StartCol}] = struct{}{}
	distances[Point{m.StartRow, m.StartCol}] = 0

	directions := [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			np := Point{p.Row + d[0], p.Col + d[1]}
			if _, ok := visited[np]; ok {
				continue
			}
			if np.Row < 0 || np.Row >= m.Height || np.Col < 0 || np.Col >= m.Width {
				continue
			}
			if m.Cells[np.Row][np.Col] == CellTypeWall {
				continue
			}

			visited[np] = struct{}{}
			distances[np] = distances[p] + 1
			queue = append(queue, np)
		}
	}

	return distances[Point{m.EndRow, m.EndCol}]
}

func ParseInput(fileName string) (Maze, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Maze{}, err
	}
	defer file.Close()

	var maze Maze
	var startRow, startCol, endRow, endCol int
	cells := make([][]CellType, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]CellType, 0)
		for _, c := range line {
			var cellType CellType
			switch c {
			case '#':
				cellType = CellTypeWall
			case '.':
				cellType = CellTypeEmpty
			case 'S':
				cellType = CellTypeEmpty
				startRow = len(cells)
				startCol = len(row)
			case 'E':
				cellType = CellTypeEmpty
				endRow = len(cells)
				endCol = len(row)
			}
			row = append(row, cellType)
		}
		cells = append(cells, row)
	}

	maze.Width = len(cells[0])
	maze.Height = len(cells)
	maze.Cells = cells

	maze.StartRow = startRow
	maze.StartCol = startCol
	maze.EndRow = endRow
	maze.EndCol = endCol

	return maze, nil
}
