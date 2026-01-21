/*
Problem Description
Write a program to find the factorial of the given number A using recursion.
Note: The factorial of a number N is defined as the product of the numbers from 1 to N.


Example Input
Input 1:
 A = 4
Input 2:
 A = 1


Example Output
Output 1:
 24
Output 2:
 1


Example Explanation
Explanation 1:
 Factorial of 4 = 4 * 3 * 2 * 1 = 24
Explanation 2:
 Factorial of 1 = 1
*/

package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1)
}

func Solve(A int) int {
	return factorial(A)
}

func main() {
	A := 5

	fmt.Println(Solve(A))
}
