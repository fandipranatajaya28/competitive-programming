package main

import "fmt"

func earliestFinishTime(
	landStartTime []int,
	landDuration []int,
	waterStartTime []int,
	waterDuration []int,
) int {
	// Try both possible orders:
	// 1. land ride first, then water ride
	// 2. water ride first, then land ride
	landThenWater := calc(landStartTime, landDuration, waterStartTime, waterDuration)
	waterThenLand := calc(waterStartTime, waterDuration, landStartTime, landDuration)

	return min(landThenWater, waterThenLand)
}

// calc returns the earliest finish time if we must take:
// one ride from the first category first,
// then one ride from the second category.
func calc(
	firstStart []int,
	firstDuration []int,
	secondStart []int,
	secondDuration []int,
) int {
	// Step 1:
	// Find the earliest possible finish time among all first rides.
	//
	// If we choose ride i as the first ride:
	// finish time = firstStart[i] + firstDuration[i]
	//
	// We only need the minimum finish time.
	earliestFirstFinish := int(1e9)

	for i := 0; i < len(firstStart); i++ {
		finish := firstStart[i] + firstDuration[i]
		earliestFirstFinish = min(earliestFirstFinish, finish)
	}

	// Step 2:
	// Try every second ride.
	//
	// If the second ride already opened before earliestFirstFinish,
	// we can start it immediately at earliestFirstFinish.
	//
	// Otherwise, we must wait until secondStart[i].
	ans := int(1e9)

	for i := 0; i < len(secondStart); i++ {
		startSecond := max(earliestFirstFinish, secondStart[i])
		finishSecond := startSecond + secondDuration[i]

		ans = min(ans, finishSecond)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(earliestFinishTime([]int{2, 8}, []int{4, 1}, []int{6}, []int{3}))
}
