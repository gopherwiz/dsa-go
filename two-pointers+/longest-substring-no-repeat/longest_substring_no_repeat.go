package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	n := len(s)

	left := 0
	right := 0 // Initialize outside
	maxLen := 0
	for right < n { // Condition check
		char := s[right]

		// If character is in map and within current window
		if lastIdx, found := charMap[char]; found {
			// If the duplicate is inside the window [left, right],
			// jump the 'left' pointer to index + 1
			if lastIdx >= left {
				left = lastIdx + 1
			}
		}

		// Update map with current index
		charMap[char] = right

		// Update Max Length
		currentWindow := right - left + 1
		if currentWindow > maxLen {
			maxLen = currentWindow
		}

		right++ // Manual increment
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
}
