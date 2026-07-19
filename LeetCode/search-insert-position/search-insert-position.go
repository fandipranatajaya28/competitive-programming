package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// Midpoint written this way to avoid potential integer overflow.
		mid := left + (right-left)/2

		switch {
		case nums[mid] == target:
			// Exact match: return its index.
			return mid
		case nums[mid] < target:
			// Target is to the right of mid.
			left = mid + 1
		default:
			// nums[mid] > target: target is at or to the left of mid.
			right = mid - 1
		}
	}

	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
}
