/*
Problem Description
Given 2 integers A and B and an array of integers C of size N. Element C[i] represents the length of ith board.
You have to paint all N boards [C0, C1, C2, C3 … CN-1]. There are A painters available and each of them takes B units of time to paint 1 unit of the board.

Calculate and return the minimum time required to paint all boards under the constraints that any painter will only paint contiguous sections of the board.
NOTE:
1. 2 painters cannot share a board to paint. That is to say, a board cannot be painted partially by one painter, and partially by another.
2. A painter will only paint contiguous boards. This means a configuration where painter 1 paints boards 1 and 3 but not 2 is invalid.

Return the ans % 10000003.



Problem Constraints
1 <= A <= 1000
1 <= B <= 106
1 <= N <= 105
1 <= C[i] <= 106



Input Format
The first argument given is the integer A.
The second argument given is the integer B.
The third argument given is the integer array C.



Output Format
Return minimum time required to paint all boards under the constraints that any painter will only paint contiguous sections of board % 10000003.



Example Input
Input 1:
 A = 2
 B = 5
 C = [1, 10]

Input 2:
 A = 10
 B = 1
 C = [1, 8, 11, 3]


Example Output
Output 1:
 50

Output 2:
 11


Example Explanation
Explanation 1:
 Possibility 1:- One painter paints both blocks, time taken = 55 units.
 Possibility 2:- Painter 1 paints block 1, painter 2 paints block 2, time take = max(5, 50) = 50
 There are no other distinct ways to paint boards.
 ans = 50 % 10000003

Explanation 2:
 Each block is painted by a painter so, Painter 1 paints block 1, painter 2 paints block 2, painter 3 paints block 3
 and painter 4 paints block 4, time taken = max(1, 8, 11, 3) = 11
 ans = 11 % 10000003
*/

package main

import (
	"fmt"
)

// paint has A as number_of_painters, B as time_per_board_unit, C as boards_size_array
func paint(A int, B int, C []int) int {
	MOD := 10000003

	maxBoards := 0
	sumBoards := int64(0)
	for _, length := range C {
		if length > maxBoards {
			maxBoards = length
		}
		sumBoards += int64(length)
	}

	// left: minimum possible time - max board a painter might have to paint
	// right: maximum possible time - sum of all boards (all boards by 1 painter)
	left := int64(maxBoards)
	right := sumBoards
	ans := sumBoards // assume ans as worst and improve

	for left <= right {
		mid := left + (right-left)/2

		// Logic: If painters needed for 'mid' is <= available A, then 'mid' is possible
		if paintersNeeded(C, mid) <= A {
			ans = mid
			right = mid - 1 // reduce timeLimit to use more painters or get more painting output per painter
		} else {
			left = mid + 1 // increase timeLimit to use fewer painters
		}
	}

	// Apply multiplier B and MOD
	// We use int64 for calculation to handle B up to 10^6 and ans up to 10^11
	finalAns := (ans % int64(MOD)) * (int64(B) % int64(MOD))
	return int(finalAns % int64(MOD))
}

func paintersNeeded(C []int, timeLimit int64) int {
	// We always need at least 1 painter if there are boards
	painters := 1
	currentTime := int64(0)

	for _, boardTime := range C {
		// If adding this boardTime exceeds the limit, assign to the next painter
		if currentTime+int64(boardTime) > timeLimit {
			painters++
			currentTime = int64(boardTime)
		} else {
			currentTime += int64(boardTime)
		}
	}

	return painters
}

func main() {
	fmt.Println(paint(2, 5, []int{1, 10}))               // Output: 50
	fmt.Println(paint(10, 1, []int{1, 8, 11, 3}))        // Output: 11
	fmt.Println(paint(4, 10, []int{884, 228, 442, 889})) // Output: 8890
}
