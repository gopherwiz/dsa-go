package main

import "fmt"

/**
 * @input A : Integer (Number of nodes)
 * @input B : 2D integer array (Directed edges)
 * * @Output Integer (1 if cycle present, else 0)
 */
func solve(A int, B [][]int) int {
	// 1. Build Adjacency List
	// Nodes are 1-indexed, so we use A + 1
	graph := make([][]int, A+1)
	for _, edge := range B {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
	}

	visited := make([]bool, A+1)
	path := make([]bool, A+1)

	// 2. Define DFS as a nested closure to capture local variables
	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = true
		path[node] = true

		for _, neighbor := range graph[node] {
			// If neighbor is in current recursion path, cycle found
			if path[neighbor] {
				return true
			}
			// If neighbor not visited, recurse
			if !visited[neighbor] {
				if dfs(neighbor) { // must not return dfs(neighbor) as it won't iterate over all neighbors
					return true
				}
			}
		}

		// Backtrack: Remove node from current recursion stack
		path[node] = false
		return false
	}

	// 3. Iterate through all nodes to handle disconnected components
	for i := 1; i <= A; i++ {
		if !visited[i] {
			if dfs(i) {
				return 1
			}
		}
	}

	return 0
}

func main() {
	// Test Case 1: Cycle present
	fmt.Println(solve(5, [][]int{{1, 2}, {4, 1}, {2, 4}, {3, 4}, {5, 2}, {1, 3}})) // Output: 1

	// Test Case 2: No cycle
	fmt.Println(solve(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})) // Output: 0
}
