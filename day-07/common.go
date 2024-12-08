package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Target  int
	Numbers []int
}

func (e Equation) IsValid() bool {
	if len(e.Numbers) == 0 {
		return false
	}

	if len(e.Numbers) == 1 {
		return e.Numbers[0] == e.Target
	}

	// Plus case
	plusEquation := Equation{
		Target:  e.Target - e.Numbers[len(e.Numbers)-1],
		Numbers: e.Numbers[:len(e.Numbers)-1],
	}

	if plusEquation.IsValid() {
		return true
	}

	// Multiply case
	if e.Target%e.Numbers[len(e.Numbers)-1] != 0 {
		return false
	}

	multiplyEquation := Equation{
		Target:  e.Target / e.Numbers[len(e.Numbers)-1],
		Numbers: e.Numbers[:len(e.Numbers)-1],
	}

	return multiplyEquation.IsValid()
}

func ParseInput(fileName string) ([]Equation, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var equations []Equation

	for scanner.Scan() {
		// 45370174: 926 900 3 314 79
		// 1361257567489: 4 2 6 7 52 959 8 351 2
		// parse the line using colon and spaces

		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		targetStr := line[:colonIndex]
		numbersStr := strings.Split(strings.TrimSpace(line[colonIndex+2:]), " ")

		target, err := strconv.Atoi(targetStr)
		if err != nil {
			return nil, err
		}

		var numbers []int
		for _, numStr := range numbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}

		equations = append(equations, Equation{Target: target, Numbers: numbers})
	}

	return equations, nil
}
