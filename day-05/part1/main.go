package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-05"
)

func main() {
	orderRules, updates, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	var answer int

	for _, update := range updates {
		hasViolation := false

		for _, rule := range orderRules {
			if update.HasRuleViolation(rule) {
				hasViolation = true
				break
			}
		}

		if !hasViolation {
			answer += update.GetMiddleElement()
		}
	}

	fmt.Println(answer) // 6612
}
