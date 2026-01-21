package main

import "fmt"

func search(A []int, B int) int {
	l, r := 0, len(A)-1

	for l <= r {
		mid := l + (r-l)/2

		if A[mid] == B {
			return mid
		}

		// Identify which half is sorted
		if A[l] <= A[mid] {
			// Left side is sorted (like [4, 5, 6, 7])
			if B >= A[l] && B < A[mid] {
				r = mid - 1 // Target is in the sorted left part
			} else {
				l = mid + 1 // Target must be in the right part
			}
		} else {
			// Right side is sorted (like [0, 1, 2])
			if B > A[mid] && B <= A[r] {
				l = mid + 1 // Target is in the sorted right part
			} else {
				r = mid - 1 // Target must be in the left part
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 4)) // Output: 0
	fmt.Println(search([]int{9, 10, 3, 5, 6, 8}, 5))   // Output: 3
}
