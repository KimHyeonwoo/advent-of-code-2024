package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Stone int

func (s Stone) Process() []Stone {
	if s == 0 {
		return []Stone{1}
	}

	stoneStr := strconv.Itoa(int(s))

	if len(stoneStr)%2 == 0 {
		leftStoneStr := stoneStr[:len(stoneStr)/2]
		rightStoneStr := stoneStr[len(stoneStr)/2:]

		leftStone, _ := strconv.Atoi(leftStoneStr)
		rightStone, _ := strconv.Atoi(rightStoneStr)

		return []Stone{Stone(leftStone), Stone(rightStone)}
	}

	return []Stone{s * 2024}
}

func ParseInput(fileName string) ([]Stone, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var stones []Stone

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // Reads the entire line

		elements := strings.Split(strings.TrimSpace(line), " ")

		for _, element := range elements {
			stone, convertErr := strconv.Atoi(element)
			if convertErr != nil {
				return nil, convertErr
			}
			stones = append(stones, Stone(stone))
		}
	}

	return stones, nil
}
