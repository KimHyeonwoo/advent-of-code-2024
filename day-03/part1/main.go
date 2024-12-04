package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	memory := string(bytes)

	re := regexp.MustCompile("mul\\((?P<left>\\d{1,3}),(?P<right>\\d{1,3})\\)")
	matches := re.FindAllString(memory, -1)

	sum := 0

	for _, match := range matches {
		leftStr := re.FindStringSubmatch(match)[1]
		rightStr := re.FindStringSubmatch(match)[2]
		left, convertErr := strconv.Atoi(leftStr)
		if convertErr != nil {
			panic(convertErr)
		}
		right, convertErr := strconv.Atoi(rightStr)
		if convertErr != nil {
			panic(convertErr)
		}

		sum += left * right
	}

	fmt.Println(sum) // 189527826
}
