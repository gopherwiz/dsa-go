/*
Problem Description
Write a function that takes an integer and returns the number of 1 bits present in its binary representation.


Example Input
Input 1:
11
Input 2:
6


Example Output
Output 1:
3
Output 2:
2


Example Explanation
Explaination 1:
11 is represented as 1011 in binary.
Explaination 2:
6 is represented as 110 in binary.
*/

package main

import "fmt"

func Solve(A int) any {
	ct := 0

	for A > 0 {
		if A&1 == 1 {
			ct++
		}
		A = A >> 1
	}

	return ct
}

func main() {
	A := 256

	fmt.Println(getBinary(A))
	fmt.Println(Solve(A))
}

func getBinary(A int) any {
	var binary []int
	for A > 0 {
		binary = append([]int{A & 1}, binary...)
		A = A >> 1
	}

	return binary
}
