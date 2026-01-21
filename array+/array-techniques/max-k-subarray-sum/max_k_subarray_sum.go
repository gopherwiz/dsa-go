/*
   Max subarray sum in all k length subarrays ~ sliding-window or prefix
*/

package main

import (
	"fmt"
	"math"
)

func SolveWithPrefix(A []int, k int) int {
	n := len(A)

	prefix := make([]int64, n)
	prefix[0] = int64(A[0])
	for i := 1; i < len(A); i++ {
		prefix[i] = prefix[i-1] + int64(A[i])
	}

	var sum int64
	maxSum := int64(math.MinInt64)
	for i := 0; i <= n-k; i++ {
		start := i
		end := i + k - 1 // The subarray starting at `i` will end at `i + k - 1`.

		if i == 0 {
			sum = prefix[end]
		} else {
			sum = prefix[end] - prefix[start-1]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return int(maxSum)
}

func SolveWithSlidingWindow(A []int, k int) int {
	n := len(A)

	start := 0
	end := k - 1
	sum := int64(0)
	maxSum := int64(math.MinInt64)

	// first window sum
	for i := start; i <= end; i++ {
		sum += int64(A[i])
	}

	// next windows
	for end < n-1 {
		sum = sum - int64(A[start]) + int64(A[end+1])

		if sum > maxSum {
			maxSum = sum
		}

		start++
		end++
	}

	return int(maxSum)
}

func main() {
	A := []int{-7, 1, 5, -1, -4, 15, 0}
	k := 3

	fmt.Println(SolveWithPrefix(A, k))
	fmt.Println(SolveWithSlidingWindow(A, k))
}
