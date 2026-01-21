package main

import "fmt"

/**
 * @input A : Integer (Number of nodes)
 * @input B : 2D integer array (Edges)
 * @input C : Source Node
 * @input D : Destination Node
 * * @Output Integer (Shortest distance or -1 if unreachable)
 */
func solve(A int, B [][]int, C int, D int) int {
	// 1. Build Adjacency List
	// Using A+1 because nodes are typically 1-indexed
	graph := make([][]int, A+1)
	for _, edge := range B {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // Undirected graph
	}

	// 2. BFS Setup
	type Item struct {
		node int
		dist int
	}

	// visited array to keep track of explored nodes
	visited := make([]bool, A+1)

	// Create a queue and push the starting node C
	queue := []Item{{node: C, dist: 0}}
	visited[C] = true

	// 3. BFS Traversal
	for len(queue) > 0 {
		// Pop the front element
		curr := queue[0]
		queue = queue[1:]

		// If we reached the target, return the distance
		if curr.node == D {
			return curr.dist
		}

		// Check all neighbors
		for _, neighbor := range graph[curr.node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, Item{node: neighbor, dist: curr.dist + 1})
			}
		}
	}

	// If the queue is empty and we never hit D
	return -1
}

func main() {
	// Example Graph: 1 - 2, 2 - 3, 3 - 4
	A := 4
	B := [][]int{{1, 2}, {2, 3}, {3, 4}}
	C, D := 1, 4

	fmt.Printf("Shortest distance from %d to %d is: %d\n", C, D, solve(A, B, C, D))
	// Output: 3
}
