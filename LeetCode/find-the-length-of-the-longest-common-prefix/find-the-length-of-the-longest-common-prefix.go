package main

import "fmt"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixMap := make(map[int]int) // map[prefix]length

	// Store all prefixes with their lengths
	for _, num := range arr1 {
		length := 0
		temp := num
		for temp > 0 {
			length++
			temp /= 10
		}

		// Generate all prefixes
		for num > 0 {
			prefixMap[num] = length
			num /= 10
			length--
		}
	}

	maxLen := 0

	// Find longest matching prefix
	for _, num := range arr2 {
		for num > 0 {
			if length, exists := prefixMap[num]; exists {
				if length > maxLen {
					maxLen = length
				}
				break // Found longest for this number
			}
			num /= 10
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000})) // 3
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4})) // 0
}
