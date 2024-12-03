package common

import (
	"fmt"
	"io"
	"os"
)

func ParseInput(fileName string) ([]int, []int, error) {
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
