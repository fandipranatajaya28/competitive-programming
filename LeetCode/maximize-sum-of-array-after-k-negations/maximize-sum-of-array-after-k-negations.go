package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	var (
		maxSum         int
		remainingFlips = k
	)
	for _, num := range nums {
		if remainingFlips > 0 && num < 0 {
			num = -num
			remainingFlips--
		}
		maxSum += num
	}

	if remainingFlips%2 == 1 {
		maxSum -= 2*abs(nums[len(nums)-1])
	}

	return maxSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}
