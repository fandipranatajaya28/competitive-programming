package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	combined := s + s

	if strings.Contains(combined, goal) {
		return true
	}

	return false
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
}
