package main

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	low := 0
	high := nums[len(nums)-1] - nums[0]
	var minResult int
	for low <= high {
		mid := low + (high-low)/2
		if canFormPairs(nums, p, mid) {
			high = mid - 1
			minResult = mid
		} else {
			low = mid + 1
		}
	}

	return minResult
}

func canFormPairs(nums []int, p int, mid int) bool {
	count := 0
	i := 0
	for i < len(nums)-1 && count < p {
		if nums[i+1]-nums[i] <= mid {
			count++
			i += 2
		} else {
			i++
		}
	}
	return count >= p
}

func main() {
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 2))
}
