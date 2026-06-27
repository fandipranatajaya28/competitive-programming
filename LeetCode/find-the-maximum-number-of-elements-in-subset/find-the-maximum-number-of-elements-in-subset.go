package main

import "fmt"

func maximumLength(nums []int) int {
	// Count frequency of each number.
	// Use int64 because we will square numbers.
	count := make(map[int64]int)

	for _, num := range nums {
		count[int64(num)]++
	}

	answer := 1

	for num, freq := range count {
		// Special case:
		// 1^2 is still 1, so the sequence is only made of 1s.
		//
		// We need an odd length because the valid shape has:
		// pair + pair + ... + center
		//
		// If we have 5 ones, we can use 5.
		// If we have 4 ones, we can only use 3.
		if num == 1 {
			if freq%2 == 1 {
				answer = max(answer, freq)
			} else {
				answer = max(answer, freq-1)
			}
			continue
		}

		length := 0
		curr := num

		// For each level, we need 2 copies:
		//
		// x appears twice
		// x^2 appears twice
		// x^4 appears twice
		// ...
		//
		// Until we choose one final center.
		for count[curr] >= 2 {
			length += 2
			curr = curr * curr
		}

		// Now count[curr] is either:
		// 1 -> we can use it as the center
		// 0 -> no valid center, so the last pair cannot be used
		if count[curr] >= 1 {
			length++
		} else {
			// Example:
			// nums has [2, 2], but no 4.
			//
			// We cannot form [2, 2] because valid length must be odd.
			// The best is just [2], length 1.
			length--
		}

		answer = max(answer, length)
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{5, 4, 1, 2, 2, 4, 16}

	fmt.Println(maximumLength(nums)) // Output: 5
}
