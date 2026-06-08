package main

import "fmt"

// func pivotArray(nums []int, pivot int) []int {
// 	// Result array.
// 	// We append numbers in the required order.
// 	ans := []int{}

// 	// First pass:
// 	// Add all numbers smaller than pivot.
// 	// Their original relative order is preserved.
// 	for _, num := range nums {
// 		if num < pivot {
// 			ans = append(ans, num)
// 		}
// 	}

// 	// Second pass:
// 	// Add all numbers equal to pivot.
// 	for _, num := range nums {
// 		if num == pivot {
// 			ans = append(ans, num)
// 		}
// 	}

// 	// Third pass:
// 	// Add all numbers greater than pivot.
// 	// Their original relative order is also preserved.
// 	for _, num := range nums {
// 		if num > pivot {
// 			ans = append(ans, num)
// 		}
// 	}

// 	return ans
// }

func pivotArray(nums []int, pivot int) []int {
	n := len(nums)

	// Result array with the same length as nums.
	ans := make([]int, n)

	// left points to the next position for numbers smaller than pivot.
	left := 0

	// right points to the next position for numbers greater than pivot.
	right := n - 1

	// Process from both sides.
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		// Scan from left to right.
		// This keeps the relative order of numbers smaller than pivot.
		if nums[i] < pivot {
			ans[left] = nums[i]
			left++
		}

		// Scan from right to left.
		// We place greater numbers from the back.
		// This keeps their relative order in the final array.
		if nums[j] > pivot {
			ans[right] = nums[j]
			right--
		}
	}

	// Fill the remaining middle positions with pivot.
	for left <= right {
		ans[left] = pivot
		left++
	}

	return ans
}

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
}
