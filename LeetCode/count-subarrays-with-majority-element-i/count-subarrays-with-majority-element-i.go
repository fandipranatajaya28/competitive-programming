package main

import "fmt"

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)

	// Convert nums into prefix sum form.
	//
	// For each element:
	// nums[i] == target  => +1
	// nums[i] != target  => -1
	//
	// If a subarray sum is positive, then target appears more
	// than non-target elements, so target is the majority.
	prefix := make([]int, n+1)

	for i := 0; i < n; i++ {
		score := -1

		if nums[i] == target {
			score = 1
		}

		prefix[i+1] = prefix[i] + score
	}

	answer := 0

	// Check every subarray nums[left...right].
	//
	// subarraySum = prefix[right+1] - prefix[left]
	//
	// If subarraySum > 0, target is the majority element
	// in this subarray.
	for left := 0; left < n; left++ {
		for right := left; right < n; right++ {
			subarraySum := prefix[right+1] - prefix[left]

			if subarraySum > 0 {
				answer++
			}
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 2, 3}
	target := 2

	fmt.Println(countMajoritySubarrays(nums, target)) // Output: 5
}
