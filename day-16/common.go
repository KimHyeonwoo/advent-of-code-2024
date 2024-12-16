package common

import (
	"bufio"
	"math"
	"os"
	"sort"
)

type Point struct {
	Row int
	Col int
}

type PointWithDir struct {
	Point
	Dir int
}

type Data struct {
	Point
	Dir  int
	Cost int
}

type DataByCost []Data

func (d DataByCost) Len() int {
	return len(d)
}

func (d DataByCost) Less(i, j int) bool {
	return d[i].Cost < d[j].Cost
}

func (d DataByCost) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type CellType int

const (
	CellTypeWall CellType = iota
	CellTypeEmpty
	CellTypeStart
	CellTypeEnd
)

type Maze struct {
	Width  int
	Height int
	Cells  [][]CellType
	Start  Point
	End    Point
}

var dirs = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func (m Maze) Solve() (int, int) {
	dp := make([][][]int, m.Height)
	for i := range dp {
		dp[i] = make([][]int, m.Width)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}

	dp[m.Start.Row][m.Start.Col][0] = 0
	for i := 1; i < 4; i++ {
		dp[m.Start.Row][m.Start.Col][i] = 1000
	}

	candidates := make([]Data, 0)
	candidates = append(candidates, Data{m.Start, 0, 0})
	candidates = append(candidates, Data{m.Start, 1, 1000})
	candidates = append(candidates, Data{m.Start, 2, 1000})
	candidates = append(candidates, Data{m.Start, 3, 1000})

	visited := make(map[PointWithDir]struct{})
	visited[PointWithDir{m.Start, 0}] = struct{}{}
	visited[PointWithDir{m.Start, 1}] = struct{}{}
	visited[PointWithDir{m.Start, 2}] = struct{}{}
	visited[PointWithDir{m.Start, 3}] = struct{}{}

	for {
		if len(candidates) == 0 {
			break
		}

		sort.Sort(DataByCost(candidates))

		data := candidates[0]
		candidates = candidates[1:]

		row := data.Row
		col := data.Col
		cost := data.Cost
		dir := data.Dir

		for i := range dirs {
			if i == dir {
				continue
			}

			newPointWithDir := PointWithDir{Point{row, col}, i}

			if _, ok := visited[newPointWithDir]; ok {
				continue
			}

			newCost := cost + 1000

			candidates = append(candidates, Data{Point{row, col}, i, newCost})
			dp[row][col][i] = newCost
			visited[newPointWithDir] = struct{}{}
		}

		newRow := row + dirs[dir].Row
		newCol := col + dirs[dir].Col

		newPointWithDir := PointWithDir{Point{newRow, newCol}, dir}

		if _, ok := visited[newPointWithDir]; ok {
			continue
		}

		if m.Cells[newRow][newCol] == CellTypeWall {
			continue
		}

		newCost := cost + 1

		candidates = append(candidates, Data{Point{newRow, newCol}, dir, newCost})
		dp[newRow][newCol][dir] = newCost
		visited[newPointWithDir] = struct{}{}
	}

	minCost := math.MaxInt
	minCostDir := -1
	for i := 0; i < 4; i++ {
		currCost := dp[m.End.Row][m.End.Col][i]
		if currCost != 0 && currCost < minCost {
			minCost = dp[m.End.Row][m.End.Col][i]
			minCostDir = i
		}
	}

	pointVisited := make(map[Point]struct{})
	pointVisited[m.End] = struct{}{}
	queue := make([]PointWithDir, 0)
	queue = append(queue, PointWithDir{m.End, minCostDir})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		row := p.Row
		col := p.Col
		dir := p.Dir
		cost := dp[row][col][dir]

		for i := range dirs {
			if i == dir {
				continue
			}

			if dp[row][col][i] == cost-1000 {
				queue = append(queue, PointWithDir{Point{row, col}, i})
				pointVisited[Point{row, col}] = struct{}{}
			}
		}

		newRow := row - dirs[dir].Row
		newCol := col - dirs[dir].Col

		if dp[newRow][newCol][dir] == cost-1 {
			queue = append(queue, PointWithDir{Point{newRow, newCol}, dir})
			pointVisited[Point{newRow, newCol}] = struct{}{}
		}
	}

	count := 0
	for range pointVisited {
		count++
	}

	return minCost, count
}

func ParseInput(fileName string) (Maze, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Maze{}, err
	}
	defer file.Close()

	var cells [][]CellType
	var start, end Point

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var row []CellType
		for _, c := range line {
			switch c {
			case '#':
				row = append(row, CellTypeWall)
			case '.':
				row = append(row, CellTypeEmpty)
			case 'S':
				row = append(row, CellTypeStart)
				start = Point{len(cells), len(row) - 1}
			case 'E':
				row = append(row, CellTypeEnd)
				end = Point{len(cells), len(row) - 1}
			}
		}
		cells = append(cells, row)
	}

	if err := scanner.Err(); err != nil {
		return Maze{}, err
	}

	return Maze{
		Width:  len(cells[0]),
		Height: len(cells),
		Cells:  cells,
		Start:  start,
		End:    end,
	}, nil
}
