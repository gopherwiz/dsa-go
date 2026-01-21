/*
Problem Description
You are given an array of N integers, A1, A2 ,..., AN and an integer B. Return the of count of distinct numbers in all windows of size B.
Formally, return an array of size N-B+1 where i'th element in this array contains number of distinct elements in sequence Ai, Ai+1 ,..., Ai+B-1.
NOTE: if B > N, return an empty array.


Problem Constraints
1 <= N <= 106

1 <= A[i] <= 109


Input Format
First argument is an integer array A
Second argument is an integer B.


Output Format
Return an integer array.


Example Input
Input 1:
 A = [1, 2, 1, 3, 4, 3]
 B = 3
Input 2:
 A = [1, 1, 2, 2]
 B = 1


Example Output
Output 1:
 [2, 3, 3, 2]
Output 2:
 [1, 1, 1, 1]


Example Explanation
Explanation 1:
 A=[1, 2, 1, 3, 4, 3] and B = 3
 All windows of size B are
 [1, 2, 1]
 [2, 1, 3]
 [1, 3, 4]
 [3, 4, 3]
 So, we return an array [2, 3, 3, 2].
Explanation 2:
 Window size is 1, so the output array is [1, 1, 1, 1].
*/

package main

import "fmt"

func Solve(A []int, B int) []int {
	n := len(A)
	hashmap := make(map[int]int)

	start := 0
	end := B - 1
	for i := start; i <= end; i++ {
		hashmap[A[i]]++
	}

	result := []int{len(hashmap)}
	start++
	end++

	for end < n {
		hashmap[A[end]]++
		hashmap[A[start-1]]--
		if hashmap[A[start-1]] == 0 {
			delete(hashmap, A[start-1])
		}

		result = append(result, len(hashmap))

		start++
		end++
	}

	return result
}

func main() {
	A := []int{1, 2, 1, 3, 4, 3}
	B := 3
	fmt.Println(Solve(A, B))
}
