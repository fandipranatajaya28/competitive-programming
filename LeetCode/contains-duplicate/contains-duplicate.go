package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	duplicateMap := make(map[int]bool)
	for _, num := range nums {
		if _, isExist := duplicateMap[num]; isExist {
			return true
		}
		duplicateMap[num] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
