package common

import (
	"bufio"
	"os"
	"strings"
)

type CellType int

const (
	CellTypeWall CellType = iota
	CellTypeEmpty
	CellTypeBox
	CellTypeRobot
)

func (c CellType) Pushable() bool {
	return c == CellTypeRobot || c == CellTypeBox
}

type Movement uint8

const (
	MovementUp    Movement = '^'
	MovementDown  Movement = 'v'
	MovementLeft  Movement = '<'
	MovementRight Movement = '>'
)

type Warehouse struct {
	Cells    [][]CellType
	Width    int
	Height   int
	RobotRow int
	RobotCol int
}

func (w *Warehouse) Up() bool {
	row := w.RobotRow - 1
	col := w.RobotCol

	for w.Cells[row][col].Pushable() {
		row--
	}

	if w.Cells[row][col] == CellTypeWall {
		return false
	}

	for row < w.RobotRow {
		w.Cells[row][col] = w.Cells[row+1][col]
		row++
	}

	w.Cells[row][col] = CellTypeEmpty

	w.RobotRow--

	return true
}

func (w *Warehouse) Down() bool {
	row := w.RobotRow + 1
	col := w.RobotCol

	for w.Cells[row][col].Pushable() {
		row++
	}

	if w.Cells[row][col] == CellTypeWall {
		return false
	}

	for row > w.RobotRow {
		w.Cells[row][col] = w.Cells[row-1][col]
		row--
	}

	w.Cells[row][col] = CellTypeEmpty

	w.RobotRow++

	return true
}

func (w *Warehouse) Left() bool {
	row := w.RobotRow
	col := w.RobotCol - 1

	for w.Cells[row][col].Pushable() {
		col--
	}

	if w.Cells[row][col] == CellTypeWall {
		return false
	}

	for col < w.RobotCol {
		w.Cells[row][col] = w.Cells[row][col+1]
		col++
	}

	w.Cells[row][col] = CellTypeEmpty

	w.RobotCol--

	return true
}

func (w *Warehouse) Right() bool {
	row := w.RobotRow
	col := w.RobotCol + 1

	for w.Cells[row][col].Pushable() {
		col++
	}

	if w.Cells[row][col] == CellTypeWall {
		return false
	}

	for col > w.RobotCol {
		w.Cells[row][col] = w.Cells[row][col-1]
		col--
	}

	w.Cells[row][col] = CellTypeEmpty

	w.RobotCol++

	return true
}

func (w *Warehouse) GPS() int {
	var gps int

	for row, rowCells := range w.Cells {
		for col, cell := range rowCells {
			if cell == CellTypeBox {
				gps += row*100 + col
			}
		}
	}

	return gps
}

func ParseInput(fileName string) (Warehouse, []Movement, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Warehouse{}, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cells := make([][]CellType, 0)
	var width, height, robotRow, robotCol int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			break
		}

		row := make([]CellType, 0)

		for _, c := range line {
			if c == '#' {
				row = append(row, CellTypeWall)
			} else if c == '.' {
				row = append(row, CellTypeEmpty)
			} else if c == 'O' {
				row = append(row, CellTypeBox)
			} else if c == '@' {
				robotRow = len(cells)
				robotCol = len(row)
				row = append(row, CellTypeRobot)
			}
		}

		cells = append(cells, row)
		width = len(row)
		height++
	}

	movements := make([]Movement, 0)
	for scanner.Scan() {
		movementsStr := scanner.Text()
		for _, c := range movementsStr {
			movements = append(movements, Movement(c))
		}
	}

	return Warehouse{
		Cells:    cells,
		Width:    width,
		Height:   height,
		RobotRow: robotRow,
		RobotCol: robotCol,
	}, movements, nil
}
