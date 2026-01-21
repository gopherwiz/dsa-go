/*
Problem Description
Given a number A, we need to find the sum of its digits using recursion.


Example Input
Input 1:
 A = 46
Input 2:
 A = 11


Example Output
Output 1:
 10
Output 2:
 2


Example Explanation
Explanation 1:
 Sum of digits of 46 = 4 + 6 = 10
Explanation 2:
 Sum of digits of 11 = 1 + 1 = 2
*/

package main

import "fmt"

func sum(num int) int {
	if num <= 9 {
		return num
	}

	return sum(num/10) + num%10
}

func Solve(A int) int {
	return sum(A)
}

func main() {
	A := 123456789

	fmt.Println(Solve(A))
}
