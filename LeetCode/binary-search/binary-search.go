package main

import "fmt"

func search(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   = (left + right) / 2
	)

	for left <= right {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
			mid = (left + right) / 2
		} else {
			right = mid - 1
			mid = (left + right) / 2
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	firstTarget := 1
	secondTarget := 5
	fmt.Println(search(nums, firstTarget))
	fmt.Println(search(nums, secondTarget))
}
