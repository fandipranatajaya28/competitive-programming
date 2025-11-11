package main

import "fmt"

func searchRange(nums []int, target int) []int {
	// Recursive version
	// lowRange := binarySearch(nums, target, 0, len(nums)-1, true)
	// highRange := binarySearch(nums, target, 0, len(nums)-1, false)

	// Iterative version
	lowRange := lowRangeBinarySearch(nums, target)
	highRange := highRangeBinarySearch(nums, target)
	return []int{lowRange, highRange}
}

// Recursive Version
func binarySearch(nums []int, target, low, high int, findLow bool) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if nums[mid] > target {
		return binarySearch(nums, target, low, mid-1, findLow)
	} else if nums[mid] < target {
		return binarySearch(nums, target, mid+1, high, findLow)
	}

	// nums[mid] == target
	if findLow {
		// Keep searching to the left for the first occurrence
		left := binarySearch(nums, target, low, mid-1, findLow)
		if left != -1 {
			return left
		}
	} else {
		// Keep searching to the right for the last occurrence
		right := binarySearch(nums, target, mid+1, high, findLow)
		if right != -1 {
			return right
		}
	}

	return mid
}

// Iterative version
func lowRangeBinarySearch(nums []int, target int) (index int) {
	low := 0
	high := len(nums) - 1
	index = -1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}

		if nums[mid] == target {
			index = mid
		}
	}

	return
}

// Iterative version
func highRangeBinarySearch(nums []int, target int) (index int) {
	low := 0
	high := len(nums) - 1
	index = -1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}

		if nums[mid] == target {
			index = mid
		}
	}

	return
}

func main() {
	fmt.Println(searchRange([]int{8, 8, 8, 8, 8, 8}, 8))
}
