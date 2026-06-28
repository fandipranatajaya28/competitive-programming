package main

import (
	"fmt"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	maxElement := 1
	sort.Ints(arr)
	arr[0] = 1

	for i := 1; i < len(arr); i++ {
		if (arr[i] - arr[i-1]) > 1 {
			arr[i] = arr[i-1] + 1
		}
		maxElement = arr[i]
	}

	return maxElement
}

func main() {
	fmt.Println(maximumElementAfterDecrementingAndRearranging([]int{1, 2, 3, 4, 5}))
}
