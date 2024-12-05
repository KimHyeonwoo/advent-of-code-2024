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
			continue
		}

		relatedRules := orderRules.GetRulesFor(update)
		topology, topologyErr := relatedRules.GetTopology()
		if topologyErr != nil {
			panic(topologyErr)
		}
		answer += topology[len(topology)/2]
	}

	fmt.Println(answer) // 4944
}
