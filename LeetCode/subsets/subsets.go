package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int // Store all subsets
	var current []int  // Current subset being built

	var dfs func(start int)
	dfs = func(start int) {
		// SAVE current subset (make a copy!)
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)

		// Try adding each remaining number
		for i := start; i < len(nums); i++ {
			// CHOOSE: Add nums[i] to current subset
			current = append(current, nums[i])

			// EXPLORE: Generate all subsets that include nums[i]
			dfs(i + 1)

			// BACKTRACK: Remove nums[i], try without it
			current = current[:len(current)-1]
		}
	}

	dfs(0)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
