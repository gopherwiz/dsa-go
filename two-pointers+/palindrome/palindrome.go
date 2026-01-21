package main

import "fmt"

// IsPalindrome returns true if the slice is a palindrome.
func IsPalindrome(A []int) bool {
	n := len(A)
	if n <= 1 {
		return true
	}

	left := 0
	right := n - 1

	for left < right {
		if A[left] != A[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	// Case 1: Odd length palindrome
	fmt.Println("[1, 2, 3, 2, 1]:", IsPalindrome([]int{1, 2, 3, 2, 1})) // true

	// Case 2: Even length palindrome
	fmt.Println("[1, 2, 2, 1]:", IsPalindrome([]int{1, 2, 2, 1})) // true

	// Case 3: Not a palindrome
	fmt.Println("[1, 2, 3]:", IsPalindrome([]int{1, 2, 3})) // false
}
