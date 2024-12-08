package common

import (
	"bufio"
	"os"
)

type Point struct {
	Row int
	Col int
}

type Map struct {
	Width    int
	Height   int
	Antennas map[rune][]Point
}

func ParseInput(fileName string) (Map, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Map{}, err
	}
	defer file.Close()

	antennas := make(map[rune][]Point)
	var width int
	var height int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], Point{Row: i, Col: height})
			}
		}
		width = len(line)
		height++
	}

	if err := scanner.Err(); err != nil {
		return Map{}, err
	}

	return Map{
		Width:    width,
		Height:   height,
		Antennas: antennas,
	}, nil
}
