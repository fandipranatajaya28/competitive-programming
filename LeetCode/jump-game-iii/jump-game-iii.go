package main

import "fmt"

func canReach(arr []int, start int) bool {
	isVisited := make([]bool, len(arr))
	queue := []int{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if isVisited[curr] {
			continue
		}

		isVisited[curr] = true

		if arr[curr] == 0 {
			return true
		}

		left := curr - arr[curr]
		right := curr + arr[curr]

		if left >= 0 && !isVisited[left] {
			queue = append(queue, left)
		}

		if right < len(arr) && !isVisited[right] {
			queue = append(queue, right)
		}
	}

	return false
}

func main() {
	fmt.Println(canReach([]int{0, 3, 0, 6, 3, 3, 4}, 6))
}
