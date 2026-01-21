/*
Problem Description
You are given an array A of N elements. Sort the given array in increasing order of number of distinct factors of each element, i.e., element having the least number of factors should be the first to be displayed and the number having highest number of factors should be the last one. If 2 elements have same number of factors, then number with less value should come first.

Note: You cannot use any extra space

Problem Constraints
1 <= N <= 104
1 <= A[i] <= 104

Input Format
First argument A is an array of integers.

Output Format
Return an array of integers.

Example Input
Input 1:
A = [6, 8, 9]
Input 2:
A = [2, 4, 7]

Example Output
Output 1:
[9, 6, 8]
Output 2:
[2, 7, 4]

Example Explanation
For Input 1:
The number 9 has 3 factors, 6 has 4 factors and 8 has 4 factors.
For Input 2:
The number 2 has 2 factors, 7 has 2 factors and 4 has 3 factors.
*/

package main

import (
	"fmt"
	"sort"
)

func countFactors(n int) int {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				count += 1 // Perfect square, count only once (e.g., 3 for 9)
			} else {
				count += 2 // Count both i and n/i (e.g., 2 and 4 for 8)
			}
		}
	}

	return count
}

func solve(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		f1 := countFactors(A[i])
		f2 := countFactors(A[j])

		// Condition 1: Sort by number of factors
		if f1 != f2 {
			return f1 < f2
		}
		// Condition 2: If factors are same, sort by value
		return A[i] < A[j]
	})

	return A
}

func main() {
	fmt.Println(solve([]int{6, 8, 9})) // Output: [9 6 8]
	fmt.Println(solve([]int{2, 4, 7})) // Output: [2 7 4]
}
