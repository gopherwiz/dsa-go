package main

import "fmt"

// isPalindrome checks for a palindrome using raw byte indexing.
// Note: This approach only works correctly for single-byte (ASCII) characters.
func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		// Compare bytes at the left and right positions
		if s[l] != s[r] {
			return false
		}

		// Move pointers toward the middle
		l++
		r--
	}

	return true
}

func main() {
	// Test Case 1: Odd length palindrome
	s1 := "racecar"
	fmt.Printf("'%s' is palindrome: %v\n", s1, isPalindrome(s1)) // true

	// Test Case 2: Even length palindrome
	s2 := "abba"
	fmt.Printf("'%s' is palindrome: %v\n", s2, isPalindrome(s2)) // true

	// Test Case 3: Not a palindrome
	s3 := "golang"
	fmt.Printf("'%s' is palindrome: %v\n", s3, isPalindrome(s3)) // false

	// Test Case 4: Empty or single char
	fmt.Printf("'' is palindrome: %v\n", isPalindrome(""))   // true
	fmt.Printf("'a' is palindrome: %v\n", isPalindrome("a")) // true
}
