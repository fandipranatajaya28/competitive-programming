package main

import "fmt"

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// Map: stop -> list of route indices that service this stop
	stopToRoutes := make(map[int][]int)
	for i, route := range routes {
		for _, stop := range route {
			stopToRoutes[stop] = append(stopToRoutes[stop], i)
		}
	}

	// BFS over routes
	visitedRoute := make([]bool, len(routes))
	q := []int{}    // queue of route indices
	dist := []int{} // parallel queue: buses taken so far to reach that route

	for _, r := range stopToRoutes[source] {
		visitedRoute[r] = true
		q = append(q, r)
		dist = append(dist, 1) // boarding first bus
	}

	// Standard BFS
	for len(q) > 0 {
		r := q[0]
		steps := dist[0]
		q, dist = q[1:], dist[1:]

		// If this route contains the target, we are done.
		for _, stop := range routes[r] {
			if stop == target {
				return steps
			}
		}

		// From every stop on this route, you can transfer to any other route serving that stop.
		for _, stop := range routes[r] {
			for _, nr := range stopToRoutes[stop] {
				if !visitedRoute[nr] {
					visitedRoute[nr] = true
					q = append(q, nr)
					dist = append(dist, steps+1)
				}
			}
			// Optional micro-opt: clear to save work later
			stopToRoutes[stop] = nil
		}
	}

	return -1
}

func main() {
	routes := [][]int{
		{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13},
	}
	source := 15
	target := 12

	fmt.Println(numBusesToDestination(routes, source, target)) // output -1

	routes = [][]int{
		{1, 2, 7}, {3, 6, 7},
	}
	source = 1
	target = 6

	fmt.Println(numBusesToDestination(routes, source, target)) // output 2
}
