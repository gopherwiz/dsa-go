/*
Problem Description
Implement the next permutation, which rearranges numbers into the numerically next greater permutation of numbers for a given array A of size N.
If such arrangement is not possible, it must be rearranged as the lowest possible order, i.e., sorted in ascending order.
NOTE:
The replacement must be in-place, do not allocate extra memory.
DO NOT USE LIBRARY FUNCTION FOR NEXT PERMUTATION. Use of Library functions will disqualify your submission retroactively and will give you penalty points.


Problem Constraints
1 <= N <= 5 * 105
1 <= A[i] <= 109


Input Format
The first and the only argument of input has an array of integers, A.


Output Format
Return an array of integers, representing the next permutation of the given array.


Example Input
Input 1:
 A = [1, 2, 3]
Input 2:
 A = [3, 2, 1]


Example Output
Output 1:
 [1, 3, 2]
Output 2:
 [1, 2, 3]


Example Explanation
Explanation 1:
 Next permutaion of [1, 2, 3] will be [1, 3, 2].
Explanation 2:
 No arrangement is possible such that the number are arranged into the numerically next greater permutation of numbers.
 So will rearranges it in the lowest possible order.
*/

package main

import "fmt"

func NextPermutation(A []int) []int {
	n := len(A)
	if n <= 1 {
		return A
	}

	// 1. Find the first decreasing element from the right {2 4 5 1 9 8 7}
	pivot := -1
	for i := n - 2; i >= 0; i-- {
		if A[i] < A[i+1] {
			pivot = i
			break
		}
	}

	// 2. If pivot exists, find the successor to swap with
	if pivot != -1 {
		successor := -1
		for i := n - 1; i > pivot; i-- {
			if A[i] > A[pivot] {
				successor = i
				break
			}
		}
		// Swap pivot and successor
		A[pivot], A[successor] = A[successor], A[pivot]
	}

	// 3. Reverse the portion to the right of the pivot
	// (If pivot was -1, this reverses the whole array)
	reverse(A, pivot+1, n-1)

	return A
}

func reverse(A []int, start, end int) {
	for start < end {
		A[start], A[end] = A[end], A[start]
		start++
		end--
	}
}

func main() {
	A1 := []int{1, 2, 3}
	fmt.Println("Next of [1, 2, 3]:", NextPermutation(A1)) // [1, 3, 2]

	A2 := []int{3, 2, 1}
	fmt.Println("Next of [3, 2, 1]:", NextPermutation(A2)) // [1, 2, 3]

	A3 := []int{1, 5, 8, 4, 7, 6, 5, 3, 1}
	fmt.Println("Next of complex array:", NextPermutation(A3))
}
