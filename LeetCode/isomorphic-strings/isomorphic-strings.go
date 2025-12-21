package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Store the t char pair in s map
	sMap := make(map[byte]byte)
	// Store the s char pair in t map
	tMap := make(map[byte]byte)

	// Check for each char whether the mapped pair is missmatch or not
	for idx := 0; idx < len(s); idx++ {
		tPair, isExist := sMap[s[idx]]
		if isExist && tPair != t[idx] {
			return false
		}
		sMap[s[idx]] = t[idx]

		sPair, isExist := tMap[t[idx]]
		if isExist && sPair != s[idx] {
			return false
		}
		tMap[t[idx]] = s[idx]
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}
