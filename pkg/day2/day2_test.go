package day2_test

import (
	"io"
	"strings"
	"testing"

	"github.com/dermorz/advent-of-code-2020/pkg/day2"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		r         io.Reader
		expected  int
		expectErr bool
	}{
		"example": {
			strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`),
			2,
			false,
		},
		"empty input": {
			strings.NewReader(``),
			0,
			false,
		},
		"pattern not matched": {
			strings.NewReader(`foo`),
			0,
			true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := day2.PartOne(tt.r)
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
			strings.NewReader(`1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`),
			1,
			false,
		},
		"empty input": {
			strings.NewReader(``),
			0,
			false,
		},
		"pattern not matched": {
			strings.NewReader(`foo`),
			0,
			true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := day2.PartTwo(tt.r)
			if tt.expectErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
