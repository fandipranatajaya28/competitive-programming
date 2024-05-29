package main

import "fmt"

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(arr []int, left int, right int) {
	if left < right {
		// find the middle point
		mid := (left + right) / 2

		// divide the left and right subarrays
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)

		// combine and sort the left and right subarrays
		merge(arr, left, mid, right)
	}
	return
}

func merge(arr []int, left int, mid int, right int) {
	// find the left and right subarrays size
	leftArrSize := mid - left + 1
	rightArrSize := right - mid

	// Create temporary left and right subarrays
	leftArr := make([]int, leftArrSize)
	rightArr := make([]int, rightArrSize)

	// Copy the values to left subarray
	for i := 0; i < leftArrSize; i++ {
		leftArr[i] = arr[left+i]
	}

	// Copy the values to right subarray
	for j := 0; j < rightArrSize; j++ {
		rightArr[j] = arr[mid+1+j]
	}

	// Initiate cursors for both subarrays
	leftArrCursor := 0
	rightArrCursor := 0
	mainArrCursor := left

	// Compare and sort each element in both subarrays and copy it to main array ascendingly
	for leftArrCursor < leftArrSize && rightArrCursor < rightArrSize {
		if leftArr[leftArrCursor] <= rightArr[rightArrCursor] {
			arr[mainArrCursor] = leftArr[leftArrCursor]
			leftArrCursor++
		} else {
			arr[mainArrCursor] = rightArr[rightArrCursor]
			rightArrCursor++
		}
		mainArrCursor++
	}

	// Copy the remaining left subarray elements to main array
	for leftArrCursor < leftArrSize {
		arr[mainArrCursor] = leftArr[leftArrCursor]
		leftArrCursor++
		mainArrCursor++
	}

	// Copy the remaining right subarray elements to main array
	for rightArrCursor < rightArrSize {
		arr[mainArrCursor] = rightArr[rightArrCursor]
		rightArrCursor++
		mainArrCursor++
	}
}

func main() {
	arr := []int{-2, 3, -5}
	fmt.Println(sortArray(arr))
	return
}
