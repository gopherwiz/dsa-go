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

import "fmt"

const MOD = 1000000007

func Solve(A []int) int {
	if len(A) == 0 {
		return 0
	}
	// Create a temporary buffer once to avoid repeated allocations in merge
	temp := make([]int, len(A))
	return mergeSort(A, temp, 0, len(A)-1)
}

func mergeSort(A, temp []int, l, r int) int {
	if l >= r {
		return 0
	}

	m := l + (r-l)/2

	// Total count = left inversions + right inversions + merge inversions
	count := 0
	count = (count + mergeSort(A, temp, l, m)) % MOD
	count = (count + mergeSort(A, temp, m+1, r)) % MOD
	count = (count + merge(A, temp, l, m, r)) % MOD

	return count
}

func merge(A, temp []int, l, m, r int) int {
	// Equivalent to your vector<int> left and right creation
	// But we use the single 'temp' buffer to save memory
	for i := l; i <= r; i++ {
		temp[i] = A[i]
	}

	i := l     // Left half pointer
	j := m + 1 // Right half pointer
	k := l     // Original array pointer
	count := 0

	for i <= m && j <= r {
		if temp[i] <= temp[j] {
			A[k] = temp[i]
			i++
		} else {
			// INVERSION DETECTED
			// If temp[i] > temp[j], then temp[i...m] are all > temp[j]
			count = (count + (m - i + 1)) % MOD
			A[k] = temp[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i <= m {
		A[k] = temp[i]
		i++
		k++
	}
	// (Note: elements in the right half are already in place if i reaches m)

	return count
}

func main() {
	A := []int{1, 3, 2, 3, 1}
	fmt.Println(Solve(A)) // Expected Output: 3
}
