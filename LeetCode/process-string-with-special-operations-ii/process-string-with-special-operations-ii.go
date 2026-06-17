package main

import "fmt"

func processStr(s string, k int64) byte {
	// The problem specifically asks to create this variable.
	tibrelkano := s

	n := len(tibrelkano)

	// lengths[i] = length of result after processing first i characters.
	//
	// lengths[0] = 0, before processing anything.
	// lengths[i+1] = length after processing tibrelkano[i].
	lengths := make([]int64, n+1)

	for i := 0; i < n; i++ {
		ch := tibrelkano[i]
		prevLen := lengths[i]

		if ch >= 'a' && ch <= 'z' {
			// Letter operation:
			// append one character.
			lengths[i+1] = prevLen + 1

		} else if ch == '*' {
			// '*' operation:
			// delete last character if result is not empty.
			if prevLen > 0 {
				lengths[i+1] = prevLen - 1
			} else {
				lengths[i+1] = 0
			}

		} else if ch == '#' {
			// '#' operation:
			// duplicate current result.
			lengths[i+1] = prevLen * 2

		} else if ch == '%' {
			// '%' operation:
			// reverse current result.
			// Length does not change.
			lengths[i+1] = prevLen
		}
	}

	finalLen := lengths[n]

	// If k is outside the final string, return '.'.
	if k >= finalLen {
		return '.'
	}

	// Work backward from the final result.
	//
	// Instead of reconstructing the string, we reverse the effect
	// of each operation and update k to the corresponding index
	// before that operation.
	for i := n - 1; i >= 0; i-- {
		ch := tibrelkano[i]

		beforeLen := lengths[i]
		afterLen := lengths[i+1]

		if ch >= 'a' && ch <= 'z' {
			// Forward:
			// previousResult + ch
			//
			// If k points to the newly appended character,
			// then this is the answer.
			if k == afterLen-1 {
				return ch
			}

			// Otherwise, k points to an older character,
			// so index stays the same when moving backward.

		} else if ch == '*' {
			// Forward:
			// delete last char if exists.
			//
			// The final string after '*' contains only characters
			// that already existed before '*', so k stays the same
			// when moving backward.
			_ = beforeLen

		} else if ch == '#' {
			// Forward:
			// result = result + result
			//
			// If k is in the second copy, map it back to the first copy.
			// Example:
			// before = "abc", after = "abcabc"
			// k = 4 maps to 1.
			if beforeLen > 0 {
				k %= beforeLen
			}

		} else if ch == '%' {
			// Forward:
			// reverse result.
			//
			// If length is L, index k after reverse came from:
			// L - 1 - k
			k = afterLen - 1 - k
		}
	}

	return '.'
}

func main() {
	fmt.Println(string(processStr("a#b%*", 1)))  // Output: "a"
	fmt.Println(string(processStr("cd%#*#", 3))) // Output: "d"
	fmt.Println(string(processStr("z*#", 0)))    // Output: "."
}
