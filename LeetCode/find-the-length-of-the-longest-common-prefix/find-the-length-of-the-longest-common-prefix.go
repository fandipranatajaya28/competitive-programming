package main

import "fmt"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	var (
		lcpLen      int
		prefixesMap = make(map[int]bool)
	)

	for _, num := range arr1 {
		for num != 0 {
			prefixesMap[num] = true
			num /= 10
		}
	}

	for _, num := range arr2 {
		for num != 0 {
			if prefixesMap[num] {
				len := intLen(num)
				if len > lcpLen {
					lcpLen = len
				}
			}
			num /= 10
		}
	}

	return lcpLen
}

func intLen(num int) int {
	var len int
	for num != 0 {
		len++
		num /= 10
	}
	return len
}

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 10, 100}, []int{1000}))
}
