package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var result [][]int // Stores all unique subsets
	var current []int  // Current subset being built (backtracking state)

	// DFS function to generate all subsets starting from index 'start'
	var dfs func(start int)
	dfs = func(start int) {
		// Add current subset to result
		// Must copy because 'current' slice will be modified during backtracking
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)

		// Try including each number from 'start' onwards
		for i := start; i < len(nums); i++ {
			// Skip duplicates at the same recursion level
			// Example: [1, 2₁, 2₂] at level starting from index 1
			//   - i=1 (start=1): use 2₁  ✓
			//   - i=2 (start=1): skip 2₂ (i > start && 2₂ == 2₁) ✓
			// This prevents generating [1,2] twice
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// CHOOSE: include nums[i] in the current subset
			current = append(current, nums[i])

			// EXPLORE: recursively generate subsets with nums[i] included
			dfs(i + 1)

			// BACKTRACK: remove nums[i] to explore subsets without it
			current = current[:len(current)-1]
		}
	}

	// Sort array to group duplicate elements together
	// This is required for the duplicate-skipping logic to work
	sort.Slice(nums, func(x, y int) bool {
		return nums[x] < nums[y]
	})

	// Start generating subsets from index 0
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
