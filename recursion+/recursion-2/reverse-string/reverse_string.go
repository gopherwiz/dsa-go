/*
Given a string, reverse it using recursion
*/

package main

import "fmt"

func ReverseString(s string) string {
	// Convert to rune slice to handle multi-byte characters (Unicode)
	runes := []rune(s)

	// Start recursion from the outer edges
	reverseRecursive(runes, 0, len(runes)-1)

	return string(runes)
}

func reverseRecursive(runes []rune, left int, right int) {
	// Base Case: When pointers meet or cross in the middle
	if left >= right {
		return
	}

	// 1. Swap the elements at left and right
	runes[left], runes[right] = runes[right], runes[left]

	// 2. Recursive call: Move inward toward the middle
	reverseRecursive(runes, left+1, right-1)
}

func main() {
	input := "PropertyFinder"
	fmt.Printf("Original: %s\n", input)
	fmt.Printf("Reversed: %s\n", ReverseString(input))
}
