package main

import "fmt"

func assignEdgeWeights(edges [][]int) int {
	const MOD int64 = 1_000_000_007

	// Number of nodes = number of edges + 1 because this is a tree.
	n := len(edges) + 1

	// Build adjacency list for the undirected tree.
	graph := make([][]int, n+1)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// BFS from root node 1 to find the maximum depth.
	queue := []int{1}
	visited := make([]bool, n+1)
	visited[1] = true

	// depth represents number of edges from root to current BFS level.
	depth := 0
	maxDepth := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Since BFS processes level by level,
			// current depth is the depth of this node.
			maxDepth = depth

			for _, next := range graph[node] {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}

		depth++
	}

	// If maxDepth == 0, there is no edge.
	// But constraints have n >= 2, so normally maxDepth >= 1.
	if maxDepth == 0 {
		return 0
	}

	// Number of valid assignments for maxDepth edges:
	// 2^(maxDepth - 1)
	return int(modPow(2, maxDepth-1, MOD))
}

func modPow(base int64, exp int, mod int64) int64 {
	result := int64(1)

	for exp > 0 {
		if exp%2 == 1 {
			result = result * base % mod
		}

		base = base * base % mod
		exp /= 2
	}

	return result
}

func main() {
	fmt.Println(assignEdgeWeights([][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}))
}
