package main

import "fmt"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	// Find the largest asteroid mass.
	// This lets us create a counting array only as large as needed.
	maxAsteroid := 0
	for _, asteroid := range asteroids {
		if asteroid > maxAsteroid {
			maxAsteroid = asteroid
		}
	}

	// count[x] stores how many asteroids have mass x.
	// This is a counting-sort style approach.
	count := make([]int, maxAsteroid+1)

	for _, asteroid := range asteroids {
		count[asteroid]++
	}

	// Use int64 because mass can grow after destroying asteroids.
	currMass := int64(mass)

	// Process asteroids from smallest mass to largest mass.
	// Greedy idea:
	// If we cannot destroy the smallest remaining asteroid,
	// then we cannot destroy any larger asteroid either.
	for asteroidMass := 1; asteroidMass <= maxAsteroid; asteroidMass++ {
		freq := count[asteroidMass]

		// No asteroid with this mass, skip it.
		if freq == 0 {
			continue
		}

		// If current mass is smaller than this asteroid mass,
		// we cannot destroy it, so the answer is false.
		if currMass < int64(asteroidMass) {
			return false
		}

		// Destroy all asteroids with the same mass.
		// Since currMass >= asteroidMass before destroying the first one,
		// after each destroyed asteroid, currMass only gets larger,
		// so all asteroids with this same mass can be destroyed safely.
		currMass += int64(asteroidMass * freq)
	}

	// All asteroids were destroyed successfully.
	return true
}

func main() {
	fmt.Println(asteroidsDestroyed(10, []int{3, 9, 19, 5, 21}))
}
