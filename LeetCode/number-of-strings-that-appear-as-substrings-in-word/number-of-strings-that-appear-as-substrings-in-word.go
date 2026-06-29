package main

import "fmt"

func numOfStrings(patterns []string, word string) int {
	count := 0

	for _, pattern := range patterns {
		// Check whether this pattern appears inside word.
		if contains(word, pattern) {
			count++
		}
	}

	return count
}

// contains returns true if pattern appears as a substring in word.
//
// Example:
// word    = "leetcode"
// pattern = "code"
// result  = true
func contains(word string, pattern string) bool {
	n := len(word)
	m := len(pattern)

	// If pattern is longer than word, it cannot be a substring.
	if m > n {
		return false
	}

	// Try every possible starting position in word.
	for start := 0; start <= n-m; start++ {
		match := true

		// Compare word[start:start+m] with pattern.
		for i := 0; i < m; i++ {
			if word[start+i] != pattern[i] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

func main() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"

	fmt.Println(numOfStrings(patterns, word)) // Output: 3
}
