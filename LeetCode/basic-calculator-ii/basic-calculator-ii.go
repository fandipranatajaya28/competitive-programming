package main

import "fmt"

func calculate(s string) int {
	var result int
	lenStr := len(s)
	op := '+'
	num := 0
	prev := 0

	for i := 0; i < lenStr; i++ {
		currChar := s[i]
		if currChar >= '0' && currChar <= '9' {
			num = num*10 + int(currChar-'0')
		}
		if currChar == '+' || currChar == '-' || currChar == '*' || currChar == '/' || i == lenStr-1 {
			switch op {
			case '+':
				result += prev
				prev = num
			case '-':
				result += prev
				prev = -num
			case '*':
				prev *= num
			case '/':
				prev /= num
			}
			op = rune(currChar)
			num = 0
		}
	}

	return result + prev
}

func calculateStackVersion(s string) int {
	var result int
	lenStr := len(s)
	stack := []int{0}
	op := '+'
	num := 0

	for i := 0; i < lenStr; i++ {
		currChar := s[i]
		if currChar >= '0' && currChar <= '9' {
			num = num*10 + int(currChar-'0')
		}
		if currChar == '+' || currChar == '-' || currChar == '*' || currChar == '/' || i == lenStr-1 {
			lenStack := len(stack)
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[lenStack-1] = stack[lenStack-1] * num
			case '/':
				stack[lenStack-1] = stack[lenStack-1] / num
			}
			op = rune(currChar)
			num = 0
		}
	}

	for _, elem := range stack {
		result += elem
	}

	return result
}

func main() {
	fmt.Println(calculate(" 3+5 / 2 "))
}
