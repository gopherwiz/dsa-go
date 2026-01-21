package main

import (
	"fmt"
	"sort"
)

func hasPairWithSum(A []int, k int) bool {
	// Step 1: Array must be sorted for two-pointer to work
	// If the problem guarantees a sorted array, you can skip this.
	sort.Ints(A)

	l := 0
	r := len(A) - 1

	// We use l < r because we need TWO distinct elements
	for l < r {
		currentSum := A[l] + A[r]

		if currentSum == k {
			return true
		}

		if currentSum < k {
			l++ // Need a larger sum
		} else {
			r-- // Need a smaller sum
		}
	}

	return false
}

func main() {
	A1 := []int{1, 2, 4, 7, 11, 15}
	k1 := 15
	fmt.Printf("Target %d: %v\n", k1, hasPairWithSum(A1, k1)) // Output: true (4+11)

	A2 := []int{1, 4, 4, 5}
	k2 := 8
	fmt.Printf("Target %d: %v\n", k2, hasPairWithSum(A2, k2)) // Output: true (4+4)
}
