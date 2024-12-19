package common

import (
	"bufio"
	"os"
	"strings"
)

type Towel string

var trueCache = make(map[Towel]struct{})
var falseCache = make(map[Towel]struct{})

func (t Towel) CanConstruct(candidates []Towel) bool {
	if len(t) == 0 {
		return true
	}

	if _, ok := trueCache[t]; ok {
		return true
	}

	if _, ok := falseCache[t]; ok {
		return false
	}

	for _, candidate := range candidates {
		if strings.HasPrefix(string(t), string(candidate)) {
			if Towel(string(t)[len(candidate):]).CanConstruct(candidates) {
				trueCache[t] = struct{}{}
				return true
			}
		}
	}

	falseCache[t] = struct{}{}
	return false
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
