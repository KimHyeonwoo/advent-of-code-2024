package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type OrderRule struct {
	Before int
	After  int
}

type OrderRules []OrderRule

func (o OrderRules) GetRulesFor(update Update) OrderRules {
	var rules OrderRules
	for _, rule := range o {
		if slices.Contains(update.Pages, rule.Before) && slices.Contains(update.Pages, rule.After) {
			rules = append(rules, rule)
		}
	}

	return rules
}

func (o OrderRules) GetTopology() ([]int, error) {
	inDegrees := make(map[int]int)
	nodes := make(map[int]struct{})

	for _, rule := range o {
		inDegrees[rule.After]++
		nodes[rule.Before] = struct{}{}
		nodes[rule.After] = struct{}{}
	}

	var queue []int
	for node := range nodes {
		if _, ok := inDegrees[node]; !ok {
			inDegrees[node] = 0
		}

		if inDegrees[node] == 0 {
			queue = append(queue, node)
		}
	}

	var sorted []int

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, rule := range o {
			if rule.Before == node {
				inDegrees[rule.After]--
				if inDegrees[rule.After] == 0 {
					queue = append(queue, rule.After)
				}
			}
		}
	}

	if len(sorted) != len(inDegrees) {
		fmt.Println(sorted)
		return nil, errors.New("cycle detected")
	}

	return sorted, nil
}

type Update struct {
	Pages []int
}

func (u Update) HasRuleViolation(rule OrderRule) bool {
	beforeIndex := -1
	afterIndex := -1
	for i, page := range u.Pages {
		if page == rule.Before {
			beforeIndex = i
		}
		if page == rule.After {
			afterIndex = i
		}
	}

	return beforeIndex != -1 && afterIndex != -1 && beforeIndex > afterIndex
}

func (u Update) GetMiddleElement() int {
	return u.Pages[len(u.Pages)/2]
}

func ParseInput(fileName string) (OrderRules, []Update, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var isOrderRule = true
	var orderRules OrderRules
	var updates []Update

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			isOrderRule = false
			continue
		}

		if isOrderRule {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				return nil, nil, errors.New("invalid input")
			}
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			orderRules = append(orderRules, OrderRule{Before: before, After: after})
		} else {
			parts := strings.Split(line, ",")
			var pages []int
			for _, part := range parts {
				page, _ := strconv.Atoi(part)
				pages = append(pages, page)
			}
			updates = append(updates, Update{Pages: pages})
		}
	}

	return orderRules, updates, nil
}
