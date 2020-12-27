package day1_test

import (
	"io"
	"strings"
	"testing"

	"github.com/dermorz/advent-of-code-2020/pkg/day1"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := map[string]struct {
		r         io.Reader
		expected  int
		expectErr bool
	}{
		"example": {
			strings.NewReader(`1721
979
366
299
675
1456`),
			514579,
			false,
		},
		"empty input": {
			strings.NewReader(``),
			0,
			true,
		},
		"not a number": {
			strings.NewReader(`foo`),
			0,
			true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := day1.PartOne(tt.r)
			if tt.expectErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
