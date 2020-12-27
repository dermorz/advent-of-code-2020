package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dermorz/advent-of-code-2020/pkg/day1"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	day := flag.Int("day", 1, "Day of of December (1-24)")
	part := flag.Int("part", 1, "Part 1 or 2")

	flag.Parse()

	if *day < 1 || *day > 24 {
		return errors.New("day must be between 1 and 24")
	}
	if *part < 1 || *part > 2 {
		return errors.New("part must be 1 or 2")
	}

	var f func(io.Reader) (int, error)
	var result int
	switch *day {
	case 1:
		switch *part {
		case 1:
			f = day1.PartOne
		default:
			fmt.Printf("Day %d Part %d not implemented yet.", *day, *part)
			os.Exit(1)
		}
	default:
		fmt.Printf("Day %d not implemented yet.", *day)
		os.Exit(1)
	}

	input, err := os.Open(fmt.Sprintf("./input/%02d.txt", *day))
	if err != nil {
		return err
	}
	defer input.Close()

	result, err = f(input)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
