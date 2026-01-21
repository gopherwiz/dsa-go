/*
Problem Description
Given an array of integers, every element appears thrice except for one, which occurs once.

Find that element that does not appear thrice.

NOTE: Your algorithm should have a linear runtime complexity.

Could you implement it without using extra memory?


Example Input
Input 1:
 A = [1, 2, 4, 3, 3, 2, 2, 3, 1, 1]
Input 2:
 A = [0, 0, 0, 1]


Example Output
Output 1:
 4
Output 2:
 1


Example Explanation
Explanation 1:
 4 occurs exactly once in Input 1.
 1 occurs exactly once in Input 2.
*/

package main

import "fmt"

func Solve(A []int) int {
	k := 3 // If k is even, we can directly xor all array elements and get the number occurring once

	result := 0
	for i := 0; i < 32; i++ {
		setBitCount := 0
		for _, num := range A {
			if num&(1<<i) != 0 {
				setBitCount++
			}
		}

		if setBitCount%k != 0 {
			result = result | (1 << i)
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}

	fmt.Println(Solve(A))
}
