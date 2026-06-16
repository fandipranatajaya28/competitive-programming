package main

import "fmt"

func processStr(s string) string {
	// Use a byte slice as a mutable string builder.
	// This is more efficient than repeatedly doing result += string(c).
	result := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'a' && ch <= 'z' {
			// Normal lowercase letter:
			// append it to the current result.
			result = append(result, ch)

		} else if ch == '*' {
			// '*' operation:
			// remove the last character if result is not empty.
			if len(result) > 0 {
				result = result[:len(result)-1]
			}

		} else if ch == '#' {
			// '#' operation:
			// duplicate the current result and append it to itself.
			//
			// Example:
			// result = "ab"
			// after '#': "abab"
			result = append(result, result...)

		} else if ch == '%' {
			// '%' operation:
			// reverse the current result.
			for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
				result[l], result[r] = result[r], result[l]
			}
		}
	}

	return string(result)
}

func main() {
	s := "a#b%*"

	fmt.Println(processStr(s)) // Output: "ba"
}
