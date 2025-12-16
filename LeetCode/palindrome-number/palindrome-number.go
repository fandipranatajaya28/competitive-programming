package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	reversed := 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num = num / 10
	}

	return reversed == x
}

func main() {
	fmt.Println(isPalindrome(121))
}
