package main

import (
	"fmt"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-17"
)

func backtrack(currentAnswer int64, length int, computer common.Computer) {
	if length > len(computer.Instructions) {
		fmt.Println(currentAnswer)
		return
	}

	for mod := int64(0); mod < 8; mod++ {
		newComputer := common.Computer{
			RegisterA:          currentAnswer*8 + mod,
			RegisterB:          0,
			RegisterC:          0,
			InstructionPointer: 0,
			Instructions:       computer.Instructions,
		}
		results := newComputer.Execute()

		mismatch := false

		for i := range results {
			if results[i] != computer.Instructions[i+len(computer.Instructions)-length] {
				mismatch = true
				break
			}
		}

		if mismatch {
			continue
		}

		backtrack(currentAnswer*8+mod, length+1, computer)
	}
}

func main() {
	computer, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	backtrack(0, 1, computer) // 107413700225434
}
