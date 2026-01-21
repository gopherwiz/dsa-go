package main

// SolveS is the entry point for the sorting logic
func SolveS(A []int) []int {
	return mergeSortS(A)
}

// mergeSortS is the main function that handles the recursion
func mergeSortS(A []int) []int {
	// Base case: arrays with 0 or 1 element are already sorted
	if len(A) <= 1 {
		return A
	}

	// 1. Divide: Find the midpoint
	mid := len(A) / 2

	// 2. Conquer: Recursively sort left and right halves
	// Note: A[:mid] is indices 0 to mid-1, A[mid:] is mid to end
	left := mergeSortS(A[:mid])
	right := mergeSortS(A[mid:])

	// 3. Combine: Merge the sorted halves
	return mergeS(left, right)
}

// merge is the helper function that combines two sorted slices
func mergeS(left, right []int) []int {
	size := len(left) + len(right)
	result := make([]int, size)

	i, j, k := 0, 0, 0

	// Compare elements from both halves and pick the smaller one
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// If there are remaining elements in the left slice, add them
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	// If there are remaining elements in the right slice, add them
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

//func main() {
//	A := []int{38, 27, 43, 3, 9, 82, 10}
//	fmt.Println("Original:", A)
//
//	sortedA := SolveS(A)
//	fmt.Println("Sorted:  ", sortedA)
//}
