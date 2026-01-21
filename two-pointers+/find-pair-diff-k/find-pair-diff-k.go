package main

import (
	"fmt"
	"sort"
)

func solve(A []int, B int) int {
	sort.Ints(A)
	n := len(A)

	count := 0
	left := 0
	right := 1
	for right < n {
		diff := A[right] - A[left]

		if diff == B && left != right {
			count++

			// Capture the values we just used
			valL, valR := A[left], A[right]

			// Move pointers and SKIP all duplicates of these values
			// This ensures we only count UNIQUE pairs
			for left < n && A[left] == valL {
				left++
			}
			for right < n && A[right] == valR {
				right++
			}
		} else if diff < B {
			right++
		} else {
			left++
		}

		// Ensure pointers don't overlap
		if left == right {
			right++
		}
	}

	return count
}

func main() {
	// Test Case 1: Standard
	fmt.Println(solve([]int{1, 5, 3, 4, 2}, 3)) // Output: 2

	// Test Case 2: Duplicates
	// Pairs are (0,4), (4,8), (8,12), (12,16), (16,20)
	fmt.Println(solve([]int{8, 12, 16, 4, 0, 20}, 4)) // Output: 5

	// Test Case 3: B = 0 with many duplicates
	// Only unique pairs (1,1) and (2,2) should be counted
	fmt.Println(solve([]int{1, 1, 1, 2, 2}, 0)) // Output: 2
}
