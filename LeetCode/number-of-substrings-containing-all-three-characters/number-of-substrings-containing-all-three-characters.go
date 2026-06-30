package main

// numberOfSubstrings counts substrings that contain at least one 'a', 'b', and 'c'.
//
// Key idea: fix the right end of the substring. Track the last position where
// each of 'a', 'b', 'c' appeared. If all three have appeared, then any starting
// index from 0 up to min(lastA, lastB, lastC) produces a valid substring ending
// at the current right index. That gives min(...)+1 valid substrings for this right end.
//
// Time: O(n), Space: O(1).
func numberOfSubstrings(s string) int {
	// last[c] = most recent index where character ('a'+c) was seen; -1 if never.
	last := [3]int{-1, -1, -1}
	count := 0

	for i := 0; i < len(s); i++ {
		// Record the current index for this character.
		last[s[i]-'a'] = i

		// The earliest of the three last-seen positions limits how far left
		// the start can go while still covering all three characters.
		minPos := last[0]
		if last[1] < minPos {
			minPos = last[1]
		}
		if last[2] < minPos {
			minPos = last[2]
		}

		// If any character hasn't appeared yet, minPos == -1 and we add 0.
		// Otherwise, starts 0..minPos are all valid → minPos+1 substrings.
		count += minPos + 1
	}

	return count
}

func main() {
	// Example test cases from the problem.
	println(numberOfSubstrings("abcabc")) // 10
	println(numberOfSubstrings("aaacb"))  // 3
	println(numberOfSubstrings("abc"))    // 1
}
