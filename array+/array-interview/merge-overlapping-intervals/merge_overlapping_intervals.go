/*
Problem Description
Given a collection of intervals, merge all overlapping intervals.

Example Input
Input 1:
[1,3],[2,6],[8,10],[15,18]

Example Output
Output 1:
[1,6],[8,10],[15,18]

Example Explanation
Explanation 1:
Merge intervals [1,3] and [2,6] -> [1,6].
so, the required answer after merging is [1,6],[8,10],[15,18].
No more overlapping intervals present.
*/

package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start, end int
}

func Solve(A []Interval) []Interval {
	n := len(A)

	sort.Slice(A,
		func(i, j int) bool {
			return A[i].start < A[j].start
		})

	start := A[0].start
	end := A[0].end
	var result []Interval
	for i := 1; i < n; i++ {
		if A[i].start > end { // Non overlapping condition
			result = append(result, Interval{start: start, end: end})
			start = A[i].start
			end = A[i].end
		} else { // Overlapping condition
			start = getMin(start, A[i].start) // Optional as already sorted A as per starts
			end = getMax(end, A[i].end)
		}
	}

	// Need it for the last one as append is done after setting start & end
	result = append(result, Interval{start: start, end: end})

	return result
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	intervals := []Interval{
		{start: 1, end: 3},
		{start: 2, end: 6},
		{start: 8, end: 10},
		{start: 15, end: 18},
	}

	fmt.Println(Solve(intervals))
}
