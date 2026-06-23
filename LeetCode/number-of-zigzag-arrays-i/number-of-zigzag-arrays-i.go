package main

import "fmt"

func zigZagArrays(n int, l int, r int) int {
	const MOD int = 1_000_000_007

	// Number of possible values in range [l, r].
	m := r - l + 1

	// If length is 1, every value from l to r is valid.
	if n == 1 {
		return m
	}

	// If only one value exists and n > 1,
	// we cannot create a valid array because adjacent values cannot be equal.
	if m == 1 {
		return 0
	}

	// up[i]:
	// number of arrays ending with value (l+i),
	// where the last movement is upward.
	//
	// Example:
	// ..., 3, 5
	// last move is up because 3 < 5.
	up := make([]int, m)

	// down[i]:
	// number of arrays ending with value (l+i),
	// where the last movement is downward.
	//
	// Example:
	// ..., 5, 3
	// last move is down because 5 > 3.
	down := make([]int, m)

	// Initialize arrays of length 2.
	//
	// For ending value x:
	// up[x] = number of previous values smaller than x
	// down[x] = number of previous values greater than x
	for i := 0; i < m; i++ {
		up[i] = i           // values before it are smaller
		down[i] = m - 1 - i // values after it are greater
	}

	// Build arrays from length 3 to n.
	for length := 3; length <= n; length++ {
		newUp := make([]int, m)
		newDown := make([]int, m)

		// prefixDown[i] = sum of down[0..i-1]
		// This helps calculate:
		// newUp[i] = sum of down[j] where j < i
		prefixDown := make([]int, m+1)

		// prefixUp[i] = sum of up[0..i-1]
		prefixUp := make([]int, m+1)

		for i := 0; i < m; i++ {
			prefixDown[i+1] = (prefixDown[i] + down[i]) % MOD
			prefixUp[i+1] = (prefixUp[i] + up[i]) % MOD
		}

		for i := 0; i < m; i++ {
			// To end with an upward move into value i,
			// previous value must be smaller than i,
			// and the previous move must be downward.
			//
			// Example:
			// 5, 2, 4
			// previous move: 5 -> 2 is down
			// current move:  2 -> 4 is up
			newUp[i] = prefixDown[i]

			// To end with a downward move into value i,
			// previous value must be greater than i,
			// and the previous move must be upward.
			//
			// Example:
			// 2, 5, 3
			// previous move: 2 -> 5 is up
			// current move:  5 -> 3 is down
			newDown[i] = (prefixUp[m] - prefixUp[i+1] + MOD) % MOD
		}

		up = newUp
		down = newDown
	}

	answer := 0

	// Sum all valid arrays ending with any value
	// and with either last direction.
	for i := 0; i < m; i++ {
		answer = (answer + up[i]) % MOD
		answer = (answer + down[i]) % MOD
	}

	return answer
}

func main() {
	fmt.Println(zigZagArrays(3, 1, 3))
}
