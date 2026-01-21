package main

import "fmt"

func findFirstOccurrence(A []int, target int) int {
	l, r := 0, len(A)-1
	result := -1 // Default if not found

	for l <= r {
		mid := l + (r-l)/2

		if A[mid] == target {
			result = mid // Potential answer found
			r = mid - 1  // Keep searching left for earlier occurrence
		} else if A[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 4, 4, 4, 5, 6, 7}
	target := 4

	index := findFirstOccurrence(A, target)
	if index != -1 {
		fmt.Printf("First occurrence of %d is at index %d\n", target, index)
	} else {
		fmt.Println("Target not found")
	}
}
