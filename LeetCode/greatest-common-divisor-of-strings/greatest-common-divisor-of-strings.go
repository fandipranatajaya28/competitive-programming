package main

import "fmt"

func gcdOfStrings(s1 string, s2 string) string {
	if s1+s2 != s2+s1 {
		return ""
	}
	x := gcd(len(s1), len(s2))
	return s1[:x]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	a := "ABABABAB"
	b := "ABAB"
	fmt.Println(gcdOfStrings(a, b))
}
