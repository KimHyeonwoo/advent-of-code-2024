package common

import (
	"fmt"
	"io"
	"os"
)

type Board struct {
	Board  [][]rune
	Width  int
	Height int
}

func (b Board) HasXMAS(x, y, dx, dy int) bool {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return false
	}

	if b.Board[y][x] != 'X' {
		return false
	}

	if x+dx < 0 || x+dx >= b.Width || y+dy < 0 || y+dy >= b.Height {
		return false
	}

	if b.Board[y+dy][x+dx] != 'M' {
		return false
	}

	if x+2*dx < 0 || x+2*dx >= b.Width || y+2*dy < 0 || y+2*dy >= b.Height {
		return false
	}

	if b.Board[y+2*dy][x+2*dx] != 'A' {
		return false
	}

	if x+3*dx < 0 || x+3*dx >= b.Width || y+3*dy < 0 || y+3*dy >= b.Height {
		return false
	}

	if b.Board[y+3*dy][x+3*dx] != 'S' {
		return false
	}

	return true
}

func ParseInput(fileName string) (Board, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Board{}, err
	}
	defer file.Close()

	var board [][]rune
	for {
		var row string
		_, err := fmt.Fscanln(file, &row)
		if err != nil {
			if err == io.EOF {
				break
			}
			return Board{}, err
		}
		board = append(board, []rune(row))
	}

	width := len(board[0])
	height := len(board)

	return Board{Board: board, Width: width, Height: height}, nil
}
