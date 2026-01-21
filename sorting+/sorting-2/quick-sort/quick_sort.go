package main

import "fmt"

func Solve(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	quickSort(A, 0, len(A)-1)

	return A
}

func quickSort(A []int, l, r int) {
	if l >= r {
		return
	}

	pi := partition(A, l, r)

	quickSort(A, l, pi-1)
	quickSort(A, pi+1, r)
}

func partition(A []int, l, r int) int {
	pivotIndex := l // first element as pivot
	pivot := A[l]

	for l <= r {
		if A[l] <= pivot {
			l++
		} else if A[r] > pivot {
			r--
		} else {
			A[l], A[r] = A[r], A[l]
		}
	}

	A[pivotIndex], A[r] = A[r], A[pivotIndex]

	return r
}

func main() {
	A := []int{5, 4, 3, 2, 1, 7, 7}
	fmt.Println(Solve(A))
}
