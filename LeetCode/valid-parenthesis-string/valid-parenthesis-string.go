package main

import "fmt"

func checkValidString(s string) bool {
	var (
		stack    []int
		wildcard []int
	)

	for idx, char := range s {
		if char == '(' {
			stack = append(stack, idx)
			continue
		} else if char == '*' {
			wildcard = append(wildcard, idx)
			continue
		}
		if len(stack) == 0 && len(wildcard) == 0 {
			return false
		}
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
			continue
		}
		wildcard = wildcard[:len(wildcard)-1]
	}

	index := 0

	for i := 0; i < len(stack); i++ {
		found := false
		for j := index; j < len(wildcard); j++ {
			if wildcard[j] > stack[i] {
				index = j + 1
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func main() {
	var s1 = "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"
	fmt.Println(checkValidString(s1))
	// true

	var s2 = "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"
	fmt.Println(checkValidString(s2))
	// false
}
