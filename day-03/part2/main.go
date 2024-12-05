package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type ActionType int

const (
	ActionDo ActionType = iota
	ActionDont
	ActionMultiply
)

type Action struct {
	Type  ActionType
	Index int
	Left  int
	Right int
}

type Actions []Action

func (a Actions) Len() int {
	return len(a)
}

func (a Actions) Less(i, j int) bool {
	return a[i].Index < a[j].Index
}

func (a Actions) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

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

	actions := make(Actions, 0)

	re := regexp.MustCompile("mul\\((?P<left>\\d{1,3}),(?P<right>\\d{1,3})\\)")
	matches := re.FindAllStringSubmatch(memory, -1)
	matchIndexes := re.FindAllStringIndex(memory, -1)

	for matchIdx, match := range matches {
		leftStr := match[1]
		rightStr := match[2]

		left, convertErr := strconv.Atoi(leftStr)
		if convertErr != nil {
			panic(convertErr)
		}
		right, convertErr := strconv.Atoi(rightStr)
		if convertErr != nil {
			panic(convertErr)
		}

		idx := matchIndexes[matchIdx][0]
		actions = append(actions, Action{
			Type:  ActionMultiply,
			Index: idx,
			Left:  left,
			Right: right,
		})
	}

	doRe := regexp.MustCompile("do\\(\\)")
	doIndexes := doRe.FindAllStringIndex(memory, -1)
	for _, idx := range doIndexes {
		actions = append(actions, Action{
			Type:  ActionDo,
			Index: idx[0],
		})
	}

	dontRe := regexp.MustCompile("don't\\(\\)")
	dontIndexes := dontRe.FindAllStringIndex(memory, -1)
	for _, idx := range dontIndexes {
		actions = append(actions, Action{
			Type:  ActionDont,
			Index: idx[0],
		})
	}

	sort.Sort(actions)

	sum := 0
	canMultiply := true
	for _, action := range actions {
		switch action.Type {
		case ActionDo:
			canMultiply = true
		case ActionDont:
			canMultiply = false
		case ActionMultiply:
			if canMultiply {
				sum += action.Left * action.Right
			}
		}
	}

	fmt.Println(sum) // 63013756
}
