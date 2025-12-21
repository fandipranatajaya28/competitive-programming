package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	hashMap := make(map[rune]int)
	for _, char := range magazine {
		hashMap[char]++
	}

	for _, char := range ransomNote {
		if hashMap[char] <= 0 {
			return false
		}
		hashMap[char]--
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
