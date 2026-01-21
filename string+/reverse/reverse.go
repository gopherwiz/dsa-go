package main

import "fmt"

// reverseString takes a string and returns its reversed version.
// Using bytes makes it efficient for ASCII text.
func reverseString(s string) string {
	// 1. Convert to a mutable byte slice
	// Strings in Go are read-only; slices are changeable.
	bytes := []byte(s)

	l := 0
	r := len(bytes) - 1

	// 2. Standard two-pointer swap
	for l < r {
		// Swap the bytes in-place
		bytes[l], bytes[r] = bytes[r], bytes[l]

		// Move pointers toward the middle
		l++
		r--
	}

	// 3. Convert the byte slice back to a string
	return string(bytes)
}

func main() {
	// Test Case 1: Standard word
	input1 := "golang"
	fmt.Printf("Original: %s | Reversed: %s\n", input1, reverseString(input1))

	// Test Case 2: Palindrome (should look the same)
	input2 := "racecar"
	fmt.Printf("Original: %s | Reversed: %s\n", input2, reverseString(input2))

	// Test Case 3: Mixed case and numbers
	input3 := "Go123!"
	fmt.Printf("Original: %s | Reversed: %s\n", input3, reverseString(input3))
}
