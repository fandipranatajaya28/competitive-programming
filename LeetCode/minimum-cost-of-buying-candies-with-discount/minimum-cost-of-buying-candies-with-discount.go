package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	var minCost int

	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	for idx, candyCost := range cost {
		if idx%3 == 2 {
			continue
		}
		minCost += candyCost
	}

	return minCost
}

func main() {
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
}
