package main

import (
	"fmt"
)

// The NextGreaterElement finds the next greater element to the right of each element.
func NextGreaterElement(A []int) []int {
	n := len(A)
	if n == 0 {
		return []int{}
	}

	// G will store the result array
	G := make([]int, n)
	// stack will store elements that could be the next greater for previous elements
	stack := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		current := A[i]

		// 1. Maintain monotonicity:
		// Remove elements from the stack that are NOT greater than current.
		// These elements are useless now because 'current' is greater and further left.
		for len(stack) > 0 && stack[len(stack)-1] <= current {
			stack = stack[:len(stack)-1]
		}

		// 2. Identify the result for index i
		if len(stack) == 0 {
			G[i] = -1 // No greater element found to the right
		} else {
			G[i] = stack[len(stack)-1] // Top of stack is the next greater
		}

		// 3. Push current element to be a candidate for indices to the left
		stack = append(stack, current)
	}

	return G
}

func main() {
	// Case 1: [4, 5, 2, 10]
	A1 := []int{4, 5, 2, 10}
	fmt.Printf("Input: %v | Output: %v\n", A1, NextGreaterElement(A1))
	// Expected: [5, 10, 10, -1]

	// Case 1: [3, 2, 1]
	A2 := []int{3, 2, 1}
	fmt.Printf("Input: %v | Output: %v\n", A2, NextGreaterElement(A2))
	// Expected: [-1, -1, -1]
}
