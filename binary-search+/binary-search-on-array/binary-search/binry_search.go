package main

import "fmt"

func search(A []int, target int) int {
	l, r := 0, len(A)-1

	for l <= r {
		mid := l + (r-l)/2

		if A[mid] == target {
			return mid // Match found, exit immediately
		}

		if A[mid] < target {
			l = mid + 1 // Target is in the right half
		} else {
			r = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not present
}

func main() {
	A := []int{1, 3, 5, 7, 9, 11}
	target := 7

	index := search(A, target)
	fmt.Printf("Index of %d: %d\n", target, index)
}
