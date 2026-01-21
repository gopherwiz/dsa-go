package main

import (
	"fmt"
	"math"
)

// safeGet using math constants
func safeGet(A []int, index int) int {
	if index >= 0 && index < len(A) {
		return A[index]
	}
	// Use MaxInt for "Inifinity" or MinInt for "-Infinity"
	return math.MinInt // forces a direction to find peak
}

func findPeak(A []int) int {
	n := len(A)
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2

		curr := A[mid]
		// Using safeGet to handle boundaries automatically
		prev := safeGet(A, mid-1)
		next := safeGet(A, mid+1)

		if curr > prev && curr > next {
			return mid
		}

		if next > curr {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	A := []int{1, 2, 3, 1}
	fmt.Println("Peak index:", findPeak(A))
}
