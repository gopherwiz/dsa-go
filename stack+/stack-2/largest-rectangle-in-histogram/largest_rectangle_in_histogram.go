package main

import "fmt"

/*
largestRectangleArea is an optimal solution where we do things on the fly
we have previous context to know pse
we have gained forward context when we pop back and have known nse

Another simplified approach is to calculate the pse & nse arrays in separate passes
And then do the final pass and get pse & nse for a given element to calculate the area

Area = (r - l + 1) - 2 {not to include the smaller bars}
Area = r - l - 1
*/
func largestRectangleArea(A []int) int {
	n := len(A)

	stack := make([]int, n)
	maxArea := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && A[stack[len(stack)-1]] > A[i] {
			element := A[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			nse := i
			pse := -1
			if len(stack) > 0 {
				pse = stack[len(stack)-1]
			}

			area := element * (nse - pse - 1)
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		element := A[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		nse := n
		pse := -1
		if len(stack) > 0 {
			pse = stack[len(stack)-1]
		}

		area := element * (nse - pse - 1)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
	// Case 1: [2, 1, 5, 6, 2, 3]
	A1 := []int{2, 1, 5, 6, 2, 3}
	fmt.Printf("Input: %v | Output: %v\n", A1, largestRectangleArea(A1))
	// Expected: 10
}
