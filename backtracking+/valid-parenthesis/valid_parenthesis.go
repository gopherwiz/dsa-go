/*
Problem Description
Given an integer A pairs of parentheses, write a function to generate all combinations of well-formed parentheses of length 2*A.



Problem Constraints
1 <= A <= 10



Input Format
First and only argument is integer A.



Output Format
Return a sorted list of all possible parenthesis.



Example Input
Input 1:

A = 3
Input 2:
A = 1


Example Output
Output 1:
[ "((()))", "(()())", "(())()", "()(())", "()()()" ]

Output 2:
[ "()" ]


Example Explanation
Explanation 1:
 All paranthesis are given in the output list.

Explanation 2:
 All paranthesis are given in the output list.
*/

/*
In professional Go development, strings.Builder is used when you are moving in one direction (adding only).
Because backtracking requires frequent "popping" from the end,
a []byte slice (like in your previous code) is actually the superior choice.

Feature                []byte                  Backtracking,strings.Builder
Adding                "append(b, '(')"         sb.WriteByte('(')
Removing               b[:len(b)-1] (O(1))     No built-in method (Hard)
Final Result           string(b) (Allocates)   sb.String() (Allocates)
Recommendation         Best for Backtracking   Best for simple concatenation
*/
package main

import "fmt"

func Solve(A int) []string {
	var result []string
	backtrack(&result, []byte{}, 0, 0, A)
	return result
}

// True Backtracking (Memory Efficient)
func backtrack(res *[]string, current []byte, open, close, max int) {
	if len(current) == max*2 {
		*res = append(*res, string(current))
		return
	}

	if open < max {
		current = append(current, '(')
		backtrack(res, current, open+1, close, max)
		current = current[:len(current)-1] // Undo (The "Back" in Backtrack)
	}
	if close < open {
		current = append(current, ')')
		backtrack(res, current, open, close+1, max)
		current = current[:len(current)-1] // Undo
	}
}

func main() {
	A := 3
	fmt.Printf("Valid parentheses for N=%d:\n%v\n", A, Solve(A))
}
