package lazy_landscaper

import (
	"errors"
)

var ErrStupidInput = errors.New("bleh")

func Solution(A []int) (int, error) {
	evenPeak := false
	oddPeak := false
	for k, v := range A {
		if v == 1 {
			evenPeak = k%2 != 0 || evenPeak
			oddPeak = k%2 == 0 || oddPeak
		}
	}

	if evenPeak && oddPeak {
		return 0, ErrStupidInput
	}

	if evenPeak {
		return getEvenPeakTotal(A), nil
	}

	if oddPeak {
		return getOddPeakTotal(A), nil
	}

	return max(getOddPeakTotal(A), getEvenPeakTotal(A)), nil
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func getOddPeakTotal(A []int) int {
	cuts := 0
	for i := 0; i < len(A)-1; i++ {
		if i%2 == 0 {
			if A[i] >= A[i+1] {
				A[i] = 1
				cuts++
			}
		}

		if i%2 != 0 {
			if A[i] < A[i+1] {
				A[i+1] = 1
				cuts++
			}
		}
	}
	return cuts
}

func getEvenPeakTotal(A []int) int {
	cuts := 0
	for i := 0; i < len(A)-1; i++ {
		if i%2 != 0 {
			if A[i] > A[i+1] {
				A[i] = 1
				cuts++
			}
		}

		if i%2 == 0 {
			if A[i] <= A[i+1] {
				A[i+1] = 1
				cuts++
			}
		}
	}
	return cuts
}
