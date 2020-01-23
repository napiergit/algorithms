package lazy_landscaper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"solutions/lazy_landscaper"
)

// https://cs.stackexchange.com/questions/116854/minimum-number-of-tree-cuts-so-that-each-pair-of-trees-alternates-between-strict
// https://www.geeksforgeeks.org/convert-to-strictly-increasing-integer-array-with-minimum-changes/

func TestSolution(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
		err      error
	}{
		{[]int{1, 2, 3, 4, 5}, 2, nil}, // Odd peak
		{[]int{1, 1}, 0, lazy_landscaper.ErrStupidInput},
		{[]int{1, 2, 1}, 0, nil},                   // Odd peak
		{[]int{3, 1, 3, 5, 6, 7, 8, 9}, 3, nil},    // Even peak
		{[]int{3, 1, 3, 5, 6, 2, 8, 9}, 2, nil},    // Even peak
		{[]int{8, 8, 8, 8, 8, 8, 8, 8}, 4, nil},    // Either
		{[]int{8, 8, 8, 8, 8, 8, 8, 8, 8}, 4, nil}, // Even peak
	}

	for k, tc := range testCases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			actual, err := lazy_landscaper.Solution(tc.A)
			assert.Equal(t, tc.err, err)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
