package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int // Stores all subsets (e.g., [[], [1], [2], [1,2]])
	var current []int  // Tracks current subset during exploration

	// DFS to generate all subsets starting from index 'start'
	var dfs func(start int)
	dfs = func(start int) {
		// Save current subset to result
		// Must copy because 'current' will be modified during backtracking
		// Example: without copy, all subsets would reference the same array
		temp := make([]int, len(current))
		copy(temp, current)
		result = append(result, temp)

		// For each remaining element, decide: include it or skip it
		// 'start' prevents duplicates (e.g., [1,2] and [2,1] are the same subset)
		for i := start; i < len(nums); i++ {
			// CHOOSE: Include nums[i] in current subset
			current = append(current, nums[i])

			// EXPLORE: Recursively build subsets that include nums[i]
			// Use i+1 to only consider elements after current position
			dfs(i + 1)

			// BACKTRACK: Remove nums[i] to explore paths without it
			// This restores 'current' to try other combinations
			current = current[:len(current)-1]
		}
	}

	// Generate all subsets starting from index 0
	dfs(0)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

/*
[
  [],
  [1],
  [1, 2],
  [1, 2, 3],
  [1, 3],
  [2],
  [2, 3],
  [3]
]
```

---

## **Visual Decision Tree**
```
                         []
                         │
        ┌────────────────┼────────────────┐
        │                │                │
       [1]              [2]              [3]
        │                │
    ┌───┴───┐        ┌───┘
    │       │        │
  [1,2]   [1,3]    [2,3]
    │
    │
 [1,2,3]
```

## **Execution Order**
```
1.  Add []
2.  Choose 1      → current = [1]
3.    Add [1]
4.    Choose 2    → current = [1,2]
5.      Add [1,2]
6.      Choose 3  → current = [1,2,3]
7.        Add [1,2,3]
8.      Backtrack → current = [1,2]
9.    Backtrack   → current = [1]
10.   Choose 3    → current = [1,3]
11.     Add [1,3]
12.   Backtrack   → current = [1]
13. Backtrack     → current = []
14. Choose 2      → current = [2]
15.   Add [2]
16.   Choose 3    → current = [2,3]
17.     Add [2,3]
18.   Backtrack   → current = [2]
19. Backtrack     → current = []
20. Choose 3      → current = [3]
21.   Add [3]
22. Backtrack     → current = []
```

---

## **Table of Changes**

| Step | Call | current | Action | result |
|------|------|---------|--------|--------|
| 1 | dfs(0) | `[]` | Add [] | `[[]]` |
| 2 | dfs(0) | `[1]` | Choose 1 | `[[]]` |
| 3 | dfs(1) | `[1]` | Add [1] | `[[], [1]]` |
| 4 | dfs(1) | `[1,2]` | Choose 2 | `[[], [1]]` |
| 5 | dfs(2) | `[1,2]` | Add [1,2] | `[[], [1], [1,2]]` |
| 6 | dfs(2) | `[1,2,3]` | Choose 3 | `[[], [1], [1,2]]` |
| 7 | dfs(3) | `[1,2,3]` | Add [1,2,3] | `[[], [1], [1,2], [1,2,3]]` |
| 8 | dfs(3) | `[1,2,3]` | Return | `[[], [1], [1,2], [1,2,3]]` |
| 9 | dfs(2) | `[1,2]` | Backtrack | `[[], [1], [1,2], [1,2,3]]` |
| 10 | dfs(2) | `[1,2]` | Return | `[[], [1], [1,2], [1,2,3]]` |
| 11 | dfs(1) | `[1]` | Backtrack | `[[], [1], [1,2], [1,2,3]]` |
| 12 | dfs(1) | `[1,3]` | Choose 3 | `[[], [1], [1,2], [1,2,3]]` |
| 13 | dfs(3) | `[1,3]` | Add [1,3] | `[[], [1], [1,2], [1,2,3], [1,3]]` |
| 14 | dfs(3) | `[1,3]` | Return | `[[], [1], [1,2], [1,2,3], [1,3]]` |
| 15 | dfs(1) | `[1]` | Backtrack | `[[], [1], [1,2], [1,2,3], [1,3]]` |
| 16 | dfs(1) | `[1]` | Return | `[[], [1], [1,2], [1,2,3], [1,3]]` |
| 17 | dfs(0) | `[]` | Backtrack | `[[], [1], [1,2], [1,2,3], [1,3]]` |
| 18 | dfs(0) | `[2]` | Choose 2 | `[[], [1], [1,2], [1,2,3], [1,3]]` |
| 19 | dfs(2) | `[2]` | Add [2] | `[[], [1], [1,2], [1,2,3], [1,3], [2]]` |
| 20 | dfs(2) | `[2,3]` | Choose 3 | `[[], [1], [1,2], [1,2,3], [1,3], [2]]` |
| 21 | dfs(3) | `[2,3]` | Add [2,3] | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]` |
| 22 | dfs(3) | `[2,3]` | Return | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]` |
| 23 | dfs(2) | `[2]` | Backtrack | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]` |
| 24 | dfs(2) | `[2]` | Return | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]` |
| 25 | dfs(0) | `[]` | Backtrack | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]` |
| 26 | dfs(0) | `[3]` | Choose 3 | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]` |
| 27 | dfs(3) | `[3]` | Add [3] | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]` |
| 28 | dfs(3) | `[3]` | Return | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]` |
| 29 | dfs(0) | `[]` | Backtrack | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]` |
| 30 | dfs(0) | `[]` | Done | `[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]` |

---

## **Key Pattern to Remember**

For each number, you're making a **binary choice**:
1. **Include it** → add to current, recurse
2. **Skip it** → backtrack, try next

This creates all **2^n** possible combinations!
```
3 numbers → 2^3 = 8 subsets
*/
