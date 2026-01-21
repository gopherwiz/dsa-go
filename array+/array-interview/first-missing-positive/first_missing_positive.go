/*
Problem Description
Given an unsorted integer array, A of size N. Find the first missing positive integer.
Note: Your algorithm should run in O(n) time and use constant space.


Example Input
Input 1:
[1, 2, 0]
Input 2:
[3, 4, -1, 1]
Input 3:
[-8, -7, -6]


Example Output
Output 1:
3
Output 2:
2
Output 3:
1


Example Explanation
Explanation 1:
A = [1, 2, 0]
First positive integer missing from the array is 3.
Explanation 2:
A = [3, 4, -1, 1]
First positive integer missing from the array is 2.
Explanation 3:
A = [-8, -7, -6]
First positive integer missing from the array is 1.
*/

package main

import (
	"fmt"
	"math"
)

func Solve(A []int) int {
	n := len(A)

	arr := make([]bool, n+1)
	for i := 0; i < n; i++ {
		number := A[i]
		if number > 0 && number <= n {
			arr[number] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !arr[i] {
			return i
		}
	}

	return n + 1
}

func SolveInPlace(A []int) int {
	n := len(A)

	for i := 0; i < n; i++ {
		if A[i] <= 0 {
			A[i] = math.MaxInt // Handle negatives upfront by converting all of them to MaxInt to fail index range check
		}
	}

	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(A[i])) - 1)
		if (index >= 0 && index <= n-1) && A[index] > 0 { // If {number(1 to n) -> index(0 to n-1)} is within array range(n) AND if not already marked
			A[index] *= -1
		}
	}

	for i := 0; i < n; i++ {
		if A[i] > 0 {
			return i + 1 // Common scenario where some number in the middle is missing
		}
	}

	return n + 1 // Edge case scenario where all numbers present in sequence, and the missing number is the next number
}

func main() {
	A := []int{3, 4, -1, 1}

	fmt.Println(Solve(A))
	fmt.Println(SolveInPlace(A))
}
