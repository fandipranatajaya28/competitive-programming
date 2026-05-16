package main

import "fmt"

func firstUniqChar(s string) int {
	uniqCharMap := make(map[rune]int)
	for _, char := range s {
		uniqCharMap[char]++
	}
	for i, char := range s {
		if uniqCharMap[char] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
