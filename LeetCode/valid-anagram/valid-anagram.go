package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charCounts [26]int
	for i := 0; i < len(s); i++ {
		charCounts[s[i]-'a']++
		charCounts[t[i]-'a']--
	}

	for _, charCount := range charCounts {
		if charCount != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
