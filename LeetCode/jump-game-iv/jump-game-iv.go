package main

import "fmt"

func minJumps(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 0
	}

	sameNumIdxMap := make(map[int][]int)
	for idx, num := range arr {
		sameNumIdxMap[num] = append(sameNumIdxMap[num], idx)
	}

	queue := []int{0}
	isVisited := make([]bool, n)
	isVisited[0] = true
	minDistance := 0

	for len(queue) > 0 {
		currLevelSize := len(queue)

		for i := 0; i < currLevelSize; i++ {
			currPos := queue[0]
			queue = queue[1:]

			if currPos == n-1 {
				return minDistance
			}

			// Try teleport to same-value indices
			for _, next := range sameNumIdxMap[arr[currPos]] {
				if !isVisited[next] {
					isVisited[next] = true
					queue = append(queue, next)
				}
			}

			// Important optimization:
			// Avoid processing the same value group again.
			delete(sameNumIdxMap, arr[currPos])

			// Try right
			if currPos+1 < n && !isVisited[currPos+1] {
				isVisited[currPos+1] = true
				queue = append(queue, currPos+1)
			}

			// Try left
			if currPos-1 >= 0 && !isVisited[currPos-1] {
				isVisited[currPos-1] = true
				queue = append(queue, currPos-1)
			}
		}

		minDistance++
	}

	return -1
}

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}
