package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func (r Report) IsSafe() bool {
	if len(r.levels) < 2 {
		return true
	}

	firstTwoDifference := r.levels[1] - r.levels[0]

	for i := 1; i < len(r.levels); i++ {
		difference := r.levels[i] - r.levels[i-1]
		if difference > 3 || difference < -3 {
			return false
		}

		if firstTwoDifference*difference <= 0 {
			return false
		}
	}

	return true
}

func ParseInput(fileName string) ([]Report, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports []Report
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // Reads the entire line

		elements := strings.Split(line, " ")
		levels := make([]int, 0, len(elements))

		for _, element := range elements {
			level, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			levels = append(levels, level)
		}

		reports = append(reports, Report{levels: levels})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
