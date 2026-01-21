/*
Given an array of integers A, find and return the product array of the same size where the ith element of the product array will be equal to the product of all the elements divided by the ith element of the array.
Note: It is always possible to form the product array with integer (32 bit) values. Solve it without using the division operator.

For Example
Input 1:
    A = [1, 2, 3, 4, 5]
Output 1:
    [120, 60, 40, 30, 24]

Input 2:
    A = [5, 1, 10, 1]
Output 2:
    [10, 50, 5, 50]
*/

package main

import "fmt"

func SolveWithPrefixSuffix(A []int) []int {
	n := len(A)

	prefix := make([]int64, n)
	prefix[0] = int64(A[0])
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * int64(A[i])
	}

	suffix := make([]int64, n)
	suffix[n-1] = int64(A[n-1])
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * int64(A[i])
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			result[i] = int(suffix[1])
		} else if i == n-1 {
			result[i] = int(prefix[n-2])
		} else {
			result[i] = int(prefix[i-1] * suffix[i+1])
		}
	}

	return result
}

func main() {
	A := []int{5, 1, 10, 1}

	fmt.Println(SolveWithPrefixSuffix(A))
}
