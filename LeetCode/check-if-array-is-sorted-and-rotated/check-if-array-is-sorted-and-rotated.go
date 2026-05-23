package main

import "fmt"

func check(nums []int) bool {
	var numOfOrderBreaks int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			numOfOrderBreaks++
		}
	}

	if nums[len(nums)-1] > nums[0] {
		numOfOrderBreaks++
	}

	return numOfOrderBreaks < 2
}

func main() {
	fmt.Println(check([]int{3, 4, 5, 1, 2}))
}
