package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var result [][]int // Stores all unique subsets
	var current []int  // Current subset being built

	var dfs func(start int)
	dfs = func(start int) {
		// Save current subset (make a copy to avoid reference issues)
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)

		// Try adding each remaining number
		for i := start; i < len(nums); i++ {
			// Skip duplicates at the same level to avoid duplicate subsets
			// i > start: ensures we don't skip the first element at this level
			// nums[i] == nums[i-1]: current element is same as previous
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// Choose: add nums[i] to current subset
			current = append(current, nums[i])

			// Explore: recursively build subsets that include nums[i]
			dfs(i + 1)

			// Backtrack: remove nums[i] to try other combinations
			current = current[:len(current)-1]
		}
	}

	// Sort to group duplicates together (required for duplicate detection)
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	// Start building subsets from index 0
	dfs(0)
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

/* Visual Decision Tree
```
                         []
                         │
        ┌────────────────┼────────────┐
        │                │            │
       [1]              [2]          [2] ← SKIPPED (duplicate at same level)
        │                │
    ┌───┴───┐            │
    │       │            │
  [1,2]  [1,2] ← SKIPPED  [2,2]
    │
    │
 [1,2,2]

*/
