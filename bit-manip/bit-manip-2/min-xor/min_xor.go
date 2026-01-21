/*
Problem Description
Given an integer array A of N integers, find the pair of integers in the array which have minimum XOR value. Report the minimum XOR value.


Example Input
Input 1:
 A = [0, 2, 5, 7]
Input 2:
 A = [0, 4, 7, 9]


Example Output
Output 1:
 2
Output 2:
 3


Example Explanation
Explanation 1:
 0 xor 2 = 2
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

// When you XOR two numbers, you are finding where their bits are different
// If the numbers are close in value, their binary patterns will be very similar, especially on the left side where the bits are most important
// To get a small XOR result, the two numbers must be very close to each other in value

func Solve(A []int) int {
	n := len(A)
	sort.Ints(A) // By sorting the array, we guarantee that the numbers that are closest in value will become immediate neighbors in the list

	minXor := math.MaxInt
	for i := 0; i < n-1; i++ {
		xor := A[i] ^ A[i+1]
		if xor < minXor {
			minXor = xor
		}
	}

	return minXor
}

func main() {
	A := []int{0, 2, 5, 7}

	fmt.Println(Solve(A))
}
