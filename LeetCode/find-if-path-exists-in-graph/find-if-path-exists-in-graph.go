package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	var (
		adjList   = make([][]int, n)
		queue     []int
		isVisited = make([]bool, n)
	)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	queue = append(queue, source)
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if isVisited[currNode] {
			continue
		}

		isVisited[currNode] = true

		if currNode == destination {
			return true
		}

		queue = append(queue, adjList[currNode]...)
	}

	return false
}

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
}
