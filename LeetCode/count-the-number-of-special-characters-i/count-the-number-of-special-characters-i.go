package main

import (
	"fmt"
	"math/bits"
)

// func numberOfSpecialChars(word string) int {
// 	lower := [26]bool{}
// 	upper := [26]bool{}

// 	for _, c := range word {
// 		if c >= 'a' && c <= 'z' {
// 			lower[c-'a'] = true
// 		} else {
// 			upper[c-'A'] = true
// 		}
// 	}

// 	res := 0
// 	for i := 0; i < 26; i++ {
// 		if lower[i] && upper[i] {
// 			res++
// 		}
// 	}

// 	return res
// }

func numberOfSpecialChars(word string) int {
	var lowerMask, upperMask uint32

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			lowerMask |= 1 << (c - 'a')
		} else {
			upperMask |= 1 << (c - 'A')
		}
	}

	commonMask := lowerMask & upperMask

	return bits.OnesCount32(commonMask)
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
}
