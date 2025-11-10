package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

func quickSort(nums []int, low int, high int) {
	if low < high {
		/* pi is partitioning index,
		   arr[p] is now at right place */
		pi := partition(nums, low, high)

		// Separately sort elements before
		// partition and after partition
		quickSort(nums, low, pi)
		quickSort(nums, pi+1, high)
	}
}

// QuickSort using Hoare's partition scheme.
/*
	This function takes first element as pivot, and places
	all the elements greater than the pivot on the left side
	and all the elements smaller than the pivot on
	the right side. It returns the index of the last element
	on the greater side
*/
func partition(nums []int, low int, high int) int {
	// choose the pivot
	pivot := nums[low]

	i := low - 1
	j := high + 1

	for {
		// Find leftmost element smaller than
		// or equal to pivot
		i++
		for nums[i] > pivot {
			i++
		}

		// Find rightmost element greater than
		// or equal to pivot
		j--
		for nums[j] < pivot {
			j--
		}

		// If two pointers met.
		if i >= j {
			return j
		}

		swap(nums, i, j)
	}
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}
