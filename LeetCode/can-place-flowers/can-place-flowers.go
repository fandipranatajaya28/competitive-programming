package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}

	if n == 0 {
		return true
	}

	remainingPlot := 0

	for idx := 0; idx < len(flowerbed); idx++ {
		isLeftPlotEmpty := true
		if idx > 0 && flowerbed[idx-1] == 1 {
			isLeftPlotEmpty = false
		}
		isRightPlotEmpty := true
		if idx < len(flowerbed)-1 && flowerbed[idx+1] == 1 {
			isRightPlotEmpty = false
		}
		if isLeftPlotEmpty && isRightPlotEmpty && flowerbed[idx] == 0 {
			remainingPlot++
			idx++
		}
	}

	return remainingPlot >= n
}

func main() {
	flowerbed := []int{0, 1, 0, 1, 0, 1, 0, 0}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
