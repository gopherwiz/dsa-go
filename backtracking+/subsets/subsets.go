/*
Problem Description
Given a set of distinct integers A, return all possible subsets.
NOTE:
Elements in a subset must be in non-descending order.
The solution set must not contain duplicate subsets.
Also, the subsets should be sorted in ascending ( lexicographic ) order.
The initial list is not necessarily sorted.


Problem Constraints
1 <= |A| <= 16
INTMIN <= A[i] <= INTMAX


Input Format
First and only argument of input contains a single integer array A.


Output Format
Return a vector of vectors denoting the answer.


Example Input
Input 1:
A = [1]
Input 2:
A = [1, 2, 3]


Example Output
Output 1:
[
    []
    [1]
]
Output 2:
[
 []
 [1]
 [1, 2]
 [1, 2, 3]
 [1, 3]
 [2]
 [2, 3]
 [3]
]


Example Explanation
Explanation 1:
 You can see that these are all possible subsets.
Explanation 2:
You can see that these are all possible subsets.
*/

package main

import (
	"fmt"
	"sort"
)

func Solve(A []int) [][]int {
	sort.Ints(A)

	var subsets [][]int
	backtrack(&subsets, []int{}, 0, A)

	sort.Slice(subsets, func(i, j int) bool {
		subset1 := subsets[i]
		subset2 := subsets[j]
		for k := 0; k < len(subset1) && k < len(subset2); k++ {
			if subset1[k] != subset2[k] {
				return subset1[k] < subset2[k]
			}
		}

		return len(subset1) < len(subset2)
	})

	return subsets
}

func backtrack(subsets *[][]int, subset []int, index int, A []int) {
	if index == len(A) {
		subsetCopy := append([]int(nil), subset...) // IMPORTANT: Must copy the slice because 'subset' is reused in BackTracking
		//subsetCopy := make([]int, len(subset))
		//copy(subsetCopy, subset)

		*subsets = append(*subsets, subsetCopy)

		return
	}

	// take A[index]
	subset = append(subset, A[index])
	backtrack(subsets, subset, index+1, A)
	subset = subset[:len(subset)-1]

	// don't take A[index]
	backtrack(subsets, subset, index+1, A)
}

//func backtrack(result *[][]int, current []int, index int, A []int) {
//	// Add a copy of the current state immediately
//	currentCopy := append([]int(nil), current...)
//	*result = append(*result, currentCopy)
//
//	for i := index; i < len(A); i++ {
//		current = append(current, A[i])    // "Pick"
//		backtrack(result, current, i+1, A) // Recurse
//		current = current[:len(current)-1] // "Don't Pick" (Backtrack)
//	}
//}

func main() {
	A := []int{1, 2, 3}
	result := Solve(A)
	fmt.Println(result)
}
