package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var (
		result          [][]string
		anagramGroupMap = make(map[string][]string)
	)

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		anagramGroupMap[string(b)] = append(anagramGroupMap[string(b)], str)
	}

	for _, anagramGroup := range anagramGroupMap {
		result = append(result, anagramGroup)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
