package main

import "fmt"

func maxTotalValue(nums []int, k int) int64 {
	var (
		answer int64
		minNum = nums[0]
		maxNum = nums[0]
	)

	for _, num := range nums {
		if num < minNum {
			minNum = num
		}

		if num > maxNum {
			maxNum = num
		}
	}

	answer = int64(maxNum-minNum) * int64(k)

	return answer
}

func main() {
	fmt.Println(maxTotalValue([]int{4, 2, 5, 1}, 3))
}
