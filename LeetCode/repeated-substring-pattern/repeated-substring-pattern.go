package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	/**
	  abcabc
	  abcabcabcabc

	  b c abcabc a b
	*/

	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)

	// Use string rotation
	// sRotation := s
	// for i := 0; i < len(s)-1; i++ {
	// 	sRotation = string(sRotation[len(s)-1]) + sRotation[:len(s)-1]
	// 	if sRotation == s {
	// 		return true
	// 	}
	// }

	// return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
