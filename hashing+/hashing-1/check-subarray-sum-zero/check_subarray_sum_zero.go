/*
Problem Description
Given an array of integers A, find and return whether the given array contains a non-empty subarray with a sum equal to 0.
If the given array contains a sub-array with sum zero return 1, else return 0.

Problem Constraints
1 <= |A| <= 100000
-10^9 <= A[i] <= 10^9

Input Format
The only argument given is the integer array A.

Output Format
Return whether the given array contains a subarray with a sum equal to 0.

Example Input
Input 1:

	A = [1, 2, 3, 4, 5]

Input 2:

	A = [4, -1, 1]

Example Output
Output 1:

	0

Output 2:

	1

Example Explanation
Explanation 1:

	No subarray has sum 0.

Explanation 2:

	The subarray [-1, 1] has sum 0.
*/

/*
prefix[i] == 0
prefix[j] - prefix[i-1] == 0
*/

package main

import "fmt"

// SolveWithSubarraysSumCheck is Bruteforce
func SolveWithSubarraysSumCheck(A []int) int {
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += A[k]
			}

			if sum == 0 {
				return 1
			}
		}
	}

	return 0
}

func SolveWithPrefix(A []int) int {
	for i := 1; i < len(A); i++ {
		A[i] = A[i] + A[i-1]
	}

	hashset := make(map[int]struct{})
	for _, ps := range A {
		if ps == 0 {
			return 1
		}

		if _, exists := hashset[ps]; exists {
			return 1
		} else {
			hashset[ps] = struct{}{}
		}
	}

	return 0
}

func SolveWithCarryForward(A []int) int {
	hashset := make(map[int]struct{})
	hashset[0] = struct{}{} // 0 added once, it should not appear as total now

	total := 0
	for _, v := range A {
		total += v

		if _, exists := hashset[total]; exists {
			return 1
		} else {
			hashset[total] = struct{}{}
		}
	}

	return 0
}

func main() {
	A := []int{2, 1, -1, 3, 12}
	fmt.Println(SolveWithCarryForward(A))
}
