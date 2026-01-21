/*
Problem Description
Write a recursive function that checks whether string A is a palindrome or Not.
Return 1 if the string A is a palindrome, else return 0.
Note: A palindrome is a string that's the same when read forward and backward.


Example Input
Input 1:
 A = "naman"
Input 2:
 A = "strings"


Example Output
Output 1:
 1
Output 2:
 0


Example Explanation
Explanation 1:
 "naman" is a palindomic string, so return 1.
Explanation 2:
 "strings" is not a palindrome, so return 0.
*/

/*
   ----------------------------------------------------------------
   -- String Concatenation: Go vs. Java
   ----------------------------------------------------------------
   --
   -- Both Go and Java have immutable strings, but their underlying
   -- implementation leads to different performance characteristics
   -- and the need for builder patterns.
   --
   --  - Go `string`: A value type containing a header (a pointer
   --    to data and a length). Passed by copying this small header.
   --
   --  - Java `String`: An immutable object. Passed by copying the
   --    reference to the object.
   --
   -- In both languages, concatenating strings in a loop using `+`
   -- is inefficient because it creates a new string object or
   -- underlying array in every single iteration.
   --
   --    Go's `strings.Builder` is the high-performance way to build strings.
   --    Use Java's `StringBuilder` to avoid high memory allocation and garbage collection overhead.
   ----------------------------------------------------------------
*/

package main

import "fmt"

func Solve(A string) int {
	if check(A, 0, len(A)-1) {
		return 1
	}

	return 0
}

func check(s string, start, end int) bool {
	if start >= end {
		return true
	}
	if s[start] != s[end] {
		return false
	}

	return check(s, start+1, end-1)
}

func main() {
	A := "nitin"

	fmt.Println(Solve(A))
}
