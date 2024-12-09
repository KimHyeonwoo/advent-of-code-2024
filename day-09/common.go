package common

import (
	"bufio"
	"os"
	"strconv"
)

type Disk struct {
	Blocks []int
}

func (d *Disk) Compact() {
	leftIndex := 0
	rightIndex := len(d.Blocks) - 1

	for leftIndex < rightIndex {
		if d.Blocks[leftIndex] != -1 {
			leftIndex++
			continue
		}

		if d.Blocks[rightIndex] == -1 {
			rightIndex--
			continue
		}

		d.Blocks[leftIndex], d.Blocks[rightIndex] = d.Blocks[rightIndex], d.Blocks[leftIndex]
		leftIndex++
		rightIndex--
	}
}

func (d *Disk) ComputeChecksum() int {
	var checksum int

	for idx, block := range d.Blocks {
		if block == -1 {
			continue
		}
		checksum += idx * block
	}

	return checksum
}

func ParseInput(fileName string) (Disk, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Disk{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var blocks []int
		var isFile = true
		var fileIdx = 0

		for _, char := range line {
			length, err := strconv.Atoi(string(char))
			if err != nil {
				return Disk{}, err
			}
			if isFile {
				for i := 0; i < length; i++ {
					blocks = append(blocks, fileIdx)
				}
				isFile = false
				fileIdx++
			} else {
				for i := 0; i < length; i++ {
					blocks = append(blocks, -1)
				}
				isFile = true
			}
		}

		return Disk{blocks}, nil
	}

	return Disk{}, nil
}
