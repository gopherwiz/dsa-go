/*
Problem Description
You are given an array A of integers of size N.
Your task is to find the equilibrium index of the given array
The equilibrium index of an array is an index such that the sum of elements at lower indexes is equal to the sum of elements at higher indexes.
If there are no elements that are at lower indexes or at higher indexes, then the corresponding sum of elements is considered as 0.

Note:
Array indexing starts from 0.
If there is no equilibrium index then return -1.
If there are more than one equilibrium indexes then return the minimum index.

Example Input
Input 1:
A = [-7, 1, 5, 2, -4, 3, 0]
Input 2:
A = [1, 2, 3]


Example Output
Output 1:
3
Output 2:
-1
*/

package main

import "fmt"

/*
SolveWithPrefix
For getting a range [i,j] sum in a prefix array
if i == 0, sum = prefix[j]
else, sum = prefix[j] - prefix[i-1]
*/
func SolveWithPrefix(A []int) int {
	n := len(A)

	prefix := make([]int64, n)
	prefix[0] = int64(A[0])
	for i := 1; i < len(A); i++ {
		prefix[i] = prefix[i-1] + int64(A[i])
	}

	for i, _ := range A {
		var leftSum int64 // Sum of [0, i-1]

		if i == 0 {
			leftSum = 0
		} else {
			leftSum = prefix[i-1]
		}

		rightSum := prefix[n-1] - prefix[i] // Sum of [i+1, n-1]

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

func SolveWithCarryForward(A []int) int {
	var totalSum int64
	for _, num := range A {
		totalSum += int64(num)
	}

	var leftSum int64 = 0
	for i, num := range A {
		rightSum := totalSum - leftSum - int64(num)

		if leftSum == rightSum {
			return i
		}

		leftSum += int64(num)
	}

	return -1
}

func main() {
	A := []int{-7, 1, 5, 2, -4, 3, 0}

	fmt.Println(SolveWithPrefix(A))
	fmt.Println(SolveWithCarryForward(A))
}
