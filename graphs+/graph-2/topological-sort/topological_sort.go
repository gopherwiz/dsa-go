package main

import "fmt"

/**
 * @input A : Integer (Number of nodes)
 * @input B : 2D integer array (Edges)
 * @Output 1D integer array
 */
func solve(A int, B [][]int) []int {
	// 1. Initialize Graph and In-degrees
	graph := make([][]int, A+1)
	inDegree := make([]int, A+1)
	for _, edge := range B {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		inDegree[v]++
	}

	// 2. Initialize Queue with all nodes having in-degree 0
	var queue []int
	for i := 1; i <= A; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. Process nodes
	result := make([]int, 0, A)
	for len(queue) > 0 {
		// Pop front of queue
		curr := queue[0]
		queue = queue[1:]

		result = append(result, curr)

		// Reduce in-degree of neighbors
		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			// If in-degree becomes 0, add to queue
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 4. Cycle Check
	if len(result) == A {
		return result
	}

	return []int{} // Return empty if there's a cycle
}

func main() {
	A := 6
	B := [][]int{{6, 3}, {6, 1}, {5, 1}, {5, 2}, {3, 4}, {4, 2}}
	fmt.Println("Topo Order:", solve(A, B))
}
