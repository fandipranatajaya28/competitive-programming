package main

import (
	"fmt"
)

func totalWaviness(num1 int64, num2 int64) int64 {
	return countUpTo(num2) - countUpTo(num1-1)
}

type Result struct {
	count int64 // how many numbers can be formed from this state
	total int64 // total waviness from all numbers formed from this state
}

type State struct {
	pos      int
	started  bool
	lenState int

	// Last two digits.
	// Only meaningful when lenState == 2.
	prevPrev int
	prev     int
}

func countUpTo(n int64) int64 {
	if n <= 0 {
		return 0
	}

	digits := toDigits(n)
	memo := make(map[State]Result)

	var dfs func(pos int, tight bool, started bool, lenState int, prevPrev int, prev int) Result

	dfs = func(pos int, tight bool, started bool, lenState int, prevPrev int, prev int) Result {
		// Finished building one number.
		if pos == len(digits) {
			return Result{
				count: 1,
				total: 0,
			}
		}

		// Memoize only when tight == false.
		// If tight is true, the current prefix is constrained by n,
		// so the result depends on the upper bound.
		state := State{
			pos:      pos,
			started:  started,
			lenState: lenState,
			prevPrev: prevPrev,
			prev:     prev,
		}

		if !tight {
			if val, ok := memo[state]; ok {
				return val
			}
		}

		limit := 9
		if tight {
			limit = digits[pos]
		}

		res := Result{}

		for d := 0; d <= limit; d++ {
			nextTight := tight && d == limit

			// Skip leading zeroes.
			// Example: 000123 should be treated as 123,
			// not as a number with leading zero digits.
			if !started && d == 0 {
				child := dfs(pos+1, nextTight, false, 0, 0, 0)

				res.count += child.count
				res.total += child.total
				continue
			}

			addWave := int64(0)

			// If we already have at least two digits,
			// appending digit d lets us check whether prev is wavy:
			//
			// prevPrev, prev, d
			//
			// Example:
			// 1, 3, 2 => 3 is a peak
			// 3, 1, 2 => 1 is a valley
			if lenState == 2 {
				if (prev > prevPrev && prev > d) ||
					(prev < prevPrev && prev < d) {
					addWave = 1
				}
			}

			nextLenState := lenState
			nextPrevPrev := prevPrev
			nextPrev := prev

			if lenState == 0 {
				// First real digit.
				nextLenState = 1
				nextPrev = d
			} else if lenState == 1 {
				// Second real digit.
				nextLenState = 2
				nextPrevPrev = prev
				nextPrev = d
			} else {
				// Already have at least two digits.
				// Shift the last two digits.
				nextLenState = 2
				nextPrevPrev = prev
				nextPrev = d
			}

			child := dfs(pos+1, nextTight, true, nextLenState, nextPrevPrev, nextPrev)

			// addWave applies to every number generated from child.
			res.count += child.count
			res.total += child.total + addWave*child.count
		}

		if !tight {
			memo[state] = res
		}

		return res
	}

	return dfs(0, true, false, 0, 0, 0).total
}

func toDigits(n int64) []int {
	digits := []int{}

	for n > 0 {
		digits = append(digits, int(n%10))
		n /= 10
	}

	// Reverse to most-significant digit first.
	for l, r := 0, len(digits)-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}

	return digits
}

func main() {
	fmt.Println(totalWaviness(120, 130))
}
