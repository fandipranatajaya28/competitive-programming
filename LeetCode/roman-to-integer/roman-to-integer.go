package main

import "fmt"

func romanToInt(s string) int {
	romanToIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var resNum int

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanToIntMap[s[i]] < romanToIntMap[s[i+1]] {
			resNum -= romanToIntMap[s[i]]
			continue
		}
		resNum += romanToIntMap[s[i]]
	}

	return resNum
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
