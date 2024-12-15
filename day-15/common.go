package common

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type CellType int

const (
	CellTypeWall CellType = iota
	CellTypeEmpty
	CellTypeBox
	CellTypeRobot
	CellTypeBoxLeft
	CellTypeBoxRight
)

func (c CellType) Pushable() bool {
	return c == CellTypeRobot || c == CellTypeBox || c == CellTypeBoxLeft || c == CellTypeBoxRight
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

func (w *Warehouse) GetWiderWarehouse() WiderWarehouse {
	cells := make([][]CellType, w.Height)
	for i := 0; i < w.Height; i++ {
		cells[i] = make([]CellType, w.Width*2)
	}

	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if w.Cells[i][j] == CellTypeWall {
				cells[i][j*2] = CellTypeWall
				cells[i][j*2+1] = CellTypeWall
			} else if w.Cells[i][j] == CellTypeEmpty {
				cells[i][j*2] = CellTypeEmpty
				cells[i][j*2+1] = CellTypeEmpty
			} else if w.Cells[i][j] == CellTypeBox {
				cells[i][j*2] = CellTypeBoxLeft
				cells[i][j*2+1] = CellTypeBoxRight
			} else {
				cells[i][j*2] = CellTypeRobot
				cells[i][j*2+1] = CellTypeEmpty
			}
		}
	}

	return WiderWarehouse{
		Cells:    cells,
		Width:    w.Width * 2,
		Height:   w.Height,
		RobotRow: w.RobotRow,
		RobotCol: w.RobotCol * 2,
	}
}

type WiderWarehouse struct {
	Cells    [][]CellType
	Width    int
	Height   int
	RobotRow int
	RobotCol int
}

type Point struct {
	Row int
	Col int
}

type PointsByRow []Point

func (p PointsByRow) Len() int {
	return len(p)
}

func (p PointsByRow) Less(i, j int) bool {
	return p[i].Row < p[j].Row
}

func (p PointsByRow) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type PointsByRowDesc []Point

func (p PointsByRowDesc) Len() int {
	return len(p)
}

func (p PointsByRowDesc) Less(i, j int) bool {
	return p[i].Row > p[j].Row
}

func (p PointsByRowDesc) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (w *WiderWarehouse) Up() bool {
	row := w.RobotRow - 1
	affectedColumns := make(map[int]struct{})
	affectedColumns[w.RobotCol] = struct{}{}

	affected := make([]Point, 0)
	affected = append(affected, Point{w.RobotRow, w.RobotCol})

	for {
		allEmpty := true
		for col := range affectedColumns {
			if w.Cells[row][col] == CellTypeWall {
				return false
			}
			if w.Cells[row][col] != CellTypeEmpty {
				allEmpty = false
			}
		}

		if allEmpty {
			sort.Sort(PointsByRow(affected))
			for _, point := range affected {
				w.Cells[point.Row-1][point.Col] = w.Cells[point.Row][point.Col]
				w.Cells[point.Row][point.Col] = CellTypeEmpty
			}

			w.RobotRow--
			return true
		}

		nextAffectedColumns := make(map[int]struct{})
		for col := range affectedColumns {
			if w.Cells[row][col] == CellTypeBoxRight {
				nextAffectedColumns[col-1] = struct{}{}
				nextAffectedColumns[col] = struct{}{}
			}

			if w.Cells[row][col] == CellTypeBoxLeft {
				nextAffectedColumns[col+1] = struct{}{}
				nextAffectedColumns[col] = struct{}{}
			}
		}

		for col := range nextAffectedColumns {
			affected = append(affected, Point{row, col})
		}

		affectedColumns = nextAffectedColumns

		row--
	}
}

func (w *WiderWarehouse) Down() bool {
	row := w.RobotRow + 1
	affectedColumns := make(map[int]struct{})
	affectedColumns[w.RobotCol] = struct{}{}

	affected := make([]Point, 0)
	affected = append(affected, Point{w.RobotRow, w.RobotCol})

	for {
		allEmpty := true
		for col := range affectedColumns {
			if w.Cells[row][col] == CellTypeWall {
				return false
			}
			if w.Cells[row][col] != CellTypeEmpty {
				allEmpty = false
			}
		}

		if allEmpty {
			sort.Sort(PointsByRowDesc(affected))
			for _, point := range affected {
				w.Cells[point.Row+1][point.Col] = w.Cells[point.Row][point.Col]
				w.Cells[point.Row][point.Col] = CellTypeEmpty
			}

			w.RobotRow++
			return true
		}

		nextAffectedColumns := make(map[int]struct{})
		for col := range affectedColumns {
			if w.Cells[row][col] == CellTypeBoxRight {
				nextAffectedColumns[col-1] = struct{}{}
				nextAffectedColumns[col] = struct{}{}
			}

			if w.Cells[row][col] == CellTypeBoxLeft {
				nextAffectedColumns[col+1] = struct{}{}
				nextAffectedColumns[col] = struct{}{}
			}
		}

		for col := range nextAffectedColumns {
			affected = append(affected, Point{row, col})
		}

		affectedColumns = nextAffectedColumns

		row++
	}
}

func (w *WiderWarehouse) Left() bool {
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

func (w *WiderWarehouse) Right() bool {
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

func (w *WiderWarehouse) Print() {
	for _, row := range w.Cells {
		for _, cell := range row {
			if cell == CellTypeWall {
				fmt.Print("#")
			} else if cell == CellTypeEmpty {
				fmt.Print(".")
			} else if cell == CellTypeBoxLeft {
				fmt.Print("[")
			} else if cell == CellTypeBoxRight {
				fmt.Print("]")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", w.Width))
}

func (w *WiderWarehouse) GPS() int {
	var gps int

	for row, rowCells := range w.Cells {
		for col, cell := range rowCells {
			if cell == CellTypeBoxLeft {
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
