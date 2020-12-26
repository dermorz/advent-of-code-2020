// Package day1 solves Report Repair
package day1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// PartOne of the day1 challenge
func PartOne(r io.Reader) (int, error) {
	numbers, err := scanInts(r)
	if err != nil {
		return 0, fmt.Errorf("scanning ints from input: %w", err)
	}
	for i := range numbers {
		for j := i; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i] * numbers[j], nil
			}
		}
	}

	return 0, nil
}

func scanInts(r io.Reader) ([]int, error) {
	var numbers []int
	scanner := bufio.NewScanner(r)

	for {
		scanner.Scan()
		number := scanner.Text()
		if len(number) == 0 {
			break
		}
		n, err := strconv.Atoi(number)
		if err != nil {
			return nil, fmt.Errorf("converting %v to int: %w", number, err)
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}
