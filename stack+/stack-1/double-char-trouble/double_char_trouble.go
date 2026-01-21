package main

import (
	"fmt"
)

// solve removes consecutive identical pairs from the string A.
func solve(A string) string {
	// We use a byte slice as a stack.
	// Since the problem involves ASCII characters, []byte is efficient.
	stack := make([]byte, 0, len(A))

	for i := 0; i < len(A); i++ {
		curr := A[i]

		// If stack is not empty and the top element matches the current element
		if len(stack) > 0 && stack[len(stack)-1] == curr {
			// Pop the element (remove the pair)
			stack = stack[:len(stack)-1]
		} else {
			// Push the current element onto the stack
			stack = append(stack, curr)
		}
	}

	// In Go, converting a []byte slice to a string is straightforward.
	// Because we processed from left to right, the stack is already in order.
	return string(stack)
}

func main() {
	// Test Case 1: "abccbc" -> "abbc" -> "ac"
	input1 := "abccbc"
	fmt.Printf("Input: %s | Output: %s\n", input1, solve(input1))

	// Test Case 2: "ab" -> "ab"
	input2 := "ab"
	fmt.Printf("Input: %s | Output: %s\n", input2, solve(input2))

	// Test Case 3: "aaaaaaaa" -> ""
	input3 := "aaaaaaaa"
	fmt.Printf("Input: %s | Output: %s\n", input3, solve(input3))
}
