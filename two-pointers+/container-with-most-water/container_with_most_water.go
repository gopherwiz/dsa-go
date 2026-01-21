package main

import (
	"fmt"
)

func maxArea(A []int) int {
	n := len(A)

	// Edge Case: At least two lines are required to form a container
	if n < 2 {
		return 0
	}

	left := 0
	right := n - 1
	maxA := 0
	for left < right {
		// Calculate width
		width := right - left

		// Calculate height (limited by the shorter line)
		height := 0
		if A[left] < A[right] {
			height = A[left]
		} else {
			height = A[right]
		}

		// Calculate area and update maximum
		currentArea := width * height
		if currentArea > maxA {
			maxA = currentArea
		}

		// LOGIC: Move the pointer that points to the shorter line
		/*
			Crucial Decision: Which pointer do you move?
				If you move the pointer pointing to the taller line,
				the height of your new container will still be limited by the shorter line.
				You gain nothing but lose width.

				If you move the pointer pointing to the shorter line,
				you have a chance of finding a significantly taller line
				that compensates for the loss in width.
		*/
		if A[left] < A[right] {
			left++
		} else {
			right--
		}
	}

	return maxA
}

func main() {
	// Test Case 1: [1, 5, 4, 3]
	// Width (5 to 3) is 2. Height is min(5, 3) = 3. Area = 6.
	fmt.Printf("Input: [1, 5, 4, 3] | Output: %d\n", maxArea([]int{1, 5, 4, 3}))

	// Test Case 2: Standard case
	fmt.Printf("Input: [1, 8, 6, 2, 5, 4, 8, 3, 7] | Output: %d\n", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	// Test Case 3: Single element
	fmt.Printf("Input: [1] | Output: %d\n", maxArea([]int{1}))
}
