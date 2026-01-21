/*
Problem Description
Given an array A of length N where all the elements are distinct and are in the range [1, N+2].
Two numbers from the range [1, N+2] are missing from the array A. Find the two missing numbers.


Example Input
Input 1:
A = [3, 2, 4]
Input 2:
A = [5, 1, 3, 6]


Example Output
Output 1:
[1, 5]
Output 2:
[2, 4]


Example Explanation
For Input 1:
The missing numbers are 1 and 5.
For Input 2:
The missing numbers are 2 and 4.
*/

package main

import (
	"fmt"
	"sort"
)

func Solve(A []int) []int {
	n := len(A)

	xorAll := 0
	for num := 1; num <= n+2; num++ { // Extra XOR with [1,n+2] to create duplicates & xor them to 0
		xorAll ^= num
	}
	for _, num := range A {
		xorAll ^= num
	}

	// Find any set bit index in xorAll - set bit implies that at that position, the two single numbers have opposite bits
	setBitIndex := 0
	for xorAll > 0 {
		if xorAll&1 == 1 {
			break
		}

		xorAll = xorAll >> 1
		setBitIndex++
	}

	// Segregate numbers into 2 groups based on set bit & xor the groups to turn duplicate numbers in a group to 0
	num1 := 0
	num2 := 0
	for num := 1; num <= n+2; num++ { // Extra XOR with [1,n+2] to create duplicates & xor them to 0
		if num&(1<<setBitIndex) != 0 {
			num1 = num1 ^ num
		} else {
			num2 = num2 ^ num
		}
	}
	for _, num := range A {
		if num&(1<<setBitIndex) != 0 {
			num1 = num1 ^ num
		} else {
			num2 = num2 ^ num
		}
	}

	result := []int{num1, num2}
	sort.Ints(result)

	return result
}

func main() {
	A := []int{5, 1, 3, 6}

	fmt.Println(Solve(A))
}
