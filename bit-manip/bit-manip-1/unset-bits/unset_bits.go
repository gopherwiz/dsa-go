/*
Problem Description
Given an integer A. Unset B bits from the right of A in binary.

For example, if A = 93 and B = 4, the binary representation of A is 1011101.
If we unset the rightmost 4 bits, we get the binary number 1010000, which is equal to the decimal value 80.


Example Input
Input 1:-
A = 25
B = 3
Input 2:-
A = 37
B = 3


Example Output
Output 1:-
24
Output 2:-
32


Example Explanation
Explanation 1:-
A = 11001 to 11000
Explantio 2:-
A = 100101 to 100000
*/

package main

import "fmt"

func Solve(A int64, B int) int64 {
	for i := 0; i < B; i++ {
		if A&(1<<i) != 0 {
			A = A ^ (1 << i)
		}
	}

	return A
}

func main() {
	A := int64(25) // 11001
	B := 3

	fmt.Println(Solve(A, B))
}
