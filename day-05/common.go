package common

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type OrderRule struct {
	Before int
	After  int
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

func ParseInput(fileName string) ([]OrderRule, []Update, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var isOrderRule = true
	var orderRules []OrderRule
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
