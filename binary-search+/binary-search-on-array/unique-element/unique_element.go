package main

import "fmt"

func solve(A []int) int {
	n := len(A)
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2

		// 1. Check if mid is the unique element
		// Instead of INT_MAX, we explicitly check if we are at the edges
		isLeftDifferent := (mid == 0) || (A[mid] != A[mid-1])
		isRightDifferent := (mid == n-1) || (A[mid] != A[mid+1])

		if isLeftDifferent && isRightDifferent {
			return A[mid]
		}

		// 2. Decide which half to search
		// Find the first index of the pair mid belongs to
		firstIdx := mid
		if mid > 0 && A[mid] == A[mid-1] {
			firstIdx = mid - 1
		}

		// Property: Pairs before the single element always start at an EVEN index
		if firstIdx%2 == 0 {
			// Rule holds, single element is on the right
			l = firstIdx + 2
		} else {
			// Rule broken, single element is on the left
			r = firstIdx - 1
		}
	}
	return -1
}

func main() {
	A := []int{1, 1, 2, 2, 7, 9, 9}
	fmt.Println("Single Element:", solve(A))
}
