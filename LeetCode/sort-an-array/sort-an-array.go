package main

import "fmt"

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	var (
		merged   []int
		leftIdx  int
		rightIdx int
	)

	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] <= right[rightIdx] {
			merged = append(merged, left[leftIdx])
			leftIdx++
			continue
		}
		merged = append(merged, right[rightIdx])
		rightIdx++
	}

	merged = append(merged, left[leftIdx:]...)
	merged = append(merged, right[rightIdx:]...)

	return merged
}

func main() {
	arr := []int{-2, 3, -5}
	fmt.Println(sortArray(arr))
}
