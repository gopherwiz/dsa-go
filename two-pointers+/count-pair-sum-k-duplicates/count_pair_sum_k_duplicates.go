package main

import (
	"fmt"
	"sort"
)

func countPairsWithSum(A []int, k int) int {
	// 1. Array must be sorted for the two-pointer approach
	// If the problem guarantees a sorted array, you can skip this.
	sort.Ints(A)

	// 2. Define MOD as a literal integer to avoid float64 type errors
	const MOD int64 = 1000000007

	l, r := 0, len(A)-1
	var totalCount int64 = 0

	for l < r {
		currentSum := A[l] + A[r]

		if currentSum < k {
			// Sum is too small, move left pointer forward
			l++
		} else if currentSum > k {
			// Sum is too large, move right pointer backward
			r--
		} else {
			// Found a match!
			if A[l] == A[r] {
				// Special Case: All elements between l and r are the same
				// Example: [2, 2, 2, 2], k=4
				n := int64(r - l + 1)
				// Combination formula: nC2 = n * (n-1) / 2
				pairs := (n * (n - 1) / 2) % MOD
				totalCount = (totalCount + pairs) % MOD

				// Since all elements are the same, no more unique pairs exist
				break
			} else {
				// Case: A[l] and A[r] are different; count occurrences of each
				leftVal := A[l]
				cntL := int64(0)
				for A[l] == leftVal {
					cntL++
					l++
				}

				rightVal := A[r]
				cntR := int64(0)
				for A[r] == rightVal {
					cntR++
					r--
				}

				// The total pairs for these specific values is cntL * cntR
				pairs := (cntL * cntR) % MOD
				totalCount = (totalCount + pairs) % MOD

				// Pointers are already moved to the next unique values by the inner loops
			}
		}
	}

	return int(totalCount)
}

func main() {
	// Example 1: Standard duplicates (3 ones, 2 twos) -> 3 * 2 = 6 pairs
	A1 := []int{1, 2, 1, 2, 1}
	k1 := 3
	fmt.Printf("Example 1 Count: %d\n", countPairsWithSum(A1, k1))

	// Example 2: Identical elements (4 twos) -> 4C2 = 6 pairs
	A2 := []int{2, 2, 2, 2}
	k2 := 4
	fmt.Printf("Example 2 Count: %d\n", countPairsWithSum(A2, k2))

	// Example 3: Multiple distinct pairs (1,9) and (3,7) -> 1 + 1 = 2 pairs
	A3 := []int{1, 3, 7, 9}
	k3 := 10
	fmt.Printf("Example 3 Count: %d\n", countPairsWithSum(A3, k3))
}
