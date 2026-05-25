package main

import "fmt"

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)

	// dp[i] means: can we reach index i?
	dp := make([]bool, n)

	// We always start from index 0.
	dp[0] = true

	// count represents how many reachable indices exist
	// inside the current valid jump window.
	//
	// For index i, we need a previous index j such that:
	// i - maxJump <= j <= i - minJump
	//
	// Instead of checking all j one by one,
	// we maintain the count of reachable j's in that window.
	count := 0

	// Start from minJump because indices before minJump
	// cannot be reached from index 0.
	for i := minJump; i < n; i++ {
		// Add the new left-side candidate into the window.
		//
		// For current i, index i-minJump is now close enough
		// to jump from.
		if dp[i-minJump] {
			count++
		}

		// Remove the index that is now too far away.
		//
		// If i > maxJump, then index i-maxJump-1
		// is no longer able to jump to i.
		if i > maxJump && dp[i-maxJump-1] {
			count--
		}

		// We can reach i if:
		// 1. s[i] == '0', because we can only land on '0'
		// 2. count > 0, meaning at least one reachable previous index
		//    can jump to i.
		if s[i] == '0' && count > 0 {
			dp[i] = true
		}
	}

	// Return whether the last index is reachable.
	return dp[n-1]
}

func main() {
	fmt.Println(canReach("011010", 2, 3))
}
