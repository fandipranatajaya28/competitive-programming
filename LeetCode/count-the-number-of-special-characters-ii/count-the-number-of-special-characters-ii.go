package main

import "fmt"

func numberOfSpecialChars(word string) int {
	lowerLast := make([]int, 26)
	upperFirst := make([]int, 26)

	for i := 0; i < 26; i++ {
		lowerLast[i] = -1
		upperFirst[i] = -1
	}

	for idx, c := range word {
		if c >= 'a' && c <= 'z' {
			lowerLast[c-'a'] = idx
		} else {
			pos := c - 'A'
			if upperFirst[pos] == -1 {
				upperFirst[pos] = idx
			}
		}
	}

	res := 0
	for i := 0; i < 26; i++ {
		if lowerLast[i] != -1 &&
			upperFirst[i] != -1 &&
			lowerLast[i] < upperFirst[i] {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
}
