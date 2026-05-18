package main

import (
	"fmt"
	"strconv"
)

// func separateDigits(nums []int) []int {
// 	ans := []int{}

// 	for _, num := range nums {
// 		digits := getDigits(num)
// 		ans = append(ans, digits...)
// 	}

// 	return ans
// }

// func getDigits(num int) []int {
// 	temp := []int{}

// 	for num > 0 {
// 		temp = append(temp, num%10)
// 		num /= 10
// 	}

// 	for l, r := 0, len(temp)-1; l < r; l, r = l+1, r-1 {
// 		temp[l], temp[r] = temp[r], temp[l]
// 	}

// 	return temp
// }

func separateDigits(nums []int) []int {
	ans := []int{}

	for _, num := range nums {
		s := strconv.Itoa(num)

		for _, ch := range s {
			ans = append(ans, int(ch-'0'))
		}
	}

	return ans
}

func main() {
	fmt.Println(separateDigits([]int{13, 25, 83, 77}))
}
