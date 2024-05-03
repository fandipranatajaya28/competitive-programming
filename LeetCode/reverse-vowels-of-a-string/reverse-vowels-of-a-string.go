package main

import "fmt"

func reverseVowels(s string) string {
	result := ""
	reversedVowels := []string{}

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if isVowels(c) {
			reversedVowels = append(reversedVowels, c)
		}
	}

	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if isVowels(c) {
			c = reversedVowels[len(reversedVowels)-1]
			reversedVowels = reversedVowels[:len(reversedVowels)-1]
		}
		result += c
	}

	return result
}

func isVowels(s string) bool {
	switch s {
	case "a", "A", "i", "I", "u", "U", "e", "E", "o", "O":
		return true
	}
	return false
}

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))
}
