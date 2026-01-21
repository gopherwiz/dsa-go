/*
Problem Description
Given an integer A, generate a square matrix filled with elements from 1 to A^2 in spiral order and return the generated square matrix.


Example Input
Input 1:
1
Input 2:
2
Input 3:
5


Example Output
Output 1:
[ [1] ]
Output 2:
[ [1, 2],
  [4, 3] ]
Output 3:
[ [1,   2,  3,  4, 5],
  [16, 17, 18, 19, 6],
  [15, 24, 25, 20, 7],
  [14, 23, 22, 21, 8],
  [13, 12, 11, 10, 9] ]


Example Explanation
Explanation 1:
Only 1 is to be arranged.
Explanation 2:
1 --> 2
      |
      |
4<--- 3
*/

package main

import "fmt"

func Solve(A int) [][]int {
	result := make([][]int, A) // Rows init
	for i := range result {
		result[i] = make([]int, A) // Cols init
	}

	r := 0
	c := 0
	ct := 1
	for A > 0 {
		for i := 1; i < A; i++ {
			result[r][c] = ct
			ct++
			c++
		}

		for i := 1; i < A; i++ {
			result[r][c] = ct
			ct++
			r++
		}

		for i := 1; i < A; i++ {
			result[r][c] = ct
			ct++
			c--
		}

		for i := 1; i < A; i++ {
			result[r][c] = ct
			ct++
			r--
		}

		if A == 1 {
			result[r][c] = ct
		}

		r++
		c++
		A -= 2
	}

	return result
}

func main() {
	A := 5

	arr := Solve(A)
	for i := 0; i < A; i++ {
		for j := 0; j < A; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
