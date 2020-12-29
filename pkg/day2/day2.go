// Package day2 solves Password Philosophy
package day2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Password struct {
	value  string
	letter string
	low    int
	high   int
}

// PartOne of the day2 challenge
func PartOne(r io.Reader) (int, error) {
	var valid int

	passwords, err := scanPasswords(r)
	if err != nil {
		return 0, fmt.Errorf("scanning ints from input: %w", err)
	}

	for _, pw := range passwords {
		c := strings.Count(pw.value, pw.letter)
		if c >= pw.low && c <= pw.high {
			valid++
		}
	}

	return valid, nil
}

// PartTwo of the day2 challenge
func PartTwo(r io.Reader) (int, error) {
	var valid int

	passwords, err := scanPasswords(r)
	if err != nil {
		return 0, fmt.Errorf("scanning ints from input: %w", err)
	}

	for _, pw := range passwords {
		if (string(pw.value[pw.low-1]) == pw.letter) != (string(pw.value[pw.high-1]) == pw.letter) {
			valid++
		}
	}

	return valid, nil
}

func scanPasswords(r io.Reader) ([]Password, error) {
	var passords []Password
	scanner := bufio.NewScanner(r)

	var re = regexp.MustCompile(`(?m)(\d+)-(\d+) (\w): (\w+)`)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if matches == nil {
			return nil, errors.New("data pattern unexpected")
		}
		// strconv.Atoi() cannot result in errors here
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		pw := Password{
			value:  matches[4],
			letter: matches[3],
			low:    min,
			high:   max,
		}
		passords = append(passords, pw)
	}
	return passords, nil
}
