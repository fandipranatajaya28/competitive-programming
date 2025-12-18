package main

import (
	"fmt"
)

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i := 0; i < n; i++ {
		ans[i*2] = nums[i]
		ans[i*2+1] = nums[n+i]
	}
	return ans
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
