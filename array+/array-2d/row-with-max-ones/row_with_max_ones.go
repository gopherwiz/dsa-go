/*
Problem Description
Given a binary sorted matrix A of size N x N. Find the row with the maximum number of 1.
NOTE:
If two rows have the maximum number of 1 then return the row which has a lower index.
Rows are numbered from top to bottom and columns are numbered from left to right.
Assume 0-based indexing.
Assume each row to be sorted by values.
Expected time complexity is O(rows + columns).


Example Input
Input 1:
 A = [   [0, 1, 1]
         [0, 0, 1]
         [0, 1, 1]   ]
Input 2:
 A = [   [0, 0, 0, 0]
         [0, 0, 0, 1]
         [0, 0, 1, 1]
         [0, 1, 1, 1]    ]


Example Output
Output 1:
 0
Output 2:
 3


Example Explanation
Explanation 1:
 Row 0 has maximum number of 1s.
Explanation 2:
 Row 3 has maximum number of 1s.
*/

package main

import (
	"fmt"
)

func Solve(A [][]int) int {
	n := len(A)

	r := 0
	c := n - 1
	index := 0
	for r < n && c >= 0 {
		if A[r][c] == 1 {
			index = r
			c--
		} else if A[r][c] == 0 {
			r++
		}
	}

	return index
}

func main() {
	A := [][]int{{0, 0, 0, 1}, {0, 0, 1, 1}, {0, 1, 1, 1}, {0, 1, 1, 1}}

	fmt.Println(Solve(A))
}
