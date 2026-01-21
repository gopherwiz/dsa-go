package main

import (
	"fmt"
	"sort"
)

func countPairsUnique(A []int, k int) int {
	// Step 1: Sorting is required for the two-pointer approach
	// If the problem guarantees a sorted array, you can skip this.
	sort.Ints(A)

	l, r := 0, len(A)-1
	count := 0

	for l < r {
		currentSum := A[l] + A[r]

		if currentSum == k {
			count++
			l++ // Move both pointers since elements are unique
			r--
		} else if currentSum < k {
			l++ // Sum too small, increase the smaller value
		} else {
			r-- // Sum too large, decrease the larger value
		}
	}

	return count
}

func main() {
	A := []int{1, 2, 3, 4, 5, 6}
	k := 7
	// Expected pairs: (1,6), (2,5), (3,4)
	fmt.Printf("Number of pairs with sum %d: %d\n", k, countPairsUnique(A, k))
}
