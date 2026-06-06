package main

import (
	"fmt"
)

func leftRightDifference(nums []int) []int {
	answer := make([]int, len(nums))

	leftSum := nums[0]
	for i := 1; i < len(nums); i++ {
		answer[i] = leftSum
		leftSum += nums[i]
	}

	rightSum := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		answer[i] = abs(answer[i] - rightSum)
		rightSum += nums[i]
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3}))
}
