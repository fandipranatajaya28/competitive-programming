package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	var (
		minEatingSpeed int
		highK          int
		lowK           = 1
	)

	for _, pile := range piles {
		if pile > highK {
			highK = pile
		}
	}

	for lowK <= highK {
		midK := lowK + (highK-lowK)/2
		if canFinishEating(piles, h, midK) {
			highK = midK - 1
			minEatingSpeed = midK
		} else {
			lowK = midK + 1
		}
	}

	return minEatingSpeed
}

func canFinishEating(piles []int, h int, k int) bool {
	var totalHours int
	for _, pile := range piles {
		pileEatenHours := math.Ceil(float64(pile) / float64(k))
		totalHours += int(pileEatenHours)
	}
	return totalHours <= h
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
