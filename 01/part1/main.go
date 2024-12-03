package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	left, right, err := parseInput("input")
	if err != nil {
		panic(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	var answer int

	for i := 0; i < len(left); i++ {
		answer += distance(left[i], right[i])
	}

	fmt.Println(answer) // 1579939
}
