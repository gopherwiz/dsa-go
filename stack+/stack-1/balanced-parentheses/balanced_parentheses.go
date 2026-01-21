package main

import (
	"fmt"
)

// solve checks if the parenthesis sequence is balanced.
func solve(A string) int {
	// A slice of runes (characters) to act as our stack.
	// Since the constraint is small (100), this is very efficient.
	var stack []rune

	// Map to store matching pairs for quick lookup.
	// This is more idiomatic in Go than multiple if/else functions.
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range A {
		// Check if the character is a closing bracket
		if opening, isClosing := pairs[char]; isClosing {
			// If it's a closing bracket:
			// 1. Stack must not be empty
			// 2. The top of the stack must match the required opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return 0
			}

			// Pop the element from the stack
			stack = stack[:len(stack)-1]
		} else {
			// If it's an opening bracket, push it onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets were matched correctly.
	if len(stack) == 0 {
		return 1
	}

	return 0
}

func main() {
	// Test Case 1: Balanced
	fmt.Printf("Input: {([])} | Output: %d\n", solve("{([])}"))

	// Test Case 2: Unbalanced (Missing closing)
	fmt.Printf("Input: (){ | Output: %d\n", solve("(){"))

	// Test Case 3: Balanced (Multiple sets)
	fmt.Printf("Input: ()[] | Output: %d\n", solve("()[]"))

	// Test Case 4: Wrong Order
	fmt.Printf("Input: ([)] | Output: %d\n", solve("([)]"))
}
