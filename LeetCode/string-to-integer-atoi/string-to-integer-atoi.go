package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	i := 0
	n := len(s)

	// Step 1: skip leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	// Step 2: check sign
	sign := 1
	if i < n && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Step 3: read digits
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Step 4: check overflow before it happens
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}

func main() {
	fmt.Println(myAtoi("-042"))
}
