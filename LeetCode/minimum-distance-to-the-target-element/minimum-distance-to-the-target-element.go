package main

import "fmt"

func getMinDistance(nums []int, target int, start int) int {
	if nums[start] == target {
		return 0
	}

	var minDistance = len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if nums[i] != target {
			continue
		}
		distance := abs(i, start)
		if distance < minDistance {
			minDistance = distance
		}
	}

	return minDistance
}

func abs(a int, b int) int {
	temp := a - b
	if temp < 0 {
		return -temp
	}
	return temp
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(getMinDistance(nums, 5, 3))
}
