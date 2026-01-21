/*
Problem Description
Given an array A. For every pair of indices i and j (i != j), find the maximum A[i] & A[j].

Example Input
Input 1:-
A = [53, 39, 88]
Input 2:-
A = [38, 44, 84, 12]

Example Output
Output 1:-
37
Output 2:-
36

Example Explanation
Explanation 1:-
53 & 39 = 37
39 & 88 = 0
53 & 88 = 16
Maximum among all these pairs is 37
Explanation 2:-
Maximum bitwise and among all pairs is (38, 44) = 36
*/

package main

import "fmt"

func Solve(A []int) int {
	result := 0
	for i := 31; i >= 0; i-- {
		setBitCount := 0
		for _, num := range A {
			if num&(1<<i) != 0 {
				setBitCount++
			}
		}

		if setBitCount >= 2 {
			result = result | (1 << i)

			for index, num := range A {
				if num&(1<<i) == 0 {
					A[index] = 0
				}
			}
		}
	}

	return result
}

func main() {
	A := []int{38, 44, 84, 12}

	fmt.Println(Solve(A))
}
