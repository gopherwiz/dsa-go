package main

import "fmt"

// reverse flips the elements of an integer slice in-place
func reverse(s []int) {
	l := 0
	r := len(s) - 1

	for l < r {
		// Swap elements
		s[l], s[r] = s[r], s[l]

		// Move pointers toward the middle
		l++
		r--
	}
}

func main() {
	// 1. Initialize a slice of integers
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("Original slice:", nums)

	// 2. Call the reverse function
	// This modifies 'nums' directly in memory
	reverse(nums)

	// 3. Print the result to verify
	fmt.Println("Reversed slice:", nums)

	// --- Additional Test Cases ---

	// Test Case: Odd number of elements
	oddNums := []int{1, 2, 3}
	reverse(oddNums)
	fmt.Println("Odd reversal:  ", oddNums) // Expected: [3 2 1]

	// Test Case: Empty or single element
	single := []int{100}
	reverse(single)
	fmt.Println("Single element:", single) // Expected: [100]
}
