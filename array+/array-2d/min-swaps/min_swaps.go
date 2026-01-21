/*
Problem Description
Given an array of integers A and an integer B, find and return the minimum number of swaps required to bring all the numbers less than or equal to B together.
Note: It is possible to swap any two elements, not necessarily consecutive.

# Example Input

Input 1:

	A = [1, 12, 10, 3, 14, 10, 5]
	B = 8

Input 2:

	A = [5, 17, 100, 11]
	B = 20

# Example Output

Output 1:

	2

Output 2:

	1

# Example Explanation

Explanation 1:

	A = [1, 12, 10, 3, 14, 10, 5]
	After swapping  12 and 3, A => [1, 3, 10, 12, 14, 10, 5].
	After swapping  the first occurence of 10 and 5, A => [1, 3, 5, 12, 14, 10, 10].
	Now, all elements less than or equal to 8 are together.

Explanation 2:

	A = [5, 17, 100, 11]
	After swapping 100 and 11, A => [5, 17, 11, 100].
	Now, all elements less than or equal to 20 are together.
*/

package main

import "fmt"

func Solve(A []int, B int) int {
	n := len(A)

	k := 0
	for i := 0; i < n; i++ {
		if A[i] <= B {
			k++
		}
	}

	start := 0
	end := k - 1
	swapCount := 0
	for i := start; i <= end; i++ {
		if A[i] > B {
			swapCount++ // Needs to be swapped in the current window
		}
	}

	minSwap := swapCount
	for end < n-1 {
		if A[start] > B { // Exclude from front
			swapCount--
		}

		if A[end+1] > B { // Include from back
			swapCount++
		}

		if swapCount < minSwap {
			minSwap = swapCount
		}
		start++
		end++
	}

	return minSwap
}

func main() {
	A := []int{5, 17, 100, 11}
	B := 20

	fmt.Println(Solve(A, B))
}
