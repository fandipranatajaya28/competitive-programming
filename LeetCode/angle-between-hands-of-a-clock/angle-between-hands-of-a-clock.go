package main

import "fmt"

func angleClock(hour int, minutes int) float64 {
	// Convert hour to 12-hour position.
	// Example:
	// hour = 12 should behave like 0 degrees, not 360 degrees.
	hour = hour % 12

	// Hour hand:
	// - moves 30 degrees per hour
	// - also moves 0.5 degrees per minute
	//
	// Example:
	// 3:30
	// hour angle = 3*30 + 30*0.5 = 90 + 15 = 105
	hourAngle := float64(hour)*30.0 + float64(minutes)*0.5

	// Minute hand:
	// - moves 6 degrees per minute
	//
	// Example:
	// 30 minutes = 30*6 = 180 degrees
	minuteAngle := float64(minutes) * 6.0

	// Get absolute difference between both hand positions.
	diff := absFloat(hourAngle - minuteAngle)

	// There are always two angles between clock hands.
	// Example:
	// diff = 270 degrees
	// smaller angle = 360 - 270 = 90 degrees
	if diff > 180.0 {
		return 360.0 - diff
	}

	return diff
}

func absFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(angleClock(12, 30)) // 165
	fmt.Println(angleClock(3, 30))  // 75
	fmt.Println(angleClock(3, 15))  // 7.5
}
