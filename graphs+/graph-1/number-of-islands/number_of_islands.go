package main

import "fmt"

/**
 * solve counts the number of connected groups of 1's (islands).
 * Two 1's are connected if they are adjacent horizontally, vertically, or diagonally.
 */
func solve(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	rows := len(A)
	cols := len(A[0])
	islandCount := 0

	// We'll define DFS as a nested function
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Base Case: check boundaries and if the cell is water (0)
		if r < 0 || r >= rows || c < 0 || c >= cols || A[r][c] == 0 {
			return
		}

		// Mark current cell as visited by "sinking" it (setting it to 0)
		A[r][c] = 0

		// Explore all 8 neighbors (including diagonals)
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				dfs(r+i, c+j)
			}
		}
	}

	// Main loop to scan the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if A[r][c] == 1 {
				// We found a piece of land, start an island discovery
				islandCount++
				dfs(r, c)
			}
		}
	}

	return islandCount
}

func main() {
	// Example 1: Simple 3x3 grid
	// Output should be 2
	grid1 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	}
	fmt.Println("Islands in Example 1:", solve(grid1))

	// Example 2: Larger 5x5 grid
	// Output should be 5
	grid2 := [][]int{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
	}
	fmt.Println("Islands in Example 2:", solve(grid2))
}
