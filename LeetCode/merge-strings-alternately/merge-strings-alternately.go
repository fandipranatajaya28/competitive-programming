package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	mergedWord := ""
	limit := len(word1)
	if len(word2) > limit {
		limit = len(word2)
	}
	for i := 0; i < limit; i++ {
		if i < len(word1) {
			mergedWord += string(word1[i])
		}
		if i < len(word2) {
			mergedWord += string(word2[i])
		}
	}
	return mergedWord
}

func main() {
	a := "abcd"
	b := "pq"
	fmt.Println(mergeAlternately(a, b))
}
