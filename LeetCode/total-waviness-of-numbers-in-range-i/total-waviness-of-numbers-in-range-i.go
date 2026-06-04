package main

import (
	"fmt"
)

// func totalWaviness(num1 int, num2 int) int {
// 	ans := 0

// 	for x := num1; x <= num2; x++ {
// 		s := strconv.Itoa(x)

// 		for i := 1; i < len(s)-1; i++ {
// 			if (s[i] > s[i-1] && s[i] > s[i+1]) ||
// 				(s[i] < s[i-1] && s[i] < s[i+1]) {
// 				ans++
// 			}
// 		}
// 	}

// 	return ans
// }

func totalWaviness(num1 int, num2 int) int {
	// We only need to precompute up to num2.
	// pref[i] = total waviness of all numbers from 1 to i.
	pref := make([]int, num2+1)

	// dp[i] = waviness of number i.
	dp := make([]int, num2+1)

	// Numbers below 100 have fewer than 3 digits,
	// so their waviness is always 0.
	for i := 100; i <= num2; i++ {
		// Get the last 3 digits of i.
		//
		// Example:
		// i = 12345
		// d1 = 5
		// d2 = 4
		// d3 = 3
		d1 := i % 10
		d2 := (i / 10) % 10
		d3 := (i / 100) % 10

		wave := 0

		// d2 is the middle digit of the last 3 digits.
		// It is counted if it is a peak or valley:
		//
		// peak:   d3 < d2 > d1
		// valley: d3 > d2 < d1
		if (d2 > d3 && d2 > d1) || (d2 < d3 && d2 < d1) {
			wave = 1
		}

		// dp[i/10] is the waviness of i without the last digit.
		//
		// Adding the last digit only creates one new possible middle digit:
		// the tens digit d2.
		//
		// So:
		// dp[i] = dp[i/10] + whether d2 is wavy.
		dp[i] = dp[i/10] + wave

		// Prefix sum:
		// total waviness from 1 to i.
		pref[i] = pref[i-1] + dp[i]
	}

	// Important:
	// If num2 < 100, the loop above does not fill pref[i].
	// But all values are 0 anyway, so this is fine.
	return pref[num2] - pref[num1-1]
}

func main() {
	fmt.Println(totalWaviness(120, 130))
}
