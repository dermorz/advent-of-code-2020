// Package day3 solves Toboggan Trajectory
package day3

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
)

const (
	tree = "#"
)

// PartOne of the day3 challenge
func PartOne(r io.Reader) (int, error) {
	input, err := scanMap(r)
	if err != nil {
		return 0, fmt.Errorf("scanning input: %w", err)
	}
	return treesOnSlope(input, 3, 1), nil
}

// PartTwo of the day3 challenge
func PartTwo(r io.Reader) (int, error) {
	input, err := scanMap(r)
	if err != nil {
		return 0, fmt.Errorf("scanning input: %w", err)
	}
	return treesOnSlope(input, 1, 1) *
		treesOnSlope(input, 3, 1) *
		treesOnSlope(input, 5, 1) *
		treesOnSlope(input, 7, 1) *
		treesOnSlope(input, 1, 2), nil
}

func treesOnSlope(input []string, right, down int) int {
	ring := initRing(len(input[0]))
	var trees int
	for i := 0; i < len(input); i = i + down {
		line := input[i]
		j := ring.Value.(int)
		fmt.Println(j)
		if string(line[j]) == tree {
			trees++
		}
		ring = ring.Move(right)
	}
	return trees
}

func initRing(n int) *ring.Ring {
	ring := ring.New(n)
	for i := 0; i < ring.Len(); i++ {
		ring.Value = i
		ring = ring.Next()
	}
	return ring
}

func scanMap(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
