package main

import (
	"fmt"
	"strings"
)

func longestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		oddPalindrome := expandFromCenter(s, i, i)
		evenPalindrome := expandFromCenter(s, i, i+1)

		if len(oddPalindrome) > len(result) {
			result = oddPalindrome
		}

		if len(evenPalindrome) > len(result) {
			result = evenPalindrome
		}
	}
	return result
}

func expandFromCenter(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// Using Manacher's algorithm https://www.youtube.com/watch?v=6Bq8j2dhzJc
func longestPalindromeManacher(s string) string {
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)
	P := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		if R > i {
			P[i] = min(R-i, P[2*C-i])
		}
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
