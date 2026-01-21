package main

import (
	"fmt"
	"sort"
)

// ThreeSumClosest finds the sum of three integers closest to target B.
func ThreeSumClosest(A []int, B int) int {
	n := len(A)

	// Edge Case: If we have exactly 3 elements, return their sum.
	if n == 3 {
		return A[0] + A[1] + A[2]
	}

	sort.Ints(A)                     // Sort the array to enable the two-pointer technique.
	closestSum := A[0] + A[1] + A[2] // Initialize closestSum with the first possible triplet.
	for i := 0; i < n-2; i++ {
		// Optimization: Skip duplicate values for the first element to save time.
		if i > 0 && A[i] == A[i-1] {
			continue
		}

		left := i + 1
		right := n - 1
		for left < right {
			currentSum := A[i] + A[left] + A[right]

			// If we found the exact sum, return it immediately.
			if currentSum == B {
				return currentSum
			}

			// Update closestSum if the current triplet is closer to B than the previous best.
			if Abs(B-currentSum) < Abs(B-closestSum) {
				closestSum = currentSum
			}

			// Move pointers based on how currentSum compares to target B.
			if currentSum < B {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}

// Abs is a helper function because math.Abs takes and returns float64.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 1
	A1 := []int{-1, 2, 1, -4}
	B1 := 1
	fmt.Printf("Input: %v, Target: %d | Closest Sum: %d\n", A1, B1, ThreeSumClosest(A1, B1)) // Output: 2

	// Example 2
	A2 := []int{1, 2, 3}
	B2 := 6
	fmt.Printf("Input: %v, Target: %d | Closest Sum: %d\n", A2, B2, ThreeSumClosest(A2, B2)) // Output: 6

	// Large values case
	A3 := []int{10, 2, 4, 1, 8, 3}
	B3 := 100
	fmt.Printf("Input: %v, Target: %d | Closest Sum: %d\n", A3, B3, ThreeSumClosest(A3, B3)) // Output: 22 (10+8+4)
}
