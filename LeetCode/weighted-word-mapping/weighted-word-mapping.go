package main

import "fmt"

func mapWordWeights(words []string, weights []int) string {
	var result string

	for _, word := range words {
		weight := 0
		for _, letter := range word {
			weight += weights[letter-'a']
		}
		result += string(rune('z' - weight%26))
	}

	return result
}

func main() {
	fmt.Println(mapWordWeights([]string{"abcd", "def", "xyz"}, []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}))
}
