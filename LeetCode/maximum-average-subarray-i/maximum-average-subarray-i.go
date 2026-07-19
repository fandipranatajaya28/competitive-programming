package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	// Sum the first window explicitly, then seed the answer with it.
	tempTotal := 0
	for i := 0; i < k; i++ {
		tempTotal += nums[i]
	}
	answer := float64(tempTotal) / float64(k)

	// Slide: add the new element, drop the one leaving the window.
	for i := k; i < len(nums); i++ {
		tempTotal += nums[i] - nums[i-k]
		if avg := float64(tempTotal) / float64(k); avg > answer {
			answer = avg
		}
	}

	return answer
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
