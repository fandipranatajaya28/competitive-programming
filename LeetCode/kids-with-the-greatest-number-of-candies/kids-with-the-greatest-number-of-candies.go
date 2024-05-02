package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandy := 0
	result := []bool{}

	for _, candy := range candies {
		if candy > maxCandy {
			maxCandy = candy
		}
	}

	for _, candy := range candies {
		result = append(result, (candy+extraCandies) >= maxCandy)
	}

	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
