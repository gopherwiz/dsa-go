/*
Problem Description
You are given two integers A and B.
Return 1 if B-th bit in A is set
Return 0 if B-th bit in A is unset
Note:
The bit position is 0-indexed, which means that the least significant bit (LSB) has index 0.


Example Input
Input 1:
A = 4
B = 1
Input 2:
A = 5
B = 2


Example Output
Output 1:
0
Output 2:
1


Example Explanation
For Input 1:
Given N = 4 which is 100 in binary. The 1-st bit is unset
so we return 0
For Input 2:
Given N = 5 which is 101 in binary. The 2-nd bit is set
so we return 1
*/

package main

import "fmt"

func Solve(A int, B int) int {
	if isSet(A, B) {
		return 1
	}

	return 0
}

func main() {
	A := 256
	B := 8

	fmt.Println(Solve(A, B))
}

func isSet(n int, i int) bool {
	return n&(1<<i) != 0
}

func set(n int, i int) int {
	return n | (1 << i)
}

func unset(n int, i int) int {
	if isSet(n, i) {
		return toggle(n, i)
	}

	return n
}

func toggle(n int, i int) int {
	return n ^ (1 << i)
}
