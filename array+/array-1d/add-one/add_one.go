/*
Problem Description
Given a non-negative number represented as an array of digits, add 1 to the number ( increment the number represented by the digits ).
The digits are stored such that the most significant digit is at the head of the list.
NOTE: Certain things are intentionally left unclear in this question which you should practice asking the interviewer. For example: for this problem, the following are some good questions to ask :
Q: Can the input have 0's before the most significant digit. Or, in other words, is 0 1 2 3 a valid input?
A: For the purpose of this question, YES
Q: Can the output have 0's before the most significant digit? Or, in other words, is 0 1 2 4 a valid output?
A: For the purpose of this question, NO. Even if the input has zeroes before the most significant digit.


Example Input
Input 1:
[1, 2, 3]


Example Output
Output 1:
[1, 2, 4]


Example Explanation
Explanation 1:
Given vector is [1, 2, 3].
The returned vector should be [1, 2, 4] as 123 + 1 = 124.
*/

/*
   ----------------------------------------------------------------
   -- The Go Slice Operator `[:]`
   ----------------------------------------------------------------
   --
   -- The operator `a[low:high]` creates a new slice from a portion
   -- of an existing slice or array.
   --
   -- It includes elements from index `low` up to, but not
   -- including, index `high`.
   --
   -- IMPORTANT: The new slice shares the same underlying array as
   -- the original. Modifying the sub-slice will modify the
   -- original slice.
   --
   ----------------------------------------------------------------

   //-- Syntax Variations --//

   a[low:high]  ->  Elements from index `low` to `high-1`
   a[:high]     ->  Elements from the start (index 0) to `high-1`
   a[low:]      ->  Elements from index `low` to the end
   a[:]         ->  A slice of the entire original slice
*/

/*
   ----------------------------------------------------------------
   -- Prepending to a Slice in Go
   ----------------------------------------------------------------
   --
   -- Go does not have a dedicated `prepend` function. The idiomatic
   -- and most efficient way to add elements to the beginning of a
   -- slice is to use the built-in `append` function combined with
   -- the `...` operator.
   --
   -- The `...` operator (the "spread" operator) unpacks a slice
   -- into individual elements.
   --
   ----------------------------------------------------------------

   //-- Syntax --//
   newSlice := append(elementsToPrepend, originalSlice...)
*/

package main

import (
	"fmt"
)

func Solve(A []int) []int {
	n := len(A)
	carry := 1

	// Core logic
	for i := n - 1; i >= 0; i-- {
		sum := A[i] + carry
		A[i] = sum % 10
		carry = sum / 10

		if carry == 0 { // Optional early exit
			break
		}
	}

	// If carry was propagated to the last digit, need to add it as a new digit to the front
	if carry > 0 {
		A = append([]int{carry}, A...) // Prepend the carry
	}

	// Remove leading zeroes
	for i := 0; i < len(A); i++ {
		if A[i] != 0 {
			return A[i:]
		}
	}

	return A
}

func main() {
	A := []int{0, 0, 9, 9, 9} // {0, 0, 9, 9, 9}

	fmt.Println(Solve(A))
}
