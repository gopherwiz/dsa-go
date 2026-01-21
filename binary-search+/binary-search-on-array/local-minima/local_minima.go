package main

import (
	"fmt"
	"math"
)

// safeGet returns math.MaxInt for out-of-bounds to treat boundaries as +infinity
func safeGet(A []int, index int) int {
	if index >= 0 && index < len(A) {
		return A[index]
	}

	return math.MaxInt // forces a direction to find valley
}

func findLocalMinimum(A []int) int {
	n := len(A)
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2

		curr := A[mid]
		prev := safeGet(A, mid-1)
		next := safeGet(A, mid+1)

		// Check if mid is a local minimum
		if curr < prev && curr < next {
			return mid
		}

		// If the element to the right is smaller, the slope goes down to the right
		if next < curr {
			l = mid + 1
		} else {
			// Otherwise, the slope must go down to the left
			r = mid - 1
		}
	}
	return -1
}

func main() {
	// Example 1: Minimum in the middle
	A1 := []int{9, 6, 3, 4, 5, 7, 8}
	fmt.Println("Local Minimum Index:", findLocalMinimum(A1)) // Output: 2 (value 3)

	// Example 2: Minimum at the start
	A2 := []int{1, 5, 10, 20}
	fmt.Println("Local Minimum Index:", findLocalMinimum(A2)) // Output: 0 (value 1)

	// Example 3: Minimum at the end
	A3 := []int{10, 8, 6, 2}
	fmt.Println("Local Minimum Index:", findLocalMinimum(A3)) // Output: 3 (value 2)
}
