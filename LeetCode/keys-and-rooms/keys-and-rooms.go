package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	isVisitedMap := make(map[int]bool)
	visitRooms(0, rooms, isVisitedMap)
	for i := 0; i < len(rooms); i++ {
		if !isVisitedMap[i] {
			return false
		}
	}
	return true
}

func visitRooms(room int, rooms [][]int, isVisitedMap map[int]bool) {
	if isVisitedMap[room] {
		return
	}
	isVisitedMap[room] = true
	for _, key := range rooms[room] {
		visitRooms(key, rooms, isVisitedMap)
	}
}

func main() {
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}
