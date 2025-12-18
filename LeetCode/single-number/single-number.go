package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
