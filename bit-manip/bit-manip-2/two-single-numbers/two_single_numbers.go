/*
Problem Description
Given an array of positive integers A, two integers appear only once, and all the other integers appear twice.
Find the two integers that appear only once.
Note: Return the two numbers in ascending order.


Example Input
Input 1:
A = [1, 2, 3, 1, 2, 4]
Input 2:
A = [1, 2]


Example Output
Output 1:
[3, 4]
Output 2:
[1, 2]


Example Explanation
Explanation 1:
3 and 4 appear only once.
Explanation 2:
1 and 2 appear only once.
*/

package main

import (
	"fmt"
	"sort"
)

func Solve(A []int) []int {
	xorAll := 0
	for i := 0; i < len(A); i++ {
		xorAll ^= A[i]
	}

	// Find any set bit index in xorAll - set bit implies that at that position, the two single numbers have opposite bits
	setBitIndex := 0
	for xorAll > 0 {
		if xorAll&1 == 1 {
			break
		}

		xorAll = xorAll >> 1
		setBitIndex++
	}

	// Segregate numbers into 2 groups based on set bit & xor the groups to turn duplicate numbers in a group to 0
	num1 := 0
	num2 := 0
	for _, num := range A {
		if num&(1<<setBitIndex) != 0 {
			num1 = num1 ^ num
		} else {
			num2 = num2 ^ num
		}
	}

	result := []int{num1, num2}
	sort.Ints(result)

	return result
}

func main() {
	A := []int{5, 5, 4, 4, 1, 2}

	fmt.Println(Solve(A))
}
