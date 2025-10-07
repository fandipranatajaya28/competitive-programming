package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var maxLength int

	if len(s) < 2 {
		return len(s)
	}

	for i := 0; i < len(s); i++ {
		var (
			hashMap    = make(map[byte]bool)
			tempLength = 1
		)
		hashMap[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if hashMap[s[j]] {
				break
			}
			hashMap[s[j]] = true
			tempLength++
		}
		if tempLength > maxLength {
			maxLength = tempLength
		}
	}

	return maxLength
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
