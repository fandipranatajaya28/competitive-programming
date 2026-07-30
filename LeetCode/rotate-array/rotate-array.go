package main

import "fmt"

// Rotate Array:
// Rotate nums to the RIGHT by k steps, in place.
// Example: [1,2,3,4,5,6,7], k=3  ->  [5,6,7,1,2,3,4]
//
// Triple-reversal trick (O(1) extra space):
//   1. Reverse the entire array.
//   2. Reverse the first k elements.
//   3. Reverse the remaining n-k elements.
//
// Why it works: step 1 moves the last k elements to the front and the first
// n-k elements to the back, but each block ends up internally reversed.
// Steps 2 and 3 un-reverse each block in its new position.
//
//   start:        [1,2,3,4,5,6,7]
//   reverse all:  [7,6,5,4,3,2,1]
//   reverse[0:3]: [5,6,7,4,3,2,1]
//   reverse[3:]:  [5,6,7,1,2,3,4]  <- done
//
// Time: O(n). Space: O(1).
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	// k may exceed n. Rotating by n is a no-op, so only the remainder matters.
	// (Without this, the index math below would go out of range.)
	k %= n
	if k == 0 {
		return
	}

	reverse(nums, 0, n-1) // step 1: whole array
	reverse(nums, 0, k-1) // step 2: first k elements
	reverse(nums, k, n-1) // step 3: the rest
}

// reverse reverses nums[left..right] in place using two pointers walking inward.
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	// Example test cases.
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a) // [5 6 7 1 2 3 4]

	b := []int{-1, -100, 3, 99}
	rotate(b, 2)
	fmt.Println(b) // [3 99 -1 -100]

	// k larger than the array length: 7 % 4 == 3, so this rotates by 3.
	c := []int{1, 2, 3, 4}
	rotate(c, 7)
	fmt.Println(c) // [2 3 4 1]

	// k == 0: unchanged.
	d := []int{1, 2, 3}
	rotate(d, 0)
	fmt.Println(d) // [1 2 3]
}
