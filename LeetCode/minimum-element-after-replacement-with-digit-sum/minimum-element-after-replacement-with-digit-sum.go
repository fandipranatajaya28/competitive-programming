package main

import (
	"fmt"
	"math"
)

func minElement(nums []int) int {
	var minElem = math.MaxInt32
	for i, n := range nums {
		nums[i] = 0
		for n > 0 {
			nums[i] += n % 10
			n /= 10
		}
		if nums[i] < minElem {
			minElem = nums[i]
		}
	}
	return minElem
}

func main() {
	fmt.Println(minElement([]int{111, 19, 199}))
}
