package main

import (
	"fmt"
)

func minReorder(n int, connections [][]int) int {
	isVisited := make(map[int]bool)
	graph := make([][]int, n)
	for _, val := range connections {
		// positive is directed outside the source, negative is directed into the source
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], -val[0])
	}
	queue := []int{}
	for _, val := range graph[0] {
		queue = append(queue, val)
	}

	ans := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			// if the node is directed outside the source then the edge need to be reordered
			if curr > 0 && !isVisited[abs(curr)] {
				ans++
			}

			if len(graph[abs(curr)]) > 0 && !isVisited[abs(curr)] {
				for _, val := range graph[abs(curr)] {
					queue = append(queue, val)
				}
			}

			isVisited[abs(curr)] = true
		}
	}
	return ans
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	n := 6
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	fmt.Println(minReorder(n, connections))
	return
}
