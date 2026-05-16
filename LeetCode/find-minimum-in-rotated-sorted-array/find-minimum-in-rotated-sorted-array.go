package main

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid // mid could be the answer, don't exclude it
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func main() {
	fmt.Println([]int{4, 5, 6, 7, 0, 1, 2})
}
