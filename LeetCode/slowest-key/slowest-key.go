package main

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	sKey := keysPressed[0]
	max := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if ((releaseTimes[i] - releaseTimes[i-1]) > max) || ((releaseTimes[i]-releaseTimes[i-1]) == max && keysPressed[i] > sKey) {
			max = releaseTimes[i] - releaseTimes[i-1]
			sKey = keysPressed[i]
		}
	}
	return sKey
}

func main() {
	releaseTimes := []int{12, 23, 36, 46, 62}
	keysPressed := "spuda"
	fmt.Println(string(slowestKey(releaseTimes, keysPressed)))
	return
}
