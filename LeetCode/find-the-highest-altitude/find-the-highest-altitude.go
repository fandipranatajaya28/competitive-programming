package main

import "fmt"

func largestAltitude(gain []int) int {
	var (
		highestAlt int
		currAlt    int
	)

	for _, pointGain := range gain {
		currAlt += pointGain
		if currAlt > highestAlt {
			highestAlt = currAlt
		}
	}

	return highestAlt
}

func main() {
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
