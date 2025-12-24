package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	minLen := 201
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	var result string

	for i := 0; i < minLen; i++ {
		isCommonChar := true
		commonChar := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != commonChar {
				isCommonChar = false
				break
			}
		}
		if isCommonChar {
			result += string(commonChar)
			continue
		}
		break
	}

	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
