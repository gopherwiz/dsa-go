/*
/*
Problem Description
Implement pow(A, B) % C.
In other words, given A, B and C, Find (AB % C).
Note: The remainders on division cannot be negative. In other words, make sure the answer you return is non-negative.


Problem Constraints
-109 <= A <= 109
0 <= B <= 109
1 <= C <= 109


Input Format
Given three integers A, B, C.


Output Format
Return an integer.


Example Input
Input 1:
A = 2
B = 3
C = 3

Input 2:
A = 3
B = 3
C = 1


Example Output
Output 1:
2

Output 2:
0


Example Explanation
Explanation 1:
2^3 % 3 = 8 % 3 = 2

Explanation 2:
3^3 % 1 = 27 % 1 = 0
*/

package main

import "fmt"

func fastPow(base int, power int, mod int) int {
	if base == 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	half := fastPow(base, power/2, mod)

	result := (half * half) % mod

	if power%2 == 1 {
		result = (base * result) % mod
	}

	return result
}

func Solve(A int, B int, C int) int {
	result := fastPow(A, B, C)

	if result < 0 {
		return result + C
	}

	return result
}

func main() {
	A := 2
	B := 3
	C := 3

	fmt.Println(Solve(A, B, C))
}
