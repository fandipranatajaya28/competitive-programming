package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid // min is in left half (inclusive)
		} else if nums[mid] == nums[right] {
			right-- // can't tell, safely shrink
		} else {
			left = mid + 1 // min is in right half
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}
