/*
Problem Description
You are given a binary string A(i.e., with characters 0 and 1) consisting of characters A1, A2, ..., AN. In a single operation, you can choose two indices, L and R, such that 1 ≤ L ≤ R ≤ N and flip the characters AL, AL+1, ..., AR. By flipping, we mean changing character 0 to 1 and vice-versa.
Your aim is to perform ATMOST one operation such that in the final string number of 1s is maximized.
If you don't want to perform the operation, return an empty array. Else, return an array consisting of two elements denoting L and R. If there are multiple solutions, return the lexicographically smallest pair of L and R.
NOTE: Pair (a, b) is lexicographically smaller than pair (c, d) if a < c or, if a == c and b < d.


Example Input
Input 1:
A = "010"
Input 2:
A = "111"


Example Output
Output 1:
[1, 1]
Output 2:
[]


Example Explanation
Explanation 1:
A = "010"

Pair of [L, R] | Final string
_______________|_____________
[1 1]          | "110"
[1 2]          | "100"
[1 3]          | "101"
[2 2]          | "000"
[2 3]          | "001"

We see that two pairs [1, 1] and [1, 3] give same number of 1s in final string. So, we return [1, 1].
Explanation 2:
No operation can give us more than three 1s in final string. So, we return empty array [].
*/

/*
   ----------------------------------------------------------------
   -- The Go `strings` Package: A Quick Reference
   ----------------------------------------------------------------
   --
   -- The `strings` package provides a rich set of functions
   -- for string manipulation.
   --
   ----------------------------------------------------------------

   //-- Searching --//

   // Contains: Checks if a substring is present.
   // func Contains(s, substr string) bool
   // Example: strings.Contains("hello world", "world") // true

   // HasPrefix: Checks if a string starts with a prefix.
   // func HasPrefix(s, prefix string) bool
   // Example: strings.HasPrefix("file.pdf", "file") // true

   // HasSuffix: Checks if a string ends with a suffix.
   // func HasSuffix(s, suffix string) bool
   // Example: strings.HasSuffix("photo.jpg", ".jpg") // true

   // Index: Finds the first index of a substring.
   // func Index(s, substr string) int
   // Example: strings.Index("banana", "na") // 2


   //-- Splitting & Joining --//

   // Split: Splits a string into a slice of strings.
   // func Split(s, sep string) []string
   // Example: strings.Split("a,b,c", ",") // []string{"a", "b", "c"}

   // Join: Joins a slice of strings into a single string.
   // func Join(elems []string, sep string) string
   // Example: strings.Join([]string{"a", "b"}, "-") // "a-b"


   //-- Modification --//

   // Replace: Replaces occurrences of a substring.
   // func Replace(s, old, new string, n int) string
   // Example: strings.Replace("one two two", "two", "three", 1) // "one three two"

   // ToLower / ToUpper: Converts the case of a string.
   // func ToLower(s string) string
   // func ToUpper(s string) string
   // Example: strings.ToUpper("hello") // "HELLO"

   // TrimSpace: Removes leading and trailing whitespace.
   // func TrimSpace(s string) string
   // Example: strings.TrimSpace("  hello world  ") // "hello world"
*/

package main

import (
	"fmt"
	"math"
)

func Solve(A string) []int {
	n := len(A)

	arr := make([]int, n) // 001001 -> {1, 1, -1, 1, 1, -1}
	for index, char := range A {
		if string(char) == "1" {
			arr[index] = -1
		} else {
			arr[index] = 1
		}
	}

	sum := int64(0)
	maxSum := int64(math.MinInt64)
	l := 0
	r := 0
	start := 0
	for i := 0; i < n; i++ {
		sum += int64(arr[i])
		if sum > maxSum {
			maxSum = sum
			l = start
			r = i
		}

		if sum < 0 {
			sum = 0
			start = i + 1 // Later on, apply this start to l only if there is a maxSum
		}
	}

	if maxSum > 0 {
		return []int{l + 1, r + 1}
	}

	return []int{}
}

func main() {
	A := "100101101" // 100101101 001001

	fmt.Println(Solve(A))
}
