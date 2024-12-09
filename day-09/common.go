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

func (d *Disk) CompactV2() {
	fileIdx := d.Blocks[len(d.Blocks)-1]

	for i := fileIdx; i >= 0; i-- {
		var fileStartIdx, fileEndIdx int

		for j := 0; j < len(d.Blocks); j++ {
			if d.Blocks[j] == i {
				fileStartIdx = j
				break
			}
		}

		for j := len(d.Blocks) - 1; j >= 0; j-- {
			if d.Blocks[j] == i {
				fileEndIdx = j
				break
			}
		}

		fileLength := fileEndIdx - fileStartIdx + 1

		// Find consecutive empty blocks with length of fileLength
		var emptyCount int
		for j := 0; j < fileStartIdx; j++ {
			if d.Blocks[j] == -1 {
				emptyCount++
			} else {
				emptyCount = 0
			}

			if emptyCount == fileLength {
				for k := j; k > j-fileLength; k-- {
					d.Blocks[k] = i
				}
				for k := fileStartIdx; k <= fileEndIdx; k++ {
					d.Blocks[k] = -1
				}
				break
			}
		}
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
