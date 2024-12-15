package common

import (
	"bufio"
	"fmt"
	"os"
)

type Machine struct {
	AX     int64
	AY     int64
	BX     int64
	BY     int64
	PrizeX int64
	PrizeY int64
}

func (m Machine) GetPrize() (int64, bool) {
	var cost int64 = 500

	var x int64
	var y int64

	for x = 0; x <= 100; x++ {
		for y = 0; y <= 100; y++ {
			if m.AX*x+m.BX*y == m.PrizeX && m.AY*x+m.BY*y == m.PrizeY {
				if 3*x+y < cost {
					cost = 3*x + y
				}
			}
		}
	}

	if cost == 500 {
		return 0, false
	}

	return cost, true
}

func (m Machine) GetActualPrize() (int64, bool) {
	PrizeX := m.PrizeX + 10000000000000
	PrizeY := m.PrizeY + 10000000000000

	numerator := PrizeX*m.BY - PrizeY*m.BX
	denominator := m.AX*m.BY - m.AY*m.BX

	if denominator == 0 {
		return 0, false
	}

	if numerator%denominator != 0 {
		return 0, false
	}

	x := numerator / denominator
	if x < 0 {
		return 0, false
	}

	numerator = PrizeX - m.AX*x
	denominator = m.BX

	if denominator == 0 {
		return 0, false
	}

	if numerator%denominator != 0 {
		return 0, false
	}

	y := numerator / denominator
	if y < 0 {
		return 0, false
	}

	return 3*x + y, true
}

func ParseInput(fileName string) ([]Machine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var machine Machine

		fmt.Sscanf(scanner.Text(), "Button A: X+%d, Y+%d", &machine.AX, &machine.AY)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d", &machine.BX, &machine.BY)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d", &machine.PrizeX, &machine.PrizeY)
		scanner.Scan()

		machines = append(machines, machine)
	}

	return machines, nil
}
