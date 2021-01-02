package day3_test

import (
	"io"
	"strings"
	"testing"

	"github.com/dermorz/advent-of-code-2020/pkg/day3"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		r         io.Reader
		expected  int
		expectErr bool
	}{
		"example": {
			strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`),
			7,
			false,
		},
		"empty input": {
			strings.NewReader(``),
			0,
			false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := day3.PartOne(tt.r)
			if tt.expectErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]struct {
		r         io.Reader
		expected  int
		expectErr bool
	}{
		"example": {
			strings.NewReader(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`),
			336,
			false,
		},
		"empty input": {
			strings.NewReader(``),
			0,
			false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := day3.PartTwo(tt.r)
			if tt.expectErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
