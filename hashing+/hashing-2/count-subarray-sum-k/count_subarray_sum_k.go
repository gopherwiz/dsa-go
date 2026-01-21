/*
Problem Description
Given an array of integers A and an integer B.
Find the total number of subarrays having sum equals to B.


Problem Constraints
 1 <= length of the array <= 50000
-1000 <= A[i] <= 1000


Input Format
The first argument given is the integer array A.
The second argument given is integer B.


Output Format
Return the total number of subarrays having sum equals to B.


Example Input
Input 1:
A = [1, 0, 1]
B = 1
Input 2:
A = [0, 0, 0]
B = 0


Example Output
Output 1:
4
Output 2:
6


Example Explanation
Explanation 1:
[1], [1, 0], [0, 1] and [1] are four subarrays having sum 1.
Explanation 1:
All the possible subarrays having sum 0.
*/

/*
prefix[i] == k
prefix[j] - prefix[i-1] == k
*/

package main

import "fmt"

/*
SUBARRAY SUM (O(N) Logic): "Looking Backwards"
==============================================
Equation: P[j] - P[i] = B  =>  P[i] = P[j] - B

Logic:
1. I am at the END of a potential subarray (P[j]).
2. I ask: "Has there been a STARTING point (P[i]) in the past that makes this work?"
3. P[i] must be exactly (currentSum - B).

Single-Pass Execution:
- As you loop, check: `count += hashmap[currentSum - B]`
- Then, save the current sum: `hashmap[currentSum]++`

Why this works:
By checking the map BEFORE adding the current sum, you are searching
only the "past" (valid starting points). Initializing the map with
{0: 1} handles subarrays that start from index 0.
*/

func SolveWithPrefix(A []int, B int) int {
	n := len(A)

	// NOTE* panics if A is empty @ prefix[0] = A[0]
	//prefix := make([]int, n)
	//prefix[0] = A[0]
	//for i := 1; i < n; i++ {
	//	prefix[i] = prefix[i-1] + A[i]
	//}

	prefix := make([]int, n+1)
	prefix[0] = 0
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + A[i]
	}

	count := 0
	hashmap := make(map[int]int)
	hashmap[0] = 1
	for i := 0; i < n; i++ {
		pj := prefix[i]
		
		if _, exists := hashmap[pj-B]; exists {
			count += hashmap[pj-B]
		}

		hashmap[pj]++
	}

	return count
}

func Solve(A []int, B int) int {
	hashmap := make(map[int]int)
	hashmap[0] = 1

	total := 0
	count := 0
	for _, num := range A {
		total += num

		// NOTE* hashmap handles this @ hashmap[0] = 1
		//if total == B {
		//	count++
		//}

		if freq, exists := hashmap[total-B]; exists {
			count += freq
		}

		hashmap[total]++
	}

	return count
}

func main() {
	A := []int{1, 0, 1}
	B := 1
	fmt.Println(SolveWithPrefix(A, B))
}
