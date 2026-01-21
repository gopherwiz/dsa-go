/*
Problem Description
Given an array of integers A. If i < j and A[i] > A[j], then the pair (i, j) is called an inversion of A. Find the total number of inversions of A modulo (109 + 7).


Problem Constraints
1 <= length of the array <= 105
1 <= A[i] <= 109


Input Format
The only argument given is the integer array A.


Output Format
Return the number of inversions of A modulo (109 + 7).


Example Input
Input 1:
A = [1, 3, 2]
Input 2:
A = [3, 4, 1, 2]


Example Output
Output 1:
1
Output 2:
4


Example Explanation
Explanation 1:
The pair (1, 2) is an inversion as 1 < 2 and A[1] > A[2]
Explanation 2:
The pair (0, 2) is an inversion as 0 < 2 and A[0] > A[2]
The pair (0, 3) is an inversion as 0 < 3 and A[0] > A[3]
The pair (1, 2) is an inversion as 1 < 2 and A[1] > A[2]
The pair (1, 3) is an inversion as 1 < 3 and A[1] > A[3]
*/

package main

// SolveInversions returns the total number of inversions in the array
func SolveInversions(A []int) int {
	_, count := mergeSortWithCount(A)
	return count
}

func mergeSortWithCount(A []int) ([]int, int) {
	if len(A) <= 1 {
		return A, 0
	}

	mid := len(A) / 2

	// Recursively get counts from both halves
	leftSorted, leftCount := mergeSortWithCount(A[:mid])
	rightSorted, rightCount := mergeSortWithCount(A[mid:])

	// Merge and get the "split" inversions
	merged, splitCount := mergeAndCount(leftSorted, rightSorted)

	return merged, leftCount + rightCount + splitCount
}

func mergeAndCount(left, right []int) ([]int, int) {
	result := make([]int, len(left)+len(right))
	i, j, k, count := 0, 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			// INVERSION DETECTED:
			// If left[i] > right[j], then right[j] is smaller than
			// left[i], left[i+1]... left[end]
			result[k] = right[j]

			// The number of inversions is the number of
			// elements remaining in the left slice.
			count += len(left) - i
			j++
		}
		k++
	}

	// Append remainders
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result, count
}
