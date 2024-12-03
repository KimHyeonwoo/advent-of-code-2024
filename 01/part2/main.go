package main

import (
	"fmt"
	"io"
	"os"
)

func parseInput(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left []int
	var right []int

	for {
		var l, r int
		_, err := fmt.Fscanln(file, &l, &r)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}
		left = append(left, l)
		right = append(right, r)
	}

	return left, right, nil
}

func main() {
	left, right, err := parseInput("input")
	if err != nil {
		panic(err)
	}

	leftOccurrenceCount := make(map[int]int)
	rightOccurrenceCount := make(map[int]int)

	for _, l := range left {
		leftOccurrenceCount[l]++
	}

	for _, r := range right {
		rightOccurrenceCount[r]++
	}

	var answer int

	for l, count := range leftOccurrenceCount {
		answer += l * count * rightOccurrenceCount[l]
	}

	fmt.Println(answer) // 20351745
}
