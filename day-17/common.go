package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	RegisterA          int64
	RegisterB          int64
	RegisterC          int64
	InstructionPointer int64
	Instructions       []int64
}

func (c *Computer) getLiteralOperand(operand int64) int64 {
	return operand
}

func (c *Computer) getComboOperand(operand int64) int64 {
	if operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return c.RegisterA
	case 5:
		return c.RegisterB
	case 6:
		return c.RegisterC
	default:
		return 0
	}
}

func (c *Computer) Execute() []int64 {
	var results []int64
	length := int64(len(c.Instructions))
	for c.InstructionPointer < length {
		operator := c.Instructions[c.InstructionPointer]
		operand := c.Instructions[c.InstructionPointer+1]

		switch operator {
		case 0:
			numerator := c.RegisterA
			operand = c.getComboOperand(operand)
			for range operand {
				numerator /= 2
			}
			c.RegisterA = numerator
			c.InstructionPointer += 2
		case 1:
			lhs := c.RegisterB
			rhs := c.getLiteralOperand(operand)
			c.RegisterB = lhs ^ rhs
			c.InstructionPointer += 2
		case 2:
			operand = c.getComboOperand(operand)
			c.RegisterB = operand % 8
			c.InstructionPointer += 2
		case 3:
			a := c.RegisterA
			if a == 0 {
				c.InstructionPointer += 2
				continue
			}
			target := c.getLiteralOperand(operand)
			c.InstructionPointer = target
		case 4:
			lhs := c.RegisterB
			rhs := c.RegisterC
			c.RegisterB = lhs ^ rhs
			c.InstructionPointer += 2
		case 5:
			operand = c.getComboOperand(operand)
			result := operand % 8
			results = append(results, result)
			c.InstructionPointer += 2
		case 6:
			numerator := c.RegisterA
			operand = c.getComboOperand(operand)
			for range operand {
				numerator /= 2
			}
			c.RegisterB = numerator
			c.InstructionPointer += 2
		case 7:
			numerator := c.RegisterA
			operand = c.getComboOperand(operand)
			for range operand {
				numerator /= 2
			}
			c.RegisterC = numerator
			c.InstructionPointer += 2
		}
	}

	return results
}

func ParseInput(fileName string) (Computer, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Computer{}, err
	}
	defer file.Close()

	var computer Computer
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A: ") {
			fmt.Sscanf(line, "Register A: %d", &computer.RegisterA)
		}
		if strings.HasPrefix(line, "Register B: ") {
			fmt.Sscanf(line, "Register B: %d", &computer.RegisterB)
		}
		if strings.HasPrefix(line, "Register C: ") {
			fmt.Sscanf(line, "Register C: %d", &computer.RegisterC)
		}
		if strings.HasPrefix(line, "Program: ") {
			parts := strings.Split(strings.TrimPrefix(line, "Program: "), ",")
			for _, part := range parts {
				instruction, _ := strconv.Atoi(part)
				computer.Instructions = append(computer.Instructions, int64(instruction))
			}
		}
	}
	computer.InstructionPointer = 0

	return computer, nil
}
