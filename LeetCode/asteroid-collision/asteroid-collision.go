package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	// stack stores asteroids that are still alive.
	stack := []int{}

	for _, asteroid := range asteroids {
		// alive tells whether the current asteroid survives
		// after possible collisions.
		alive := true

		// Collision only happens when:
		// 1. There is an asteroid in the stack.
		// 2. The top asteroid moves right: stack[len(stack)-1] > 0.
		// 3. Current asteroid moves left: asteroid < 0.
		//
		// Example:
		// stack top = 10, asteroid = -5
		// They move toward each other, so they collide.
		for alive && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]

			if top < -asteroid {
				// Top asteroid is smaller, so it explodes.
				// Current asteroid may continue colliding with previous asteroids.
				stack = stack[:len(stack)-1]
			} else if top == -asteroid {
				// Same size, both explode.
				stack = stack[:len(stack)-1]
				alive = false
			} else {
				// Top asteroid is larger, current asteroid explodes.
				alive = false
			}
		}

		// If current asteroid survives all collisions,
		// add it to the stack.
		if alive {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{3, 5, -6, 2, -1, 4}))
}
