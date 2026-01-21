package main

import "fmt"

func solve(A []int, B int) []int {
	n := len(A)

	l, r := 0, 0
	currentSum := 0
	// Use a single loop to manage the 'right' boundary
	for r < n {
		// 1. Expand: Add the current element to the sum
		currentSum += A[r]

		// 2. Contract: While the sum is too large, move the left pointer
		// This keeps the "ladder" clean by ensuring currentSum <= B
		for currentSum > B && l < r {
			currentSum -= A[l]
			l++
		}

		// 3. Match: Check if we've hit the target
		if currentSum == B {
			result := make([]int, r-l+1)
			copy(result, A[l:r+1])

			return result
		}

		// 4. Continue: Increment right to expand in the next iteration
		r++
	}

	return []int{-1}
}

func main() {
	A := []int{1, 2, 3, 4, 5}
	B := 5
	fmt.Println("Result:", solve(A, B)) // Output: [2, 3]
}
