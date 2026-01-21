package main

import "fmt"

/**
 * @input A : 2D integer array
 * @Output Integer
 */
func solve(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	rows := len(A)
	cols := len(A[0])
	freshCount := 0

	type Orange struct {
		r, c, time int
	}
	var queue []Orange

	// 1. Initial scan to find all rotten oranges and count fresh ones
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if A[r][c] == 2 {
				queue = append(queue, Orange{r, c, 0})
			} else if A[r][c] == 1 {
				freshCount++
			}
		}
	}

	// If there are no fresh oranges, it takes 0 minutes
	if freshCount == 0 {
		return 0
	}

	maxTime := 0
	// Directions: Right, Down, Left, Up
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	// 2. Multi-source BFS
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Update maxTime with the time of the current orange being processed
		if curr.time > maxTime {
			maxTime = curr.time
		}

		for i := 0; i < 4; i++ {
			nr, nc := curr.r+dr[i], curr.c+dc[i]

			// If neighbor is within bounds and is a fresh orange
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && A[nr][nc] == 1 {
				// Rot the orange
				A[nr][nc] = 2
				freshCount--
				queue = append(queue, Orange{nr, nc, curr.time + 1})
			}
		}
	}

	// 3. If any fresh oranges remain, return -1
	if freshCount > 0 {
		return -1
	}

	return maxTime
}

func main() {
	// Example 1
	grid1 := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println("Example 1 Minutes:", solve(grid1)) // Output: 4

	// Example 2
	grid2 := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	fmt.Println("Example 2 Minutes:", solve(grid2)) // Output: -1
}
