package main

import (
	"fmt"
)

// The NearestSmallerElement finds the nearest smaller element to the left of each element.
func NearestSmallerElement(A []int) []int {
	n := len(A)
	if n == 0 {
		return []int{}
	}

	// G will store the result array
	G := make([]int, n)
	// stack will store elements that could be the nearest smaller for future elements
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		current := A[i]

		// 1. Maintain monotonicity:
		// Remove elements from the stack that are NOT smaller than current.
		// These elements are useless now because 'current' is smaller and further right.
		for len(stack) > 0 && stack[len(stack)-1] >= current {
			stack = stack[:len(stack)-1]
		}

		// 2. Identify the result for index i
		if len(stack) == 0 {
			G[i] = -1 // No smaller element found to the left
		} else {
			G[i] = stack[len(stack)-1] // Top of stack is the nearest smaller
		}

		// 3. Push current element to be a candidate for indices to the right
		stack = append(stack, current)
	}

	return G
}

func main() {
	// Case 1: [4, 5, 2, 10, 8]
	A1 := []int{4, 5, 2, 10, 8}
	fmt.Printf("Input: %v | Output: %v\n", A1, NearestSmallerElement(A1))
	// Expected: [-1, 4, -1, 2, 2]

	// Case 2: Strictly decreasing [3, 2, 1]
	A2 := []int{3, 2, 1}
	fmt.Printf("Input: %v | Output: %v\n", A2, NearestSmallerElement(A2))
	// Expected: [-1, -1, -1]

	// Case 3: Strictly increasing [1, 2, 3]
	A3 := []int{1, 2, 3}
	fmt.Printf("Input: %v | Output: %v\n", A3, NearestSmallerElement(A3))
	// Expected: [-1, 1, 2]
}
