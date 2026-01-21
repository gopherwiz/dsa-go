/*
Key Takeaway: You use indices to divide the work (Conquer),
but you still need a temporary "waiting room" (Buffer) to combine the work
because you can't sort the elements in their original slots without stepping on each other's toes.
*/

package main

import "fmt"

// Solve is the entry point for the sorting logic
func Solve(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	// Create one auxiliary buffer to be used for all merge operations.
	// This keeps space complexity at O(N) instead of O(N log N).
	temp := make([]int, len(A))

	mergeSort(A, temp, 0, len(A)-1)
	return A
}

// mergeSort handles the recursive division of the array
func mergeSort(A, temp []int, left, right int) {
	if left >= right {
		return
	}

	// Calculate mid safely to avoid potential integer overflow
	mid := left + (right-left)/2

	// Conquer: Sort left and right halves
	mergeSort(A, temp, left, mid)
	mergeSort(A, temp, mid+1, right)

	// Combine: Merge the two sorted halves
	merge(A, temp, left, mid, right)
}

// merge combines two sorted portions of A[left...mid] and A[mid+1...right]
func merge(A, temp []int, left, mid, right int) {
	// Copy current segment into temp buffer
	for i := left; i <= right; i++ {
		temp[i] = A[i]
	}

	i := left    // Initial index of first sub-array
	j := mid + 1 // Initial index of second sub-array
	k := left    // Initial index of merged array in A

	// Standard merge logic using the temp buffer to pick the smallest element
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			A[k] = temp[i]
			i++
		} else {
			A[k] = temp[j]
			j++
		}
		k++
	}

	// Copy any remaining elements from the left side
	for i <= mid {
		A[k] = temp[i]
		i++
		k++
	}
	// Note: Right side elements are already in their place in A
	// if we don't finish them, so no need for a second 'remaining' loop.
}

func main() {
	A := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array:", A)

	Solve(A)
	fmt.Println("Sorted array:  ", A)
}
