/*
Problem Description
Given an array A of length N, your task is to find the maximum possible sum of any non-empty contiguous subarray.

In other words, among all possible subarrays of A, determine the one that yields the highest sum and return that sum.

Example Input
Input 1:
 A = [1, 2, 3, 4, -10]
Input 2:
 A = [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Example Output
Output 1:
 10
Output 2:
 6

Example Explanation
Explanation 1:
 The subarray [1, 2, 3, 4] has the maximum possible sum of 10.
Explanation 2:
 The subarray [4,-1,2,1] has the maximum possible sum of 6.
*/

package main

import (
	"fmt"
	"math"
)

func SolveWithPrefix(A []int) int {
	n := len(A)

	prefix := make([]int64, n)
	prefix[0] = int64(A[0])
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + int64(A[i])
	}

	sum := int64(0)
	maxSum := int64(math.MinInt64)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if i == 0 {
				sum = prefix[j]
			} else {
				sum = prefix[j] - prefix[i-1]
			}
			
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return int(maxSum)
}

func SolveWithKadane(A []int) int {
	n := len(A)

	sum := int64(0)
	maxSum := int64(math.MinInt64)
	for i := 0; i < n; i++ {
		sum += int64(A[i])
		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return int(maxSum)
}

func main() {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(SolveWithPrefix(A))
	fmt.Println(SolveWithKadane(A))
}
