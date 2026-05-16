package main

import (
	"fmt"
)

func isGood(nums []int) bool {
	maxNum := len(nums) - 1 // In a good array, max must equal len-1
	freq := make([]int, maxNum+1)

	for _, num := range nums {
		if num < 1 || num > maxNum {
			return false // Early exit: out of range number
		}
		freq[num]++
		if freq[num] > 2 || (num != maxNum && freq[num] > 1) {
			return false // Early exit: too many duplicates
		}
	}

	return freq[maxNum] == 2
}

func main() {
	fmt.Println(isGood([]int{1, 3, 3, 2}))
}
