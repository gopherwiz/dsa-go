/*
Problem Description
Given a matrix of integers A of size N x M and an integer B.
In the given matrix every row and column is sorted in non-decreasing order. Find and return the position of B in the matrix in the given form:
If A[i][j] = B then return (i * 1009 + j)
If B is not present return -1.

Note 1: Rows are numbered from top to bottom and columns are numbered from left to right.
Note 2: If there are multiple B in A then return the smallest value of i*1009 +j such that A[i][j]=B.
Note 3: Expected time complexity is linear
Note 4: Use 1-based indexing


Example Input
Input 1:-
A = [[1, 2, 3]
     [4, 5, 6]
     [7, 8, 9]]
B = 2
Input 2:-
A = [[1, 2]
     [3, 3]]
B = 3


Example Output
Output 1:-
1011
Output 2:-
2019


Example Explanation
Expanation 1:-
A[1][2] = 2
1 * 1009 + 2 = 1011
Explanation 2:-
A[2][1] = 3
2 * 1009 + 1 = 2019
A[2][2] = 3
2 * 1009 + 2 = 2020
The minimum value is 2019
*/

package main

import (
	"fmt"
	"math"
)

func Solve(A [][]int, B int) int {
	R := len(A)
	C := len(A[0])

	r := 0
	c := C - 1
	minResult := math.MaxInt
	for r < R && c >= 0 {
		item := A[r][c]
		if item == B {
			result := (r+1)*1009 + (c + 1)
			if result < minResult {
				minResult = result
			}
			c-- // Since we want to find another match at lower {r,c} if possible
		} else if item < B {
			r++
		} else {
			c--
		}
	}

	if minResult == math.MaxInt {
		return -1
	}

	return minResult
}

func main() {
	A := [][]int{{2, 8, 8, 8}, {2, 8, 8, 8}, {2, 8, 8, 8}} //{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	B := 8                                                 // 2

	fmt.Println(Solve(A, B))
}
