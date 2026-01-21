/*
Problem Description
Given an integer array A of size N denoting collection of numbers , return all possible permutations.

NOTE:

No two entries in the permutation sequence should be the same.
For the purpose of this problem, assume that all the numbers in the collection are unique.
Return the answer in any order
WARNING: DO NOT USE LIBRARY FUNCTION FOR GENERATING PERMUTATIONS.
Example : next_permutations in C++ / itertools.permutations in python.
If you do, we will disqualify your submission retroactively and give you penalty points.


Problem Constraints
1 <= N <= 9



Input Format
Only argument is an integer array A of size N.



Output Format
Return a 2-D array denoting all possible permutation of the array.



Example Input
A = [1, 2, 3]


Example Output
[ [1, 2, 3]
  [1, 3, 2]
  [2, 1, 3]
  [2, 3, 1]
  [3, 1, 2]
  [3, 2, 1] ]


Example Explanation
All the possible permutation of array [1, 2, 3].
*/

package main

import (
	"fmt"
)

func Solve(A []int) [][]int {
	var permutations [][]int

	backtrack(&permutations, []int{}, make([]bool, len(A)), A)

	return permutations
}

func backtrack(result *[][]int, permutation []int, visited []bool, A []int) {
	if len(permutation) == len(A) {
		*result = append(*result, append([]int(nil), permutation...))
		return
	}

	for i := 0; i < len(A); i++ {
		if !visited[i] {
			// Do
			visited[i] = true
			permutation = append(permutation, A[i])

			backtrack(result, permutation, visited, A)

			// Backtrack
			visited[i] = false
			permutation = permutation[:len(permutation)-1]
		}
	}
}

func main() {
	A := []int{1, 2, 3}
	result := Solve(A)
	fmt.Println(result)
}
