package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	count := make([]int, 26)

	for _, char := range text {
		count[char-'a']++
	}

	b := count['b'-'a']
	a := count['a'-'a']
	l := count['l'-'a'] / 2
	o := count['o'-'a'] / 2
	n := count['n'-'a']

	return min(b, min(a, min(l, min(o, n))))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}
