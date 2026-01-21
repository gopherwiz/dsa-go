package main

import "fmt"

/**
 * solve detects a cycle using BFS (Kahn's Algorithm).
 * @input A : Integer (Number of nodes)
 * @input B : 2D integer array (Directed edges)
 * @Output Integer (1 if cycle present, else 0)
 */
func solve(A int, B [][]int) int {
	// 1. Build Adjacency List and Calculate In-Degrees
	graph := make([][]int, A+1)
	inDegree := make([]int, A+1)
	for _, edge := range B {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		inDegree[v]++
	}

	// 2. Initialize Queue with all nodes having In-Degree 0
	var queue []int
	for i := 1; i <= A; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. Process nodes and count how many we can reach
	nodesProcessed := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		nodesProcessed++

		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			// If dependency count hits zero, it can now be processed
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 4. Cycle Detection Logic:
	// If we couldn't process all nodes, there must be a cycle.
	if nodesProcessed == A {
		return 0 // No cycle
	}
	return 1 // Cycle present
}

func main() {
	// Case 1: Cycle present (1 -> 2 -> 3 -> 1)
	A1 := 3
	B1 := [][]int{{1, 2}, {2, 3}, {3, 1}}
	fmt.Printf("Test Case 1 (Cycle): %d\n", solve(A1, B1)) // Expected: 1

	// Case 2: No Cycle (1 -> 2 -> 3)
	A2 := 3
	B2 := [][]int{{1, 2}, {2, 3}}
	fmt.Printf("Test Case 2 (No Cycle): %d\n", solve(A2, B2)) // Expected: 0
}
