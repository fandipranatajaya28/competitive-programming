package main

import "fmt"

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	const MOD int64 = 1_000_000_007

	n := len(edges) + 1

	// Build graph because edges represent an undirected tree.
	graph := make([][]int, n+1)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// LOG is enough for binary lifting up to n nodes.
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}

	// up[j][node] = the 2^j-th ancestor of node.
	up := make([][]int, LOG)
	for i := 0; i < LOG; i++ {
		up[i] = make([]int, n+1)
	}

	// depth[node] = distance in edges from root node 1.
	depth := make([]int, n+1)

	// BFS from root 1 to fill depth[] and up[0][].
	queue := []int{1}
	visited := make([]bool, n+1)
	visited[1] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, next := range graph[node] {
			if visited[next] {
				continue
			}

			visited[next] = true
			depth[next] = depth[node] + 1

			// Direct parent of next is node.
			up[0][next] = node

			queue = append(queue, next)
		}
	}

	// Build binary lifting table.
	for j := 1; j < LOG; j++ {
		for node := 1; node <= n; node++ {
			midAncestor := up[j-1][node]
			up[j][node] = up[j-1][midAncestor]
		}
	}

	// Precompute powers of 2.
	// pow2[i] = 2^i mod MOD.
	pow2 := make([]int64, n+1)
	pow2[0] = 1

	for i := 1; i <= n; i++ {
		pow2[i] = pow2[i-1] * 2 % MOD
	}

	// Helper function to find Lowest Common Ancestor of u and v.
	lca := func(u int, v int) int {
		// Make sure u is deeper or equal.
		if depth[u] < depth[v] {
			u, v = v, u
		}

		// Lift u up until both nodes have the same depth.
		diff := depth[u] - depth[v]
		for j := 0; j < LOG; j++ {
			if diff&(1<<j) != 0 {
				u = up[j][u]
			}
		}

		// If v was ancestor of u.
		if u == v {
			return u
		}

		// Lift both nodes together while their ancestors differ.
		for j := LOG - 1; j >= 0; j-- {
			if up[j][u] != up[j][v] {
				u = up[j][u]
				v = up[j][v]
			}
		}

		// Now their parent is the LCA.
		return up[0][u]
	}

	ans := make([]int, len(queries))

	for i, query := range queries {
		u := query[0]
		v := query[1]

		ancestor := lca(u, v)

		// Number of edges on path u -> v.
		dist := depth[u] + depth[v] - 2*depth[ancestor]

		// If u == v, there are no edges.
		// Sum is 0, which is even, so there are 0 valid assignments.
		if dist == 0 {
			ans[i] = 0
			continue
		}

		// For dist edges:
		// each edge can be 1 or 2.
		// We need odd total sum.
		// Only edges assigned weight 1 affect parity.
		// Exactly half of all 2^dist assignments have odd parity.
		// So answer = 2^(dist-1).
		ans[i] = int(pow2[dist-1])
	}

	return ans
}

func main() {
	fmt.Println(assignEdgeWeights([][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}, [][]int{{1, 4}, {3, 4}, {2, 5}}))
}
