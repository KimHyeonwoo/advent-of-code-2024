package common

import (
	"bufio"
	"os"
)

type Garden struct {
	Width     int
	Height    int
	ValueRune [][]rune
	ValueInt  [][]int
}

type Point struct {
	Row int
	Col int
}

func (g Garden) ParseValue() {
	visited := make([][]bool, g.Height)
	for i := 0; i < g.Height; i++ {
		visited[i] = make([]bool, g.Width)
	}

	value := 1
	queue := make([]Point, 0)
	dir := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for row := 0; row < g.Height; row++ {
		for col := 0; col < g.Width; col++ {
			if visited[row][col] {
				continue
			}

			visited[row][col] = true
			g.ValueInt[row+1][col+1] = value
			queue = append(queue, Point{Row: row, Col: col})

			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]

				for _, d := range dir {
					nr := p.Row + d.Row
					nc := p.Col + d.Col

					if nr < 0 || nr >= g.Height || nc < 0 || nc >= g.Width {
						continue
					}

					if visited[nr][nc] {
						continue
					}

					if g.ValueRune[nr][nc] != g.ValueRune[p.Row][p.Col] {
						continue
					}

					visited[nr][nc] = true
					g.ValueInt[nr+1][nc+1] = value
					queue = append(queue, Point{Row: nr, Col: nc})
				}
			}

			value++
		}
	}
}

func ParseInput(fileName string) (Garden, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Garden{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineValues := make([]rune, 0)
		for _, r := range line {
			lineValues = append(lineValues, r)
		}
		values = append(values, lineValues)
	}

	valueInt := make([][]int, len(values)+2)
	for i := 0; i < len(values)+2; i++ {
		valueInt[i] = make([]int, len(values[0])+2)
	}

	return Garden{
		Width:     len(values[0]),
		Height:    len(values),
		ValueRune: values,
		ValueInt:  valueInt,
	}, nil
}
