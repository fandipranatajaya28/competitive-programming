package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		word, isExist := patternToWord[pattern[i]]
		if isExist && word != words[i] {
			return false
		}
		patternToWord[pattern[i]] = words[i]

		patternMapped, isExist := wordToPattern[words[i]]
		if isExist && patternMapped != pattern[i] {
			return false
		}
		wordToPattern[words[i]] = pattern[i]
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
