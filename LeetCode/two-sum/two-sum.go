package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int)
	for idx, num := range nums {
		remainder := target - num
		if remainderIdx, isExist := numMap[remainder]; isExist {
			return []int{idx, remainderIdx}
		}
		numMap[num] = idx
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
