package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0

	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfit = max(maxProfit, price-minPrice)
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
