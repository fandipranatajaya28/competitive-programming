package main

import "fmt"

func maxArea(height []int) int {
	var (
		leftPtr  = 0
		rightPtr = len(height) - 1
		area     int
	)

	if len(height) < 2 {
		return area
	}

	for rightPtr > leftPtr {
		minHeight := height[leftPtr]
		if height[rightPtr] < minHeight {
			minHeight = height[rightPtr]
		}

		tempArea := minHeight * (rightPtr - leftPtr)
		if tempArea > area {
			area = tempArea
		}

		if height[leftPtr] < height[rightPtr] {
			leftPtr++
		} else {
			rightPtr--
		}
	}

	return area
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
