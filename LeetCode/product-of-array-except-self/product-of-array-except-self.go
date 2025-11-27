package main

import "fmt"

// Create a new array where each element is the product of all the elements in the original array except for the one at its own index
func productExceptSelf(nums []int) []int {
	var ans = make([]int, len(nums))

	prefixProduct := 1
	for i := range len(nums) {
		ans[i] = prefixProduct
		prefixProduct *= nums[i]
	}

	suffixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
