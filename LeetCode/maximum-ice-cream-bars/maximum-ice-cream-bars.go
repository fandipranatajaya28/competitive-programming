package main

import "fmt"

func maxIceCream(costs []int, coins int) int {
	// Find the maximum cost so we only allocate the count array as large as needed.
	maxCost := 0
	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}

	// count[c] = how many ice cream bars have cost c.
	count := make([]int, maxCost+1)

	for _, cost := range costs {
		count[cost]++
	}

	bought := 0

	// Try to buy ice cream bars from cheapest to most expensive.
	for cost := 1; cost <= maxCost; cost++ {
		freq := count[cost]

		// No ice cream bar with this cost.
		if freq == 0 {
			continue
		}

		// If we can buy all bars with this cost, buy them all at once.
		totalCost := cost * freq

		if coins >= totalCost {
			coins -= totalCost
			bought += freq
		} else {
			// Otherwise, buy as many as possible with remaining coins.
			bought += coins / cost
			break
		}
	}

	return bought
}

func main() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7

	fmt.Println(maxIceCream(costs, coins)) // Output: 4
}
