/*
Problem Description
You have a set of non-overlapping intervals. You are given a new interval [start, end], insert this new interval into the set of intervals (merge if necessary).
You may assume that the intervals were initially sorted according to their start times.

Example Input
Input 1:
Given intervals [1, 3], [6, 9] insert and merge [2, 5] .
Input 2:
Given intervals [1, 3], [6, 9] insert and merge [2, 6] .

Example Output
Output 1:

	[ [1, 5], [6, 9] ]

Output 2:

	[ [1, 9] ]

Example Explanation
Explanation 1:
(2,5) does not completely merge the given intervals
Explanation 2:
(2,6) completely merges the given intervals
*/

package main

import (
	"fmt"
)

type Interval struct {
	start, end int
}

func Solve(intervals []Interval, newInterval Interval) []Interval {
	n := len(intervals)

	L := newInterval.start
	R := newInterval.end
	var result []Interval
	for i := 0; i < n; i++ {
		start := intervals[i].start
		end := intervals[i].end
		if end < L { // Non overlapping interval towards the left
			result = append(result, Interval{start: start, end: end})
			// Exit here is an edge case scenario, where [L, R] does not overlap and is simply pushed at the end
		} else if start > R { // Non overlapping interval towards the right, the new interval is already merged in the left
			result = append(result, Interval{start: L, end: R})
			for j := i; j < n; j++ {
				result = append(result, intervals[j])
			}
			return result
			// Exit here is the most common scenario, where [L, R] overlaps and is merged somewhere in the middle
		} else { // Overlapping interval - not appending [L, R] prematurely as there might be more overlaps
			L = getMin(L, start)
			R = getMax(R, end)
			// Exit here is an edge case scenario, where [L, R] overlaps and is merged till the very end
		}
	}

	result = append(result, Interval{start: L, end: R}) // Needed when exiting from 1st and 3rd condition

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
		{start: 6, end: 9},
	}
	interval := Interval{start: 2, end: 5}

	fmt.Println(Solve(intervals, interval))
}
