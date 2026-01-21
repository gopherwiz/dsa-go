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

/* GPT
That's a very insightful question — you're right to think about the possibility space being exponential, like 2^N, but let's clarify why the time complexity is not exactly O(2^N), and why O(C(n) * n) is a more accurate bound in this case.

👇 Let's break it down:
✅ At each step, you can choose:

Add a '(' if you haven't used up all n
Add a ')' if you haven’t exceeded the number of '(' added
That feels like two options — hence, 2^k possibilities for k = 2n steps, right?

But not quite — here's the catch:

❌ Not all 2^2n sequences are valid
For example, for n = 2:

Total 4-character combinations from '(' and ')' is 2^4 = 16
But only 2 of those are valid parentheses combinations: "(())" and "()()"
That's a tiny fraction — and this is exactly where Catalan numbers come in.

📈 Total valid combinations: Catalan number C(n)
Instead of brute-forcing all 2^(2n) paths, we prune the invalid ones as we go using backtracking:

We never go down a path where ) would exceed ( — those are instantly skipped
So even though the theoretical tree has 2^(2n) nodes, we never explore most of it
Hence, the total number of recursive calls is actually bounded by O(C(n)), and building each sequence of length 2n costs O(n).

✅ So total time = O(C(n) * n), not O(2^(2n))
🤔 Why not call it exponential?
We still can — Catalan numbers grow exponentially, roughly like:

C(n) ~ 4^n / (n^(3/2) * sqrt(pi))
So while it’s not 2^(2n), it’s still exponential — just much less than brute-force.
*/

package main

import "fmt"

func Solve(n int) []string {
	var result []string
	parenthesis(&result, "", 0, 0, n)
	return result
}

func parenthesis(result *[]string, current string, open int, close int, max int) {
	if open == max && close == max {
		*result = append(*result, current)
		return
	}

	if open > max || close > open {
		return
	}

	parenthesis(result, current+"(", open+1, close, max)
	parenthesis(result, current+")", open, close+1, max)
}

func main() {
	A := 3
	result := Solve(A)
	fmt.Println(result)
}
