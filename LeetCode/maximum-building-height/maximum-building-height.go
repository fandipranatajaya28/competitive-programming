package main

import (
	"fmt"
	"sort"
)

func maxBuilding(n int, restrictions [][]int) int {
	// Building 1 must have height 0.
	// Add it as a restriction.
	restrictions = append(restrictions, []int{1, 0})

	// If building n has no explicit restriction, the maximum possible height
	// from building 1 to building n is n-1.
	//
	// Add building n as a restriction too, so we can calculate every segment
	// between adjacent restricted buildings.
	restrictions = append(restrictions, []int{n, n - 1})

	// Sort restrictions by building index.
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	// Left to right pass:
	// A building cannot be more than distance higher than the previous
	// restricted building.
	//
	// Example:
	// previous restriction: building 2 height <= 1
	// current building: 5
	// distance = 3
	// so current height cannot exceed 1 + 3 = 4
	for i := 1; i < len(restrictions); i++ {
		prevBuilding := restrictions[i-1][0]
		prevHeight := restrictions[i-1][1]

		currBuilding := restrictions[i][0]
		currHeight := restrictions[i][1]

		distance := currBuilding - prevBuilding

		restrictions[i][1] = min(currHeight, prevHeight+distance)
	}

	// Right to left pass:
	// Same idea, but now constraints from the right side also limit the left.
	for i := len(restrictions) - 2; i >= 0; i-- {
		nextBuilding := restrictions[i+1][0]
		nextHeight := restrictions[i+1][1]

		currBuilding := restrictions[i][0]
		currHeight := restrictions[i][1]

		distance := nextBuilding - currBuilding

		restrictions[i][1] = min(currHeight, nextHeight+distance)
	}

	answer := 0

	// Now every adjacent pair of restrictions is valid.
	// For each segment between two restricted buildings, calculate the highest
	// possible peak inside that segment.
	for i := 1; i < len(restrictions); i++ {
		leftBuilding := restrictions[i-1][0]
		leftHeight := restrictions[i-1][1]

		rightBuilding := restrictions[i][0]
		rightHeight := restrictions[i][1]

		distance := rightBuilding - leftBuilding

		// Suppose we move from left to right.
		// Height can increase by at most 1 each step from leftHeight.
		// Also, it must be able to decrease back to rightHeight.
		//
		// The highest peak between two endpoints is:
		//
		// (leftHeight + rightHeight + distance) / 2
		//
		// Example:
		// left height = 1, right height = 2, distance = 4
		//
		// possible shape:
		// 1, 2, 3, 3, 2
		//
		// peak = (1 + 2 + 4) / 2 = 3
		peak := (leftHeight + rightHeight + distance) / 2

		answer = max(answer, peak)
	}

	return answer
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
	n := 5
	restrictions := [][]int{
		{2, 1},
		{4, 1},
	}

	fmt.Println(maxBuilding(n, restrictions)) // Output: 2
}
