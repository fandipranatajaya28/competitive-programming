package main

import "fmt"

func maxJumps(arr []int, d int) int {
	n := len(arr)

	// dp[i] stores the maximum number of indices we can visit
	// starting from index i.
	// -1 means not calculated yet.
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		// If we already calculated the answer for index i,
		// return it directly to avoid repeated work.
		if dp[i] != -1 {
			return dp[i]
		}

		// At minimum, we can visit the current index itself.
		ans := 1

		// Try jumping to the right.
		// We can jump at most d steps and cannot go outside the array.
		for j := i + 1; j <= min(n-1, i+d); j++ {
			// If arr[j] is greater than or equal to arr[i],
			// then j blocks the path.
			// We cannot jump to j or beyond j in this direction.
			if arr[j] >= arr[i] {
				break
			}

			// Jump to j, then continue DFS from j.
			ans = max(ans, 1+dfs(j))
		}

		// Try jumping to the left.
		// We can jump at most d steps and cannot go outside the array.
		for j := i - 1; j >= max(0, i-d); j-- {
			// Same blocking rule:
			// if arr[j] >= arr[i], we cannot jump to j or beyond it.
			if arr[j] >= arr[i] {
				break
			}

			// Jump to j, then continue DFS from j.
			ans = max(ans, 1+dfs(j))
		}

		// Save the result for index i before returning.
		dp[i] = ans

		return ans
	}

	answer := 1

	// Since we can start from any index,
	// calculate the best answer among all starting positions.
	for i := 0; i < n; i++ {
		answer = max(answer, dfs(i))
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
}
