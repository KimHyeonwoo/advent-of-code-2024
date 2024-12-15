package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	PositionRow int
	PositionCol int
	VelocityRow int
	VelocityCol int
}

func (r *Robot) Move(width, height int) {
	r.PositionRow = (r.PositionRow + r.VelocityRow + height) % height
	r.PositionCol = (r.PositionCol + r.VelocityCol + width) % width
}

func (r *Robot) MoveWidthTimes(width, height int) {
	r.PositionRow = (r.PositionRow + r.VelocityRow*width + height*width) % height
}

func ParseInput(fileName string) ([]Robot, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var robots []Robot
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimPrefix(line, "p=")
		line = strings.ReplaceAll(line, " v=", ",")

		parts := strings.Split(line, ",")
		positionRow, _ := strconv.Atoi(parts[1])
		positionCol, _ := strconv.Atoi(parts[0])
		velocityRow, _ := strconv.Atoi(parts[3])
		velocityCol, _ := strconv.Atoi(parts[2])

		robot := Robot{
			PositionRow: positionRow,
			PositionCol: positionCol,
			VelocityRow: velocityRow,
			VelocityCol: velocityCol,
		}

		robots = append(robots, robot)
	}

	return robots, nil
}
