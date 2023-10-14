package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) (result []int) {
	// class to store value and its original index
	type number struct {
		value int
		index int
	}

	var (
		left, right int
		numsArr     []number
	)

	for i, v := range nums {
		numsArr = append(numsArr, number{
			value: v,
			index: i,
		})
	}

	// sort the elements
	sort.Slice(numsArr, func(i, j int) bool {
		return numsArr[i].value < numsArr[j].value
	})

	// init pointers
	left = 0
	right = len(numsArr) - 1

	for left < right {
		if numsArr[left].value+numsArr[right].value == target {
			result = append(result, numsArr[left].index, numsArr[right].index)
			return
		} else if numsArr[left].value+numsArr[right].value < target {
			left++
		} else { // numsArr[left].value+numsArr[right].value > target
			right--
		}
	}
	return
}

func main() {
	nums := []int{11, 15, 2, 7}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
