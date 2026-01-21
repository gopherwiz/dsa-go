/*
Problem Description
You are given an integer array A of size N.
You have to perform B operations. In one operation, you can remove either the leftmost or the rightmost element of the array A.
Find and return the maximum possible sum of the B elements that were removed after the B operations.
NOTE: Suppose B = 3, and array A contains 10 elements, then you can:
Remove 3 elements from front and 0 elements from the back, OR
Remove 2 elements from front and 1 element from the back, OR
Remove 1 element from front and 2 elements from the back, OR
Remove 0 elements from front and 3 elements from the back.

Example Input
Input 1:
 A = [5, -2, 3 , 1, 2]
 B = 3
Input 2:
 A = [ 2, 3, -1, 4, 2, 1 ]
 B = 4

Example Output
Output 1:
 8
Output 2:
 9

Example Explanation
Explanation 1:
 Remove element 5 from front and element (1, 2) from back so we get 5 + 1 + 2 = 8
Explanation 2:
 Remove the first element and the last 3 elements. So we get 2 + 4 + 2 + 1 = 9
*/

package main

import (
	"fmt"
)

func SolveWithPrefixSuffix(A []int, B int) int {
	n := len(A)

	prefix := make([]int64, n)
	prefix[0] = int64(A[0])
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + int64(A[i])
	}

	suffix := make([]int64, n)
	suffix[n-1] = int64(A[n-1])
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] + int64(A[i])
	}

	maxSum := getMax(prefix[B-1], suffix[n-B]) // This is the initial edge cases, picking all from front vs. all from back
	for i := 1; i < B; i++ {
		frontPick := prefix[i-1]
		backPick := suffix[n-B+i]
		pick := frontPick + backPick
		maxSum = getMax(maxSum, pick)
	}

	return int(maxSum)
}

func getMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func SolveWithSlidingWindow(A []int, B int) int {
	n := len(A)
	k := n - B // Sliding window length is not B, it is (n - B)

	totalSum := int64(0)
	for i := 0; i < n; i++ {
		totalSum += int64(A[i])
	}

	// We start with the first window
	start := 0
	end := k - 1
	sum := int64(0)
	for i := start; i <= end; i++ {
		sum += int64(A[i])
	}

	minSum := sum
	for end < n-1 {
		sum += int64(A[end+1])
		sum -= int64(A[start])

		if sum < minSum {
			minSum = sum
		}

		start++
		end++
	}

	return int(totalSum - minSum)
}

func main() {
	A := []int{2, 3, -1, 4, 2, 1}
	B := 4

	fmt.Println(SolveWithPrefixSuffix(A, B))
	fmt.Println(SolveWithSlidingWindow(A, B))
}
