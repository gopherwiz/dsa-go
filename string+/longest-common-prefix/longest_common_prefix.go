package main

import (
	"fmt"
)

/**
 * longestCommonPrefix finds the longest common starting substring
 * among an array of strings.
 * * Time Complexity: O(S) - where S is the sum of all characters in all strings.
 * Space Complexity: O(1) - we use the existing memory of the strings.
 */
func longestCommonPrefix(strs []string) string {
	// 1. If the input slice is empty, there is no common prefix
	if len(strs) == 0 {
		return ""
	}

	// 2. Iterate through each character of the first string (the "reference" string)
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		// 3. Compare this character with the character at the same position
		//    in all other strings.
		for j := 1; j < len(strs); j++ {

			// IF: We reached the end of the current string 'strs[j]'
			// OR: The character at index 'i' doesn't match
			// THEN: The prefix ends exactly before index 'i'
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	// 4. If we finish the loop, the entire first string is the common prefix
	return strs[0]
}

func main() {
	// Test Case 1: Product Categories with common naming convention
	case1 := []string{"Electronics-Mobile", "Electronics-Laptop", "Electronics-Tablet"}
	fmt.Printf("Input: %v\nResult: %q\n\n", case1, longestCommonPrefix(case1))

	// Test Case 2: Mixed categories with no common prefix
	case2 := []string{"Home-Decor", "Furniture", "Electronics"}
	fmt.Printf("Input: %v\nResult: %q\n\n", case2, longestCommonPrefix(case2))

	// Test Case 3: Strings with a very short common prefix
	case3 := []string{"dog", "racecar", "car"}
	fmt.Printf("Input: %v\nResult: %q\n\n", case3, longestCommonPrefix(case3))

	// Test Case 4: One string is a prefix of all others
	case4 := []string{"inter", "interview", "internal"}
	fmt.Printf("Input: %v\nResult: %q\n\n", case4, longestCommonPrefix(case4))

	// Test Case 5: Single string input
	case5 := []string{"onlyOne"}
	fmt.Printf("Input: %v\nResult: %q\n\n", case5, longestCommonPrefix(case5))
}
