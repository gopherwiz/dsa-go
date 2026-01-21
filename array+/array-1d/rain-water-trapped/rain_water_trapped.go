/*
Problem Description
Imagine a histogram where the bars' heights are given by the array A. Each bar is of uniform width, which is 1 unit. When it rains, water will accumulate in the valleys between the bars.

Your task is to calculate the total amount of water that can be trapped in these valleys.

Example:

For the Array A = [5, 4, 1, 4, 3, 2, 7], the total amount of rain water trapped is 11.

Example Input
Input 1:
A = [0, 1, 0, 2]
Input 2:
A = [1, 2]

Example Output
Output 1:
1
Output 2:
0

Example Explanation
Explanation 1:
1 unit is trapped on top of the 3rd element.
Rain Water Histogram
Explanation 2:
No water is trapped.
*/

package main

import (
	"fmt"
)

func SolveWithPrefixSuffix(A []int) int {
	n := len(A)

	prefixMaxHeight := make([]int, n)
	prefixMaxHeight[0] = A[0]
	for i := 1; i < n; i++ {
		if A[i] > prefixMaxHeight[i-1] {
			prefixMaxHeight[i] = A[i]
		} else {
			prefixMaxHeight[i] = prefixMaxHeight[i-1]
		}
	}

	suffixMaxHeight := make([]int, n)
	suffixMaxHeight[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		if A[i] > suffixMaxHeight[i+1] {
			suffixMaxHeight[i] = A[i]
		} else {
			suffixMaxHeight[i] = suffixMaxHeight[i+1]
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		leftHeight := prefixMaxHeight[i]
		rightHeight := suffixMaxHeight[i]
		minHeight := getMin(leftHeight, rightHeight)

		waterHeight := minHeight - A[i]
		total += waterHeight
	}

	return total
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	A := []int{5, 4, 1, 4, 3, 2, 7}

	fmt.Println(SolveWithPrefixSuffix(A))
}
