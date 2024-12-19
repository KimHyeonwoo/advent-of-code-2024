package common

import (
	"bufio"
	"os"
	"strings"
)

type Towel string

func (t Towel) Construct(candidates []Towel) int {
	if len(t) == 0 {
		return 0
	}

	dp := make([]int, len(t)+1)
	dp[0] = 1

	for i := 0; i <= len(t); i++ {
		for _, candidate := range candidates {
			if i+len(candidate) <= len(t) && t[i:i+len(candidate)] == candidate {
				dp[i+len(candidate)] += dp[i]
			}
		}
	}

	return dp[len(t)]
}

func ParseInput(filename string) ([]Towel, []Towel, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	candidates := make([]Towel, 0)
	targets := make([]Towel, 0)

	candidateMode := true

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			candidateMode = false
			continue
		}

		if candidateMode {
			splits := strings.Split(line, ",")
			for _, split := range splits {
				candidates = append(candidates, Towel(strings.TrimSpace(split)))
			}
		} else {
			targets = append(targets, Towel(strings.TrimSpace(line)))
		}
	}

	return candidates, targets, nil
}
